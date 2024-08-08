// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: aserto/management/v2/control_plane.proto

package management

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
	ControlPlane_ListInstanceRegistrations_FullMethodName = "/aserto.management.v2.ControlPlane/ListInstanceRegistrations"
	ControlPlane_ExecCommand_FullMethodName               = "/aserto.management.v2.ControlPlane/ExecCommand"
)

// ControlPlaneClient is the client API for ControlPlane service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControlPlaneClient interface {
	ListInstanceRegistrations(ctx context.Context, in *ListInstanceRegistrationsRequest, opts ...grpc.CallOption) (*ListInstanceRegistrationsResponse, error)
	ExecCommand(ctx context.Context, in *ExecCommandRequest, opts ...grpc.CallOption) (*ExecCommandResponse, error)
}

type controlPlaneClient struct {
	cc grpc.ClientConnInterface
}

func NewControlPlaneClient(cc grpc.ClientConnInterface) ControlPlaneClient {
	return &controlPlaneClient{cc}
}

func (c *controlPlaneClient) ListInstanceRegistrations(ctx context.Context, in *ListInstanceRegistrationsRequest, opts ...grpc.CallOption) (*ListInstanceRegistrationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListInstanceRegistrationsResponse)
	err := c.cc.Invoke(ctx, ControlPlane_ListInstanceRegistrations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlPlaneClient) ExecCommand(ctx context.Context, in *ExecCommandRequest, opts ...grpc.CallOption) (*ExecCommandResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExecCommandResponse)
	err := c.cc.Invoke(ctx, ControlPlane_ExecCommand_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlPlaneServer is the server API for ControlPlane service.
// All implementations should embed UnimplementedControlPlaneServer
// for forward compatibility.
type ControlPlaneServer interface {
	ListInstanceRegistrations(context.Context, *ListInstanceRegistrationsRequest) (*ListInstanceRegistrationsResponse, error)
	ExecCommand(context.Context, *ExecCommandRequest) (*ExecCommandResponse, error)
}

// UnimplementedControlPlaneServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedControlPlaneServer struct{}

func (UnimplementedControlPlaneServer) ListInstanceRegistrations(context.Context, *ListInstanceRegistrationsRequest) (*ListInstanceRegistrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstanceRegistrations not implemented")
}
func (UnimplementedControlPlaneServer) ExecCommand(context.Context, *ExecCommandRequest) (*ExecCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecCommand not implemented")
}
func (UnimplementedControlPlaneServer) testEmbeddedByValue() {}

// UnsafeControlPlaneServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControlPlaneServer will
// result in compilation errors.
type UnsafeControlPlaneServer interface {
	mustEmbedUnimplementedControlPlaneServer()
}

func RegisterControlPlaneServer(s grpc.ServiceRegistrar, srv ControlPlaneServer) {
	// If the following call pancis, it indicates UnimplementedControlPlaneServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ControlPlane_ServiceDesc, srv)
}

func _ControlPlane_ListInstanceRegistrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstanceRegistrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneServer).ListInstanceRegistrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlPlane_ListInstanceRegistrations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneServer).ListInstanceRegistrations(ctx, req.(*ListInstanceRegistrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlPlane_ExecCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlPlaneServer).ExecCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlPlane_ExecCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlPlaneServer).ExecCommand(ctx, req.(*ExecCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ControlPlane_ServiceDesc is the grpc.ServiceDesc for ControlPlane service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ControlPlane_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.management.v2.ControlPlane",
	HandlerType: (*ControlPlaneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListInstanceRegistrations",
			Handler:    _ControlPlane_ListInstanceRegistrations_Handler,
		},
		{
			MethodName: "ExecCommand",
			Handler:    _ControlPlane_ExecCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aserto/management/v2/control_plane.proto",
}
