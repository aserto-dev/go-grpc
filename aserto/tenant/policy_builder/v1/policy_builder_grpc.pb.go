// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package policy_builder

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

// PolicyBuilderClient is the client API for PolicyBuilder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyBuilderClient interface {
	ListPolicyBuilders(ctx context.Context, in *ListPolicyBuildersRequest, opts ...grpc.CallOption) (*ListPolicyBuildersResponse, error)
	CreatePolicyBuilder(ctx context.Context, in *CreatePolicyBuilderRequest, opts ...grpc.CallOption) (*CreatePolicyBuilderResponse, error)
	DeletePolicyBuilder(ctx context.Context, in *DeletePolicyBuilderRequest, opts ...grpc.CallOption) (*DeletePolicyBuilderResponse, error)
}

type policyBuilderClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyBuilderClient(cc grpc.ClientConnInterface) PolicyBuilderClient {
	return &policyBuilderClient{cc}
}

func (c *policyBuilderClient) ListPolicyBuilders(ctx context.Context, in *ListPolicyBuildersRequest, opts ...grpc.CallOption) (*ListPolicyBuildersResponse, error) {
	out := new(ListPolicyBuildersResponse)
	err := c.cc.Invoke(ctx, "/aserto.tenant.policy_builder.v1.PolicyBuilder/ListPolicyBuilders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyBuilderClient) CreatePolicyBuilder(ctx context.Context, in *CreatePolicyBuilderRequest, opts ...grpc.CallOption) (*CreatePolicyBuilderResponse, error) {
	out := new(CreatePolicyBuilderResponse)
	err := c.cc.Invoke(ctx, "/aserto.tenant.policy_builder.v1.PolicyBuilder/CreatePolicyBuilder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyBuilderClient) DeletePolicyBuilder(ctx context.Context, in *DeletePolicyBuilderRequest, opts ...grpc.CallOption) (*DeletePolicyBuilderResponse, error) {
	out := new(DeletePolicyBuilderResponse)
	err := c.cc.Invoke(ctx, "/aserto.tenant.policy_builder.v1.PolicyBuilder/DeletePolicyBuilder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyBuilderServer is the server API for PolicyBuilder service.
// All implementations should embed UnimplementedPolicyBuilderServer
// for forward compatibility
type PolicyBuilderServer interface {
	ListPolicyBuilders(context.Context, *ListPolicyBuildersRequest) (*ListPolicyBuildersResponse, error)
	CreatePolicyBuilder(context.Context, *CreatePolicyBuilderRequest) (*CreatePolicyBuilderResponse, error)
	DeletePolicyBuilder(context.Context, *DeletePolicyBuilderRequest) (*DeletePolicyBuilderResponse, error)
}

// UnimplementedPolicyBuilderServer should be embedded to have forward compatible implementations.
type UnimplementedPolicyBuilderServer struct {
}

func (UnimplementedPolicyBuilderServer) ListPolicyBuilders(context.Context, *ListPolicyBuildersRequest) (*ListPolicyBuildersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPolicyBuilders not implemented")
}
func (UnimplementedPolicyBuilderServer) CreatePolicyBuilder(context.Context, *CreatePolicyBuilderRequest) (*CreatePolicyBuilderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolicyBuilder not implemented")
}
func (UnimplementedPolicyBuilderServer) DeletePolicyBuilder(context.Context, *DeletePolicyBuilderRequest) (*DeletePolicyBuilderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicyBuilder not implemented")
}

// UnsafePolicyBuilderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolicyBuilderServer will
// result in compilation errors.
type UnsafePolicyBuilderServer interface {
	mustEmbedUnimplementedPolicyBuilderServer()
}

func RegisterPolicyBuilderServer(s grpc.ServiceRegistrar, srv PolicyBuilderServer) {
	s.RegisterService(&PolicyBuilder_ServiceDesc, srv)
}

func _PolicyBuilder_ListPolicyBuilders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPolicyBuildersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyBuilderServer).ListPolicyBuilders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.tenant.policy_builder.v1.PolicyBuilder/ListPolicyBuilders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyBuilderServer).ListPolicyBuilders(ctx, req.(*ListPolicyBuildersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyBuilder_CreatePolicyBuilder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePolicyBuilderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyBuilderServer).CreatePolicyBuilder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.tenant.policy_builder.v1.PolicyBuilder/CreatePolicyBuilder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyBuilderServer).CreatePolicyBuilder(ctx, req.(*CreatePolicyBuilderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyBuilder_DeletePolicyBuilder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyBuilderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyBuilderServer).DeletePolicyBuilder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.tenant.policy_builder.v1.PolicyBuilder/DeletePolicyBuilder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyBuilderServer).DeletePolicyBuilder(ctx, req.(*DeletePolicyBuilderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PolicyBuilder_ServiceDesc is the grpc.ServiceDesc for PolicyBuilder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PolicyBuilder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.tenant.policy_builder.v1.PolicyBuilder",
	HandlerType: (*PolicyBuilderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPolicyBuilders",
			Handler:    _PolicyBuilder_ListPolicyBuilders_Handler,
		},
		{
			MethodName: "CreatePolicyBuilder",
			Handler:    _PolicyBuilder_CreatePolicyBuilder_Handler,
		},
		{
			MethodName: "DeletePolicyBuilder",
			Handler:    _PolicyBuilder_DeletePolicyBuilder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aserto/tenant/policy_builder/v1/policy_builder.proto",
}