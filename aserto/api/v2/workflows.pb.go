// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: aserto/api/v2/workflows.proto

package api

import (
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

type WorkflowState int32

const (
	WorkflowState_WORKFLOW_STATE_UNKNOWN WorkflowState = 0 // Value not set.
	WorkflowState_WORKFLOW_STATE_RUNNING WorkflowState = 1 // Workflow is running.
	WorkflowState_WORKFLOW_STATE_DONE    WorkflowState = 2 // Workflow has completed.
)

// Enum value maps for WorkflowState.
var (
	WorkflowState_name = map[int32]string{
		0: "WORKFLOW_STATE_UNKNOWN",
		1: "WORKFLOW_STATE_RUNNING",
		2: "WORKFLOW_STATE_DONE",
	}
	WorkflowState_value = map[string]int32{
		"WORKFLOW_STATE_UNKNOWN": 0,
		"WORKFLOW_STATE_RUNNING": 1,
		"WORKFLOW_STATE_DONE":    2,
	}
)

func (x WorkflowState) Enum() *WorkflowState {
	p := new(WorkflowState)
	*p = x
	return p
}

func (x WorkflowState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkflowState) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_api_v2_workflows_proto_enumTypes[0].Descriptor()
}

func (WorkflowState) Type() protoreflect.EnumType {
	return &file_aserto_api_v2_workflows_proto_enumTypes[0]
}

func (x WorkflowState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkflowState.Descriptor instead.
func (WorkflowState) EnumDescriptor() ([]byte, []int) {
	return file_aserto_api_v2_workflows_proto_rawDescGZIP(), []int{0}
}

type ActivityResponseState int32

const (
	ActivityResponseState_ACTIVITY_RESPONSE_STATE_UNKNOWN   ActivityResponseState = 0 // Value not set.
	ActivityResponseState_ACTIVITY_RESPONSE_STATE_CREATED   ActivityResponseState = 1 // Resouce managed by the activity was created.
	ActivityResponseState_ACTIVITY_RESPONSE_STATE_DELETED   ActivityResponseState = 2 // Resouce managed by the activity was deleted.
	ActivityResponseState_ACTIVITY_RESPONSE_STATE_UPDATED   ActivityResponseState = 3 // Resouce managed by the activity was updated.
	ActivityResponseState_ACTIVITY_RESPONSE_STATE_UNCHANGED ActivityResponseState = 4 // Resouce managed by the activity was unchanged.
)

// Enum value maps for ActivityResponseState.
var (
	ActivityResponseState_name = map[int32]string{
		0: "ACTIVITY_RESPONSE_STATE_UNKNOWN",
		1: "ACTIVITY_RESPONSE_STATE_CREATED",
		2: "ACTIVITY_RESPONSE_STATE_DELETED",
		3: "ACTIVITY_RESPONSE_STATE_UPDATED",
		4: "ACTIVITY_RESPONSE_STATE_UNCHANGED",
	}
	ActivityResponseState_value = map[string]int32{
		"ACTIVITY_RESPONSE_STATE_UNKNOWN":   0,
		"ACTIVITY_RESPONSE_STATE_CREATED":   1,
		"ACTIVITY_RESPONSE_STATE_DELETED":   2,
		"ACTIVITY_RESPONSE_STATE_UPDATED":   3,
		"ACTIVITY_RESPONSE_STATE_UNCHANGED": 4,
	}
)

func (x ActivityResponseState) Enum() *ActivityResponseState {
	p := new(ActivityResponseState)
	*p = x
	return p
}

func (x ActivityResponseState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActivityResponseState) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_api_v2_workflows_proto_enumTypes[1].Descriptor()
}

func (ActivityResponseState) Type() protoreflect.EnumType {
	return &file_aserto_api_v2_workflows_proto_enumTypes[1]
}

