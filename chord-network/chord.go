package main

import (
	"context"
	"log"
	"math/big"
	"sync"

	"chord-network/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/benchmark/latency"
)

type Chord struct {
	*proto.Node 

	predecessor *proto.Node
	predecessorLock sync.RWMutex

	fingerTable []*proto.Node
	fingerLock sync.RWMutex

	successorList []*proto.Node // keep track of log(n) nearest successor for recovery
	succLock sync.RWMutex

	stopChan chan struct{} 

	store     Store
	storeLock sync.RWMutex

	grpcServer          *grpc.Server         // this provides services
	connectionsPool     map[string]*GRPCConn // this takes care of all connections
	connLock sync.RWMutex

	tracer *Tracer
	tracerLock sync.RWMutex

	config *Config

	logger *log.Logger
	network *latency.Network
}

// Returns hashed values in []byte
func (c *Chord) hash(IP string) []byte {
	
	h := c.config.Hash()
	
	h.Write([]byte(IP))

	idInt := big.NewInt(0)
	idInt.SetBytes(h.Sum(nil)) // Sum() returns []byte, convert it into BigInt

	maxVal := big.NewInt(0)
	maxVal.Exp(big.NewInt(2), big.NewInt(int64(c.config.ringSize)), nil) // calculate 2^m
	
	idInt.Mod(idInt, maxVal)  // mod id to make it to be [0, 2^m - 1]
	if idInt.Cmp(big.NewInt(0)) == 0 {
		return []byte{0}
	}

	return idInt.Bytes()
}

func (c *Chord) Get(ctx context.Context, key string) (string, error) {
	
	c.storeLock.RLock()
	defer c.storeLock.RLock()
	
	_, ok := c.store.Get(ctx, key)
	if ok != nil {
		return "", ok
	}
	return key, nil
}

func (c *Chord) Put(ctx context.Context, key string) {
	
	c.storeLock.RLock()
	defer c.storeLock.RLock()

	c.store.Put(ctx, key)
}

func (c *Chord) GetSuccessor() *proto.Node {
	
	c.fingerLock.RLock()
	defer c.fingerLock.RUnlock()

	return c.fingerTable[0]
}

func (c *Chord) GetPredecessor() *proto.Node {
	
	c.predecessorLock.RLock()
	defer c.predecessorLock.RUnlock()

	return c.predecessor
}

// Returns the closest finger based on fingerTablex
func (c *Chord) FindClosestPrecedingNode(id []byte) *proto.Node {
	
	c.fingerLock.RLock()
	defer c.fingerLock.RUnlock()

	for i := c.config.ringSize - 1; i >= 0; i-- {
		
		if c.fingerTable[i] != nil {
			
			if between(c.fingerTable[i].Id, c.Id, id) {
				
				return c.fingerTable[i]
			}
		}
	}

	return c.Node
}

func (c *Chord) FindPredecessor(ctx context.Context, id []byte) (*proto.Node, error) {
	
	closest := c.FindClosestPrecedingNode(id)
	
	if idsEqual(closest.Id, c.Id) {
		return closest, nil
	}

	closestSucc, err := c.getSuccessor(ctx, closest) // get closest successor
	if err != nil {
		return nil, err
	}

	c.tracer.traceNode(closest.Id)

	for !betweenRightInclusive(id, closest.Id, closestSucc.Id) {
		
		closest, err := c.findClosestPrecedingNode(ctx, closest, id)
		if err != nil {
			return nil, err
		}

		closestSucc, err = c.getSuccessor(ctx, closest) // get closest successor
		if err != nil {
			return nil, err
		}

		c.tracer.traceNode(closest.Id)
	}

	return closest, nil
}

// For fix finger tables
func (c *Chord) FFindPredecessor(ctx context.Context, id []byte) (*proto.Node, error) {
	
	closest := c.FindClosestPrecedingNode(id)
	
	if idsEqual(closest.Id, c.Id) {
		return closest, nil
	}

	closestSucc, err := c.getSuccessor(ctx, closest) // get closest successor
	if err != nil {
		return nil, err
	}

	for !betweenRightInclusive(id, closest.Id, closestSucc.Id) {
		
		closest, err := c.findClosestPrecedingNode(ctx, closest, id)
		if err != nil {
			return nil, err
		}

		closestSucc, err = c.getSuccessor(ctx, closest) // get closest successor
		if err != nil {
			return nil, err
		}
	}

	return closest, nil
}

func (c *Chord) FindSuccessor(ctx context.Context, id []byte) (*proto.Node, error) {
	
	c.tracerLock.Lock()
	defer c.tracerLock.Unlock()

	c.tracer.startTracer(c.Id, id)
	
	pred, err := c.FindPredecessor(ctx, id)
	if err != nil {
		return nil, err
	}

	successor, err := c.getSuccessor(ctx, pred)
	if err != nil {
		return nil, err
	}

	c.tracer.traceNode(successor.Id)
	c.tracer.endTracer(successor.Id)

	return successor, nil
}

// For fix finger tables
func (c *Chord) FFindSuccessor(ctx context.Context, id []byte) (*proto.Node, error) {

	c.tracer.startTracer(c.Id, id)
	
	pred, err := c.FindPredecessor(ctx, id)
	if err != nil {
		return nil, err
	}

	successor, err := c.getSuccessor(ctx, pred)
	if err != nil {
		return nil, err
	}

	return successor, nil
}

func (c *Chord) fingerStart(i int) []byte {
	
	currID := new(big.Int).SetBytes(c.Id)

	offset := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(i)), nil)

	max := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64(c.config.ringSize)), nil)

	start := new(big.Int).Add(currID, offset)
	
	start.Mod(start, max)

	if len(start.Bytes()) == 0 {
		return []byte{0}
	}

	return start.Bytes()
}

// Periodically refresh finger table 
func (c *Chord) FixFingers(ctx context.Context, i int) (int, error) {
	
	i = (i + 1) % c.config.ringSize
	
	fingerStart := c.fingerStart(i)
	
	finger, err := c.FFindSuccessor(ctx, fingerStart)
	if err != nil {
		return 0, err
	}

	c.fingerLock.Lock()
	c.fingerTable[i] = finger
	c.fingerLock.Unlock()

	return i, nil
}
