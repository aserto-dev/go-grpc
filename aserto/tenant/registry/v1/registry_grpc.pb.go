// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package registry

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

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistryClient interface {
	ListOrgs(ctx context.Context, in *ListOrgsRequest, opts ...grpc.CallOption) (*ListOrgsResponse, error)
	ListPolicyRepos(ctx context.Context, in *ListPolicyReposRequest, opts ...grpc.CallOption) (*ListPolicyReposResponse, error)
	DeletePolicyRepo(ctx context.Context, in *DeletePolicyRepoRequest, opts ...grpc.CallOption) (*DeletePolicyRepoResponse, error)
	ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsResponse, error)
	GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*GetTagResponse, error)
	CreatePolicyRepo(ctx context.Context, in *CreatePolicyRepoRequest, opts ...grpc.CallOption) (*CreatePolicyRepoResponse, error)
}

type registryClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryClient(cc grpc.ClientConnInterface) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) ListOrgs(ctx context.Context, in *ListOrgsRequest, opts ...grpc.CallOption) (*ListOrgsResponse, error) {
	out := new(ListOrgsResponse)
	err := c.cc.Invoke(ctx, "/aserto.tenant.registry.v1.Registry/ListOrgs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListPolicyRepos(ctx context.Context, in *ListPolicyReposRequest, opts ...grpc.CallOption) (*ListPolicyReposResponse, error) {
	out := new(ListPolicyReposResponse)
	err := c.cc.Invoke(ctx, "/aserto.tenant.registry.v1.Registry/ListPolicyRepos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) DeletePolicyRepo(ctx context.Context, in *DeletePolicyRepoRequest, opts ...grpc.CallOption) (*DeletePolicyRepoResponse, error) {
	out := new(DeletePolicyRepoResponse)
	err := c.cc.Invoke(ctx, "/aserto.tenant.registry.v1.Registry/DeletePolicyRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListTags(ctx context.Context, in *ListTagsRequest, opts ...grpc.CallOption) (*ListTagsResponse, error) {
	out := new(ListTagsResponse)
	err := c.cc.Invoke(ctx, "/aserto.tenant.registry.v1.Registry/ListTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetTag(ctx context.Context, in *GetTagRequest, opts ...grpc.CallOption) (*GetTagResponse, error) {
	out := new(GetTagResponse)
	err := c.cc.Invoke(ctx, "/aserto.tenant.registry.v1.Registry/GetTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) CreatePolicyRepo(ctx context.Context, in *CreatePolicyRepoRequest, opts ...grpc.CallOption) (*CreatePolicyRepoResponse, error) {
	out := new(CreatePolicyRepoResponse)
	err := c.cc.Invoke(ctx, "/aserto.tenant.registry.v1.Registry/CreatePolicyRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
// All implementations should embed UnimplementedRegistryServer
// for forward compatibility
type RegistryServer interface {
	ListOrgs(context.Context, *ListOrgsRequest) (*ListOrgsResponse, error)
	ListPolicyRepos(context.Context, *ListPolicyReposRequest) (*ListPolicyReposResponse, error)
	DeletePolicyRepo(context.Context, *DeletePolicyRepoRequest) (*DeletePolicyRepoResponse, error)
	ListTags(context.Context, *ListTagsRequest) (*ListTagsResponse, error)
	GetTag(context.Context, *GetTagRequest) (*GetTagResponse, error)
	CreatePolicyRepo(context.Context, *CreatePolicyRepoRequest) (*CreatePolicyRepoResponse, error)
}

// UnimplementedRegistryServer should be embedded to have forward compatible implementations.
type UnimplementedRegistryServer struct {
}

func (UnimplementedRegistryServer) ListOrgs(context.Context, *ListOrgsRequest) (*ListOrgsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrgs not implemented")
}
func (UnimplementedRegistryServer) ListPolicyRepos(context.Context, *ListPolicyReposRequest) (*ListPolicyReposResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPolicyRepos not implemented")
}
func (UnimplementedRegistryServer) DeletePolicyRepo(context.Context, *DeletePolicyRepoRequest) (*DeletePolicyRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicyRepo not implemented")
}
func (UnimplementedRegistryServer) ListTags(context.Context, *ListTagsRequest) (*ListTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTags not implemented")
}
func (UnimplementedRegistryServer) GetTag(context.Context, *GetTagRequest) (*GetTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTag not implemented")
}
func (UnimplementedRegistryServer) CreatePolicyRepo(context.Context, *CreatePolicyRepoRequest) (*CreatePolicyRepoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolicyRepo not implemented")
}

// UnsafeRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistryServer will
// result in compilation errors.
type UnsafeRegistryServer interface {
	mustEmbedUnimplementedRegistryServer()
}

func RegisterRegistryServer(s grpc.ServiceRegistrar, srv RegistryServer) {
	s.RegisterService(&Registry_ServiceDesc, srv)
}

func _Registry_ListOrgs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrgsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListOrgs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.tenant.registry.v1.Registry/ListOrgs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListOrgs(ctx, req.(*ListOrgsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListPolicyRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPolicyReposRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListPolicyRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.tenant.registry.v1.Registry/ListPolicyRepos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListPolicyRepos(ctx, req.(*ListPolicyReposRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_DeletePolicyRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).DeletePolicyRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.tenant.registry.v1.Registry/DeletePolicyRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).DeletePolicyRepo(ctx, req.(*DeletePolicyRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.tenant.registry.v1.Registry/ListTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListTags(ctx, req.(*ListTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.tenant.registry.v1.Registry/GetTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetTag(ctx, req.(*GetTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_CreatePolicyRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePolicyRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).CreatePolicyRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.tenant.registry.v1.Registry/CreatePolicyRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).CreatePolicyRepo(ctx, req.(*CreatePolicyRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Registry_ServiceDesc is the grpc.ServiceDesc for Registry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Registry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.tenant.registry.v1.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListOrgs",
			Handler:    _Registry_ListOrgs_Handler,
		},
		{
			MethodName: "ListPolicyRepos",
			Handler:    _Registry_ListPolicyRepos_Handler,
		},
		{
			MethodName: "DeletePolicyRepo",
			Handler:    _Registry_DeletePolicyRepo_Handler,
		},
		{
			MethodName: "ListTags",
			Handler:    _Registry_ListTags_Handler,
		},
		{
			MethodName: "GetTag",
			Handler:    _Registry_GetTag_Handler,
		},
		{
			MethodName: "CreatePolicyRepo",
			Handler:    _Registry_CreatePolicyRepo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aserto/tenant/registry/v1/registry.proto",
}
