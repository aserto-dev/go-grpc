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
	ListImages(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error)
	RemoveImage(ctx context.Context, in *RemoveImageRequest, opts ...grpc.CallOption) (*RemoveImageResponse, error)
	SetImageVisibility(ctx context.Context, in *SetImageVisibilityRequest, opts ...grpc.CallOption) (*SetImageVisibilityResponse, error)
}

type registryClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryClient(cc grpc.ClientConnInterface) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) ListImages(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error) {
	out := new(ListImagesResponse)
	err := c.cc.Invoke(ctx, "/aserto.registry.v1.Registry/ListImages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) RemoveImage(ctx context.Context, in *RemoveImageRequest, opts ...grpc.CallOption) (*RemoveImageResponse, error) {
	out := new(RemoveImageResponse)
	err := c.cc.Invoke(ctx, "/aserto.registry.v1.Registry/RemoveImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) SetImageVisibility(ctx context.Context, in *SetImageVisibilityRequest, opts ...grpc.CallOption) (*SetImageVisibilityResponse, error) {
	out := new(SetImageVisibilityResponse)
	err := c.cc.Invoke(ctx, "/aserto.registry.v1.Registry/SetImageVisibility", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
// All implementations should embed UnimplementedRegistryServer
// for forward compatibility
type RegistryServer interface {
	ListImages(context.Context, *ListImagesRequest) (*ListImagesResponse, error)
	RemoveImage(context.Context, *RemoveImageRequest) (*RemoveImageResponse, error)
	SetImageVisibility(context.Context, *SetImageVisibilityRequest) (*SetImageVisibilityResponse, error)
}

// UnimplementedRegistryServer should be embedded to have forward compatible implementations.
type UnimplementedRegistryServer struct {
}

func (UnimplementedRegistryServer) ListImages(context.Context, *ListImagesRequest) (*ListImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImages not implemented")
}
func (UnimplementedRegistryServer) RemoveImage(context.Context, *RemoveImageRequest) (*RemoveImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveImage not implemented")
}
func (UnimplementedRegistryServer) SetImageVisibility(context.Context, *SetImageVisibilityRequest) (*SetImageVisibilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetImageVisibility not implemented")
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

func _Registry_ListImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.registry.v1.Registry/ListImages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListImages(ctx, req.(*ListImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_RemoveImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).RemoveImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.registry.v1.Registry/RemoveImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).RemoveImage(ctx, req.(*RemoveImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_SetImageVisibility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetImageVisibilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).SetImageVisibility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aserto.registry.v1.Registry/SetImageVisibility",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).SetImageVisibility(ctx, req.(*SetImageVisibilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Registry_ServiceDesc is the grpc.ServiceDesc for Registry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Registry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.registry.v1.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListImages",
			Handler:    _Registry_ListImages_Handler,
		},
		{
			MethodName: "RemoveImage",
			Handler:    _Registry_RemoveImage_Handler,
		},
		{
			MethodName: "SetImageVisibility",
			Handler:    _Registry_SetImageVisibility_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aserto/registry/v1/registry.proto",
}