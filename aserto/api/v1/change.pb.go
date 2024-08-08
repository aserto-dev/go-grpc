// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: aserto/api/v1/change.proto

package api

import (
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

type ChangeType int32

const (
	ChangeType_CHANGE_TYPE_UNKNOWN ChangeType = 0
	ChangeType_CHANGE_TYPE_CREATED ChangeType = 1
	ChangeType_CHANGE_TYPE_UPDATED ChangeType = 2
	ChangeType_CHANGE_TYPE_DELETED ChangeType = 3
)

// Enum value maps for ChangeType.
var (
	ChangeType_name = map[int32]string{
		0: "CHANGE_TYPE_UNKNOWN",
		1: "CHANGE_TYPE_CREATED",
		2: "CHANGE_TYPE_UPDATED",
		3: "CHANGE_TYPE_DELETED",
	}
	ChangeType_value = map[string]int32{
		"CHANGE_TYPE_UNKNOWN": 0,
		"CHANGE_TYPE_CREATED": 1,
		"CHANGE_TYPE_UPDATED": 2,
		"CHANGE_TYPE_DELETED": 3,
	}
)

func (x ChangeType) Enum() *ChangeType {
	p := new(ChangeType)
	*p = x
	return p
}

func (x ChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_api_v1_change_proto_enumTypes[0].Descriptor()
}

func (ChangeType) Type() protoreflect.EnumType {
	return &file_aserto_api_v1_change_proto_enumTypes[0]
}

func (x ChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeType.Descriptor instead.
func (ChangeType) EnumDescriptor() ([]byte, []int) {
	return file_aserto_api_v1_change_proto_rawDescGZIP(), []int{0}
}

var File_aserto_api_v1_change_proto protoreflect.FileDescriptor

var file_aserto_api_v1_change_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2a, 0x70, 0x0a, 0x0a, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x43,
	0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0x41, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69,
	0xaa, 0x02, 0x0d, 0x41, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_api_v1_change_proto_rawDescOnce sync.Once
	file_aserto_api_v1_change_proto_rawDescData = file_aserto_api_v1_change_proto_rawDesc
)

func file_aserto_api_v1_change_proto_rawDescGZIP() []byte {
	file_aserto_api_v1_change_proto_rawDescOnce.Do(func() {
		file_aserto_api_v1_change_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_api_v1_change_proto_rawDescData)
	})
	return file_aserto_api_v1_change_proto_rawDescData
}

var file_aserto_api_v1_change_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_api_v1_change_proto_goTypes = []any{
	(ChangeType)(0), // 0: aserto.api.v1.ChangeType
}
var file_aserto_api_v1_change_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aserto_api_v1_change_proto_init() }
func file_aserto_api_v1_change_proto_init() {
	if File_aserto_api_v1_change_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aserto_api_v1_change_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_api_v1_change_proto_goTypes,
		DependencyIndexes: file_aserto_api_v1_change_proto_depIdxs,
		EnumInfos:         file_aserto_api_v1_change_proto_enumTypes,
	}.Build()
	File_aserto_api_v1_change_proto = out.File
	file_aserto_api_v1_change_proto_rawDesc = nil
	file_aserto_api_v1_change_proto_goTypes = nil
	file_aserto_api_v1_change_proto_depIdxs = nil
}
