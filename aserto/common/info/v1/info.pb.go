// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/common/info/v1/info.proto

package info

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type InfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	mi := &file_aserto_common_info_v1_info_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_common_info_v1_info_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_aserto_common_info_v1_info_proto_rawDescGZIP(), []int{0}
}

type InfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	System        *SystemInfo            `protobuf:"bytes,1,opt,name=system,proto3" json:"system,omitempty"`
	Version       *VersionInfo           `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Build         *BuildInfo             `protobuf:"bytes,3,opt,name=build,proto3" json:"build,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	mi := &file_aserto_common_info_v1_info_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_common_info_v1_info_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_aserto_common_info_v1_info_proto_rawDescGZIP(), []int{1}
}

func (x *InfoResponse) GetSystem() *SystemInfo {
	if x != nil {
		return x.System
	}
	return nil
}

func (x *InfoResponse) GetVersion() *VersionInfo {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *InfoResponse) GetBuild() *BuildInfo {
	if x != nil {
		return x.Build
	}
	return nil
}

type SystemInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InstanceId    string                 `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SystemInfo) Reset() {
	*x = SystemInfo{}
	mi := &file_aserto_common_info_v1_info_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SystemInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemInfo) ProtoMessage() {}

func (x *SystemInfo) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_common_info_v1_info_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemInfo.ProtoReflect.Descriptor instead.
func (*SystemInfo) Descriptor() ([]byte, []int) {
	return file_aserto_common_info_v1_info_proto_rawDescGZIP(), []int{2}
}

func (x *SystemInfo) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *SystemInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type VersionInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	System        int32                  `protobuf:"varint,2,opt,name=system,proto3" json:"system,omitempty"`
	Schema        string                 `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VersionInfo) Reset() {
	*x = VersionInfo{}
	mi := &file_aserto_common_info_v1_info_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VersionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionInfo) ProtoMessage() {}

func (x *VersionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_common_info_v1_info_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionInfo.ProtoReflect.Descriptor instead.
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return file_aserto_common_info_v1_info_proto_rawDescGZIP(), []int{3}
}

func (x *VersionInfo) GetSystem() int32 {
	if x != nil {
		return x.System
	}
	return 0
}

func (x *VersionInfo) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

type BuildInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Commit        string                 `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	Date          string                 `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Os            string                 `protobuf:"bytes,4,opt,name=os,proto3" json:"os,omitempty"`
	Arch          string                 `protobuf:"bytes,5,opt,name=arch,proto3" json:"arch,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BuildInfo) Reset() {
	*x = BuildInfo{}
	mi := &file_aserto_common_info_v1_info_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BuildInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildInfo) ProtoMessage() {}

func (x *BuildInfo) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_common_info_v1_info_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildInfo.ProtoReflect.Descriptor instead.
func (*BuildInfo) Descriptor() ([]byte, []int) {
	return file_aserto_common_info_v1_info_proto_rawDescGZIP(), []int{4}
}

func (x *BuildInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *BuildInfo) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

func (x *BuildInfo) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *BuildInfo) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *BuildInfo) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

var File_aserto_common_info_v1_info_proto protoreflect.FileDescriptor

const file_aserto_common_info_v1_info_proto_rawDesc = "" +
	"\n" +
	" aserto/common/info/v1/info.proto\x12\x15aserto.common.info.v1\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\r\n" +
	"\vInfoRequest\"\xbf\x01\n" +
	"\fInfoResponse\x129\n" +
	"\x06system\x18\x01 \x01(\v2!.aserto.common.info.v1.SystemInfoR\x06system\x12<\n" +
	"\aversion\x18\x02 \x01(\v2\".aserto.common.info.v1.VersionInfoR\aversion\x126\n" +
	"\x05build\x18\x03 \x01(\v2 .aserto.common.info.v1.BuildInfoR\x05build\"L\n" +
	"\n" +
	"SystemInfo\x12\x1f\n" +
	"\vinstance_id\x18\x01 \x01(\tR\n" +
	"instanceId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x02 \x01(\tR\tcreatedAt\"=\n" +
	"\vVersionInfo\x12\x16\n" +
	"\x06system\x18\x02 \x01(\x05R\x06system\x12\x16\n" +
	"\x06schema\x18\x03 \x01(\tR\x06schema\"u\n" +
	"\tBuildInfo\x12\x18\n" +
	"\aversion\x18\x01 \x01(\tR\aversion\x12\x16\n" +
	"\x06commit\x18\x02 \x01(\tR\x06commit\x12\x12\n" +
	"\x04date\x18\x03 \x01(\tR\x04date\x12\x0e\n" +
	"\x02os\x18\x04 \x01(\tR\x02os\x12\x12\n" +
	"\x04arch\x18\x05 \x01(\tR\x04arch2\xb8\x01\n" +
	"\x04Info\x12\xaf\x01\n" +
	"\x04Info\x12\".aserto.common.info.v1.InfoRequest\x1a#.aserto.common.info.v1.InfoResponse\"^\x92AG\n" +
	"\x04info\x12\rInfo endpoint\x1a#Return endpoint versio information.*\tinfo.infob\x00\x82\xd3\xe4\x93\x02\x0e\x12\f/api/v1/infoBd\x92A'*\x01\x022\x10application/json:\x10application/jsonZ8github.com/aserto-dev/go-grpc/aserto/common/info/v1;infob\x06proto3"

var (
	file_aserto_common_info_v1_info_proto_rawDescOnce sync.Once
	file_aserto_common_info_v1_info_proto_rawDescData []byte
)

func file_aserto_common_info_v1_info_proto_rawDescGZIP() []byte {
	file_aserto_common_info_v1_info_proto_rawDescOnce.Do(func() {
		file_aserto_common_info_v1_info_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_common_info_v1_info_proto_rawDesc), len(file_aserto_common_info_v1_info_proto_rawDesc)))
	})
	return file_aserto_common_info_v1_info_proto_rawDescData
}

var file_aserto_common_info_v1_info_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_aserto_common_info_v1_info_proto_goTypes = []any{
	(*InfoRequest)(nil),  // 0: aserto.common.info.v1.InfoRequest
	(*InfoResponse)(nil), // 1: aserto.common.info.v1.InfoResponse
	(*SystemInfo)(nil),   // 2: aserto.common.info.v1.SystemInfo
	(*VersionInfo)(nil),  // 3: aserto.common.info.v1.VersionInfo
	(*BuildInfo)(nil),    // 4: aserto.common.info.v1.BuildInfo
}
var file_aserto_common_info_v1_info_proto_depIdxs = []int32{
	2, // 0: aserto.common.info.v1.InfoResponse.system:type_name -> aserto.common.info.v1.SystemInfo
	3, // 1: aserto.common.info.v1.InfoResponse.version:type_name -> aserto.common.info.v1.VersionInfo
	4, // 2: aserto.common.info.v1.InfoResponse.build:type_name -> aserto.common.info.v1.BuildInfo
	0, // 3: aserto.common.info.v1.Info.Info:input_type -> aserto.common.info.v1.InfoRequest
	1, // 4: aserto.common.info.v1.Info.Info:output_type -> aserto.common.info.v1.InfoResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_aserto_common_info_v1_info_proto_init() }
func file_aserto_common_info_v1_info_proto_init() {
	if File_aserto_common_info_v1_info_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_common_info_v1_info_proto_rawDesc), len(file_aserto_common_info_v1_info_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_common_info_v1_info_proto_goTypes,
		DependencyIndexes: file_aserto_common_info_v1_info_proto_depIdxs,
		MessageInfos:      file_aserto_common_info_v1_info_proto_msgTypes,
	}.Build()
	File_aserto_common_info_v1_info_proto = out.File
	file_aserto_common_info_v1_info_proto_goTypes = nil
	file_aserto_common_info_v1_info_proto_depIdxs = nil
}
