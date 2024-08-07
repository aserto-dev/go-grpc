// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: aserto/api/v1/policy_builder.proto

package api

import (
	_ "github.com/aserto-dev/go-grpc/aserto/options/v1"
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

type PolicyBuilder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Configuration for source code
	SccConnectionId      string `protobuf:"bytes,2,opt,name=scc_connection_id,json=sccConnectionId,proto3" json:"scc_connection_id,omitempty"`
	SccRepo              string `protobuf:"bytes,3,opt,name=scc_repo,json=sccRepo,proto3" json:"scc_repo,omitempty"`
	SccOrg               string `protobuf:"bytes,4,opt,name=scc_org,json=sccOrg,proto3" json:"scc_org,omitempty"`
	RegistryConnectionId string `protobuf:"bytes,5,opt,name=registry_connection_id,json=registryConnectionId,proto3" json:"registry_connection_id,omitempty"`
	RegistryRepo         string `protobuf:"bytes,6,opt,name=registry_repo,json=registryRepo,proto3" json:"registry_repo,omitempty"`
	RegistryOrg          string `protobuf:"bytes,7,opt,name=registry_org,json=registryOrg,proto3" json:"registry_org,omitempty"`
}

func (x *PolicyBuilder) Reset() {
	*x = PolicyBuilder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v1_policy_builder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyBuilder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyBuilder) ProtoMessage() {}

func (x *PolicyBuilder) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v1_policy_builder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyBuilder.ProtoReflect.Descriptor instead.
func (*PolicyBuilder) Descriptor() ([]byte, []int) {
	return file_aserto_api_v1_policy_builder_proto_rawDescGZIP(), []int{0}
}

func (x *PolicyBuilder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PolicyBuilder) GetSccConnectionId() string {
	if x != nil {
		return x.SccConnectionId
	}
	return ""
}

func (x *PolicyBuilder) GetSccRepo() string {
	if x != nil {
		return x.SccRepo
	}
	return ""
}

func (x *PolicyBuilder) GetSccOrg() string {
	if x != nil {
		return x.SccOrg
	}
	return ""
}

func (x *PolicyBuilder) GetRegistryConnectionId() string {
	if x != nil {
		return x.RegistryConnectionId
	}
	return ""
}

func (x *PolicyBuilder) GetRegistryRepo() string {
	if x != nil {
		return x.RegistryRepo
	}
	return ""
}

func (x *PolicyBuilder) GetRegistryOrg() string {
	if x != nil {
		return x.RegistryOrg
	}
	return ""
}

var File_aserto_api_v1_policy_builder_proto protoreflect.FileDescriptor

var file_aserto_api_v1_policy_builder_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8f, 0x02, 0x0a, 0x0d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04,
	0x90, 0xb5, 0x18, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x11, 0x73, 0x63, 0x63, 0x5f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0x90, 0xb5, 0x18, 0x07, 0x52, 0x0f, 0x73, 0x63, 0x63, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63,
	0x63, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x63,
	0x63, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x63, 0x63, 0x5f, 0x6f, 0x72, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x63, 0x4f, 0x72, 0x67, 0x12, 0x3a,
	0x0a, 0x16, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04,
	0x90, 0xb5, 0x18, 0x07, 0x52, 0x14, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x6f, 0x72, 0x67, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x4f,
	0x72, 0x67, 0x42, 0x41, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x3b, 0x61, 0x70, 0x69, 0xaa, 0x02, 0x0d, 0x41, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x41,
	0x50, 0x49, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_api_v1_policy_builder_proto_rawDescOnce sync.Once
	file_aserto_api_v1_policy_builder_proto_rawDescData = file_aserto_api_v1_policy_builder_proto_rawDesc
)

func file_aserto_api_v1_policy_builder_proto_rawDescGZIP() []byte {
	file_aserto_api_v1_policy_builder_proto_rawDescOnce.Do(func() {
		file_aserto_api_v1_policy_builder_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_api_v1_policy_builder_proto_rawDescData)
	})
	return file_aserto_api_v1_policy_builder_proto_rawDescData
}

var file_aserto_api_v1_policy_builder_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_aserto_api_v1_policy_builder_proto_goTypes = []any{
	(*PolicyBuilder)(nil), // 0: aserto.api.v1.PolicyBuilder
}
var file_aserto_api_v1_policy_builder_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aserto_api_v1_policy_builder_proto_init() }
func file_aserto_api_v1_policy_builder_proto_init() {
	if File_aserto_api_v1_policy_builder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_api_v1_policy_builder_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PolicyBuilder); i {
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
			RawDescriptor: file_aserto_api_v1_policy_builder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_api_v1_policy_builder_proto_goTypes,
		DependencyIndexes: file_aserto_api_v1_policy_builder_proto_depIdxs,
		MessageInfos:      file_aserto_api_v1_policy_builder_proto_msgTypes,
	}.Build()
	File_aserto_api_v1_policy_builder_proto = out.File
	file_aserto_api_v1_policy_builder_proto_rawDesc = nil
	file_aserto_api_v1_policy_builder_proto_goTypes = nil
	file_aserto_api_v1_policy_builder_proto_depIdxs = nil
}
