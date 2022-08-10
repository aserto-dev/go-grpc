// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: aserto/tenant/v2/source.proto

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

type CreateSourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source           *v2.Source `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	WorkflowFileName string     `protobuf:"bytes,2,opt,name=workflow_file_name,json=workflowFileName,proto3" json:"workflow_file_name,omitempty"`
}

func (x *CreateSourceRequest) Reset() {
	*x = CreateSourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_v2_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSourceRequest) ProtoMessage() {}

func (x *CreateSourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSourceRequest.ProtoReflect.Descriptor instead.
func (*CreateSourceRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_source_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSourceRequest) GetSource() *v2.Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *CreateSourceRequest) GetWorkflowFileName() string {
	if x != nil {
		return x.WorkflowFileName
	}
	return ""
}

type CreateSourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *v2.Source `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *CreateSourceResponse) Reset() {
	*x = CreateSourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_v2_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSourceResponse) ProtoMessage() {}

func (x *CreateSourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSourceResponse.ProtoReflect.Descriptor instead.
func (*CreateSourceResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_source_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSourceResponse) GetSource() *v2.Source {
	if x != nil {
		return x.Source
	}
	return nil
}

type GetSourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyId string `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
}

func (x *GetSourceRequest) Reset() {
	*x = GetSourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_v2_source_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSourceRequest) ProtoMessage() {}

func (x *GetSourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSourceRequest.ProtoReflect.Descriptor instead.
func (*GetSourceRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_source_proto_rawDescGZIP(), []int{2}
}

func (x *GetSourceRequest) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

type GetSourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *v2.Source `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *GetSourceResponse) Reset() {
	*x = GetSourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_v2_source_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSourceResponse) ProtoMessage() {}

func (x *GetSourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSourceResponse.ProtoReflect.Descriptor instead.
func (*GetSourceResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_source_proto_rawDescGZIP(), []int{3}
}

func (x *GetSourceResponse) GetSource() *v2.Source {
	if x != nil {
		return x.Source
	}
	return nil
}

type DeleteSourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyId string `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
}

func (x *DeleteSourceRequest) Reset() {
	*x = DeleteSourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_v2_source_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSourceRequest) ProtoMessage() {}

func (x *DeleteSourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSourceRequest.ProtoReflect.Descriptor instead.
func (*DeleteSourceRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_source_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSourceRequest) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

type DeleteSourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSourceResponse) Reset() {
	*x = DeleteSourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_v2_source_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSourceResponse) ProtoMessage() {}

func (x *DeleteSourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSourceResponse.ProtoReflect.Descriptor instead.
func (*DeleteSourceResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_source_proto_rawDescGZIP(), []int{5}
}

type UpdateSourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *v2.Source `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Fields *v1.Fields `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
}

func (x *UpdateSourceRequest) Reset() {
	*x = UpdateSourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_v2_source_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSourceRequest) ProtoMessage() {}

func (x *UpdateSourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSourceRequest.ProtoReflect.Descriptor instead.
func (*UpdateSourceRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_source_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateSourceRequest) GetSource() *v2.Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *UpdateSourceRequest) GetFields() *v1.Fields {
	if x != nil {
		return x.Fields
	}
	return nil
}

type UpdateSourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *v2.Source `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *UpdateSourceResponse) Reset() {
	*x = UpdateSourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_v2_source_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSourceResponse) ProtoMessage() {}

func (x *UpdateSourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSourceResponse.ProtoReflect.Descriptor instead.
func (*UpdateSourceResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_source_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateSourceResponse) GetSource() *v2.Source {
	if x != nil {
		return x.Source
	}
	return nil
}

var File_aserto_tenant_v2_source_proto protoreflect.FileDescriptor

var file_aserto_tenant_v2_source_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f,
	0x76, 0x32, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76,
	0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22,
	0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64,
	0x22, 0x42, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0x38, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04,
	0x90, 0xb5, 0x18, 0x04, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x22, 0x16,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x73, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x45, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x32, 0xa2, 0x07, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0xe6, 0x01,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e,
	0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8f, 0x01, 0x92, 0x41, 0x69, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x0b, 0x47, 0x65, 0x74, 0x20, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x1a,
	0x26, 0x47, 0x65, 0x74, 0x20, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x63, 0x6f, 0x64, 0x65,
	0x20, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x2a, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x67, 0x65, 0x74, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03,
	0x4a, 0x57, 0x54, 0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xdc, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76,
	0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7d, 0x92, 0x41, 0x5b, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x1a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x2a, 0x14, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x17, 0x0a,
	0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0xdc, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7d, 0x92, 0x41, 0x57, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x20, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x1a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x20, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x2a, 0x14, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57,
	0x54, 0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x2a, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0xf0, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x90, 0x01, 0x92, 0x41, 0x5b, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x1a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x2a, 0x14, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x17, 0x0a, 0x07,
	0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x32, 0x22, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0xbd, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65,
	0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x3b, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x92, 0x41, 0x82, 0x01, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x40, 0x0a, 0x1a,
	0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x13, 0x08, 0x02, 0x1a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x0a, 0x22, 0x0a, 0x08, 0x54, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x08, 0x02, 0x1a, 0x10, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2d, 0x69, 0x64, 0x20, 0x02, 0x62, 0x17,
	0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_tenant_v2_source_proto_rawDescOnce sync.Once
	file_aserto_tenant_v2_source_proto_rawDescData = file_aserto_tenant_v2_source_proto_rawDesc
)

