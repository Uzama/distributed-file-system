package main

import (
	"context"
	"errors"
	"log"
	"math/big"
	"net"
	"os"
	"sync"
	"time"

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

func newChord(config *Config, ip string, joinNode string) (*Chord, error) {
	
	ctx := context.Background()

	chord := &Chord{
		Node:   new(proto.Node),
		config: config,
	}

	chord.Ip = ip
	chord.Id = chord.hash(ip)

	chord.fingerTable = make([]*proto.Node, chord.config.ringSize)
	chord.stopChan = make(chan struct{})
	chord.connectionsPool = make(map[string]*GRPCConn)
	chord.store = NewStore()
	chord.tracer = MakeTracer()
	chord.logger = log.New(os.Stderr, "logger: ", log.Ltime|log.Lshortfile) 

	chord.network = &latency.Network{Kbps: 100 * 1024, Latency: 2 * time.Millisecond, MTU: 1500}
	l, err := net.Listen("tcp", ip)
	if err != nil {
		chord.logger.Println(err)
		return nil, err
	}
	
	// listener := chord.network.Listener(l) // listener with latency injected

	chord.grpcServer = grpc.NewServer()

	proto.RegisterCommunicationServer(chord.grpcServer, chord)

	err = chord.join(ctx, &proto.Node{Ip: joinNode})

	if err != nil {
		chord.logger.Println(err)
		return nil, err
	}

	// info := benchmark.ServerInfo{Type: "protobuf", Listener: l}

	// benchmark.StartServer(info)

	go chord.grpcServer.Serve(l)

	go func() {
		
		ticker := time.NewTicker(chord.config.stabilizeTime)
		
		for {
			select {
			case <-ticker.C:
				err := chord.stabilize(ctx)
				if err != nil {
					chord.logger.Println(err)
				}
			case <-chord.stopChan:
				ticker.Stop()
				chord.logger.Printf("%s %d stopping stabilize\n", chord.Ip, chord.Id)
				return
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(chord.config.fixFingerTime)
		i := 0
		for {
			select {
			case <-ticker.C:
				i, err = chord.fixFingers(ctx, i)
				if err != nil {
					chord.logger.Println(err)
				}
			case <-chord.stopChan:
				ticker.Stop()
				chord.logger.Printf("%s %d stopping fixing fingers\n", chord.Ip, chord.Id)
				return
			}
		}
	}()

	return chord, nil
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

func (c *Chord) get(ctx context.Context, key string) (string, error) {
	
	c.storeLock.RLock()
	defer c.storeLock.RLock()
	
	_, ok := c.store.Get(ctx, key)
	if ok != nil {
		return "", ok
	}
	return key, nil
}

func (c *Chord) put(ctx context.Context, key string) {
	
	c.storeLock.RLock()
	defer c.storeLock.RLock()

	c.store.Put(ctx, key)
}

// Takes in a key, returns the ip address of the node that should store the key
func (c *Chord) lookup(ctx context.Context, key string) (string, error) {
	
	keyHased := c.hash(key)
	
	successor, err := c.findSuccessor(ctx, keyHased) 
	if err != nil {                           
		c.logger.Println(err)
		return "", err
	}

	return successor.Ip, nil
}

func (c *Chord) getSuccessor() *proto.Node {
	
	c.fingerLock.RLock()
	defer c.fingerLock.RUnlock()

	return c.fingerTable[0]
}

func (c *Chord) getPredecessor() *proto.Node {
	
	c.predecessorLock.RLock()
	defer c.predecessorLock.RUnlock()

	return c.predecessor
}

// Returns the closest finger based on fingerTablex
func (c *Chord) findClosestPrecedingNode(ctx context.Context, id []byte) *proto.Node {
	
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

func (c *Chord) findPredecessor(ctx context.Context, id []byte) (*proto.Node, error) {
	
	closest := c.findClosestPrecedingNode(ctx, id)
	
	if idsEqual(closest.Id, c.Id) {
		return closest, nil
	}

	closestSucc, err := c._getSuccessor(ctx, closest) // get closest successor
	if err != nil {
		return nil, err
	}

	c.tracer.traceNode(closest.Id)

	for !betweenRightInclusive(id, closest.Id, closestSucc.Id) {
		
		closest, err := c._findClosestPrecedingNode(ctx, closest, id)
		if err != nil {
			return nil, err
		}

		closestSucc, err = c._getSuccessor(ctx, closest) // get closest successor
		if err != nil {
			return nil, err
		}

		c.tracer.traceNode(closest.Id)
	}

	return closest, nil
}

// For fix finger tables
func (c *Chord) ffindPredecessor(ctx context.Context, id []byte) (*proto.Node, error) {
	
	closest := c.findClosestPrecedingNode(ctx, id)
	
	if idsEqual(closest.Id, c.Id) {
		return closest, nil
	}

	closestSucc, err := c._getSuccessor(ctx, closest) // get closest successor
	if err != nil {
		return nil, err
	}

	for !betweenRightInclusive(id, closest.Id, closestSucc.Id) {
		
		closest, err := c._findClosestPrecedingNode(ctx, closest, id)
		if err != nil {
			return nil, err
		}

		closestSucc, err = c._getSuccessor(ctx, closest) // get closest successor
		if err != nil {
			return nil, err
		}
	}

	return closest, nil
}

func (c *Chord) findSuccessor(ctx context.Context, id []byte) (*proto.Node, error) {
	
	c.tracerLock.Lock()
	defer c.tracerLock.Unlock()

	c.tracer.startTracer(c.Id, id)
	
	pred, err := c.findPredecessor(ctx, id)
	if err != nil {
		return nil, err
	}

	successor, err := c._getSuccessor(ctx, pred)
	if err != nil {
		return nil, err
	}

	c.tracer.traceNode(successor.Id)
	c.tracer.endTracer(successor.Id)

	return successor, nil
}

// For fix finger tables
func (c *Chord) ffindSuccessor(ctx context.Context, id []byte) (*proto.Node, error) {

	c.tracer.startTracer(c.Id, id)
	
	pred, err := c.ffindPredecessor(ctx, id)
	if err != nil {
		return nil, err
	}

	successor, err := c._getSuccessor(ctx, pred)
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
func (c *Chord) fixFingers(ctx context.Context, i int) (int, error) {
	
	i = (i + 1) % c.config.ringSize
	
	fingerStart := c.fingerStart(i)
	
	finger, err := c.ffindSuccessor(ctx, fingerStart)
	if err != nil {
		return 0, err
	}

	c.fingerLock.Lock()
	c.fingerTable[i] = finger
	c.fingerLock.Unlock()

	return i, nil
}

func (c *Chord) StopFixFingers() {
	close(c.stopChan)
}

// Rransfer keys from current node to target
func (c *Chord) transferKeys(ctx context.Context, target *proto.Node, start []byte, end []byte) (int, error) {
	
	c.storeLock.Lock()
	defer c.storeLock.Unlock()

	count := 0
	
	for _, key := range c.store.GetKeys(ctx) {
		
		hashedKey := c.hash(key)
		
		if betweenRightInclusive(hashedKey, start, end) {

			err := c._put(ctx, target.Ip, key)
			if err != nil {
				return count, err
			}

			count++

			c.store.Delete(ctx, key)
		}
	}

	return count, nil
}

func (c *Chord) notify(ctx context.Context, potentialPredecessor *proto.Node) error {
	
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
				c.transferKeys(ctx, c.predecessor, prevPredecessor.Id, c.predecessor.Id) // transfer our key to new predecessor
			}
		}
	}

	return nil
}

// Periodically verify node's immediate successor
func (c *Chord) stabilize(ctx context.Context) error {
	
	successor := c.getSuccessor()
	
	x, err := c._getPredecessor(ctx, successor)

	if x == nil || err != nil {
		c.logger.Println("RPC fails or the successor node died")
		return err
	}

	// the pred of our succ is nil, it hasn't updated it pred, still notify
	if x.Id == nil {
		_, err = c._notify(ctx, c.getSuccessor(), c.Node)
		return err
	}

	// found new successor
	if between(x.Id, c.Id, successor.Id) {
		c.fingerLock.Lock()
		c.fingerTable[0] = x
		c.fingerLock.Unlock()
	}

	_, err = c._notify(ctx, c.getSuccessor(), c.Node)
	
	return err
}

func (c *Chord) join(ctx context.Context, joinNode *proto.Node) error {
	
	c.fingerLock.Lock()
	c.fingerLock.Unlock()

	c.predecessor = nil
	
	if joinNode.Ip == "" { // first node in the ring
		c.fingerTable[0] = c.Node
		return nil
	}

	successor, err := c._findSuccessor(ctx, joinNode, c.Id)
	if err != nil {
		return err 
	}

	if idsEqual(successor.Id, c.Id) {
		return errors.New("Node with same ID already exists in the ring")
	}
	
	c.fingerTable[0] = successor

	return nil
}

func (c *Chord) leave(ctx context.Context) {
	
	c.logger.Printf("%s leaving the ring\n", c.Id)

	c.StopFixFingers()
	
	c.grpcServer.GracefulStop()

	successor := c.getSuccessor()
	pred := c.getPredecessor()

	if !idsEqual(successor.Id, c.Id) && pred != nil {
		
		count, _ := c.transferKeys(ctx, successor, pred.Id, c.Id)
		c.logger.Printf("number of transfered keys: %d\n", count)

		c._setPredecessor(ctx, successor, pred)
		c._setSuccessor(ctx, pred, successor)
	}

	c.stop()
}

// Gracefully stops the instance
func (c *Chord) stop() {
	
	c.logger.Printf("%s stopping outgoing connections\n", c.Ip)
	
	c.connLock.Lock()
	defer c.connLock.Unlock()

	for _, connection := range c.connectionsPool {
		
		err := connection.Conn.Close()
		if err != nil {
			c.logger.Println(err)
		}
	}
}


