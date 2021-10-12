// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.17.3
// source: aserto/registry/v1/registry.proto

package registry

import (
	v1 "github.com/aserto-dev/go-grpc/aserto/api/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type ListImagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListImagesRequest) Reset() {
	*x = ListImagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_registry_v1_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImagesRequest) ProtoMessage() {}

func (x *ListImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_registry_v1_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListImagesRequest.ProtoReflect.Descriptor instead.
func (*ListImagesRequest) Descriptor() ([]byte, []int) {
	return file_aserto_registry_v1_registry_proto_rawDescGZIP(), []int{0}
}

type ListImagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images []*v1.PolicyImage `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
}

func (x *ListImagesResponse) Reset() {
	*x = ListImagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_registry_v1_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImagesResponse) ProtoMessage() {}

func (x *ListImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_registry_v1_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListImagesResponse.ProtoReflect.Descriptor instead.
func (*ListImagesResponse) Descriptor() ([]byte, []int) {
	return file_aserto_registry_v1_registry_proto_rawDescGZIP(), []int{1}
}

func (x *ListImagesResponse) GetImages() []*v1.PolicyImage {
	if x != nil {
		return x.Images
	}
	return nil
}

type RemoveImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image        string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Tag          string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	Organization string `protobuf:"bytes,3,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (x *RemoveImageRequest) Reset() {
	*x = RemoveImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_registry_v1_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveImageRequest) ProtoMessage() {}

func (x *RemoveImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_registry_v1_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveImageRequest.ProtoReflect.Descriptor instead.
func (*RemoveImageRequest) Descriptor() ([]byte, []int) {
	return file_aserto_registry_v1_registry_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveImageRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *RemoveImageRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *RemoveImageRequest) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

type RemoveImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveImageResponse) Reset() {
	*x = RemoveImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_registry_v1_registry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveImageResponse) ProtoMessage() {}

