// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: aserto/tenant/policy/v1/policy.proto

package policy

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
	Policy_ListPolicyRefs_FullMethodName       = "/aserto.tenant.policy.v1.Policy/ListPolicyRefs"
	Policy_CreatePolicyRef_FullMethodName      = "/aserto.tenant.policy.v1.Policy/CreatePolicyRef"
	Policy_DeletePolicyRef_FullMethodName      = "/aserto.tenant.policy.v1.Policy/DeletePolicyRef"
	Policy_UpdatePolicyRef_FullMethodName      = "/aserto.tenant.policy.v1.Policy/UpdatePolicyRef"
	Policy_OPADiscovery_FullMethodName         = "/aserto.tenant.policy.v1.Policy/OPADiscovery"
	Policy_OPAInstanceDiscovery_FullMethodName = "/aserto.tenant.policy.v1.Policy/OPAInstanceDiscovery"
)

// PolicyClient is the client API for Policy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyClient interface {
	ListPolicyRefs(ctx context.Context, in *ListPolicyRefsRequest, opts ...grpc.CallOption) (*ListPolicyRefsResponse, error)
	CreatePolicyRef(ctx context.Context, in *CreatePolicyRefRequest, opts ...grpc.CallOption) (*CreatePolicyRefResponse, error)
	DeletePolicyRef(ctx context.Context, in *DeletePolicyRefRequest, opts ...grpc.CallOption) (*DeletePolicyRefResponse, error)
	UpdatePolicyRef(ctx context.Context, in *UpdatePolicyRefRequest, opts ...grpc.CallOption) (*UpdatePolicyRefResponse, error)
	OPADiscovery(ctx context.Context, in *OPADiscoveryRequest, opts ...grpc.CallOption) (*OPADiscoveryResponse, error)
	OPAInstanceDiscovery(ctx context.Context, in *OPAInstanceDiscoveryRequest, opts ...grpc.CallOption) (*OPAInstanceDiscoveryResponse, error)
}

type policyClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyClient(cc grpc.ClientConnInterface) PolicyClient {
	return &policyClient{cc}
}

