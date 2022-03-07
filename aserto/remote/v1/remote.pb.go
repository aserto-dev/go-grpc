// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: aserto/remote/v1/remote.proto

package remote

import (
	errdetails "google.golang.org/genproto/googleapis/rpc/errdetails"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TriggerImageDiscoveryClientCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization string `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	Repository   string `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
}

func (x *TriggerImageDiscoveryClientCmd) Reset() {
	*x = TriggerImageDiscoveryClientCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_remote_v1_remote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerImageDiscoveryClientCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerImageDiscoveryClientCmd) ProtoMessage() {}

func (x *TriggerImageDiscoveryClientCmd) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_remote_v1_remote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerImageDiscoveryClientCmd.ProtoReflect.Descriptor instead.
func (*TriggerImageDiscoveryClientCmd) Descriptor() ([]byte, []int) {
	return file_aserto_remote_v1_remote_proto_rawDescGZIP(), []int{0}
}

func (x *TriggerImageDiscoveryClientCmd) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *TriggerImageDiscoveryClientCmd) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

type TriggerImageDiscoveryClientResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TriggerImageDiscoveryClientResp) Reset() {
	*x = TriggerImageDiscoveryClientResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_remote_v1_remote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerImageDiscoveryClientResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerImageDiscoveryClientResp) ProtoMessage() {}

func (x *TriggerImageDiscoveryClientResp) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_remote_v1_remote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerImageDiscoveryClientResp.ProtoReflect.Descriptor instead.
func (*TriggerImageDiscoveryClientResp) Descriptor() ([]byte, []int) {
	return file_aserto_remote_v1_remote_proto_rawDescGZIP(), []int{1}
}

type RegisterAuthorizerServerCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorizerId string `protobuf:"bytes,1,opt,name=authorizer_id,json=authorizerId,proto3" json:"authorizer_id,omitempty"`
}

func (x *RegisterAuthorizerServerCmd) Reset() {
	*x = RegisterAuthorizerServerCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_remote_v1_remote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAuthorizerServerCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAuthorizerServerCmd) ProtoMessage() {}

func (x *RegisterAuthorizerServerCmd) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_remote_v1_remote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAuthorizerServerCmd.ProtoReflect.Descriptor instead.
func (*RegisterAuthorizerServerCmd) Descriptor() ([]byte, []int) {
	return file_aserto_remote_v1_remote_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterAuthorizerServerCmd) GetAuthorizerId() string {
	if x != nil {
		return x.AuthorizerId
	}
	return ""
}

type RegisterAuthorizerServerResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterAuthorizerServerResp) Reset() {
	*x = RegisterAuthorizerServerResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_remote_v1_remote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAuthorizerServerResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAuthorizerServerResp) ProtoMessage() {}

func (x *RegisterAuthorizerServerResp) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_remote_v1_remote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAuthorizerServerResp.ProtoReflect.Descriptor instead.
func (*RegisterAuthorizerServerResp) Descriptor() ([]byte, []int) {
	return file_aserto_remote_v1_remote_proto_rawDescGZIP(), []int{3}
}

type CommandStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandId string `protobuf:"bytes,1,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	// Types that are assignable to Data:
	//	*CommandStreamRequest_Error
	//	*CommandStreamRequest_TriggerImageDiscovery
	//	*CommandStreamRequest_RegisterAuthorizer
	Data isCommandStreamRequest_Data `protobuf_oneof:"data"`
}

func (x *CommandStreamRequest) Reset() {
	*x = CommandStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_remote_v1_remote_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandStreamRequest) ProtoMessage() {}

func (x *CommandStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_remote_v1_remote_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandStreamRequest.ProtoReflect.Descriptor instead.
func (*CommandStreamRequest) Descriptor() ([]byte, []int) {
	return file_aserto_remote_v1_remote_proto_rawDescGZIP(), []int{4}
}

func (x *CommandStreamRequest) GetCommandId() string {
	if x != nil {
		return x.CommandId
	}
	return ""
}

func (m *CommandStreamRequest) GetData() isCommandStreamRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *CommandStreamRequest) GetError() *errdetails.ErrorInfo {
	if x, ok := x.GetData().(*CommandStreamRequest_Error); ok {
		return x.Error
	}
	return nil
}

func (x *CommandStreamRequest) GetTriggerImageDiscovery() *TriggerImageDiscoveryClientResp {
	if x, ok := x.GetData().(*CommandStreamRequest_TriggerImageDiscovery); ok {
		return x.TriggerImageDiscovery
	}
	return nil
}

func (x *CommandStreamRequest) GetRegisterAuthorizer() *RegisterAuthorizerServerCmd {
	if x, ok := x.GetData().(*CommandStreamRequest_RegisterAuthorizer); ok {
		return x.RegisterAuthorizer
	}
	return nil
}

type isCommandStreamRequest_Data interface {
	isCommandStreamRequest_Data()
}

type CommandStreamRequest_Error struct {
	Error *errdetails.ErrorInfo `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type CommandStreamRequest_TriggerImageDiscovery struct {
	TriggerImageDiscovery *TriggerImageDiscoveryClientResp `protobuf:"bytes,3,opt,name=trigger_image_discovery,json=triggerImageDiscovery,proto3,oneof"`
}

