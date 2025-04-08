// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/idpplugin/v1/idpplugin.proto

package idpplugin

import (
	v11 "github.com/aserto-dev/go-grpc/aserto/api/v1"
	v1 "github.com/aserto-dev/go-grpc/aserto/common/info/v1"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type OperationType int32

const (
	OperationType_OPERATION_TYPE_UNKNOWN OperationType = 0
	OperationType_OPERATION_TYPE_IMPORT  OperationType = 1
	OperationType_OPERATION_TYPE_EXPORT  OperationType = 2
	OperationType_OPERATION_TYPE_DELETE  OperationType = 3
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "OPERATION_TYPE_UNKNOWN",
		1: "OPERATION_TYPE_IMPORT",
		2: "OPERATION_TYPE_EXPORT",
		3: "OPERATION_TYPE_DELETE",
	}
	OperationType_value = map[string]int32{
		"OPERATION_TYPE_UNKNOWN": 0,
		"OPERATION_TYPE_IMPORT":  1,
		"OPERATION_TYPE_EXPORT":  2,
		"OPERATION_TYPE_DELETE":  3,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_idpplugin_v1_idpplugin_proto_enumTypes[0].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_aserto_idpplugin_v1_idpplugin_proto_enumTypes[0]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_aserto_idpplugin_v1_idpplugin_proto_rawDescGZIP(), []int{0}
}

type InfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_aserto_idpplugin_v1_idpplugin_proto_rawDescGZIP(), []int{0}
}

type InfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Build         *v1.BuildInfo          `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Configs       []*v11.ConfigElement   `protobuf:"bytes,3,rep,name=configs,proto3" json:"configs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_aserto_idpplugin_v1_idpplugin_proto_rawDescGZIP(), []int{1}
}

func (x *InfoResponse) GetBuild() *v1.BuildInfo {
	if x != nil {
		return x.Build
	}
	return nil
}

func (x *InfoResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *InfoResponse) GetConfigs() []*v11.ConfigElement {
	if x != nil {
		return x.Configs
	}
	return nil
}

type ImportRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*ImportRequest_Config
	//	*ImportRequest_User
	//	*ImportRequest_UserExt
	Data          isImportRequest_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImportRequest) Reset() {
	*x = ImportRequest{}
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportRequest) ProtoMessage() {}

func (x *ImportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportRequest.ProtoReflect.Descriptor instead.
func (*ImportRequest) Descriptor() ([]byte, []int) {
	return file_aserto_idpplugin_v1_idpplugin_proto_rawDescGZIP(), []int{2}
}

func (x *ImportRequest) GetData() isImportRequest_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ImportRequest) GetConfig() *structpb.Struct {
	if x != nil {
		if x, ok := x.Data.(*ImportRequest_Config); ok {
			return x.Config
		}
	}
	return nil
}

func (x *ImportRequest) GetUser() *v11.User {
	if x != nil {
		if x, ok := x.Data.(*ImportRequest_User); ok {
			return x.User
		}
	}
	return nil
}

func (x *ImportRequest) GetUserExt() *v11.UserExt {
	if x != nil {
		if x, ok := x.Data.(*ImportRequest_UserExt); ok {
			return x.UserExt
		}
	}
	return nil
}

type isImportRequest_Data interface {
	isImportRequest_Data()
}

type ImportRequest_Config struct {
	Config *structpb.Struct `protobuf:"bytes,1,opt,name=config,proto3,oneof"`
}

type ImportRequest_User struct {
	User *v11.User `protobuf:"bytes,2,opt,name=user,proto3,oneof"`
}

type ImportRequest_UserExt struct {
	UserExt *v11.UserExt `protobuf:"bytes,3,opt,name=user_ext,json=userExt,proto3,oneof"`
}

func (*ImportRequest_Config) isImportRequest_Data() {}

func (*ImportRequest_User) isImportRequest_Data() {}

func (*ImportRequest_UserExt) isImportRequest_Data() {}

type ImportResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Error         *status.Status         `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Stats         *v11.UserProcessStats  `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImportResponse) Reset() {
	*x = ImportResponse{}
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportResponse) ProtoMessage() {}

