// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/tenant/connection/v1/connection.proto

package connection

import (
	v1 "github.com/aserto-dev/go-grpc/aserto/api/v1"
	_ "github.com/aserto-dev/go-grpc/aserto/options/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

type ListConnectionsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kind          v1.ProviderKind        `protobuf:"varint,1,opt,name=kind,proto3,enum=aserto.api.v1.ProviderKind" json:"kind,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListConnectionsRequest) Reset() {
	*x = ListConnectionsRequest{}
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListConnectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConnectionsRequest) ProtoMessage() {}

func (x *ListConnectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConnectionsRequest.ProtoReflect.Descriptor instead.
func (*ListConnectionsRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP(), []int{0}
}

func (x *ListConnectionsRequest) GetKind() v1.ProviderKind {
	if x != nil {
		return x.Kind
	}
	return v1.ProviderKind(0)
}

type ListConnectionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*v1.Connection       `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListConnectionsResponse) Reset() {
	*x = ListConnectionsResponse{}
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListConnectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConnectionsResponse) ProtoMessage() {}

func (x *ListConnectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConnectionsResponse.ProtoReflect.Descriptor instead.
func (*ListConnectionsResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP(), []int{1}
}

func (x *ListConnectionsResponse) GetResults() []*v1.Connection {
	if x != nil {
		return x.Results
	}
	return nil
}

type GetConnectionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetConnectionRequest) Reset() {
	*x = GetConnectionRequest{}
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectionRequest) ProtoMessage() {}

