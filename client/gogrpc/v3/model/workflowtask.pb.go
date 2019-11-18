// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/workflowtask.proto

package model // import "github.com/mactaggart/conductor/client/gogrpc/v3/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkflowTask struct {
	Name                           string                                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TaskReferenceName              string                                    `protobuf:"bytes,2,opt,name=task_reference_name,json=taskReferenceName,proto3" json:"task_reference_name,omitempty"`
	Description                    string                                    `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	InputParameters                map[string]*any.Any                       `protobuf:"bytes,4,rep,name=input_parameters,json=inputParameters,proto3" json:"input_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Type                           string                                    `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	DynamicTaskNameParam           string                                    `protobuf:"bytes,6,opt,name=dynamic_task_name_param,json=dynamicTaskNameParam,proto3" json:"dynamic_task_name_param,omitempty"`
	CaseValueParam                 string                                    `protobuf:"bytes,7,opt,name=case_value_param,json=caseValueParam,proto3" json:"case_value_param,omitempty"`
	CaseExpression                 string                                    `protobuf:"bytes,8,opt,name=case_expression,json=caseExpression,proto3" json:"case_expression,omitempty"`
	ScriptExpression               string                                    `protobuf:"bytes,22,opt,name=script_expression,json=scriptExpression,proto3" json:"script_expression,omitempty"`
	DecisionCases                  map[string]*WorkflowTask_WorkflowTaskList `protobuf:"bytes,9,rep,name=decision_cases,json=decisionCases,proto3" json:"decision_cases,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DynamicForkTasksParam          string                                    `protobuf:"bytes,10,opt,name=dynamic_fork_tasks_param,json=dynamicForkTasksParam,proto3" json:"dynamic_fork_tasks_param,omitempty"`
	DynamicForkTasksInputParamName string                                    `protobuf:"bytes,11,opt,name=dynamic_fork_tasks_input_param_name,json=dynamicForkTasksInputParamName,proto3" json:"dynamic_fork_tasks_input_param_name,omitempty"`
	DefaultCase                    []*WorkflowTask                           `protobuf:"bytes,12,rep,name=default_case,json=defaultCase,proto3" json:"default_case,omitempty"`
	ForkTasks                      []*WorkflowTask_WorkflowTaskList          `protobuf:"bytes,13,rep,name=fork_tasks,json=forkTasks,proto3" json:"fork_tasks,omitempty"`
	StartDelay                     int32                                     `protobuf:"varint,14,opt,name=start_delay,json=startDelay,proto3" json:"start_delay,omitempty"`
	SubWorkflowParam               *SubWorkflowParams                        `protobuf:"bytes,15,opt,name=sub_workflow_param,json=subWorkflowParam,proto3" json:"sub_workflow_param,omitempty"`
	JoinOn                         []string                                  `protobuf:"bytes,16,rep,name=join_on,json=joinOn,proto3" json:"join_on,omitempty"`
	Sink                           string                                    `protobuf:"bytes,17,opt,name=sink,proto3" json:"sink,omitempty"`
	Optional                       bool                                      `protobuf:"varint,18,opt,name=optional,proto3" json:"optional,omitempty"`
	TaskDefinition                 *TaskDef                                  `protobuf:"bytes,19,opt,name=task_definition,json=taskDefinition,proto3" json:"task_definition,omitempty"`
	RateLimited                    bool                                      `protobuf:"varint,20,opt,name=rate_limited,json=rateLimited,proto3" json:"rate_limited,omitempty"`
	DefaultExclusiveJoinTask       []string                                  `protobuf:"bytes,21,rep,name=default_exclusive_join_task,json=defaultExclusiveJoinTask,proto3" json:"default_exclusive_join_task,omitempty"`
	AsyncComplete                  bool                                      `protobuf:"varint,23,opt,name=async_complete,json=asyncComplete,proto3" json:"async_complete,omitempty"`
	LoopCondition                  string                                    `protobuf:"bytes,24,opt,name=loop_condition,json=loopCondition,proto3" json:"loop_condition,omitempty"`
	LoopOver                       []*WorkflowTask                           `protobuf:"bytes,25,rep,name=loop_over,json=loopOver,proto3" json:"loop_over,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}                                  `json:"-"`
	XXX_unrecognized               []byte                                    `json:"-"`
	XXX_sizecache                  int32                                     `json:"-"`
}

