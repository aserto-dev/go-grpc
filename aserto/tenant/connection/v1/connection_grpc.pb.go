// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: aserto/tenant/connection/v1/connection.proto

package connection

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
	Connection_ListConnections_FullMethodName     = "/aserto.tenant.connection.v1.Connection/ListConnections"
	Connection_GetConnection_FullMethodName       = "/aserto.tenant.connection.v1.Connection/GetConnection"
	Connection_CreateConnection_FullMethodName    = "/aserto.tenant.connection.v1.Connection/CreateConnection"
	Connection_UpdateConnection_FullMethodName    = "/aserto.tenant.connection.v1.Connection/UpdateConnection"
	Connection_DeleteConnection_FullMethodName    = "/aserto.tenant.connection.v1.Connection/DeleteConnection"
	Connection_VerifyConnection_FullMethodName    = "/aserto.tenant.connection.v1.Connection/VerifyConnection"
	Connection_RotateSecret_FullMethodName        = "/aserto.tenant.connection.v1.Connection/RotateSecret"
	Connection_ConnectionAvailable_FullMethodName = "/aserto.tenant.connection.v1.Connection/ConnectionAvailable"
)

// ConnectionClient is the client API for Connection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionClient interface {
	ListConnections(ctx context.Context, in *ListConnectionsRequest, opts ...grpc.CallOption) (*ListConnectionsResponse, error)
	GetConnection(ctx context.Context, in *GetConnectionRequest, opts ...grpc.CallOption) (*GetConnectionResponse, error)
	CreateConnection(ctx context.Context, in *CreateConnectionRequest, opts ...grpc.CallOption) (*CreateConnectionResponse, error)
	UpdateConnection(ctx context.Context, in *UpdateConnectionRequest, opts ...grpc.CallOption) (*UpdateConnectionResponse, error)
	DeleteConnection(ctx context.Context, in *DeleteConnectionRequest, opts ...grpc.CallOption) (*DeleteConnectionResponse, error)
	VerifyConnection(ctx context.Context, in *VerifyConnectionRequest, opts ...grpc.CallOption) (*VerifyConnectionResponse, error)
	RotateSecret(ctx context.Context, in *RotateSecretRequest, opts ...grpc.CallOption) (*RotateSecretResponse, error)
	ConnectionAvailable(ctx context.Context, in *ConnectionAvailableRequest, opts ...grpc.CallOption) (*ConnectionAvailableResponse, error)
}

type connectionClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionClient(cc grpc.ClientConnInterface) ConnectionClient {
	return &connectionClient{cc}
}

