package main

import (
	"chord-network/proto"
	"context"
	"time"

	"google.golang.org/grpc"
)

type GRPCConn struct {
	IP     string
	Client proto.CommunicationClient 
	Conn   *grpc.ClientConn
}

// Returns a grpc.ClientConn
func Dial(ip string, options ...grpc.DialOption) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return grpc.DialContext(ctx, ip, options...)
}