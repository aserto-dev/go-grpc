// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: aserto/funnel/v1/funnel.proto

package funnel

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
	Funnel_RunWorkflow_FullMethodName   = "/aserto.funnel.v1.Funnel/RunWorkflow"
	Funnel_StartWorkflow_FullMethodName = "/aserto.funnel.v1.Funnel/StartWorkflow"
	Funnel_StopWorkflow_FullMethodName  = "/aserto.funnel.v1.Funnel/StopWorkflow"
)

// FunnelClient is the client API for Funnel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FunnelClient interface {
	RunWorkflow(ctx context.Context, in *RunWorkflowRequest, opts ...grpc.CallOption) (*RunWorkflowResponse, error)
	StartWorkflow(ctx context.Context, in *StartWorkflowRequest, opts ...grpc.CallOption) (*StartWorkflowResponse, error)
	StopWorkflow(ctx context.Context, in *StopWorkflowRequest, opts ...grpc.CallOption) (*StopWorkflowResponse, error)
}

type funnelClient struct {
	cc grpc.ClientConnInterface
}

func NewFunnelClient(cc grpc.ClientConnInterface) FunnelClient {
	return &funnelClient{cc}
}

func (c *funnelClient) RunWorkflow(ctx context.Context, in *RunWorkflowRequest, opts ...grpc.CallOption) (*RunWorkflowResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RunWorkflowResponse)
	err := c.cc.Invoke(ctx, Funnel_RunWorkflow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *funnelClient) StartWorkflow(ctx context.Context, in *StartWorkflowRequest, opts ...grpc.CallOption) (*StartWorkflowResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartWorkflowResponse)
	err := c.cc.Invoke(ctx, Funnel_StartWorkflow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *funnelClient) StopWorkflow(ctx context.Context, in *StopWorkflowRequest, opts ...grpc.CallOption) (*StopWorkflowResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StopWorkflowResponse)
	err := c.cc.Invoke(ctx, Funnel_StopWorkflow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FunnelServer is the server API for Funnel service.
// All implementations should embed UnimplementedFunnelServer
// for forward compatibility.
type FunnelServer interface {
	RunWorkflow(context.Context, *RunWorkflowRequest) (*RunWorkflowResponse, error)
	StartWorkflow(context.Context, *StartWorkflowRequest) (*StartWorkflowResponse, error)
	StopWorkflow(context.Context, *StopWorkflowRequest) (*StopWorkflowResponse, error)
}

// UnimplementedFunnelServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFunnelServer struct{}

func (UnimplementedFunnelServer) RunWorkflow(context.Context, *RunWorkflowRequest) (*RunWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunWorkflow not implemented")
}
func (UnimplementedFunnelServer) StartWorkflow(context.Context, *StartWorkflowRequest) (*StartWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartWorkflow not implemented")
}
func (UnimplementedFunnelServer) StopWorkflow(context.Context, *StopWorkflowRequest) (*StopWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopWorkflow not implemented")
}
func (UnimplementedFunnelServer) testEmbeddedByValue() {}

// UnsafeFunnelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FunnelServer will
// result in compilation errors.
type UnsafeFunnelServer interface {
	mustEmbedUnimplementedFunnelServer()
}

func RegisterFunnelServer(s grpc.ServiceRegistrar, srv FunnelServer) {
	// If the following call pancis, it indicates UnimplementedFunnelServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Funnel_ServiceDesc, srv)
}

func _Funnel_RunWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunnelServer).RunWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Funnel_RunWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunnelServer).RunWorkflow(ctx, req.(*RunWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Funnel_StartWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunnelServer).StartWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Funnel_StartWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunnelServer).StartWorkflow(ctx, req.(*StartWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Funnel_StopWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FunnelServer).StopWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Funnel_StopWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FunnelServer).StopWorkflow(ctx, req.(*StopWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Funnel_ServiceDesc is the grpc.ServiceDesc for Funnel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Funnel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.funnel.v1.Funnel",
	HandlerType: (*FunnelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunWorkflow",
			Handler:    _Funnel_RunWorkflow_Handler,
		},
		{
			MethodName: "StartWorkflow",
			Handler:    _Funnel_StartWorkflow_Handler,
		},
		{
			MethodName: "StopWorkflow",
			Handler:    _Funnel_StopWorkflow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aserto/funnel/v1/funnel.proto",
}
