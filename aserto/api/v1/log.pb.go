// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: aserto/api/v1/log.proto

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

// Log Level, describes the level of the log from the services.
type LogLevel int32

const (
	LogLevel_LOG_LEVEL_UNKNOWN LogLevel = 0
	LogLevel_LOG_LEVEL_TRACE   LogLevel = 1
	LogLevel_LOG_LEVEL_DEBUG   LogLevel = 2
	LogLevel_LOG_LEVEL_INFO    LogLevel = 3
	LogLevel_LOG_LEVEL_WARN    LogLevel = 4
	LogLevel_LOG_LEVEL_ERROR   LogLevel = 5
	LogLevel_LOG_LEVEL_FATAL   LogLevel = 6
	LogLevel_LOG_LEVEL_PANIC   LogLevel = 7
)

// Enum value maps for LogLevel.
var (
	LogLevel_name = map[int32]string{
		0: "LOG_LEVEL_UNKNOWN",
		1: "LOG_LEVEL_TRACE",
		2: "LOG_LEVEL_DEBUG",
		3: "LOG_LEVEL_INFO",
		4: "LOG_LEVEL_WARN",
		5: "LOG_LEVEL_ERROR",
		6: "LOG_LEVEL_FATAL",
		7: "LOG_LEVEL_PANIC",
	}
	LogLevel_value = map[string]int32{
		"LOG_LEVEL_UNKNOWN": 0,
		"LOG_LEVEL_TRACE":   1,
		"LOG_LEVEL_DEBUG":   2,
		"LOG_LEVEL_INFO":    3,
		"LOG_LEVEL_WARN":    4,
		"LOG_LEVEL_ERROR":   5,
		"LOG_LEVEL_FATAL":   6,
		"LOG_LEVEL_PANIC":   7,
	}
)

func (x LogLevel) Enum() *LogLevel {
	p := new(LogLevel)
	*p = x
	return p
}

func (x LogLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_api_v1_log_proto_enumTypes[0].Descriptor()
}

func (LogLevel) Type() protoreflect.EnumType {
	return &file_aserto_api_v1_log_proto_enumTypes[0]
}

func (x LogLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogLevel.Descriptor instead.
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return file_aserto_api_v1_log_proto_rawDescGZIP(), []int{0}
}

var File_aserto_api_v1_log_proto protoreflect.FileDescriptor

var file_aserto_api_v1_log_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2a, 0xb2, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x67,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f,
	0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x45, 0x10,
	0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x44,
	0x45, 0x42, 0x55, 0x47, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x4f,
	0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x10, 0x04, 0x12, 0x13,
	0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c,
	0x5f, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x50, 0x41, 0x4e, 0x49, 0x43, 0x10, 0x07, 0x42, 0x41, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69,
	0xaa, 0x02, 0x0d, 0x41, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_api_v1_log_proto_rawDescOnce sync.Once
	file_aserto_api_v1_log_proto_rawDescData = file_aserto_api_v1_log_proto_rawDesc
)

func file_aserto_api_v1_log_proto_rawDescGZIP() []byte {
	file_aserto_api_v1_log_proto_rawDescOnce.Do(func() {
		file_aserto_api_v1_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_api_v1_log_proto_rawDescData)
	})
	return file_aserto_api_v1_log_proto_rawDescData
}

var file_aserto_api_v1_log_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_api_v1_log_proto_goTypes = []interface{}{
	(LogLevel)(0), // 0: aserto.api.v1.LogLevel
}
var file_aserto_api_v1_log_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aserto_api_v1_log_proto_init() }
func file_aserto_api_v1_log_proto_init() {
	if File_aserto_api_v1_log_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aserto_api_v1_log_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_api_v1_log_proto_goTypes,
		DependencyIndexes: file_aserto_api_v1_log_proto_depIdxs,
		EnumInfos:         file_aserto_api_v1_log_proto_enumTypes,
	}.Build()
	File_aserto_api_v1_log_proto = out.File
	file_aserto_api_v1_log_proto_rawDesc = nil
	file_aserto_api_v1_log_proto_goTypes = nil
	file_aserto_api_v1_log_proto_depIdxs = nil
}
