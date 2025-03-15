// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: pb/catalog.proto

package catalog

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RestaurantService_CreateRestaurant_FullMethodName  = "/catalog.RestaurantService/CreateRestaurant"
	RestaurantService_GetRestaurants_FullMethodName    = "/catalog.RestaurantService/GetRestaurants"
	RestaurantService_GetRestaurantById_FullMethodName = "/catalog.RestaurantService/GetRestaurantById"
)

// RestaurantServiceClient is the client API for RestaurantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RestaurantServiceClient interface {
	CreateRestaurant(ctx context.Context, in *CreateRestaurantRequest, opts ...grpc.CallOption) (*Restaurant, error)
	GetRestaurants(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetRestaurantsResponse, error)
	GetRestaurantById(ctx context.Context, in *GetRestaurantByIdRequest, opts ...grpc.CallOption) (*Restaurant, error)
}

type restaurantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRestaurantServiceClient(cc grpc.ClientConnInterface) RestaurantServiceClient {
	return &restaurantServiceClient{cc}
}

func (c *restaurantServiceClient) CreateRestaurant(ctx context.Context, in *CreateRestaurantRequest, opts ...grpc.CallOption) (*Restaurant, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Restaurant)
	err := c.cc.Invoke(ctx, RestaurantService_CreateRestaurant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) GetRestaurants(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetRestaurantsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRestaurantsResponse)
	err := c.cc.Invoke(ctx, RestaurantService_GetRestaurants_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantServiceClient) GetRestaurantById(ctx context.Context, in *GetRestaurantByIdRequest, opts ...grpc.CallOption) (*Restaurant, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Restaurant)
	err := c.cc.Invoke(ctx, RestaurantService_GetRestaurantById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestaurantServiceServer is the server API for RestaurantService service.
// All implementations must embed UnimplementedRestaurantServiceServer
// for forward compatibility.
type RestaurantServiceServer interface {
	CreateRestaurant(context.Context, *CreateRestaurantRequest) (*Restaurant, error)
	GetRestaurants(context.Context, *emptypb.Empty) (*GetRestaurantsResponse, error)
	GetRestaurantById(context.Context, *GetRestaurantByIdRequest) (*Restaurant, error)
	mustEmbedUnimplementedRestaurantServiceServer()
}

// UnimplementedRestaurantServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRestaurantServiceServer struct{}

func (UnimplementedRestaurantServiceServer) CreateRestaurant(context.Context, *CreateRestaurantRequest) (*Restaurant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRestaurant not implemented")
}
func (UnimplementedRestaurantServiceServer) GetRestaurants(context.Context, *emptypb.Empty) (*GetRestaurantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurants not implemented")
}
func (UnimplementedRestaurantServiceServer) GetRestaurantById(context.Context, *GetRestaurantByIdRequest) (*Restaurant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurantById not implemented")
}
func (UnimplementedRestaurantServiceServer) mustEmbedUnimplementedRestaurantServiceServer() {}
func (UnimplementedRestaurantServiceServer) testEmbeddedByValue()                           {}

// UnsafeRestaurantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RestaurantServiceServer will
// result in compilation errors.
type UnsafeRestaurantServiceServer interface {
	mustEmbedUnimplementedRestaurantServiceServer()
}

func RegisterRestaurantServiceServer(s grpc.ServiceRegistrar, srv RestaurantServiceServer) {
	// If the following call pancis, it indicates UnimplementedRestaurantServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RestaurantService_ServiceDesc, srv)
}

func _RestaurantService_CreateRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).CreateRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_CreateRestaurant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).CreateRestaurant(ctx, req.(*CreateRestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_GetRestaurants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).GetRestaurants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_GetRestaurants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).GetRestaurants(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestaurantService_GetRestaurantById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRestaurantByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServiceServer).GetRestaurantById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestaurantService_GetRestaurantById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServiceServer).GetRestaurantById(ctx, req.(*GetRestaurantByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RestaurantService_ServiceDesc is the grpc.ServiceDesc for RestaurantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RestaurantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "catalog.RestaurantService",
	HandlerType: (*RestaurantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRestaurant",
			Handler:    _RestaurantService_CreateRestaurant_Handler,
		},
		{
			MethodName: "GetRestaurants",
			Handler:    _RestaurantService_GetRestaurants_Handler,
		},
		{
			MethodName: "GetRestaurantById",
			Handler:    _RestaurantService_GetRestaurantById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/catalog.proto",
}

const (
	ItemService_CreateItem_FullMethodName             = "/catalog.ItemService/CreateItem"
	ItemService_GetItemsByRestaurantId_FullMethodName = "/catalog.ItemService/GetItemsByRestaurantId"
	ItemService_GetItemById_FullMethodName            = "/catalog.ItemService/GetItemById"
)

// ItemServiceClient is the client API for ItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItemServiceClient interface {
	CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*MenuItem, error)
	GetItemsByRestaurantId(ctx context.Context, in *GetItemsByRestaurantIdRequest, opts ...grpc.CallOption) (*GetItemsByRestaurantIdResponse, error)
	GetItemById(ctx context.Context, in *GetItemByIdRequest, opts ...grpc.CallOption) (*MenuItem, error)
}

type itemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewItemServiceClient(cc grpc.ClientConnInterface) ItemServiceClient {
	return &itemServiceClient{cc}
}

func (c *itemServiceClient) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*MenuItem, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MenuItem)
	err := c.cc.Invoke(ctx, ItemService_CreateItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetItemsByRestaurantId(ctx context.Context, in *GetItemsByRestaurantIdRequest, opts ...grpc.CallOption) (*GetItemsByRestaurantIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetItemsByRestaurantIdResponse)
	err := c.cc.Invoke(ctx, ItemService_GetItemsByRestaurantId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetItemById(ctx context.Context, in *GetItemByIdRequest, opts ...grpc.CallOption) (*MenuItem, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MenuItem)
	err := c.cc.Invoke(ctx, ItemService_GetItemById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemServiceServer is the server API for ItemService service.
// All implementations must embed UnimplementedItemServiceServer
// for forward compatibility.
type ItemServiceServer interface {
	CreateItem(context.Context, *CreateItemRequest) (*MenuItem, error)
	GetItemsByRestaurantId(context.Context, *GetItemsByRestaurantIdRequest) (*GetItemsByRestaurantIdResponse, error)
	GetItemById(context.Context, *GetItemByIdRequest) (*MenuItem, error)
	mustEmbedUnimplementedItemServiceServer()
}

// UnimplementedItemServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedItemServiceServer struct{}

func (UnimplementedItemServiceServer) CreateItem(context.Context, *CreateItemRequest) (*MenuItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (UnimplementedItemServiceServer) GetItemsByRestaurantId(context.Context, *GetItemsByRestaurantIdRequest) (*GetItemsByRestaurantIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemsByRestaurantId not implemented")
}
func (UnimplementedItemServiceServer) GetItemById(context.Context, *GetItemByIdRequest) (*MenuItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemById not implemented")
}
func (UnimplementedItemServiceServer) mustEmbedUnimplementedItemServiceServer() {}
func (UnimplementedItemServiceServer) testEmbeddedByValue()                     {}

// UnsafeItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItemServiceServer will
// result in compilation errors.
type UnsafeItemServiceServer interface {
	mustEmbedUnimplementedItemServiceServer()
}

func RegisterItemServiceServer(s grpc.ServiceRegistrar, srv ItemServiceServer) {
	// If the following call pancis, it indicates UnimplementedItemServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ItemService_ServiceDesc, srv)
}

func _ItemService_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemService_CreateItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).CreateItem(ctx, req.(*CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetItemsByRestaurantId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemsByRestaurantIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetItemsByRestaurantId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemService_GetItemsByRestaurantId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetItemsByRestaurantId(ctx, req.(*GetItemsByRestaurantIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetItemById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetItemById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemService_GetItemById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetItemById(ctx, req.(*GetItemByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ItemService_ServiceDesc is the grpc.ServiceDesc for ItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "catalog.ItemService",
	HandlerType: (*ItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateItem",
			Handler:    _ItemService_CreateItem_Handler,
		},
		{
			MethodName: "GetItemsByRestaurantId",
			Handler:    _ItemService_GetItemsByRestaurantId_Handler,
		},
		{
			MethodName: "GetItemById",
			Handler:    _ItemService_GetItemById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/catalog.proto",
}
