// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/tenant/policy_builder/v1/policy_builder.proto

package policy_builder

import (
	v1 "github.com/aserto-dev/go-grpc/aserto/api/v1"
	_ "github.com/aserto-dev/go-grpc/aserto/options/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type CreatePolicyBuilderRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	PolicyBuilder    *v1.PolicyBuilder      `protobuf:"bytes,1,opt,name=policy_builder,json=policyBuilder,proto3" json:"policy_builder,omitempty"`
	ForceReconnect   bool                   `protobuf:"varint,2,opt,name=force_reconnect,json=forceReconnect,proto3" json:"force_reconnect,omitempty"`
	WorkflowFileName string                 `protobuf:"bytes,3,opt,name=workflow_file_name,json=workflowFileName,proto3" json:"workflow_file_name,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreatePolicyBuilderRequest) Reset() {
	*x = CreatePolicyBuilderRequest{}
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePolicyBuilderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyBuilderRequest) ProtoMessage() {}

func (x *CreatePolicyBuilderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyBuilderRequest.ProtoReflect.Descriptor instead.
func (*CreatePolicyBuilderRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePolicyBuilderRequest) GetPolicyBuilder() *v1.PolicyBuilder {
	if x != nil {
		return x.PolicyBuilder
	}
	return nil
}

func (x *CreatePolicyBuilderRequest) GetForceReconnect() bool {
	if x != nil {
		return x.ForceReconnect
	}
	return false
}

func (x *CreatePolicyBuilderRequest) GetWorkflowFileName() string {
	if x != nil {
		return x.WorkflowFileName
	}
	return ""
}

type CreatePolicyBuilderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePolicyBuilderResponse) Reset() {
	*x = CreatePolicyBuilderResponse{}
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePolicyBuilderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyBuilderResponse) ProtoMessage() {}

func (x *CreatePolicyBuilderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyBuilderResponse.ProtoReflect.Descriptor instead.
func (*CreatePolicyBuilderResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePolicyBuilderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListPolicyBuildersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RegistryRepo  string                 `protobuf:"bytes,1,opt,name=registry_repo,json=registryRepo,proto3" json:"registry_repo,omitempty"`
	RegistryOrg   string                 `protobuf:"bytes,2,opt,name=registry_org,json=registryOrg,proto3" json:"registry_org,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPolicyBuildersRequest) Reset() {
	*x = ListPolicyBuildersRequest{}
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPolicyBuildersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPolicyBuildersRequest) ProtoMessage() {}

func (x *ListPolicyBuildersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPolicyBuildersRequest.ProtoReflect.Descriptor instead.
func (*ListPolicyBuildersRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescGZIP(), []int{2}
}

func (x *ListPolicyBuildersRequest) GetRegistryRepo() string {
	if x != nil {
		return x.RegistryRepo
	}
	return ""
}

func (x *ListPolicyBuildersRequest) GetRegistryOrg() string {
	if x != nil {
		return x.RegistryOrg
	}
	return ""
}

type ListPolicyBuildersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*v1.PolicyBuilder    `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPolicyBuildersResponse) Reset() {
	*x = ListPolicyBuildersResponse{}
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPolicyBuildersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPolicyBuildersResponse) ProtoMessage() {}

func (x *ListPolicyBuildersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPolicyBuildersResponse.ProtoReflect.Descriptor instead.
func (*ListPolicyBuildersResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescGZIP(), []int{3}
}

func (x *ListPolicyBuildersResponse) GetResults() []*v1.PolicyBuilder {
	if x != nil {
		return x.Results
	}
	return nil
}

type DeletePolicyBuilderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePolicyBuilderRequest) Reset() {
	*x = DeletePolicyBuilderRequest{}
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePolicyBuilderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicyBuilderRequest) ProtoMessage() {}

func (x *DeletePolicyBuilderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicyBuilderRequest.ProtoReflect.Descriptor instead.
func (*DeletePolicyBuilderRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePolicyBuilderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePolicyBuilderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        *emptypb.Empty         `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeletePolicyBuilderResponse) Reset() {
	*x = DeletePolicyBuilderResponse{}
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeletePolicyBuilderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicyBuilderResponse) ProtoMessage() {}

