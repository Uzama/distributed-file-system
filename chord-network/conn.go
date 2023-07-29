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

func (c *Chord) connectToRemote(remoteIP string) (proto.CommunicationClient, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()

	grpc, ok := c.connectionsPool[remoteIP]
	if ok {
		c.connLock.RUnlock()
		return grpc.Client, nil
	}

	c.logger.Printf("connect-to-temote: dialing to %s\n", remoteIP)

	conn, err := Dial(remoteIP, c.config.DialOptions...)
	if err != nil {
		c.logger.Printf("connect-to-temote: dialing failed to %s\n", err.Error())
		return nil, err
	}

	c.logger.Println("connect-to-temote: dialed success")

	client := proto.NewCommunicationClient(conn)
	grpc = &GRPCConn{remoteIP, client, conn}

	c.connectionsPool[remoteIP] = grpc

	return client, nil
}

func (c *Chord) notify(remote *proto.Node, potentialPredecessor *proto.Node) (*proto.NN, error) {
	client, err := c.connectToRemote(remote.Ip)
	if err != nil {
		return nil, err
	}

	result, err := client.Notify(context.Background(), potentialPredecessor)
	return result, err
}