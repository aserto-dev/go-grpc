// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: aserto/tenant/onboarding/v1/onboarding.proto

package onboarding

import (
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

type TenantAvailability int32

const (
	TenantAvailability_TENANT_AVAILABILITY_UNKNOWN     TenantAvailability = 0
	TenantAvailability_TENANT_AVAILABILITY_AVAILABLE   TenantAvailability = 1
	TenantAvailability_TENANT_AVAILABILITY_UNAVAILABLE TenantAvailability = 2
	TenantAvailability_TENANT_AVAILABILITY_INVALID     TenantAvailability = 3
	TenantAvailability_TENANT_AVAILABILITY_PROFANE     TenantAvailability = 4
	TenantAvailability_TENANT_AVAILABILITY_RESERVED    TenantAvailability = 5
)

// Enum value maps for TenantAvailability.
var (
	TenantAvailability_name = map[int32]string{
		0: "TENANT_AVAILABILITY_UNKNOWN",
		1: "TENANT_AVAILABILITY_AVAILABLE",
		2: "TENANT_AVAILABILITY_UNAVAILABLE",
		3: "TENANT_AVAILABILITY_INVALID",
		4: "TENANT_AVAILABILITY_PROFANE",
		5: "TENANT_AVAILABILITY_RESERVED",
	}
	TenantAvailability_value = map[string]int32{
		"TENANT_AVAILABILITY_UNKNOWN":     0,
		"TENANT_AVAILABILITY_AVAILABLE":   1,
		"TENANT_AVAILABILITY_UNAVAILABLE": 2,
		"TENANT_AVAILABILITY_INVALID":     3,
		"TENANT_AVAILABILITY_PROFANE":     4,
		"TENANT_AVAILABILITY_RESERVED":    5,
	}
)

func (x TenantAvailability) Enum() *TenantAvailability {
	p := new(TenantAvailability)
	*p = x
	return p
}

func (x TenantAvailability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TenantAvailability) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_tenant_onboarding_v1_onboarding_proto_enumTypes[0].Descriptor()
}

func (TenantAvailability) Type() protoreflect.EnumType {
	return &file_aserto_tenant_onboarding_v1_onboarding_proto_enumTypes[0]
}

func (x TenantAvailability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TenantAvailability.Descriptor instead.
func (TenantAvailability) EnumDescriptor() ([]byte, []int) {
	return file_aserto_tenant_onboarding_v1_onboarding_proto_rawDescGZIP(), []int{0}
}

type ClaimTenantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Personal         bool   `protobuf:"varint,2,opt,name=personal,proto3" json:"personal,omitempty"`
	DefaultArtifacts bool   `protobuf:"varint,3,opt,name=default_artifacts,json=defaultArtifacts,proto3" json:"default_artifacts,omitempty"`
}

func (x *ClaimTenantRequest) Reset() {
	*x = ClaimTenantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimTenantRequest) ProtoMessage() {}

func (x *ClaimTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimTenantRequest.ProtoReflect.Descriptor instead.
func (*ClaimTenantRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_onboarding_v1_onboarding_proto_rawDescGZIP(), []int{0}
}

func (x *ClaimTenantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClaimTenantRequest) GetPersonal() bool {
	if x != nil {
		return x.Personal
	}
	return false
}

func (x *ClaimTenantRequest) GetDefaultArtifacts() bool {
	if x != nil {
		return x.DefaultArtifacts
	}
	return false
}

type ClaimTenantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ClaimTenantResponse) Reset() {
	*x = ClaimTenantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimTenantResponse) ProtoMessage() {}

