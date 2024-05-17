// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: aserto/tenant/account/v1/account.proto

package account

import (
	v1 "github.com/aserto-dev/go-grpc/aserto/api/v1"
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

type SignupAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Recaptcha string `protobuf:"bytes,2,opt,name=recaptcha,proto3" json:"recaptcha,omitempty"`
}

func (x *SignupAccountRequest) Reset() {
	*x = SignupAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupAccountRequest) ProtoMessage() {}

func (x *SignupAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupAccountRequest.ProtoReflect.Descriptor instead.
func (*SignupAccountRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_account_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *SignupAccountRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignupAccountRequest) GetRecaptcha() string {
	if x != nil {
		return x.Recaptcha
	}
	return ""
}

type SignupAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignupAccountResponse) Reset() {
	*x = SignupAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupAccountResponse) ProtoMessage() {}

func (x *SignupAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupAccountResponse.ProtoReflect.Descriptor instead.
func (*SignupAccountResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_account_v1_account_proto_rawDescGZIP(), []int{1}
}

type GetAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAccountRequest) Reset() {
	*x = GetAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountRequest) ProtoMessage() {}

func (x *GetAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountRequest.ProtoReflect.Descriptor instead.
func (*GetAccountRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_account_v1_account_proto_rawDescGZIP(), []int{2}
}

type GetAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *v1.Account `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GetAccountResponse) Reset() {
	*x = GetAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountResponse) ProtoMessage() {}

func (x *GetAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountResponse.ProtoReflect.Descriptor instead.
func (*GetAccountResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_account_v1_account_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccountResponse) GetResult() *v1.Account {
	if x != nil {
		return x.Result
	}
	return nil
}

type ListInvitesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListInvitesRequest) Reset() {
	*x = ListInvitesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvitesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvitesRequest) ProtoMessage() {}

func (x *ListInvitesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvitesRequest.ProtoReflect.Descriptor instead.
func (*ListInvitesRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_account_v1_account_proto_rawDescGZIP(), []int{4}
}

type TenantInvite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invite *v1.Invite `protobuf:"bytes,1,opt,name=invite,proto3" json:"invite,omitempty"`
	Tenant *v1.Tenant `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (x *TenantInvite) Reset() {
	*x = TenantInvite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantInvite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantInvite) ProtoMessage() {}

func (x *TenantInvite) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantInvite.ProtoReflect.Descriptor instead.
func (*TenantInvite) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_account_v1_account_proto_rawDescGZIP(), []int{5}
}

func (x *TenantInvite) GetInvite() *v1.Invite {
	if x != nil {
		return x.Invite
	}
	return nil
}

func (x *TenantInvite) GetTenant() *v1.Tenant {
	if x != nil {
		return x.Tenant
	}
	return nil
}

type ListInvitesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*TenantInvite `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListInvitesResponse) Reset() {
	*x = ListInvitesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvitesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvitesResponse) ProtoMessage() {}

func (x *ListInvitesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvitesResponse.ProtoReflect.Descriptor instead.
func (*ListInvitesResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_account_v1_account_proto_rawDescGZIP(), []int{6}
}

func (x *ListInvitesResponse) GetResults() []*TenantInvite {
	if x != nil {
		return x.Results
	}
	return nil
}

type UpdateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *v1.Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *UpdateAccountRequest) Reset() {
	*x = UpdateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountRequest) ProtoMessage() {}

func (x *UpdateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountRequest.ProtoReflect.Descriptor instead.
func (*UpdateAccountRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_account_v1_account_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateAccountRequest) GetAccount() *v1.Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type UpdateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAccountResponse) Reset() {
	*x = UpdateAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountResponse) ProtoMessage() {}

func (x *UpdateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_account_v1_account_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountResponse.ProtoReflect.Descriptor instead.
func (*UpdateAccountResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_account_v1_account_proto_rawDescGZIP(), []int{8}
}

var File_aserto_tenant_account_v1_account_proto protoreflect.FileDescriptor

var file_aserto_tenant_account_v1_account_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x14, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68,
	0x61, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x44, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6c, 0x0a, 0x0c, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x52, 0x06, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x57, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x22, 0x48, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x17, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd2, 0x08, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0xaa, 0x02, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x2b, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc0, 0x01, 0x92, 0x41,
	0xa5, 0x01, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x47, 0x65, 0x74,
	0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x1a, 0x68, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x69,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x67, 0x69,
	0x76, 0x65, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x3a, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x69, 0x64, 0x2c, 0x20, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x20, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x20, 0x61, 0x6e, 0x64, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x20, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x20, 0x69, 0x64, 0x2e, 0x2a, 0x13, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x09, 0x0a, 0x07,
	0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x87,
	0x02, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x12, 0x2c,
	0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9a, 0x01, 0x92, 0x41,
	0x78, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x20, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x1a, 0x3e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x73, 0x20, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x20, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x20, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x2a, 0x14, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x62, 0x09,
	0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12,
	0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x12, 0x86, 0x02, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x2e, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x93, 0x01, 0x92, 0x41,
	0x70, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x32, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x2a, 0x16,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x09, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12,
	0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x86, 0x02, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x2e, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x93, 0x01, 0x92, 0x41, 0x6f, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x12, 0x16, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x53, 0x69, 0x67, 0x6e,
	0x20, 0x75, 0x70, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x6e, 0x20, 0x41, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x20, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x2a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x62, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a,
	0x01, 0x2a, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x42, 0x93, 0x01, 0x92, 0x41, 0x50,
	0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x1c, 0x0a, 0x1a, 0x0a, 0x03, 0x4a, 0x57, 0x54,
	0x12, 0x13, 0x08, 0x02, 0x1a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x09, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00,
	0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_tenant_account_v1_account_proto_rawDescOnce sync.Once
	file_aserto_tenant_account_v1_account_proto_rawDescData = file_aserto_tenant_account_v1_account_proto_rawDesc
)

func file_aserto_tenant_account_v1_account_proto_rawDescGZIP() []byte {
	file_aserto_tenant_account_v1_account_proto_rawDescOnce.Do(func() {
		file_aserto_tenant_account_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_tenant_account_v1_account_proto_rawDescData)
	})
	return file_aserto_tenant_account_v1_account_proto_rawDescData
}

var file_aserto_tenant_account_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_aserto_tenant_account_v1_account_proto_goTypes = []interface{}{
	(*SignupAccountRequest)(nil),  // 0: aserto.tenant.account.v1.SignupAccountRequest
	(*SignupAccountResponse)(nil), // 1: aserto.tenant.account.v1.SignupAccountResponse
	(*GetAccountRequest)(nil),     // 2: aserto.tenant.account.v1.GetAccountRequest
	(*GetAccountResponse)(nil),    // 3: aserto.tenant.account.v1.GetAccountResponse
	(*ListInvitesRequest)(nil),    // 4: aserto.tenant.account.v1.ListInvitesRequest
	(*TenantInvite)(nil),          // 5: aserto.tenant.account.v1.TenantInvite
	(*ListInvitesResponse)(nil),   // 6: aserto.tenant.account.v1.ListInvitesResponse
	(*UpdateAccountRequest)(nil),  // 7: aserto.tenant.account.v1.UpdateAccountRequest
	(*UpdateAccountResponse)(nil), // 8: aserto.tenant.account.v1.UpdateAccountResponse
	(*v1.Account)(nil),            // 9: aserto.api.v1.Account
	(*v1.Invite)(nil),             // 10: aserto.api.v1.Invite
	(*v1.Tenant)(nil),             // 11: aserto.api.v1.Tenant
}
var file_aserto_tenant_account_v1_account_proto_depIdxs = []int32{
	9,  // 0: aserto.tenant.account.v1.GetAccountResponse.result:type_name -> aserto.api.v1.Account
	10, // 1: aserto.tenant.account.v1.TenantInvite.invite:type_name -> aserto.api.v1.Invite
	11, // 2: aserto.tenant.account.v1.TenantInvite.tenant:type_name -> aserto.api.v1.Tenant
	5,  // 3: aserto.tenant.account.v1.ListInvitesResponse.results:type_name -> aserto.tenant.account.v1.TenantInvite
	9,  // 4: aserto.tenant.account.v1.UpdateAccountRequest.account:type_name -> aserto.api.v1.Account
	2,  // 5: aserto.tenant.account.v1.Account.GetAccount:input_type -> aserto.tenant.account.v1.GetAccountRequest
	4,  // 6: aserto.tenant.account.v1.Account.ListInvites:input_type -> aserto.tenant.account.v1.ListInvitesRequest
	7,  // 7: aserto.tenant.account.v1.Account.UpdateAccount:input_type -> aserto.tenant.account.v1.UpdateAccountRequest
	0,  // 8: aserto.tenant.account.v1.Account.SignupAccount:input_type -> aserto.tenant.account.v1.SignupAccountRequest
	3,  // 9: aserto.tenant.account.v1.Account.GetAccount:output_type -> aserto.tenant.account.v1.GetAccountResponse
	6,  // 10: aserto.tenant.account.v1.Account.ListInvites:output_type -> aserto.tenant.account.v1.ListInvitesResponse
	8,  // 11: aserto.tenant.account.v1.Account.UpdateAccount:output_type -> aserto.tenant.account.v1.UpdateAccountResponse
	1,  // 12: aserto.tenant.account.v1.Account.SignupAccount:output_type -> aserto.tenant.account.v1.SignupAccountResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_aserto_tenant_account_v1_account_proto_init() }
func file_aserto_tenant_account_v1_account_proto_init() {
	if File_aserto_tenant_account_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_tenant_account_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupAccountRequest); i {
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
		file_aserto_tenant_account_v1_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupAccountResponse); i {
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
		file_aserto_tenant_account_v1_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountRequest); i {
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
		file_aserto_tenant_account_v1_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountResponse); i {
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
		file_aserto_tenant_account_v1_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvitesRequest); i {
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
		file_aserto_tenant_account_v1_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantInvite); i {
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
		file_aserto_tenant_account_v1_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvitesResponse); i {
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
		file_aserto_tenant_account_v1_account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccountRequest); i {
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
		file_aserto_tenant_account_v1_account_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccountResponse); i {
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
			RawDescriptor: file_aserto_tenant_account_v1_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_tenant_account_v1_account_proto_goTypes,
		DependencyIndexes: file_aserto_tenant_account_v1_account_proto_depIdxs,
		MessageInfos:      file_aserto_tenant_account_v1_account_proto_msgTypes,
	}.Build()
	File_aserto_tenant_account_v1_account_proto = out.File
	file_aserto_tenant_account_v1_account_proto_rawDesc = nil
	file_aserto_tenant_account_v1_account_proto_goTypes = nil
	file_aserto_tenant_account_v1_account_proto_depIdxs = nil
}
