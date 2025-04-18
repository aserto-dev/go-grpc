// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/tenant/v2/account.proto

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

type DeleteAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAccountRequest) Reset() {
	*x = DeleteAccountRequest{}
	mi := &file_aserto_tenant_v2_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountRequest) ProtoMessage() {}

func (x *DeleteAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountRequest.ProtoReflect.Descriptor instead.
func (*DeleteAccountRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_account_proto_rawDescGZIP(), []int{0}
}

type DeleteAccountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAccountResponse) Reset() {
	*x = DeleteAccountResponse{}
	mi := &file_aserto_tenant_v2_account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountResponse) ProtoMessage() {}

func (x *DeleteAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_account_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountResponse.ProtoReflect.Descriptor instead.
func (*DeleteAccountResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_account_proto_rawDescGZIP(), []int{1}
}

var File_aserto_tenant_v2_account_proto protoreflect.FileDescriptor

const file_aserto_tenant_v2_account_proto_rawDesc = "" +
	"\n" +
	"\x1easerto/tenant/v2/account.proto\x12\x10aserto.tenant.v2\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\x16\n" +
	"\x14DeleteAccountRequest\"\x17\n" +
	"\x15DeleteAccountResponse2\xd6\x01\n" +
	"\aAccount\x12\xca\x01\n" +
	"\rDeleteAccount\x12&.aserto.tenant.v2.DeleteAccountRequest\x1a'.aserto.tenant.v2.DeleteAccountResponse\"h\x92AM\n" +
	"\aaccount\x12\x0fDelete account.\x1a\x0fDelete account.*\x15tenant.delete_accountb\t\n" +
	"\a\n" +
	"\x03JWT\x12\x00\x82\xd3\xe4\x93\x02\x12*\x10/api/v2/accountsB\x8a\x01\x92AP*\x01\x022\x10application/json:\x10application/jsonZ\x1c\n" +
	"\x1a\n" +
	"\x03JWT\x12\x13\b\x02\x1a\rauthorization \x02b\t\n" +
	"\a\n" +
	"\x03JWT\x12\x00Z5github.com/aserto-dev/go-grpc/aserto/tenant/v2;tenantb\x06proto3"

var (
	file_aserto_tenant_v2_account_proto_rawDescOnce sync.Once
	file_aserto_tenant_v2_account_proto_rawDescData []byte
)

func file_aserto_tenant_v2_account_proto_rawDescGZIP() []byte {
	file_aserto_tenant_v2_account_proto_rawDescOnce.Do(func() {
		file_aserto_tenant_v2_account_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_tenant_v2_account_proto_rawDesc), len(file_aserto_tenant_v2_account_proto_rawDesc)))
	})
	return file_aserto_tenant_v2_account_proto_rawDescData
}

var file_aserto_tenant_v2_account_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aserto_tenant_v2_account_proto_goTypes = []any{
	(*DeleteAccountRequest)(nil),  // 0: aserto.tenant.v2.DeleteAccountRequest
	(*DeleteAccountResponse)(nil), // 1: aserto.tenant.v2.DeleteAccountResponse
}
var file_aserto_tenant_v2_account_proto_depIdxs = []int32{
	0, // 0: aserto.tenant.v2.Account.DeleteAccount:input_type -> aserto.tenant.v2.DeleteAccountRequest
	1, // 1: aserto.tenant.v2.Account.DeleteAccount:output_type -> aserto.tenant.v2.DeleteAccountResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aserto_tenant_v2_account_proto_init() }
func file_aserto_tenant_v2_account_proto_init() {
	if File_aserto_tenant_v2_account_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_tenant_v2_account_proto_rawDesc), len(file_aserto_tenant_v2_account_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_tenant_v2_account_proto_goTypes,
		DependencyIndexes: file_aserto_tenant_v2_account_proto_depIdxs,
		MessageInfos:      file_aserto_tenant_v2_account_proto_msgTypes,
	}.Build()
	File_aserto_tenant_v2_account_proto = out.File
	file_aserto_tenant_v2_account_proto_goTypes = nil
	file_aserto_tenant_v2_account_proto_depIdxs = nil
}
