// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/api/v1/invite.proto

package api

import (
	_ "github.com/aserto-dev/go-grpc/aserto/options/v1"
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

type InviteStatus int32

const (
	InviteStatus_INVITE_STATUS_UNKNOWN  InviteStatus = 0 // Unknown invite status
	InviteStatus_INVITE_STATUS_ACCEPTED InviteStatus = 1 // Normal attribute
	InviteStatus_INVITE_STATUS_DECLINED InviteStatus = 2 // User declined the invitation
	InviteStatus_INVITE_STATUS_EXPIRED  InviteStatus = 3 // Invitation expired
	InviteStatus_INVITE_STATUS_CANCELED InviteStatus = 4 // Invitation was canceled
	InviteStatus_INVITE_STATUS_ACTIVE   InviteStatus = 5 // Invitation is active
)

// Enum value maps for InviteStatus.
var (
	InviteStatus_name = map[int32]string{
		0: "INVITE_STATUS_UNKNOWN",
		1: "INVITE_STATUS_ACCEPTED",
		2: "INVITE_STATUS_DECLINED",
		3: "INVITE_STATUS_EXPIRED",
		4: "INVITE_STATUS_CANCELED",
		5: "INVITE_STATUS_ACTIVE",
	}
	InviteStatus_value = map[string]int32{
		"INVITE_STATUS_UNKNOWN":  0,
		"INVITE_STATUS_ACCEPTED": 1,
		"INVITE_STATUS_DECLINED": 2,
		"INVITE_STATUS_EXPIRED":  3,
		"INVITE_STATUS_CANCELED": 4,
		"INVITE_STATUS_ACTIVE":   5,
	}
)

func (x InviteStatus) Enum() *InviteStatus {
	p := new(InviteStatus)
	*p = x
	return p
}

func (x InviteStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InviteStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_api_v1_invite_proto_enumTypes[0].Descriptor()
}

func (InviteStatus) Type() protoreflect.EnumType {
	return &file_aserto_api_v1_invite_proto_enumTypes[0]
}

func (x InviteStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InviteStatus.Descriptor instead.
func (InviteStatus) EnumDescriptor() ([]byte, []int) {
	return file_aserto_api_v1_invite_proto_rawDescGZIP(), []int{0}
}

type Invite struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                          // Unique ID of the invite
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`                                    // Email of the invitee
	Status        InviteStatus           `protobuf:"varint,3,opt,name=status,proto3,enum=aserto.api.v1.InviteStatus" json:"status,omitempty"` // Whether the invite was accepted/declined
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`           // When was the invite created?
	RespondedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=responded_at,json=respondedAt,proto3" json:"responded_at,omitempty"`     // When was the invite accepted/declined?
	InvitedBy     string                 `protobuf:"bytes,6,opt,name=invited_by,json=invitedBy,proto3" json:"invited_by,omitempty"`           // Account ID for the inviter
	Role          string                 `protobuf:"bytes,7,opt,name=role,proto3" json:"role,omitempty"`                                      // Role assumed by the invitee on acceptance
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Invite) Reset() {
	*x = Invite{}
	mi := &file_aserto_api_v1_invite_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Invite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invite) ProtoMessage() {}

func (x *Invite) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v1_invite_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invite.ProtoReflect.Descriptor instead.
func (*Invite) Descriptor() ([]byte, []int) {
	return file_aserto_api_v1_invite_proto_rawDescGZIP(), []int{0}
}

func (x *Invite) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Invite) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Invite) GetStatus() InviteStatus {
	if x != nil {
		return x.Status
	}
	return InviteStatus_INVITE_STATUS_UNKNOWN
}

func (x *Invite) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Invite) GetRespondedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RespondedAt
	}
	return nil
}

func (x *Invite) GetInvitedBy() string {
	if x != nil {
		return x.InvitedBy
	}
	return ""
}

func (x *Invite) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_aserto_api_v1_invite_proto protoreflect.FileDescriptor

const file_aserto_api_v1_invite_proto_rawDesc = "" +
	"\n" +
	"\x1aaserto/api/v1/invite.proto\x12\raserto.api.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1baserto/options/v1/ids.proto\"\x96\x02\n" +
	"\x06Invite\x12\x14\n" +
	"\x02id\x18\x01 \x01(\tB\x04\x90\xb5\x18\bR\x02id\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x123\n" +
	"\x06status\x18\x03 \x01(\x0e2\x1b.aserto.api.v1.InviteStatusR\x06status\x129\n" +
	"\n" +
	"created_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12=\n" +
	"\fresponded_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\vrespondedAt\x12\x1d\n" +
	"\n" +
	"invited_by\x18\x06 \x01(\tR\tinvitedBy\x12\x12\n" +
	"\x04role\x18\a \x01(\tR\x04role*\xb2\x01\n" +
	"\fInviteStatus\x12\x19\n" +
	"\x15INVITE_STATUS_UNKNOWN\x10\x00\x12\x1a\n" +
	"\x16INVITE_STATUS_ACCEPTED\x10\x01\x12\x1a\n" +
	"\x16INVITE_STATUS_DECLINED\x10\x02\x12\x19\n" +
	"\x15INVITE_STATUS_EXPIRED\x10\x03\x12\x1a\n" +
	"\x16INVITE_STATUS_CANCELED\x10\x04\x12\x18\n" +
	"\x14INVITE_STATUS_ACTIVE\x10\x05BAZ/github.com/aserto-dev/go-grpc/aserto/api/v1;api\xaa\x02\rAserto.API.V1b\x06proto3"

var (
	file_aserto_api_v1_invite_proto_rawDescOnce sync.Once
	file_aserto_api_v1_invite_proto_rawDescData []byte
)

func file_aserto_api_v1_invite_proto_rawDescGZIP() []byte {
	file_aserto_api_v1_invite_proto_rawDescOnce.Do(func() {
		file_aserto_api_v1_invite_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_api_v1_invite_proto_rawDesc), len(file_aserto_api_v1_invite_proto_rawDesc)))
	})
	return file_aserto_api_v1_invite_proto_rawDescData
}

var file_aserto_api_v1_invite_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_api_v1_invite_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_aserto_api_v1_invite_proto_goTypes = []any{
	(InviteStatus)(0),             // 0: aserto.api.v1.InviteStatus
	(*Invite)(nil),                // 1: aserto.api.v1.Invite
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_aserto_api_v1_invite_proto_depIdxs = []int32{
	0, // 0: aserto.api.v1.Invite.status:type_name -> aserto.api.v1.InviteStatus
	2, // 1: aserto.api.v1.Invite.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: aserto.api.v1.Invite.responded_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_aserto_api_v1_invite_proto_init() }
func file_aserto_api_v1_invite_proto_init() {
	if File_aserto_api_v1_invite_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_api_v1_invite_proto_rawDesc), len(file_aserto_api_v1_invite_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_api_v1_invite_proto_goTypes,
		DependencyIndexes: file_aserto_api_v1_invite_proto_depIdxs,
		EnumInfos:         file_aserto_api_v1_invite_proto_enumTypes,
		MessageInfos:      file_aserto_api_v1_invite_proto_msgTypes,
	}.Build()
	File_aserto_api_v1_invite_proto = out.File
	file_aserto_api_v1_invite_proto_goTypes = nil
	file_aserto_api_v1_invite_proto_depIdxs = nil
}
