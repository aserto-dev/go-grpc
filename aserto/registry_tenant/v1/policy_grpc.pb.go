// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// PolicyClient is the client API for Policy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyClient interface {
	GetPolicyImage(ctx context.Context, in *GetPolicyImageRequest, opts ...grpc.CallOption) (*GetPolicyImageResponse, error)
	ListPolicyImages(ctx context.Context, in *ListPolicyImagesRequest, opts ...grpc.CallOption) (*ListPolicyImagesResponse, error)
	ListPublicPolicyImages(ctx context.Context, in *ListPublicPolicyImagesRequest, opts ...grpc.CallOption) (*ListPublicPolicyImagesResponse, error)
	CreatePolicyImage(ctx context.Context, in *CreatePolicyImageRequest, opts ...grpc.CallOption) (*CreatePolicyImageResponse, error)
	DeletePolicyImage(ctx context.Context, in *DeletePolicyImageRequest, opts ...grpc.CallOption) (*DeletePolicyImageResponse, error)
	UpdatePolicyImage(ctx context.Context, in *UpdatePolicyImageRequest, opts ...grpc.CallOption) (*UpdatePolicyImageResponse, error)
}

type policyClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyClient(cc grpc.ClientConnInterface) PolicyClient {
	return &policyClient{cc}
}

func (c *policyClient) GetPolicyImage(ctx context.Context, in *GetPolicyImageRequest, opts ...grpc.CallOption) (*GetPolicyImageResponse, error) {
	out := new(GetPolicyImageResponse)
	err := c.cc.Invoke(ctx, "/aserto.registry_tenant.v1.Policy/GetPolicyImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) ListPolicyImages(ctx context.Context, in *ListPolicyImagesRequest, opts ...grpc.CallOption) (*ListPolicyImagesResponse, error) {
	out := new(ListPolicyImagesResponse)
	err := c.cc.Invoke(ctx, "/aserto.registry_tenant.v1.Policy/ListPolicyImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) ListPublicPolicyImages(ctx context.Context, in *ListPublicPolicyImagesRequest, opts ...grpc.CallOption) (*ListPublicPolicyImagesResponse, error) {
	out := new(ListPublicPolicyImagesResponse)
	err := c.cc.Invoke(ctx, "/aserto.registry_tenant.v1.Policy/ListPublicPolicyImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) CreatePolicyImage(ctx context.Context, in *CreatePolicyImageRequest, opts ...grpc.CallOption) (*CreatePolicyImageResponse, error) {
	out := new(CreatePolicyImageResponse)
	err := c.cc.Invoke(ctx, "/aserto.registry_tenant.v1.Policy/CreatePolicyImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) DeletePolicyImage(ctx context.Context, in *DeletePolicyImageRequest, opts ...grpc.CallOption) (*DeletePolicyImageResponse, error) {
	out := new(DeletePolicyImageResponse)
	err := c.cc.Invoke(ctx, "/aserto.registry_tenant.v1.Policy/DeletePolicyImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) UpdatePolicyImage(ctx context.Context, in *UpdatePolicyImageRequest, opts ...grpc.CallOption) (*UpdatePolicyImageResponse, error) {
	out := new(UpdatePolicyImageResponse)
	err := c.cc.Invoke(ctx, "/aserto.registry_tenant.v1.Policy/UpdatePolicyImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyServer is the server API for Policy service.
// All implementations should embed UnimplementedPolicyServer
// for forward compatibility
type PolicyServer interface {
	GetPolicyImage(context.Context, *GetPolicyImageRequest) (*GetPolicyImageResponse, error)
	ListPolicyImages(context.Context, *ListPolicyImagesRequest) (*ListPolicyImagesResponse, error)
	ListPublicPolicyImages(context.Context, *ListPublicPolicyImagesRequest) (*ListPublicPolicyImagesResponse, error)
	CreatePolicyImage(context.Context, *CreatePolicyImageRequest) (*CreatePolicyImageResponse, error)
	DeletePolicyImage(context.Context, *DeletePolicyImageRequest) (*DeletePolicyImageResponse, error)
	UpdatePolicyImage(context.Context, *UpdatePolicyImageRequest) (*UpdatePolicyImageResponse, error)
}

// UnimplementedPolicyServer should be embedded to have forward compatible implementations.
type UnimplementedPolicyServer struct {
}

func (UnimplementedPolicyServer) GetPolicyImage(context.Context, *GetPolicyImageRequest) (*GetPolicyImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicyImage not implemented")
}
func (UnimplementedPolicyServer) ListPolicyImages(context.Context, *ListPolicyImagesRequest) (*ListPolicyImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPolicyImages not implemented")
}
func (UnimplementedPolicyServer) ListPublicPolicyImages(context.Context, *ListPublicPolicyImagesRequest) (*ListPublicPolicyImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPublicPolicyImages not implemented")
}
func (UnimplementedPolicyServer) CreatePolicyImage(context.Context, *CreatePolicyImageRequest) (*CreatePolicyImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolicyImage not implemented")
}
func (UnimplementedPolicyServer) DeletePolicyImage(context.Context, *DeletePolicyImageRequest) (*DeletePolicyImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicyImage not implemented")
}
func (UnimplementedPolicyServer) UpdatePolicyImage(context.Context, *UpdatePolicyImageRequest) (*UpdatePolicyImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePolicyImage not implemented")
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

func _Policy_GetPolicyImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).GetPolicyImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.registry_tenant.v1.Policy/GetPolicyImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).GetPolicyImage(ctx, req.(*GetPolicyImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_ListPolicyImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPolicyImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).ListPolicyImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.registry_tenant.v1.Policy/ListPolicyImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).ListPolicyImages(ctx, req.(*ListPolicyImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_ListPublicPolicyImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPublicPolicyImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).ListPublicPolicyImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.registry_tenant.v1.Policy/ListPublicPolicyImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).ListPublicPolicyImages(ctx, req.(*ListPublicPolicyImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_CreatePolicyImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePolicyImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).CreatePolicyImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.registry_tenant.v1.Policy/CreatePolicyImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).CreatePolicyImage(ctx, req.(*CreatePolicyImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_DeletePolicyImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).DeletePolicyImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.registry_tenant.v1.Policy/DeletePolicyImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).DeletePolicyImage(ctx, req.(*DeletePolicyImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_UpdatePolicyImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePolicyImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).UpdatePolicyImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.registry_tenant.v1.Policy/UpdatePolicyImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).UpdatePolicyImage(ctx, req.(*UpdatePolicyImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Policy_ServiceDesc is the grpc.ServiceDesc for Policy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Policy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.registry_tenant.v1.Policy",
	HandlerType: (*PolicyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPolicyImage",
			Handler:    _Policy_GetPolicyImage_Handler,
		},
		{
			MethodName: "ListPolicyImages",
			Handler:    _Policy_ListPolicyImages_Handler,
		},
		{
			MethodName: "ListPublicPolicyImages",
			Handler:    _Policy_ListPublicPolicyImages_Handler,
		},
		{
			MethodName: "CreatePolicyImage",
			Handler:    _Policy_CreatePolicyImage_Handler,
		},
		{
			MethodName: "DeletePolicyImage",
			Handler:    _Policy_DeletePolicyImage_Handler,
		},
		{
			MethodName: "UpdatePolicyImage",
			Handler:    _Policy_UpdatePolicyImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aserto/registry_tenant/v1/policy.proto",
}