func (m *WorkflowTask) Reset()         { *m = WorkflowTask{} }
func (m *WorkflowTask) String() string { return proto.CompactTextString(m) }
func (*WorkflowTask) ProtoMessage()    {}
func (*WorkflowTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflowtask_111d53343898d8d9, []int{0}
}
func (m *WorkflowTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowTask.Unmarshal(m, b)
}
func (m *WorkflowTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowTask.Marshal(b, m, deterministic)
}
func (dst *WorkflowTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowTask.Merge(dst, src)
}
func (m *WorkflowTask) XXX_Size() int {
	return xxx_messageInfo_WorkflowTask.Size(m)
}
func (m *WorkflowTask) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowTask.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowTask proto.InternalMessageInfo

func (m *WorkflowTask) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkflowTask) GetTaskReferenceName() string {
	if m != nil {
		return m.TaskReferenceName
	}
	return ""
}

func (m *WorkflowTask) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *WorkflowTask) GetInputParameters() map[string]*any.Any {
	if m != nil {
		return m.InputParameters
	}
	return nil
}

func (m *WorkflowTask) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *WorkflowTask) GetDynamicTaskNameParam() string {
	if m != nil {
		return m.DynamicTaskNameParam
	}
	return ""
}

func (m *WorkflowTask) GetCaseValueParam() string {
	if m != nil {
		return m.CaseValueParam
	}
	return ""
}

func (m *WorkflowTask) GetCaseExpression() string {
	if m != nil {
		return m.CaseExpression
	}
	return ""
}

func (m *WorkflowTask) GetScriptExpression() string {
	if m != nil {
		return m.ScriptExpression
	}
	return ""
}

func (m *WorkflowTask) GetDecisionCases() map[string]*WorkflowTask_WorkflowTaskList {
	if m != nil {
		return m.DecisionCases
	}
	return nil
}

func (m *WorkflowTask) GetDynamicForkTasksParam() string {
	if m != nil {
		return m.DynamicForkTasksParam
	}
	return ""
}

func (m *WorkflowTask) GetDynamicForkTasksInputParamName() string {
	if m != nil {
		return m.DynamicForkTasksInputParamName
	}
	return ""
}

func (m *WorkflowTask) GetDefaultCase() []*WorkflowTask {
	if m != nil {
		return m.DefaultCase
	}
	return nil
}

func (m *WorkflowTask) GetForkTasks() []*WorkflowTask_WorkflowTaskList {
	if m != nil {
		return m.ForkTasks
	}
	return nil
}

func (m *WorkflowTask) GetStartDelay() int32 {
	if m != nil {
		return m.StartDelay
	}
	return 0
}

func (m *WorkflowTask) GetSubWorkflowParam() *SubWorkflowParams {
	if m != nil {
		return m.SubWorkflowParam
	}
	return nil
}

func (m *WorkflowTask) GetJoinOn() []string {
	if m != nil {
		return m.JoinOn
	}
	return nil
}

func (m *WorkflowTask) GetSink() string {
	if m != nil {
		return m.Sink
	}
	return ""
}

func (m *WorkflowTask) GetOptional() bool {
	if m != nil {
		return m.Optional
	}
	return false
}

func (m *WorkflowTask) GetTaskDefinition() *TaskDef {
	if m != nil {
		return m.TaskDefinition
	}
	return nil
}

func (m *WorkflowTask) GetRateLimited() bool {
	if m != nil {
		return m.RateLimited
	}
	return false
}

func (m *WorkflowTask) GetDefaultExclusiveJoinTask() []string {
	if m != nil {
		return m.DefaultExclusiveJoinTask
	}
	return nil
}

func (m *WorkflowTask) GetAsyncComplete() bool {
	if m != nil {
		return m.AsyncComplete
	}
	return false
}

func (m *WorkflowTask) GetLoopCondition() string {
	if m != nil {
		return m.LoopCondition
	}
	return ""
}

func (m *WorkflowTask) GetLoopOver() []*WorkflowTask {
	if m != nil {
		return m.LoopOver
	}
	return nil
}

