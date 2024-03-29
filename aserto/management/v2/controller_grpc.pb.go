// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: aserto/management/v2/controller.proto

package management

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
	Controller_CommandStream_FullMethodName = "/aserto.management.v2.Controller/CommandStream"
)

// ControllerClient is the client API for Controller service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControllerClient interface {
	CommandStream(ctx context.Context, in *CommandStreamRequest, opts ...grpc.CallOption) (Controller_CommandStreamClient, error)
}

type controllerClient struct {
	cc grpc.ClientConnInterface
}

func NewControllerClient(cc grpc.ClientConnInterface) ControllerClient {
	return &controllerClient{cc}
}

func (c *controllerClient) CommandStream(ctx context.Context, in *CommandStreamRequest, opts ...grpc.CallOption) (Controller_CommandStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Controller_ServiceDesc.Streams[0], Controller_CommandStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &controllerCommandStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Controller_CommandStreamClient interface {
	Recv() (*CommandStreamResponse, error)
	grpc.ClientStream
}

type controllerCommandStreamClient struct {
	grpc.ClientStream
}

func (x *controllerCommandStreamClient) Recv() (*CommandStreamResponse, error) {
	m := new(CommandStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ControllerServer is the server API for Controller service.
// All implementations should embed UnimplementedControllerServer
// for forward compatibility
type ControllerServer interface {
	CommandStream(*CommandStreamRequest, Controller_CommandStreamServer) error
}

// UnimplementedControllerServer should be embedded to have forward compatible implementations.
type UnimplementedControllerServer struct {
}

func (UnimplementedControllerServer) CommandStream(*CommandStreamRequest, Controller_CommandStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CommandStream not implemented")
}

// UnsafeControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControllerServer will
// result in compilation errors.
type UnsafeControllerServer interface {
	mustEmbedUnimplementedControllerServer()
}

func RegisterControllerServer(s grpc.ServiceRegistrar, srv ControllerServer) {
	s.RegisterService(&Controller_ServiceDesc, srv)
}

func _Controller_CommandStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CommandStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControllerServer).CommandStream(m, &controllerCommandStreamServer{stream})
}

type Controller_CommandStreamServer interface {
	Send(*CommandStreamResponse) error
	grpc.ServerStream
}

type controllerCommandStreamServer struct {
	grpc.ServerStream
}

func (x *controllerCommandStreamServer) Send(m *CommandStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Controller_ServiceDesc is the grpc.ServiceDesc for Controller service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Controller_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.management.v2.Controller",
	HandlerType: (*ControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CommandStream",
			Handler:       _Controller_CommandStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "aserto/management/v2/controller.proto",
}
