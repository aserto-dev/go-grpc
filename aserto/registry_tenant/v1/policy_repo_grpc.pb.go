// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aserto/registry_tenant/v1/policy_repo.proto

package registry_tenant

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

const (
	PolicyRepo_GetPolicyRepo_FullMethodName         = "/aserto.registry_tenant.v1.PolicyRepo/GetPolicyRepo"
	PolicyRepo_ListPolicyRepos_FullMethodName       = "/aserto.registry_tenant.v1.PolicyRepo/ListPolicyRepos"
	PolicyRepo_ListPublicPolicyRepos_FullMethodName = "/aserto.registry_tenant.v1.PolicyRepo/ListPublicPolicyRepos"
	PolicyRepo_CreatePolicyRepo_FullMethodName      = "/aserto.registry_tenant.v1.PolicyRepo/CreatePolicyRepo"
	PolicyRepo_DeletePolicyRepo_FullMethodName      = "/aserto.registry_tenant.v1.PolicyRepo/DeletePolicyRepo"
	PolicyRepo_UpdatePolicyRepo_FullMethodName      = "/aserto.registry_tenant.v1.PolicyRepo/UpdatePolicyRepo"
)

// PolicyRepoClient is the client API for PolicyRepo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyRepoClient interface {
	GetPolicyRepo(ctx context.Context, in *GetPolicyRepoRequest, opts ...grpc.CallOption) (*GetPolicyRepoResponse, error)
	ListPolicyRepos(ctx context.Context, in *ListPolicyReposRequest, opts ...grpc.CallOption) (*ListPolicyReposResponse, error)
	ListPublicPolicyRepos(ctx context.Context, in *ListPublicPolicyReposRequest, opts ...grpc.CallOption) (*ListPublicPolicyReposResponse, error)
	CreatePolicyRepo(ctx context.Context, in *CreatePolicyRepoRequest, opts ...grpc.CallOption) (*CreatePolicyRepoResponse, error)
	DeletePolicyRepo(ctx context.Context, in *DeletePolicyRepoRequest, opts ...grpc.CallOption) (*DeletePolicyRepoResponse, error)
	UpdatePolicyRepo(ctx context.Context, in *UpdatePolicyRepoRequest, opts ...grpc.CallOption) (*UpdatePolicyRepoResponse, error)
}

type policyRepoClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyRepoClient(cc grpc.ClientConnInterface) PolicyRepoClient {
	return &policyRepoClient{cc}
}