func (x *DeletePolicyBuilderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicyBuilderResponse.ProtoReflect.Descriptor instead.
func (*DeletePolicyBuilderResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePolicyBuilderResponse) GetResult() *emptypb.Empty {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_aserto_tenant_policy_builder_v1_policy_builder_proto protoreflect.FileDescriptor

const file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDesc = "" +
	"\n" +
	"4aserto/tenant/policy_builder/v1/policy_builder.proto\x12\x1faserto.tenant.policy_builder.v1\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\"aserto/api/v1/policy_builder.proto\x1a\x1baserto/options/v1/ids.proto\"\xb8\x01\n" +
	"\x1aCreatePolicyBuilderRequest\x12C\n" +
	"\x0epolicy_builder\x18\x01 \x01(\v2\x1c.aserto.api.v1.PolicyBuilderR\rpolicyBuilder\x12'\n" +
	"\x0fforce_reconnect\x18\x02 \x01(\bR\x0eforceReconnect\x12,\n" +
	"\x12workflow_file_name\x18\x03 \x01(\tR\x10workflowFileName\"3\n" +
	"\x1bCreatePolicyBuilderResponse\x12\x14\n" +
	"\x02id\x18\x01 \x01(\tB\x04\x90\xb5\x18\tR\x02id\"c\n" +
	"\x19ListPolicyBuildersRequest\x12#\n" +
	"\rregistry_repo\x18\x01 \x01(\tR\fregistryRepo\x12!\n" +
	"\fregistry_org\x18\x02 \x01(\tR\vregistryOrg\"T\n" +
	"\x1aListPolicyBuildersResponse\x126\n" +
	"\aresults\x18\x01 \x03(\v2\x1c.aserto.api.v1.PolicyBuilderR\aresults\"2\n" +
	"\x1aDeletePolicyBuilderRequest\x12\x14\n" +
	"\x02id\x18\x01 \x01(\tB\x04\x90\xb5\x18\tR\x02id\"M\n" +
	"\x1bDeletePolicyBuilderResponse\x12.\n" +
	"\x06result\x18\x01 \x01(\v2\x16.google.protobuf.EmptyR\x06result2\xe3\a\n" +
	"\rPolicyBuilder\x12\xb8\x02\n" +
	"\x12ListPolicyBuilders\x12:.aserto.tenant.policy_builder.v1.ListPolicyBuildersRequest\x1a;.aserto.tenant.policy_builder.v1.ListPolicyBuildersResponse\"\xa8\x01\x92A\x80\x01\n" +
	"\x0epolicy builder\x12\"List policy builders (DEPRECATED).\x1a\x14List policy builders*\x1bpolicy.list_policy_buildersb\x17\n" +
	"\a\n" +
	"\x03JWT\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02\x1e\x12\x1c/api/v1/tenant/policybuilder\x12\xcf\x02\n" +
	"\x13CreatePolicyBuilder\x12;.aserto.tenant.policy_builder.v1.CreatePolicyBuilderRequest\x1a<.aserto.tenant.policy_builder.v1.CreatePolicyBuilderResponse\"\xbc\x01\x92A\x84\x01\n" +
	"\x0epolicy builder\x12#Create policy builder (DEPRECATED).\x1a\x16Create policy builder.*\x1cpolicy.create_policy_builderb\x17\n" +
	"\a\n" +
	"\x03JWT\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02.:\x0epolicy_builder\"\x1c/api/v1/tenant/policybuilder\x12\xc4\x02\n" +
	"\x13DeletePolicyBuilder\x12;.aserto.tenant.policy_builder.v1.DeletePolicyBuilderRequest\x1a<.aserto.tenant.policy_builder.v1.DeletePolicyBuilderResponse\"\xb1\x01\x92A\x84\x01\n" +
	"\x0epolicy builder\x12#Remove policy builder (DEPRECATED).\x1a\x16Remove policy builder.*\x1cpolicy.delete_policy_builderb\x17\n" +
	"\a\n" +
	"\x03JWT\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02#*!/api/v1/tenant/policybuilder/{id}B\xd4\x01\x92A\x82\x01*\x01\x022\x10application/json:\x10application/jsonZ@\n" +
	"\x1a\n" +
	"\x03JWT\x12\x13\b\x02\x1a\rauthorization \x02\n" +
	"\"\n" +
	"\bTenantID\x12\x16\b\x02\x1a\x10aserto-tenant-id \x02b\x17\n" +
	"\a\n" +
	"\x03JWT\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00ZLgithub.com/aserto-dev/go-grpc/aserto/tenant/policy_builder/v1;policy_builderb\x06proto3"

var (
	file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescOnce sync.Once
	file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescData []byte
)

func file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescGZIP() []byte {
	file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescOnce.Do(func() {
		file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDesc), len(file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDesc)))
	})
	return file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDescData
}

