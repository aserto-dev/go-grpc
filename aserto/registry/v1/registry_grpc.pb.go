// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aserto/registry/v1/registry.proto

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

const (
	Registry_ListPublicOrgs_FullMethodName      = "/aserto.registry.v1.Registry/ListPublicOrgs"
	Registry_ListPublicImages_FullMethodName    = "/aserto.registry.v1.Registry/ListPublicImages"
	Registry_ListImages_FullMethodName          = "/aserto.registry.v1.Registry/ListImages"
	Registry_ListOrgs_FullMethodName            = "/aserto.registry.v1.Registry/ListOrgs"
	Registry_RemoveImage_FullMethodName         = "/aserto.registry.v1.Registry/RemoveImage"
	Registry_CreateImage_FullMethodName         = "/aserto.registry.v1.Registry/CreateImage"
	Registry_SetImageVisibility_FullMethodName  = "/aserto.registry.v1.Registry/SetImageVisibility"
	Registry_GetReadAccessToken_FullMethodName  = "/aserto.registry.v1.Registry/GetReadAccessToken"
	Registry_GetWriteAccessToken_FullMethodName = "/aserto.registry.v1.Registry/GetWriteAccessToken"
	Registry_ListTagsWithDetails_FullMethodName = "/aserto.registry.v1.Registry/ListTagsWithDetails"
	Registry_ListDigests_FullMethodName         = "/aserto.registry.v1.Registry/ListDigests"
	Registry_RepoAvailable_FullMethodName       = "/aserto.registry.v1.Registry/RepoAvailable"
)

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistryClient interface {
	ListPublicOrgs(ctx context.Context, in *ListPublicOrgsRequest, opts ...grpc.CallOption) (*ListPublicOrgsResponse, error)
	ListPublicImages(ctx context.Context, in *ListPublicImagesRequest, opts ...grpc.CallOption) (*ListPublicImagesResponse, error)
	ListImages(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error)
	ListOrgs(ctx context.Context, in *ListOrgsRequest, opts ...grpc.CallOption) (*ListOrgsResponse, error)
	RemoveImage(ctx context.Context, in *RemoveImageRequest, opts ...grpc.CallOption) (*RemoveImageResponse, error)
	CreateImage(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*CreateImageResponse, error)
	SetImageVisibility(ctx context.Context, in *SetImageVisibilityRequest, opts ...grpc.CallOption) (*SetImageVisibilityResponse, error)
	GetReadAccessToken(ctx context.Context, in *GetReadAccessTokenRequest, opts ...grpc.CallOption) (*GetReadAccessTokenResponse, error)
	GetWriteAccessToken(ctx context.Context, in *GetWriteAccessTokenRequest, opts ...grpc.CallOption) (*GetWriteAccessTokenResponse, error)
	ListTagsWithDetails(ctx context.Context, in *ListTagsWithDetailsRequest, opts ...grpc.CallOption) (*ListTagsWithDetailsResponse, error)
	ListDigests(ctx context.Context, in *ListDigestsRequest, opts ...grpc.CallOption) (*ListDigestsResponse, error)
	RepoAvailable(ctx context.Context, in *RepoAvailableRequest, opts ...grpc.CallOption) (*RepoAvailableResponse, error)
}

type registryClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistryClient(cc grpc.ClientConnInterface) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) ListPublicOrgs(ctx context.Context, in *ListPublicOrgsRequest, opts ...grpc.CallOption) (*ListPublicOrgsResponse, error) {
	out := new(ListPublicOrgsResponse)
	err := c.cc.Invoke(ctx, Registry_ListPublicOrgs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListPublicImages(ctx context.Context, in *ListPublicImagesRequest, opts ...grpc.CallOption) (*ListPublicImagesResponse, error) {
	out := new(ListPublicImagesResponse)
	err := c.cc.Invoke(ctx, Registry_ListPublicImages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListImages(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error) {
	out := new(ListImagesResponse)
	err := c.cc.Invoke(ctx, Registry_ListImages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListOrgs(ctx context.Context, in *ListOrgsRequest, opts ...grpc.CallOption) (*ListOrgsResponse, error) {
	out := new(ListOrgsResponse)
	err := c.cc.Invoke(ctx, Registry_ListOrgs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) RemoveImage(ctx context.Context, in *RemoveImageRequest, opts ...grpc.CallOption) (*RemoveImageResponse, error) {
	out := new(RemoveImageResponse)
	err := c.cc.Invoke(ctx, Registry_RemoveImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) CreateImage(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*CreateImageResponse, error) {
	out := new(CreateImageResponse)
	err := c.cc.Invoke(ctx, Registry_CreateImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) SetImageVisibility(ctx context.Context, in *SetImageVisibilityRequest, opts ...grpc.CallOption) (*SetImageVisibilityResponse, error) {
	out := new(SetImageVisibilityResponse)
	err := c.cc.Invoke(ctx, Registry_SetImageVisibility_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetReadAccessToken(ctx context.Context, in *GetReadAccessTokenRequest, opts ...grpc.CallOption) (*GetReadAccessTokenResponse, error) {
	out := new(GetReadAccessTokenResponse)
	err := c.cc.Invoke(ctx, Registry_GetReadAccessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) GetWriteAccessToken(ctx context.Context, in *GetWriteAccessTokenRequest, opts ...grpc.CallOption) (*GetWriteAccessTokenResponse, error) {
	out := new(GetWriteAccessTokenResponse)
	err := c.cc.Invoke(ctx, Registry_GetWriteAccessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListTagsWithDetails(ctx context.Context, in *ListTagsWithDetailsRequest, opts ...grpc.CallOption) (*ListTagsWithDetailsResponse, error) {
	out := new(ListTagsWithDetailsResponse)
	err := c.cc.Invoke(ctx, Registry_ListTagsWithDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) ListDigests(ctx context.Context, in *ListDigestsRequest, opts ...grpc.CallOption) (*ListDigestsResponse, error) {
	out := new(ListDigestsResponse)
	err := c.cc.Invoke(ctx, Registry_ListDigests_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryClient) RepoAvailable(ctx context.Context, in *RepoAvailableRequest, opts ...grpc.CallOption) (*RepoAvailableResponse, error) {
	out := new(RepoAvailableResponse)
	err := c.cc.Invoke(ctx, Registry_RepoAvailable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
// All implementations should embed UnimplementedRegistryServer
// for forward compatibility
type RegistryServer interface {
	ListPublicOrgs(context.Context, *ListPublicOrgsRequest) (*ListPublicOrgsResponse, error)
	ListPublicImages(context.Context, *ListPublicImagesRequest) (*ListPublicImagesResponse, error)
	ListImages(context.Context, *ListImagesRequest) (*ListImagesResponse, error)
	ListOrgs(context.Context, *ListOrgsRequest) (*ListOrgsResponse, error)
	RemoveImage(context.Context, *RemoveImageRequest) (*RemoveImageResponse, error)
	CreateImage(context.Context, *CreateImageRequest) (*CreateImageResponse, error)
	SetImageVisibility(context.Context, *SetImageVisibilityRequest) (*SetImageVisibilityResponse, error)
	GetReadAccessToken(context.Context, *GetReadAccessTokenRequest) (*GetReadAccessTokenResponse, error)
	GetWriteAccessToken(context.Context, *GetWriteAccessTokenRequest) (*GetWriteAccessTokenResponse, error)
	ListTagsWithDetails(context.Context, *ListTagsWithDetailsRequest) (*ListTagsWithDetailsResponse, error)
	ListDigests(context.Context, *ListDigestsRequest) (*ListDigestsResponse, error)
	RepoAvailable(context.Context, *RepoAvailableRequest) (*RepoAvailableResponse, error)
}

// UnimplementedRegistryServer should be embedded to have forward compatible implementations.
type UnimplementedRegistryServer struct {
}

func (UnimplementedRegistryServer) ListPublicOrgs(context.Context, *ListPublicOrgsRequest) (*ListPublicOrgsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPublicOrgs not implemented")
}
func (UnimplementedRegistryServer) ListPublicImages(context.Context, *ListPublicImagesRequest) (*ListPublicImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPublicImages not implemented")
}
func (UnimplementedRegistryServer) ListImages(context.Context, *ListImagesRequest) (*ListImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImages not implemented")
}
func (UnimplementedRegistryServer) ListOrgs(context.Context, *ListOrgsRequest) (*ListOrgsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrgs not implemented")
}
func (UnimplementedRegistryServer) RemoveImage(context.Context, *RemoveImageRequest) (*RemoveImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveImage not implemented")
}
func (UnimplementedRegistryServer) CreateImage(context.Context, *CreateImageRequest) (*CreateImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateImage not implemented")
}
func (UnimplementedRegistryServer) SetImageVisibility(context.Context, *SetImageVisibilityRequest) (*SetImageVisibilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetImageVisibility not implemented")
}
func (UnimplementedRegistryServer) GetReadAccessToken(context.Context, *GetReadAccessTokenRequest) (*GetReadAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReadAccessToken not implemented")
}
func (UnimplementedRegistryServer) GetWriteAccessToken(context.Context, *GetWriteAccessTokenRequest) (*GetWriteAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWriteAccessToken not implemented")
}
func (UnimplementedRegistryServer) ListTagsWithDetails(context.Context, *ListTagsWithDetailsRequest) (*ListTagsWithDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTagsWithDetails not implemented")
}
func (UnimplementedRegistryServer) ListDigests(context.Context, *ListDigestsRequest) (*ListDigestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDigests not implemented")
}
func (UnimplementedRegistryServer) RepoAvailable(context.Context, *RepoAvailableRequest) (*RepoAvailableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepoAvailable not implemented")
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

func _Registry_ListPublicOrgs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPublicOrgsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListPublicOrgs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_ListPublicOrgs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListPublicOrgs(ctx, req.(*ListPublicOrgsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListPublicImages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPublicImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListPublicImages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_ListPublicImages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListPublicImages(ctx, req.(*ListPublicImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: Registry_ListImages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListImages(ctx, req.(*ListImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: Registry_ListOrgs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListOrgs(ctx, req.(*ListOrgsRequest))
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
		FullMethod: Registry_RemoveImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).RemoveImage(ctx, req.(*RemoveImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_CreateImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).CreateImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_CreateImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).CreateImage(ctx, req.(*CreateImageRequest))
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
		FullMethod: Registry_SetImageVisibility_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).SetImageVisibility(ctx, req.(*SetImageVisibilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetReadAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReadAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetReadAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_GetReadAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetReadAccessToken(ctx, req.(*GetReadAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_GetWriteAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWriteAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).GetWriteAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_GetWriteAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).GetWriteAccessToken(ctx, req.(*GetWriteAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListTagsWithDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagsWithDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListTagsWithDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_ListTagsWithDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListTagsWithDetails(ctx, req.(*ListTagsWithDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_ListDigests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDigestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).ListDigests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_ListDigests_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).ListDigests(ctx, req.(*ListDigestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registry_RepoAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).RepoAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registry_RepoAvailable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).RepoAvailable(ctx, req.(*RepoAvailableRequest))
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
			MethodName: "ListPublicOrgs",
			Handler:    _Registry_ListPublicOrgs_Handler,
		},
		{
			MethodName: "ListPublicImages",
			Handler:    _Registry_ListPublicImages_Handler,
		},
		{
			MethodName: "ListImages",
			Handler:    _Registry_ListImages_Handler,
		},
		{
			MethodName: "ListOrgs",
			Handler:    _Registry_ListOrgs_Handler,
		},
		{
			MethodName: "RemoveImage",
			Handler:    _Registry_RemoveImage_Handler,
		},
		{
			MethodName: "CreateImage",
			Handler:    _Registry_CreateImage_Handler,
		},
		{
			MethodName: "SetImageVisibility",
			Handler:    _Registry_SetImageVisibility_Handler,
		},
		{
			MethodName: "GetReadAccessToken",
			Handler:    _Registry_GetReadAccessToken_Handler,
		},
		{
			MethodName: "GetWriteAccessToken",
			Handler:    _Registry_GetWriteAccessToken_Handler,
		},
		{
			MethodName: "ListTagsWithDetails",
			Handler:    _Registry_ListTagsWithDetails_Handler,
		},
		{
			MethodName: "ListDigests",
			Handler:    _Registry_ListDigests_Handler,
		},
		{
			MethodName: "RepoAvailable",
			Handler:    _Registry_RepoAvailable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aserto/registry/v1/registry.proto",
}
