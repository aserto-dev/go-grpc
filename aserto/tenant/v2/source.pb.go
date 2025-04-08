// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/tenant/v2/source.proto

package tenant

import (
	v1 "github.com/aserto-dev/go-grpc/aserto/api/v1"
	v2 "github.com/aserto-dev/go-grpc/aserto/api/v2"
	_ "github.com/aserto-dev/go-grpc/aserto/options/v1"
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

type CreateSourceRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Source           *v2.Source             `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	WorkflowFileName string                 `protobuf:"bytes,2,opt,name=workflow_file_name,json=workflowFileName,proto3" json:"workflow_file_name,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreateSourceRequest) Reset() {
	*x = CreateSourceRequest{}
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSourceRequest) ProtoMessage() {}

func (x *CreateSourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSourceRequest.ProtoReflect.Descriptor instead.
func (*CreateSourceRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_source_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSourceRequest) GetSource() *v2.Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *CreateSourceRequest) GetWorkflowFileName() string {
	if x != nil {
		return x.WorkflowFileName
	}
	return ""
}

type CreateSourceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Source        *v2.Source             `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSourceResponse) Reset() {
	*x = CreateSourceResponse{}
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSourceResponse) ProtoMessage() {}

func (x *CreateSourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSourceResponse.ProtoReflect.Descriptor instead.
func (*CreateSourceResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_source_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSourceResponse) GetSource() *v2.Source {
	if x != nil {
		return x.Source
	}
	return nil
}

type GetSourceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PolicyId      string                 `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSourceRequest) Reset() {
	*x = GetSourceRequest{}
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSourceRequest) ProtoMessage() {}

func (x *GetSourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSourceRequest.ProtoReflect.Descriptor instead.
func (*GetSourceRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_source_proto_rawDescGZIP(), []int{2}
}

func (x *GetSourceRequest) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

type GetSourceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Source        *v2.Source             `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSourceResponse) Reset() {
	*x = GetSourceResponse{}
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSourceResponse) ProtoMessage() {}

func (x *GetSourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSourceResponse.ProtoReflect.Descriptor instead.
func (*GetSourceResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_source_proto_rawDescGZIP(), []int{3}
}

func (x *GetSourceResponse) GetSource() *v2.Source {
	if x != nil {
		return x.Source
	}
	return nil
}

type DeleteSourceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PolicyId      string                 `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteSourceRequest) Reset() {
	*x = DeleteSourceRequest{}
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSourceRequest) ProtoMessage() {}

func (x *DeleteSourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSourceRequest.ProtoReflect.Descriptor instead.
func (*DeleteSourceRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_source_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSourceRequest) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

type DeleteSourceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteSourceResponse) Reset() {
	*x = DeleteSourceResponse{}
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSourceResponse) ProtoMessage() {}

func (x *DeleteSourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSourceResponse.ProtoReflect.Descriptor instead.
func (*DeleteSourceResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_source_proto_rawDescGZIP(), []int{5}
}

type UpdateSourceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Source        *v2.Source             `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Fields        *v1.Fields             `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateSourceRequest) Reset() {
	*x = UpdateSourceRequest{}
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSourceRequest) ProtoMessage() {}

func (x *UpdateSourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSourceRequest.ProtoReflect.Descriptor instead.
func (*UpdateSourceRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_source_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateSourceRequest) GetSource() *v2.Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *UpdateSourceRequest) GetFields() *v1.Fields {
	if x != nil {
		return x.Fields
	}
	return nil
}

type UpdateSourceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Source        *v2.Source             `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateSourceResponse) Reset() {
	*x = UpdateSourceResponse{}
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSourceResponse) ProtoMessage() {}

func (x *UpdateSourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_source_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSourceResponse.ProtoReflect.Descriptor instead.
func (*UpdateSourceResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_source_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateSourceResponse) GetSource() *v2.Source {
	if x != nil {
		return x.Source
	}
	return nil
}

var File_aserto_tenant_v2_source_proto protoreflect.FileDescriptor

const file_aserto_tenant_v2_source_proto_rawDesc = "" +
	"\n" +
	"\x1daserto/tenant/v2/source.proto\x12\x10aserto.tenant.v2\x1a\x1aaserto/api/v1/fields.proto\x1a\x1aaserto/api/v2/source.proto\x1a\x1baserto/options/v1/ids.proto\"r\n" +
	"\x13CreateSourceRequest\x12-\n" +
	"\x06source\x18\x01 \x01(\v2\x15.aserto.api.v2.SourceR\x06source\x12,\n" +
	"\x12workflow_file_name\x18\x02 \x01(\tR\x10workflowFileName\"E\n" +
	"\x14CreateSourceResponse\x12-\n" +
	"\x06source\x18\x01 \x01(\v2\x15.aserto.api.v2.SourceR\x06source\"/\n" +
	"\x10GetSourceRequest\x12\x1b\n" +
	"\tpolicy_id\x18\x01 \x01(\tR\bpolicyId\"B\n" +
	"\x11GetSourceResponse\x12-\n" +
	"\x06source\x18\x01 \x01(\v2\x15.aserto.api.v2.SourceR\x06source\"8\n" +
	"\x13DeleteSourceRequest\x12!\n" +
	"\tpolicy_id\x18\x01 \x01(\tB\x04\x90\xb5\x18\x04R\bpolicyId\"\x16\n" +
	"\x14DeleteSourceResponse\"s\n" +
	"\x13UpdateSourceRequest\x12-\n" +
	"\x06source\x18\x01 \x01(\v2\x15.aserto.api.v2.SourceR\x06source\x12-\n" +
	"\x06fields\x18\x02 \x01(\v2\x15.aserto.api.v1.FieldsR\x06fields\"E\n" +
	"\x14UpdateSourceResponse\x12-\n" +
	"\x06source\x18\x01 \x01(\v2\x15.aserto.api.v2.SourceR\x06source2\xfb\x02\n" +
	"\x06Source\x12T\n" +
	"\tGetSource\x12\".aserto.tenant.v2.GetSourceRequest\x1a#.aserto.tenant.v2.GetSourceResponse\x12]\n" +
	"\fCreateSource\x12%.aserto.tenant.v2.CreateSourceRequest\x1a&.aserto.tenant.v2.CreateSourceResponse\x12]\n" +
	"\fDeleteSource\x12%.aserto.tenant.v2.DeleteSourceRequest\x1a&.aserto.tenant.v2.DeleteSourceResponse\x12]\n" +
	"\fUpdateSource\x12%.aserto.tenant.v2.UpdateSourceRequest\x1a&.aserto.tenant.v2.UpdateSourceResponseB7Z5github.com/aserto-dev/go-grpc/aserto/tenant/v2;tenantb\x06proto3"

var (
	file_aserto_tenant_v2_source_proto_rawDescOnce sync.Once
	file_aserto_tenant_v2_source_proto_rawDescData []byte
)

func file_aserto_tenant_v2_source_proto_rawDescGZIP() []byte {
	file_aserto_tenant_v2_source_proto_rawDescOnce.Do(func() {
		file_aserto_tenant_v2_source_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_tenant_v2_source_proto_rawDesc), len(file_aserto_tenant_v2_source_proto_rawDesc)))
	})
	return file_aserto_tenant_v2_source_proto_rawDescData
}

var file_aserto_tenant_v2_source_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_aserto_tenant_v2_source_proto_goTypes = []any{
	(*CreateSourceRequest)(nil),  // 0: aserto.tenant.v2.CreateSourceRequest
	(*CreateSourceResponse)(nil), // 1: aserto.tenant.v2.CreateSourceResponse
	(*GetSourceRequest)(nil),     // 2: aserto.tenant.v2.GetSourceRequest
	(*GetSourceResponse)(nil),    // 3: aserto.tenant.v2.GetSourceResponse
	(*DeleteSourceRequest)(nil),  // 4: aserto.tenant.v2.DeleteSourceRequest
	(*DeleteSourceResponse)(nil), // 5: aserto.tenant.v2.DeleteSourceResponse
	(*UpdateSourceRequest)(nil),  // 6: aserto.tenant.v2.UpdateSourceRequest
	(*UpdateSourceResponse)(nil), // 7: aserto.tenant.v2.UpdateSourceResponse
	(*v2.Source)(nil),            // 8: aserto.api.v2.Source
	(*v1.Fields)(nil),            // 9: aserto.api.v1.Fields
}
var file_aserto_tenant_v2_source_proto_depIdxs = []int32{
	8,  // 0: aserto.tenant.v2.CreateSourceRequest.source:type_name -> aserto.api.v2.Source
	8,  // 1: aserto.tenant.v2.CreateSourceResponse.source:type_name -> aserto.api.v2.Source
	8,  // 2: aserto.tenant.v2.GetSourceResponse.source:type_name -> aserto.api.v2.Source
	8,  // 3: aserto.tenant.v2.UpdateSourceRequest.source:type_name -> aserto.api.v2.Source
	9,  // 4: aserto.tenant.v2.UpdateSourceRequest.fields:type_name -> aserto.api.v1.Fields
	8,  // 5: aserto.tenant.v2.UpdateSourceResponse.source:type_name -> aserto.api.v2.Source
	2,  // 6: aserto.tenant.v2.Source.GetSource:input_type -> aserto.tenant.v2.GetSourceRequest
	0,  // 7: aserto.tenant.v2.Source.CreateSource:input_type -> aserto.tenant.v2.CreateSourceRequest
	4,  // 8: aserto.tenant.v2.Source.DeleteSource:input_type -> aserto.tenant.v2.DeleteSourceRequest
	6,  // 9: aserto.tenant.v2.Source.UpdateSource:input_type -> aserto.tenant.v2.UpdateSourceRequest
	3,  // 10: aserto.tenant.v2.Source.GetSource:output_type -> aserto.tenant.v2.GetSourceResponse
	1,  // 11: aserto.tenant.v2.Source.CreateSource:output_type -> aserto.tenant.v2.CreateSourceResponse
	5,  // 12: aserto.tenant.v2.Source.DeleteSource:output_type -> aserto.tenant.v2.DeleteSourceResponse
	7,  // 13: aserto.tenant.v2.Source.UpdateSource:output_type -> aserto.tenant.v2.UpdateSourceResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_aserto_tenant_v2_source_proto_init() }
func file_aserto_tenant_v2_source_proto_init() {
	if File_aserto_tenant_v2_source_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_tenant_v2_source_proto_rawDesc), len(file_aserto_tenant_v2_source_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_tenant_v2_source_proto_goTypes,
		DependencyIndexes: file_aserto_tenant_v2_source_proto_depIdxs,
		MessageInfos:      file_aserto_tenant_v2_source_proto_msgTypes,
	}.Build()
	File_aserto_tenant_v2_source_proto = out.File
	file_aserto_tenant_v2_source_proto_goTypes = nil
	file_aserto_tenant_v2_source_proto_depIdxs = nil
}
