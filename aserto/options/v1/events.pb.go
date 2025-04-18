// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/options/v1/events.proto

package options

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

type Event struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	SubjectTemplate string                 `protobuf:"bytes,2,opt,name=subject_template,json=subjectTemplate,proto3" json:"subject_template,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_aserto_options_v1_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_options_v1_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_aserto_options_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetSubjectTemplate() string {
	if x != nil {
		return x.SubjectTemplate
	}
	return ""
}

var file_aserto_options_v1_events_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*Event)(nil),
		Field:         50095,
		Name:          "aserto.options.v1.event",
		Tag:           "bytes,50095,opt,name=event",
		Filename:      "aserto/options/v1/events.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50095,
		Name:          "aserto.options.v1.tag",
		Tag:           "bytes,50095,opt,name=tag",
		Filename:      "aserto/options/v1/events.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional aserto.options.v1.Event event = 50095;
	E_Event = &file_aserto_options_v1_events_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string tag = 50095;
	E_Tag = &file_aserto_options_v1_events_proto_extTypes[1]
)

var File_aserto_options_v1_events_proto protoreflect.FileDescriptor

const file_aserto_options_v1_events_proto_rawDesc = "" +
	"\n" +
	"\x1easerto/options/v1/events.proto\x12\x11aserto.options.v1\x1a google/protobuf/descriptor.proto\"2\n" +
	"\x05Event\x12)\n" +
	"\x10subject_template\x18\x02 \x01(\tR\x0fsubjectTemplate:Q\n" +
	"\x05event\x12\x1f.google.protobuf.MessageOptions\x18\xaf\x87\x03 \x01(\v2\x18.aserto.options.v1.EventR\x05event:1\n" +
	"\x03tag\x12\x1d.google.protobuf.FieldOptions\x18\xaf\x87\x03 \x01(\tR\x03tagB9Z7github.com/aserto-dev/go-grpc/aserto/options/v1;optionsb\x06proto3"

var (
	file_aserto_options_v1_events_proto_rawDescOnce sync.Once
	file_aserto_options_v1_events_proto_rawDescData []byte
)

func file_aserto_options_v1_events_proto_rawDescGZIP() []byte {
	file_aserto_options_v1_events_proto_rawDescOnce.Do(func() {
		file_aserto_options_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_options_v1_events_proto_rawDesc), len(file_aserto_options_v1_events_proto_rawDesc)))
	})
	return file_aserto_options_v1_events_proto_rawDescData
}

var file_aserto_options_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_aserto_options_v1_events_proto_goTypes = []any{
	(*Event)(nil),                       // 0: aserto.options.v1.Event
	(*descriptorpb.MessageOptions)(nil), // 1: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 2: google.protobuf.FieldOptions
}
var file_aserto_options_v1_events_proto_depIdxs = []int32{
	1, // 0: aserto.options.v1.event:extendee -> google.protobuf.MessageOptions
	2, // 1: aserto.options.v1.tag:extendee -> google.protobuf.FieldOptions
	0, // 2: aserto.options.v1.event:type_name -> aserto.options.v1.Event
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aserto_options_v1_events_proto_init() }
func file_aserto_options_v1_events_proto_init() {
	if File_aserto_options_v1_events_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_options_v1_events_proto_rawDesc), len(file_aserto_options_v1_events_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_aserto_options_v1_events_proto_goTypes,
		DependencyIndexes: file_aserto_options_v1_events_proto_depIdxs,
		MessageInfos:      file_aserto_options_v1_events_proto_msgTypes,
		ExtensionInfos:    file_aserto_options_v1_events_proto_extTypes,
	}.Build()
	File_aserto_options_v1_events_proto = out.File
	file_aserto_options_v1_events_proto_goTypes = nil
	file_aserto_options_v1_events_proto_depIdxs = nil
}
