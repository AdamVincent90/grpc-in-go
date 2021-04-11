// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: proto/calculate.proto

package calculate

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Sums struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum1 int32 `protobuf:"varint,1,opt,name=sum1,proto3" json:"sum1,omitempty"`
	Sum2 int32 `protobuf:"varint,2,opt,name=sum2,proto3" json:"sum2,omitempty"`
}

func (x *Sums) Reset() {
	*x = Sums{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sums) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sums) ProtoMessage() {}

func (x *Sums) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sums.ProtoReflect.Descriptor instead.
func (*Sums) Descriptor() ([]byte, []int) {
	return file_proto_calculate_proto_rawDescGZIP(), []int{0}
}

func (x *Sums) GetSum1() int32 {
	if x != nil {
		return x.Sum1
	}
	return 0
}

func (x *Sums) GetSum2() int32 {
	if x != nil {
		return x.Sum2
	}
	return 0
}

type SquareRootRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number float32 `protobuf:"fixed32,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *SquareRootRequest) Reset() {
	*x = SquareRootRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquareRootRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareRootRequest) ProtoMessage() {}

func (x *SquareRootRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareRootRequest.ProtoReflect.Descriptor instead.
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculate_proto_rawDescGZIP(), []int{1}
}

func (x *SquareRootRequest) GetNumber() float32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type SquareRootResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SquareRootResponse) Reset() {
	*x = SquareRootResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquareRootResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareRootResponse) ProtoMessage() {}

func (x *SquareRootResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareRootResponse.ProtoReflect.Descriptor instead.
func (*SquareRootResponse) Descriptor() ([]byte, []int) {
	return file_proto_calculate_proto_rawDescGZIP(), []int{2}
}

func (x *SquareRootResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sums *Sums `protobuf:"bytes,1,opt,name=sums,proto3" json:"sums,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculate_proto_rawDescGZIP(), []int{3}
}

func (x *AddRequest) GetSums() *Sums {
	if x != nil {
		return x.Sums
	}
	return nil
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_proto_calculate_proto_rawDescGZIP(), []int{4}
}

func (x *AddResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type PrimeNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *PrimeNumberRequest) Reset() {
	*x = PrimeNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberRequest) ProtoMessage() {}