func (x ActivityResponseState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActivityResponseState.Descriptor instead.
func (ActivityResponseState) EnumDescriptor() ([]byte, []int) {
	return file_aserto_api_v2_workflows_proto_rawDescGZIP(), []int{1}
}

type WorkflowRunType int32

const (
	WorkflowRunType_WORKFLOW_RUN_TYPE_UNKNOWN WorkflowRunType = 0 // Value not set.
	WorkflowRunType_WORKFLOW_RUN_TYPE_NOOP    WorkflowRunType = 1 // Workflow run is a no-op.
	WorkflowRunType_WORKFLOW_RUN_TYPE_EXECUTE WorkflowRunType = 2 // Workflow run is set to execute.
)

// Enum value maps for WorkflowRunType.
var (
	WorkflowRunType_name = map[int32]string{
		0: "WORKFLOW_RUN_TYPE_UNKNOWN",
		1: "WORKFLOW_RUN_TYPE_NOOP",
		2: "WORKFLOW_RUN_TYPE_EXECUTE",
	}
	WorkflowRunType_value = map[string]int32{
		"WORKFLOW_RUN_TYPE_UNKNOWN": 0,
		"WORKFLOW_RUN_TYPE_NOOP":    1,
		"WORKFLOW_RUN_TYPE_EXECUTE": 2,
	}
)

func (x WorkflowRunType) Enum() *WorkflowRunType {
	p := new(WorkflowRunType)
	*p = x
	return p
}

func (x WorkflowRunType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkflowRunType) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_api_v2_workflows_proto_enumTypes[2].Descriptor()
}

func (WorkflowRunType) Type() protoreflect.EnumType {
	return &file_aserto_api_v2_workflows_proto_enumTypes[2]
}

func (x WorkflowRunType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkflowRunType.Descriptor instead.
func (WorkflowRunType) EnumDescriptor() ([]byte, []int) {
	return file_aserto_api_v2_workflows_proto_rawDescGZIP(), []int{2}
}

type WorkflowEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowId   string           `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	RunId        string           `protobuf:"bytes,2,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	WorkflowName string           `protobuf:"bytes,3,opt,name=workflow_name,json=workflowName,proto3" json:"workflow_name,omitempty"`
	Status       WorkflowState    `protobuf:"varint,4,opt,name=status,proto3,enum=aserto.api.v2.WorkflowState" json:"status,omitempty"`
	Event        *structpb.Struct `protobuf:"bytes,5,opt,name=event,proto3" json:"event,omitempty"`
	Message      string           `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *WorkflowEvent) Reset() {
	*x = WorkflowEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v2_workflows_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowEvent) ProtoMessage() {}

func (x *WorkflowEvent) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v2_workflows_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowEvent.ProtoReflect.Descriptor instead.
func (*WorkflowEvent) Descriptor() ([]byte, []int) {
	return file_aserto_api_v2_workflows_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowEvent) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *WorkflowEvent) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *WorkflowEvent) GetWorkflowName() string {
	if x != nil {
		return x.WorkflowName
	}
	return ""
}

func (x *WorkflowEvent) GetStatus() WorkflowState {
	if x != nil {
		return x.Status
	}
	return WorkflowState_WORKFLOW_STATE_UNKNOWN
}

func (x *WorkflowEvent) GetEvent() *structpb.Struct {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *WorkflowEvent) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ActivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activity string                `protobuf:"bytes,1,opt,name=activity,proto3" json:"activity,omitempty"`
	State    ActivityResponseState `protobuf:"varint,2,opt,name=state,proto3,enum=aserto.api.v2.ActivityResponseState" json:"state,omitempty"`
	Response *structpb.Value       `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *ActivityResponse) Reset() {
	*x = ActivityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v2_workflows_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityResponse) ProtoMessage() {}

