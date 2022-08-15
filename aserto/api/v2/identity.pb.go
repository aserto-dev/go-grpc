// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: aserto/api/v2/identity.proto

package api

import (
	v1 "github.com/aserto-dev/go-grpc/aserto/api/v1"
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

type IdentityProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind         v1.IdentityKind `protobuf:"varint,1,opt,name=kind,proto3,enum=aserto.api.v1.IdentityKind" json:"kind,omitempty"`          // identity kind [email|username|uid|pid|dn|phone]
	Provider     string          `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`                                   // identity provider name
	Verified     bool            `protobuf:"varint,3,opt,name=verified,proto3" json:"verified,omitempty"`                                  // identity has been verified (false when not explicitly specified)
	ConnectionId *string         `protobuf:"bytes,4,opt,name=connection_id,json=connectionId,proto3,oneof" json:"connection_id,omitempty"` // IDP connection id which owns the object instance
}

func (x *IdentityProperties) Reset() {
	*x = IdentityProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v2_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityProperties) ProtoMessage() {}

func (x *IdentityProperties) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v2_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityProperties.ProtoReflect.Descriptor instead.
func (*IdentityProperties) Descriptor() ([]byte, []int) {
	return file_aserto_api_v2_identity_proto_rawDescGZIP(), []int{0}
}

func (x *IdentityProperties) GetKind() v1.IdentityKind {
	if x != nil {
		return x.Kind
	}
	return v1.IdentityKind(0)
}

func (x *IdentityProperties) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *IdentityProperties) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *IdentityProperties) GetConnectionId() string {
	if x != nil && x.ConnectionId != nil {
		return *x.ConnectionId
	}
	return ""
}

var File_aserto_api_v2_identity_proto protoreflect.FileDescriptor

var file_aserto_api_v2_identity_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x1a, 0x18, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x12, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2f,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x42, 0x41, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x32, 0x3b, 0x61, 0x70, 0x69, 0xaa, 0x02, 0x0d, 0x41, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x41, 0x50, 0x49, 0x2e, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_api_v2_identity_proto_rawDescOnce sync.Once
	file_aserto_api_v2_identity_proto_rawDescData = file_aserto_api_v2_identity_proto_rawDesc
)

func file_aserto_api_v2_identity_proto_rawDescGZIP() []byte {
	file_aserto_api_v2_identity_proto_rawDescOnce.Do(func() {
		file_aserto_api_v2_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_api_v2_identity_proto_rawDescData)
	})
	return file_aserto_api_v2_identity_proto_rawDescData
}

var file_aserto_api_v2_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_aserto_api_v2_identity_proto_goTypes = []interface{}{
	(*IdentityProperties)(nil), // 0: aserto.api.v2.IdentityProperties
	(v1.IdentityKind)(0),       // 1: aserto.api.v1.IdentityKind
}
var file_aserto_api_v2_identity_proto_depIdxs = []int32{
	1, // 0: aserto.api.v2.IdentityProperties.kind:type_name -> aserto.api.v1.IdentityKind
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aserto_api_v2_identity_proto_init() }
func file_aserto_api_v2_identity_proto_init() {
	if File_aserto_api_v2_identity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_api_v2_identity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityProperties); i {
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
	file_aserto_api_v2_identity_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aserto_api_v2_identity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_api_v2_identity_proto_goTypes,
		DependencyIndexes: file_aserto_api_v2_identity_proto_depIdxs,
		MessageInfos:      file_aserto_api_v2_identity_proto_msgTypes,
	}.Build()
	File_aserto_api_v2_identity_proto = out.File
	file_aserto_api_v2_identity_proto_rawDesc = nil
	file_aserto_api_v2_identity_proto_goTypes = nil
	file_aserto_api_v2_identity_proto_depIdxs = nil
}
