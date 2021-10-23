// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.17.3
// source: aserto/api/v1/connection.proto

package api

import (
	_ "github.com/aserto-dev/go-grpc/aserto/options/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type ProviderKind int32

const (
	ProviderKind_PROVIDER_KIND_UNKNOWN         ProviderKind = 0 // Unknown provider kind
	ProviderKind_PROVIDER_KIND_IDP             ProviderKind = 1 // Identity Provider
	ProviderKind_PROVIDER_KIND_SCC             ProviderKind = 2 // Source Code Control Providers
	ProviderKind_PROVIDER_KIND_POLICY_REGISTRY ProviderKind = 3 // Provider for a registry of policies
	ProviderKind_PROVIDER_KIND_AUTHORIZER      ProviderKind = 4 // Provider of an authorization service
)

// Enum value maps for ProviderKind.
var (
	ProviderKind_name = map[int32]string{
		0: "PROVIDER_KIND_UNKNOWN",
		1: "PROVIDER_KIND_IDP",
		2: "PROVIDER_KIND_SCC",
		3: "PROVIDER_KIND_POLICY_REGISTRY",
		4: "PROVIDER_KIND_AUTHORIZER",
	}
	ProviderKind_value = map[string]int32{
		"PROVIDER_KIND_UNKNOWN":         0,
		"PROVIDER_KIND_IDP":             1,
		"PROVIDER_KIND_SCC":             2,
		"PROVIDER_KIND_POLICY_REGISTRY": 3,
		"PROVIDER_KIND_AUTHORIZER":      4,
	}
)

func (x ProviderKind) Enum() *ProviderKind {
	p := new(ProviderKind)
	*p = x
	return p
}

func (x ProviderKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProviderKind) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_api_v1_connection_proto_enumTypes[0].Descriptor()
}

func (ProviderKind) Type() protoreflect.EnumType {
	return &file_aserto_api_v1_connection_proto_enumTypes[0]
}

func (x ProviderKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProviderKind.Descriptor instead.
func (ProviderKind) EnumDescriptor() ([]byte, []int) {
	return file_aserto_api_v1_connection_proto_rawDescGZIP(), []int{0}
}

type ConnectionType int32

const (
	ConnectionType_CONNECTION_TYPE_UNKNOWN ConnectionType = 0
	ConnectionType_CONNECTION_TYPE_SIMPLE  ConnectionType = 1
	ConnectionType_CONNECTION_TYPE_OAUTH   ConnectionType = 2
)

// Enum value maps for ConnectionType.
var (
	ConnectionType_name = map[int32]string{
		0: "CONNECTION_TYPE_UNKNOWN",
		1: "CONNECTION_TYPE_SIMPLE",
		2: "CONNECTION_TYPE_OAUTH",
	}
	ConnectionType_value = map[string]int32{
		"CONNECTION_TYPE_UNKNOWN": 0,
		"CONNECTION_TYPE_SIMPLE":  1,
		"CONNECTION_TYPE_OAUTH":   2,
	}
)

func (x ConnectionType) Enum() *ConnectionType {
	p := new(ConnectionType)
	*p = x
	return p
}

func (x ConnectionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionType) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_api_v1_connection_proto_enumTypes[1].Descriptor()
}

func (ConnectionType) Type() protoreflect.EnumType {
	return &file_aserto_api_v1_connection_proto_enumTypes[1]
}

