// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/api/v1/nats.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type NatsMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Payload       *anypb.Any             `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Metadata      map[string]string      `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NatsMessage) Reset() {
	*x = NatsMessage{}
	mi := &file_aserto_api_v1_nats_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NatsMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NatsMessage) ProtoMessage() {}

func (x *NatsMessage) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v1_nats_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NatsMessage.ProtoReflect.Descriptor instead.
func (*NatsMessage) Descriptor() ([]byte, []int) {
	return file_aserto_api_v1_nats_proto_rawDescGZIP(), []int{0}
}

func (x *NatsMessage) GetPayload() *anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *NatsMessage) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_aserto_api_v1_nats_proto protoreflect.FileDescriptor

const file_aserto_api_v1_nats_proto_rawDesc = "" +
	"\n" +
	"\x18aserto/api/v1/nats.proto\x12\raserto.api.v1\x1a\x19google/protobuf/any.proto\"\xc0\x01\n" +
	"\vNatsMessage\x12.\n" +
	"\apayload\x18\x01 \x01(\v2\x14.google.protobuf.AnyR\apayload\x12D\n" +
	"\bmetadata\x18\x02 \x03(\v2(.aserto.api.v1.NatsMessage.MetadataEntryR\bmetadata\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01BAZ/github.com/aserto-dev/go-grpc/aserto/api/v1;api\xaa\x02\rAserto.API.V1b\x06proto3"

var (
	file_aserto_api_v1_nats_proto_rawDescOnce sync.Once
	file_aserto_api_v1_nats_proto_rawDescData []byte
)

func file_aserto_api_v1_nats_proto_rawDescGZIP() []byte {
	file_aserto_api_v1_nats_proto_rawDescOnce.Do(func() {
		file_aserto_api_v1_nats_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_api_v1_nats_proto_rawDesc), len(file_aserto_api_v1_nats_proto_rawDesc)))
	})
	return file_aserto_api_v1_nats_proto_rawDescData
}

var file_aserto_api_v1_nats_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aserto_api_v1_nats_proto_goTypes = []any{
	(*NatsMessage)(nil), // 0: aserto.api.v1.NatsMessage
	nil,                 // 1: aserto.api.v1.NatsMessage.MetadataEntry
	(*anypb.Any)(nil),   // 2: google.protobuf.Any
}
var file_aserto_api_v1_nats_proto_depIdxs = []int32{
	2, // 0: aserto.api.v1.NatsMessage.payload:type_name -> google.protobuf.Any
	1, // 1: aserto.api.v1.NatsMessage.metadata:type_name -> aserto.api.v1.NatsMessage.MetadataEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_aserto_api_v1_nats_proto_init() }
func file_aserto_api_v1_nats_proto_init() {
	if File_aserto_api_v1_nats_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_api_v1_nats_proto_rawDesc), len(file_aserto_api_v1_nats_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_api_v1_nats_proto_goTypes,
		DependencyIndexes: file_aserto_api_v1_nats_proto_depIdxs,
		MessageInfos:      file_aserto_api_v1_nats_proto_msgTypes,
	}.Build()
	File_aserto_api_v1_nats_proto = out.File
	file_aserto_api_v1_nats_proto_goTypes = nil
	file_aserto_api_v1_nats_proto_depIdxs = nil
}