type CommandStreamRequest_RegisterAuthorizer struct {
	RegisterAuthorizer *RegisterAuthorizerServerCmd `protobuf:"bytes,4,opt,name=register_authorizer,json=registerAuthorizer,proto3,oneof"`
}

func (*CommandStreamRequest_Error) isCommandStreamRequest_Data() {}

func (*CommandStreamRequest_TriggerImageDiscovery) isCommandStreamRequest_Data() {}

func (*CommandStreamRequest_RegisterAuthorizer) isCommandStreamRequest_Data() {}

type CommandStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandId string `protobuf:"bytes,1,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	// Types that are assignable to Data:
	//	*CommandStreamResponse_Error
	//	*CommandStreamResponse_TriggerImageDiscovery
	//	*CommandStreamResponse_RegisterAuthorizer
	Data isCommandStreamResponse_Data `protobuf_oneof:"data"`
}

func (x *CommandStreamResponse) Reset() {
	*x = CommandStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_remote_v1_remote_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandStreamResponse) ProtoMessage() {}

func (x *CommandStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_remote_v1_remote_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandStreamResponse.ProtoReflect.Descriptor instead.
func (*CommandStreamResponse) Descriptor() ([]byte, []int) {
	return file_aserto_remote_v1_remote_proto_rawDescGZIP(), []int{5}
}

func (x *CommandStreamResponse) GetCommandId() string {
	if x != nil {
		return x.CommandId
	}
	return ""
}

func (m *CommandStreamResponse) GetData() isCommandStreamResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *CommandStreamResponse) GetError() *errdetails.ErrorInfo {
	if x, ok := x.GetData().(*CommandStreamResponse_Error); ok {
		return x.Error
	}
	return nil
}

func (x *CommandStreamResponse) GetTriggerImageDiscovery() *TriggerImageDiscoveryClientCmd {
	if x, ok := x.GetData().(*CommandStreamResponse_TriggerImageDiscovery); ok {
		return x.TriggerImageDiscovery
	}
	return nil
}

func (x *CommandStreamResponse) GetRegisterAuthorizer() *RegisterAuthorizerServerResp {
	if x, ok := x.GetData().(*CommandStreamResponse_RegisterAuthorizer); ok {
		return x.RegisterAuthorizer
	}
	return nil
}

type isCommandStreamResponse_Data interface {
	isCommandStreamResponse_Data()
}

type CommandStreamResponse_Error struct {
	Error *errdetails.ErrorInfo `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type CommandStreamResponse_TriggerImageDiscovery struct {
	TriggerImageDiscovery *TriggerImageDiscoveryClientCmd `protobuf:"bytes,3,opt,name=trigger_image_discovery,json=triggerImageDiscovery,proto3,oneof"`
}

type CommandStreamResponse_RegisterAuthorizer struct {
	RegisterAuthorizer *RegisterAuthorizerServerResp `protobuf:"bytes,4,opt,name=register_authorizer,json=registerAuthorizer,proto3,oneof"`
}

func (*CommandStreamResponse_Error) isCommandStreamResponse_Data() {}

func (*CommandStreamResponse_TriggerImageDiscovery) isCommandStreamResponse_Data() {}

func (*CommandStreamResponse_RegisterAuthorizer) isCommandStreamResponse_Data() {}

var File_aserto_remote_v1_remote_proto protoreflect.FileDescriptor

var file_aserto_remote_v1_remote_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x64, 0x0a, 0x1e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x43, 0x6d, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x21, 0x0a, 0x1f, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x42, 0x0a, 0x1b, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6d, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1e,
	0x0a, 0x1c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0xbb,
	0x02, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x6b, 0x0a, 0x17, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x48, 0x00, 0x52, 0x15, 0x74, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x12, 0x60, 0x0a, 0x13, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6d, 0x64, 0x48, 0x00,
	0x52, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xbc, 0x02, 0x0a,
	0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x6a, 0x0a, 0x17, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x6d, 0x64, 0x48, 0x00, 0x52, 0x15, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x12, 0x61, 0x0a, 0x13, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x48, 0x00, 0x52,
	0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x70, 0x0a, 0x06, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x66, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x26, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x4a, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0xaa, 0x02, 0x10, 0x41, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_aserto_remote_v1_remote_proto_rawDescOnce sync.Once
	file_aserto_remote_v1_remote_proto_rawDescData = file_aserto_remote_v1_remote_proto_rawDesc
)

