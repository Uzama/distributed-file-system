package main

import (
	"chord-network/proto"
	"context"
)


func (c *Chord) Notify(ctx context.Context, potentialPred *proto.Node) (*proto.NN, error) {
	
	return &proto.NN{}, nil
}


func (c *Chord) FindSuccessor(ctx context.Context, id *proto.ID) (*proto.Node, error) {
	
	return nil, nil
}


func (c *Chord) FindClosestPrecedingNode(ctx context.Context, id *proto.ID) (*proto.Node, error) {
	
	return nil, nil
}


func (c *Chord) GetSuccessor(context.Context, *proto.NN) (*proto.Node, error) {

	return nil, nil
}


func (c *Chord) GetPredecessor(context.Context, *proto.NN) (*proto.Node, error) {
	
	return nil, nil
}


func (c *Chord) SetPredecessor(ctx context.Context, pred *proto.Node) (*proto.NN, error) {
	

	return &proto.NN{}, nil
}


func (c *Chord) SetSuccessor(ctx context.Context, succ *proto.Node) (*proto.NN, error) {
	
	return &proto.NN{}, nil
}

func (c *Chord) Get(ctx context.Context, args *proto.GetRequest) (*proto.GetReply, error) {
	
	return nil, nil
}

func (c *Chord) Put(ctx context.Context, args *proto.PutRequest) (*proto.PutReply, error) {
	return nil, nil
}