// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: aserto/api/v2/management.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// InstanceInfo has data about a running aserto instance that is registered with the control
// plane.
type InstanceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyId     string `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`             // The id of the policy associated with the instance.
	PolicyLabel  string `protobuf:"bytes,2,opt,name=policy_label,json=policyLabel,proto3" json:"policy_label,omitempty"`    // The label of the policy associated with the instance.
	RemoteHost   string `protobuf:"bytes,3,opt,name=remote_host,json=remoteHost,proto3" json:"remote_host,omitempty"`       // The name of the host where the instance is running.
	ConnectionId string `protobuf:"bytes,4,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"` // The connection an edge authorizer used to connect to the control plane.
}

func (x *InstanceInfo) Reset() {
	*x = InstanceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v2_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceInfo) ProtoMessage() {}

func (x *InstanceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v2_management_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceInfo.ProtoReflect.Descriptor instead.
func (*InstanceInfo) Descriptor() ([]byte, []int) {
	return file_aserto_api_v2_management_proto_rawDescGZIP(), []int{0}
}

func (x *InstanceInfo) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

func (x *InstanceInfo) GetPolicyLabel() string {
	if x != nil {
		return x.PolicyLabel
	}
	return ""
}

func (x *InstanceInfo) GetRemoteHost() string {
	if x != nil {
		return x.RemoteHost
	}
	return ""
}

func (x *InstanceInfo) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

// InstanceRegistration has data about an instance's registration with the control plane.
type InstanceRegistration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`     // Unique id of the instance registration, generated by the control plane.
	Info *InstanceInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"` // Information about the registered instance.
}

func (x *InstanceRegistration) Reset() {
	*x = InstanceRegistration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v2_management_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceRegistration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceRegistration) ProtoMessage() {}

func (x *InstanceRegistration) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v2_management_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceRegistration.ProtoReflect.Descriptor instead.
func (*InstanceRegistration) Descriptor() ([]byte, []int) {
	return file_aserto_api_v2_management_proto_rawDescGZIP(), []int{1}
}

