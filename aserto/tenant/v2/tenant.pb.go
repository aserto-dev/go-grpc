// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/tenant/v2/tenant.proto

package tenant

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

type DeleteTenantRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTenantRequest) Reset() {
	*x = DeleteTenantRequest{}
	mi := &file_aserto_tenant_v2_tenant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTenantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTenantRequest) ProtoMessage() {}

func (x *DeleteTenantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_tenant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTenantRequest.ProtoReflect.Descriptor instead.
func (*DeleteTenantRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_tenant_proto_rawDescGZIP(), []int{0}
}

type DeleteTenantResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTenantResponse) Reset() {
	*x = DeleteTenantResponse{}
	mi := &file_aserto_tenant_v2_tenant_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTenantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTenantResponse) ProtoMessage() {}

func (x *DeleteTenantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_tenant_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTenantResponse.ProtoReflect.Descriptor instead.
func (*DeleteTenantResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_tenant_proto_rawDescGZIP(), []int{1}
}

var File_aserto_tenant_v2_tenant_proto protoreflect.FileDescriptor

const file_aserto_tenant_v2_tenant_proto_rawDesc = "" +
	"\n" +
	"\x1daserto/tenant/v2/tenant.proto\x12\x10aserto.tenant.v2\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\x15\n" +
	"\x13DeleteTenantRequest\"\x16\n" +
	"\x14DeleteTenantResponse2\xdb\x01\n" +
	"\x06Tenant\x12\xd0\x01\n" +
	"\fDeleteTenant\x12%.aserto.tenant.v2.DeleteTenantRequest\x1a&.aserto.tenant.v2.DeleteTenantResponse\"q\x92AW\n" +
	"\x06tenant\x12\x0eDelete tenant.\x1a\x0eDelete tenant.*\x14tenant.delete_tenantb\x17\n" +
	"\a\n" +
	"\x03JWT\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02\x11*\x0f/api/v2/tenantsB\xbd\x01\x92A\x82\x01*\x01\x022\x10application/json:\x10application/jsonZ@\n" +
	"\x1a\n" +
	"\x03JWT\x12\x13\b\x02\x1a\rauthorization \x02\n" +
	"\"\n" +
	"\bTenantID\x12\x16\b\x02\x1a\x10aserto-tenant-id \x02b\x17\n" +
	"\a\n" +
	"\x03JWT\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00Z5github.com/aserto-dev/go-grpc/aserto/tenant/v2;tenantb\x06proto3"

var (
	file_aserto_tenant_v2_tenant_proto_rawDescOnce sync.Once
	file_aserto_tenant_v2_tenant_proto_rawDescData []byte
)

func file_aserto_tenant_v2_tenant_proto_rawDescGZIP() []byte {
	file_aserto_tenant_v2_tenant_proto_rawDescOnce.Do(func() {
		file_aserto_tenant_v2_tenant_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_tenant_v2_tenant_proto_rawDesc), len(file_aserto_tenant_v2_tenant_proto_rawDesc)))
	})
	return file_aserto_tenant_v2_tenant_proto_rawDescData
}

var file_aserto_tenant_v2_tenant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aserto_tenant_v2_tenant_proto_goTypes = []any{
	(*DeleteTenantRequest)(nil),  // 0: aserto.tenant.v2.DeleteTenantRequest
	(*DeleteTenantResponse)(nil), // 1: aserto.tenant.v2.DeleteTenantResponse
}
var file_aserto_tenant_v2_tenant_proto_depIdxs = []int32{
	0, // 0: aserto.tenant.v2.Tenant.DeleteTenant:input_type -> aserto.tenant.v2.DeleteTenantRequest
	1, // 1: aserto.tenant.v2.Tenant.DeleteTenant:output_type -> aserto.tenant.v2.DeleteTenantResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aserto_tenant_v2_tenant_proto_init() }
func file_aserto_tenant_v2_tenant_proto_init() {
	if File_aserto_tenant_v2_tenant_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_tenant_v2_tenant_proto_rawDesc), len(file_aserto_tenant_v2_tenant_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_tenant_v2_tenant_proto_goTypes,
		DependencyIndexes: file_aserto_tenant_v2_tenant_proto_depIdxs,
		MessageInfos:      file_aserto_tenant_v2_tenant_proto_msgTypes,
	}.Build()
	File_aserto_tenant_v2_tenant_proto = out.File
	file_aserto_tenant_v2_tenant_proto_goTypes = nil
	file_aserto_tenant_v2_tenant_proto_depIdxs = nil
}
