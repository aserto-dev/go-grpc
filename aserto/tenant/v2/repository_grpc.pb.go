// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: aserto/tenant/v2/repository.proto

package tenant

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
	Repository_GetRepository_FullMethodName    = "/aserto.tenant.v2.Repository/GetRepository"
	Repository_CreateRepository_FullMethodName = "/aserto.tenant.v2.Repository/CreateRepository"
	Repository_DeleteRepository_FullMethodName = "/aserto.tenant.v2.Repository/DeleteRepository"
	Repository_UpdateRepository_FullMethodName = "/aserto.tenant.v2.Repository/UpdateRepository"
)

// RepositoryClient is the client API for Repository service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepositoryClient interface {
	GetRepository(ctx context.Context, in *GetRepositoryRequest, opts ...grpc.CallOption) (*GetRepositoryResponse, error)
	CreateRepository(ctx context.Context, in *CreateRepositoryRequest, opts ...grpc.CallOption) (*CreateRepositoryResponse, error)
	DeleteRepository(ctx context.Context, in *DeleteRepositoryRequest, opts ...grpc.CallOption) (*DeleteRepositoryResponse, error)
	UpdateRepository(ctx context.Context, in *UpdateRepositoryRequest, opts ...grpc.CallOption) (*UpdateRepositoryResponse, error)
}

type repositoryClient struct {
	cc grpc.ClientConnInterface
}

func NewRepositoryClient(cc grpc.ClientConnInterface) RepositoryClient {
	return &repositoryClient{cc}
}

func (c *repositoryClient) GetRepository(ctx context.Context, in *GetRepositoryRequest, opts ...grpc.CallOption) (*GetRepositoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRepositoryResponse)
	err := c.cc.Invoke(ctx, Repository_GetRepository_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) CreateRepository(ctx context.Context, in *CreateRepositoryRequest, opts ...grpc.CallOption) (*CreateRepositoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRepositoryResponse)
	err := c.cc.Invoke(ctx, Repository_CreateRepository_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) DeleteRepository(ctx context.Context, in *DeleteRepositoryRequest, opts ...grpc.CallOption) (*DeleteRepositoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRepositoryResponse)
	err := c.cc.Invoke(ctx, Repository_DeleteRepository_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) UpdateRepository(ctx context.Context, in *UpdateRepositoryRequest, opts ...grpc.CallOption) (*UpdateRepositoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateRepositoryResponse)
	err := c.cc.Invoke(ctx, Repository_UpdateRepository_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryServer is the server API for Repository service.
// All implementations should embed UnimplementedRepositoryServer
// for forward compatibility.
type RepositoryServer interface {
	GetRepository(context.Context, *GetRepositoryRequest) (*GetRepositoryResponse, error)
	CreateRepository(context.Context, *CreateRepositoryRequest) (*CreateRepositoryResponse, error)
	DeleteRepository(context.Context, *DeleteRepositoryRequest) (*DeleteRepositoryResponse, error)
	UpdateRepository(context.Context, *UpdateRepositoryRequest) (*UpdateRepositoryResponse, error)
}

// UnimplementedRepositoryServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRepositoryServer struct{}

func (UnimplementedRepositoryServer) GetRepository(context.Context, *GetRepositoryRequest) (*GetRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepository not implemented")
}
func (UnimplementedRepositoryServer) CreateRepository(context.Context, *CreateRepositoryRequest) (*CreateRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepository not implemented")
}
func (UnimplementedRepositoryServer) DeleteRepository(context.Context, *DeleteRepositoryRequest) (*DeleteRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRepository not implemented")
}
func (UnimplementedRepositoryServer) UpdateRepository(context.Context, *UpdateRepositoryRequest) (*UpdateRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRepository not implemented")
}
func (UnimplementedRepositoryServer) testEmbeddedByValue() {}

// UnsafeRepositoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepositoryServer will
// result in compilation errors.
type UnsafeRepositoryServer interface {
	mustEmbedUnimplementedRepositoryServer()
}

func RegisterRepositoryServer(s grpc.ServiceRegistrar, srv RepositoryServer) {
	// If the following call pancis, it indicates UnimplementedRepositoryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Repository_ServiceDesc, srv)
}

func _Repository_GetRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).GetRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repository_GetRepository_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).GetRepository(ctx, req.(*GetRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_CreateRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).CreateRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repository_CreateRepository_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).CreateRepository(ctx, req.(*CreateRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_DeleteRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).DeleteRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repository_DeleteRepository_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).DeleteRepository(ctx, req.(*DeleteRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_UpdateRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).UpdateRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repository_UpdateRepository_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).UpdateRepository(ctx, req.(*UpdateRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Repository_ServiceDesc is the grpc.ServiceDesc for Repository service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Repository_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.tenant.v2.Repository",
	HandlerType: (*RepositoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRepository",
			Handler:    _Repository_GetRepository_Handler,
		},
		{
			MethodName: "CreateRepository",
			Handler:    _Repository_CreateRepository_Handler,
		},
		{
			MethodName: "DeleteRepository",
			Handler:    _Repository_DeleteRepository_Handler,
		},
		{
			MethodName: "UpdateRepository",
			Handler:    _Repository_UpdateRepository_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aserto/tenant/v2/repository.proto",
}