func (c *policyClient) ListPolicyRefs(ctx context.Context, in *ListPolicyRefsRequest, opts ...grpc.CallOption) (*ListPolicyRefsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPolicyRefsResponse)
	err := c.cc.Invoke(ctx, Policy_ListPolicyRefs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) CreatePolicyRef(ctx context.Context, in *CreatePolicyRefRequest, opts ...grpc.CallOption) (*CreatePolicyRefResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePolicyRefResponse)
	err := c.cc.Invoke(ctx, Policy_CreatePolicyRef_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) DeletePolicyRef(ctx context.Context, in *DeletePolicyRefRequest, opts ...grpc.CallOption) (*DeletePolicyRefResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePolicyRefResponse)
	err := c.cc.Invoke(ctx, Policy_DeletePolicyRef_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) UpdatePolicyRef(ctx context.Context, in *UpdatePolicyRefRequest, opts ...grpc.CallOption) (*UpdatePolicyRefResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePolicyRefResponse)
	err := c.cc.Invoke(ctx, Policy_UpdatePolicyRef_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) OPADiscovery(ctx context.Context, in *OPADiscoveryRequest, opts ...grpc.CallOption) (*OPADiscoveryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OPADiscoveryResponse)
	err := c.cc.Invoke(ctx, Policy_OPADiscovery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) OPAInstanceDiscovery(ctx context.Context, in *OPAInstanceDiscoveryRequest, opts ...grpc.CallOption) (*OPAInstanceDiscoveryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OPAInstanceDiscoveryResponse)
	err := c.cc.Invoke(ctx, Policy_OPAInstanceDiscovery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyServer is the server API for Policy service.
// All implementations should embed UnimplementedPolicyServer
// for forward compatibility
type PolicyServer interface {
	ListPolicyRefs(context.Context, *ListPolicyRefsRequest) (*ListPolicyRefsResponse, error)
	CreatePolicyRef(context.Context, *CreatePolicyRefRequest) (*CreatePolicyRefResponse, error)
	DeletePolicyRef(context.Context, *DeletePolicyRefRequest) (*DeletePolicyRefResponse, error)
	UpdatePolicyRef(context.Context, *UpdatePolicyRefRequest) (*UpdatePolicyRefResponse, error)
	OPADiscovery(context.Context, *OPADiscoveryRequest) (*OPADiscoveryResponse, error)
	OPAInstanceDiscovery(context.Context, *OPAInstanceDiscoveryRequest) (*OPAInstanceDiscoveryResponse, error)
}

// UnimplementedPolicyServer should be embedded to have forward compatible implementations.
type UnimplementedPolicyServer struct {
}

func (UnimplementedPolicyServer) ListPolicyRefs(context.Context, *ListPolicyRefsRequest) (*ListPolicyRefsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPolicyRefs not implemented")
}
func (UnimplementedPolicyServer) CreatePolicyRef(context.Context, *CreatePolicyRefRequest) (*CreatePolicyRefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolicyRef not implemented")
}
func (UnimplementedPolicyServer) DeletePolicyRef(context.Context, *DeletePolicyRefRequest) (*DeletePolicyRefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicyRef not implemented")
}
func (UnimplementedPolicyServer) UpdatePolicyRef(context.Context, *UpdatePolicyRefRequest) (*UpdatePolicyRefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePolicyRef not implemented")
}
func (UnimplementedPolicyServer) OPADiscovery(context.Context, *OPADiscoveryRequest) (*OPADiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OPADiscovery not implemented")
}
func (UnimplementedPolicyServer) OPAInstanceDiscovery(context.Context, *OPAInstanceDiscoveryRequest) (*OPAInstanceDiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OPAInstanceDiscovery not implemented")
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

func _Policy_ListPolicyRefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPolicyRefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).ListPolicyRefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Policy_ListPolicyRefs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).ListPolicyRefs(ctx, req.(*ListPolicyRefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_CreatePolicyRef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePolicyRefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).CreatePolicyRef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Policy_CreatePolicyRef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).CreatePolicyRef(ctx, req.(*CreatePolicyRefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_DeletePolicyRef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyRefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).DeletePolicyRef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Policy_DeletePolicyRef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).DeletePolicyRef(ctx, req.(*DeletePolicyRefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_UpdatePolicyRef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePolicyRefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).UpdatePolicyRef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Policy_UpdatePolicyRef_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).UpdatePolicyRef(ctx, req.(*UpdatePolicyRefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_OPADiscovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OPADiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).OPADiscovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Policy_OPADiscovery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).OPADiscovery(ctx, req.(*OPADiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_OPAInstanceDiscovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OPAInstanceDiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).OPAInstanceDiscovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Policy_OPAInstanceDiscovery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).OPAInstanceDiscovery(ctx, req.(*OPAInstanceDiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Policy_ServiceDesc is the grpc.ServiceDesc for Policy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Policy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.tenant.policy.v1.Policy",
	HandlerType: (*PolicyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPolicyRefs",
			Handler:    _Policy_ListPolicyRefs_Handler,
		},
		{
			MethodName: "CreatePolicyRef",
			Handler:    _Policy_CreatePolicyRef_Handler,
		},
		{
			MethodName: "DeletePolicyRef",
			Handler:    _Policy_DeletePolicyRef_Handler,
		},
		{
			MethodName: "UpdatePolicyRef",
			Handler:    _Policy_UpdatePolicyRef_Handler,
		},
		{
			MethodName: "OPADiscovery",
			Handler:    _Policy_OPADiscovery_Handler,
		},
		{
			MethodName: "OPAInstanceDiscovery",
			Handler:    _Policy_OPAInstanceDiscovery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aserto/tenant/policy/v1/policy.proto",
}