func (x *ActivityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v2_workflows_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityResponse.ProtoReflect.Descriptor instead.
func (*ActivityResponse) Descriptor() ([]byte, []int) {
	return file_aserto_api_v2_workflows_proto_rawDescGZIP(), []int{1}
}

func (x *ActivityResponse) GetActivity() string {
	if x != nil {
		return x.Activity
	}
	return ""
}

func (x *ActivityResponse) GetState() ActivityResponseState {
	if x != nil {
		return x.State
	}
	return ActivityResponseState_ACTIVITY_RESPONSE_STATE_UNKNOWN
}

func (x *ActivityResponse) GetResponse() *structpb.Value {
	if x != nil {
		return x.Response
	}
	return nil
}

type WorkflowOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunType WorkflowRunType `protobuf:"varint,1,opt,name=run_type,json=runType,proto3,enum=aserto.api.v2.WorkflowRunType" json:"run_type,omitempty"`
}

func (x *WorkflowOptions) Reset() {
	*x = WorkflowOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v2_workflows_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowOptions) ProtoMessage() {}

func (x *WorkflowOptions) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v2_workflows_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowOptions.ProtoReflect.Descriptor instead.
func (*WorkflowOptions) Descriptor() ([]byte, []int) {
	return file_aserto_api_v2_workflows_proto_rawDescGZIP(), []int{2}
}

func (x *WorkflowOptions) GetRunType() WorkflowRunType {
	if x != nil {
		return x.RunType
	}
	return WorkflowRunType_WORKFLOW_RUN_TYPE_UNKNOWN
}

type WorkflowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workflow          string              `protobuf:"bytes,1,opt,name=workflow,proto3" json:"workflow,omitempty"`
	Options           *WorkflowOptions    `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	ActivityResult    []*ActivityResponse `protobuf:"bytes,3,rep,name=activity_result,json=activityResult,proto3" json:"activity_result,omitempty"`
	WorkflowResponses []*WorkflowResponse `protobuf:"bytes,4,rep,name=workflow_responses,json=workflowResponses,proto3" json:"workflow_responses,omitempty"`
}

func (x *WorkflowResponse) Reset() {
	*x = WorkflowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_api_v2_workflows_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowResponse) ProtoMessage() {}

func (x *WorkflowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_api_v2_workflows_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowResponse.ProtoReflect.Descriptor instead.
func (*WorkflowResponse) Descriptor() ([]byte, []int) {
	return file_aserto_api_v2_workflows_proto_rawDescGZIP(), []int{3}
}

func (x *WorkflowResponse) GetWorkflow() string {
	if x != nil {
		return x.Workflow
	}
	return ""
}

func (x *WorkflowResponse) GetOptions() *WorkflowOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *WorkflowResponse) GetActivityResult() []*ActivityResponse {
	if x != nil {
		return x.ActivityResult
	}
	return nil
}

func (x *WorkflowResponse) GetWorkflowResponses() []*WorkflowResponse {
	if x != nil {
		return x.WorkflowResponses
	}
	return nil
}

var File_aserto_api_v2_workflows_proto protoreflect.FileDescriptor

var file_aserto_api_v2_workflows_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01, 0x0a,
	0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x10, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x0a, 0x0f, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39,
	0x0a, 0x08, 0x72, 0x75, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x07, 0x72, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x82, 0x02, 0x0a, 0x10, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x38, 0x0a, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x48, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x4e,
	0x0a, 0x12, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x11, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2a, 0x60,
	0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x16, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x57,
	0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x55,
	0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x57, 0x4f, 0x52, 0x4b, 0x46,
	0x4c, 0x4f, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02,
	0x2a, 0xd2, 0x01, 0x0a, 0x15, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x23, 0x0a, 0x1f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x50,
	0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59,
	0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x25,
	0x0a, 0x21, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f,
	0x4e, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x44, 0x10, 0x04, 0x2a, 0x6b, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x52, 0x75, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x57, 0x4f, 0x52, 0x4b,
	0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x52, 0x55, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x57, 0x4f, 0x52, 0x4b, 0x46,
	0x4c, 0x4f, 0x57, 0x5f, 0x52, 0x55, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4f,
	0x50, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f,
	0x52, 0x55, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45,
	0x10, 0x02, 0x42, 0x41, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x32, 0x3b, 0x61, 0x70, 0x69, 0xaa, 0x02, 0x0d, 0x41, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x41,
	0x50, 0x49, 0x2e, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_api_v2_workflows_proto_rawDescOnce sync.Once
	file_aserto_api_v2_workflows_proto_rawDescData = file_aserto_api_v2_workflows_proto_rawDesc
)

func file_aserto_api_v2_workflows_proto_rawDescGZIP() []byte {
	file_aserto_api_v2_workflows_proto_rawDescOnce.Do(func() {
		file_aserto_api_v2_workflows_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_api_v2_workflows_proto_rawDescData)
	})
	return file_aserto_api_v2_workflows_proto_rawDescData
}

var file_aserto_api_v2_workflows_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_aserto_api_v2_workflows_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_aserto_api_v2_workflows_proto_goTypes = []any{
	(WorkflowState)(0),         // 0: aserto.api.v2.WorkflowState
	(ActivityResponseState)(0), // 1: aserto.api.v2.ActivityResponseState
	(WorkflowRunType)(0),       // 2: aserto.api.v2.WorkflowRunType
	(*WorkflowEvent)(nil),      // 3: aserto.api.v2.WorkflowEvent
	(*ActivityResponse)(nil),   // 4: aserto.api.v2.ActivityResponse
	(*WorkflowOptions)(nil),    // 5: aserto.api.v2.WorkflowOptions
	(*WorkflowResponse)(nil),   // 6: aserto.api.v2.WorkflowResponse
	(*structpb.Struct)(nil),    // 7: google.protobuf.Struct
	(*structpb.Value)(nil),     // 8: google.protobuf.Value
}
var file_aserto_api_v2_workflows_proto_depIdxs = []int32{
	0, // 0: aserto.api.v2.WorkflowEvent.status:type_name -> aserto.api.v2.WorkflowState
	7, // 1: aserto.api.v2.WorkflowEvent.event:type_name -> google.protobuf.Struct
	1, // 2: aserto.api.v2.ActivityResponse.state:type_name -> aserto.api.v2.ActivityResponseState
	8, // 3: aserto.api.v2.ActivityResponse.response:type_name -> google.protobuf.Value
	2, // 4: aserto.api.v2.WorkflowOptions.run_type:type_name -> aserto.api.v2.WorkflowRunType
	5, // 5: aserto.api.v2.WorkflowResponse.options:type_name -> aserto.api.v2.WorkflowOptions
	4, // 6: aserto.api.v2.WorkflowResponse.activity_result:type_name -> aserto.api.v2.ActivityResponse
	6, // 7: aserto.api.v2.WorkflowResponse.workflow_responses:type_name -> aserto.api.v2.WorkflowResponse
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_aserto_api_v2_workflows_proto_init() }
func file_aserto_api_v2_workflows_proto_init() {
	if File_aserto_api_v2_workflows_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_api_v2_workflows_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*WorkflowEvent); i {
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
		file_aserto_api_v2_workflows_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ActivityResponse); i {
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
		file_aserto_api_v2_workflows_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*WorkflowOptions); i {
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
		file_aserto_api_v2_workflows_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*WorkflowResponse); i {
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
			RawDescriptor: file_aserto_api_v2_workflows_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_api_v2_workflows_proto_goTypes,
		DependencyIndexes: file_aserto_api_v2_workflows_proto_depIdxs,
		EnumInfos:         file_aserto_api_v2_workflows_proto_enumTypes,
		MessageInfos:      file_aserto_api_v2_workflows_proto_msgTypes,
	}.Build()
	File_aserto_api_v2_workflows_proto = out.File
	file_aserto_api_v2_workflows_proto_rawDesc = nil
	file_aserto_api_v2_workflows_proto_goTypes = nil
	file_aserto_api_v2_workflows_proto_depIdxs = nil
}
