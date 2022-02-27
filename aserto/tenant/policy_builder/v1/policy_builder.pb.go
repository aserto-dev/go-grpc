// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.19.1
// source: aserto/tenant/policy_builder/v1/policy_builder.proto

package policy_builder

import (
	v1 "github.com/aserto-dev/go-grpc/aserto/api/v1"
	_ "github.com/aserto-dev/go-grpc/aserto/options/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreatePolicyBuilderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyBuilder  *v1.PolicyBuilder `protobuf:"bytes,1,opt,name=policy_builder,json=policyBuilder,proto3" json:"policy_builder,omitempty"`
	ForceReconnect bool              `protobuf:"varint,2,opt,name=force_reconnect,json=forceReconnect,proto3" json:"force_reconnect,omitempty"`
}

func (x *CreatePolicyBuilderRequest) Reset() {
	*x = CreatePolicyBuilderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePolicyBuilderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyBuilderRequest) ProtoMessage() {}

func (x *CreatePolicyBuilderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyBuilderRequest.ProtoReflect.Descriptor instead.
func (*CreatePolicyBuilderRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePolicyBuilderRequest) GetPolicyBuilder() *v1.PolicyBuilder {
	if x != nil {
		return x.PolicyBuilder
	}
	return nil
}

func (x *CreatePolicyBuilderRequest) GetForceReconnect() bool {
	if x != nil {
		return x.ForceReconnect
	}
	return false
}

type CreatePolicyBuilderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreatePolicyBuilderResponse) Reset() {
	*x = CreatePolicyBuilderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePolicyBuilderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyBuilderResponse) ProtoMessage() {}