var file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_aserto_tenant_policy_builder_v1_policy_builder_proto_goTypes = []any{
	(*CreatePolicyBuilderRequest)(nil),  // 0: aserto.tenant.policy_builder.v1.CreatePolicyBuilderRequest
	(*CreatePolicyBuilderResponse)(nil), // 1: aserto.tenant.policy_builder.v1.CreatePolicyBuilderResponse
	(*ListPolicyBuildersRequest)(nil),   // 2: aserto.tenant.policy_builder.v1.ListPolicyBuildersRequest
	(*ListPolicyBuildersResponse)(nil),  // 3: aserto.tenant.policy_builder.v1.ListPolicyBuildersResponse
	(*DeletePolicyBuilderRequest)(nil),  // 4: aserto.tenant.policy_builder.v1.DeletePolicyBuilderRequest
	(*DeletePolicyBuilderResponse)(nil), // 5: aserto.tenant.policy_builder.v1.DeletePolicyBuilderResponse
	(*v1.PolicyBuilder)(nil),            // 6: aserto.api.v1.PolicyBuilder
	(*emptypb.Empty)(nil),               // 7: google.protobuf.Empty
}
var file_aserto_tenant_policy_builder_v1_policy_builder_proto_depIdxs = []int32{
	6, // 0: aserto.tenant.policy_builder.v1.CreatePolicyBuilderRequest.policy_builder:type_name -> aserto.api.v1.PolicyBuilder
	6, // 1: aserto.tenant.policy_builder.v1.ListPolicyBuildersResponse.results:type_name -> aserto.api.v1.PolicyBuilder
	7, // 2: aserto.tenant.policy_builder.v1.DeletePolicyBuilderResponse.result:type_name -> google.protobuf.Empty
	2, // 3: aserto.tenant.policy_builder.v1.PolicyBuilder.ListPolicyBuilders:input_type -> aserto.tenant.policy_builder.v1.ListPolicyBuildersRequest
	0, // 4: aserto.tenant.policy_builder.v1.PolicyBuilder.CreatePolicyBuilder:input_type -> aserto.tenant.policy_builder.v1.CreatePolicyBuilderRequest
	4, // 5: aserto.tenant.policy_builder.v1.PolicyBuilder.DeletePolicyBuilder:input_type -> aserto.tenant.policy_builder.v1.DeletePolicyBuilderRequest
	3, // 6: aserto.tenant.policy_builder.v1.PolicyBuilder.ListPolicyBuilders:output_type -> aserto.tenant.policy_builder.v1.ListPolicyBuildersResponse
	1, // 7: aserto.tenant.policy_builder.v1.PolicyBuilder.CreatePolicyBuilder:output_type -> aserto.tenant.policy_builder.v1.CreatePolicyBuilderResponse
	5, // 8: aserto.tenant.policy_builder.v1.PolicyBuilder.DeletePolicyBuilder:output_type -> aserto.tenant.policy_builder.v1.DeletePolicyBuilderResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_aserto_tenant_policy_builder_v1_policy_builder_proto_init() }
func file_aserto_tenant_policy_builder_v1_policy_builder_proto_init() {
	if File_aserto_tenant_policy_builder_v1_policy_builder_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDesc), len(file_aserto_tenant_policy_builder_v1_policy_builder_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_tenant_policy_builder_v1_policy_builder_proto_goTypes,
		DependencyIndexes: file_aserto_tenant_policy_builder_v1_policy_builder_proto_depIdxs,
		MessageInfos:      file_aserto_tenant_policy_builder_v1_policy_builder_proto_msgTypes,
	}.Build()
	File_aserto_tenant_policy_builder_v1_policy_builder_proto = out.File
	file_aserto_tenant_policy_builder_v1_policy_builder_proto_goTypes = nil
	file_aserto_tenant_policy_builder_v1_policy_builder_proto_depIdxs = nil
}