func file_aserto_remote_v1_remote_proto_rawDescGZIP() []byte {
	file_aserto_remote_v1_remote_proto_rawDescOnce.Do(func() {
		file_aserto_remote_v1_remote_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_remote_v1_remote_proto_rawDescData)
	})
	return file_aserto_remote_v1_remote_proto_rawDescData
}

var file_aserto_remote_v1_remote_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_aserto_remote_v1_remote_proto_goTypes = []interface{}{
	(*TriggerImageDiscoveryClientCmd)(nil),  // 0: aserto.remote.v1.TriggerImageDiscoveryClientCmd
	(*TriggerImageDiscoveryClientResp)(nil), // 1: aserto.remote.v1.TriggerImageDiscoveryClientResp
	(*RegisterAuthorizerServerCmd)(nil),     // 2: aserto.remote.v1.RegisterAuthorizerServerCmd
	(*RegisterAuthorizerServerResp)(nil),    // 3: aserto.remote.v1.RegisterAuthorizerServerResp
	(*CommandStreamRequest)(nil),            // 4: aserto.remote.v1.CommandStreamRequest
	(*CommandStreamResponse)(nil),           // 5: aserto.remote.v1.CommandStreamResponse
	(*errdetails.ErrorInfo)(nil),            // 6: google.rpc.ErrorInfo
}
var file_aserto_remote_v1_remote_proto_depIdxs = []int32{
	6, // 0: aserto.remote.v1.CommandStreamRequest.error:type_name -> google.rpc.ErrorInfo
	1, // 1: aserto.remote.v1.CommandStreamRequest.trigger_image_discovery:type_name -> aserto.remote.v1.TriggerImageDiscoveryClientResp
	2, // 2: aserto.remote.v1.CommandStreamRequest.register_authorizer:type_name -> aserto.remote.v1.RegisterAuthorizerServerCmd
	6, // 3: aserto.remote.v1.CommandStreamResponse.error:type_name -> google.rpc.ErrorInfo
	0, // 4: aserto.remote.v1.CommandStreamResponse.trigger_image_discovery:type_name -> aserto.remote.v1.TriggerImageDiscoveryClientCmd
	3, // 5: aserto.remote.v1.CommandStreamResponse.register_authorizer:type_name -> aserto.remote.v1.RegisterAuthorizerServerResp
	4, // 6: aserto.remote.v1.Remote.CommandStream:input_type -> aserto.remote.v1.CommandStreamRequest
	5, // 7: aserto.remote.v1.Remote.CommandStream:output_type -> aserto.remote.v1.CommandStreamResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_aserto_remote_v1_remote_proto_init() }
func file_aserto_remote_v1_remote_proto_init() {
	if File_aserto_remote_v1_remote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_remote_v1_remote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerImageDiscoveryClientCmd); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aserto_remote_v1_remote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerImageDiscoveryClientResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aserto_remote_v1_remote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAuthorizerServerCmd); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aserto_remote_v1_remote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAuthorizerServerResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aserto_remote_v1_remote_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandStreamRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aserto_remote_v1_remote_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandStreamResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_aserto_remote_v1_remote_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*CommandStreamRequest_Error)(nil),
		(*CommandStreamRequest_TriggerImageDiscovery)(nil),
		(*CommandStreamRequest_RegisterAuthorizer)(nil),
	}
	file_aserto_remote_v1_remote_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*CommandStreamResponse_Error)(nil),
		(*CommandStreamResponse_TriggerImageDiscovery)(nil),
		(*CommandStreamResponse_RegisterAuthorizer)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aserto_remote_v1_remote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_remote_v1_remote_proto_goTypes,
		DependencyIndexes: file_aserto_remote_v1_remote_proto_depIdxs,
		MessageInfos:      file_aserto_remote_v1_remote_proto_msgTypes,
	}.Build()
	File_aserto_remote_v1_remote_proto = out.File
	file_aserto_remote_v1_remote_proto_rawDesc = nil
	file_aserto_remote_v1_remote_proto_goTypes = nil
	file_aserto_remote_v1_remote_proto_depIdxs = nil
}
