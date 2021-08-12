// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package policy

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PolicyClient is the client API for Policy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyClient interface {
	ListPolicies(ctx context.Context, in *ListPoliciesRequest, opts ...grpc.CallOption) (*ListPoliciesResponse, error)
	GetPolicies(ctx context.Context, in *GetPoliciesRequest, opts ...grpc.CallOption) (*GetPoliciesResponse, error)
	GetModule(ctx context.Context, in *GetModuleRequest, opts ...grpc.CallOption) (*GetModuleResponse, error)
}

type policyClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyClient(cc grpc.ClientConnInterface) PolicyClient {
	return &policyClient{cc}
}

func (c *policyClient) ListPolicies(ctx context.Context, in *ListPoliciesRequest, opts ...grpc.CallOption) (*ListPoliciesResponse, error) {
	out := new(ListPoliciesResponse)
	err := c.cc.Invoke(ctx, "/aserto.authorizer.policy.v1.Policy/ListPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) GetPolicies(ctx context.Context, in *GetPoliciesRequest, opts ...grpc.CallOption) (*GetPoliciesResponse, error) {
	out := new(GetPoliciesResponse)
	err := c.cc.Invoke(ctx, "/aserto.authorizer.policy.v1.Policy/GetPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) GetModule(ctx context.Context, in *GetModuleRequest, opts ...grpc.CallOption) (*GetModuleResponse, error) {
	out := new(GetModuleResponse)
	err := c.cc.Invoke(ctx, "/aserto.authorizer.policy.v1.Policy/GetModule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyServer is the server API for Policy service.
// All implementations should embed UnimplementedPolicyServer
// for forward compatibility
type PolicyServer interface {
	ListPolicies(context.Context, *ListPoliciesRequest) (*ListPoliciesResponse, error)
	GetPolicies(context.Context, *GetPoliciesRequest) (*GetPoliciesResponse, error)
	GetModule(context.Context, *GetModuleRequest) (*GetModuleResponse, error)
}

// UnimplementedPolicyServer should be embedded to have forward compatible implementations.
type UnimplementedPolicyServer struct {
}

func (UnimplementedPolicyServer) ListPolicies(context.Context, *ListPoliciesRequest) (*ListPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPolicies not implemented")
}
func (UnimplementedPolicyServer) GetPolicies(context.Context, *GetPoliciesRequest) (*GetPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicies not implemented")
}
func (UnimplementedPolicyServer) GetModule(context.Context, *GetModuleRequest) (*GetModuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModule not implemented")
}

// UnsafePolicyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolicyServer will
// result in compilation errors.
type UnsafePolicyServer interface {
	mustEmbedUnimplementedPolicyServer()
}

func RegisterPolicyServer(s grpc.ServiceRegistrar, srv PolicyServer) {
	s.RegisterService(&Policy_ServiceDesc, srv)
}

func _Policy_ListPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).ListPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.authorizer.policy.v1.Policy/ListPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).ListPolicies(ctx, req.(*ListPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_GetPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).GetPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.authorizer.policy.v1.Policy/GetPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).GetPolicies(ctx, req.(*GetPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_GetModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).GetModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.authorizer.policy.v1.Policy/GetModule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).GetModule(ctx, req.(*GetModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Policy_ServiceDesc is the grpc.ServiceDesc for Policy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Policy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.authorizer.policy.v1.Policy",
	HandlerType: (*PolicyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPolicies",
			Handler:    _Policy_ListPolicies_Handler,
		},
		{
			MethodName: "GetPolicies",
			Handler:    _Policy_GetPolicies_Handler,
		},
		{
			MethodName: "GetModule",
			Handler:    _Policy_GetModule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aserto/authorizer/policy/v1/policy.proto",
}
