package main

import (
	"log"
	"sync"

	"chord-network/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/benchmark/latency"
)

type GRPCConn struct {
	IP     string
	Client proto.CommunicationClient 
	Conn   *grpc.ClientConn
}

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

	// config *Config

	logger *log.Logger
	network *latency.Network
}