func (x ConnectionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionType.Descriptor instead.
func (ConnectionType) EnumDescriptor() ([]byte, []int) {
	return file_aserto_api_v1_connection_proto_rawDescGZIP(), []int{1}
}

type ConfigElementKind int32

const (
	ConfigElementKind_CONFIG_ELEMENT_KIND_UNKNOWN     ConfigElementKind = 0 // Unknown configuration element kind
	ConfigElementKind_CONFIG_ELEMENT_KIND_ATTRIBUTE   ConfigElementKind = 1 // Normal attribute
	ConfigElementKind_CONFIG_ELEMENT_KIND_SECRET      ConfigElementKind = 2 // Secret
	ConfigElementKind_CONFIG_ELEMENT_KIND_CERTIFICATE ConfigElementKind = 3 // Certificate
)

// Enum value maps for ConfigElementKind.
var (
	ConfigElementKind_name = map[int32]string{
		0: "CONFIG_ELEMENT_KIND_UNKNOWN",
		1: "CONFIG_ELEMENT_KIND_ATTRIBUTE",
		2: "CONFIG_ELEMENT_KIND_SECRET",
		3: "CONFIG_ELEMENT_KIND_CERTIFICATE",
	}
	ConfigElementKind_value = map[string]int32{
		"CONFIG_ELEMENT_KIND_UNKNOWN":     0,
		"CONFIG_ELEMENT_KIND_ATTRIBUTE":   1,
		"CONFIG_ELEMENT_KIND_SECRET":      2,
		"CONFIG_ELEMENT_KIND_CERTIFICATE": 3,
	}
)

func (x ConfigElementKind) Enum() *ConfigElementKind {
	p := new(ConfigElementKind)
	*p = x
	return p
}

func (x ConfigElementKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigElementKind) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_api_v1_connection_proto_enumTypes[2].Descriptor()
}

func (ConfigElementKind) Type() protoreflect.EnumType {
	return &file_aserto_api_v1_connection_proto_enumTypes[2]
}

func (x ConfigElementKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigElementKind.Descriptor instead.
func (ConfigElementKind) EnumDescriptor() ([]byte, []int) {
	return file_aserto_api_v1_connection_proto_rawDescGZIP(), []int{2}
}

type ConfigElementType int32

const (
	ConfigElementType_CONFIG_ELEMENT_TYPE_UNKNOWN ConfigElementType = 0 // Unknown configuration element type
	ConfigElementType_CONFIG_ELEMENT_TYPE_STRING  ConfigElementType = 1 // Normal attribute
	ConfigElementType_CONFIG_ELEMENT_TYPE_INTEGER ConfigElementType = 2 // Secret
	ConfigElementType_CONFIG_ELEMENT_TYPE_BOOLEAN ConfigElementType = 3 // Certificate
)

// Enum value maps for ConfigElementType.
var (
	ConfigElementType_name = map[int32]string{
		0: "CONFIG_ELEMENT_TYPE_UNKNOWN",
		1: "CONFIG_ELEMENT_TYPE_STRING",
		2: "CONFIG_ELEMENT_TYPE_INTEGER",
		3: "CONFIG_ELEMENT_TYPE_BOOLEAN",
	}
	ConfigElementType_value = map[string]int32{
		"CONFIG_ELEMENT_TYPE_UNKNOWN": 0,
		"CONFIG_ELEMENT_TYPE_STRING":  1,
		"CONFIG_ELEMENT_TYPE_INTEGER": 2,
		"CONFIG_ELEMENT_TYPE_BOOLEAN": 3,
	}
)

func (x ConfigElementType) Enum() *ConfigElementType {
	p := new(ConfigElementType)
	*p = x
	return p
}

func (x ConfigElementType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigElementType) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_api_v1_connection_proto_enumTypes[3].Descriptor()
}

func (ConfigElementType) Type() protoreflect.EnumType {
	return &file_aserto_api_v1_connection_proto_enumTypes[3]
}

func (x ConfigElementType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigElementType.Descriptor instead.
func (ConfigElementType) EnumDescriptor() ([]byte, []int) {
	return file_aserto_api_v1_connection_proto_rawDescGZIP(), []int{3}
}

type DisplayMode int32

const (
	DisplayMode_DISPLAY_MODE_UNKNOWN DisplayMode = 0
	DisplayMode_DISPLAY_MODE_NORMAL  DisplayMode = 1
	DisplayMode_DISPLAY_MODE_MASKED  DisplayMode = 2
)

// Enum value maps for DisplayMode.
var (
	DisplayMode_name = map[int32]string{
		0: "DISPLAY_MODE_UNKNOWN",
		1: "DISPLAY_MODE_NORMAL",
		2: "DISPLAY_MODE_MASKED",
	}
	DisplayMode_value = map[string]int32{
		"DISPLAY_MODE_UNKNOWN": 0,
		"DISPLAY_MODE_NORMAL":  1,
		"DISPLAY_MODE_MASKED":  2,
	}
)

func (x DisplayMode) Enum() *DisplayMode {
	p := new(DisplayMode)
	*p = x
	return p
}

func (x DisplayMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DisplayMode) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_api_v1_connection_proto_enumTypes[4].Descriptor()
}

