// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: aserto/api/v1/name_availability.proto

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

type NameAvailability int32

const (
	NameAvailability_NAME_AVAILABILITY_UNKNOWN     NameAvailability = 0
	NameAvailability_NAME_AVAILABILITY_AVAILABLE   NameAvailability = 1
	NameAvailability_NAME_AVAILABILITY_UNAVAILABLE NameAvailability = 2
	NameAvailability_NAME_AVAILABILITY_INVALID     NameAvailability = 3
	NameAvailability_NAME_AVAILABILITY_PROFANE     NameAvailability = 4
	NameAvailability_NAME_AVAILABILITY_RESERVED    NameAvailability = 5
)

// Enum value maps for NameAvailability.
var (
	NameAvailability_name = map[int32]string{
		0: "NAME_AVAILABILITY_UNKNOWN",
		1: "NAME_AVAILABILITY_AVAILABLE",
		2: "NAME_AVAILABILITY_UNAVAILABLE",
		3: "NAME_AVAILABILITY_INVALID",
		4: "NAME_AVAILABILITY_PROFANE",
		5: "NAME_AVAILABILITY_RESERVED",
	}
	NameAvailability_value = map[string]int32{
		"NAME_AVAILABILITY_UNKNOWN":     0,
		"NAME_AVAILABILITY_AVAILABLE":   1,
		"NAME_AVAILABILITY_UNAVAILABLE": 2,
		"NAME_AVAILABILITY_INVALID":     3,
		"NAME_AVAILABILITY_PROFANE":     4,
		"NAME_AVAILABILITY_RESERVED":    5,
	}
)

func (x NameAvailability) Enum() *NameAvailability {
	p := new(NameAvailability)
	*p = x
	return p
}

func (x NameAvailability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NameAvailability) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_api_v1_name_availability_proto_enumTypes[0].Descriptor()
}

func (NameAvailability) Type() protoreflect.EnumType {
	return &file_aserto_api_v1_name_availability_proto_enumTypes[0]
}

func (x NameAvailability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NameAvailability.Descriptor instead.
func (NameAvailability) EnumDescriptor() ([]byte, []int) {
	return file_aserto_api_v1_name_availability_proto_rawDescGZIP(), []int{0}
}

var File_aserto_api_v1_name_availability_proto protoreflect.FileDescriptor

var file_aserto_api_v1_name_availability_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2a, 0xd3, 0x01, 0x0a, 0x10, 0x4e, 0x61, 0x6d, 0x65, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x19, 0x4e,
	0x41, 0x4d, 0x45, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x4e, 0x41,
	0x4d, 0x45, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f,
	0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x4e,
	0x41, 0x4d, 0x45, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59,
	0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x1d,
	0x0a, 0x19, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x49, 0x4c,
	0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x03, 0x12, 0x1d, 0x0a,
	0x19, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x49, 0x4c, 0x49,
	0x54, 0x59, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x41, 0x4e, 0x45, 0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a,
	0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x44, 0x10, 0x05, 0x42, 0x41, 0x5a, 0x2f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0xaa,
	0x02, 0x0d, 0x41, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_api_v1_name_availability_proto_rawDescOnce sync.Once
	file_aserto_api_v1_name_availability_proto_rawDescData = file_aserto_api_v1_name_availability_proto_rawDesc
)

func file_aserto_api_v1_name_availability_proto_rawDescGZIP() []byte {
	file_aserto_api_v1_name_availability_proto_rawDescOnce.Do(func() {
		file_aserto_api_v1_name_availability_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_api_v1_name_availability_proto_rawDescData)
	})
	return file_aserto_api_v1_name_availability_proto_rawDescData
}

var file_aserto_api_v1_name_availability_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_api_v1_name_availability_proto_goTypes = []interface{}{
	(NameAvailability)(0), // 0: aserto.api.v1.NameAvailability
}
var file_aserto_api_v1_name_availability_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aserto_api_v1_name_availability_proto_init() }
func file_aserto_api_v1_name_availability_proto_init() {
	if File_aserto_api_v1_name_availability_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aserto_api_v1_name_availability_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_api_v1_name_availability_proto_goTypes,
		DependencyIndexes: file_aserto_api_v1_name_availability_proto_depIdxs,
		EnumInfos:         file_aserto_api_v1_name_availability_proto_enumTypes,
	}.Build()
	File_aserto_api_v1_name_availability_proto = out.File
	file_aserto_api_v1_name_availability_proto_rawDesc = nil
	file_aserto_api_v1_name_availability_proto_goTypes = nil
	file_aserto_api_v1_name_availability_proto_depIdxs = nil
}
