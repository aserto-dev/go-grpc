// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: tenant/v2/policy.proto

package tenant

import (
	v1 "github.com/aserto-dev/go-grpc/aserto/api/v1"
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

type CreatePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy *v2.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *CreatePolicyRequest) Reset() {
	*x = CreatePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v2_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyRequest) ProtoMessage() {}

func (x *CreatePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v2_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyRequest.ProtoReflect.Descriptor instead.
func (*CreatePolicyRequest) Descriptor() ([]byte, []int) {
	return file_tenant_v2_policy_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePolicyRequest) GetPolicy() *v2.Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type CreatePolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy *v2.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *CreatePolicyResponse) Reset() {
	*x = CreatePolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v2_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyResponse) ProtoMessage() {}

func (x *CreatePolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v2_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyResponse.ProtoReflect.Descriptor instead.
func (*CreatePolicyResponse) Descriptor() ([]byte, []int) {
	return file_tenant_v2_policy_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePolicyResponse) GetPolicy() *v2.Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

type ListPolicyFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NameStartsWith string `protobuf:"bytes,1,opt,name=name_starts_with,json=nameStartsWith,proto3" json:"name_starts_with,omitempty"`
}

func (x *ListPolicyFilter) Reset() {
	*x = ListPolicyFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v2_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPolicyFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPolicyFilter) ProtoMessage() {}

func (x *ListPolicyFilter) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v2_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPolicyFilter.ProtoReflect.Descriptor instead.
func (*ListPolicyFilter) Descriptor() ([]byte, []int) {
	return file_tenant_v2_policy_proto_rawDescGZIP(), []int{2}
}

func (x *ListPolicyFilter) GetNameStartsWith() string {
	if x != nil {
		return x.NameStartsWith
	}
	return ""
}

type ListPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   *v1.PaginationRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Filter *ListPolicyFilter     `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListPolicyRequest) Reset() {
	*x = ListPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v2_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPolicyRequest) ProtoMessage() {}

func (x *ListPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v2_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPolicyRequest.ProtoReflect.Descriptor instead.
func (*ListPolicyRequest) Descriptor() ([]byte, []int) {
	return file_tenant_v2_policy_proto_rawDescGZIP(), []int{3}
}

func (x *ListPolicyRequest) GetPage() *v1.PaginationRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListPolicyRequest) GetFilter() *ListPolicyFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*v2.Policy           `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Page    *v1.PaginationResponse `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListPolicyResponse) Reset() {
	*x = ListPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v2_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPolicyResponse) ProtoMessage() {}

func (x *ListPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v2_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPolicyResponse.ProtoReflect.Descriptor instead.
func (*ListPolicyResponse) Descriptor() ([]byte, []int) {
	return file_tenant_v2_policy_proto_rawDescGZIP(), []int{4}
}

func (x *ListPolicyResponse) GetResults() []*v2.Policy {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListPolicyResponse) GetPage() *v1.PaginationResponse {
	if x != nil {
		return x.Page
	}
	return nil
}

type DeletePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePolicyRequest) Reset() {
	*x = DeletePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v2_policy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicyRequest) ProtoMessage() {}

func (x *DeletePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v2_policy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicyRequest.ProtoReflect.Descriptor instead.
func (*DeletePolicyRequest) Descriptor() ([]byte, []int) {
	return file_tenant_v2_policy_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePolicyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePolicyResponse) Reset() {
	*x = DeletePolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v2_policy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicyResponse) ProtoMessage() {}

func (x *DeletePolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v2_policy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicyResponse.ProtoReflect.Descriptor instead.
func (*DeletePolicyResponse) Descriptor() ([]byte, []int) {
	return file_tenant_v2_policy_proto_rawDescGZIP(), []int{6}
}

type UpdatePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy *v2.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	Fields *v1.Fields `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
}

func (x *UpdatePolicyRequest) Reset() {
	*x = UpdatePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v2_policy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePolicyRequest) ProtoMessage() {}

