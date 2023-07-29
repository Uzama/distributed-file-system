// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: proto/communication.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Communication_FindSuccessor_FullMethodName            = "/main.Communication/FindSuccessor"
	Communication_FindClosestPrecedingNode_FullMethodName = "/main.Communication/FindClosestPrecedingNode"
	Communication_GetPredecessor_FullMethodName           = "/main.Communication/GetPredecessor"
	Communication_GetSuccessor_FullMethodName             = "/main.Communication/GetSuccessor"
	Communication_SetPredecessor_FullMethodName           = "/main.Communication/SetPredecessor"
	Communication_SetSuccessor_FullMethodName             = "/main.Communication/SetSuccessor"
	Communication_Notify_FullMethodName                   = "/main.Communication/Notify"
	Communication_Get_FullMethodName                      = "/main.Communication/Get"
	Communication_Put_FullMethodName                      = "/main.Communication/Put"
	Communication_ListKeys_FullMethodName                 = "/main.Communication/ListKeys"
)

// CommunicationClient is the client API for Communication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommunicationClient interface {
	FindSuccessor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Node, error)
	FindClosestPrecedingNode(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Node, error)
	GetPredecessor(ctx context.Context, in *NN, opts ...grpc.CallOption) (*Node, error)
	GetSuccessor(ctx context.Context, in *NN, opts ...grpc.CallOption) (*Node, error)
	SetPredecessor(ctx context.Context, in *Node, opts ...grpc.CallOption) (*NN, error)
	SetSuccessor(ctx context.Context, in *Node, opts ...grpc.CallOption) (*NN, error)
	Notify(ctx context.Context, in *Node, opts ...grpc.CallOption) (*NN, error)
	// Store
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutReply, error)
	ListKeys(ctx context.Context, in *ListKeysRequest, opts ...grpc.CallOption) (*ListKeysReply, error)
}

type communicationClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunicationClient(cc grpc.ClientConnInterface) CommunicationClient {
	return &communicationClient{cc}
}

func (c *communicationClient) FindSuccessor(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, Communication_FindSuccessor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) FindClosestPrecedingNode(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, Communication_FindClosestPrecedingNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) GetPredecessor(ctx context.Context, in *NN, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, Communication_GetPredecessor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) GetSuccessor(ctx context.Context, in *NN, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, Communication_GetSuccessor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) SetPredecessor(ctx context.Context, in *Node, opts ...grpc.CallOption) (*NN, error) {
	out := new(NN)
	err := c.cc.Invoke(ctx, Communication_SetPredecessor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) SetSuccessor(ctx context.Context, in *Node, opts ...grpc.CallOption) (*NN, error) {
	out := new(NN)
	err := c.cc.Invoke(ctx, Communication_SetSuccessor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) Notify(ctx context.Context, in *Node, opts ...grpc.CallOption) (*NN, error) {
	out := new(NN)
	err := c.cc.Invoke(ctx, Communication_Notify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, Communication_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutReply, error) {
	out := new(PutReply)
	err := c.cc.Invoke(ctx, Communication_Put_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) ListKeys(ctx context.Context, in *ListKeysRequest, opts ...grpc.CallOption) (*ListKeysReply, error) {
	out := new(ListKeysReply)
	err := c.cc.Invoke(ctx, Communication_ListKeys_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicationServer is the server API for Communication service.
// All implementations must embed UnimplementedCommunicationServer
// for forward compatibility
type CommunicationServer interface {
	FindSuccessor(context.Context, *ID) (*Node, error)
	FindClosestPrecedingNode(context.Context, *ID) (*Node, error)
	GetPredecessor(context.Context, *NN) (*Node, error)
	GetSuccessor(context.Context, *NN) (*Node, error)
	SetPredecessor(context.Context, *Node) (*NN, error)
	SetSuccessor(context.Context, *Node) (*NN, error)
	Notify(context.Context, *Node) (*NN, error)
	// Store
	Get(context.Context, *GetRequest) (*GetReply, error)
	Put(context.Context, *PutRequest) (*PutReply, error)
	ListKeys(context.Context, *ListKeysRequest) (*ListKeysReply, error)
	mustEmbedUnimplementedCommunicationServer()
}

// UnimplementedCommunicationServer must be embedded to have forward compatible implementations.
type UnimplementedCommunicationServer struct {
}

func (UnimplementedCommunicationServer) FindSuccessor(context.Context, *ID) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessor not implemented")
}
func (UnimplementedCommunicationServer) FindClosestPrecedingNode(context.Context, *ID) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindClosestPrecedingNode not implemented")
}
func (UnimplementedCommunicationServer) GetPredecessor(context.Context, *NN) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPredecessor not implemented")
}
func (UnimplementedCommunicationServer) GetSuccessor(context.Context, *NN) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuccessor not implemented")
}
func (UnimplementedCommunicationServer) SetPredecessor(context.Context, *Node) (*NN, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPredecessor not implemented")
}
func (UnimplementedCommunicationServer) SetSuccessor(context.Context, *Node) (*NN, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSuccessor not implemented")
}
func (UnimplementedCommunicationServer) Notify(context.Context, *Node) (*NN, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedCommunicationServer) Get(context.Context, *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCommunicationServer) Put(context.Context, *PutRequest) (*PutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedCommunicationServer) ListKeys(context.Context, *ListKeysRequest) (*ListKeysReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKeys not implemented")
}
func (UnimplementedCommunicationServer) mustEmbedUnimplementedCommunicationServer() {}

// UnsafeCommunicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommunicationServer will
// result in compilation errors.
type UnsafeCommunicationServer interface {
	mustEmbedUnimplementedCommunicationServer()
}

func RegisterCommunicationServer(s grpc.ServiceRegistrar, srv CommunicationServer) {
	s.RegisterService(&Communication_ServiceDesc, srv)
}

func _Communication_FindSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).FindSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Communication_FindSuccessor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).FindSuccessor(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_FindClosestPrecedingNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).FindClosestPrecedingNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Communication_FindClosestPrecedingNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).FindClosestPrecedingNode(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_GetPredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).GetPredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Communication_GetPredecessor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).GetPredecessor(ctx, req.(*NN))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_GetSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NN)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).GetSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Communication_GetSuccessor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).GetSuccessor(ctx, req.(*NN))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_SetPredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).SetPredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Communication_SetPredecessor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).SetPredecessor(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_SetSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).SetSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Communication_SetSuccessor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).SetSuccessor(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Communication_Notify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).Notify(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Communication_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Communication_Put_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_ListKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).ListKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Communication_ListKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).ListKeys(ctx, req.(*ListKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Communication_ServiceDesc is the grpc.ServiceDesc for Communication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Communication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.Communication",
	HandlerType: (*CommunicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindSuccessor",
			Handler:    _Communication_FindSuccessor_Handler,
		},
		{
			MethodName: "FindClosestPrecedingNode",
			Handler:    _Communication_FindClosestPrecedingNode_Handler,
		},
		{
			MethodName: "GetPredecessor",
			Handler:    _Communication_GetPredecessor_Handler,
		},
		{
			MethodName: "GetSuccessor",
			Handler:    _Communication_GetSuccessor_Handler,
		},
		{
			MethodName: "SetPredecessor",
			Handler:    _Communication_SetPredecessor_Handler,
		},
		{
			MethodName: "SetSuccessor",
			Handler:    _Communication_SetSuccessor_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Communication_Notify_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Communication_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Communication_Put_Handler,
		},
		{
			MethodName: "ListKeys",
			Handler:    _Communication_ListKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/communication.proto",
}