func (DisplayMode) Type() protoreflect.EnumType {
	return &file_aserto_api_v1_connection_proto_enumTypes[4]
}

func (x DisplayMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DisplayMode.Descriptor instead.
func (DisplayMode) EnumDescriptor() ([]byte, []int) {
	return file_aserto_api_v1_connection_proto_rawDescGZIP(), []int{4}
}

type Provider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                                                                                                                            // provider short name
	Description       string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`                                                                                                                              // provider description
	Kind              ProviderKind      `protobuf:"varint,4,opt,name=kind,proto3,enum=aserto.api.v1.ProviderKind" json:"kind,omitempty"`                                                                                                           // provider kind enum [scc|idp]
	ConnectionType    ConnectionType    `protobuf:"varint,5,opt,name=connection_type,json=connectionType,proto3,enum=aserto.api.v1.ConnectionType" json:"connection_type,omitempty"`                                                               // connection type enum [simple|oauth]
	DisplayAttributes map[string]string `protobuf:"bytes,6,rep,name=display_attributes,json=displayAttributes,proto3" json:"display_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // display attributes
	Config            []*ConfigElement  `protobuf:"bytes,7,rep,name=config,proto3" json:"config,omitempty"`                                                                                                                                        // JSON struct with all configurable settings (schema)
}

func (x *Provider) Reset() {
	*x = Provider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v1_connection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Provider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Provider) ProtoMessage() {}

func (x *Provider) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v1_connection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Provider.ProtoReflect.Descriptor instead.
func (*Provider) Descriptor() ([]byte, []int) {
	return file_aserto_api_v1_connection_proto_rawDescGZIP(), []int{0}
}

func (x *Provider) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Provider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Provider) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Provider) GetKind() ProviderKind {
	if x != nil {
		return x.Kind
	}
	return ProviderKind_PROVIDER_KIND_UNKNOWN
}

func (x *Provider) GetConnectionType() ConnectionType {
	if x != nil {
		return x.ConnectionType
	}
	return ConnectionType_CONNECTION_TYPE_UNKNOWN
}

func (x *Provider) GetDisplayAttributes() map[string]string {
	if x != nil {
		return x.DisplayAttributes
	}
	return nil
}

func (x *Provider) GetConfig() []*ConfigElement {
	if x != nil {
		return x.Config
	}
	return nil
}

type Connection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                                              // connection instance ID (unique withing tenant scope)
	Name               string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                                          // user defined name
	Description        string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`                                            // user defined description
	Kind               ProviderKind           `protobuf:"varint,4,opt,name=kind,proto3,enum=aserto.api.v1.ProviderKind" json:"kind,omitempty"`                         // provider kind enum [scc|idp]
	ProviderId         string                 `protobuf:"bytes,5,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`                            // provider ID used by connection
	Verified           bool                   `protobuf:"varint,6,opt,name=verified,proto3" json:"verified,omitempty"`                                                 // flag indicating if the connection is validated and wired up
	System             bool                   `protobuf:"varint,7,opt,name=system,proto3" json:"system,omitempty"`                                                     // flag indicating if this is a system connection
	Config             *structpb.Struct       `protobuf:"bytes,8,opt,name=config,proto3" json:"config,omitempty"`                                                      // JSON struct with config values to be stored in Vault?, these would not be returned on a get
	Metadata           *Metadata              `protobuf:"bytes,9,opt,name=metadata,proto3" json:"metadata,omitempty"`                                                  // generic metadata results block
	LastVerificationAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=last_verification_at,json=lastVerificationAt,proto3" json:"last_verification_at,omitempty"` // last verification time (UTC)
}

func (x *Connection) Reset() {
	*x = Connection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v1_connection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connection) ProtoMessage() {}

func (x *Connection) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v1_connection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connection.ProtoReflect.Descriptor instead.
func (*Connection) Descriptor() ([]byte, []int) {
	return file_aserto_api_v1_connection_proto_rawDescGZIP(), []int{1}
}

func (x *Connection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Connection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Connection) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Connection) GetKind() ProviderKind {
	if x != nil {
		return x.Kind
	}
	return ProviderKind_PROVIDER_KIND_UNKNOWN
}

func (x *Connection) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

func (x *Connection) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *Connection) GetSystem() bool {
	if x != nil {
		return x.System
	}
	return false
}

