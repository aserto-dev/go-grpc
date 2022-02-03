// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: aserto/api/v1/nats.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NatsMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload  *anypb.Any        `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NatsMessage) Reset() {
	*x = NatsMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v1_nats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NatsMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NatsMessage) ProtoMessage() {}

func (x *NatsMessage) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v1_nats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_aserto_api_v1_nats_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc0, 0x01, 0x0a, 0x0b, 0x4e, 0x61, 0x74, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x44, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x61, 0x74, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x41, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0xaa, 0x02, 0x0d, 0x41, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x41, 0x50, 0x49, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_api_v1_nats_proto_rawDescOnce sync.Once
	file_aserto_api_v1_nats_proto_rawDescData = file_aserto_api_v1_nats_proto_rawDesc
)

func file_aserto_api_v1_nats_proto_rawDescGZIP() []byte {
	file_aserto_api_v1_nats_proto_rawDescOnce.Do(func() {
		file_aserto_api_v1_nats_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_api_v1_nats_proto_rawDescData)
	})
	return file_aserto_api_v1_nats_proto_rawDescData
}

var file_aserto_api_v1_nats_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aserto_api_v1_nats_proto_goTypes = []interface{}{
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
	file_aserto_api_v1_metadata_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_aserto_api_v1_nats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NatsMessage); i {
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
			RawDescriptor: file_aserto_api_v1_nats_proto_rawDesc,
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
	file_aserto_api_v1_nats_proto_rawDesc = nil
	file_aserto_api_v1_nats_proto_goTypes = nil
	file_aserto_api_v1_nats_proto_depIdxs = nil
}
