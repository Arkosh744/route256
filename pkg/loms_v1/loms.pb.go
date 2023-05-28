// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: loms.proto

package loms_v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type OrderStatus int32

const (
	OrderStatus_NEW              OrderStatus = 0
	OrderStatus_AWAITING_PAYMENT OrderStatus = 1
	OrderStatus_FAILED           OrderStatus = 2
	OrderStatus_PAID             OrderStatus = 3
	OrderStatus_CANCELLED        OrderStatus = 4
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "NEW",
		1: "AWAITING_PAYMENT",
		2: "FAILED",
		3: "PAID",
		4: "CANCELLED",
	}
	OrderStatus_value = map[string]int32{
		"NEW":              0,
		"AWAITING_PAYMENT": 1,
		"FAILED":           2,
		"PAID":             3,
		"CANCELLED":        4,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_loms_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_loms_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{0}
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User int64   `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Item []*Item `protobuf:"bytes,2,rep,name=Item,proto3" json:"Item,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *CreateOrderRequest) GetItem() []*Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku   uint32 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{1}
}

func (x *Item) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *Item) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderResponse) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type ListOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status OrderStatus `protobuf:"varint,1,opt,name=status,proto3,enum=route256.loms.OrderStatus" json:"status,omitempty"`
	User   int64       `protobuf:"varint,2,opt,name=user,proto3" json:"user,omitempty"`
	Item   []*Item     `protobuf:"bytes,3,rep,name=Item,proto3" json:"Item,omitempty"`
}

func (x *ListOrdersResponse) Reset() {
	*x = ListOrdersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersResponse) ProtoMessage() {}

func (x *ListOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersResponse.ProtoReflect.Descriptor instead.
func (*ListOrdersResponse) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{3}
}

func (x *ListOrdersResponse) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_NEW
}

func (x *ListOrdersResponse) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *ListOrdersResponse) GetItem() []*Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type StocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku uint32 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *StocksRequest) Reset() {
	*x = StocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StocksRequest) ProtoMessage() {}

func (x *StocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StocksRequest.ProtoReflect.Descriptor instead.
func (*StocksRequest) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{4}
}

func (x *StocksRequest) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

type StocksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stocks []*Stock `protobuf:"bytes,1,rep,name=stocks,proto3" json:"stocks,omitempty"`
}

func (x *StocksResponse) Reset() {
	*x = StocksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StocksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StocksResponse) ProtoMessage() {}

func (x *StocksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StocksResponse.ProtoReflect.Descriptor instead.
func (*StocksResponse) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{5}
}

func (x *StocksResponse) GetStocks() []*Stock {
	if x != nil {
		return x.Stocks
	}
	return nil
}

type Stock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WarehouseID int64  `protobuf:"varint,1,opt,name=warehouseID,proto3" json:"warehouseID,omitempty"`
	Count       uint64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Stock) Reset() {
	*x = Stock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stock) ProtoMessage() {}

func (x *Stock) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stock.ProtoReflect.Descriptor instead.
func (*Stock) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{6}
}

func (x *Stock) GetWarehouseID() int64 {
	if x != nil {
		return x.WarehouseID
	}
	return 0
}

func (x *Stock) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type OrderIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *OrderIDRequest) Reset() {
	*x = OrderIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loms_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderIDRequest) ProtoMessage() {}