func (x *Connection) GetConfig() *structpb.Struct {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Connection) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Connection) GetLastVerificationAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastVerificationAt
	}
	return nil
}

type ConfigElement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Kind        ConfigElementKind `protobuf:"varint,2,opt,name=kind,proto3,enum=aserto.api.v1.ConfigElementKind" json:"kind,omitempty"`
	Type        ConfigElementType `protobuf:"varint,3,opt,name=type,proto3,enum=aserto.api.v1.ConfigElementType" json:"type,omitempty"`
	Name        string            `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description string            `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Usage       string            `protobuf:"bytes,6,opt,name=usage,proto3" json:"usage,omitempty"`
	Mode        DisplayMode       `protobuf:"varint,7,opt,name=mode,proto3,enum=aserto.api.v1.DisplayMode" json:"mode,omitempty"`
	ReadOnly    bool              `protobuf:"varint,8,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	Generated   bool              `protobuf:"varint,9,opt,name=generated,proto3" json:"generated,omitempty"`
}

func (x *ConfigElement) Reset() {
	*x = ConfigElement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v1_connection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigElement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigElement) ProtoMessage() {}

func (x *ConfigElement) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v1_connection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigElement.ProtoReflect.Descriptor instead.
func (*ConfigElement) Descriptor() ([]byte, []int) {
	return file_aserto_api_v1_connection_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigElement) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ConfigElement) GetKind() ConfigElementKind {
	if x != nil {
		return x.Kind
	}
	return ConfigElementKind_CONFIG_ELEMENT_KIND_UNKNOWN
}

func (x *ConfigElement) GetType() ConfigElementType {
	if x != nil {
		return x.Type
	}
	return ConfigElementType_CONFIG_ELEMENT_TYPE_UNKNOWN
}

func (x *ConfigElement) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConfigElement) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ConfigElement) GetUsage() string {
	if x != nil {
		return x.Usage
	}
	return ""
}

func (x *ConfigElement) GetMode() DisplayMode {
	if x != nil {
		return x.Mode
	}
	return DisplayMode_DISPLAY_MODE_UNKNOWN
}

func (x *ConfigElement) GetReadOnly() bool {
	if x != nil {
		return x.ReadOnly
	}
	return false
}

func (x *ConfigElement) GetGenerated() bool {
	if x != nil {
		return x.Generated
	}
	return false
}

var File_aserto_api_v1_connection_proto protoreflect.FileDescriptor

var file_aserto_api_v1_connection_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x03, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x04, 0x90, 0xb5, 0x18, 0x06, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x12, 0x46, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5d, 0x0a, 0x12, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x1a, 0x44, 0x0a, 0x16, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x98, 0x03, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0x90, 0xb5, 0x18, 0x07, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2f, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1b, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x12, 0x25, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x90, 0xb5, 0x18, 0x06, 0x52, 0x0a, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x2f, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x33,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x4c, 0x0a, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x6c,
	0x61, 0x73, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x74, 0x22, 0xc2, 0x02, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2a, 0x98, 0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x56, 0x49,
	0x44, 0x45, 0x52, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x4b,
	0x49, 0x4e, 0x44, 0x5f, 0x49, 0x44, 0x50, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x52, 0x4f,
	0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x53, 0x43, 0x43, 0x10, 0x02,
	0x12, 0x21, 0x0a, 0x1d, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x4b, 0x49, 0x4e,
	0x44, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52,
	0x59, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f,
	0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x52, 0x10,
	0x04, 0x2a, 0x64, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x53, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4f, 0x41, 0x55, 0x54, 0x48, 0x10, 0x02, 0x2a, 0x9c, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1f, 0x0a,
	0x1b, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x21,
	0x0a, 0x1d, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x10,
	0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x45, 0x4c, 0x45, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x10,
	0x02, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x45, 0x4c, 0x45, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49,
	0x43, 0x41, 0x54, 0x45, 0x10, 0x03, 0x2a, 0x96, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b,
	0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1e, 0x0a,
	0x1a, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1f, 0x0a,
	0x1b, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x10, 0x02, 0x12, 0x1f,
	0x0a, 0x1b, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x03, 0x2a,
	0x59, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x14, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x49, 0x53, 0x50,
	0x4c, 0x41, 0x59, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10,
	0x01, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x4d, 0x4f, 0x44,
	0x45, 0x5f, 0x4d, 0x41, 0x53, 0x4b, 0x45, 0x44, 0x10, 0x02, 0x42, 0x41, 0x5a, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d,
	0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0xaa, 0x02, 0x0d,
	0x41, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_api_v1_connection_proto_rawDescOnce sync.Once
	file_aserto_api_v1_connection_proto_rawDescData = file_aserto_api_v1_connection_proto_rawDesc
)

func file_aserto_api_v1_connection_proto_rawDescGZIP() []byte {
	file_aserto_api_v1_connection_proto_rawDescOnce.Do(func() {
		file_aserto_api_v1_connection_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_api_v1_connection_proto_rawDescData)
	})
	return file_aserto_api_v1_connection_proto_rawDescData
}

var file_aserto_api_v1_connection_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_aserto_api_v1_connection_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_aserto_api_v1_connection_proto_goTypes = []interface{}{
	(ProviderKind)(0),             // 0: aserto.api.v1.ProviderKind
	(ConnectionType)(0),           // 1: aserto.api.v1.ConnectionType
	(ConfigElementKind)(0),        // 2: aserto.api.v1.ConfigElementKind
	(ConfigElementType)(0),        // 3: aserto.api.v1.ConfigElementType
	(DisplayMode)(0),              // 4: aserto.api.v1.DisplayMode
	(*Provider)(nil),              // 5: aserto.api.v1.Provider
	(*Connection)(nil),            // 6: aserto.api.v1.Connection
	(*ConfigElement)(nil),         // 7: aserto.api.v1.ConfigElement
	nil,                           // 8: aserto.api.v1.Provider.DisplayAttributesEntry
	(*structpb.Struct)(nil),       // 9: google.protobuf.Struct
	(*Metadata)(nil),              // 10: aserto.api.v1.Metadata
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_aserto_api_v1_connection_proto_depIdxs = []int32{
	0,  // 0: aserto.api.v1.Provider.kind:type_name -> aserto.api.v1.ProviderKind
	1,  // 1: aserto.api.v1.Provider.connection_type:type_name -> aserto.api.v1.ConnectionType
	8,  // 2: aserto.api.v1.Provider.display_attributes:type_name -> aserto.api.v1.Provider.DisplayAttributesEntry
	7,  // 3: aserto.api.v1.Provider.config:type_name -> aserto.api.v1.ConfigElement
	0,  // 4: aserto.api.v1.Connection.kind:type_name -> aserto.api.v1.ProviderKind
	9,  // 5: aserto.api.v1.Connection.config:type_name -> google.protobuf.Struct
	10, // 6: aserto.api.v1.Connection.metadata:type_name -> aserto.api.v1.Metadata
	11, // 7: aserto.api.v1.Connection.last_verification_at:type_name -> google.protobuf.Timestamp
	2,  // 8: aserto.api.v1.ConfigElement.kind:type_name -> aserto.api.v1.ConfigElementKind
	3,  // 9: aserto.api.v1.ConfigElement.type:type_name -> aserto.api.v1.ConfigElementType
	4,  // 10: aserto.api.v1.ConfigElement.mode:type_name -> aserto.api.v1.DisplayMode
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_aserto_api_v1_connection_proto_init() }
func file_aserto_api_v1_connection_proto_init() {
	if File_aserto_api_v1_connection_proto != nil {
		return
	}
	file_aserto_api_v1_metadata_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_aserto_api_v1_connection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Provider); i {
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
		file_aserto_api_v1_connection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connection); i {
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
		file_aserto_api_v1_connection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigElement); i {
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
			RawDescriptor: file_aserto_api_v1_connection_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_api_v1_connection_proto_goTypes,
		DependencyIndexes: file_aserto_api_v1_connection_proto_depIdxs,
		EnumInfos:         file_aserto_api_v1_connection_proto_enumTypes,
		MessageInfos:      file_aserto_api_v1_connection_proto_msgTypes,
	}.Build()
	File_aserto_api_v1_connection_proto = out.File
	file_aserto_api_v1_connection_proto_rawDesc = nil
	file_aserto_api_v1_connection_proto_goTypes = nil
	file_aserto_api_v1_connection_proto_depIdxs = nil
}