func (x *GetConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectionRequest.ProtoReflect.Descriptor instead.
func (*GetConnectionRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP(), []int{2}
}

func (x *GetConnectionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetConnectionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        *v1.Connection         `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetConnectionResponse) Reset() {
	*x = GetConnectionResponse{}
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConnectionResponse) ProtoMessage() {}

func (x *GetConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConnectionResponse.ProtoReflect.Descriptor instead.
func (*GetConnectionResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP(), []int{3}
}

func (x *GetConnectionResponse) GetResult() *v1.Connection {
	if x != nil {
		return x.Result
	}
	return nil
}

type CreateConnectionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Connection    *v1.Connection         `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateConnectionRequest) Reset() {
	*x = CreateConnectionRequest{}
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConnectionRequest) ProtoMessage() {}

func (x *CreateConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConnectionRequest.ProtoReflect.Descriptor instead.
func (*CreateConnectionRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP(), []int{4}
}

func (x *CreateConnectionRequest) GetConnection() *v1.Connection {
	if x != nil {
		return x.Connection
	}
	return nil
}

type CreateConnectionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateConnectionResponse) Reset() {
	*x = CreateConnectionResponse{}
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConnectionResponse) ProtoMessage() {}

func (x *CreateConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConnectionResponse.ProtoReflect.Descriptor instead.
func (*CreateConnectionResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP(), []int{5}
}

func (x *CreateConnectionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateConnectionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Connection    *v1.Connection         `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	Force         bool                   `protobuf:"varint,2,opt,name=force,proto3" json:"force,omitempty"` // Use to force an update for readonly fields
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateConnectionRequest) Reset() {
	*x = UpdateConnectionRequest{}
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConnectionRequest) ProtoMessage() {}

func (x *UpdateConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConnectionRequest.ProtoReflect.Descriptor instead.
func (*UpdateConnectionRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateConnectionRequest) GetConnection() *v1.Connection {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *UpdateConnectionRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

type UpdateConnectionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateConnectionResponse) Reset() {
	*x = UpdateConnectionResponse{}
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConnectionResponse) ProtoMessage() {}

func (x *UpdateConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConnectionResponse.ProtoReflect.Descriptor instead.
func (*UpdateConnectionResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateConnectionResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteConnectionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteConnectionRequest) Reset() {
	*x = DeleteConnectionRequest{}
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConnectionRequest) ProtoMessage() {}

func (x *DeleteConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConnectionRequest.ProtoReflect.Descriptor instead.
func (*DeleteConnectionRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteConnectionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteConnectionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       *emptypb.Empty         `protobuf:"bytes,1,opt,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteConnectionResponse) Reset() {
	*x = DeleteConnectionResponse{}
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConnectionResponse) ProtoMessage() {}

func (x *DeleteConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConnectionResponse.ProtoReflect.Descriptor instead.
func (*DeleteConnectionResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteConnectionResponse) GetResults() *emptypb.Empty {
	if x != nil {
		return x.Results
	}
	return nil
}

type VerifyConnectionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyConnectionRequest) Reset() {
	*x = VerifyConnectionRequest{}
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyConnectionRequest) ProtoMessage() {}

func (x *VerifyConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyConnectionRequest.ProtoReflect.Descriptor instead.
func (*VerifyConnectionRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP(), []int{10}
}

func (x *VerifyConnectionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type VerifyConnectionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        *status.Status         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyConnectionResponse) Reset() {
	*x = VerifyConnectionResponse{}
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyConnectionResponse) ProtoMessage() {}

func (x *VerifyConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyConnectionResponse.ProtoReflect.Descriptor instead.
func (*VerifyConnectionResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP(), []int{11}
}

func (x *VerifyConnectionResponse) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type RotateSecretRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SecretKey     string                 `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RotateSecretRequest) Reset() {
	*x = RotateSecretRequest{}
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RotateSecretRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RotateSecretRequest) ProtoMessage() {}

func (x *RotateSecretRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RotateSecretRequest.ProtoReflect.Descriptor instead.
func (*RotateSecretRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP(), []int{12}
}

func (x *RotateSecretRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RotateSecretRequest) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

type RotateSecretResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        *v1.Connection         `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RotateSecretResponse) Reset() {
	*x = RotateSecretResponse{}
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RotateSecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RotateSecretResponse) ProtoMessage() {}

func (x *RotateSecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RotateSecretResponse.ProtoReflect.Descriptor instead.
func (*RotateSecretResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP(), []int{13}
}

func (x *RotateSecretResponse) GetResult() *v1.Connection {
	if x != nil {
		return x.Result
	}
	return nil
}

type ConnectionAvailableRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectionAvailableRequest) Reset() {
	*x = ConnectionAvailableRequest{}
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectionAvailableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionAvailableRequest) ProtoMessage() {}

func (x *ConnectionAvailableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionAvailableRequest.ProtoReflect.Descriptor instead.
func (*ConnectionAvailableRequest) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP(), []int{14}
}

func (x *ConnectionAvailableRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ConnectionAvailableResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Availability  v1.NameAvailability    `protobuf:"varint,1,opt,name=availability,proto3,enum=aserto.api.v1.NameAvailability" json:"availability,omitempty"`
	Reason        string                 `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectionAvailableResponse) Reset() {
	*x = ConnectionAvailableResponse{}
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectionAvailableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionAvailableResponse) ProtoMessage() {}

func (x *ConnectionAvailableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_tenant_connection_v1_connection_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionAvailableResponse.ProtoReflect.Descriptor instead.
func (*ConnectionAvailableResponse) Descriptor() ([]byte, []int) {
	return file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP(), []int{15}
}

func (x *ConnectionAvailableResponse) GetAvailability() v1.NameAvailability {
	if x != nil {
		return x.Availability
	}
	return v1.NameAvailability(0)
}

func (x *ConnectionAvailableResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_aserto_tenant_connection_v1_connection_proto protoreflect.FileDescriptor

const file_aserto_tenant_connection_v1_connection_proto_rawDesc = "" +
	"\n" +
	",aserto/tenant/connection/v1/connection.proto\x12\x1baserto.tenant.connection.v1\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x17google/rpc/status.proto\x1a\x1easerto/api/v1/connection.proto\x1a%aserto/api/v1/name_availability.proto\x1a\x1baserto/options/v1/ids.proto\"I\n" +
	"\x16ListConnectionsRequest\x12/\n" +
	"\x04kind\x18\x01 \x01(\x0e2\x1b.aserto.api.v1.ProviderKindR\x04kind\"N\n" +
	"\x17ListConnectionsResponse\x123\n" +
	"\aresults\x18\x01 \x03(\v2\x19.aserto.api.v1.ConnectionR\aresults\",\n" +
	"\x14GetConnectionRequest\x12\x14\n" +
	"\x02id\x18\x01 \x01(\tB\x04\x90\xb5\x18\aR\x02id\"J\n" +
	"\x15GetConnectionResponse\x121\n" +
	"\x06result\x18\x01 \x01(\v2\x19.aserto.api.v1.ConnectionR\x06result\"T\n" +
	"\x17CreateConnectionRequest\x129\n" +
	"\n" +
	"connection\x18\x01 \x01(\v2\x19.aserto.api.v1.ConnectionR\n" +
	"connection\"0\n" +
	"\x18CreateConnectionResponse\x12\x14\n" +
	"\x02id\x18\x01 \x01(\tB\x04\x90\xb5\x18\aR\x02id\"j\n" +
	"\x17UpdateConnectionRequest\x129\n" +
	"\n" +
	"connection\x18\x01 \x01(\v2\x19.aserto.api.v1.ConnectionR\n" +
	"connection\x12\x14\n" +
	"\x05force\x18\x02 \x01(\bR\x05force\"0\n" +
	"\x18UpdateConnectionResponse\x12\x14\n" +
	"\x02id\x18\x01 \x01(\tB\x04\x90\xb5\x18\aR\x02id\"/\n" +
	"\x17DeleteConnectionRequest\x12\x14\n" +
	"\x02id\x18\x01 \x01(\tB\x04\x90\xb5\x18\aR\x02id\"L\n" +
	"\x18DeleteConnectionResponse\x120\n" +
	"\aresults\x18\x01 \x01(\v2\x16.google.protobuf.EmptyR\aresults\"/\n" +
	"\x17VerifyConnectionRequest\x12\x14\n" +
	"\x02id\x18\x01 \x01(\tB\x04\x90\xb5\x18\aR\x02id\"F\n" +
	"\x18VerifyConnectionResponse\x12*\n" +
	"\x06status\x18\x01 \x01(\v2\x12.google.rpc.StatusR\x06status\"J\n" +
	"\x13RotateSecretRequest\x12\x14\n" +
	"\x02id\x18\x01 \x01(\tB\x04\x90\xb5\x18\aR\x02id\x12\x1d\n" +
	"\n" +
	"secret_key\x18\x02 \x01(\tR\tsecretKey\"I\n" +
	"\x14RotateSecretResponse\x121\n" +
	"\x06result\x18\x01 \x01(\v2\x19.aserto.api.v1.ConnectionR\x06result\"0\n" +
	"\x1aConnectionAvailableRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"z\n" +
	"\x1bConnectionAvailableResponse\x12C\n" +
	"\favailability\x18\x01 \x01(\x0e2\x1f.aserto.api.v1.NameAvailabilityR\favailability\x12\x16\n" +
	"\x06reason\x18\x02 \x01(\tR\x06reason2\xc6\x14\n" +
	"\n" +
	"Connection\x12\xb2\x02\n" +
	"\x0fListConnections\x123.aserto.tenant.connection.v1.ListConnectionsRequest\x1a4.aserto.tenant.connection.v1.ListConnectionsResponse\"\xb3\x01\x92A\x8d\x01\n" +
	"\n" +
	"connection\x12\x10List connections\x1a7Returns the collection of connections for given tenant.*\x1bconnection.list_connectionsb\x17\n" +
	"\a\n" +
	"\x03JWT\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02\x1c\x12\x1a/api/v1/tenant/connections\x12\xb4\x02\n" +
	"\rGetConnection\x121.aserto.tenant.connection.v1.GetConnectionRequest\x1a2.aserto.tenant.connection.v1.GetConnectionResponse\"\xbb\x01\x92A\x90\x01\n" +
	"\n" +
	"connection\x12\x0eGet connection\x1a>Returns the connection definition for the given connection id.*\x19connection.get_connectionb\x17\n" +
	"\a\n" +
	"\x03JWT\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02!\x12\x1f/api/v1/tenant/connections/{id}\x12\xc9\x02\n" +
	"\x10CreateConnection\x124.aserto.tenant.connection.v1.CreateConnectionRequest\x1a5.aserto.tenant.connection.v1.CreateConnectionResponse\"\xc7\x01\x92A\x95\x01\n" +
	"\n" +
	"connection\x12\x11Create connection\x1a=Creates a new connection instance of a given connection kind.*\x1cconnection.create_connectionb\x17\n" +
	"\a\n" +
	"\x03JWT\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02(:\n" +
	"connection\"\x1a/api/v1/tenant/connections\x12\xde\x02\n" +
	"\x10UpdateConnection\x124.aserto.tenant.connection.v1.UpdateConnectionRequest\x1a5.aserto.tenant.connection.v1.UpdateConnectionResponse\"\xdc\x01\x92A\x9a\x01\n" +
	"\n" +
	"connection\x12\x11Update connection\x1aBUpdate existing connection definition for the given connection id.*\x1cconnection.update_connectionb\x17\n" +
	"\a\n" +
	"\x03JWT\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x028:\n" +
	"connection\x1a*/api/v1/tenant/connections/{connection.id}\x12\xc3\x02\n" +
	"\x10DeleteConnection\x124.aserto.tenant.connection.v1.DeleteConnectionRequest\x1a5.aserto.tenant.connection.v1.DeleteConnectionResponse\"\xc1\x01\x92A\x96\x01\n" +
	"\n" +
	"connection\x12\x11Delete connection\x1a>Removes the connection definition for the given connection id.*\x1cconnection.delete_connectionb\x17\n" +
	"\a\n" +
	"\x03JWT\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02!*\x1f/api/v1/tenant/connections/{id}\x12\xcb\x02\n" +
	"\x10VerifyConnection\x124.aserto.tenant.connection.v1.VerifyConnectionRequest\x1a5.aserto.tenant.connection.v1.VerifyConnectionResponse\"\xc9\x01\x92A\x97\x01\n" +
	"\n" +
	"connection\x12\x11Verify connection\x1a?Verifiy the configuration settings for the given connection id.*\x1cconnection.verify_connectionb\x17\n" +
	"\a\n" +
	"\x03JWT\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02(\"&/api/v1/tenant/connections/{id}/verify\x12\xaf\x02\n" +
	"\fRotateSecret\x120.aserto.tenant.connection.v1.RotateSecretRequest\x1a1.aserto.tenant.connection.v1.RotateSecretResponse\"\xb9\x01\x92A{\n" +
	"\n" +
	"connection\x12\x18Rotate connection secret\x1a Rotate a generated secret value.*\x18connection.rotate_secretb\x17\n" +
	"\a\n" +
	"\x03JWT\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x025\x1a3/api/v1/tenant/connections/{id}/{secret_key}/rotate\x12\xd8\x02\n" +
	"\x13ConnectionAvailable\x127.aserto.tenant.connection.v1.ConnectionAvailableRequest\x1a8.aserto.tenant.connection.v1.ConnectionAvailableResponse\"\xcd\x01\x92A\x96\x01\n" +
	"\n" +
	"connection\x12\x1dConnection availability check\x1a/Verifies if given connection name is available.*\x1fconnection.connection_availableb\x17\n" +
	"\a\n" +
	"\x03JWT\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02-\x12+/api/v1/tenant/connections/available/{name}B\xcc\x01\x92A\x82\x01*\x01\x022\x10application/json:\x10application/jsonZ@\n" +
	"\x1a\n" +
	"\x03JWT\x12\x13\b\x02\x1a\rauthorization \x02\n" +
	"\"\n" +
	"\bTenantID\x12\x16\b\x02\x1a\x10aserto-tenant-id \x02b\x17\n" +
	"\a\n" +
	"\x03JWT\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00ZDgithub.com/aserto-dev/go-grpc/aserto/tenant/connection/v1;connectionb\x06proto3"

var (
	file_aserto_tenant_connection_v1_connection_proto_rawDescOnce sync.Once
	file_aserto_tenant_connection_v1_connection_proto_rawDescData []byte
)

func file_aserto_tenant_connection_v1_connection_proto_rawDescGZIP() []byte {
	file_aserto_tenant_connection_v1_connection_proto_rawDescOnce.Do(func() {
		file_aserto_tenant_connection_v1_connection_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_tenant_connection_v1_connection_proto_rawDesc), len(file_aserto_tenant_connection_v1_connection_proto_rawDesc)))
	})
	return file_aserto_tenant_connection_v1_connection_proto_rawDescData
}

var file_aserto_tenant_connection_v1_connection_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_aserto_tenant_connection_v1_connection_proto_goTypes = []any{
	(*ListConnectionsRequest)(nil),      // 0: aserto.tenant.connection.v1.ListConnectionsRequest
	(*ListConnectionsResponse)(nil),     // 1: aserto.tenant.connection.v1.ListConnectionsResponse
	(*GetConnectionRequest)(nil),        // 2: aserto.tenant.connection.v1.GetConnectionRequest
	(*GetConnectionResponse)(nil),       // 3: aserto.tenant.connection.v1.GetConnectionResponse
	(*CreateConnectionRequest)(nil),     // 4: aserto.tenant.connection.v1.CreateConnectionRequest
	(*CreateConnectionResponse)(nil),    // 5: aserto.tenant.connection.v1.CreateConnectionResponse
	(*UpdateConnectionRequest)(nil),     // 6: aserto.tenant.connection.v1.UpdateConnectionRequest
	(*UpdateConnectionResponse)(nil),    // 7: aserto.tenant.connection.v1.UpdateConnectionResponse
	(*DeleteConnectionRequest)(nil),     // 8: aserto.tenant.connection.v1.DeleteConnectionRequest
	(*DeleteConnectionResponse)(nil),    // 9: aserto.tenant.connection.v1.DeleteConnectionResponse
	(*VerifyConnectionRequest)(nil),     // 10: aserto.tenant.connection.v1.VerifyConnectionRequest
	(*VerifyConnectionResponse)(nil),    // 11: aserto.tenant.connection.v1.VerifyConnectionResponse
	(*RotateSecretRequest)(nil),         // 12: aserto.tenant.connection.v1.RotateSecretRequest
	(*RotateSecretResponse)(nil),        // 13: aserto.tenant.connection.v1.RotateSecretResponse
	(*ConnectionAvailableRequest)(nil),  // 14: aserto.tenant.connection.v1.ConnectionAvailableRequest
	(*ConnectionAvailableResponse)(nil), // 15: aserto.tenant.connection.v1.ConnectionAvailableResponse
	(v1.ProviderKind)(0),                // 16: aserto.api.v1.ProviderKind
	(*v1.Connection)(nil),               // 17: aserto.api.v1.Connection
	(*emptypb.Empty)(nil),               // 18: google.protobuf.Empty
	(*status.Status)(nil),               // 19: google.rpc.Status
	(v1.NameAvailability)(0),            // 20: aserto.api.v1.NameAvailability
}
var file_aserto_tenant_connection_v1_connection_proto_depIdxs = []int32{
	16, // 0: aserto.tenant.connection.v1.ListConnectionsRequest.kind:type_name -> aserto.api.v1.ProviderKind
	17, // 1: aserto.tenant.connection.v1.ListConnectionsResponse.results:type_name -> aserto.api.v1.Connection
	17, // 2: aserto.tenant.connection.v1.GetConnectionResponse.result:type_name -> aserto.api.v1.Connection
	17, // 3: aserto.tenant.connection.v1.CreateConnectionRequest.connection:type_name -> aserto.api.v1.Connection
	17, // 4: aserto.tenant.connection.v1.UpdateConnectionRequest.connection:type_name -> aserto.api.v1.Connection
	18, // 5: aserto.tenant.connection.v1.DeleteConnectionResponse.results:type_name -> google.protobuf.Empty
	19, // 6: aserto.tenant.connection.v1.VerifyConnectionResponse.status:type_name -> google.rpc.Status
	17, // 7: aserto.tenant.connection.v1.RotateSecretResponse.result:type_name -> aserto.api.v1.Connection
	20, // 8: aserto.tenant.connection.v1.ConnectionAvailableResponse.availability:type_name -> aserto.api.v1.NameAvailability
	0,  // 9: aserto.tenant.connection.v1.Connection.ListConnections:input_type -> aserto.tenant.connection.v1.ListConnectionsRequest
	2,  // 10: aserto.tenant.connection.v1.Connection.GetConnection:input_type -> aserto.tenant.connection.v1.GetConnectionRequest
	4,  // 11: aserto.tenant.connection.v1.Connection.CreateConnection:input_type -> aserto.tenant.connection.v1.CreateConnectionRequest
	6,  // 12: aserto.tenant.connection.v1.Connection.UpdateConnection:input_type -> aserto.tenant.connection.v1.UpdateConnectionRequest
	8,  // 13: aserto.tenant.connection.v1.Connection.DeleteConnection:input_type -> aserto.tenant.connection.v1.DeleteConnectionRequest
	10, // 14: aserto.tenant.connection.v1.Connection.VerifyConnection:input_type -> aserto.tenant.connection.v1.VerifyConnectionRequest
	12, // 15: aserto.tenant.connection.v1.Connection.RotateSecret:input_type -> aserto.tenant.connection.v1.RotateSecretRequest
	14, // 16: aserto.tenant.connection.v1.Connection.ConnectionAvailable:input_type -> aserto.tenant.connection.v1.ConnectionAvailableRequest
	1,  // 17: aserto.tenant.connection.v1.Connection.ListConnections:output_type -> aserto.tenant.connection.v1.ListConnectionsResponse
	3,  // 18: aserto.tenant.connection.v1.Connection.GetConnection:output_type -> aserto.tenant.connection.v1.GetConnectionResponse
	5,  // 19: aserto.tenant.connection.v1.Connection.CreateConnection:output_type -> aserto.tenant.connection.v1.CreateConnectionResponse
	7,  // 20: aserto.tenant.connection.v1.Connection.UpdateConnection:output_type -> aserto.tenant.connection.v1.UpdateConnectionResponse
	9,  // 21: aserto.tenant.connection.v1.Connection.DeleteConnection:output_type -> aserto.tenant.connection.v1.DeleteConnectionResponse
	11, // 22: aserto.tenant.connection.v1.Connection.VerifyConnection:output_type -> aserto.tenant.connection.v1.VerifyConnectionResponse
	13, // 23: aserto.tenant.connection.v1.Connection.RotateSecret:output_type -> aserto.tenant.connection.v1.RotateSecretResponse
	15, // 24: aserto.tenant.connection.v1.Connection.ConnectionAvailable:output_type -> aserto.tenant.connection.v1.ConnectionAvailableResponse
	17, // [17:25] is the sub-list for method output_type
	9,  // [9:17] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_aserto_tenant_connection_v1_connection_proto_init() }
func file_aserto_tenant_connection_v1_connection_proto_init() {
	if File_aserto_tenant_connection_v1_connection_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_tenant_connection_v1_connection_proto_rawDesc), len(file_aserto_tenant_connection_v1_connection_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_tenant_connection_v1_connection_proto_goTypes,
		DependencyIndexes: file_aserto_tenant_connection_v1_connection_proto_depIdxs,
		MessageInfos:      file_aserto_tenant_connection_v1_connection_proto_msgTypes,
	}.Build()
	File_aserto_tenant_connection_v1_connection_proto = out.File
	file_aserto_tenant_connection_v1_connection_proto_goTypes = nil
	file_aserto_tenant_connection_v1_connection_proto_depIdxs = nil
}
