// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: water_potability.proto

package rpc

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
	WaterPotabilityService_PredictWaterPotability_FullMethodName = "/water_potability.WaterPotabilityService/PredictWaterPotability"
)

// WaterPotabilityServiceClient is the client API for WaterPotabilityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WaterPotabilityServiceClient interface {
	PredictWaterPotability(ctx context.Context, in *PredictWaterPotabilityRequest, opts ...grpc.CallOption) (*PredictWaterPotabilityResponse, error)
}

type waterPotabilityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWaterPotabilityServiceClient(cc grpc.ClientConnInterface) WaterPotabilityServiceClient {
	return &waterPotabilityServiceClient{cc}
}

func (c *waterPotabilityServiceClient) PredictWaterPotability(ctx context.Context, in *PredictWaterPotabilityRequest, opts ...grpc.CallOption) (*PredictWaterPotabilityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PredictWaterPotabilityResponse)
	err := c.cc.Invoke(ctx, WaterPotabilityService_PredictWaterPotability_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WaterPotabilityServiceServer is the server API for WaterPotabilityService service.
// All implementations must embed UnimplementedWaterPotabilityServiceServer
// for forward compatibility.
type WaterPotabilityServiceServer interface {
	PredictWaterPotability(context.Context, *PredictWaterPotabilityRequest) (*PredictWaterPotabilityResponse, error)
	mustEmbedUnimplementedWaterPotabilityServiceServer()
}

// UnimplementedWaterPotabilityServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWaterPotabilityServiceServer struct{}

func (UnimplementedWaterPotabilityServiceServer) PredictWaterPotability(context.Context, *PredictWaterPotabilityRequest) (*PredictWaterPotabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PredictWaterPotability not implemented")
}
func (UnimplementedWaterPotabilityServiceServer) mustEmbedUnimplementedWaterPotabilityServiceServer() {
}
func (UnimplementedWaterPotabilityServiceServer) testEmbeddedByValue() {}

// UnsafeWaterPotabilityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WaterPotabilityServiceServer will
// result in compilation errors.
type UnsafeWaterPotabilityServiceServer interface {
	mustEmbedUnimplementedWaterPotabilityServiceServer()
}

func RegisterWaterPotabilityServiceServer(s grpc.ServiceRegistrar, srv WaterPotabilityServiceServer) {
	// If the following call pancis, it indicates UnimplementedWaterPotabilityServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WaterPotabilityService_ServiceDesc, srv)
}

func _WaterPotabilityService_PredictWaterPotability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictWaterPotabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaterPotabilityServiceServer).PredictWaterPotability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WaterPotabilityService_PredictWaterPotability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaterPotabilityServiceServer).PredictWaterPotability(ctx, req.(*PredictWaterPotabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WaterPotabilityService_ServiceDesc is the grpc.ServiceDesc for WaterPotabilityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WaterPotabilityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "water_potability.WaterPotabilityService",
	HandlerType: (*WaterPotabilityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PredictWaterPotability",
			Handler:    _WaterPotabilityService_PredictWaterPotability_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "water_potability.proto",
}
