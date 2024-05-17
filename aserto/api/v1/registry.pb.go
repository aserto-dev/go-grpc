// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: aserto/api/v1/registry.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegistryRepo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Org  string `protobuf:"bytes,2,opt,name=org,proto3" json:"org,omitempty"`
}

func (x *RegistryRepo) Reset() {
	*x = RegistryRepo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v1_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryRepo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryRepo) ProtoMessage() {}

func (x *RegistryRepo) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v1_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryRepo.ProtoReflect.Descriptor instead.
func (*RegistryRepo) Descriptor() ([]byte, []int) {
	return file_aserto_api_v1_registry_proto_rawDescGZIP(), []int{0}
}

func (x *RegistryRepo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegistryRepo) GetOrg() string {
	if x != nil {
		return x.Org
	}
	return ""
}

type RegistryRepoTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Digest      string                    `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	Size        int64                     `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	CreatedAt   *timestamppb.Timestamp    `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Annotations []*RegistryRepoAnnotation `protobuf:"bytes,5,rep,name=annotations,proto3" json:"annotations,omitempty"`
}

func (x *RegistryRepoTag) Reset() {
	*x = RegistryRepoTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v1_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryRepoTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryRepoTag) ProtoMessage() {}

func (x *RegistryRepoTag) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v1_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryRepoTag.ProtoReflect.Descriptor instead.
func (*RegistryRepoTag) Descriptor() ([]byte, []int) {
	return file_aserto_api_v1_registry_proto_rawDescGZIP(), []int{1}
}

func (x *RegistryRepoTag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegistryRepoTag) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *RegistryRepoTag) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *RegistryRepoTag) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *RegistryRepoTag) GetAnnotations() []*RegistryRepoAnnotation {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type RegistryRepoAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *RegistryRepoAnnotation) Reset() {
	*x = RegistryRepoAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v1_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryRepoAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryRepoAnnotation) ProtoMessage() {}

func (x *RegistryRepoAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v1_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryRepoAnnotation.ProtoReflect.Descriptor instead.
func (*RegistryRepoAnnotation) Descriptor() ([]byte, []int) {
	return file_aserto_api_v1_registry_proto_rawDescGZIP(), []int{2}
}

func (x *RegistryRepoAnnotation) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RegistryRepoAnnotation) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type RegistryRepoDigest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Digest    string                 `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	Tags      []string               `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *RegistryRepoDigest) Reset() {
	*x = RegistryRepoDigest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v1_registry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryRepoDigest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryRepoDigest) ProtoMessage() {}

func (x *RegistryRepoDigest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v1_registry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryRepoDigest.ProtoReflect.Descriptor instead.
func (*RegistryRepoDigest) Descriptor() ([]byte, []int) {
	return file_aserto_api_v1_registry_proto_rawDescGZIP(), []int{3}
}

func (x *RegistryRepoDigest) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *RegistryRepoDigest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *RegistryRepoDigest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_aserto_api_v1_registry_proto protoreflect.FileDescriptor

var file_aserto_api_v1_registry_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34,
	0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x72, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6f, 0x72, 0x67, 0x22, 0xd5, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x70, 0x6f, 0x54, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x47, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x70, 0x6f, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x40, 0x0a, 0x16,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7b,
	0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x44, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x41, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0xaa, 0x02,
	0x0d, 0x41, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_api_v1_registry_proto_rawDescOnce sync.Once
	file_aserto_api_v1_registry_proto_rawDescData = file_aserto_api_v1_registry_proto_rawDesc
)

func file_aserto_api_v1_registry_proto_rawDescGZIP() []byte {
	file_aserto_api_v1_registry_proto_rawDescOnce.Do(func() {
		file_aserto_api_v1_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_api_v1_registry_proto_rawDescData)
	})
	return file_aserto_api_v1_registry_proto_rawDescData
}

var file_aserto_api_v1_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_aserto_api_v1_registry_proto_goTypes = []interface{}{
	(*RegistryRepo)(nil),           // 0: aserto.api.v1.RegistryRepo
	(*RegistryRepoTag)(nil),        // 1: aserto.api.v1.RegistryRepoTag
	(*RegistryRepoAnnotation)(nil), // 2: aserto.api.v1.RegistryRepoAnnotation
	(*RegistryRepoDigest)(nil),     // 3: aserto.api.v1.RegistryRepoDigest
	(*timestamppb.Timestamp)(nil),  // 4: google.protobuf.Timestamp
}
var file_aserto_api_v1_registry_proto_depIdxs = []int32{
	4, // 0: aserto.api.v1.RegistryRepoTag.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: aserto.api.v1.RegistryRepoTag.annotations:type_name -> aserto.api.v1.RegistryRepoAnnotation
	4, // 2: aserto.api.v1.RegistryRepoDigest.created_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_aserto_api_v1_registry_proto_init() }
func file_aserto_api_v1_registry_proto_init() {
	if File_aserto_api_v1_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_api_v1_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryRepo); i {
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
		file_aserto_api_v1_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryRepoTag); i {
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
		file_aserto_api_v1_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryRepoAnnotation); i {
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
		file_aserto_api_v1_registry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryRepoDigest); i {
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
			RawDescriptor: file_aserto_api_v1_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_api_v1_registry_proto_goTypes,
		DependencyIndexes: file_aserto_api_v1_registry_proto_depIdxs,
		MessageInfos:      file_aserto_api_v1_registry_proto_msgTypes,
	}.Build()
	File_aserto_api_v1_registry_proto = out.File
	file_aserto_api_v1_registry_proto_rawDesc = nil
	file_aserto_api_v1_registry_proto_goTypes = nil
	file_aserto_api_v1_registry_proto_depIdxs = nil
}