func (x *RemoveImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_registry_v1_registry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveImageResponse.ProtoReflect.Descriptor instead.
func (*RemoveImageResponse) Descriptor() ([]byte, []int) {
	return file_aserto_registry_v1_registry_proto_rawDescGZIP(), []int{3}
}

type SetImageVisibilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Public       bool   `protobuf:"varint,1,opt,name=public,proto3" json:"public,omitempty"`
	Image        string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Organization string `protobuf:"bytes,3,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (x *SetImageVisibilityRequest) Reset() {
	*x = SetImageVisibilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_registry_v1_registry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetImageVisibilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetImageVisibilityRequest) ProtoMessage() {}

func (x *SetImageVisibilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_registry_v1_registry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetImageVisibilityRequest.ProtoReflect.Descriptor instead.
func (*SetImageVisibilityRequest) Descriptor() ([]byte, []int) {
	return file_aserto_registry_v1_registry_proto_rawDescGZIP(), []int{4}
}

func (x *SetImageVisibilityRequest) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

func (x *SetImageVisibilityRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *SetImageVisibilityRequest) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

type SetImageVisibilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Public bool `protobuf:"varint,1,opt,name=public,proto3" json:"public,omitempty"`
}

func (x *SetImageVisibilityResponse) Reset() {
	*x = SetImageVisibilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_registry_v1_registry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetImageVisibilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetImageVisibilityResponse) ProtoMessage() {}

func (x *SetImageVisibilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_registry_v1_registry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetImageVisibilityResponse.ProtoReflect.Descriptor instead.
func (*SetImageVisibilityResponse) Descriptor() ([]byte, []int) {
	return file_aserto_registry_v1_registry_proto_rawDescGZIP(), []int{5}
}

func (x *SetImageVisibilityResponse) GetPublic() bool {
	if x != nil {
		return x.Public
	}
	return false
}

var File_aserto_registry_v1_registry_proto protoreflect.FileDescriptor

var file_aserto_registry_v1_registry_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x32, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x22, 0x60, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74,
	0x61, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6d, 0x0a,
	0x19, 0x53, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x34, 0x0a, 0x1a,
	0x53, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x32, 0xfa, 0x06, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12,
	0xfc, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x25,
	0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9e, 0x01,
	0x92, 0x41, 0x7c, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x20, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x31, 0x4c, 0x69, 0x73, 0x74,
	0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x20, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x20, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x2a, 0x14, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x62, 0x1a, 0x0a, 0x0a, 0x0a, 0x06, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12,
	0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0xfe,
	0x01, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x26,
	0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x9d, 0x01, 0x92, 0x41, 0x64, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12,
	0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x20, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x73, 0x20, 0x61, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x20,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x2a, 0x15, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x62, 0x1a, 0x0a,
	0x0a, 0x0a, 0x06, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x2a,
	0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x7b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x7d, 0x12,
	0xed, 0x02, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x56, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x2d, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xf7, 0x01, 0x92, 0x41, 0xb2, 0x01, 0x0a, 0x08, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x21, 0x4d, 0x61, 0x6b, 0x65, 0x73, 0x20, 0x61, 0x6e,
	0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x20, 0x6f, 0x72,
	0x20, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2e, 0x1a, 0x48, 0x49, 0x66, 0x20, 0x61, 0x6e,
	0x20, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x20, 0x69, 0x73, 0x20, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2c, 0x20, 0x69, 0x74, 0x20, 0x6d, 0x65, 0x61, 0x6e, 0x73, 0x20, 0x69, 0x74, 0x27, 0x73, 0x20,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2d, 0x72, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x20,
	0x41, 0x6e, 0x79, 0x6f, 0x6e, 0x65, 0x20, 0x63, 0x61, 0x6e, 0x20, 0x72, 0x65, 0x61, 0x64, 0x20,
	0x69, 0x74, 0x2e, 0x2a, 0x1d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x73, 0x65,
	0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x62, 0x1a, 0x0a, 0x0a, 0x0a, 0x06, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x00,
	0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x3b, 0x22, 0x39, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x7b, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x7d, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42,
	0xdc, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0xaa, 0x02, 0x12,
	0x41, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x56, 0x31, 0x92, 0x41, 0x88, 0x01, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x43, 0x0a,
	0x1d, 0x0a, 0x06, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x13, 0x08, 0x02, 0x1a, 0x0d, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x0a, 0x22,
	0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x08, 0x02, 0x1a, 0x10,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2d, 0x69, 0x64,
	0x20, 0x02, 0x62, 0x1a, 0x0a, 0x0a, 0x0a, 0x06, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x00,
	0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_registry_v1_registry_proto_rawDescOnce sync.Once
	file_aserto_registry_v1_registry_proto_rawDescData = file_aserto_registry_v1_registry_proto_rawDesc
)

func file_aserto_registry_v1_registry_proto_rawDescGZIP() []byte {
	file_aserto_registry_v1_registry_proto_rawDescOnce.Do(func() {
		file_aserto_registry_v1_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_registry_v1_registry_proto_rawDescData)
	})
	return file_aserto_registry_v1_registry_proto_rawDescData
}

var file_aserto_registry_v1_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_aserto_registry_v1_registry_proto_goTypes = []interface{}{
	(*ListImagesRequest)(nil),          // 0: aserto.registry.v1.ListImagesRequest
	(*ListImagesResponse)(nil),         // 1: aserto.registry.v1.ListImagesResponse
	(*RemoveImageRequest)(nil),         // 2: aserto.registry.v1.RemoveImageRequest
	(*RemoveImageResponse)(nil),        // 3: aserto.registry.v1.RemoveImageResponse
	(*SetImageVisibilityRequest)(nil),  // 4: aserto.registry.v1.SetImageVisibilityRequest
	(*SetImageVisibilityResponse)(nil), // 5: aserto.registry.v1.SetImageVisibilityResponse
	(*v1.PolicyImage)(nil),             // 6: aserto.api.v1.PolicyImage
}
var file_aserto_registry_v1_registry_proto_depIdxs = []int32{
	6, // 0: aserto.registry.v1.ListImagesResponse.images:type_name -> aserto.api.v1.PolicyImage
	0, // 1: aserto.registry.v1.Registry.ListImages:input_type -> aserto.registry.v1.ListImagesRequest
	2, // 2: aserto.registry.v1.Registry.RemoveImage:input_type -> aserto.registry.v1.RemoveImageRequest
	4, // 3: aserto.registry.v1.Registry.SetImageVisibility:input_type -> aserto.registry.v1.SetImageVisibilityRequest
	1, // 4: aserto.registry.v1.Registry.ListImages:output_type -> aserto.registry.v1.ListImagesResponse
	3, // 5: aserto.registry.v1.Registry.RemoveImage:output_type -> aserto.registry.v1.RemoveImageResponse
	5, // 6: aserto.registry.v1.Registry.SetImageVisibility:output_type -> aserto.registry.v1.SetImageVisibilityResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aserto_registry_v1_registry_proto_init() }
func file_aserto_registry_v1_registry_proto_init() {
	if File_aserto_registry_v1_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_registry_v1_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListImagesRequest); i {
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
		file_aserto_registry_v1_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListImagesResponse); i {
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
		file_aserto_registry_v1_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveImageRequest); i {
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
		file_aserto_registry_v1_registry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveImageResponse); i {
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
		file_aserto_registry_v1_registry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetImageVisibilityRequest); i {
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
		file_aserto_registry_v1_registry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetImageVisibilityResponse); i {
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
			RawDescriptor: file_aserto_registry_v1_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_registry_v1_registry_proto_goTypes,
		DependencyIndexes: file_aserto_registry_v1_registry_proto_depIdxs,
		MessageInfos:      file_aserto_registry_v1_registry_proto_msgTypes,
	}.Build()
	File_aserto_registry_v1_registry_proto = out.File
	file_aserto_registry_v1_registry_proto_rawDesc = nil
	file_aserto_registry_v1_registry_proto_goTypes = nil
	file_aserto_registry_v1_registry_proto_depIdxs = nil
}
