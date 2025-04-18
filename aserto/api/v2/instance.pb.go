// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/api/v2/instance.proto

package api

import (
	_ "github.com/aserto-dev/go-grpc/aserto/options/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type InstanceType int32

const (
	InstanceType_INSTANCE_TYPE_UNKNOWN         InstanceType = 0 // Value not set.
	InstanceType_INSTANCE_TYPE_HOSTED          InstanceType = 1 // Instance runs in a hosted environment.
	InstanceType_INSTANCE_TYPE_EDGE_AUTHORIZER InstanceType = 2 // Instance runs in the user's network.
)

// Enum value maps for InstanceType.
var (
	InstanceType_name = map[int32]string{
		0: "INSTANCE_TYPE_UNKNOWN",
		1: "INSTANCE_TYPE_HOSTED",
		2: "INSTANCE_TYPE_EDGE_AUTHORIZER",
	}
	InstanceType_value = map[string]int32{
		"INSTANCE_TYPE_UNKNOWN":         0,
		"INSTANCE_TYPE_HOSTED":          1,
		"INSTANCE_TYPE_EDGE_AUTHORIZER": 2,
	}
)

func (x InstanceType) Enum() *InstanceType {
	p := new(InstanceType)
	*p = x
	return p
}

func (x InstanceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstanceType) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_api_v2_instance_proto_enumTypes[0].Descriptor()
}

func (InstanceType) Type() protoreflect.EnumType {
	return &file_aserto_api_v2_instance_proto_enumTypes[0]
}

func (x InstanceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstanceType.Descriptor instead.
func (InstanceType) EnumDescriptor() ([]byte, []int) {
	return file_aserto_api_v2_instance_proto_rawDescGZIP(), []int{0}
}

type Instance struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	PolicyId        string                 `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`                                              // Policy ID.
	Label           string                 `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`                                                                    // Unique human-readable label for the instance.
	Tag             string                 `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`                                                                        // Default repository tag to be run by the instance.
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`                                           // Timestamp of when the instance was created.
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`                                           // Timestamp of when the instance was last updated.
	InstanceType    InstanceType           `protobuf:"varint,6,opt,name=instance_type,json=instanceType,proto3,enum=aserto.api.v2.InstanceType" json:"instance_type,omitempty"` // Type of instance (hosted or edge).
	DecisionLogging *bool                  `protobuf:"varint,7,opt,name=decision_logging,json=decisionLogging,proto3,oneof" json:"decision_logging,omitempty"`                  // Whether to enable decision logging.
	VersionHash     string                 `protobuf:"bytes,8,opt,name=version_hash,json=versionHash,proto3" json:"version_hash,omitempty"`                                     // Sha256 string of policy_id, label, tag and decision_logging concatenated values.
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Instance) Reset() {
	*x = Instance{}
	mi := &file_aserto_api_v2_instance_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instance) ProtoMessage() {}

func (x *Instance) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v2_instance_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instance.ProtoReflect.Descriptor instead.
func (*Instance) Descriptor() ([]byte, []int) {
	return file_aserto_api_v2_instance_proto_rawDescGZIP(), []int{0}
}

func (x *Instance) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

func (x *Instance) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Instance) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *Instance) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Instance) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Instance) GetInstanceType() InstanceType {
	if x != nil {
		return x.InstanceType
	}
	return InstanceType_INSTANCE_TYPE_UNKNOWN
}

func (x *Instance) GetDecisionLogging() bool {
	if x != nil && x.DecisionLogging != nil {
		return *x.DecisionLogging
	}
	return false
}

func (x *Instance) GetVersionHash() string {
	if x != nil {
		return x.VersionHash
	}
	return ""
}

var File_aserto_api_v2_instance_proto protoreflect.FileDescriptor

const file_aserto_api_v2_instance_proto_rawDesc = "" +
	"\n" +
	"\x1caserto/api/v2/instance.proto\x12\raserto.api.v2\x1a\x1baserto/options/v1/ids.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\x9a\x03\n" +
	"\bInstance\x12!\n" +
	"\tpolicy_id\x18\x01 \x01(\tB\x04\x90\xb5\x18\x04R\bpolicyId\x12\x14\n" +
	"\x05label\x18\x02 \x01(\tR\x05label\x12\x10\n" +
	"\x03tag\x18\x03 \x01(\tR\x03tag\x129\n" +
	"\n" +
	"created_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12@\n" +
	"\rinstance_type\x18\x06 \x01(\x0e2\x1b.aserto.api.v2.InstanceTypeR\finstanceType\x12.\n" +
	"\x10decision_logging\x18\a \x01(\bH\x00R\x0fdecisionLogging\x88\x01\x01\x12!\n" +
	"\fversion_hash\x18\b \x01(\tR\vversionHash:#\x92A \n" +
	"\x1e\xd2\x01\x05label\xd2\x01\x03tag\xd2\x01\rinstance_typeB\x13\n" +
	"\x11_decision_logging*f\n" +
	"\fInstanceType\x12\x19\n" +
	"\x15INSTANCE_TYPE_UNKNOWN\x10\x00\x12\x18\n" +
	"\x14INSTANCE_TYPE_HOSTED\x10\x01\x12!\n" +
	"\x1dINSTANCE_TYPE_EDGE_AUTHORIZER\x10\x02BAZ/github.com/aserto-dev/go-grpc/aserto/api/v2;api\xaa\x02\rAserto.API.V2b\x06proto3"

var (
	file_aserto_api_v2_instance_proto_rawDescOnce sync.Once
	file_aserto_api_v2_instance_proto_rawDescData []byte
)

func file_aserto_api_v2_instance_proto_rawDescGZIP() []byte {
	file_aserto_api_v2_instance_proto_rawDescOnce.Do(func() {
		file_aserto_api_v2_instance_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_api_v2_instance_proto_rawDesc), len(file_aserto_api_v2_instance_proto_rawDesc)))
	})
	return file_aserto_api_v2_instance_proto_rawDescData
}

var file_aserto_api_v2_instance_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_api_v2_instance_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_aserto_api_v2_instance_proto_goTypes = []any{
	(InstanceType)(0),             // 0: aserto.api.v2.InstanceType
	(*Instance)(nil),              // 1: aserto.api.v2.Instance
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_aserto_api_v2_instance_proto_depIdxs = []int32{
	2, // 0: aserto.api.v2.Instance.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: aserto.api.v2.Instance.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: aserto.api.v2.Instance.instance_type:type_name -> aserto.api.v2.InstanceType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_aserto_api_v2_instance_proto_init() }
func file_aserto_api_v2_instance_proto_init() {
	if File_aserto_api_v2_instance_proto != nil {
		return
	}
	file_aserto_api_v2_instance_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_api_v2_instance_proto_rawDesc), len(file_aserto_api_v2_instance_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_api_v2_instance_proto_goTypes,
		DependencyIndexes: file_aserto_api_v2_instance_proto_depIdxs,
		EnumInfos:         file_aserto_api_v2_instance_proto_enumTypes,
		MessageInfos:      file_aserto_api_v2_instance_proto_msgTypes,
	}.Build()
	File_aserto_api_v2_instance_proto = out.File
	file_aserto_api_v2_instance_proto_goTypes = nil
	file_aserto_api_v2_instance_proto_depIdxs = nil
}
