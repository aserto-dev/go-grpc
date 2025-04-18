// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/api/v1/log.proto

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

const file_aserto_api_v1_log_proto_rawDesc = "" +
	"\n" +
	"\x17aserto/api/v1/log.proto\x12\raserto.api.v1*\xb2\x01\n" +
	"\bLogLevel\x12\x15\n" +
	"\x11LOG_LEVEL_UNKNOWN\x10\x00\x12\x13\n" +
	"\x0fLOG_LEVEL_TRACE\x10\x01\x12\x13\n" +
	"\x0fLOG_LEVEL_DEBUG\x10\x02\x12\x12\n" +
	"\x0eLOG_LEVEL_INFO\x10\x03\x12\x12\n" +
	"\x0eLOG_LEVEL_WARN\x10\x04\x12\x13\n" +
	"\x0fLOG_LEVEL_ERROR\x10\x05\x12\x13\n" +
	"\x0fLOG_LEVEL_FATAL\x10\x06\x12\x13\n" +
	"\x0fLOG_LEVEL_PANIC\x10\aBAZ/github.com/aserto-dev/go-grpc/aserto/api/v1;api\xaa\x02\rAserto.API.V1b\x06proto3"

var (
	file_aserto_api_v1_log_proto_rawDescOnce sync.Once
	file_aserto_api_v1_log_proto_rawDescData []byte
)

func file_aserto_api_v1_log_proto_rawDescGZIP() []byte {
	file_aserto_api_v1_log_proto_rawDescOnce.Do(func() {
		file_aserto_api_v1_log_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_api_v1_log_proto_rawDesc), len(file_aserto_api_v1_log_proto_rawDesc)))
	})
	return file_aserto_api_v1_log_proto_rawDescData
}

var file_aserto_api_v1_log_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_api_v1_log_proto_goTypes = []any{
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_api_v1_log_proto_rawDesc), len(file_aserto_api_v1_log_proto_rawDesc)),
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
	file_aserto_api_v1_log_proto_goTypes = nil
	file_aserto_api_v1_log_proto_depIdxs = nil
}