func (x *OrderIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loms_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderIDRequest.ProtoReflect.Descriptor instead.
func (*OrderIDRequest) Descriptor() ([]byte, []int) {
	return file_loms_proto_rawDescGZIP(), []int{7}
}

func (x *OrderIDRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

var File_loms_proto protoreflect.FileDescriptor

var file_loms_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c, 0x6f,
	0x6d, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x3b, 0x0a,
	0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x12, 0x21, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x2a, 0x06, 0x10, 0xff, 0xff,
	0x03, 0x20, 0x00, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x30, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x85, 0x01, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c,
	0x6f, 0x6d, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x04, 0x49,
	0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04,
	0x49, 0x74, 0x65, 0x6d, 0x22, 0x21, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52,
	0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x3f, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x12, 0x20, 0x0a, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x51, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x45, 0x57, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x41, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e,
	0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x50, 0x41, 0x49, 0x44, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x32, 0xfd, 0x02, 0x0a, 0x04, 0x4c, 0x6f, 0x6d,
	0x73, 0x12, 0x54, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x21, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c, 0x6f, 0x6d, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c,
	0x6f, 0x6d, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e,
	0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c,
	0x6f, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x61, 0x79, 0x65, 0x64, 0x12, 0x1d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e,
	0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x44, 0x0a, 0x0b, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x45, 0x0a, 0x06, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1c, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x32, 0x35, 0x36, 0x2e, 0x6c, 0x6f, 0x6d, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x32, 0x35, 0x36, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x6f, 0x6d, 0x73, 0x5f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loms_proto_rawDescOnce sync.Once
	file_loms_proto_rawDescData = file_loms_proto_rawDesc
)

func file_loms_proto_rawDescGZIP() []byte {
	file_loms_proto_rawDescOnce.Do(func() {
		file_loms_proto_rawDescData = protoimpl.X.CompressGZIP(file_loms_proto_rawDescData)
	})
	return file_loms_proto_rawDescData
}

var file_loms_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_loms_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_loms_proto_goTypes = []interface{}{
	(OrderStatus)(0),            // 0: route256.loms.OrderStatus
	(*CreateOrderRequest)(nil),  // 1: route256.loms.CreateOrderRequest
	(*Item)(nil),                // 2: route256.loms.Item
	(*CreateOrderResponse)(nil), // 3: route256.loms.CreateOrderResponse
	(*ListOrdersResponse)(nil),  // 4: route256.loms.ListOrdersResponse
	(*StocksRequest)(nil),       // 5: route256.loms.StocksRequest
	(*StocksResponse)(nil),      // 6: route256.loms.StocksResponse
	(*Stock)(nil),               // 7: route256.loms.Stock
	(*OrderIDRequest)(nil),      // 8: route256.loms.OrderIDRequest
	(*empty.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_loms_proto_depIdxs = []int32{
	2, // 0: route256.loms.CreateOrderRequest.Item:type_name -> route256.loms.Item
	0, // 1: route256.loms.ListOrdersResponse.status:type_name -> route256.loms.OrderStatus
	2, // 2: route256.loms.ListOrdersResponse.Item:type_name -> route256.loms.Item
	7, // 3: route256.loms.StocksResponse.stocks:type_name -> route256.loms.Stock
	1, // 4: route256.loms.Loms.CreateOrder:input_type -> route256.loms.CreateOrderRequest
	8, // 5: route256.loms.Loms.ListOrder:input_type -> route256.loms.OrderIDRequest
	8, // 6: route256.loms.Loms.OrderPayed:input_type -> route256.loms.OrderIDRequest
	8, // 7: route256.loms.Loms.CancelOrder:input_type -> route256.loms.OrderIDRequest
	5, // 8: route256.loms.Loms.Stocks:input_type -> route256.loms.StocksRequest
	3, // 9: route256.loms.Loms.CreateOrder:output_type -> route256.loms.CreateOrderResponse
	4, // 10: route256.loms.Loms.ListOrder:output_type -> route256.loms.ListOrdersResponse
	9, // 11: route256.loms.Loms.OrderPayed:output_type -> google.protobuf.Empty
	9, // 12: route256.loms.Loms.CancelOrder:output_type -> google.protobuf.Empty
	6, // 13: route256.loms.Loms.Stocks:output_type -> route256.loms.StocksResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_loms_proto_init() }
func file_loms_proto_init() {
	if File_loms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_loms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_loms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
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
		file_loms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrdersResponse); i {
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
		file_loms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StocksRequest); i {
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
		file_loms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StocksResponse); i {
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
		file_loms_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stock); i {
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
		file_loms_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderIDRequest); i {
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
			RawDescriptor: file_loms_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_loms_proto_goTypes,
		DependencyIndexes: file_loms_proto_depIdxs,
		EnumInfos:         file_loms_proto_enumTypes,
		MessageInfos:      file_loms_proto_msgTypes,
	}.Build()
	File_loms_proto = out.File
	file_loms_proto_rawDesc = nil
	file_loms_proto_goTypes = nil
	file_loms_proto_depIdxs = nil
}
