// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.17.3
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
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Build       *v1.BuildInfo        `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty"`
	Description string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Configs     []*v11.ConfigElement `protobuf:"bytes,3,rep,name=configs,proto3" json:"configs,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*ImportRequest_Config
	//	*ImportRequest_User
	//	*ImportRequest_UserExt
	Data isImportRequest_Data `protobuf_oneof:"data"`
}

func (x *ImportRequest) Reset() {
	*x = ImportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportRequest) ProtoMessage() {}

func (x *ImportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (m *ImportRequest) GetData() isImportRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *ImportRequest) GetConfig() *structpb.Struct {
	if x, ok := x.GetData().(*ImportRequest_Config); ok {
		return x.Config
	}
	return nil
}

func (x *ImportRequest) GetUser() *v11.User {
	if x, ok := x.GetData().(*ImportRequest_User); ok {
		return x.User
	}
	return nil
}

func (x *ImportRequest) GetUserExt() *v11.UserExt {
	if x, ok := x.GetData().(*ImportRequest_UserExt); ok {
		return x.UserExt
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *status.Status `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ImportResponse) Reset() {
	*x = ImportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportResponse) ProtoMessage() {}

func (x *ImportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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

type ExportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *structpb.Struct `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ExportRequest) Reset() {
	*x = ExportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportRequest) ProtoMessage() {}

func (x *ExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*ExportResponse_User
	//	*ExportResponse_UserExt
	//	*ExportResponse_Error
	Data isExportResponse_Data `protobuf_oneof:"data"`
}

func (x *ExportResponse) Reset() {
	*x = ExportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportResponse) ProtoMessage() {}

func (x *ExportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (m *ExportResponse) GetData() isExportResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *ExportResponse) GetUser() *v11.User {
	if x, ok := x.GetData().(*ExportResponse_User); ok {
		return x.User
	}
	return nil
}

func (x *ExportResponse) GetUserExt() *v11.UserExt {
	if x, ok := x.GetData().(*ExportResponse_UserExt); ok {
		return x.UserExt
	}
	return nil
}

func (x *ExportResponse) GetError() *status.Status {
	if x, ok := x.GetData().(*ExportResponse_Error); ok {
		return x.Error
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*DeleteRequest_Config
	//	*DeleteRequest_UserId
	Data isDeleteRequest_Data `protobuf_oneof:"data"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (m *DeleteRequest) GetData() isDeleteRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *DeleteRequest) GetConfig() *structpb.Struct {
	if x, ok := x.GetData().(*DeleteRequest_Config); ok {
		return x.Config
	}
	return nil
}

func (x *DeleteRequest) GetUserId() string {
	if x, ok := x.GetData().(*DeleteRequest_UserId); ok {
		return x.UserId
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *status.Status `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *structpb.Struct `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
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

type ValidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ValidateResponse) Reset() {
	*x = ValidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateResponse) ProtoMessage() {}

func (x *ValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_aserto_idpplugin_v1_idpplugin_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x69, 0x64, 0x70, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x70, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64,
	0x70, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a,
	0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa0, 0x01, 0x0a,
	0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a,
	0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22,
	0xaa, 0x01, 0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x29, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x33, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x74, 0x48, 0x00, 0x52, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x45, 0x78, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x0e,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x40, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xa4, 0x01, 0x0a, 0x0e, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x48, 0x00, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x45,
	0x78, 0x74, 0x48, 0x00, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x45, 0x78, 0x74, 0x12, 0x2a, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x65, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x42, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x3e, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xb1, 0x03, 0x0a, 0x06, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x12, 0x4b, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x2e, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x70, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x70, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x55, 0x0a, 0x06, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x2e, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x70, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x70, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x53, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x22, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x70, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69,
	0x64, 0x70, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x55, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x69, 0x64, 0x70, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x70, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28,
	0x01, 0x30, 0x01, 0x12, 0x57, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x24, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x70, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69,
	0x64, 0x70, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3d, 0x5a, 0x3b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2f, 0x69, 0x64, 0x70, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x3b, 0x69, 0x64, 0x70, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_aserto_idpplugin_v1_idpplugin_proto_rawDescOnce sync.Once
	file_aserto_idpplugin_v1_idpplugin_proto_rawDescData = file_aserto_idpplugin_v1_idpplugin_proto_rawDesc
)

func file_aserto_idpplugin_v1_idpplugin_proto_rawDescGZIP() []byte {
	file_aserto_idpplugin_v1_idpplugin_proto_rawDescOnce.Do(func() {
		file_aserto_idpplugin_v1_idpplugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_idpplugin_v1_idpplugin_proto_rawDescData)
	})
	return file_aserto_idpplugin_v1_idpplugin_proto_rawDescData
}