func (x *UpdatePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v2_policy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePolicyRequest.ProtoReflect.Descriptor instead.
func (*UpdatePolicyRequest) Descriptor() ([]byte, []int) {
	return file_tenant_v2_policy_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePolicyRequest) GetPolicy() *v2.Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

func (x *UpdatePolicyRequest) GetFields() *v1.Fields {
	if x != nil {
		return x.Fields
	}
	return nil
}

type UpdatePolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy *v2.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *UpdatePolicyResponse) Reset() {
	*x = UpdatePolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tenant_v2_policy_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePolicyResponse) ProtoMessage() {}

func (x *UpdatePolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tenant_v2_policy_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePolicyResponse.ProtoReflect.Descriptor instead.
func (*UpdatePolicyResponse) Descriptor() ([]byte, []int) {
	return file_tenant_v2_policy_proto_rawDescGZIP(), []int{8}
}

func (x *UpdatePolicyResponse) GetPolicy() *v2.Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

var File_tenant_v2_policy_proto protoreflect.FileDescriptor

var file_tenant_v2_policy_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x22, 0x45, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x3c, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x28,
	0x0a, 0x10, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x77, 0x69,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x22, 0x85, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x22, 0x7c, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x35, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x2b,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0x90, 0xb5, 0x18, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x73, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2d, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x45, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x32,
	0xf8, 0x06, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0xc7, 0x01, 0x0a, 0x0a, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x23, 0x2e, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76,
	0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6e, 0x92, 0x41, 0x53, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x0e, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x2a,
	0x12, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x0a, 0x0c,
	0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x12, 0xdd, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x25, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7e, 0x92, 0x41, 0x5b, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x1a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x2a, 0x14, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x17, 0x0a, 0x07, 0x0a,
	0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x44, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x3a, 0x06, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0xd6, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x25, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x77, 0x92, 0x41, 0x57, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x1a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x2a, 0x14, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54,
	0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xea, 0x01,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x25,
	0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76,
	0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8a, 0x01,
	0x92, 0x41, 0x5b, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x1a, 0x10, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x2a,
	0x14, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00,
	0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x26, 0x32, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x69,
	0x64, 0x7d, 0x3a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0xbd, 0x01, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x3b, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x92, 0x41, 0x82, 0x01, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a,
	0x40, 0x0a, 0x1a, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x13, 0x08, 0x02, 0x1a, 0x0d, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x0a, 0x22, 0x0a,
	0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x08, 0x02, 0x1a, 0x10, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2d, 0x69, 0x64, 0x20,
	0x02, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tenant_v2_policy_proto_rawDescOnce sync.Once
	file_tenant_v2_policy_proto_rawDescData = file_tenant_v2_policy_proto_rawDesc
)

func file_tenant_v2_policy_proto_rawDescGZIP() []byte {
	file_tenant_v2_policy_proto_rawDescOnce.Do(func() {
		file_tenant_v2_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_tenant_v2_policy_proto_rawDescData)
	})
	return file_tenant_v2_policy_proto_rawDescData
}

