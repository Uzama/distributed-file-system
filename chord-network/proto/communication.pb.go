// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: proto/communication.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Node contains a node ID and ip address.
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ip string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_communication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_proto_communication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_proto_communication_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Node) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type NN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NN) Reset() {
	*x = NN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_communication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NN) ProtoMessage() {}

func (x *NN) ProtoReflect() protoreflect.Message {
	mi := &file_proto_communication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NN.ProtoReflect.Descriptor instead.
func (*NN) Descriptor() ([]byte, []int) {
	return file_proto_communication_proto_rawDescGZIP(), []int{1}
}

// ID contains a node ID
type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_communication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_communication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_proto_communication_proto_rawDescGZIP(), []int{2}
}

func (x *ID) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

// For KV store
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_communication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_communication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_proto_communication_proto_rawDescGZIP(), []int{3}
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetReply) Reset() {
	*x = GetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_communication_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReply) ProtoMessage() {}

func (x *GetReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_communication_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReply.ProtoReflect.Descriptor instead.
func (*GetReply) Descriptor() ([]byte, []int) {
	return file_proto_communication_proto_rawDescGZIP(), []int{4}
}

func (x *GetReply) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type PutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *PutRequest) Reset() {
	*x = PutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_communication_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRequest) ProtoMessage() {}

func (x *PutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_communication_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRequest.ProtoReflect.Descriptor instead.
func (*PutRequest) Descriptor() ([]byte, []int) {
	return file_proto_communication_proto_rawDescGZIP(), []int{5}
}

func (x *PutRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type PutReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PutReply) Reset() {
	*x = PutReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_communication_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutReply) ProtoMessage() {}

func (x *PutReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_communication_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutReply.ProtoReflect.Descriptor instead.
func (*PutReply) Descriptor() ([]byte, []int) {
	return file_proto_communication_proto_rawDescGZIP(), []int{6}
}

type ListKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start []byte `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   []byte `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *ListKeysRequest) Reset() {
	*x = ListKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_communication_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKeysRequest) ProtoMessage() {}

func (x *ListKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_communication_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKeysRequest.ProtoReflect.Descriptor instead.
func (*ListKeysRequest) Descriptor() ([]byte, []int) {
	return file_proto_communication_proto_rawDescGZIP(), []int{7}
}

func (x *ListKeysRequest) GetStart() []byte {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *ListKeysRequest) GetEnd() []byte {
	if x != nil {
		return x.End
	}
	return nil
}

type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_communication_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_proto_communication_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_proto_communication_proto_rawDescGZIP(), []int{8}
}

func (x *Key) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ListKeysReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []*Key `protobuf:"bytes,1,rep,name=key,proto3" json:"key,omitempty"`
}

func (x *ListKeysReply) Reset() {
	*x = ListKeysReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_communication_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKeysReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKeysReply) ProtoMessage() {}

func (x *ListKeysReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_communication_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKeysReply.ProtoReflect.Descriptor instead.
func (*ListKeysReply) Descriptor() ([]byte, []int) {
	return file_proto_communication_proto_rawDescGZIP(), []int{9}
}

func (x *ListKeysReply) GetKey() []*Key {
	if x != nil {
		return x.Key
	}
	return nil
}

var File_proto_communication_proto protoreflect.FileDescriptor

var file_proto_communication_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69,
	0x6e, 0x22, 0x26, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x04, 0x0a, 0x02, 0x4e, 0x4e, 0x22,
	0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x1c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x1e, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x0a, 0x0a, 0x08, 0x50, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x39, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x17, 0x0a, 0x03, 0x4b, 0x65,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x2c, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x32, 0xae, 0x03, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x12, 0x08, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x49, 0x44, 0x1a, 0x0a,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x18, 0x46, 0x69,
	0x6e, 0x64, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x73, 0x74, 0x50, 0x72, 0x65, 0x63, 0x65, 0x64, 0x69,
	0x6e, 0x67, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x49, 0x44,
	0x1a, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x65, 0x64, 0x65, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x08,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4e, 0x4e, 0x1a, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x12, 0x08, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4e, 0x4e, 0x1a, 0x0a,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x53, 0x65,
	0x74, 0x50, 0x72, 0x65, 0x64, 0x65, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x0a, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x08, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x4e, 0x4e, 0x12, 0x24, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x12, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x08,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4e, 0x4e, 0x12, 0x1e, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x08,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4e, 0x4e, 0x12, 0x27, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x27, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x08, 0x4c, 0x69,
	0x73, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x67, 0x65,
	0x6e, 0x3b, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_communication_proto_rawDescOnce sync.Once
	file_proto_communication_proto_rawDescData = file_proto_communication_proto_rawDesc
)

