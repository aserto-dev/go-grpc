// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/api/v2/source.proto

package api

import (
	_ "github.com/aserto-dev/go-grpc/aserto/options/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Source struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PolicyId      string                 `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	ConnectionId  string                 `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	Org           string                 `protobuf:"bytes,3,opt,name=org,proto3" json:"org,omitempty"`
	Repo          string                 `protobuf:"bytes,4,opt,name=repo,proto3" json:"repo,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	VersionHash   string                 `protobuf:"bytes,7,opt,name=version_hash,json=versionHash,proto3" json:"version_hash,omitempty"` // Sha256 string of policy_id, connection_id, org and repo concatenated values.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Source) Reset() {
	*x = Source{}
	mi := &file_aserto_api_v2_source_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v2_source_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_aserto_api_v2_source_proto_rawDescGZIP(), []int{0}
}

func (x *Source) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

func (x *Source) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

func (x *Source) GetOrg() string {
	if x != nil {
		return x.Org
	}
	return ""
}

func (x *Source) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *Source) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Source) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Source) GetVersionHash() string {
	if x != nil {
		return x.VersionHash
	}
	return ""
}

var File_aserto_api_v2_source_proto protoreflect.FileDescriptor

const file_aserto_api_v2_source_proto_rawDesc = "" +
	"\n" +
	"\x1aaserto/api/v2/source.proto\x12\raserto.api.v2\x1a\x1baserto/options/v1/ids.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xb9\x02\n" +
	"\x06Source\x12!\n" +
	"\tpolicy_id\x18\x01 \x01(\tB\x04\x90\xb5\x18\x04R\bpolicyId\x12)\n" +
	"\rconnection_id\x18\x02 \x01(\tB\x04\x90\xb5\x18\aR\fconnectionId\x12\x10\n" +
	"\x03org\x18\x03 \x01(\tR\x03org\x12\x12\n" +
	"\x04repo\x18\x04 \x01(\tR\x04repo\x129\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12!\n" +
	"\fversion_hash\x18\a \x01(\tR\vversionHash:\"\x92A\x1f\n" +
	"\x1d\xd2\x01\rconnection_id\xd2\x01\x03org\xd2\x01\x04repoBAZ/github.com/aserto-dev/go-grpc/aserto/api/v2;api\xaa\x02\rAserto.API.V2b\x06proto3"

var (
	file_aserto_api_v2_source_proto_rawDescOnce sync.Once
	file_aserto_api_v2_source_proto_rawDescData []byte
)

func file_aserto_api_v2_source_proto_rawDescGZIP() []byte {
	file_aserto_api_v2_source_proto_rawDescOnce.Do(func() {
		file_aserto_api_v2_source_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_api_v2_source_proto_rawDesc), len(file_aserto_api_v2_source_proto_rawDesc)))
	})
	return file_aserto_api_v2_source_proto_rawDescData
}

var file_aserto_api_v2_source_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_aserto_api_v2_source_proto_goTypes = []any{
	(*Source)(nil),                // 0: aserto.api.v2.Source
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_aserto_api_v2_source_proto_depIdxs = []int32{
	1, // 0: aserto.api.v2.Source.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: aserto.api.v2.Source.updated_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_aserto_api_v2_source_proto_init() }
func file_aserto_api_v2_source_proto_init() {
	if File_aserto_api_v2_source_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_api_v2_source_proto_rawDesc), len(file_aserto_api_v2_source_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_api_v2_source_proto_goTypes,
		DependencyIndexes: file_aserto_api_v2_source_proto_depIdxs,
		MessageInfos:      file_aserto_api_v2_source_proto_msgTypes,
	}.Build()
	File_aserto_api_v2_source_proto = out.File
	file_aserto_api_v2_source_proto_goTypes = nil
	file_aserto_api_v2_source_proto_depIdxs = nil
}