func (x *PrimeNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberRequest.ProtoReflect.Descriptor instead.
func (*PrimeNumberRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculate_proto_rawDescGZIP(), []int{5}
}

func (x *PrimeNumberRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type PrimeNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PrimeNumberResponse) Reset() {
	*x = PrimeNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeNumberResponse) ProtoMessage() {}

func (x *PrimeNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeNumberResponse.ProtoReflect.Descriptor instead.
func (*PrimeNumberResponse) Descriptor() ([]byte, []int) {
	return file_proto_calculate_proto_rawDescGZIP(), []int{6}
}

func (x *PrimeNumberResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type ClientStreamNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num float32 `protobuf:"fixed32,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *ClientStreamNumberRequest) Reset() {
	*x = ClientStreamNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculate_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamNumberRequest) ProtoMessage() {}

func (x *ClientStreamNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculate_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamNumberRequest.ProtoReflect.Descriptor instead.
func (*ClientStreamNumberRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculate_proto_rawDescGZIP(), []int{7}
}

func (x *ClientStreamNumberRequest) GetNum() float32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type ClientStreamNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num float32 `protobuf:"fixed32,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *ClientStreamNumberResponse) Reset() {
	*x = ClientStreamNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculate_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamNumberResponse) ProtoMessage() {}

func (x *ClientStreamNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculate_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamNumberResponse.ProtoReflect.Descriptor instead.
func (*ClientStreamNumberResponse) Descriptor() ([]byte, []int) {
	return file_proto_calculate_proto_rawDescGZIP(), []int{8}
}

func (x *ClientStreamNumberResponse) GetNum() float32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type StreamNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *StreamNumberRequest) Reset() {
	*x = StreamNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculate_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamNumberRequest) ProtoMessage() {}

func (x *StreamNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculate_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamNumberRequest.ProtoReflect.Descriptor instead.
func (*StreamNumberRequest) Descriptor() ([]byte, []int) {
	return file_proto_calculate_proto_rawDescGZIP(), []int{9}
}

func (x *StreamNumberRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type StreamNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *StreamNumberResponse) Reset() {
	*x = StreamNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_calculate_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamNumberResponse) ProtoMessage() {}

func (x *StreamNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_calculate_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamNumberResponse.ProtoReflect.Descriptor instead.
func (*StreamNumberResponse) Descriptor() ([]byte, []int) {
	return file_proto_calculate_proto_rawDescGZIP(), []int{10}
}

func (x *StreamNumberResponse) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_proto_calculate_proto protoreflect.FileDescriptor

var file_proto_calculate_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x22, 0x2e, 0x0a, 0x04, 0x53, 0x75, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x75,
	0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x75, 0x6d, 0x31, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x75,
	0x6d, 0x32, 0x22, 0x2b, 0x0a, 0x11, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x2c, 0x0a, 0x12, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x31, 0x0a,
	0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x73,
	0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x75, 0x6d, 0x73, 0x52, 0x04, 0x73, 0x75, 0x6d, 0x73,
	0x22, 0x25, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x26, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x6d, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22,
	0x2d, 0x0a, 0x13, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2d,
	0x0a, 0x19, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e,
	0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x2e, 0x0a,
	0x1a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e,
	0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x27, 0x0a,
	0x13, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x28, 0x0a, 0x14, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d,
	0x32, 0xa6, 0x03, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x15, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a,
	0x12, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x50,
	0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x60, 0x0a, 0x0d, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x52, 0x0a, 0x09, 0x4d, 0x61, 0x78, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x4b, 0x0a, 0x0a,
	0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_calculate_proto_rawDescOnce sync.Once
	file_proto_calculate_proto_rawDescData = file_proto_calculate_proto_rawDesc
)

func file_proto_calculate_proto_rawDescGZIP() []byte {
	file_proto_calculate_proto_rawDescOnce.Do(func() {
		file_proto_calculate_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_calculate_proto_rawDescData)
	})
	return file_proto_calculate_proto_rawDescData
}

var file_proto_calculate_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_calculate_proto_goTypes = []interface{}{
	(*Sums)(nil),                       // 0: calculate.Sums
	(*SquareRootRequest)(nil),          // 1: calculate.SquareRootRequest
	(*SquareRootResponse)(nil),         // 2: calculate.SquareRootResponse
	(*AddRequest)(nil),                 // 3: calculate.AddRequest
	(*AddResponse)(nil),                // 4: calculate.AddResponse
	(*PrimeNumberRequest)(nil),         // 5: calculate.PrimeNumberRequest
	(*PrimeNumberResponse)(nil),        // 6: calculate.PrimeNumberResponse
	(*ClientStreamNumberRequest)(nil),  // 7: calculate.ClientStreamNumberRequest
	(*ClientStreamNumberResponse)(nil), // 8: calculate.ClientStreamNumberResponse
	(*StreamNumberRequest)(nil),        // 9: calculate.StreamNumberRequest
	(*StreamNumberResponse)(nil),       // 10: calculate.StreamNumberResponse
}
var file_proto_calculate_proto_depIdxs = []int32{
	0,  // 0: calculate.AddRequest.sums:type_name -> calculate.Sums
	3,  // 1: calculate.CalculateService.Add:input_type -> calculate.AddRequest
	5,  // 2: calculate.CalculateService.PrimeDecomposition:input_type -> calculate.PrimeNumberRequest
	7,  // 3: calculate.CalculateService.AverageNumber:input_type -> calculate.ClientStreamNumberRequest
	9,  // 4: calculate.CalculateService.MaxNumber:input_type -> calculate.StreamNumberRequest
	1,  // 5: calculate.CalculateService.SquareRoot:input_type -> calculate.SquareRootRequest
	4,  // 6: calculate.CalculateService.Add:output_type -> calculate.AddResponse
	6,  // 7: calculate.CalculateService.PrimeDecomposition:output_type -> calculate.PrimeNumberResponse
	8,  // 8: calculate.CalculateService.AverageNumber:output_type -> calculate.ClientStreamNumberResponse
	10, // 9: calculate.CalculateService.MaxNumber:output_type -> calculate.StreamNumberResponse
	2,  // 10: calculate.CalculateService.SquareRoot:output_type -> calculate.SquareRootResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_calculate_proto_init() }
func file_proto_calculate_proto_init() {
	if File_proto_calculate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_calculate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sums); i {
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
		file_proto_calculate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SquareRootRequest); i {
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
		file_proto_calculate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SquareRootResponse); i {
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
		file_proto_calculate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_proto_calculate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
		file_proto_calculate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberRequest); i {
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
		file_proto_calculate_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeNumberResponse); i {
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
		file_proto_calculate_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStreamNumberRequest); i {
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
		file_proto_calculate_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStreamNumberResponse); i {
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
		file_proto_calculate_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamNumberRequest); i {
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
		file_proto_calculate_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamNumberResponse); i {
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
			RawDescriptor: file_proto_calculate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_calculate_proto_goTypes,
		DependencyIndexes: file_proto_calculate_proto_depIdxs,
		MessageInfos:      file_proto_calculate_proto_msgTypes,
	}.Build()
	File_proto_calculate_proto = out.File
	file_proto_calculate_proto_rawDesc = nil
	file_proto_calculate_proto_goTypes = nil
	file_proto_calculate_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculateServiceClient is the client API for CalculateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculateServiceClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	PrimeDecomposition(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (CalculateService_PrimeDecompositionClient, error)
	AverageNumber(ctx context.Context, opts ...grpc.CallOption) (CalculateService_AverageNumberClient, error)
	MaxNumber(ctx context.Context, opts ...grpc.CallOption) (CalculateService_MaxNumberClient, error)
	// An error will be thrown if the argument within SquareRootRequest involves a
	// negative number, Error result would specify - INVALID_ARGUMENT
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type calculateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculateServiceClient(cc grpc.ClientConnInterface) CalculateServiceClient {
	return &calculateServiceClient{cc}
}

func (c *calculateServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/calculate.CalculateService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateServiceClient) PrimeDecomposition(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (CalculateService_PrimeDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculateService_serviceDesc.Streams[0], "/calculate.CalculateService/PrimeDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServicePrimeDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculateService_PrimeDecompositionClient interface {
	Recv() (*PrimeNumberResponse, error)
	grpc.ClientStream
}

type calculateServicePrimeDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculateServicePrimeDecompositionClient) Recv() (*PrimeNumberResponse, error) {
	m := new(PrimeNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) AverageNumber(ctx context.Context, opts ...grpc.CallOption) (CalculateService_AverageNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculateService_serviceDesc.Streams[1], "/calculate.CalculateService/AverageNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceAverageNumberClient{stream}
	return x, nil
}

type CalculateService_AverageNumberClient interface {
	Send(*ClientStreamNumberRequest) error
	CloseAndRecv() (*ClientStreamNumberResponse, error)
	grpc.ClientStream
}

type calculateServiceAverageNumberClient struct {
	grpc.ClientStream
}

func (x *calculateServiceAverageNumberClient) Send(m *ClientStreamNumberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceAverageNumberClient) CloseAndRecv() (*ClientStreamNumberResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ClientStreamNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) MaxNumber(ctx context.Context, opts ...grpc.CallOption) (CalculateService_MaxNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculateService_serviceDesc.Streams[2], "/calculate.CalculateService/MaxNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceMaxNumberClient{stream}
	return x, nil
}

type CalculateService_MaxNumberClient interface {
	Send(*StreamNumberRequest) error
	Recv() (*StreamNumberResponse, error)
	grpc.ClientStream
}

type calculateServiceMaxNumberClient struct {
	grpc.ClientStream
}

func (x *calculateServiceMaxNumberClient) Send(m *StreamNumberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceMaxNumberClient) Recv() (*StreamNumberResponse, error) {
	m := new(StreamNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/calculate.CalculateService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculateServiceServer is the server API for CalculateService service.
type CalculateServiceServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	PrimeDecomposition(*PrimeNumberRequest, CalculateService_PrimeDecompositionServer) error
	AverageNumber(CalculateService_AverageNumberServer) error
	MaxNumber(CalculateService_MaxNumberServer) error
	// An error will be thrown if the argument within SquareRootRequest involves a
	// negative number, Error result would specify - INVALID_ARGUMENT
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
}

// UnimplementedCalculateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculateServiceServer struct {
}

func (*UnimplementedCalculateServiceServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCalculateServiceServer) PrimeDecomposition(*PrimeNumberRequest, CalculateService_PrimeDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeDecomposition not implemented")
}
func (*UnimplementedCalculateServiceServer) AverageNumber(CalculateService_AverageNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method AverageNumber not implemented")
}
func (*UnimplementedCalculateServiceServer) MaxNumber(CalculateService_MaxNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method MaxNumber not implemented")
}
func (*UnimplementedCalculateServiceServer) SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}

func RegisterCalculateServiceServer(s *grpc.Server, srv CalculateServiceServer) {
	s.RegisterService(&_CalculateService_serviceDesc, srv)
}

func _CalculateService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculate.CalculateService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculateService_PrimeDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculateServiceServer).PrimeDecomposition(m, &calculateServicePrimeDecompositionServer{stream})
}

type CalculateService_PrimeDecompositionServer interface {
	Send(*PrimeNumberResponse) error
	grpc.ServerStream
}

type calculateServicePrimeDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculateServicePrimeDecompositionServer) Send(m *PrimeNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculateService_AverageNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).AverageNumber(&calculateServiceAverageNumberServer{stream})
}

