// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: aserto/management/v2/control_plane.proto

package management

import (
	v2 "github.com/aserto-dev/go-grpc/aserto/api/v2"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListInstanceRegistrationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListInstanceRegistrationsRequest) Reset() {
	*x = ListInstanceRegistrationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_management_v2_control_plane_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstanceRegistrationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstanceRegistrationsRequest) ProtoMessage() {}

func (x *ListInstanceRegistrationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_management_v2_control_plane_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstanceRegistrationsRequest.ProtoReflect.Descriptor instead.
func (*ListInstanceRegistrationsRequest) Descriptor() ([]byte, []int) {
	return file_aserto_management_v2_control_plane_proto_rawDescGZIP(), []int{0}
}

type ListInstanceRegistrationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*v2.InstanceRegistration `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *ListInstanceRegistrationsResponse) Reset() {
	*x = ListInstanceRegistrationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_management_v2_control_plane_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstanceRegistrationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstanceRegistrationsResponse) ProtoMessage() {}

func (x *ListInstanceRegistrationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_management_v2_control_plane_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstanceRegistrationsResponse.ProtoReflect.Descriptor instead.
func (*ListInstanceRegistrationsResponse) Descriptor() ([]byte, []int) {
	return file_aserto_management_v2_control_plane_proto_rawDescGZIP(), []int{1}
}

func (x *ListInstanceRegistrationsResponse) GetResult() []*v2.InstanceRegistration {
	if x != nil {
		return x.Result
	}
	return nil
}

type ExecCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Command *v2.Command `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *ExecCommandRequest) Reset() {
	*x = ExecCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_management_v2_control_plane_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecCommandRequest) ProtoMessage() {}

func (x *ExecCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_management_v2_control_plane_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecCommandRequest.ProtoReflect.Descriptor instead.
func (*ExecCommandRequest) Descriptor() ([]byte, []int) {
	return file_aserto_management_v2_control_plane_proto_rawDescGZIP(), []int{2}
}

func (x *ExecCommandRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExecCommandRequest) GetCommand() *v2.Command {
	if x != nil {
		return x.Command
	}
	return nil
}

type ExecCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *emptypb.Empty `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ExecCommandResponse) Reset() {
	*x = ExecCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_management_v2_control_plane_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecCommandResponse) ProtoMessage() {}

func (x *ExecCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_management_v2_control_plane_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecCommandResponse.ProtoReflect.Descriptor instead.
func (*ExecCommandResponse) Descriptor() ([]byte, []int) {
	return file_aserto_management_v2_control_plane_proto_rawDescGZIP(), []int{3}
}

func (x *ExecCommandResponse) GetResult() *emptypb.Empty {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_aserto_management_v2_control_plane_proto protoreflect.FileDescriptor

var file_aserto_management_v2_control_plane_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x20, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x60, 0x0a, 0x21, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x56, 0x0a, 0x12, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x45, 0x0a, 0x13, 0x45, 0x78, 0x65,
	0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2e, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0xd0, 0x05, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e,
	0x65, 0x12, 0xfe, 0x02, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x36, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0xef, 0x01, 0x92, 0x41, 0xac, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x12, 0x1b, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x20, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x2f, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x61, 0x20, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x20, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x2a, 0x34, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03,
	0x4a, 0x57, 0x54, 0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x12, 0x37, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0xbe, 0x02, 0x0a, 0x0b, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x28, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd9, 0x01, 0x92, 0x41, 0xa5, 0x01, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x12, 0x29, 0x72,
	0x75, 0x6e, 0x73, 0x20, 0x61, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x20, 0x6f, 0x6e,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x20, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x1a, 0x29, 0x52, 0x75, 0x6e, 0x73, 0x20, 0x61,
	0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x2a, 0x25, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x65, 0x78,
	0x65, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x62, 0x17, 0x0a, 0x07, 0x0a, 0x03,
	0x4a, 0x57, 0x54, 0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x22, 0x25, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x78, 0x65, 0x63,
	0x3a, 0x01, 0x2a, 0x42, 0xdc, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x3b, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0xaa, 0x02, 0x14, 0x41, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x32, 0x92, 0x41, 0x82, 0x01,
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
	file_aserto_management_v2_control_plane_proto_rawDescOnce sync.Once
	file_aserto_management_v2_control_plane_proto_rawDescData = file_aserto_management_v2_control_plane_proto_rawDesc
)

func file_aserto_management_v2_control_plane_proto_rawDescGZIP() []byte {
	file_aserto_management_v2_control_plane_proto_rawDescOnce.Do(func() {
		file_aserto_management_v2_control_plane_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_management_v2_control_plane_proto_rawDescData)
	})
	return file_aserto_management_v2_control_plane_proto_rawDescData
}

var file_aserto_management_v2_control_plane_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_aserto_management_v2_control_plane_proto_goTypes = []interface{}{
	(*ListInstanceRegistrationsRequest)(nil),  // 0: aserto.management.v2.ListInstanceRegistrationsRequest
	(*ListInstanceRegistrationsResponse)(nil), // 1: aserto.management.v2.ListInstanceRegistrationsResponse
	(*ExecCommandRequest)(nil),                // 2: aserto.management.v2.ExecCommandRequest
	(*ExecCommandResponse)(nil),               // 3: aserto.management.v2.ExecCommandResponse
	(*v2.InstanceRegistration)(nil),           // 4: aserto.api.v2.InstanceRegistration
	(*v2.Command)(nil),                        // 5: aserto.api.v2.Command
	(*emptypb.Empty)(nil),                     // 6: google.protobuf.Empty
}
var file_aserto_management_v2_control_plane_proto_depIdxs = []int32{
	4, // 0: aserto.management.v2.ListInstanceRegistrationsResponse.result:type_name -> aserto.api.v2.InstanceRegistration
	5, // 1: aserto.management.v2.ExecCommandRequest.command:type_name -> aserto.api.v2.Command
	6, // 2: aserto.management.v2.ExecCommandResponse.result:type_name -> google.protobuf.Empty
	0, // 3: aserto.management.v2.ControlPlane.ListInstanceRegistrations:input_type -> aserto.management.v2.ListInstanceRegistrationsRequest
	2, // 4: aserto.management.v2.ControlPlane.ExecCommand:input_type -> aserto.management.v2.ExecCommandRequest
	1, // 5: aserto.management.v2.ControlPlane.ListInstanceRegistrations:output_type -> aserto.management.v2.ListInstanceRegistrationsResponse
	3, // 6: aserto.management.v2.ControlPlane.ExecCommand:output_type -> aserto.management.v2.ExecCommandResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_aserto_management_v2_control_plane_proto_init() }
func file_aserto_management_v2_control_plane_proto_init() {
	if File_aserto_management_v2_control_plane_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_management_v2_control_plane_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInstanceRegistrationsRequest); i {
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
		file_aserto_management_v2_control_plane_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInstanceRegistrationsResponse); i {
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
		file_aserto_management_v2_control_plane_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecCommandRequest); i {
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
		file_aserto_management_v2_control_plane_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecCommandResponse); i {
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
			RawDescriptor: file_aserto_management_v2_control_plane_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_management_v2_control_plane_proto_goTypes,
		DependencyIndexes: file_aserto_management_v2_control_plane_proto_depIdxs,
		MessageInfos:      file_aserto_management_v2_control_plane_proto_msgTypes,
	}.Build()
	File_aserto_management_v2_control_plane_proto = out.File
	file_aserto_management_v2_control_plane_proto_rawDesc = nil
	file_aserto_management_v2_control_plane_proto_goTypes = nil
	file_aserto_management_v2_control_plane_proto_depIdxs = nil
}
