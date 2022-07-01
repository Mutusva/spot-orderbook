// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderbook.proto

package orderbookpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("orderbook.proto", fileDescriptor_aeed55a669e09e60) }

var fileDescriptor_aeed55a669e09e60 = []byte{
	// 74 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x2f, 0x4a, 0x49,
	0x2d, 0x4a, 0xca, 0xcf, 0xcf, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0x0b, 0x14,
	0x24, 0x19, 0x89, 0x71, 0x89, 0xf8, 0x83, 0xb8, 0x4e, 0xf9, 0xf9, 0xd9, 0xe9, 0x41, 0x01, 0xce,
	0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x49, 0x6c, 0x60, 0xb5, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd1, 0x28, 0xee, 0x88, 0x3e, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderBookgRPCServiceClient is the client API for OrderBookgRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderBookgRPCServiceClient interface {
}

type orderBookgRPCServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderBookgRPCServiceClient(cc *grpc.ClientConn) OrderBookgRPCServiceClient {
	return &orderBookgRPCServiceClient{cc}
}

// OrderBookgRPCServiceServer is the server API for OrderBookgRPCService service.
type OrderBookgRPCServiceServer interface {
}

// UnimplementedOrderBookgRPCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderBookgRPCServiceServer struct {
}

func RegisterOrderBookgRPCServiceServer(s *grpc.Server, srv OrderBookgRPCServiceServer) {
	s.RegisterService(&_OrderBookgRPCService_serviceDesc, srv)
}

var _OrderBookgRPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orderbookpb.OrderBookgRPCService",
	HandlerType: (*OrderBookgRPCServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "orderbook.proto",
}