func file_proto_communication_proto_rawDescGZIP() []byte {
	file_proto_communication_proto_rawDescOnce.Do(func() {
		file_proto_communication_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_communication_proto_rawDescData)
	})
	return file_proto_communication_proto_rawDescData
}

var file_proto_communication_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_communication_proto_goTypes = []interface{}{
	(*Node)(nil),            // 0: main.Node
	(*NN)(nil),              // 1: main.NN
	(*ID)(nil),              // 2: main.ID
	(*GetRequest)(nil),      // 3: main.GetRequest
	(*GetReply)(nil),        // 4: main.GetReply
	(*PutRequest)(nil),      // 5: main.PutRequest
	(*PutReply)(nil),        // 6: main.PutReply
	(*ListKeysRequest)(nil), // 7: main.ListKeysRequest
	(*Key)(nil),             // 8: main.Key
	(*ListKeysReply)(nil),   // 9: main.ListKeysReply
}
var file_proto_communication_proto_depIdxs = []int32{
	8,  // 0: main.ListKeysReply.key:type_name -> main.Key
	2,  // 1: main.Communication.FindSuccessor:input_type -> main.ID
	2,  // 2: main.Communication.FindClosestPrecedingNode:input_type -> main.ID
	1,  // 3: main.Communication.GetPredecessor:input_type -> main.NN
	1,  // 4: main.Communication.GetSuccessor:input_type -> main.NN
	0,  // 5: main.Communication.SetPredecessor:input_type -> main.Node
	0,  // 6: main.Communication.SetSuccessor:input_type -> main.Node
	0,  // 7: main.Communication.Notify:input_type -> main.Node
	3,  // 8: main.Communication.Get:input_type -> main.GetRequest
	5,  // 9: main.Communication.Put:input_type -> main.PutRequest
	7,  // 10: main.Communication.ListKeys:input_type -> main.ListKeysRequest
	0,  // 11: main.Communication.FindSuccessor:output_type -> main.Node
	0,  // 12: main.Communication.FindClosestPrecedingNode:output_type -> main.Node
	0,  // 13: main.Communication.GetPredecessor:output_type -> main.Node
	0,  // 14: main.Communication.GetSuccessor:output_type -> main.Node
	1,  // 15: main.Communication.SetPredecessor:output_type -> main.NN
	1,  // 16: main.Communication.SetSuccessor:output_type -> main.NN
	1,  // 17: main.Communication.Notify:output_type -> main.NN
	4,  // 18: main.Communication.Get:output_type -> main.GetReply
	6,  // 19: main.Communication.Put:output_type -> main.PutReply
	9,  // 20: main.Communication.ListKeys:output_type -> main.ListKeysReply
	11, // [11:21] is the sub-list for method output_type
	1,  // [1:11] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_communication_proto_init() }
func file_proto_communication_proto_init() {
	if File_proto_communication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_communication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_communication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NN); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_communication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_communication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_communication_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_communication_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_communication_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_communication_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKeysRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_communication_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_communication_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKeysReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_communication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_communication_proto_goTypes,
		DependencyIndexes: file_proto_communication_proto_depIdxs,
		MessageInfos:      file_proto_communication_proto_msgTypes,
	}.Build()
	File_proto_communication_proto = out.File
	file_proto_communication_proto_rawDesc = nil
	file_proto_communication_proto_goTypes = nil
	file_proto_communication_proto_depIdxs = nil
}