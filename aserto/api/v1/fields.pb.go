// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: aserto/api/v1/fields.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Fields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mask *fieldmaskpb.FieldMask `protobuf:"bytes,1,opt,name=mask,proto3" json:"mask,omitempty"`
}

func (x *Fields) Reset() {
	*x = Fields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v1_fields_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fields) ProtoMessage() {}

func (x *Fields) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v1_fields_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fields.ProtoReflect.Descriptor instead.
func (*Fields) Descriptor() ([]byte, []int) {
	return file_aserto_api_v1_fields_proto_rawDescGZIP(), []int{0}
}

func (x *Fields) GetMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.Mask
	}
	return nil
}

var File_aserto_api_v1_fields_proto protoreflect.FileDescriptor

var file_aserto_api_v1_fields_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a,
	0x06, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x42, 0x41, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76,
	0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0xaa, 0x02, 0x0d, 0x41, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_aserto_api_v1_fields_proto_rawDescOnce sync.Once
	file_aserto_api_v1_fields_proto_rawDescData = file_aserto_api_v1_fields_proto_rawDesc
)

func file_aserto_api_v1_fields_proto_rawDescGZIP() []byte {
	file_aserto_api_v1_fields_proto_rawDescOnce.Do(func() {
		file_aserto_api_v1_fields_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_api_v1_fields_proto_rawDescData)
	})
	return file_aserto_api_v1_fields_proto_rawDescData
}

var file_aserto_api_v1_fields_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_aserto_api_v1_fields_proto_goTypes = []any{
	(*Fields)(nil),                // 0: aserto.api.v1.Fields
	(*fieldmaskpb.FieldMask)(nil), // 1: google.protobuf.FieldMask
}
var file_aserto_api_v1_fields_proto_depIdxs = []int32{
	1, // 0: aserto.api.v1.Fields.mask:type_name -> google.protobuf.FieldMask
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aserto_api_v1_fields_proto_init() }
func file_aserto_api_v1_fields_proto_init() {
	if File_aserto_api_v1_fields_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_api_v1_fields_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Fields); i {
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
			RawDescriptor: file_aserto_api_v1_fields_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_api_v1_fields_proto_goTypes,
		DependencyIndexes: file_aserto_api_v1_fields_proto_depIdxs,
		MessageInfos:      file_aserto_api_v1_fields_proto_msgTypes,
	}.Build()
	File_aserto_api_v1_fields_proto = out.File
	file_aserto_api_v1_fields_proto_rawDesc = nil
	file_aserto_api_v1_fields_proto_goTypes = nil
	file_aserto_api_v1_fields_proto_depIdxs = nil
}
