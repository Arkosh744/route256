// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: loms.proto

package loms_v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LomsClient is the client API for Loms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LomsClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	ListOrder(ctx context.Context, in *OrderIDRequest, opts ...grpc.CallOption) (*ListOrderResponse, error)
	OrderPayed(ctx context.Context, in *OrderIDRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CancelOrder(ctx context.Context, in *OrderIDRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Stocks(ctx context.Context, in *StocksRequest, opts ...grpc.CallOption) (*StocksResponse, error)
}

type lomsClient struct {
	cc grpc.ClientConnInterface
}

func NewLomsClient(cc grpc.ClientConnInterface) LomsClient {
	return &lomsClient{cc}
}

func (c *lomsClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, "/route256.loms.Loms/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lomsClient) ListOrder(ctx context.Context, in *OrderIDRequest, opts ...grpc.CallOption) (*ListOrderResponse, error) {
	out := new(ListOrderResponse)
	err := c.cc.Invoke(ctx, "/route256.loms.Loms/ListOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lomsClient) OrderPayed(ctx context.Context, in *OrderIDRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/route256.loms.Loms/OrderPayed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lomsClient) CancelOrder(ctx context.Context, in *OrderIDRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/route256.loms.Loms/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lomsClient) Stocks(ctx context.Context, in *StocksRequest, opts ...grpc.CallOption) (*StocksResponse, error) {
	out := new(StocksResponse)
	err := c.cc.Invoke(ctx, "/route256.loms.Loms/Stocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LomsServer is the server API for Loms service.
// All implementations must embed UnimplementedLomsServer
// for forward compatibility
type LomsServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	ListOrder(context.Context, *OrderIDRequest) (*ListOrderResponse, error)
	OrderPayed(context.Context, *OrderIDRequest) (*empty.Empty, error)
	CancelOrder(context.Context, *OrderIDRequest) (*empty.Empty, error)
	Stocks(context.Context, *StocksRequest) (*StocksResponse, error)
	mustEmbedUnimplementedLomsServer()
}

// UnimplementedLomsServer must be embedded to have forward compatible implementations.
type UnimplementedLomsServer struct {
}

func (UnimplementedLomsServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedLomsServer) ListOrder(context.Context, *OrderIDRequest) (*ListOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrder not implemented")
}
func (UnimplementedLomsServer) OrderPayed(context.Context, *OrderIDRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderPayed not implemented")
}
func (UnimplementedLomsServer) CancelOrder(context.Context, *OrderIDRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedLomsServer) Stocks(context.Context, *StocksRequest) (*StocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stocks not implemented")
}
func (UnimplementedLomsServer) mustEmbedUnimplementedLomsServer() {}

// UnsafeLomsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LomsServer will
// result in compilation errors.
type UnsafeLomsServer interface {
	mustEmbedUnimplementedLomsServer()
}

func RegisterLomsServer(s grpc.ServiceRegistrar, srv LomsServer) {
	s.RegisterService(&Loms_ServiceDesc, srv)
}

func _Loms_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LomsServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route256.loms.Loms/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LomsServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loms_ListOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LomsServer).ListOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route256.loms.Loms/ListOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LomsServer).ListOrder(ctx, req.(*OrderIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loms_OrderPayed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LomsServer).OrderPayed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route256.loms.Loms/OrderPayed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LomsServer).OrderPayed(ctx, req.(*OrderIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loms_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LomsServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route256.loms.Loms/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LomsServer).CancelOrder(ctx, req.(*OrderIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loms_Stocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LomsServer).Stocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route256.loms.Loms/Stocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LomsServer).Stocks(ctx, req.(*StocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Loms_ServiceDesc is the grpc.ServiceDesc for Loms service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Loms_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "route256.loms.Loms",
	HandlerType: (*LomsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _Loms_CreateOrder_Handler,
		},
		{
			MethodName: "ListOrder",
			Handler:    _Loms_ListOrder_Handler,
		},
		{
			MethodName: "OrderPayed",
			Handler:    _Loms_OrderPayed_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Loms_CancelOrder_Handler,
		},
		{
			MethodName: "Stocks",
			Handler:    _Loms_Stocks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loms.proto",
}