type WorkflowTask_WorkflowTaskList struct {
	Tasks                []*WorkflowTask `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *WorkflowTask_WorkflowTaskList) Reset()         { *m = WorkflowTask_WorkflowTaskList{} }
func (m *WorkflowTask_WorkflowTaskList) String() string { return proto.CompactTextString(m) }
func (*WorkflowTask_WorkflowTaskList) ProtoMessage()    {}
func (*WorkflowTask_WorkflowTaskList) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflowtask_111d53343898d8d9, []int{0, 0}
}
func (m *WorkflowTask_WorkflowTaskList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowTask_WorkflowTaskList.Unmarshal(m, b)
}
func (m *WorkflowTask_WorkflowTaskList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowTask_WorkflowTaskList.Marshal(b, m, deterministic)
}
func (dst *WorkflowTask_WorkflowTaskList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowTask_WorkflowTaskList.Merge(dst, src)
}
func (m *WorkflowTask_WorkflowTaskList) XXX_Size() int {
	return xxx_messageInfo_WorkflowTask_WorkflowTaskList.Size(m)
}
func (m *WorkflowTask_WorkflowTaskList) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowTask_WorkflowTaskList.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowTask_WorkflowTaskList proto.InternalMessageInfo

func (m *WorkflowTask_WorkflowTaskList) GetTasks() []*WorkflowTask {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func init() {
	proto.RegisterType((*WorkflowTask)(nil), "conductor.proto.WorkflowTask")
	proto.RegisterMapType((map[string]*WorkflowTask_WorkflowTaskList)(nil), "conductor.proto.WorkflowTask.DecisionCasesEntry")
	proto.RegisterMapType((map[string]*any.Any)(nil), "conductor.proto.WorkflowTask.InputParametersEntry")
	proto.RegisterType((*WorkflowTask_WorkflowTaskList)(nil), "conductor.proto.WorkflowTask.WorkflowTaskList")
}

func init() {
	proto.RegisterFile("model/workflowtask.proto", fileDescriptor_workflowtask_111d53343898d8d9)
}

var fileDescriptor_workflowtask_111d53343898d8d9 = []byte{
	// 819 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x7f, 0x6f, 0xdb, 0x36,
	0x10, 0x85, 0x9b, 0x26, 0x8d, 0xcf, 0x8e, 0xed, 0x30, 0xe9, 0xc2, 0xba, 0xe8, 0xe6, 0x75, 0x18,
	0x66, 0x6c, 0x80, 0x34, 0x24, 0xd8, 0x0f, 0x14, 0x18, 0xb0, 0x36, 0xc9, 0x86, 0x6d, 0xdd, 0x1a,
	0x68, 0xc3, 0x3a, 0x0c, 0x18, 0x04, 0x5a, 0xa2, 0x3c, 0xce, 0x12, 0x29, 0x90, 0x94, 0x1b, 0x7d,
	0x89, 0x7d, 0xe6, 0x81, 0x47, 0x29, 0x51, 0x9d, 0x20, 0xe8, 0xfe, 0x23, 0xdf, 0xbd, 0x7b, 0xbc,
	0x77, 0x3a, 0x52, 0x40, 0x0b, 0x95, 0xf2, 0x3c, 0x7c, 0xa3, 0xf4, 0x2a, 0xcb, 0xd5, 0x1b, 0xcb,
	0xcc, 0x2a, 0x28, 0xb5, 0xb2, 0x8a, 0x8c, 0x13, 0x25, 0xd3, 0x2a, 0xb1, 0x4a, 0x7b, 0x60, 0xfa,
	0x68, 0xa9, 0xd4, 0x32, 0xe7, 0x21, 0xee, 0x16, 0x55, 0x16, 0x32, 0x59, 0x37, 0xa1, 0x03, 0xaf,
	0xe2, 0xb2, 0x53, 0x9e, 0x35, 0xe0, 0x13, 0x0f, 0x9a, 0x6a, 0xd1, 0xaa, 0x97, 0x4c, 0xb3, 0xc2,
	0xf8, 0xf0, 0xd3, 0x7f, 0x87, 0x30, 0x7c, 0xdd, 0x04, 0x7e, 0x63, 0x66, 0x45, 0x08, 0xdc, 0x97,
	0xac, 0xe0, 0xb4, 0x37, 0xeb, 0xcd, 0xfb, 0x11, 0xae, 0x49, 0x00, 0x07, 0x4e, 0x34, 0xd6, 0x3c,
	0xe3, 0x9a, 0xcb, 0x84, 0xc7, 0x48, 0xb9, 0x87, 0x94, 0x7d, 0x17, 0x8a, 0xda, 0xc8, 0x2f, 0x8e,
	0x3f, 0x83, 0x41, 0xca, 0x4d, 0xa2, 0x45, 0x69, 0x85, 0x92, 0x74, 0x0b, 0x79, 0x5d, 0x88, 0xfc,
	0x05, 0x13, 0x21, 0xcb, 0xca, 0xc6, 0x58, 0x0c, 0xb7, 0x5c, 0x1b, 0x7a, 0x7f, 0xb6, 0x35, 0x1f,
	0x1c, 0x1f, 0x07, 0x1b, 0x8e, 0x83, 0x6e, 0x79, 0xc1, 0x0f, 0x2e, 0xeb, 0xe2, 0x2a, 0xe9, 0x5c,
	0x5a, 0x5d, 0x47, 0x63, 0xf1, 0x36, 0xea, 0x4c, 0xd8, 0xba, 0xe4, 0x74, 0xdb, 0x9b, 0x70, 0x6b,
	0xf2, 0x05, 0x1c, 0xa5, 0xb5, 0x64, 0x85, 0x48, 0x62, 0x34, 0xe3, 0x2c, 0xf8, 0xe3, 0xe9, 0x0e,
	0xd2, 0x0e, 0x9b, 0xb0, 0x3b, 0xc7, 0xd9, 0x40, 0x3d, 0x32, 0x87, 0x49, 0xc2, 0x0c, 0x8f, 0xd7,
	0x2c, 0xaf, 0x5a, 0xfe, 0x03, 0xe4, 0x8f, 0x1c, 0xfe, 0xbb, 0x83, 0x3d, 0xf3, 0x13, 0x18, 0x23,
	0x93, 0x5f, 0x96, 0x9a, 0x1b, 0xe3, 0x9c, 0xef, 0x5e, 0x13, 0xcf, 0xaf, 0x50, 0xf2, 0x19, 0xec,
	0xfb, 0x4e, 0x74, 0xa9, 0xef, 0x21, 0x75, 0xe2, 0x03, 0x1d, 0xf2, 0x6b, 0x18, 0xa5, 0x3c, 0x11,
	0x6e, 0x1d, 0x3b, 0x1d, 0x43, 0xfb, 0xd8, 0xa7, 0xcf, 0xef, 0xee, 0xd3, 0x59, 0x93, 0x73, 0xea,
	0x52, 0x7c, 0x97, 0xf6, 0xd2, 0x2e, 0x46, 0xbe, 0x02, 0xda, 0xf6, 0x23, 0x53, 0x7a, 0x85, 0x4d,
	0x31, 0x8d, 0x41, 0xc0, 0x62, 0x1e, 0x36, 0xf1, 0xef, 0x94, 0x5e, 0x39, 0x51, 0xe3, 0x7d, 0xfe,
	0x04, 0x1f, 0xdd, 0x92, 0xd8, 0xf9, 0x9c, 0x7e, 0x3a, 0x06, 0xa8, 0xf1, 0xfe, 0xa6, 0xc6, 0xf5,
	0x07, 0xc4, 0x51, 0xf9, 0x16, 0x86, 0x29, 0xcf, 0x58, 0x95, 0x5b, 0x74, 0x47, 0x87, 0x68, 0xee,
	0xc9, 0x9d, 0xe6, 0xdc, 0x28, 0x61, 0x8a, 0x33, 0x42, 0x7e, 0x06, 0xb8, 0x2e, 0x83, 0xee, 0x61,
	0x7e, 0x70, 0x77, 0x73, 0xba, 0x9b, 0x97, 0xc2, 0xd8, 0xa8, 0x9f, 0xb5, 0xe5, 0x91, 0x0f, 0x60,
	0x60, 0x2c, 0xd3, 0x36, 0x4e, 0x79, 0xce, 0x6a, 0x3a, 0x9a, 0xf5, 0xe6, 0xdb, 0x11, 0x20, 0x74,
	0xe6, 0x10, 0x72, 0x01, 0xc4, 0x54, 0x8b, 0xb8, 0xbd, 0x4d, 0x4d, 0xc7, 0xc6, 0xb3, 0xde, 0x7c,
	0x70, 0xfc, 0xf4, 0xc6, 0xb9, 0xbf, 0x56, 0x8b, 0xf6, 0x34, 0x34, 0x6d, 0xa2, 0x89, 0xd9, 0x80,
	0xc8, 0x11, 0x3c, 0xf8, 0x47, 0x09, 0x19, 0x2b, 0x49, 0x27, 0xb3, 0xad, 0x79, 0x3f, 0xda, 0x71,
	0xdb, 0x57, 0xd2, 0x8d, 0xb1, 0x11, 0x72, 0x45, 0xf7, 0xfd, 0x18, 0xbb, 0x35, 0x99, 0xc2, 0xae,
	0xc2, 0x3b, 0xc4, 0x72, 0x4a, 0x66, 0xbd, 0xf9, 0x6e, 0x74, 0xb5, 0x27, 0xcf, 0x61, 0x8c, 0xa3,
	0x9d, 0xf2, 0x4c, 0x48, 0x81, 0x77, 0xef, 0x00, 0xeb, 0xa2, 0x37, 0xea, 0x72, 0x66, 0xcf, 0x78,
	0x16, 0x8d, 0xac, 0x5f, 0x34, 0x7c, 0xf2, 0x21, 0x0c, 0x35, 0xb3, 0x3c, 0xce, 0x45, 0x21, 0x2c,
	0x4f, 0xe9, 0x21, 0x1e, 0x31, 0x70, 0xd8, 0x4b, 0x0f, 0x91, 0x6f, 0xe0, 0x71, 0xfb, 0xc9, 0xf8,
	0x65, 0x92, 0x57, 0x46, 0xac, 0x79, 0x8c, 0x06, 0x9c, 0x16, 0x7d, 0x88, 0x16, 0x68, 0x43, 0x39,
	0x6f, 0x19, 0x3f, 0x2a, 0x21, 0xf1, 0x81, 0xf9, 0x18, 0x46, 0xcc, 0xd4, 0x32, 0x89, 0x13, 0x55,
	0x94, 0x39, 0xb7, 0x9c, 0x1e, 0xe1, 0x19, 0x7b, 0x88, 0x9e, 0x36, 0xa0, 0xa3, 0xe5, 0x4a, 0x95,
	0xb1, 0x2b, 0xdc, 0x5b, 0xa1, 0xd8, 0x85, 0x3d, 0x87, 0x9e, 0xb6, 0x20, 0x79, 0x06, 0x7d, 0xa4,
	0xa9, 0x35, 0xd7, 0xf4, 0xd1, 0xbb, 0x0c, 0xcf, 0xae, 0xe3, 0xbf, 0x5a, 0x73, 0x3d, 0xfd, 0x1e,
	0x26, 0x9b, 0x93, 0x40, 0x4e, 0x60, 0xdb, 0x0f, 0x52, 0xef, 0x5d, 0xb4, 0x3c, 0x77, 0xfa, 0x07,
	0x1c, 0xde, 0xf6, 0x2e, 0x91, 0x09, 0x6c, 0xad, 0x78, 0xdd, 0x3c, 0xa5, 0x6e, 0x49, 0x3e, 0x85,
	0x6d, 0x7c, 0x48, 0xf0, 0xed, 0x1c, 0x1c, 0x1f, 0x06, 0xfe, 0x35, 0x0f, 0xda, 0xd7, 0x3c, 0x78,
	0x2e, 0xeb, 0xc8, 0x53, 0x9e, 0xdd, 0xfb, 0xba, 0x37, 0x2d, 0x81, 0xdc, 0xbc, 0xc9, 0xb7, 0xe8,
	0x9e, 0xbd, 0xad, 0xfb, 0x7f, 0xe7, 0xff, 0xfa, 0xc4, 0x17, 0x1c, 0x1e, 0x27, 0xaa, 0x08, 0x24,
	0xb7, 0x59, 0x2e, 0x2e, 0x37, 0x75, 0x5e, 0x8c, 0xba, 0xb9, 0x17, 0x8b, 0x3f, 0xbf, 0x5c, 0x0a,
	0xfb, 0x77, 0xb5, 0x08, 0x12, 0x55, 0x84, 0x05, 0x4b, 0x2c, 0x5b, 0x2e, 0x99, 0xb6, 0xe1, 0x55,
	0x5a, 0x98, 0xe4, 0x82, 0x4b, 0x1b, 0x2e, 0xd5, 0x52, 0x97, 0x49, 0xb8, 0x3e, 0x09, 0xf1, 0x7f,
	0xb4, 0xd8, 0x41, 0xb9, 0x93, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x60, 0x6e, 0x38, 0xfa,
	0x06, 0x00, 0x00,
}