var file_tenant_v2_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_tenant_v2_policy_proto_goTypes = []interface{}{
	(*CreatePolicyRequest)(nil),   // 0: aserto.tenant.v2.CreatePolicyRequest
	(*CreatePolicyResponse)(nil),  // 1: aserto.tenant.v2.CreatePolicyResponse
	(*ListPolicyFilter)(nil),      // 2: aserto.tenant.v2.ListPolicyFilter
	(*ListPolicyRequest)(nil),     // 3: aserto.tenant.v2.ListPolicyRequest
	(*ListPolicyResponse)(nil),    // 4: aserto.tenant.v2.ListPolicyResponse
	(*DeletePolicyRequest)(nil),   // 5: aserto.tenant.v2.DeletePolicyRequest
	(*DeletePolicyResponse)(nil),  // 6: aserto.tenant.v2.DeletePolicyResponse
	(*UpdatePolicyRequest)(nil),   // 7: aserto.tenant.v2.UpdatePolicyRequest
	(*UpdatePolicyResponse)(nil),  // 8: aserto.tenant.v2.UpdatePolicyResponse
	(*v2.Policy)(nil),             // 9: aserto.api.v2.Policy
	(*v1.PaginationRequest)(nil),  // 10: aserto.api.v1.PaginationRequest
	(*v1.PaginationResponse)(nil), // 11: aserto.api.v1.PaginationResponse
	(*v1.Fields)(nil),             // 12: aserto.api.v1.Fields
}
var file_tenant_v2_policy_proto_depIdxs = []int32{
	9,  // 0: aserto.tenant.v2.CreatePolicyRequest.policy:type_name -> aserto.api.v2.Policy
	9,  // 1: aserto.tenant.v2.CreatePolicyResponse.policy:type_name -> aserto.api.v2.Policy
	10, // 2: aserto.tenant.v2.ListPolicyRequest.page:type_name -> aserto.api.v1.PaginationRequest
	2,  // 3: aserto.tenant.v2.ListPolicyRequest.filter:type_name -> aserto.tenant.v2.ListPolicyFilter
	9,  // 4: aserto.tenant.v2.ListPolicyResponse.results:type_name -> aserto.api.v2.Policy
	11, // 5: aserto.tenant.v2.ListPolicyResponse.page:type_name -> aserto.api.v1.PaginationResponse
	9,  // 6: aserto.tenant.v2.UpdatePolicyRequest.policy:type_name -> aserto.api.v2.Policy
	12, // 7: aserto.tenant.v2.UpdatePolicyRequest.fields:type_name -> aserto.api.v1.Fields
	9,  // 8: aserto.tenant.v2.UpdatePolicyResponse.policy:type_name -> aserto.api.v2.Policy
	3,  // 9: aserto.tenant.v2.Policy.ListPolicy:input_type -> aserto.tenant.v2.ListPolicyRequest
	0,  // 10: aserto.tenant.v2.Policy.CreatePolicy:input_type -> aserto.tenant.v2.CreatePolicyRequest
	5,  // 11: aserto.tenant.v2.Policy.DeletePolicy:input_type -> aserto.tenant.v2.DeletePolicyRequest
	7,  // 12: aserto.tenant.v2.Policy.UpdatePolicy:input_type -> aserto.tenant.v2.UpdatePolicyRequest
	4,  // 13: aserto.tenant.v2.Policy.ListPolicy:output_type -> aserto.tenant.v2.ListPolicyResponse
	1,  // 14: aserto.tenant.v2.Policy.CreatePolicy:output_type -> aserto.tenant.v2.CreatePolicyResponse
	6,  // 15: aserto.tenant.v2.Policy.DeletePolicy:output_type -> aserto.tenant.v2.DeletePolicyResponse
	8,  // 16: aserto.tenant.v2.Policy.UpdatePolicy:output_type -> aserto.tenant.v2.UpdatePolicyResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_tenant_v2_policy_proto_init() }
func file_tenant_v2_policy_proto_init() {
	if File_tenant_v2_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tenant_v2_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePolicyRequest); i {
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
		file_tenant_v2_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePolicyResponse); i {
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
		file_tenant_v2_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPolicyFilter); i {
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
		file_tenant_v2_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPolicyRequest); i {
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
		file_tenant_v2_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPolicyResponse); i {
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
		file_tenant_v2_policy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePolicyRequest); i {
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
		file_tenant_v2_policy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePolicyResponse); i {
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
		file_tenant_v2_policy_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePolicyRequest); i {
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
		file_tenant_v2_policy_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePolicyResponse); i {
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
			RawDescriptor: file_tenant_v2_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tenant_v2_policy_proto_goTypes,
		DependencyIndexes: file_tenant_v2_policy_proto_depIdxs,
		MessageInfos:      file_tenant_v2_policy_proto_msgTypes,
	}.Build()
	File_tenant_v2_policy_proto = out.File
	file_tenant_v2_policy_proto_rawDesc = nil
	file_tenant_v2_policy_proto_goTypes = nil
	file_tenant_v2_policy_proto_depIdxs = nil
}
