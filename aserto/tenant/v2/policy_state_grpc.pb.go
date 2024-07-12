// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: aserto/tenant/v2/policy_state.proto

package tenant

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	PolicyState_GetPolicyState_FullMethodName = "/aserto.tenant.v2.PolicyState/GetPolicyState"
	PolicyState_SetPolicyState_FullMethodName = "/aserto.tenant.v2.PolicyState/SetPolicyState"
)

// PolicyStateClient is the client API for PolicyState service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyStateClient interface {
	GetPolicyState(ctx context.Context, in *GetPolicyStateRequest, opts ...grpc.CallOption) (*GetPolicyStateResponse, error)
	SetPolicyState(ctx context.Context, in *SetPolicyStateRequest, opts ...grpc.CallOption) (*SetPolicyStateResponse, error)
}

type policyStateClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyStateClient(cc grpc.ClientConnInterface) PolicyStateClient {
	return &policyStateClient{cc}
}

func (c *policyStateClient) GetPolicyState(ctx context.Context, in *GetPolicyStateRequest, opts ...grpc.CallOption) (*GetPolicyStateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPolicyStateResponse)
	err := c.cc.Invoke(ctx, PolicyState_GetPolicyState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyStateClient) SetPolicyState(ctx context.Context, in *SetPolicyStateRequest, opts ...grpc.CallOption) (*SetPolicyStateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetPolicyStateResponse)
	err := c.cc.Invoke(ctx, PolicyState_SetPolicyState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyStateServer is the server API for PolicyState service.
// All implementations should embed UnimplementedPolicyStateServer
// for forward compatibility
type PolicyStateServer interface {
	GetPolicyState(context.Context, *GetPolicyStateRequest) (*GetPolicyStateResponse, error)
	SetPolicyState(context.Context, *SetPolicyStateRequest) (*SetPolicyStateResponse, error)
}

// UnimplementedPolicyStateServer should be embedded to have forward compatible implementations.
type UnimplementedPolicyStateServer struct {
}

func (UnimplementedPolicyStateServer) GetPolicyState(context.Context, *GetPolicyStateRequest) (*GetPolicyStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicyState not implemented")
}
func (UnimplementedPolicyStateServer) SetPolicyState(context.Context, *SetPolicyStateRequest) (*SetPolicyStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPolicyState not implemented")
}

// UnsafePolicyStateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolicyStateServer will
// result in compilation errors.
type UnsafePolicyStateServer interface {
	mustEmbedUnimplementedPolicyStateServer()
}

func RegisterPolicyStateServer(s grpc.ServiceRegistrar, srv PolicyStateServer) {
	s.RegisterService(&PolicyState_ServiceDesc, srv)
}

func _PolicyState_GetPolicyState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyStateServer).GetPolicyState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyState_GetPolicyState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyStateServer).GetPolicyState(ctx, req.(*GetPolicyStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyState_SetPolicyState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPolicyStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyStateServer).SetPolicyState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyState_SetPolicyState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyStateServer).SetPolicyState(ctx, req.(*SetPolicyStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PolicyState_ServiceDesc is the grpc.ServiceDesc for PolicyState service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PolicyState_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.tenant.v2.PolicyState",
	HandlerType: (*PolicyStateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPolicyState",
			Handler:    _PolicyState_GetPolicyState_Handler,
		},
		{
			MethodName: "SetPolicyState",
			Handler:    _PolicyState_SetPolicyState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aserto/tenant/v2/policy_state.proto",
}
