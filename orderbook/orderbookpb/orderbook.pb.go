// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderbook.proto

package orderbookpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type LimitOrderRequest struct {
	Side                 int32    `protobuf:"varint,1,opt,name=side,proto3" json:"side,omitempty"`
	OrderId              string   `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Quantity             string   `protobuf:"bytes,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price                string   `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LimitOrderRequest) Reset()         { *m = LimitOrderRequest{} }
func (m *LimitOrderRequest) String() string { return proto.CompactTextString(m) }
func (*LimitOrderRequest) ProtoMessage()    {}
func (*LimitOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeed55a669e09e60, []int{0}
}

func (m *LimitOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LimitOrderRequest.Unmarshal(m, b)
}
func (m *LimitOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LimitOrderRequest.Marshal(b, m, deterministic)
}
func (m *LimitOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitOrderRequest.Merge(m, src)
}
func (m *LimitOrderRequest) XXX_Size() int {
	return xxx_messageInfo_LimitOrderRequest.Size(m)
}
func (m *LimitOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LimitOrderRequest proto.InternalMessageInfo

func (m *LimitOrderRequest) GetSide() int32 {
	if m != nil {
		return m.Side
	}
	return 0
}

func (m *LimitOrderRequest) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *LimitOrderRequest) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

func (m *LimitOrderRequest) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

type PriceLevel struct {
	Price                string   `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	Quantity             string   `protobuf:"bytes,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PriceLevel) Reset()         { *m = PriceLevel{} }
func (m *PriceLevel) String() string { return proto.CompactTextString(m) }
func (*PriceLevel) ProtoMessage()    {}
func (*PriceLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeed55a669e09e60, []int{1}
}

func (m *PriceLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PriceLevel.Unmarshal(m, b)
}
func (m *PriceLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PriceLevel.Marshal(b, m, deterministic)
}
func (m *PriceLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceLevel.Merge(m, src)
}
func (m *PriceLevel) XXX_Size() int {
	return xxx_messageInfo_PriceLevel.Size(m)
}
func (m *PriceLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceLevel.DiscardUnknown(m)
}

var xxx_messageInfo_PriceLevel proto.InternalMessageInfo

func (m *PriceLevel) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *PriceLevel) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

type Order struct {
	Side                 int32                `protobuf:"varint,1,opt,name=side,proto3" json:"side,omitempty"`
	Id                   string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Quantity             string               `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price                string               `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeed55a669e09e60, []int{2}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetSide() int32 {
	if m != nil {
		return m.Side
	}
	return 0
}

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Order) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

func (m *Order) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

type LimitOrderResponse struct {
	Orders                   []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	Partial                  *Order   `protobuf:"bytes,2,opt,name=partial,proto3" json:"partial,omitempty"`
	PartialQuantityProcessed string   `protobuf:"bytes,3,opt,name=partial_quantity_processed,json=partialQuantityProcessed,proto3" json:"partial_quantity_processed,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *LimitOrderResponse) Reset()         { *m = LimitOrderResponse{} }
func (m *LimitOrderResponse) String() string { return proto.CompactTextString(m) }
func (*LimitOrderResponse) ProtoMessage()    {}
func (*LimitOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeed55a669e09e60, []int{3}
}

func (m *LimitOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LimitOrderResponse.Unmarshal(m, b)
}
func (m *LimitOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LimitOrderResponse.Marshal(b, m, deterministic)
}
func (m *LimitOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitOrderResponse.Merge(m, src)
}
func (m *LimitOrderResponse) XXX_Size() int {
	return xxx_messageInfo_LimitOrderResponse.Size(m)
}
func (m *LimitOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LimitOrderResponse proto.InternalMessageInfo

func (m *LimitOrderResponse) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func (m *LimitOrderResponse) GetPartial() *Order {
	if m != nil {
		return m.Partial
	}
	return nil
}

func (m *LimitOrderResponse) GetPartialQuantityProcessed() string {
	if m != nil {
		return m.PartialQuantityProcessed
	}
	return ""
}

type MarketOrderRequest struct {
	Side                 int32    `protobuf:"varint,1,opt,name=side,proto3" json:"side,omitempty"`
	Quantity             string   `protobuf:"bytes,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarketOrderRequest) Reset()         { *m = MarketOrderRequest{} }
func (m *MarketOrderRequest) String() string { return proto.CompactTextString(m) }
func (*MarketOrderRequest) ProtoMessage()    {}
func (*MarketOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeed55a669e09e60, []int{4}
}

func (m *MarketOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketOrderRequest.Unmarshal(m, b)
}
func (m *MarketOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketOrderRequest.Marshal(b, m, deterministic)
}
func (m *MarketOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketOrderRequest.Merge(m, src)
}
func (m *MarketOrderRequest) XXX_Size() int {
	return xxx_messageInfo_MarketOrderRequest.Size(m)
}
func (m *MarketOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MarketOrderRequest proto.InternalMessageInfo

func (m *MarketOrderRequest) GetSide() int32 {
	if m != nil {
		return m.Side
	}
	return 0
}

func (m *MarketOrderRequest) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

type MarketOrderResponse struct {
	Done                     []*Order `protobuf:"bytes,1,rep,name=done,proto3" json:"done,omitempty"`
	Partial                  *Order   `protobuf:"bytes,2,opt,name=partial,proto3" json:"partial,omitempty"`
	PartialQuantityProcessed string   `protobuf:"bytes,3,opt,name=partial_quantity_processed,json=partialQuantityProcessed,proto3" json:"partial_quantity_processed,omitempty"`
	QuantityLeft             string   `protobuf:"bytes,4,opt,name=quantity_left,json=quantityLeft,proto3" json:"quantity_left,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *MarketOrderResponse) Reset()         { *m = MarketOrderResponse{} }
func (m *MarketOrderResponse) String() string { return proto.CompactTextString(m) }
func (*MarketOrderResponse) ProtoMessage()    {}
func (*MarketOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeed55a669e09e60, []int{5}
}

func (m *MarketOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketOrderResponse.Unmarshal(m, b)
}
func (m *MarketOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketOrderResponse.Marshal(b, m, deterministic)
}
func (m *MarketOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketOrderResponse.Merge(m, src)
}
func (m *MarketOrderResponse) XXX_Size() int {
	return xxx_messageInfo_MarketOrderResponse.Size(m)
}
func (m *MarketOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MarketOrderResponse proto.InternalMessageInfo

func (m *MarketOrderResponse) GetDone() []*Order {
	if m != nil {
		return m.Done
	}
	return nil
}

func (m *MarketOrderResponse) GetPartial() *Order {
	if m != nil {
		return m.Partial
	}
	return nil
}

func (m *MarketOrderResponse) GetPartialQuantityProcessed() string {
	if m != nil {
		return m.PartialQuantityProcessed
	}
	return ""
}

func (m *MarketOrderResponse) GetQuantityLeft() string {
	if m != nil {
		return m.QuantityLeft
	}
	return ""
}

type CancelOrderRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelOrderRequest) Reset()         { *m = CancelOrderRequest{} }
func (m *CancelOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CancelOrderRequest) ProtoMessage()    {}
func (*CancelOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeed55a669e09e60, []int{6}
}

func (m *CancelOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelOrderRequest.Unmarshal(m, b)
}
func (m *CancelOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelOrderRequest.Marshal(b, m, deterministic)
}
func (m *CancelOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelOrderRequest.Merge(m, src)
}
func (m *CancelOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CancelOrderRequest.Size(m)
}
func (m *CancelOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelOrderRequest proto.InternalMessageInfo

func (m *CancelOrderRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CancelOrderResponse struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelOrderResponse) Reset()         { *m = CancelOrderResponse{} }
func (m *CancelOrderResponse) String() string { return proto.CompactTextString(m) }
func (*CancelOrderResponse) ProtoMessage()    {}
func (*CancelOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeed55a669e09e60, []int{7}
}

func (m *CancelOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelOrderResponse.Unmarshal(m, b)
}
func (m *CancelOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelOrderResponse.Marshal(b, m, deterministic)
}
func (m *CancelOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelOrderResponse.Merge(m, src)
}
func (m *CancelOrderResponse) XXX_Size() int {
	return xxx_messageInfo_CancelOrderResponse.Size(m)
}
func (m *CancelOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CancelOrderResponse proto.InternalMessageInfo

func (m *CancelOrderResponse) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeed55a669e09e60, []int{8}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type DepthResponse struct {
	Bids                 []*PriceLevel `protobuf:"bytes,1,rep,name=bids,proto3" json:"bids,omitempty"`
	Asks                 []*PriceLevel `protobuf:"bytes,2,rep,name=asks,proto3" json:"asks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DepthResponse) Reset()         { *m = DepthResponse{} }
func (m *DepthResponse) String() string { return proto.CompactTextString(m) }
func (*DepthResponse) ProtoMessage()    {}
func (*DepthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeed55a669e09e60, []int{9}
}

func (m *DepthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DepthResponse.Unmarshal(m, b)
}
func (m *DepthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DepthResponse.Marshal(b, m, deterministic)
}
func (m *DepthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepthResponse.Merge(m, src)
}
func (m *DepthResponse) XXX_Size() int {
	return xxx_messageInfo_DepthResponse.Size(m)
}
func (m *DepthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DepthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DepthResponse proto.InternalMessageInfo

func (m *DepthResponse) GetBids() []*PriceLevel {
	if m != nil {
		return m.Bids
	}
	return nil
}

func (m *DepthResponse) GetAsks() []*PriceLevel {
	if m != nil {
		return m.Asks
	}
	return nil
}

func init() {
	proto.RegisterType((*LimitOrderRequest)(nil), "orderbookpb.LimitOrderRequest")
	proto.RegisterType((*PriceLevel)(nil), "orderbookpb.PriceLevel")
	proto.RegisterType((*Order)(nil), "orderbookpb.Order")
	proto.RegisterType((*LimitOrderResponse)(nil), "orderbookpb.LimitOrderResponse")
	proto.RegisterType((*MarketOrderRequest)(nil), "orderbookpb.MarketOrderRequest")
	proto.RegisterType((*MarketOrderResponse)(nil), "orderbookpb.MarketOrderResponse")
	proto.RegisterType((*CancelOrderRequest)(nil), "orderbookpb.CancelOrderRequest")
	proto.RegisterType((*CancelOrderResponse)(nil), "orderbookpb.CancelOrderResponse")
	proto.RegisterType((*Empty)(nil), "orderbookpb.Empty")
	proto.RegisterType((*DepthResponse)(nil), "orderbookpb.DepthResponse")
}

func init() { proto.RegisterFile("orderbook.proto", fileDescriptor_aeed55a669e09e60) }

var fileDescriptor_aeed55a669e09e60 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x65, 0x37, 0x6e, 0xda, 0x09, 0x05, 0x75, 0x5a, 0x09, 0xe3, 0x03, 0x8d, 0x0c, 0x42,
	0x11, 0x20, 0x57, 0x0a, 0x07, 0x38, 0x20, 0x90, 0x68, 0x39, 0x20, 0x05, 0x11, 0x4c, 0x25, 0x8e,
	0x91, 0x1d, 0x4f, 0xc2, 0x2a, 0x4e, 0xd6, 0xf5, 0x6e, 0x2a, 0xf5, 0x49, 0x78, 0x09, 0x9e, 0x05,
	0xf1, 0x48, 0xc8, 0xeb, 0xb5, 0xe3, 0x6d, 0x92, 0xf6, 0xd6, 0x9b, 0xbd, 0xf3, 0xed, 0xaf, 0xf9,
	0xe7, 0x1f, 0x1b, 0x1e, 0xf1, 0x3c, 0xa1, 0x3c, 0xe6, 0x7c, 0x16, 0x64, 0x39, 0x97, 0x1c, 0x3b,
	0xf5, 0x41, 0x16, 0x7b, 0x27, 0x53, 0xce, 0xa7, 0x29, 0x9d, 0xaa, 0x52, 0xbc, 0x9c, 0x9c, 0x4a,
	0x36, 0x27, 0x21, 0xa3, 0x79, 0x56, 0xd2, 0xbe, 0x84, 0xc3, 0x01, 0x9b, 0x33, 0xf9, 0xad, 0xb8,
	0x14, 0xd2, 0xe5, 0x92, 0x84, 0x44, 0x84, 0x96, 0x60, 0x09, 0xb9, 0x56, 0xd7, 0xea, 0x39, 0xa1,
	0x7a, 0xc6, 0x27, 0xb0, 0xa7, 0x84, 0x47, 0x2c, 0x71, 0xed, 0xae, 0xd5, 0xdb, 0x0f, 0xdb, 0xea,
	0xfd, 0x4b, 0x82, 0x1e, 0xec, 0x5d, 0x2e, 0xa3, 0x85, 0x64, 0xf2, 0xda, 0xdd, 0x51, 0xa5, 0xfa,
	0x1d, 0x8f, 0xc1, 0xc9, 0x72, 0x36, 0x26, 0xb7, 0xa5, 0x0a, 0xe5, 0x8b, 0xff, 0x01, 0x60, 0x58,
	0x3c, 0x0c, 0xe8, 0x8a, 0xd2, 0x15, 0x63, 0x35, 0x18, 0x43, 0xd5, 0x36, 0x55, 0xfd, 0xdf, 0x16,
	0x38, 0xaa, 0xe3, 0x8d, 0xad, 0x3e, 0x04, 0xbb, 0x6e, 0xd2, 0x66, 0x09, 0xbe, 0x83, 0xfd, 0xda,
	0xb6, 0x6a, 0xb0, 0xd3, 0xf7, 0x82, 0x72, 0x30, 0x41, 0x35, 0x98, 0xe0, 0xa2, 0x22, 0xc2, 0x15,
	0x6c, 0xf4, 0xd0, 0xda, 0xe6, 0xcc, 0x69, 0x3a, 0xfb, 0x63, 0x01, 0x36, 0x07, 0x2a, 0x32, 0xbe,
	0x10, 0x84, 0x2f, 0x61, 0x57, 0x4d, 0x4b, 0xb8, 0x56, 0x77, 0xa7, 0xd7, 0xe9, 0x63, 0xd0, 0x48,
	0x29, 0x28, 0x59, 0x4d, 0xe0, 0x6b, 0x68, 0x67, 0x51, 0x2e, 0x59, 0x94, 0x2a, 0x0f, 0x9b, 0xe1,
	0x0a, 0xc1, 0xf7, 0xe0, 0xe9, 0xc7, 0x51, 0xd5, 0xda, 0x28, 0xcb, 0xf9, 0x98, 0x84, 0xa0, 0x44,
	0xc7, 0xe1, 0x6a, 0xe2, 0xbb, 0x06, 0x86, 0x55, 0xdd, 0x3f, 0x07, 0xfc, 0x1a, 0xe5, 0x33, 0xba,
	0x3b, 0xff, 0xdb, 0xe2, 0xf8, 0x6b, 0xc1, 0x91, 0x21, 0xa3, 0x5d, 0xbf, 0x80, 0x56, 0xc2, 0x17,
	0x74, 0x8b, 0x67, 0x55, 0xbf, 0x4f, 0xc7, 0xf8, 0x0c, 0x0e, 0xea, 0x5b, 0x29, 0x4d, 0xa4, 0xce,
	0xf5, 0x41, 0x75, 0x38, 0xa0, 0x89, 0xf4, 0x9f, 0x03, 0x9e, 0x45, 0x8b, 0x31, 0xa5, 0xc6, 0x58,
	0xca, 0xbd, 0xb2, 0xaa, 0xbd, 0xf2, 0x3f, 0xc2, 0x91, 0x41, 0x69, 0xd7, 0x3d, 0x70, 0x54, 0xf7,
	0x8a, 0xdc, 0xec, 0xa5, 0x04, 0xfc, 0x36, 0x38, 0x9f, 0xe7, 0x99, 0xbc, 0xf6, 0x19, 0x1c, 0x9c,
	0x53, 0x26, 0x7f, 0xd5, 0x1a, 0xaf, 0xa0, 0x15, 0xb3, 0xa4, 0xda, 0x96, 0xc7, 0x86, 0xc4, 0xea,
	0xcb, 0x09, 0x15, 0x54, 0xc0, 0x91, 0x98, 0x09, 0xd7, 0xbe, 0x03, 0x2e, 0xa0, 0xfe, 0x3f, 0x1b,
	0x8e, 0x55, 0x13, 0x9f, 0x38, 0x9f, 0x4d, 0xc3, 0xe1, 0xd9, 0x0f, 0xca, 0xaf, 0x8a, 0xef, 0xed,
	0x02, 0x0e, 0xf5, 0x94, 0x56, 0xfb, 0x8b, 0x4f, 0x0d, 0xb1, 0xb5, 0x3f, 0x85, 0x77, 0xb2, 0xb5,
	0xae, 0x8d, 0xfc, 0x04, 0xd4, 0xaa, 0x8d, 0x05, 0x41, 0xf3, 0xda, 0xfa, 0x06, 0x7a, 0xdd, 0xed,
	0x80, 0x16, 0x1e, 0x42, 0xa7, 0x31, 0xfc, 0x1b, 0x8a, 0xeb, 0xe1, 0xdd, 0x50, 0xdc, 0x94, 0xdb,
	0x5b, 0x70, 0x54, 0x08, 0x68, 0x26, 0xa6, 0x12, 0xf2, 0x3c, 0xe3, 0xcc, 0x08, 0x2b, 0xde, 0x55,
	0x3f, 0x91, 0x37, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x8a, 0xbe, 0x2b, 0x8b, 0x05, 0x00,
	0x00,
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
	ProcessLimitOrder(ctx context.Context, in *LimitOrderRequest, opts ...grpc.CallOption) (*LimitOrderResponse, error)
	ProcessMarketOrder(ctx context.Context, in *MarketOrderRequest, opts ...grpc.CallOption) (*MarketOrderResponse, error)
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
	Depth(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DepthResponse, error)
}

type orderBookgRPCServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderBookgRPCServiceClient(cc *grpc.ClientConn) OrderBookgRPCServiceClient {
	return &orderBookgRPCServiceClient{cc}
}

func (c *orderBookgRPCServiceClient) ProcessLimitOrder(ctx context.Context, in *LimitOrderRequest, opts ...grpc.CallOption) (*LimitOrderResponse, error) {
	out := new(LimitOrderResponse)
	err := c.cc.Invoke(ctx, "/orderbookpb.OrderBookgRPCService/ProcessLimitOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderBookgRPCServiceClient) ProcessMarketOrder(ctx context.Context, in *MarketOrderRequest, opts ...grpc.CallOption) (*MarketOrderResponse, error) {
	out := new(MarketOrderResponse)
	err := c.cc.Invoke(ctx, "/orderbookpb.OrderBookgRPCService/ProcessMarketOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderBookgRPCServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, "/orderbookpb.OrderBookgRPCService/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderBookgRPCServiceClient) Depth(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DepthResponse, error) {
	out := new(DepthResponse)
	err := c.cc.Invoke(ctx, "/orderbookpb.OrderBookgRPCService/Depth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderBookgRPCServiceServer is the server API for OrderBookgRPCService service.
type OrderBookgRPCServiceServer interface {
	ProcessLimitOrder(context.Context, *LimitOrderRequest) (*LimitOrderResponse, error)
	ProcessMarketOrder(context.Context, *MarketOrderRequest) (*MarketOrderResponse, error)
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	Depth(context.Context, *Empty) (*DepthResponse, error)
}

// UnimplementedOrderBookgRPCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderBookgRPCServiceServer struct {
}

func (*UnimplementedOrderBookgRPCServiceServer) ProcessLimitOrder(ctx context.Context, req *LimitOrderRequest) (*LimitOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessLimitOrder not implemented")
}
func (*UnimplementedOrderBookgRPCServiceServer) ProcessMarketOrder(ctx context.Context, req *MarketOrderRequest) (*MarketOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessMarketOrder not implemented")
}
func (*UnimplementedOrderBookgRPCServiceServer) CancelOrder(ctx context.Context, req *CancelOrderRequest) (*CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (*UnimplementedOrderBookgRPCServiceServer) Depth(ctx context.Context, req *Empty) (*DepthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Depth not implemented")
}

func RegisterOrderBookgRPCServiceServer(s *grpc.Server, srv OrderBookgRPCServiceServer) {
	s.RegisterService(&_OrderBookgRPCService_serviceDesc, srv)
}

func _OrderBookgRPCService_ProcessLimitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LimitOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderBookgRPCServiceServer).ProcessLimitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderbookpb.OrderBookgRPCService/ProcessLimitOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderBookgRPCServiceServer).ProcessLimitOrder(ctx, req.(*LimitOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderBookgRPCService_ProcessMarketOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderBookgRPCServiceServer).ProcessMarketOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderbookpb.OrderBookgRPCService/ProcessMarketOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderBookgRPCServiceServer).ProcessMarketOrder(ctx, req.(*MarketOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderBookgRPCService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderBookgRPCServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderbookpb.OrderBookgRPCService/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderBookgRPCServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderBookgRPCService_Depth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderBookgRPCServiceServer).Depth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderbookpb.OrderBookgRPCService/Depth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderBookgRPCServiceServer).Depth(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderBookgRPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orderbookpb.OrderBookgRPCService",
	HandlerType: (*OrderBookgRPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessLimitOrder",
			Handler:    _OrderBookgRPCService_ProcessLimitOrder_Handler,
		},
		{
			MethodName: "ProcessMarketOrder",
			Handler:    _OrderBookgRPCService_ProcessMarketOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _OrderBookgRPCService_CancelOrder_Handler,
		},
		{
			MethodName: "Depth",
			Handler:    _OrderBookgRPCService_Depth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orderbook.proto",
}