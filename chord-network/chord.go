package main

import (
	"context"
	"errors"
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
	
	pred, err := c.FFindPredecessor(ctx, id)
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

// Rransfer keys from current node to target
func (c *Chord) TransferKeys(ctx context.Context, target *proto.Node, start []byte, end []byte) (int, error) {
	
	c.storeLock.Lock()
	defer c.storeLock.Unlock()

	count := 0
	
	for _, key := range c.store.GetKeys(ctx) {
		
		hashedKey := c.hash(key)
		
		if betweenRightInclusive(hashedKey, start, end) {

			err := c.put(ctx, target.Ip, key)
			if err != nil {
				return count, err
			}

			count++

			c.store.Delete(ctx, key)
		}
	}

	return count, nil
}

func (c *Chord) Notify(ctx context.Context, potentialPredecessor *proto.Node) error {
	
	c.predecessorLock.Lock()
	defer c.predecessorLock.Unlock()

	var prevPredecessor *proto.Node
	
	if c.predecessor == nil || between(potentialPredecessor.Id, c.predecessor.Id, c.Id) {
		
		if c.predecessor != nil {
			prevPredecessor = c.predecessor
		}

		c.predecessor = potentialPredecessor
		
		if prevPredecessor != nil {
			
			if between(c.predecessor.Id, prevPredecessor.Id, c.Id) {
				c.TransferKeys(ctx, c.predecessor, prevPredecessor.Id, c.predecessor.Id) // transfer our key to new predecessor
			}
		}
	}

	return nil
}

// Periodically verify node's immediate successor
func (c *Chord) Stabilize(ctx context.Context) error {
	
	successor := c.GetSuccessor()
	
	x, err := c.getPredecessor(ctx, successor)

	if x == nil || err != nil {
		c.logger.Println("RPC fails or the successor node died")
		return err
	}

	// the pred of our succ is nil, it hasn't updated it pred, still notify
	if x.Id == nil {
		_, err = c.notify(ctx, c.GetSuccessor(), c.Node)
		return err
	}

	// found new successor
	if between(x.Id, c.Id, successor.Id) {
		c.fingerLock.Lock()
		c.fingerTable[0] = x
		c.fingerLock.Unlock()
	}

	_, err = c.notify(ctx, c.GetSuccessor(), c.Node)
	
	return err
}

func (c *Chord) Join(ctx context.Context, joinNode *proto.Node) error {
	
	c.fingerLock.Lock()
	c.fingerLock.Unlock()

	c.predecessor = nil
	
	if joinNode.Ip == "" { // first node in the ring
		c.fingerTable[0] = c.Node
		return nil
	}

	successor, err := c.findSuccessor(ctx, joinNode, c.Id)
	if err != nil {
		return err 
	}

	if idsEqual(successor.Id, c.Id) {
		return errors.New("Node with same ID already exists in the ring")
	}
	
	c.fingerTable[0] = successor

	return nil
}
