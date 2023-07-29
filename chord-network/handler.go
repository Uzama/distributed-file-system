package main

import (
	"chord-network/proto"
	"context"
)

// Notifies the given node that it might be the predecessor
func (c *Chord) Notify(ctx context.Context, potentialPred *proto.Node) (*proto.NN, error) {
	
	c.notify(ctx, potentialPred)

	return &proto.NN{}, nil
}

// Finds the successor of id
func (c *Chord) FindSuccessor(ctx context.Context, id *proto.ID) (*proto.Node, error) {
	
	succ, err := c.findSuccessor(ctx, id.Id)
	if err != nil {
		return nil, err
	}

	return succ, nil
}

// Finds the closest finger with given id
func (c *Chord) FindClosestPrecedingNode(ctx context.Context, id *proto.ID) (*proto.Node, error) {
	
	closestPrecedingNode := c.findClosestPrecedingNode(ctx, id.Id)
	
	return closestPrecedingNode, nil
}

// Gets the successor of the node
func (c *Chord) GetSuccessor(context.Context, *proto.NN) (*proto.Node, error) {
	
	succ := c.getSuccessor()
	if succ == nil {
		return &proto.Node{}, nil
	}

	return succ, nil
}

// sets successor for chord
func (c *Chord) SetSuccessor(ctx context.Context, succ *proto.Node) (*proto.NN, error) {
	
	c.fingerLock.Lock()
	defer c.fingerLock.Unlock()

	c.fingerTable[0] = succ

	return &proto.NN{}, nil
}

// Gets the predecessor of the node
func (c *Chord) GetPredecessor(context.Context, *proto.NN) (*proto.Node, error) {
	
	pred := c.getPredecessor()

	if pred == nil {
		return &proto.Node{}, nil
	}

	return pred, nil
}

// Sets predecessor for chord
func (c *Chord) SetPredecessor(ctx context.Context, pred *proto.Node) (*proto.NN, error) {
	
	c.predecessorLock.Lock()
	defer c.predecessorLock.Unlock()

	c.predecessor = pred

	return &proto.NN{}, nil
}

func (c *Chord) Get(ctx context.Context, args *proto.GetRequest) (*proto.GetReply, error) {
	
	key, err := c.get(ctx, args.Key)
	
	return &proto.GetReply{Key: key}, err
}

func (c *Chord) Put(ctx context.Context, args *proto.PutRequest) (*proto.PutReply, error) {
	
	c.put(ctx, args.Key)

	return &proto.PutReply{}, nil
}