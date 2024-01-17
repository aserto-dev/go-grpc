// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: aserto/tenant/provider/v1/provider.proto

package provider

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

type ListProvidersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind v1.ProviderKind `protobuf:"varint,1,opt,name=kind,proto3,enum=aserto.api.v1.ProviderKind" json:"kind,omitempty"`
}

func (x *ListProvidersRequest) Reset() {
	*x = ListProvidersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_provider_v1_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProvidersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProvidersRequest) ProtoMessage() {}

func (x *ListProvidersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_provider_v1_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProvidersRequest.ProtoReflect.Descriptor instead.
func (*ListProvidersRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_provider_v1_provider_proto_rawDescGZIP(), []int{0}
}

func (x *ListProvidersRequest) GetKind() v1.ProviderKind {
	if x != nil {
		return x.Kind
	}
	return v1.ProviderKind(0)
}

type ListProvidersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*v1.Provider `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListProvidersResponse) Reset() {
	*x = ListProvidersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_provider_v1_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProvidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProvidersResponse) ProtoMessage() {}

func (x *ListProvidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_provider_v1_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProvidersResponse.ProtoReflect.Descriptor instead.
func (*ListProvidersResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_provider_v1_provider_proto_rawDescGZIP(), []int{1}
}

func (x *ListProvidersResponse) GetResults() []*v1.Provider {
	if x != nil {
		return x.Results
	}
	return nil
}

type GetProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind v1.ProviderKind `protobuf:"varint,1,opt,name=kind,proto3,enum=aserto.api.v1.ProviderKind" json:"kind,omitempty"`
	Id   string          `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetProviderRequest) Reset() {
	*x = GetProviderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_provider_v1_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProviderRequest) ProtoMessage() {}

