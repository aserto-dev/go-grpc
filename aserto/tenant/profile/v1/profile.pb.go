// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: aserto/tenant/profile/v1/profile.proto

package profile

import (
	v1 "github.com/aserto-dev/go-grpc/aserto/api/v1"
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

type GetProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetProfileRequest) Reset() {
	*x = GetProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileRequest) ProtoMessage() {}

func (x *GetProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileRequest.ProtoReflect.Descriptor instead.
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_profile_v1_profile_proto_rawDescGZIP(), []int{0}
}

type GetProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant *v1.Tenant `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (x *GetProfileResponse) Reset() {
	*x = GetProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileResponse) ProtoMessage() {}

func (x *GetProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileResponse.ProtoReflect.Descriptor instead.
func (*GetProfileResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_profile_v1_profile_proto_rawDescGZIP(), []int{1}
}

func (x *GetProfileResponse) GetTenant() *v1.Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

type GetInvitesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetInvitesRequest) Reset() {
	*x = GetInvitesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvitesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvitesRequest) ProtoMessage() {}

func (x *GetInvitesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvitesRequest.ProtoReflect.Descriptor instead.
func (*GetInvitesRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_profile_v1_profile_proto_rawDescGZIP(), []int{2}
}

type GetInvitesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invites []*v1.Invite `protobuf:"bytes,1,rep,name=invites,proto3" json:"invites,omitempty"`
}

func (x *GetInvitesResponse) Reset() {
	*x = GetInvitesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvitesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvitesResponse) ProtoMessage() {}

func (x *GetInvitesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvitesResponse.ProtoReflect.Descriptor instead.
func (*GetInvitesResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_profile_v1_profile_proto_rawDescGZIP(), []int{3}
}

func (x *GetInvitesResponse) GetInvites() []*v1.Invite {
	if x != nil {
		return x.Invites
	}
	return nil
}

type InviteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"` // email from auth0
	Role  string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`   // role assigned to the user on accept
}

func (x *InviteUserRequest) Reset() {
	*x = InviteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteUserRequest) ProtoMessage() {}

func (x *InviteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[4]
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
	return file_aserto_tenant_profile_v1_profile_proto_rawDescGZIP(), []int{4}
}

func (x *InviteUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *InviteUserRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type InviteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InviteUserResponse) Reset() {
	*x = InviteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteUserResponse) ProtoMessage() {}

func (x *InviteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[5]
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
	return file_aserto_tenant_profile_v1_profile_proto_rawDescGZIP(), []int{5}
}

func (x *InviteUserResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RespondToInviteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status v1.InviteStatus `protobuf:"varint,2,opt,name=status,proto3,enum=aserto.api.v1.InviteStatus" json:"status,omitempty"`
}

func (x *RespondToInviteRequest) Reset() {
	*x = RespondToInviteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespondToInviteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespondToInviteRequest) ProtoMessage() {}

func (x *RespondToInviteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespondToInviteRequest.ProtoReflect.Descriptor instead.
func (*RespondToInviteRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_profile_v1_profile_proto_rawDescGZIP(), []int{6}
}

func (x *RespondToInviteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RespondToInviteRequest) GetStatus() v1.InviteStatus {
	if x != nil {
		return x.Status
	}
	return v1.InviteStatus(0)
}

type RespondToInviteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RespondToInviteResponse) Reset() {
	*x = RespondToInviteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespondToInviteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespondToInviteResponse) ProtoMessage() {}

func (x *RespondToInviteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespondToInviteResponse.ProtoReflect.Descriptor instead.
func (*RespondToInviteResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_profile_v1_profile_proto_rawDescGZIP(), []int{7}
}

type RemoveMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *RemoveMemberRequest) Reset() {
	*x = RemoveMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveMemberRequest) ProtoMessage() {}

func (x *RemoveMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveMemberRequest.ProtoReflect.Descriptor instead.
func (*RemoveMemberRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_profile_v1_profile_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveMemberRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type RemoveMemberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveMemberResponse) Reset() {
	*x = RemoveMemberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveMemberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveMemberResponse) ProtoMessage() {}

func (x *RemoveMemberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_profile_v1_profile_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveMemberResponse.ProtoReflect.Descriptor instead.
func (*RemoveMemberResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_profile_v1_profile_proto_rawDescGZIP(), []int{9}
}

var File_aserto_tenant_profile_v1_profile_proto protoreflect.FileDescriptor

var file_aserto_tenant_profile_v1_profile_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22,
	0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x69, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x22, 0x3d, 0x0a, 0x11, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x2a, 0x0a, 0x12, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x90, 0xb5,
	0x18, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x63, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x64, 0x54, 0x6f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x90, 0xb5,
	0x18, 0x08, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x54, 0x6f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0x90, 0xb5, 0x18, 0x01, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xbc, 0x0a, 0x0a, 0x07, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0xf3, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2b, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x89, 0x01, 0x92, 0x41, 0x68, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x12, 0x47, 0x65, 0x74, 0x20, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x1a, 0x1b, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x2a, 0x13, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00,
	0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x8b, 0x02, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x12, 0x2b, 0x2e, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa1, 0x01, 0x92, 0x41, 0x78, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0b, 0x47, 0x65, 0x74, 0x20, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x73, 0x1a, 0x32, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x20, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x2a, 0x13, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x62, 0x17, 0x0a, 0x07, 0x0a,
	0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x44, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x12, 0xfc, 0x01, 0x0a, 0x0a, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2b, 0x2e, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x92, 0x01, 0x92, 0x41, 0x66, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x15, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x20, 0x74, 0x6f, 0x20, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x1a, 0x16, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x2e, 0x2a, 0x13, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54,
	0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x12, 0xa0, 0x02, 0x0a, 0x0f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x64, 0x54, 0x6f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x30, 0x2e,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64,
	0x54, 0x6f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x64, 0x54, 0x6f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xa7, 0x01, 0x92, 0x41, 0x76, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x18, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x20, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x1a, 0x1d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x20, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x2a, 0x19, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00,
	0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x8a, 0x02, 0x0a,
	0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2d, 0x2e,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9a, 0x01, 0x92,
	0x41, 0x64, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0d, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x20, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x1a, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x20, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x2a, 0x15, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x62, 0x17, 0x0a,
	0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x2a, 0x2b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0xc6, 0x01, 0x5a, 0x3e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d,
	0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x92, 0x41, 0x82, 0x01,
	0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x40, 0x0a, 0x1a, 0x0a, 0x03, 0x4a, 0x57, 0x54,
	0x12, 0x13, 0x08, 0x02, 0x1a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x02, 0x0a, 0x22, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x16, 0x08, 0x02, 0x1a, 0x10, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x2d, 0x69, 0x64, 0x20, 0x02, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03, 0x4a,
	0x57, 0x54, 0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_tenant_profile_v1_profile_proto_rawDescOnce sync.Once
	file_aserto_tenant_profile_v1_profile_proto_rawDescData = file_aserto_tenant_profile_v1_profile_proto_rawDesc
)

func file_aserto_tenant_profile_v1_profile_proto_rawDescGZIP() []byte {
	file_aserto_tenant_profile_v1_profile_proto_rawDescOnce.Do(func() {
		file_aserto_tenant_profile_v1_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_tenant_profile_v1_profile_proto_rawDescData)
	})
	return file_aserto_tenant_profile_v1_profile_proto_rawDescData
}

var file_aserto_tenant_profile_v1_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_aserto_tenant_profile_v1_profile_proto_goTypes = []interface{}{
	(*GetProfileRequest)(nil),       // 0: aserto.tenant.profile.v1.GetProfileRequest
	(*GetProfileResponse)(nil),      // 1: aserto.tenant.profile.v1.GetProfileResponse
	(*GetInvitesRequest)(nil),       // 2: aserto.tenant.profile.v1.GetInvitesRequest
	(*GetInvitesResponse)(nil),      // 3: aserto.tenant.profile.v1.GetInvitesResponse
	(*InviteUserRequest)(nil),       // 4: aserto.tenant.profile.v1.InviteUserRequest
	(*InviteUserResponse)(nil),      // 5: aserto.tenant.profile.v1.InviteUserResponse
	(*RespondToInviteRequest)(nil),  // 6: aserto.tenant.profile.v1.RespondToInviteRequest
	(*RespondToInviteResponse)(nil), // 7: aserto.tenant.profile.v1.RespondToInviteResponse
	(*RemoveMemberRequest)(nil),     // 8: aserto.tenant.profile.v1.RemoveMemberRequest
	(*RemoveMemberResponse)(nil),    // 9: aserto.tenant.profile.v1.RemoveMemberResponse
	(*v1.Tenant)(nil),               // 10: aserto.api.v1.Tenant
	(*v1.Invite)(nil),               // 11: aserto.api.v1.Invite
	(v1.InviteStatus)(0),            // 12: aserto.api.v1.InviteStatus
}
var file_aserto_tenant_profile_v1_profile_proto_depIdxs = []int32{
	10, // 0: aserto.tenant.profile.v1.GetProfileResponse.tenant:type_name -> aserto.api.v1.Tenant
	11, // 1: aserto.tenant.profile.v1.GetInvitesResponse.invites:type_name -> aserto.api.v1.Invite
	12, // 2: aserto.tenant.profile.v1.RespondToInviteRequest.status:type_name -> aserto.api.v1.InviteStatus
	0,  // 3: aserto.tenant.profile.v1.Profile.GetProfile:input_type -> aserto.tenant.profile.v1.GetProfileRequest
	2,  // 4: aserto.tenant.profile.v1.Profile.GetInvites:input_type -> aserto.tenant.profile.v1.GetInvitesRequest
	4,  // 5: aserto.tenant.profile.v1.Profile.InviteUser:input_type -> aserto.tenant.profile.v1.InviteUserRequest
	6,  // 6: aserto.tenant.profile.v1.Profile.RespondToInvite:input_type -> aserto.tenant.profile.v1.RespondToInviteRequest
	8,  // 7: aserto.tenant.profile.v1.Profile.RemoveMember:input_type -> aserto.tenant.profile.v1.RemoveMemberRequest
	1,  // 8: aserto.tenant.profile.v1.Profile.GetProfile:output_type -> aserto.tenant.profile.v1.GetProfileResponse
	3,  // 9: aserto.tenant.profile.v1.Profile.GetInvites:output_type -> aserto.tenant.profile.v1.GetInvitesResponse
	5,  // 10: aserto.tenant.profile.v1.Profile.InviteUser:output_type -> aserto.tenant.profile.v1.InviteUserResponse
	7,  // 11: aserto.tenant.profile.v1.Profile.RespondToInvite:output_type -> aserto.tenant.profile.v1.RespondToInviteResponse
	9,  // 12: aserto.tenant.profile.v1.Profile.RemoveMember:output_type -> aserto.tenant.profile.v1.RemoveMemberResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_aserto_tenant_profile_v1_profile_proto_init() }
func file_aserto_tenant_profile_v1_profile_proto_init() {
	if File_aserto_tenant_profile_v1_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_tenant_profile_v1_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileRequest); i {
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
		file_aserto_tenant_profile_v1_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileResponse); i {
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
		file_aserto_tenant_profile_v1_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvitesRequest); i {
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
		file_aserto_tenant_profile_v1_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvitesResponse); i {
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
		file_aserto_tenant_profile_v1_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_aserto_tenant_profile_v1_profile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_aserto_tenant_profile_v1_profile_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespondToInviteRequest); i {
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
		file_aserto_tenant_profile_v1_profile_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespondToInviteResponse); i {
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
		file_aserto_tenant_profile_v1_profile_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveMemberRequest); i {
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
		file_aserto_tenant_profile_v1_profile_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveMemberResponse); i {
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
			RawDescriptor: file_aserto_tenant_profile_v1_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_tenant_profile_v1_profile_proto_goTypes,
		DependencyIndexes: file_aserto_tenant_profile_v1_profile_proto_depIdxs,
		MessageInfos:      file_aserto_tenant_profile_v1_profile_proto_msgTypes,
	}.Build()
	File_aserto_tenant_profile_v1_profile_proto = out.File
	file_aserto_tenant_profile_v1_profile_proto_rawDesc = nil
	file_aserto_tenant_profile_v1_profile_proto_goTypes = nil
	file_aserto_tenant_profile_v1_profile_proto_depIdxs = nil
}
