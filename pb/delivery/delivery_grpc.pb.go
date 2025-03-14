// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: pb/delivery.proto

package delivery

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DeliveryService_AssignDelivery_FullMethodName       = "/delivery.DeliveryService/AssignDelivery"
	DeliveryService_UpdateDeliveryStatus_FullMethodName = "/delivery.DeliveryService/UpdateDeliveryStatus"
)

// DeliveryServiceClient is the client API for DeliveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeliveryServiceClient interface {
	AssignDelivery(ctx context.Context, in *AssignDeliveryRequest, opts ...grpc.CallOption) (*AssignDeliveryResponse, error)
	UpdateDeliveryStatus(ctx context.Context, in *UpdateDeliveryStatusRequest, opts ...grpc.CallOption) (*UpdateDeliveryStatusResponse, error)
}

type deliveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeliveryServiceClient(cc grpc.ClientConnInterface) DeliveryServiceClient {
	return &deliveryServiceClient{cc}
}

func (c *deliveryServiceClient) AssignDelivery(ctx context.Context, in *AssignDeliveryRequest, opts ...grpc.CallOption) (*AssignDeliveryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AssignDeliveryResponse)
	err := c.cc.Invoke(ctx, DeliveryService_AssignDelivery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryServiceClient) UpdateDeliveryStatus(ctx context.Context, in *UpdateDeliveryStatusRequest, opts ...grpc.CallOption) (*UpdateDeliveryStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDeliveryStatusResponse)
	err := c.cc.Invoke(ctx, DeliveryService_UpdateDeliveryStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeliveryServiceServer is the server API for DeliveryService service.
// All implementations must embed UnimplementedDeliveryServiceServer
// for forward compatibility.
type DeliveryServiceServer interface {
	AssignDelivery(context.Context, *AssignDeliveryRequest) (*AssignDeliveryResponse, error)
	UpdateDeliveryStatus(context.Context, *UpdateDeliveryStatusRequest) (*UpdateDeliveryStatusResponse, error)
	mustEmbedUnimplementedDeliveryServiceServer()
}

// UnimplementedDeliveryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDeliveryServiceServer struct{}

func (UnimplementedDeliveryServiceServer) AssignDelivery(context.Context, *AssignDeliveryRequest) (*AssignDeliveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignDelivery not implemented")
}
func (UnimplementedDeliveryServiceServer) UpdateDeliveryStatus(context.Context, *UpdateDeliveryStatusRequest) (*UpdateDeliveryStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeliveryStatus not implemented")
}
func (UnimplementedDeliveryServiceServer) mustEmbedUnimplementedDeliveryServiceServer() {}
func (UnimplementedDeliveryServiceServer) testEmbeddedByValue()                         {}

// UnsafeDeliveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeliveryServiceServer will
// result in compilation errors.
type UnsafeDeliveryServiceServer interface {
	mustEmbedUnimplementedDeliveryServiceServer()
}

func RegisterDeliveryServiceServer(s grpc.ServiceRegistrar, srv DeliveryServiceServer) {
	// If the following call pancis, it indicates UnimplementedDeliveryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DeliveryService_ServiceDesc, srv)
}

func _DeliveryService_AssignDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignDeliveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).AssignDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeliveryService_AssignDelivery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).AssignDelivery(ctx, req.(*AssignDeliveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeliveryService_UpdateDeliveryStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeliveryStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServiceServer).UpdateDeliveryStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeliveryService_UpdateDeliveryStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServiceServer).UpdateDeliveryStatus(ctx, req.(*UpdateDeliveryStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeliveryService_ServiceDesc is the grpc.ServiceDesc for DeliveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeliveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "delivery.DeliveryService",
	HandlerType: (*DeliveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignDelivery",
			Handler:    _DeliveryService_AssignDelivery_Handler,
		},
		{
			MethodName: "UpdateDeliveryStatus",
			Handler:    _DeliveryService_UpdateDeliveryStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/delivery.proto",
}