func (c *policyRepoClient) GetPolicyRepo(ctx context.Context, in *GetPolicyRepoRequest, opts ...grpc.CallOption) (*GetPolicyRepoResponse, error) {
	out := new(GetPolicyRepoResponse)
	err := c.cc.Invoke(ctx, PolicyRepo_GetPolicyRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyRepoClient) ListPolicyRepos(ctx context.Context, in *ListPolicyReposRequest, opts ...grpc.CallOption) (*ListPolicyReposResponse, error) {
	out := new(ListPolicyReposResponse)
	err := c.cc.Invoke(ctx, PolicyRepo_ListPolicyRepos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyRepoClient) ListPublicPolicyRepos(ctx context.Context, in *ListPublicPolicyReposRequest, opts ...grpc.CallOption) (*ListPublicPolicyReposResponse, error) {
	out := new(ListPublicPolicyReposResponse)
	err := c.cc.Invoke(ctx, PolicyRepo_ListPublicPolicyRepos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyRepoClient) CreatePolicyRepo(ctx context.Context, in *CreatePolicyRepoRequest, opts ...grpc.CallOption) (*CreatePolicyRepoResponse, error) {
	out := new(CreatePolicyRepoResponse)
	err := c.cc.Invoke(ctx, PolicyRepo_CreatePolicyRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyRepoClient) DeletePolicyRepo(ctx context.Context, in *DeletePolicyRepoRequest, opts ...grpc.CallOption) (*DeletePolicyRepoResponse, error) {
	out := new(DeletePolicyRepoResponse)
	err := c.cc.Invoke(ctx, PolicyRepo_DeletePolicyRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyRepoClient) UpdatePolicyRepo(ctx context.Context, in *UpdatePolicyRepoRequest, opts ...grpc.CallOption) (*UpdatePolicyRepoResponse, error) {
	out := new(UpdatePolicyRepoResponse)
	err := c.cc.Invoke(ctx, PolicyRepo_UpdatePolicyRepo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyRepoServer is the server API for PolicyRepo service.
// All implementations should embed UnimplementedPolicyRepoServer
// for forward compatibility
type PolicyRepoServer interface {
	GetPolicyRepo(context.Context, *GetPolicyRepoRequest) (*GetPolicyRepoResponse, error)
	ListPolicyRepos(context.Context, *ListPolicyReposRequest) (*ListPolicyReposResponse, error)
	ListPublicPolicyRepos(context.Context, *ListPublicPolicyReposRequest) (*ListPublicPolicyReposResponse, error)
	CreatePolicyRepo(context.Context, *CreatePolicyRepoRequest) (*CreatePolicyRepoResponse, error)
	DeletePolicyRepo(context.Context, *DeletePolicyRepoRequest) (*DeletePolicyRepoResponse, error)
	UpdatePolicyRepo(context.Context, *UpdatePolicyRepoRequest) (*UpdatePolicyRepoResponse, error)
}

// UnimplementedPolicyRepoServer should be embedded to have forward compatible implementations.
type UnimplementedPolicyRepoServer struct {
}

func (UnimplementedPolicyRepoServer) GetPolicyRepo(context.Context, *GetPolicyRepoRequest) (*GetPolicyRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicyRepo not implemented")
}
func (UnimplementedPolicyRepoServer) ListPolicyRepos(context.Context, *ListPolicyReposRequest) (*ListPolicyReposResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPolicyRepos not implemented")
}
func (UnimplementedPolicyRepoServer) ListPublicPolicyRepos(context.Context, *ListPublicPolicyReposRequest) (*ListPublicPolicyReposResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPublicPolicyRepos not implemented")
}
func (UnimplementedPolicyRepoServer) CreatePolicyRepo(context.Context, *CreatePolicyRepoRequest) (*CreatePolicyRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolicyRepo not implemented")
}
func (UnimplementedPolicyRepoServer) DeletePolicyRepo(context.Context, *DeletePolicyRepoRequest) (*DeletePolicyRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicyRepo not implemented")
}
func (UnimplementedPolicyRepoServer) UpdatePolicyRepo(context.Context, *UpdatePolicyRepoRequest) (*UpdatePolicyRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePolicyRepo not implemented")
}

// UnsafePolicyRepoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolicyRepoServer will
// result in compilation errors.
type UnsafePolicyRepoServer interface {
	mustEmbedUnimplementedPolicyRepoServer()
}

func RegisterPolicyRepoServer(s grpc.ServiceRegistrar, srv PolicyRepoServer) {
	s.RegisterService(&PolicyRepo_ServiceDesc, srv)
}

func _PolicyRepo_GetPolicyRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyRepoServer).GetPolicyRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyRepo_GetPolicyRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyRepoServer).GetPolicyRepo(ctx, req.(*GetPolicyRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyRepo_ListPolicyRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPolicyReposRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyRepoServer).ListPolicyRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyRepo_ListPolicyRepos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyRepoServer).ListPolicyRepos(ctx, req.(*ListPolicyReposRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyRepo_ListPublicPolicyRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPublicPolicyReposRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyRepoServer).ListPublicPolicyRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyRepo_ListPublicPolicyRepos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyRepoServer).ListPublicPolicyRepos(ctx, req.(*ListPublicPolicyReposRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyRepo_CreatePolicyRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePolicyRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyRepoServer).CreatePolicyRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyRepo_CreatePolicyRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyRepoServer).CreatePolicyRepo(ctx, req.(*CreatePolicyRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyRepo_DeletePolicyRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyRepoServer).DeletePolicyRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyRepo_DeletePolicyRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyRepoServer).DeletePolicyRepo(ctx, req.(*DeletePolicyRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyRepo_UpdatePolicyRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePolicyRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyRepoServer).UpdatePolicyRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyRepo_UpdatePolicyRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyRepoServer).UpdatePolicyRepo(ctx, req.(*UpdatePolicyRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PolicyRepo_ServiceDesc is the grpc.ServiceDesc for PolicyRepo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PolicyRepo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.registry_tenant.v1.PolicyRepo",
	HandlerType: (*PolicyRepoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPolicyRepo",
			Handler:    _PolicyRepo_GetPolicyRepo_Handler,
		},
		{
			MethodName: "ListPolicyRepos",
			Handler:    _PolicyRepo_ListPolicyRepos_Handler,
		},
		{
			MethodName: "ListPublicPolicyRepos",
			Handler:    _PolicyRepo_ListPublicPolicyRepos_Handler,
		},
		{
			MethodName: "CreatePolicyRepo",
			Handler:    _PolicyRepo_CreatePolicyRepo_Handler,
		},
		{
			MethodName: "DeletePolicyRepo",
			Handler:    _PolicyRepo_DeletePolicyRepo_Handler,
		},
		{
			MethodName: "UpdatePolicyRepo",
			Handler:    _PolicyRepo_UpdatePolicyRepo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aserto/registry_tenant/v1/policy_repo.proto",
}
