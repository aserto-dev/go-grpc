// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/tenant/v2/instance.proto

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

type CreateInstanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Instance      *v2.Instance           `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateInstanceRequest) Reset() {
	*x = CreateInstanceRequest{}
	mi := &file_aserto_tenant_v2_instance_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInstanceRequest) ProtoMessage() {}

func (x *CreateInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_instance_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInstanceRequest.ProtoReflect.Descriptor instead.
func (*CreateInstanceRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_instance_proto_rawDescGZIP(), []int{0}
}

func (x *CreateInstanceRequest) GetInstance() *v2.Instance {
	if x != nil {
		return x.Instance
	}
	return nil
}

type CreateInstanceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Instance      *v2.Instance           `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateInstanceResponse) Reset() {
	*x = CreateInstanceResponse{}
	mi := &file_aserto_tenant_v2_instance_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInstanceResponse) ProtoMessage() {}

func (x *CreateInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_instance_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInstanceResponse.ProtoReflect.Descriptor instead.
func (*CreateInstanceResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_instance_proto_rawDescGZIP(), []int{1}
}

func (x *CreateInstanceResponse) GetInstance() *v2.Instance {
	if x != nil {
		return x.Instance
	}
	return nil
}

type ListInstanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PolicyId      string                 `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	InstanceType  v2.InstanceType        `protobuf:"varint,2,opt,name=instance_type,json=instanceType,proto3,enum=aserto.api.v2.InstanceType" json:"instance_type,omitempty"`
	Page          *v1.PaginationRequest  `protobuf:"bytes,3,opt,name=page,proto3" json:"page,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListInstanceRequest) Reset() {
	*x = ListInstanceRequest{}
	mi := &file_aserto_tenant_v2_instance_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstanceRequest) ProtoMessage() {}

func (x *ListInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_instance_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstanceRequest.ProtoReflect.Descriptor instead.
func (*ListInstanceRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_instance_proto_rawDescGZIP(), []int{2}
}

func (x *ListInstanceRequest) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

func (x *ListInstanceRequest) GetInstanceType() v2.InstanceType {
	if x != nil {
		return x.InstanceType
	}
	return v2.InstanceType(0)
}

func (x *ListInstanceRequest) GetPage() *v1.PaginationRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

type ListInstanceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        []*v2.Instance         `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	Page          *v1.PaginationResponse `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListInstanceResponse) Reset() {
	*x = ListInstanceResponse{}
	mi := &file_aserto_tenant_v2_instance_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstanceResponse) ProtoMessage() {}

func (x *ListInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_instance_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstanceResponse.ProtoReflect.Descriptor instead.
func (*ListInstanceResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_instance_proto_rawDescGZIP(), []int{3}
}

func (x *ListInstanceResponse) GetResult() []*v2.Instance {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *ListInstanceResponse) GetPage() *v1.PaginationResponse {
	if x != nil {
		return x.Page
	}
	return nil
}

type DeleteInstanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PolicyId      string                 `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	Label         string                 `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteInstanceRequest) Reset() {
	*x = DeleteInstanceRequest{}
	mi := &file_aserto_tenant_v2_instance_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInstanceRequest) ProtoMessage() {}

