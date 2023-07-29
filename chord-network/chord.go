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

	// tracer     *Tracer // for testing latency and hops
	tracerRWMu sync.RWMutex

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