type CalculateService_AverageNumberServer interface {
	SendAndClose(*ClientStreamNumberResponse) error
	Recv() (*ClientStreamNumberRequest, error)
	grpc.ServerStream
}

type calculateServiceAverageNumberServer struct {
	grpc.ServerStream
}

func (x *calculateServiceAverageNumberServer) SendAndClose(m *ClientStreamNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceAverageNumberServer) Recv() (*ClientStreamNumberRequest, error) {
	m := new(ClientStreamNumberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculateService_MaxNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).MaxNumber(&calculateServiceMaxNumberServer{stream})
}

type CalculateService_MaxNumberServer interface {
	Send(*StreamNumberResponse) error
	Recv() (*StreamNumberRequest, error)
	grpc.ServerStream
}

type calculateServiceMaxNumberServer struct {
	grpc.ServerStream
}

func (x *calculateServiceMaxNumberServer) Send(m *StreamNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceMaxNumberServer) Recv() (*StreamNumberRequest, error) {
	m := new(StreamNumberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculateService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculate.CalculateService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculate.CalculateService",
	HandlerType: (*CalculateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalculateService_Add_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _CalculateService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeDecomposition",
			Handler:       _CalculateService_PrimeDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AverageNumber",
			Handler:       _CalculateService_AverageNumber_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "MaxNumber",
			Handler:       _CalculateService_MaxNumber_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/calculate.proto",
}