func (x *GetProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_provider_v1_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProviderRequest.ProtoReflect.Descriptor instead.
func (*GetProviderRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_provider_v1_provider_proto_rawDescGZIP(), []int{2}
}

func (x *GetProviderRequest) GetKind() v1.ProviderKind {
	if x != nil {
		return x.Kind
	}
	return v1.ProviderKind(0)
}

func (x *GetProviderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProviderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*v1.Provider `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *GetProviderResponse) Reset() {
	*x = GetProviderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_provider_v1_provider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProviderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProviderResponse) ProtoMessage() {}

func (x *GetProviderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_provider_v1_provider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProviderResponse.ProtoReflect.Descriptor instead.
func (*GetProviderResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_provider_v1_provider_proto_rawDescGZIP(), []int{3}
}

func (x *GetProviderResponse) GetResults() []*v1.Provider {
	if x != nil {
		return x.Results
	}
	return nil
}

type ListProviderKindsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProviderKindsRequest) Reset() {
	*x = ListProviderKindsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_provider_v1_provider_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProviderKindsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProviderKindsRequest) ProtoMessage() {}

func (x *ListProviderKindsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_provider_v1_provider_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProviderKindsRequest.ProtoReflect.Descriptor instead.
func (*ListProviderKindsRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_provider_v1_provider_proto_rawDescGZIP(), []int{4}
}

type ListProviderKindsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []string `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListProviderKindsResponse) Reset() {
	*x = ListProviderKindsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_tenant_provider_v1_provider_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProviderKindsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProviderKindsResponse) ProtoMessage() {}

func (x *ListProviderKindsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_provider_v1_provider_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProviderKindsResponse.ProtoReflect.Descriptor instead.
func (*ListProviderKindsResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_provider_v1_provider_proto_rawDescGZIP(), []int{5}
}

func (x *ListProviderKindsResponse) GetResults() []string {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_aserto_tenant_provider_v1_provider_proto protoreflect.FileDescriptor

var file_aserto_tenant_provider_v1_provider_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x47, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x4a, 0x0a, 0x15, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x5b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x90, 0xb5, 0x18, 0x06, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x48, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x1a, 0x0a, 0x18,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32,
	0x9c, 0x06, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0xf5, 0x01, 0x0a,
	0x0d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2f,
	0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x80, 0x01, 0x92, 0x41, 0x64, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x1a, 0x24, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x2a, 0x17, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73,
	0x62, 0x09, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x73, 0x12, 0xf6, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x2d, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x87, 0x01, 0x92, 0x41, 0x66, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x12, 0x0c, 0x47, 0x65, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x1a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6d,
	0x65, 0x74, 0x61, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x69, 0x6e, 0x67, 0x20, 0x6f,
	0x66, 0x20, 0x61, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x2a, 0x15, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x62, 0x09, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x9e, 0x02,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4b, 0x69,
	0x6e, 0x64, 0x73, 0x12, 0x33, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x4b, 0x69, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9d,
	0x01, 0x92, 0x41, 0x7b, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x20, 0x6b, 0x69,
	0x6e, 0x64, 0x73, 0x1a, 0x31, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x20, 0x6b, 0x69, 0x6e, 0x64, 0x73, 0x20, 0x28, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x29, 0x2e, 0x2a, 0x1c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x6b,
	0x69, 0x6e, 0x64, 0x73, 0x62, 0x09, 0x0a, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x00, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x6b, 0x69, 0x6e, 0x64, 0x73, 0x42, 0x95,
	0x01, 0x92, 0x41, 0x50, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x1c, 0x0a, 0x1a, 0x0a,
	0x03, 0x4a, 0x57, 0x54, 0x12, 0x13, 0x08, 0x02, 0x1a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x09, 0x0a, 0x07, 0x0a, 0x03, 0x4a,
	0x57, 0x54, 0x12, 0x00, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_tenant_provider_v1_provider_proto_rawDescOnce sync.Once
	file_aserto_tenant_provider_v1_provider_proto_rawDescData = file_aserto_tenant_provider_v1_provider_proto_rawDesc
)

func file_aserto_tenant_provider_v1_provider_proto_rawDescGZIP() []byte {
	file_aserto_tenant_provider_v1_provider_proto_rawDescOnce.Do(func() {
		file_aserto_tenant_provider_v1_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_tenant_provider_v1_provider_proto_rawDescData)
	})
	return file_aserto_tenant_provider_v1_provider_proto_rawDescData
}

var file_aserto_tenant_provider_v1_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_aserto_tenant_provider_v1_provider_proto_goTypes = []interface{}{
	(*ListProvidersRequest)(nil),      // 0: aserto.tenant.provider.v1.ListProvidersRequest
	(*ListProvidersResponse)(nil),     // 1: aserto.tenant.provider.v1.ListProvidersResponse
	(*GetProviderRequest)(nil),        // 2: aserto.tenant.provider.v1.GetProviderRequest
	(*GetProviderResponse)(nil),       // 3: aserto.tenant.provider.v1.GetProviderResponse
	(*ListProviderKindsRequest)(nil),  // 4: aserto.tenant.provider.v1.ListProviderKindsRequest
	(*ListProviderKindsResponse)(nil), // 5: aserto.tenant.provider.v1.ListProviderKindsResponse
	(v1.ProviderKind)(0),              // 6: aserto.api.v1.ProviderKind
	(*v1.Provider)(nil),               // 7: aserto.api.v1.Provider
}
var file_aserto_tenant_provider_v1_provider_proto_depIdxs = []int32{
	6, // 0: aserto.tenant.provider.v1.ListProvidersRequest.kind:type_name -> aserto.api.v1.ProviderKind
	7, // 1: aserto.tenant.provider.v1.ListProvidersResponse.results:type_name -> aserto.api.v1.Provider
	6, // 2: aserto.tenant.provider.v1.GetProviderRequest.kind:type_name -> aserto.api.v1.ProviderKind
	7, // 3: aserto.tenant.provider.v1.GetProviderResponse.results:type_name -> aserto.api.v1.Provider
	0, // 4: aserto.tenant.provider.v1.Provider.ListProviders:input_type -> aserto.tenant.provider.v1.ListProvidersRequest
	2, // 5: aserto.tenant.provider.v1.Provider.GetProvider:input_type -> aserto.tenant.provider.v1.GetProviderRequest
	4, // 6: aserto.tenant.provider.v1.Provider.ListProviderKinds:input_type -> aserto.tenant.provider.v1.ListProviderKindsRequest
	1, // 7: aserto.tenant.provider.v1.Provider.ListProviders:output_type -> aserto.tenant.provider.v1.ListProvidersResponse
	3, // 8: aserto.tenant.provider.v1.Provider.GetProvider:output_type -> aserto.tenant.provider.v1.GetProviderResponse
	5, // 9: aserto.tenant.provider.v1.Provider.ListProviderKinds:output_type -> aserto.tenant.provider.v1.ListProviderKindsResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_aserto_tenant_provider_v1_provider_proto_init() }
func file_aserto_tenant_provider_v1_provider_proto_init() {
	if File_aserto_tenant_provider_v1_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_tenant_provider_v1_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProvidersRequest); i {
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
		file_aserto_tenant_provider_v1_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProvidersResponse); i {
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
		file_aserto_tenant_provider_v1_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProviderRequest); i {
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
		file_aserto_tenant_provider_v1_provider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProviderResponse); i {
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
		file_aserto_tenant_provider_v1_provider_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProviderKindsRequest); i {
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
		file_aserto_tenant_provider_v1_provider_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProviderKindsResponse); i {
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
			RawDescriptor: file_aserto_tenant_provider_v1_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_tenant_provider_v1_provider_proto_goTypes,
		DependencyIndexes: file_aserto_tenant_provider_v1_provider_proto_depIdxs,
		MessageInfos:      file_aserto_tenant_provider_v1_provider_proto_msgTypes,
	}.Build()
	File_aserto_tenant_provider_v1_provider_proto = out.File
	file_aserto_tenant_provider_v1_provider_proto_rawDesc = nil
	file_aserto_tenant_provider_v1_provider_proto_goTypes = nil
	file_aserto_tenant_provider_v1_provider_proto_depIdxs = nil
}