var file_aserto_idpplugin_v1_idpplugin_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_aserto_idpplugin_v1_idpplugin_proto_goTypes = []interface{}{
	(*InfoRequest)(nil),       // 0: aserto.idpplugin.v1.InfoRequest
	(*InfoResponse)(nil),      // 1: aserto.idpplugin.v1.InfoResponse
	(*ImportRequest)(nil),     // 2: aserto.idpplugin.v1.ImportRequest
	(*ImportResponse)(nil),    // 3: aserto.idpplugin.v1.ImportResponse
	(*ExportRequest)(nil),     // 4: aserto.idpplugin.v1.ExportRequest
	(*ExportResponse)(nil),    // 5: aserto.idpplugin.v1.ExportResponse
	(*DeleteRequest)(nil),     // 6: aserto.idpplugin.v1.DeleteRequest
	(*DeleteResponse)(nil),    // 7: aserto.idpplugin.v1.DeleteResponse
	(*ValidateRequest)(nil),   // 8: aserto.idpplugin.v1.ValidateRequest
	(*ValidateResponse)(nil),  // 9: aserto.idpplugin.v1.ValidateResponse
	(*v1.BuildInfo)(nil),      // 10: aserto.common.info.v1.BuildInfo
	(*v11.ConfigElement)(nil), // 11: aserto.api.v1.ConfigElement
	(*structpb.Struct)(nil),   // 12: google.protobuf.Struct
	(*v11.User)(nil),          // 13: aserto.api.v1.User
	(*v11.UserExt)(nil),       // 14: aserto.api.v1.UserExt
	(*status.Status)(nil),     // 15: google.rpc.Status
}
var file_aserto_idpplugin_v1_idpplugin_proto_depIdxs = []int32{
	10, // 0: aserto.idpplugin.v1.InfoResponse.build:type_name -> aserto.common.info.v1.BuildInfo
	11, // 1: aserto.idpplugin.v1.InfoResponse.configs:type_name -> aserto.api.v1.ConfigElement
	12, // 2: aserto.idpplugin.v1.ImportRequest.config:type_name -> google.protobuf.Struct
	13, // 3: aserto.idpplugin.v1.ImportRequest.user:type_name -> aserto.api.v1.User
	14, // 4: aserto.idpplugin.v1.ImportRequest.user_ext:type_name -> aserto.api.v1.UserExt
	15, // 5: aserto.idpplugin.v1.ImportResponse.error:type_name -> google.rpc.Status
	12, // 6: aserto.idpplugin.v1.ExportRequest.config:type_name -> google.protobuf.Struct
	13, // 7: aserto.idpplugin.v1.ExportResponse.user:type_name -> aserto.api.v1.User
	14, // 8: aserto.idpplugin.v1.ExportResponse.user_ext:type_name -> aserto.api.v1.UserExt
	15, // 9: aserto.idpplugin.v1.ExportResponse.error:type_name -> google.rpc.Status
	12, // 10: aserto.idpplugin.v1.DeleteRequest.config:type_name -> google.protobuf.Struct
	15, // 11: aserto.idpplugin.v1.DeleteResponse.error:type_name -> google.rpc.Status
	12, // 12: aserto.idpplugin.v1.ValidateRequest.config:type_name -> google.protobuf.Struct
	15, // 13: aserto.idpplugin.v1.ValidateResponse.status:type_name -> google.rpc.Status
	0,  // 14: aserto.idpplugin.v1.Plugin.Info:input_type -> aserto.idpplugin.v1.InfoRequest
	2,  // 15: aserto.idpplugin.v1.Plugin.Import:input_type -> aserto.idpplugin.v1.ImportRequest
	4,  // 16: aserto.idpplugin.v1.Plugin.Export:input_type -> aserto.idpplugin.v1.ExportRequest
	6,  // 17: aserto.idpplugin.v1.Plugin.Delete:input_type -> aserto.idpplugin.v1.DeleteRequest
	8,  // 18: aserto.idpplugin.v1.Plugin.Validate:input_type -> aserto.idpplugin.v1.ValidateRequest
	1,  // 19: aserto.idpplugin.v1.Plugin.Info:output_type -> aserto.idpplugin.v1.InfoResponse
	3,  // 20: aserto.idpplugin.v1.Plugin.Import:output_type -> aserto.idpplugin.v1.ImportResponse
	5,  // 21: aserto.idpplugin.v1.Plugin.Export:output_type -> aserto.idpplugin.v1.ExportResponse
	7,  // 22: aserto.idpplugin.v1.Plugin.Delete:output_type -> aserto.idpplugin.v1.DeleteResponse
	9,  // 23: aserto.idpplugin.v1.Plugin.Validate:output_type -> aserto.idpplugin.v1.ValidateResponse
	19, // [19:24] is the sub-list for method output_type
	14, // [14:19] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_aserto_idpplugin_v1_idpplugin_proto_init() }
func file_aserto_idpplugin_v1_idpplugin_proto_init() {
	if File_aserto_idpplugin_v1_idpplugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
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
		file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
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
		file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportRequest); i {
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
		file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportResponse); i {
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
		file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportRequest); i {
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
		file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportResponse); i {
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
		file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRequest); i {
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
		file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateResponse); i {
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
	file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ImportRequest_Config)(nil),
		(*ImportRequest_User)(nil),
		(*ImportRequest_UserExt)(nil),
	}
	file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ExportResponse_User)(nil),
		(*ExportResponse_UserExt)(nil),
		(*ExportResponse_Error)(nil),
	}
	file_aserto_idpplugin_v1_idpplugin_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*DeleteRequest_Config)(nil),
		(*DeleteRequest_UserId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aserto_idpplugin_v1_idpplugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_idpplugin_v1_idpplugin_proto_goTypes,
		DependencyIndexes: file_aserto_idpplugin_v1_idpplugin_proto_depIdxs,
		MessageInfos:      file_aserto_idpplugin_v1_idpplugin_proto_msgTypes,
	}.Build()
	File_aserto_idpplugin_v1_idpplugin_proto = out.File
	file_aserto_idpplugin_v1_idpplugin_proto_rawDesc = nil
	file_aserto_idpplugin_v1_idpplugin_proto_goTypes = nil
	file_aserto_idpplugin_v1_idpplugin_proto_depIdxs = nil
}