func (x *InstanceRegistration) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InstanceRegistration) GetInfo() *InstanceInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// Command represents a command that may be sent to and then processed by an instance.
type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*Command_Discovery
	//	*Command_SyncEdgeDirectory
	Data isCommand_Data `protobuf_oneof:"data"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v2_management_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v2_management_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_aserto_api_v2_management_proto_rawDescGZIP(), []int{2}
}

func (m *Command) GetData() isCommand_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Command) GetDiscovery() *DiscoveryCommand {
	if x, ok := x.GetData().(*Command_Discovery); ok {
		return x.Discovery
	}
	return nil
}

func (x *Command) GetSyncEdgeDirectory() *SyncEdgeDirectoryCommand {
	if x, ok := x.GetData().(*Command_SyncEdgeDirectory); ok {
		return x.SyncEdgeDirectory
	}
	return nil
}

type isCommand_Data interface {
	isCommand_Data()
}

type Command_Discovery struct {
	Discovery *DiscoveryCommand `protobuf:"bytes,1,opt,name=discovery,proto3,oneof"`
}

type Command_SyncEdgeDirectory struct {
	SyncEdgeDirectory *SyncEdgeDirectoryCommand `protobuf:"bytes,2,opt,name=sync_edge_directory,json=syncEdgeDirectory,proto3,oneof"`
}

func (*Command_Discovery) isCommand_Data() {}

func (*Command_SyncEdgeDirectory) isCommand_Data() {}

// Run OPA discovery immediatley
type DiscoveryCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DiscoveryCommand) Reset() {
	*x = DiscoveryCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v2_management_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoveryCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryCommand) ProtoMessage() {}

func (x *DiscoveryCommand) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v2_management_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveryCommand.ProtoReflect.Descriptor instead.
func (*DiscoveryCommand) Descriptor() ([]byte, []int) {
	return file_aserto_api_v2_management_proto_rawDescGZIP(), []int{3}
}

// Sync the edge directory
type SyncEdgeDirectoryCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SyncEdgeDirectoryCommand) Reset() {
	*x = SyncEdgeDirectoryCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v2_management_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncEdgeDirectoryCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncEdgeDirectoryCommand) ProtoMessage() {}

func (x *SyncEdgeDirectoryCommand) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v2_management_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncEdgeDirectoryCommand.ProtoReflect.Descriptor instead.
func (*SyncEdgeDirectoryCommand) Descriptor() ([]byte, []int) {
	return file_aserto_api_v2_management_proto_rawDescGZIP(), []int{4}
}

var File_aserto_api_v2_management_proto protoreflect.FileDescriptor

var file_aserto_api_v2_management_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x22,
	0x94, 0x01, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x48, 0x6f, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x14, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f,
	0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22,
	0xad, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x3f, 0x0a, 0x09, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x48,
	0x00, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x59, 0x0a, 0x13,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x45, 0x64,
	0x67, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x48, 0x00, 0x52, 0x11, 0x73, 0x79, 0x6e, 0x63, 0x45, 0x64, 0x67, 0x65, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x12, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x53, 0x79, 0x6e, 0x63, 0x45, 0x64, 0x67, 0x65, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42,
	0x41, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x3b, 0x61,
	0x70, 0x69, 0xaa, 0x02, 0x0d, 0x41, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x41, 0x50, 0x49, 0x2e,
	0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_api_v2_management_proto_rawDescOnce sync.Once
	file_aserto_api_v2_management_proto_rawDescData = file_aserto_api_v2_management_proto_rawDesc
)

func file_aserto_api_v2_management_proto_rawDescGZIP() []byte {
	file_aserto_api_v2_management_proto_rawDescOnce.Do(func() {
		file_aserto_api_v2_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_api_v2_management_proto_rawDescData)
	})
	return file_aserto_api_v2_management_proto_rawDescData
}

var file_aserto_api_v2_management_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_aserto_api_v2_management_proto_goTypes = []interface{}{
	(*InstanceInfo)(nil),             // 0: aserto.api.v2.InstanceInfo
	(*InstanceRegistration)(nil),     // 1: aserto.api.v2.InstanceRegistration
	(*Command)(nil),                  // 2: aserto.api.v2.Command
	(*DiscoveryCommand)(nil),         // 3: aserto.api.v2.DiscoveryCommand
	(*SyncEdgeDirectoryCommand)(nil), // 4: aserto.api.v2.SyncEdgeDirectoryCommand
}
var file_aserto_api_v2_management_proto_depIdxs = []int32{
	0, // 0: aserto.api.v2.InstanceRegistration.info:type_name -> aserto.api.v2.InstanceInfo
	3, // 1: aserto.api.v2.Command.discovery:type_name -> aserto.api.v2.DiscoveryCommand
	4, // 2: aserto.api.v2.Command.sync_edge_directory:type_name -> aserto.api.v2.SyncEdgeDirectoryCommand
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_aserto_api_v2_management_proto_init() }
func file_aserto_api_v2_management_proto_init() {
	if File_aserto_api_v2_management_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_api_v2_management_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceInfo); i {
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
		file_aserto_api_v2_management_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceRegistration); i {
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
		file_aserto_api_v2_management_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_aserto_api_v2_management_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoveryCommand); i {
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
		file_aserto_api_v2_management_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncEdgeDirectoryCommand); i {
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
	file_aserto_api_v2_management_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Command_Discovery)(nil),
		(*Command_SyncEdgeDirectory)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aserto_api_v2_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_api_v2_management_proto_goTypes,
		DependencyIndexes: file_aserto_api_v2_management_proto_depIdxs,
		MessageInfos:      file_aserto_api_v2_management_proto_msgTypes,
	}.Build()
	File_aserto_api_v2_management_proto = out.File
	file_aserto_api_v2_management_proto_rawDesc = nil
	file_aserto_api_v2_management_proto_goTypes = nil
	file_aserto_api_v2_management_proto_depIdxs = nil
}