func (x *CreatePolicyBuilderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyBuilderResponse.ProtoReflect.Descriptor instead.
func (*CreatePolicyBuilderResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePolicyBuilderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListPolicyBuildersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegistryRepo string `protobuf:"bytes,1,opt,name=registry_repo,json=registryRepo,proto3" json:"registry_repo,omitempty"`
	RegistryOrg  string `protobuf:"bytes,2,opt,name=registry_org,json=registryOrg,proto3" json:"registry_org,omitempty"`
}

func (x *ListPolicyBuildersRequest) Reset() {
	*x = ListPolicyBuildersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPolicyBuildersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPolicyBuildersRequest) ProtoMessage() {}

func (x *ListPolicyBuildersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPolicyBuildersRequest.ProtoReflect.Descriptor instead.
func (*ListPolicyBuildersRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescGZIP(), []int{2}
}

func (x *ListPolicyBuildersRequest) GetRegistryRepo() string {
	if x != nil {
		return x.RegistryRepo
	}
	return ""
}

func (x *ListPolicyBuildersRequest) GetRegistryOrg() string {
	if x != nil {
		return x.RegistryOrg
	}
	return ""
}

type ListPolicyBuildersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*v1.PolicyBuilder `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListPolicyBuildersResponse) Reset() {
	*x = ListPolicyBuildersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPolicyBuildersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPolicyBuildersResponse) ProtoMessage() {}

func (x *ListPolicyBuildersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPolicyBuildersResponse.ProtoReflect.Descriptor instead.
func (*ListPolicyBuildersResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescGZIP(), []int{3}
}

func (x *ListPolicyBuildersResponse) GetResults() []*v1.PolicyBuilder {
	if x != nil {
		return x.Results
	}
	return nil
}

type DeletePolicyBuilderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePolicyBuilderRequest) Reset() {
	*x = DeletePolicyBuilderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePolicyBuilderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicyBuilderRequest) ProtoMessage() {}

func (x *DeletePolicyBuilderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicyBuilderRequest.ProtoReflect.Descriptor instead.
func (*DeletePolicyBuilderRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePolicyBuilderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePolicyBuilderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *emptypb.Empty `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DeletePolicyBuilderResponse) Reset() {
	*x = DeletePolicyBuilderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePolicyBuilderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicyBuilderResponse) ProtoMessage() {}

func (x *DeletePolicyBuilderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicyBuilderResponse.ProtoReflect.Descriptor instead.
func (*DeletePolicyBuilderResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePolicyBuilderResponse) GetResult() *emptypb.Empty {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_aserto_tenant_policy_builder_v1_policy_builder_proto protoreflect.FileDescriptor

var file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDesc = []byte{
	0x0a, 0x34, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01,
	0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x52, 0x0d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x22, 0x33, 0x0a, 0x1b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x90, 0xb5, 0x18, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x63, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70,
	0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x6f, 0x72,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x4f, 0x72, 0x67, 0x22, 0x54, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x32, 0x0a, 0x1a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x90, 0xb5, 0x18, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4d,
	0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xb6, 0x07,
	0x0a, 0x0d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12,
	0xa9, 0x02, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x3a, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x99, 0x01, 0x92, 0x41, 0x72, 0x0a, 0x0e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x20, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x20, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x14, 0x4c, 0x69, 0x73,
	0x74, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x20, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x73, 0x2a, 0x1b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x62, 0x17,
	0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0xc0, 0x02, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x12, 0x3b, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3c, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xad,
	0x01, 0x92, 0x41, 0x76, 0x0a, 0x0e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x20, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x12, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x20, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x1a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x20, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x2e, 0x2a, 0x1c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08,
	0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e,
	0x3a, 0x0e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x22, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0xb5,
	0x02, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x3b, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xa2, 0x01, 0x92, 0x41, 0x76, 0x0a, 0x0e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x20,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x20,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x20, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x1a, 0x16,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x20, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x2a, 0x1c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x0a,
	0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x23, 0x2a, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xd4, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76,
	0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x92, 0x41, 0x82, 0x01, 0x2a, 0x01, 0x02, 0x32, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x5a, 0x40, 0x0a, 0x1a, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x13, 0x08, 0x02, 0x1a,
	0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02,
	0x0a, 0x22, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x08, 0x02,
	0x1a, 0x10, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2d,
	0x69, 0x64, 0x20, 0x02, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x0a,
	0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescOnce sync.Once
	file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescData = file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDesc
)

func file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescGZIP() []byte {
	file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescOnce.Do(func() {
		file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescData)
	})
	return file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescData
}

var file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_aserto_tenant_policy_builder_v1_policy_builder_proto_goTypes = []interface{}{
	(*CreatePolicyBuilderRequest)(nil),  // 0: aserto.tenant.policy_builder.v1.CreatePolicyBuilderRequest
	(*CreatePolicyBuilderResponse)(nil), // 1: aserto.tenant.policy_builder.v1.CreatePolicyBuilderResponse
	(*ListPolicyBuildersRequest)(nil),   // 2: aserto.tenant.policy_builder.v1.ListPolicyBuildersRequest
	(*ListPolicyBuildersResponse)(nil),  // 3: aserto.tenant.policy_builder.v1.ListPolicyBuildersResponse
	(*DeletePolicyBuilderRequest)(nil),  // 4: aserto.tenant.policy_builder.v1.DeletePolicyBuilderRequest
	(*DeletePolicyBuilderResponse)(nil), // 5: aserto.tenant.policy_builder.v1.DeletePolicyBuilderResponse
	(*v1.PolicyBuilder)(nil),            // 6: aserto.api.v1.PolicyBuilder
	(*emptypb.Empty)(nil),               // 7: google.protobuf.Empty
}
var file_aserto_tenant_policy_builder_v1_policy_builder_proto_depIdxs = []int32{
	6, // 0: aserto.tenant.policy_builder.v1.CreatePolicyBuilderRequest.policy_builder:type_name -> aserto.api.v1.PolicyBuilder
	6, // 1: aserto.tenant.policy_builder.v1.ListPolicyBuildersResponse.results:type_name -> aserto.api.v1.PolicyBuilder
	7, // 2: aserto.tenant.policy_builder.v1.DeletePolicyBuilderResponse.result:type_name -> google.protobuf.Empty
	2, // 3: aserto.tenant.policy_builder.v1.PolicyBuilder.ListPolicyBuilders:input_type -> aserto.tenant.policy_builder.v1.ListPolicyBuildersRequest
	0, // 4: aserto.tenant.policy_builder.v1.PolicyBuilder.CreatePolicyBuilder:input_type -> aserto.tenant.policy_builder.v1.CreatePolicyBuilderRequest
	4, // 5: aserto.tenant.policy_builder.v1.PolicyBuilder.DeletePolicyBuilder:input_type -> aserto.tenant.policy_builder.v1.DeletePolicyBuilderRequest
	3, // 6: aserto.tenant.policy_builder.v1.PolicyBuilder.ListPolicyBuilders:output_type -> aserto.tenant.policy_builder.v1.ListPolicyBuildersResponse
	1, // 7: aserto.tenant.policy_builder.v1.PolicyBuilder.CreatePolicyBuilder:output_type -> aserto.tenant.policy_builder.v1.CreatePolicyBuilderResponse
	5, // 8: aserto.tenant.policy_builder.v1.PolicyBuilder.DeletePolicyBuilder:output_type -> aserto.tenant.policy_builder.v1.DeletePolicyBuilderResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_aserto_tenant_policy_builder_v1_policy_builder_proto_init() }
func file_aserto_tenant_policy_builder_v1_policy_builder_proto_init() {
	if File_aserto_tenant_policy_builder_v1_policy_builder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePolicyBuilderRequest); i {
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
		file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePolicyBuilderResponse); i {
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
		file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPolicyBuildersRequest); i {
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
		file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPolicyBuildersResponse); i {
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
		file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePolicyBuilderRequest); i {
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
		file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePolicyBuilderResponse); i {
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
			RawDescriptor: file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_tenant_policy_builder_v1_policy_builder_proto_goTypes,
		DependencyIndexes: file_aserto_tenant_policy_builder_v1_policy_builder_proto_depIdxs,
		MessageInfos:      file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes,
	}.Build()
	File_aserto_tenant_policy_builder_v1_policy_builder_proto = out.File
	file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDesc = nil
	file_aserto_tenant_policy_builder_v1_policy_builder_proto_goTypes = nil
	file_aserto_tenant_policy_builder_v1_policy_builder_proto_depIdxs = nil
}
