// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: aserto/tenant/v2/policy_state.proto

package tenant

import (
	v2 "github.com/aserto-dev/go-grpc/aserto/api/v2"
	_ "github.com/aserto-dev/go-grpc/aserto/options/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetPolicyStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyId string `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
}

func (x *GetPolicyStateRequest) Reset() {
	*x = GetPolicyStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_v2_policy_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPolicyStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicyStateRequest) ProtoMessage() {}

func (x *GetPolicyStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_policy_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicyStateRequest.ProtoReflect.Descriptor instead.
func (*GetPolicyStateRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_policy_state_proto_rawDescGZIP(), []int{0}
}

func (x *GetPolicyStateRequest) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

type GetPolicyStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *v2.PolicyState `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *GetPolicyStateResponse) Reset() {
	*x = GetPolicyStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_v2_policy_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPolicyStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicyStateResponse) ProtoMessage() {}

func (x *GetPolicyStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_policy_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicyStateResponse.ProtoReflect.Descriptor instead.
func (*GetPolicyStateResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_policy_state_proto_rawDescGZIP(), []int{1}
}

func (x *GetPolicyStateResponse) GetState() *v2.PolicyState {
	if x != nil {
		return x.State
	}
	return nil
}

type SetPolicyStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *v2.PolicyState `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *SetPolicyStateRequest) Reset() {
	*x = SetPolicyStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_v2_policy_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPolicyStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPolicyStateRequest) ProtoMessage() {}

func (x *SetPolicyStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_policy_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPolicyStateRequest.ProtoReflect.Descriptor instead.
func (*SetPolicyStateRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_policy_state_proto_rawDescGZIP(), []int{2}
}

func (x *SetPolicyStateRequest) GetState() *v2.PolicyState {
	if x != nil {
		return x.State
	}
	return nil
}

type SetPolicyStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetPolicyStateResponse) Reset() {
	*x = SetPolicyStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_v2_policy_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPolicyStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPolicyStateResponse) ProtoMessage() {}

func (x *SetPolicyStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_policy_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPolicyStateResponse.ProtoReflect.Descriptor instead.
func (*SetPolicyStateResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_policy_state_proto_rawDescGZIP(), []int{3}
}

var File_aserto_tenant_v2_policy_state_proto protoreflect.FileDescriptor

var file_aserto_tenant_v2_policy_state_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f,
	0x76, 0x32, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x90, 0xb5, 0x18, 0x04,
	0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x49, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x18, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb2, 0x04, 0x0a, 0x0b,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x89, 0x02, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x27,
	0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76,
	0x32, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xa3, 0x01, 0x92, 0x41, 0x79, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x47, 0x65, 0x74, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x20, 0x73, 0x74, 0x61, 0x74, 0x65, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44, 0x1a, 0x1b, 0x47, 0x65,
	0x74, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x20, 0x73, 0x74, 0x61, 0x74, 0x65, 0x20, 0x62,
	0x79, 0x20, 0x69, 0x74, 0x73, 0x20, 0x49, 0x44, 0x2e, 0x2a, 0x1c, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54,
	0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x96, 0x02, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb0, 0x01,
	0x92, 0x41, 0x79, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x53, 0x65, 0x74, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x20, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x20, 0x62, 0x79, 0x20, 0x49, 0x44, 0x1a, 0x1b, 0x53, 0x65, 0x74, 0x20, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x20, 0x73, 0x74, 0x61, 0x74, 0x65, 0x20, 0x62, 0x79, 0x20, 0x69, 0x74,
	0x73, 0x20, 0x49, 0x44, 0x2e, 0x2a, 0x1c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x0a, 0x0c,
	0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2e, 0x3a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x7b,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x69, 0x64, 0x7d,
	0x42, 0xbd, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2f, 0x76, 0x32, 0x3b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x92, 0x41, 0x82, 0x01, 0x2a, 0x01,
	0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x40, 0x0a, 0x1a, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x13,
	0x08, 0x02, 0x1a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x02, 0x0a, 0x22, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x16, 0x08, 0x02, 0x1a, 0x10, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x2d, 0x69, 0x64, 0x20, 0x02, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54,
	0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_tenant_v2_policy_state_proto_rawDescOnce sync.Once
	file_aserto_tenant_v2_policy_state_proto_rawDescData = file_aserto_tenant_v2_policy_state_proto_rawDesc
)

func file_aserto_tenant_v2_policy_state_proto_rawDescGZIP() []byte {
	file_aserto_tenant_v2_policy_state_proto_rawDescOnce.Do(func() {
		file_aserto_tenant_v2_policy_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_tenant_v2_policy_state_proto_rawDescData)
	})
	return file_aserto_tenant_v2_policy_state_proto_rawDescData
}

var file_aserto_tenant_v2_policy_state_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_aserto_tenant_v2_policy_state_proto_goTypes = []interface{}{
	(*GetPolicyStateRequest)(nil),  // 0: aserto.tenant.v2.GetPolicyStateRequest
	(*GetPolicyStateResponse)(nil), // 1: aserto.tenant.v2.GetPolicyStateResponse
	(*SetPolicyStateRequest)(nil),  // 2: aserto.tenant.v2.SetPolicyStateRequest
	(*SetPolicyStateResponse)(nil), // 3: aserto.tenant.v2.SetPolicyStateResponse
	(*v2.PolicyState)(nil),         // 4: aserto.api.v2.PolicyState
}
var file_aserto_tenant_v2_policy_state_proto_depIdxs = []int32{
	4, // 0: aserto.tenant.v2.GetPolicyStateResponse.state:type_name -> aserto.api.v2.PolicyState
	4, // 1: aserto.tenant.v2.SetPolicyStateRequest.state:type_name -> aserto.api.v2.PolicyState
	0, // 2: aserto.tenant.v2.PolicyState.GetPolicyState:input_type -> aserto.tenant.v2.GetPolicyStateRequest
	2, // 3: aserto.tenant.v2.PolicyState.SetPolicyState:input_type -> aserto.tenant.v2.SetPolicyStateRequest
	1, // 4: aserto.tenant.v2.PolicyState.GetPolicyState:output_type -> aserto.tenant.v2.GetPolicyStateResponse
	3, // 5: aserto.tenant.v2.PolicyState.SetPolicyState:output_type -> aserto.tenant.v2.SetPolicyStateResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_aserto_tenant_v2_policy_state_proto_init() }
func file_aserto_tenant_v2_policy_state_proto_init() {
	if File_aserto_tenant_v2_policy_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_tenant_v2_policy_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPolicyStateRequest); i {
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
		file_aserto_tenant_v2_policy_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPolicyStateResponse); i {
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
		file_aserto_tenant_v2_policy_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPolicyStateRequest); i {
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
		file_aserto_tenant_v2_policy_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPolicyStateResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aserto_tenant_v2_policy_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_tenant_v2_policy_state_proto_goTypes,
		DependencyIndexes: file_aserto_tenant_v2_policy_state_proto_depIdxs,
		MessageInfos:      file_aserto_tenant_v2_policy_state_proto_msgTypes,
	}.Build()
	File_aserto_tenant_v2_policy_state_proto = out.File
	file_aserto_tenant_v2_policy_state_proto_rawDesc = nil
	file_aserto_tenant_v2_policy_state_proto_goTypes = nil
	file_aserto_tenant_v2_policy_state_proto_depIdxs = nil
}