func (c *connectionClient) ListConnections(ctx context.Context, in *ListConnectionsRequest, opts ...grpc.CallOption) (*ListConnectionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListConnectionsResponse)
	err := c.cc.Invoke(ctx, Connection_ListConnections_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionClient) GetConnection(ctx context.Context, in *GetConnectionRequest, opts ...grpc.CallOption) (*GetConnectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetConnectionResponse)
	err := c.cc.Invoke(ctx, Connection_GetConnection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionClient) CreateConnection(ctx context.Context, in *CreateConnectionRequest, opts ...grpc.CallOption) (*CreateConnectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateConnectionResponse)
	err := c.cc.Invoke(ctx, Connection_CreateConnection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionClient) UpdateConnection(ctx context.Context, in *UpdateConnectionRequest, opts ...grpc.CallOption) (*UpdateConnectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateConnectionResponse)
	err := c.cc.Invoke(ctx, Connection_UpdateConnection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionClient) DeleteConnection(ctx context.Context, in *DeleteConnectionRequest, opts ...grpc.CallOption) (*DeleteConnectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteConnectionResponse)
	err := c.cc.Invoke(ctx, Connection_DeleteConnection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionClient) VerifyConnection(ctx context.Context, in *VerifyConnectionRequest, opts ...grpc.CallOption) (*VerifyConnectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyConnectionResponse)
	err := c.cc.Invoke(ctx, Connection_VerifyConnection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionClient) RotateSecret(ctx context.Context, in *RotateSecretRequest, opts ...grpc.CallOption) (*RotateSecretResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RotateSecretResponse)
	err := c.cc.Invoke(ctx, Connection_RotateSecret_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionClient) ConnectionAvailable(ctx context.Context, in *ConnectionAvailableRequest, opts ...grpc.CallOption) (*ConnectionAvailableResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConnectionAvailableResponse)
	err := c.cc.Invoke(ctx, Connection_ConnectionAvailable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectionServer is the server API for Connection service.
// All implementations should embed UnimplementedConnectionServer
// for forward compatibility
type ConnectionServer interface {
	ListConnections(context.Context, *ListConnectionsRequest) (*ListConnectionsResponse, error)
	GetConnection(context.Context, *GetConnectionRequest) (*GetConnectionResponse, error)
	CreateConnection(context.Context, *CreateConnectionRequest) (*CreateConnectionResponse, error)
	UpdateConnection(context.Context, *UpdateConnectionRequest) (*UpdateConnectionResponse, error)
	DeleteConnection(context.Context, *DeleteConnectionRequest) (*DeleteConnectionResponse, error)
	VerifyConnection(context.Context, *VerifyConnectionRequest) (*VerifyConnectionResponse, error)
	RotateSecret(context.Context, *RotateSecretRequest) (*RotateSecretResponse, error)
	ConnectionAvailable(context.Context, *ConnectionAvailableRequest) (*ConnectionAvailableResponse, error)
}

// UnimplementedConnectionServer should be embedded to have forward compatible implementations.
type UnimplementedConnectionServer struct {
}

func (UnimplementedConnectionServer) ListConnections(context.Context, *ListConnectionsRequest) (*ListConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConnections not implemented")
}
func (UnimplementedConnectionServer) GetConnection(context.Context, *GetConnectionRequest) (*GetConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnection not implemented")
}
func (UnimplementedConnectionServer) CreateConnection(context.Context, *CreateConnectionRequest) (*CreateConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConnection not implemented")
}
func (UnimplementedConnectionServer) UpdateConnection(context.Context, *UpdateConnectionRequest) (*UpdateConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConnection not implemented")
}
func (UnimplementedConnectionServer) DeleteConnection(context.Context, *DeleteConnectionRequest) (*DeleteConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConnection not implemented")
}
func (UnimplementedConnectionServer) VerifyConnection(context.Context, *VerifyConnectionRequest) (*VerifyConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyConnection not implemented")
}
func (UnimplementedConnectionServer) RotateSecret(context.Context, *RotateSecretRequest) (*RotateSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RotateSecret not implemented")
}
func (UnimplementedConnectionServer) ConnectionAvailable(context.Context, *ConnectionAvailableRequest) (*ConnectionAvailableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectionAvailable not implemented")
}

// UnsafeConnectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionServer will
// result in compilation errors.
type UnsafeConnectionServer interface {
	mustEmbedUnimplementedConnectionServer()
}

func RegisterConnectionServer(s grpc.ServiceRegistrar, srv ConnectionServer) {
	s.RegisterService(&Connection_ServiceDesc, srv)
}

func _Connection_ListConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConnectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServer).ListConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Connection_ListConnections_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServer).ListConnections(ctx, req.(*ListConnectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connection_GetConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServer).GetConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Connection_GetConnection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServer).GetConnection(ctx, req.(*GetConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connection_CreateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServer).CreateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Connection_CreateConnection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServer).CreateConnection(ctx, req.(*CreateConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connection_UpdateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServer).UpdateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Connection_UpdateConnection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServer).UpdateConnection(ctx, req.(*UpdateConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connection_DeleteConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServer).DeleteConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Connection_DeleteConnection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServer).DeleteConnection(ctx, req.(*DeleteConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connection_VerifyConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServer).VerifyConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Connection_VerifyConnection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServer).VerifyConnection(ctx, req.(*VerifyConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connection_RotateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RotateSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServer).RotateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Connection_RotateSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServer).RotateSecret(ctx, req.(*RotateSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connection_ConnectionAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServer).ConnectionAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Connection_ConnectionAvailable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServer).ConnectionAvailable(ctx, req.(*ConnectionAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Connection_ServiceDesc is the grpc.ServiceDesc for Connection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Connection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.tenant.connection.v1.Connection",
	HandlerType: (*ConnectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListConnections",
			Handler:    _Connection_ListConnections_Handler,
		},
		{
			MethodName: "GetConnection",
			Handler:    _Connection_GetConnection_Handler,
		},
		{
			MethodName: "CreateConnection",
			Handler:    _Connection_CreateConnection_Handler,
		},
		{
			MethodName: "UpdateConnection",
			Handler:    _Connection_UpdateConnection_Handler,
		},
		{
			MethodName: "DeleteConnection",
			Handler:    _Connection_DeleteConnection_Handler,
		},
		{
			MethodName: "VerifyConnection",
			Handler:    _Connection_VerifyConnection_Handler,
		},
		{
			MethodName: "RotateSecret",
			Handler:    _Connection_RotateSecret_Handler,
		},
		{
			MethodName: "ConnectionAvailable",
			Handler:    _Connection_ConnectionAvailable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aserto/tenant/connection/v1/connection.proto",
}