func file_aserto_tenant_v2_source_proto_rawDescGZIP() []byte {
	file_aserto_tenant_v2_source_proto_rawDescOnce.Do(func() {
		file_aserto_tenant_v2_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_tenant_v2_source_proto_rawDescData)
	})
	return file_aserto_tenant_v2_source_proto_rawDescData
}

var file_aserto_tenant_v2_source_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_aserto_tenant_v2_source_proto_goTypes = []interface{}{
	(*CreateSourceRequest)(nil),  // 0: aserto.tenant.v2.CreateSourceRequest
	(*CreateSourceResponse)(nil), // 1: aserto.tenant.v2.CreateSourceResponse
	(*GetSourceRequest)(nil),     // 2: aserto.tenant.v2.GetSourceRequest
	(*GetSourceResponse)(nil),    // 3: aserto.tenant.v2.GetSourceResponse
	(*DeleteSourceRequest)(nil),  // 4: aserto.tenant.v2.DeleteSourceRequest
	(*DeleteSourceResponse)(nil), // 5: aserto.tenant.v2.DeleteSourceResponse
	(*UpdateSourceRequest)(nil),  // 6: aserto.tenant.v2.UpdateSourceRequest
	(*UpdateSourceResponse)(nil), // 7: aserto.tenant.v2.UpdateSourceResponse
	(*v2.Source)(nil),            // 8: aserto.api.v2.Source
	(*v1.Fields)(nil),            // 9: aserto.api.v1.Fields
}
var file_aserto_tenant_v2_source_proto_depIdxs = []int32{
	8,  // 0: aserto.tenant.v2.CreateSourceRequest.source:type_name -> aserto.api.v2.Source
	8,  // 1: aserto.tenant.v2.CreateSourceResponse.source:type_name -> aserto.api.v2.Source
	8,  // 2: aserto.tenant.v2.GetSourceResponse.source:type_name -> aserto.api.v2.Source
	8,  // 3: aserto.tenant.v2.UpdateSourceRequest.source:type_name -> aserto.api.v2.Source
	9,  // 4: aserto.tenant.v2.UpdateSourceRequest.fields:type_name -> aserto.api.v1.Fields
	8,  // 5: aserto.tenant.v2.UpdateSourceResponse.source:type_name -> aserto.api.v2.Source
	2,  // 6: aserto.tenant.v2.Source.GetSource:input_type -> aserto.tenant.v2.GetSourceRequest
	0,  // 7: aserto.tenant.v2.Source.CreateSource:input_type -> aserto.tenant.v2.CreateSourceRequest
	4,  // 8: aserto.tenant.v2.Source.DeleteSource:input_type -> aserto.tenant.v2.DeleteSourceRequest
	6,  // 9: aserto.tenant.v2.Source.UpdateSource:input_type -> aserto.tenant.v2.UpdateSourceRequest
	3,  // 10: aserto.tenant.v2.Source.GetSource:output_type -> aserto.tenant.v2.GetSourceResponse
	1,  // 11: aserto.tenant.v2.Source.CreateSource:output_type -> aserto.tenant.v2.CreateSourceResponse
	5,  // 12: aserto.tenant.v2.Source.DeleteSource:output_type -> aserto.tenant.v2.DeleteSourceResponse
	7,  // 13: aserto.tenant.v2.Source.UpdateSource:output_type -> aserto.tenant.v2.UpdateSourceResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_aserto_tenant_v2_source_proto_init() }
func file_aserto_tenant_v2_source_proto_init() {
	if File_aserto_tenant_v2_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_tenant_v2_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSourceRequest); i {
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
		file_aserto_tenant_v2_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSourceResponse); i {
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
		file_aserto_tenant_v2_source_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSourceRequest); i {
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
		file_aserto_tenant_v2_source_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSourceResponse); i {
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
		file_aserto_tenant_v2_source_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSourceRequest); i {
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
		file_aserto_tenant_v2_source_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSourceResponse); i {
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
		file_aserto_tenant_v2_source_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSourceRequest); i {
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
		file_aserto_tenant_v2_source_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSourceResponse); i {
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
			RawDescriptor: file_aserto_tenant_v2_source_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_tenant_v2_source_proto_goTypes,
		DependencyIndexes: file_aserto_tenant_v2_source_proto_depIdxs,
		MessageInfos:      file_aserto_tenant_v2_source_proto_msgTypes,
	}.Build()
	File_aserto_tenant_v2_source_proto = out.File
	file_aserto_tenant_v2_source_proto_rawDesc = nil
	file_aserto_tenant_v2_source_proto_goTypes = nil
	file_aserto_tenant_v2_source_proto_depIdxs = nil
}