func (x *DeleteInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_instance_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInstanceRequest.ProtoReflect.Descriptor instead.
func (*DeleteInstanceRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_instance_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteInstanceRequest) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

func (x *DeleteInstanceRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type DeleteInstanceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteInstanceResponse) Reset() {
	*x = DeleteInstanceResponse{}
	mi := &file_aserto_tenant_v2_instance_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInstanceResponse) ProtoMessage() {}

func (x *DeleteInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_instance_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInstanceResponse.ProtoReflect.Descriptor instead.
func (*DeleteInstanceResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_instance_proto_rawDescGZIP(), []int{5}
}

type UpdateInstanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Instance      *v2.Instance           `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Fields        *v1.Fields             `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateInstanceRequest) Reset() {
	*x = UpdateInstanceRequest{}
	mi := &file_aserto_tenant_v2_instance_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInstanceRequest) ProtoMessage() {}

func (x *UpdateInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_instance_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInstanceRequest.ProtoReflect.Descriptor instead.
func (*UpdateInstanceRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_instance_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateInstanceRequest) GetInstance() *v2.Instance {
	if x != nil {
		return x.Instance
	}
	return nil
}

func (x *UpdateInstanceRequest) GetFields() *v1.Fields {
	if x != nil {
		return x.Fields
	}
	return nil
}

type UpdateInstanceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Instance      *v2.Instance           `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateInstanceResponse) Reset() {
	*x = UpdateInstanceResponse{}
	mi := &file_aserto_tenant_v2_instance_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInstanceResponse) ProtoMessage() {}

func (x *UpdateInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_v2_instance_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInstanceResponse.ProtoReflect.Descriptor instead.
func (*UpdateInstanceResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_v2_instance_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateInstanceResponse) GetInstance() *v2.Instance {
	if x != nil {
		return x.Instance
	}
	return nil
}

var File_aserto_tenant_v2_instance_proto protoreflect.FileDescriptor

const file_aserto_tenant_v2_instance_proto_rawDesc = "" +
	"\n" +
	"\x1faserto/tenant/v2/instance.proto\x12\x10aserto.tenant.v2\x1a\x1aaserto/api/v1/fields.proto\x1a\x1caserto/api/v2/instance.proto\x1a\x1baserto/options/v1/ids.proto\x1a\x1easerto/api/v1/pagination.proto\"L\n" +
	"\x15CreateInstanceRequest\x123\n" +
	"\binstance\x18\x01 \x01(\v2\x17.aserto.api.v2.InstanceR\binstance\"M\n" +
	"\x16CreateInstanceResponse\x123\n" +
	"\binstance\x18\x01 \x01(\v2\x17.aserto.api.v2.InstanceR\binstance\"\xaa\x01\n" +
	"\x13ListInstanceRequest\x12\x1b\n" +
	"\tpolicy_id\x18\x01 \x01(\tR\bpolicyId\x12@\n" +
	"\rinstance_type\x18\x02 \x01(\x0e2\x1b.aserto.api.v2.InstanceTypeR\finstanceType\x124\n" +
	"\x04page\x18\x03 \x01(\v2 .aserto.api.v1.PaginationRequestR\x04page\"~\n" +
	"\x14ListInstanceResponse\x12/\n" +
	"\x06result\x18\x01 \x03(\v2\x17.aserto.api.v2.InstanceR\x06result\x125\n" +
	"\x04page\x18\x02 \x01(\v2!.aserto.api.v1.PaginationResponseR\x04page\"P\n" +
	"\x15DeleteInstanceRequest\x12!\n" +
	"\tpolicy_id\x18\x01 \x01(\tB\x04\x90\xb5\x18\x04R\bpolicyId\x12\x14\n" +
	"\x05label\x18\x02 \x01(\tR\x05label\"\x18\n" +
	"\x16DeleteInstanceResponse\"{\n" +
	"\x15UpdateInstanceRequest\x123\n" +
	"\binstance\x18\x01 \x01(\v2\x17.aserto.api.v2.InstanceR\binstance\x12-\n" +
	"\x06fields\x18\x02 \x01(\v2\x15.aserto.api.v1.FieldsR\x06fields\"M\n" +
	"\x16UpdateInstanceResponse\x123\n" +
	"\binstance\x18\x01 \x01(\v2\x17.aserto.api.v2.InstanceR\binstance2\x98\x03\n" +
	"\bInstance\x12]\n" +
	"\fListInstance\x12%.aserto.tenant.v2.ListInstanceRequest\x1a&.aserto.tenant.v2.ListInstanceResponse\x12c\n" +
	"\x0eCreateInstance\x12'.aserto.tenant.v2.CreateInstanceRequest\x1a(.aserto.tenant.v2.CreateInstanceResponse\x12c\n" +
	"\x0eDeleteInstance\x12'.aserto.tenant.v2.DeleteInstanceRequest\x1a(.aserto.tenant.v2.DeleteInstanceResponse\x12c\n" +
	"\x0eUpdateInstance\x12'.aserto.tenant.v2.UpdateInstanceRequest\x1a(.aserto.tenant.v2.UpdateInstanceResponseB7Z5github.com/aserto-dev/go-grpc/aserto/tenant/v2;tenantb\x06proto3"

var (
	file_aserto_tenant_v2_instance_proto_rawDescOnce sync.Once
	file_aserto_tenant_v2_instance_proto_rawDescData []byte
)

func file_aserto_tenant_v2_instance_proto_rawDescGZIP() []byte {
	file_aserto_tenant_v2_instance_proto_rawDescOnce.Do(func() {
		file_aserto_tenant_v2_instance_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_tenant_v2_instance_proto_rawDesc), len(file_aserto_tenant_v2_instance_proto_rawDesc)))
	})
	return file_aserto_tenant_v2_instance_proto_rawDescData
}

var file_aserto_tenant_v2_instance_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_aserto_tenant_v2_instance_proto_goTypes = []any{
	(*CreateInstanceRequest)(nil),  // 0: aserto.tenant.v2.CreateInstanceRequest
	(*CreateInstanceResponse)(nil), // 1: aserto.tenant.v2.CreateInstanceResponse
	(*ListInstanceRequest)(nil),    // 2: aserto.tenant.v2.ListInstanceRequest
	(*ListInstanceResponse)(nil),   // 3: aserto.tenant.v2.ListInstanceResponse
	(*DeleteInstanceRequest)(nil),  // 4: aserto.tenant.v2.DeleteInstanceRequest
	(*DeleteInstanceResponse)(nil), // 5: aserto.tenant.v2.DeleteInstanceResponse
	(*UpdateInstanceRequest)(nil),  // 6: aserto.tenant.v2.UpdateInstanceRequest
	(*UpdateInstanceResponse)(nil), // 7: aserto.tenant.v2.UpdateInstanceResponse
	(*v2.Instance)(nil),            // 8: aserto.api.v2.Instance
	(v2.InstanceType)(0),           // 9: aserto.api.v2.InstanceType
	(*v1.PaginationRequest)(nil),   // 10: aserto.api.v1.PaginationRequest
	(*v1.PaginationResponse)(nil),  // 11: aserto.api.v1.PaginationResponse
	(*v1.Fields)(nil),              // 12: aserto.api.v1.Fields
}
var file_aserto_tenant_v2_instance_proto_depIdxs = []int32{
	8,  // 0: aserto.tenant.v2.CreateInstanceRequest.instance:type_name -> aserto.api.v2.Instance
	8,  // 1: aserto.tenant.v2.CreateInstanceResponse.instance:type_name -> aserto.api.v2.Instance
	9,  // 2: aserto.tenant.v2.ListInstanceRequest.instance_type:type_name -> aserto.api.v2.InstanceType
	10, // 3: aserto.tenant.v2.ListInstanceRequest.page:type_name -> aserto.api.v1.PaginationRequest
	8,  // 4: aserto.tenant.v2.ListInstanceResponse.result:type_name -> aserto.api.v2.Instance
	11, // 5: aserto.tenant.v2.ListInstanceResponse.page:type_name -> aserto.api.v1.PaginationResponse
	8,  // 6: aserto.tenant.v2.UpdateInstanceRequest.instance:type_name -> aserto.api.v2.Instance
	12, // 7: aserto.tenant.v2.UpdateInstanceRequest.fields:type_name -> aserto.api.v1.Fields
	8,  // 8: aserto.tenant.v2.UpdateInstanceResponse.instance:type_name -> aserto.api.v2.Instance
	2,  // 9: aserto.tenant.v2.Instance.ListInstance:input_type -> aserto.tenant.v2.ListInstanceRequest
	0,  // 10: aserto.tenant.v2.Instance.CreateInstance:input_type -> aserto.tenant.v2.CreateInstanceRequest
	4,  // 11: aserto.tenant.v2.Instance.DeleteInstance:input_type -> aserto.tenant.v2.DeleteInstanceRequest
	6,  // 12: aserto.tenant.v2.Instance.UpdateInstance:input_type -> aserto.tenant.v2.UpdateInstanceRequest
	3,  // 13: aserto.tenant.v2.Instance.ListInstance:output_type -> aserto.tenant.v2.ListInstanceResponse
	1,  // 14: aserto.tenant.v2.Instance.CreateInstance:output_type -> aserto.tenant.v2.CreateInstanceResponse
	5,  // 15: aserto.tenant.v2.Instance.DeleteInstance:output_type -> aserto.tenant.v2.DeleteInstanceResponse
	7,  // 16: aserto.tenant.v2.Instance.UpdateInstance:output_type -> aserto.tenant.v2.UpdateInstanceResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_aserto_tenant_v2_instance_proto_init() }
func file_aserto_tenant_v2_instance_proto_init() {
	if File_aserto_tenant_v2_instance_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_tenant_v2_instance_proto_rawDesc), len(file_aserto_tenant_v2_instance_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_tenant_v2_instance_proto_goTypes,
		DependencyIndexes: file_aserto_tenant_v2_instance_proto_depIdxs,
		MessageInfos:      file_aserto_tenant_v2_instance_proto_msgTypes,
	}.Build()
	File_aserto_tenant_v2_instance_proto = out.File
	file_aserto_tenant_v2_instance_proto_goTypes = nil
	file_aserto_tenant_v2_instance_proto_depIdxs = nil
}
