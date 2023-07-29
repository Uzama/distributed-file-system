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

func (c *Chord) _notify(ctx context.Context, remote *proto.Node, potentialPredecessor *proto.Node) (*proto.NN, error) {
	
	client, err := c.connectToRemote(remote.Ip)
	if err != nil {
		return nil, err
	}

	result, err := client.Notify(ctx, potentialPredecessor)

	return result, err
}

func (c *Chord) _findSuccessor(ctx context.Context, remote *proto.Node, id []byte) (*proto.Node, error) {
	
	client, err := c.connectToRemote(remote.Ip)
	if err != nil {
		return nil, err
	}

	result, err := client.FindSuccessor(ctx, &proto.ID{Id: id})

	return result, err
}

// Returns closest node based on ID
func (c *Chord) _findClosestPrecedingNode(ctx context.Context, remote *proto.Node, id []byte) (*proto.Node, error) {
	
	client, err := c.connectToRemote(remote.Ip)
	if err != nil {
		return nil, err
	}

	result, err := client.FindClosestPrecedingNode(ctx, &proto.ID{Id: id})

	return result, err
}

func (c *Chord) _getSuccessor(ctx context.Context, remote *proto.Node) (*proto.Node, error) {
	
	client, err := c.connectToRemote(remote.Ip)
	if err != nil {
		return nil, err
	}

	result, err := client.GetSuccessor(ctx, &proto.NN{})

	return result, err
}

func (c *Chord) _getPredecessor(ctx context.Context, remote *proto.Node) (*proto.Node, error) {
	
	client, err := c.connectToRemote(remote.Ip)
	if err != nil {
		return nil, err
	}

	result, err := client.GetPredecessor(ctx, &proto.NN{})

	return result, err
}

func (c *Chord) _setPredecessor(ctx context.Context, remote *proto.Node, pred *proto.Node) (*proto.NN, error) {
	
	client, err := c.connectToRemote(remote.Ip)
	if err != nil {
		return nil, err
	}

	result, err := client.SetPredecessor(ctx, pred)

	return result, err
}

func (c *Chord) _setSuccessor(ctx context.Context, remote *proto.Node, succ *proto.Node) (*proto.NN, error) {
	
	client, err := c.connectToRemote(remote.Ip)
	if err != nil {
		return nil, err
	}

	result, err := client.SetSuccessor(ctx, succ)

	return result, err
}

func (c *Chord) _get(ctx context.Context, remoteIP string, key string) (string, error) {
	
	client, err := c.connectToRemote(remoteIP)
	if err != nil {
		return "", err
	}

	request := &proto.GetRequest{Key: key}
	
	result, err := client.Get(ctx, request)
	
	return result.Key, err
}

func (c *Chord) _put(ctx context.Context, remoteIP string, key string) error {
	
	client, err := c.connectToRemote(remoteIP)
	if err != nil {
		return err
	}

	request := &proto.PutRequest{Key: key}

	_, err = client.Put(ctx, request)

	return err
}