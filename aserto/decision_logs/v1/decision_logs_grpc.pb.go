// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: aserto/decision_logs/v1/decision_logs.proto

package decision_logs

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
	DecisionLogs_ListDecisionLogs_FullMethodName = "/aserto.decision_logs.v1.DecisionLogs/ListDecisionLogs"
	DecisionLogs_GetDecisionLog_FullMethodName   = "/aserto.decision_logs.v1.DecisionLogs/GetDecisionLog"
	DecisionLogs_ListUsers_FullMethodName        = "/aserto.decision_logs.v1.DecisionLogs/ListUsers"
	DecisionLogs_GetUser_FullMethodName          = "/aserto.decision_logs.v1.DecisionLogs/GetUser"
	DecisionLogs_ExecuteQuery_FullMethodName     = "/aserto.decision_logs.v1.DecisionLogs/ExecuteQuery"
	DecisionLogs_GetDecisions_FullMethodName     = "/aserto.decision_logs.v1.DecisionLogs/GetDecisions"
)

// DecisionLogsClient is the client API for DecisionLogs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DecisionLogsClient interface {
	ListDecisionLogs(ctx context.Context, in *ListDecisionLogsRequest, opts ...grpc.CallOption) (*ListDecisionLogsResponse, error)
	GetDecisionLog(ctx context.Context, in *GetDecisionLogRequest, opts ...grpc.CallOption) (*GetDecisionLogResponse, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	ExecuteQuery(ctx context.Context, in *ExecuteQueryRequest, opts ...grpc.CallOption) (*ExecuteQueryResponse, error)
	GetDecisions(ctx context.Context, in *GetDecisionsRequest, opts ...grpc.CallOption) (DecisionLogs_GetDecisionsClient, error)
}

type decisionLogsClient struct {
	cc grpc.ClientConnInterface
}

func NewDecisionLogsClient(cc grpc.ClientConnInterface) DecisionLogsClient {
	return &decisionLogsClient{cc}
}

func (c *decisionLogsClient) ListDecisionLogs(ctx context.Context, in *ListDecisionLogsRequest, opts ...grpc.CallOption) (*ListDecisionLogsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDecisionLogsResponse)
	err := c.cc.Invoke(ctx, DecisionLogs_ListDecisionLogs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *decisionLogsClient) GetDecisionLog(ctx context.Context, in *GetDecisionLogRequest, opts ...grpc.CallOption) (*GetDecisionLogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDecisionLogResponse)
	err := c.cc.Invoke(ctx, DecisionLogs_GetDecisionLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *decisionLogsClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, DecisionLogs_ListUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *decisionLogsClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, DecisionLogs_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *decisionLogsClient) ExecuteQuery(ctx context.Context, in *ExecuteQueryRequest, opts ...grpc.CallOption) (*ExecuteQueryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExecuteQueryResponse)
	err := c.cc.Invoke(ctx, DecisionLogs_ExecuteQuery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *decisionLogsClient) GetDecisions(ctx context.Context, in *GetDecisionsRequest, opts ...grpc.CallOption) (DecisionLogs_GetDecisionsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &DecisionLogs_ServiceDesc.Streams[0], DecisionLogs_GetDecisions_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &decisionLogsGetDecisionsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DecisionLogs_GetDecisionsClient interface {
	Recv() (*GetDecisionsResponse, error)
	grpc.ClientStream
}

type decisionLogsGetDecisionsClient struct {
	grpc.ClientStream
}

func (x *decisionLogsGetDecisionsClient) Recv() (*GetDecisionsResponse, error) {
	m := new(GetDecisionsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DecisionLogsServer is the server API for DecisionLogs service.
// All implementations should embed UnimplementedDecisionLogsServer
// for forward compatibility
type DecisionLogsServer interface {
	ListDecisionLogs(context.Context, *ListDecisionLogsRequest) (*ListDecisionLogsResponse, error)
	GetDecisionLog(context.Context, *GetDecisionLogRequest) (*GetDecisionLogResponse, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	ExecuteQuery(context.Context, *ExecuteQueryRequest) (*ExecuteQueryResponse, error)
	GetDecisions(*GetDecisionsRequest, DecisionLogs_GetDecisionsServer) error
}

// UnimplementedDecisionLogsServer should be embedded to have forward compatible implementations.
type UnimplementedDecisionLogsServer struct {
}

func (UnimplementedDecisionLogsServer) ListDecisionLogs(context.Context, *ListDecisionLogsRequest) (*ListDecisionLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDecisionLogs not implemented")
}
func (UnimplementedDecisionLogsServer) GetDecisionLog(context.Context, *GetDecisionLogRequest) (*GetDecisionLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDecisionLog not implemented")
}
func (UnimplementedDecisionLogsServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedDecisionLogsServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedDecisionLogsServer) ExecuteQuery(context.Context, *ExecuteQueryRequest) (*ExecuteQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteQuery not implemented")
}
func (UnimplementedDecisionLogsServer) GetDecisions(*GetDecisionsRequest, DecisionLogs_GetDecisionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDecisions not implemented")
}

// UnsafeDecisionLogsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DecisionLogsServer will
// result in compilation errors.
type UnsafeDecisionLogsServer interface {
	mustEmbedUnimplementedDecisionLogsServer()
}

func RegisterDecisionLogsServer(s grpc.ServiceRegistrar, srv DecisionLogsServer) {
	s.RegisterService(&DecisionLogs_ServiceDesc, srv)
}

func _DecisionLogs_ListDecisionLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDecisionLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecisionLogsServer).ListDecisionLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DecisionLogs_ListDecisionLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecisionLogsServer).ListDecisionLogs(ctx, req.(*ListDecisionLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DecisionLogs_GetDecisionLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDecisionLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecisionLogsServer).GetDecisionLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DecisionLogs_GetDecisionLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecisionLogsServer).GetDecisionLog(ctx, req.(*GetDecisionLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DecisionLogs_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecisionLogsServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DecisionLogs_ListUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecisionLogsServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DecisionLogs_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecisionLogsServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DecisionLogs_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecisionLogsServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DecisionLogs_ExecuteQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecisionLogsServer).ExecuteQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DecisionLogs_ExecuteQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecisionLogsServer).ExecuteQuery(ctx, req.(*ExecuteQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DecisionLogs_GetDecisions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDecisionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DecisionLogsServer).GetDecisions(m, &decisionLogsGetDecisionsServer{ServerStream: stream})
}

type DecisionLogs_GetDecisionsServer interface {
	Send(*GetDecisionsResponse) error
	grpc.ServerStream
}

type decisionLogsGetDecisionsServer struct {
	grpc.ServerStream
}

func (x *decisionLogsGetDecisionsServer) Send(m *GetDecisionsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// DecisionLogs_ServiceDesc is the grpc.ServiceDesc for DecisionLogs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DecisionLogs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aserto.decision_logs.v1.DecisionLogs",
	HandlerType: (*DecisionLogsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDecisionLogs",
			Handler:    _DecisionLogs_ListDecisionLogs_Handler,
		},
		{
			MethodName: "GetDecisionLog",
			Handler:    _DecisionLogs_GetDecisionLog_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _DecisionLogs_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _DecisionLogs_GetUser_Handler,
		},
		{
			MethodName: "ExecuteQuery",
			Handler:    _DecisionLogs_ExecuteQuery_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDecisions",
			Handler:       _DecisionLogs_GetDecisions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "aserto/decision_logs/v1/decision_logs.proto",
}