func (x *ImportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportResponse.ProtoReflect.Descriptor instead.
func (*ImportResponse) Descriptor() ([]byte, []int) {
	return file_aserto_idpplugin_v1_idpplugin_proto_rawDescGZIP(), []int{3}
}

func (x *ImportResponse) GetError() *status.Status {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ImportResponse) GetStats() *v11.UserProcessStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type ExportRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Config        *structpb.Struct       `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportRequest) Reset() {
	*x = ExportRequest{}
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportRequest) ProtoMessage() {}

func (x *ExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportRequest.ProtoReflect.Descriptor instead.
func (*ExportRequest) Descriptor() ([]byte, []int) {
	return file_aserto_idpplugin_v1_idpplugin_proto_rawDescGZIP(), []int{4}
}

func (x *ExportRequest) GetConfig() *structpb.Struct {
	if x != nil {
		return x.Config
	}
	return nil
}

type ExportResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*ExportResponse_User
	//	*ExportResponse_UserExt
	//	*ExportResponse_Error
	Data          isExportResponse_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportResponse) Reset() {
	*x = ExportResponse{}
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportResponse) ProtoMessage() {}

func (x *ExportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportResponse.ProtoReflect.Descriptor instead.
func (*ExportResponse) Descriptor() ([]byte, []int) {
	return file_aserto_idpplugin_v1_idpplugin_proto_rawDescGZIP(), []int{5}
}

func (x *ExportResponse) GetData() isExportResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ExportResponse) GetUser() *v11.User {
	if x != nil {
		if x, ok := x.Data.(*ExportResponse_User); ok {
			return x.User
		}
	}
	return nil
}

func (x *ExportResponse) GetUserExt() *v11.UserExt {
	if x != nil {
		if x, ok := x.Data.(*ExportResponse_UserExt); ok {
			return x.UserExt
		}
	}
	return nil
}

func (x *ExportResponse) GetError() *status.Status {
	if x != nil {
		if x, ok := x.Data.(*ExportResponse_Error); ok {
			return x.Error
		}
	}
	return nil
}

type isExportResponse_Data interface {
	isExportResponse_Data()
}

type ExportResponse_User struct {
	User *v11.User `protobuf:"bytes,1,opt,name=user,proto3,oneof"`
}

type ExportResponse_UserExt struct {
	UserExt *v11.UserExt `protobuf:"bytes,2,opt,name=user_ext,json=userExt,proto3,oneof"`
}

type ExportResponse_Error struct {
	Error *status.Status `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*ExportResponse_User) isExportResponse_Data() {}

func (*ExportResponse_UserExt) isExportResponse_Data() {}

func (*ExportResponse_Error) isExportResponse_Data() {}

type DeleteRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*DeleteRequest_Config
	//	*DeleteRequest_UserId
	Data          isDeleteRequest_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_aserto_idpplugin_v1_idpplugin_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRequest) GetData() isDeleteRequest_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DeleteRequest) GetConfig() *structpb.Struct {
	if x != nil {
		if x, ok := x.Data.(*DeleteRequest_Config); ok {
			return x.Config
		}
	}
	return nil
}

func (x *DeleteRequest) GetUserId() string {
	if x != nil {
		if x, ok := x.Data.(*DeleteRequest_UserId); ok {
			return x.UserId
		}
	}
	return ""
}

type isDeleteRequest_Data interface {
	isDeleteRequest_Data()
}

type DeleteRequest_Config struct {
	Config *structpb.Struct `protobuf:"bytes,1,opt,name=config,proto3,oneof"`
}

type DeleteRequest_UserId struct {
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3,oneof"`
}

func (*DeleteRequest_Config) isDeleteRequest_Data() {}

func (*DeleteRequest_UserId) isDeleteRequest_Data() {}

type DeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Error         *status.Status         `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_aserto_idpplugin_v1_idpplugin_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteResponse) GetError() *status.Status {
	if x != nil {
		return x.Error
	}
	return nil
}

type ValidateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Config        *structpb.Struct       `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	OpType        OperationType          `protobuf:"varint,2,opt,name=op_type,json=opType,proto3,enum=aserto.idpplugin.v1.OperationType" json:"op_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return file_aserto_idpplugin_v1_idpplugin_proto_rawDescGZIP(), []int{8}
}

func (x *ValidateRequest) GetConfig() *structpb.Struct {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ValidateRequest) GetOpType() OperationType {
	if x != nil {
		return x.OpType
	}
	return OperationType_OPERATION_TYPE_UNKNOWN
}

type ValidateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        *status.Status         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateResponse) Reset() {
	*x = ValidateResponse{}
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateResponse) ProtoMessage() {}

func (x *ValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateResponse.ProtoReflect.Descriptor instead.
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return file_aserto_idpplugin_v1_idpplugin_proto_rawDescGZIP(), []int{9}
}

func (x *ValidateResponse) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_aserto_idpplugin_v1_idpplugin_proto protoreflect.FileDescriptor

const file_aserto_idpplugin_v1_idpplugin_proto_rawDesc = "" +
	"\n" +
	"#aserto/idpplugin/v1/idpplugin.proto\x12\x13aserto.idpplugin.v1\x1a\x1cgoogle/protobuf/struct.proto\x1a\x17google/rpc/status.proto\x1a\x1easerto/api/v1/connection.proto\x1a aserto/common/info/v1/info.proto\x1a\x18aserto/api/v1/user.proto\"\r\n" +
	"\vInfoRequest\"\xa0\x01\n" +
	"\fInfoResponse\x126\n" +
	"\x05build\x18\x01 \x01(\v2 .aserto.common.info.v1.BuildInfoR\x05build\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x126\n" +
	"\aconfigs\x18\x03 \x03(\v2\x1c.aserto.api.v1.ConfigElementR\aconfigs\"\xaa\x01\n" +
	"\rImportRequest\x121\n" +
	"\x06config\x18\x01 \x01(\v2\x17.google.protobuf.StructH\x00R\x06config\x12)\n" +
	"\x04user\x18\x02 \x01(\v2\x13.aserto.api.v1.UserH\x00R\x04user\x123\n" +
	"\buser_ext\x18\x03 \x01(\v2\x16.aserto.api.v1.UserExtH\x00R\auserExtB\x06\n" +
	"\x04data\"q\n" +
	"\x0eImportResponse\x12(\n" +
	"\x05error\x18\x01 \x01(\v2\x12.google.rpc.StatusR\x05error\x125\n" +
	"\x05stats\x18\x02 \x01(\v2\x1f.aserto.api.v1.UserProcessStatsR\x05stats\"@\n" +
	"\rExportRequest\x12/\n" +
	"\x06config\x18\x01 \x01(\v2\x17.google.protobuf.StructR\x06config\"\xa4\x01\n" +
	"\x0eExportResponse\x12)\n" +
	"\x04user\x18\x01 \x01(\v2\x13.aserto.api.v1.UserH\x00R\x04user\x123\n" +
	"\buser_ext\x18\x02 \x01(\v2\x16.aserto.api.v1.UserExtH\x00R\auserExt\x12*\n" +
	"\x05error\x18\x03 \x01(\v2\x12.google.rpc.StatusH\x00R\x05errorB\x06\n" +
	"\x04data\"e\n" +
	"\rDeleteRequest\x121\n" +
	"\x06config\x18\x01 \x01(\v2\x17.google.protobuf.StructH\x00R\x06config\x12\x19\n" +
	"\auser_id\x18\x02 \x01(\tH\x00R\x06userIdB\x06\n" +
	"\x04data\":\n" +
	"\x0eDeleteResponse\x12(\n" +
	"\x05error\x18\x01 \x01(\v2\x12.google.rpc.StatusR\x05error\"\x7f\n" +
	"\x0fValidateRequest\x12/\n" +
	"\x06config\x18\x01 \x01(\v2\x17.google.protobuf.StructR\x06config\x12;\n" +
	"\aop_type\x18\x02 \x01(\x0e2\".aserto.idpplugin.v1.OperationTypeR\x06opType\">\n" +
	"\x10ValidateResponse\x12*\n" +
	"\x06status\x18\x01 \x01(\v2\x12.google.rpc.StatusR\x06status*|\n" +
	"\rOperationType\x12\x1a\n" +
	"\x16OPERATION_TYPE_UNKNOWN\x10\x00\x12\x19\n" +
	"\x15OPERATION_TYPE_IMPORT\x10\x01\x12\x19\n" +
	"\x15OPERATION_TYPE_EXPORT\x10\x02\x12\x19\n" +
	"\x15OPERATION_TYPE_DELETE\x10\x032\xb1\x03\n" +
	"\x06Plugin\x12K\n" +
	"\x04Info\x12 .aserto.idpplugin.v1.InfoRequest\x1a!.aserto.idpplugin.v1.InfoResponse\x12U\n" +
	"\x06Import\x12\".aserto.idpplugin.v1.ImportRequest\x1a#.aserto.idpplugin.v1.ImportResponse(\x010\x01\x12S\n" +
	"\x06Export\x12\".aserto.idpplugin.v1.ExportRequest\x1a#.aserto.idpplugin.v1.ExportResponse0\x01\x12U\n" +
	"\x06Delete\x12\".aserto.idpplugin.v1.DeleteRequest\x1a#.aserto.idpplugin.v1.DeleteResponse(\x010\x01\x12W\n" +
	"\bValidate\x12$.aserto.idpplugin.v1.ValidateRequest\x1a%.aserto.idpplugin.v1.ValidateResponseB=Z;github.com/aserto-dev/go-grpc/aserto/idpplugin/v1;idppluginb\x06proto3"

var (
	file_aserto_idpplugin_v1_idpplugin_proto_rawDescOnce sync.Once
	file_aserto_idpplugin_v1_idpplugin_proto_rawDescData []byte
)

func file_aserto_idpplugin_v1_idpplugin_proto_rawDescGZIP() []byte {
	file_aserto_idpplugin_v1_idpplugin_proto_rawDescOnce.Do(func() {
		file_aserto_idpplugin_v1_idpplugin_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_idpplugin_v1_idpplugin_proto_rawDesc), len(file_aserto_idpplugin_v1_idpplugin_proto_rawDesc)))
	})
	return file_aserto_idpplugin_v1_idpplugin_proto_rawDescData
}

var file_aserto_idpplugin_v1_idpplugin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_idpplugin_v1_idpplugin_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_aserto_idpplugin_v1_idpplugin_proto_goTypes = []any{
	(OperationType)(0),           // 0: aserto.idpplugin.v1.OperationType
	(*InfoRequest)(nil),          // 1: aserto.idpplugin.v1.InfoRequest
	(*InfoResponse)(nil),         // 2: aserto.idpplugin.v1.InfoResponse
	(*ImportRequest)(nil),        // 3: aserto.idpplugin.v1.ImportRequest
	(*ImportResponse)(nil),       // 4: aserto.idpplugin.v1.ImportResponse
	(*ExportRequest)(nil),        // 5: aserto.idpplugin.v1.ExportRequest
	(*ExportResponse)(nil),       // 6: aserto.idpplugin.v1.ExportResponse
	(*DeleteRequest)(nil),        // 7: aserto.idpplugin.v1.DeleteRequest
	(*DeleteResponse)(nil),       // 8: aserto.idpplugin.v1.DeleteResponse
	(*ValidateRequest)(nil),      // 9: aserto.idpplugin.v1.ValidateRequest
	(*ValidateResponse)(nil),     // 10: aserto.idpplugin.v1.ValidateResponse
	(*v1.BuildInfo)(nil),         // 11: aserto.common.info.v1.BuildInfo
	(*v11.ConfigElement)(nil),    // 12: aserto.api.v1.ConfigElement
	(*structpb.Struct)(nil),      // 13: google.protobuf.Struct
	(*v11.User)(nil),             // 14: aserto.api.v1.User
	(*v11.UserExt)(nil),          // 15: aserto.api.v1.UserExt
	(*status.Status)(nil),        // 16: google.rpc.Status
	(*v11.UserProcessStats)(nil), // 17: aserto.api.v1.UserProcessStats
}
var file_aserto_idpplugin_v1_idpplugin_proto_depIdxs = []int32{
	11, // 0: aserto.idpplugin.v1.InfoResponse.build:type_name -> aserto.common.info.v1.BuildInfo
	12, // 1: aserto.idpplugin.v1.InfoResponse.configs:type_name -> aserto.api.v1.ConfigElement
	13, // 2: aserto.idpplugin.v1.ImportRequest.config:type_name -> google.protobuf.Struct
	14, // 3: aserto.idpplugin.v1.ImportRequest.user:type_name -> aserto.api.v1.User
	15, // 4: aserto.idpplugin.v1.ImportRequest.user_ext:type_name -> aserto.api.v1.UserExt
	16, // 5: aserto.idpplugin.v1.ImportResponse.error:type_name -> google.rpc.Status
	17, // 6: aserto.idpplugin.v1.ImportResponse.stats:type_name -> aserto.api.v1.UserProcessStats
	13, // 7: aserto.idpplugin.v1.ExportRequest.config:type_name -> google.protobuf.Struct
	14, // 8: aserto.idpplugin.v1.ExportResponse.user:type_name -> aserto.api.v1.User
	15, // 9: aserto.idpplugin.v1.ExportResponse.user_ext:type_name -> aserto.api.v1.UserExt
	16, // 10: aserto.idpplugin.v1.ExportResponse.error:type_name -> google.rpc.Status
	13, // 11: aserto.idpplugin.v1.DeleteRequest.config:type_name -> google.protobuf.Struct
	16, // 12: aserto.idpplugin.v1.DeleteResponse.error:type_name -> google.rpc.Status
	13, // 13: aserto.idpplugin.v1.ValidateRequest.config:type_name -> google.protobuf.Struct
	0,  // 14: aserto.idpplugin.v1.ValidateRequest.op_type:type_name -> aserto.idpplugin.v1.OperationType
	16, // 15: aserto.idpplugin.v1.ValidateResponse.status:type_name -> google.rpc.Status
	1,  // 16: aserto.idpplugin.v1.Plugin.Info:input_type -> aserto.idpplugin.v1.InfoRequest
	3,  // 17: aserto.idpplugin.v1.Plugin.Import:input_type -> aserto.idpplugin.v1.ImportRequest
	5,  // 18: aserto.idpplugin.v1.Plugin.Export:input_type -> aserto.idpplugin.v1.ExportRequest
	7,  // 19: aserto.idpplugin.v1.Plugin.Delete:input_type -> aserto.idpplugin.v1.DeleteRequest
	9,  // 20: aserto.idpplugin.v1.Plugin.Validate:input_type -> aserto.idpplugin.v1.ValidateRequest
	2,  // 21: aserto.idpplugin.v1.Plugin.Info:output_type -> aserto.idpplugin.v1.InfoResponse
	4,  // 22: aserto.idpplugin.v1.Plugin.Import:output_type -> aserto.idpplugin.v1.ImportResponse
	6,  // 23: aserto.idpplugin.v1.Plugin.Export:output_type -> aserto.idpplugin.v1.ExportResponse
	8,  // 24: aserto.idpplugin.v1.Plugin.Delete:output_type -> aserto.idpplugin.v1.DeleteResponse
	10, // 25: aserto.idpplugin.v1.Plugin.Validate:output_type -> aserto.idpplugin.v1.ValidateResponse
	21, // [21:26] is the sub-list for method output_type
	16, // [16:21] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_aserto_idpplugin_v1_idpplugin_proto_init() }
func file_aserto_idpplugin_v1_idpplugin_proto_init() {
	if File_aserto_idpplugin_v1_idpplugin_proto != nil {
		return
	}
	file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[2].OneofWrappers = []any{
		(*ImportRequest_Config)(nil),
		(*ImportRequest_User)(nil),
		(*ImportRequest_UserExt)(nil),
	}
	file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[5].OneofWrappers = []any{
		(*ExportResponse_User)(nil),
		(*ExportResponse_UserExt)(nil),
		(*ExportResponse_Error)(nil),
	}
	file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[6].OneofWrappers = []any{
		(*DeleteRequest_Config)(nil),
		(*DeleteRequest_UserId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_idpplugin_v1_idpplugin_proto_rawDesc), len(file_aserto_idpplugin_v1_idpplugin_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_idpplugin_v1_idpplugin_proto_goTypes,
		DependencyIndexes: file_aserto_idpplugin_v1_idpplugin_proto_depIdxs,
		EnumInfos:         file_aserto_idpplugin_v1_idpplugin_proto_enumTypes,
		MessageInfos:      file_aserto_idpplugin_v1_idpplugin_proto_msgTypes,
	}.Build()
	File_aserto_idpplugin_v1_idpplugin_proto = out.File
	file_aserto_idpplugin_v1_idpplugin_proto_goTypes = nil
	file_aserto_idpplugin_v1_idpplugin_proto_depIdxs = nil
}
