// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: aserto/scribe/v2/scribe.proto

package scribe

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
	Scribe_WriteBatch_FullMethodName = "/aserto.scribe.v2.Scribe/WriteBatch"
)

// ScribeClient is the client API for Scribe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScribeClient interface {
	WriteBatch(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[WriteBatchRequest, WriteBatchResponse], error)
}

type scribeClient struct {
	cc grpc.ClientConnInterface
}

func NewScribeClient(cc grpc.ClientConnInterface) ScribeClient {
	return &scribeClient{cc}
}

func (c *scribeClient) WriteBatch(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[WriteBatchRequest, WriteBatchResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Scribe_ServiceDesc.Streams[0], Scribe_WriteBatch_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[WriteBatchRequest, WriteBatchResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Scribe_WriteBatchClient = grpc.BidiStreamingClient[WriteBatchRequest, WriteBatchResponse]

// ScribeServer is the server API for Scribe service.
// All implementations should embed UnimplementedScribeServer
// for forward compatibility.
type ScribeServer interface {
	WriteBatch(grpc.BidiStreamingServer[WriteBatchRequest, WriteBatchResponse]) error
}

// UnimplementedScribeServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedScribeServer struct{}

func (UnimplementedScribeServer) WriteBatch(grpc.BidiStreamingServer[WriteBatchRequest, WriteBatchResponse]) error {
	return status.Errorf(codes.Unimplemented, "method WriteBatch not implemented")
}
func (UnimplementedScribeServer) testEmbeddedByValue() {}

// UnsafeScribeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScribeServer will
// result in compilation errors.
type UnsafeScribeServer interface {
	mustEmbedUnimplementedScribeServer()
}

func RegisterScribeServer(s grpc.ServiceRegistrar, srv ScribeServer) {
	// If the following call pancis, it indicates UnimplementedScribeServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Scribe_ServiceDesc, srv)
}

func _Scribe_WriteBatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ScribeServer).WriteBatch(&grpc.GenericServerStream[WriteBatchRequest, WriteBatchResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Scribe_WriteBatchServer = grpc.BidiStreamingServer[WriteBatchRequest, WriteBatchResponse]

// Scribe_ServiceDesc is the grpc.ServiceDesc for Scribe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Scribe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.scribe.v2.Scribe",
	HandlerType: (*ScribeServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WriteBatch",
			Handler:       _Scribe_WriteBatch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "aserto/scribe/v2/scribe.proto",
}
