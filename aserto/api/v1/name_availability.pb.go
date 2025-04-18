// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/api/v1/name_availability.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
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

const file_aserto_api_v1_name_availability_proto_rawDesc = "" +
	"\n" +
	"%aserto/api/v1/name_availability.proto\x12\raserto.api.v1*\xd3\x01\n" +
	"\x10NameAvailability\x12\x1d\n" +
	"\x19NAME_AVAILABILITY_UNKNOWN\x10\x00\x12\x1f\n" +
	"\x1bNAME_AVAILABILITY_AVAILABLE\x10\x01\x12!\n" +
	"\x1dNAME_AVAILABILITY_UNAVAILABLE\x10\x02\x12\x1d\n" +
	"\x19NAME_AVAILABILITY_INVALID\x10\x03\x12\x1d\n" +
	"\x19NAME_AVAILABILITY_PROFANE\x10\x04\x12\x1e\n" +
	"\x1aNAME_AVAILABILITY_RESERVED\x10\x05BAZ/github.com/aserto-dev/go-grpc/aserto/api/v1;api\xaa\x02\rAserto.API.V1b\x06proto3"

var (
	file_aserto_api_v1_name_availability_proto_rawDescOnce sync.Once
	file_aserto_api_v1_name_availability_proto_rawDescData []byte
)

func file_aserto_api_v1_name_availability_proto_rawDescGZIP() []byte {
	file_aserto_api_v1_name_availability_proto_rawDescOnce.Do(func() {
		file_aserto_api_v1_name_availability_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_api_v1_name_availability_proto_rawDesc), len(file_aserto_api_v1_name_availability_proto_rawDesc)))
	})
	return file_aserto_api_v1_name_availability_proto_rawDescData
}

var file_aserto_api_v1_name_availability_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_api_v1_name_availability_proto_goTypes = []any{
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_api_v1_name_availability_proto_rawDesc), len(file_aserto_api_v1_name_availability_proto_rawDesc)),
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
	file_aserto_api_v1_name_availability_proto_goTypes = nil
	file_aserto_api_v1_name_availability_proto_depIdxs = nil
}