func (x *ClaimTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimTenantResponse.ProtoReflect.Descriptor instead.
func (*ClaimTenantResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_onboarding_v1_onboarding_proto_rawDescGZIP(), []int{1}
}

func (x *ClaimTenantResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TenantAvailableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TenantAvailableRequest) Reset() {
	*x = TenantAvailableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantAvailableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantAvailableRequest) ProtoMessage() {}

func (x *TenantAvailableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantAvailableRequest.ProtoReflect.Descriptor instead.
func (*TenantAvailableRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_onboarding_v1_onboarding_proto_rawDescGZIP(), []int{2}
}

func (x *TenantAvailableRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TenantAvailableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Availability TenantAvailability `protobuf:"varint,1,opt,name=availability,proto3,enum=aserto.tenant.onboarding.v1.TenantAvailability" json:"availability,omitempty"`
	Reason       string             `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *TenantAvailableResponse) Reset() {
	*x = TenantAvailableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantAvailableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantAvailableResponse) ProtoMessage() {}

func (x *TenantAvailableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantAvailableResponse.ProtoReflect.Descriptor instead.
func (*TenantAvailableResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_onboarding_v1_onboarding_proto_rawDescGZIP(), []int{3}
}

func (x *TenantAvailableResponse) GetAvailability() TenantAvailability {
	if x != nil {
		return x.Availability
	}
	return TenantAvailability_TENANT_AVAILABILITY_UNKNOWN
}

func (x *TenantAvailableResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type InviteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *InviteUserRequest) Reset() {
	*x = InviteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteUserRequest) ProtoMessage() {}

func (x *InviteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteUserRequest.ProtoReflect.Descriptor instead.
func (*InviteUserRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_onboarding_v1_onboarding_proto_rawDescGZIP(), []int{4}
}

func (x *InviteUserRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type InviteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InviteUserResponse) Reset() {
	*x = InviteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteUserResponse) ProtoMessage() {}

func (x *InviteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteUserResponse.ProtoReflect.Descriptor instead.
func (*InviteUserResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_onboarding_v1_onboarding_proto_rawDescGZIP(), []int{5}
}

var File_aserto_tenant_onboarding_v1_onboarding_proto protoreflect.FileDescriptor

var file_aserto_tenant_onboarding_v1_onboarding_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f,
	0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x6e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x6f, 0x6e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x12, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x11,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x22, 0x2b, 0x0a, 0x13, 0x43, 0x6c, 0x61,
	0x69, 0x6d, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x90, 0xb5,
	0x18, 0x02, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x16, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x17, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x53, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x38, 0x0a,
	0x11, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x90, 0xb5, 0x18, 0x01, 0x52, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0xe1, 0x01,
	0x0a, 0x12, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x5f, 0x41,
	0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x5f,
	0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x41, 0x56, 0x41,
	0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x54, 0x45, 0x4e, 0x41,
	0x4e, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f,
	0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x1f, 0x0a,
	0x1b, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x49,
	0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x03, 0x12, 0x1f,
	0x0a, 0x1b, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x41, 0x4e, 0x45, 0x10, 0x04, 0x12,
	0x20, 0x0a, 0x1c, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x44, 0x10,
	0x05, 0x32, 0x99, 0x06, 0x0a, 0x0a, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0xea, 0x01, 0x0a, 0x0b, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x12, 0x2f, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6c, 0x61, 0x69, 0x6d, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x30, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6c, 0x61, 0x69, 0x6d, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x78, 0x92, 0x41, 0x52, 0x0a, 0x0a, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x0c, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x20, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x1a, 0x12, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x20, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x20, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x2a, 0x17, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x2e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x62,
	0x09, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d,
	0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x6e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0xa9, 0x02,
	0x0a, 0x0f, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x33, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xaa, 0x01, 0x92,
	0x41, 0x7c, 0x0a, 0x0a, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x19,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x20, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x2b, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x73, 0x20, 0x69, 0x66, 0x20, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x20, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x69, 0x73, 0x20, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x2a, 0x1b, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x62, 0x09, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x25, 0x12, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x6e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0xf1, 0x01, 0x0a, 0x0a, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2e, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x81, 0x01, 0x92, 0x41, 0x59, 0x0a,
	0x0a, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x0b, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x74, 0x6f, 0x20, 0x6a, 0x6f, 0x69, 0x6e, 0x20, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x2a, 0x16, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x09, 0x0a,
	0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01,
	0x2a, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x6e, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x42, 0x99, 0x01,
	0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x6f, 0x6e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x6e, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x92, 0x41, 0x50, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x5a, 0x1c, 0x0a, 0x1a, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x13, 0x08, 0x02, 0x1a, 0x0d, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x09,
	0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_aserto_tenant_onboarding_v1_onboarding_proto_rawDescOnce sync.Once
	file_aserto_tenant_onboarding_v1_onboarding_proto_rawDescData = file_aserto_tenant_onboarding_v1_onboarding_proto_rawDesc
)

func file_aserto_tenant_onboarding_v1_onboarding_proto_rawDescGZIP() []byte {
	file_aserto_tenant_onboarding_v1_onboarding_proto_rawDescOnce.Do(func() {
		file_aserto_tenant_onboarding_v1_onboarding_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_tenant_onboarding_v1_onboarding_proto_rawDescData)
	})
	return file_aserto_tenant_onboarding_v1_onboarding_proto_rawDescData
}

var file_aserto_tenant_onboarding_v1_onboarding_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_aserto_tenant_onboarding_v1_onboarding_proto_goTypes = []interface{}{
	(TenantAvailability)(0),         // 0: aserto.tenant.onboarding.v1.TenantAvailability
	(*ClaimTenantRequest)(nil),      // 1: aserto.tenant.onboarding.v1.ClaimTenantRequest
	(*ClaimTenantResponse)(nil),     // 2: aserto.tenant.onboarding.v1.ClaimTenantResponse
	(*TenantAvailableRequest)(nil),  // 3: aserto.tenant.onboarding.v1.TenantAvailableRequest
	(*TenantAvailableResponse)(nil), // 4: aserto.tenant.onboarding.v1.TenantAvailableResponse
	(*InviteUserRequest)(nil),       // 5: aserto.tenant.onboarding.v1.InviteUserRequest
	(*InviteUserResponse)(nil),      // 6: aserto.tenant.onboarding.v1.InviteUserResponse
}
var file_aserto_tenant_onboarding_v1_onboarding_proto_depIdxs = []int32{
	0, // 0: aserto.tenant.onboarding.v1.TenantAvailableResponse.availability:type_name -> aserto.tenant.onboarding.v1.TenantAvailability
	1, // 1: aserto.tenant.onboarding.v1.Onboarding.ClaimTenant:input_type -> aserto.tenant.onboarding.v1.ClaimTenantRequest
	3, // 2: aserto.tenant.onboarding.v1.Onboarding.TenantAvailable:input_type -> aserto.tenant.onboarding.v1.TenantAvailableRequest
	5, // 3: aserto.tenant.onboarding.v1.Onboarding.InviteUser:input_type -> aserto.tenant.onboarding.v1.InviteUserRequest
	2, // 4: aserto.tenant.onboarding.v1.Onboarding.ClaimTenant:output_type -> aserto.tenant.onboarding.v1.ClaimTenantResponse
	4, // 5: aserto.tenant.onboarding.v1.Onboarding.TenantAvailable:output_type -> aserto.tenant.onboarding.v1.TenantAvailableResponse
	6, // 6: aserto.tenant.onboarding.v1.Onboarding.InviteUser:output_type -> aserto.tenant.onboarding.v1.InviteUserResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aserto_tenant_onboarding_v1_onboarding_proto_init() }
func file_aserto_tenant_onboarding_v1_onboarding_proto_init() {
	if File_aserto_tenant_onboarding_v1_onboarding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimTenantRequest); i {
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
		file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimTenantResponse); i {
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
		file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantAvailableRequest); i {
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
		file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantAvailableResponse); i {
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
		file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteUserRequest); i {
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
		file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteUserResponse); i {
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
			RawDescriptor: file_aserto_tenant_onboarding_v1_onboarding_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_tenant_onboarding_v1_onboarding_proto_goTypes,
		DependencyIndexes: file_aserto_tenant_onboarding_v1_onboarding_proto_depIdxs,
		EnumInfos:         file_aserto_tenant_onboarding_v1_onboarding_proto_enumTypes,
		MessageInfos:      file_aserto_tenant_onboarding_v1_onboarding_proto_msgTypes,
	}.Build()
	File_aserto_tenant_onboarding_v1_onboarding_proto = out.File
	file_aserto_tenant_onboarding_v1_onboarding_proto_rawDesc = nil
	file_aserto_tenant_onboarding_v1_onboarding_proto_goTypes = nil
	file_aserto_tenant_onboarding_v1_onboarding_proto_depIdxs = nil
}
