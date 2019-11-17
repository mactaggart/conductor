// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/workflowsummary.proto

package model // import "github.com/netflix/conductor/client/gogrpc/conductor/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkflowSummary struct {
	WorkflowType                     string                  `protobuf:"bytes,1,opt,name=workflow_type,json=workflowType,proto3" json:"workflow_type,omitempty"`
	Version                          int32                   `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	WorkflowId                       string                  `protobuf:"bytes,3,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	CorrelationId                    string                  `protobuf:"bytes,4,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	StartTime                        string                  `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	UpdateTime                       string                  `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	EndTime                          string                  `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Status                           Workflow_WorkflowStatus `protobuf:"varint,8,opt,name=status,proto3,enum=conductor.proto.Workflow_WorkflowStatus" json:"status,omitempty"`
	Input                            string                  `protobuf:"bytes,9,opt,name=input,proto3" json:"input,omitempty"`
	Output                           string                  `protobuf:"bytes,10,opt,name=output,proto3" json:"output,omitempty"`
	ReasonForIncompletion            string                  `protobuf:"bytes,11,opt,name=reason_for_incompletion,json=reasonForIncompletion,proto3" json:"reason_for_incompletion,omitempty"`
	ExecutionTime                    int64                   `protobuf:"varint,12,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	Event                            string                  `protobuf:"bytes,13,opt,name=event,proto3" json:"event,omitempty"`
	FailedReferenceTaskNames         string                  `protobuf:"bytes,14,opt,name=failed_reference_task_names,json=failedReferenceTaskNames,proto3" json:"failed_reference_task_names,omitempty"`
	ExternalInputPayloadStoragePath  string                  `protobuf:"bytes,15,opt,name=external_input_payload_storage_path,json=externalInputPayloadStoragePath,proto3" json:"external_input_payload_storage_path,omitempty"`
	ExternalOutputPayloadStoragePath string                  `protobuf:"bytes,16,opt,name=external_output_payload_storage_path,json=externalOutputPayloadStoragePath,proto3" json:"external_output_payload_storage_path,omitempty"`
	Priority                         int32                   `protobuf:"varint,17,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}                `json:"-"`
	XXX_unrecognized                 []byte                  `json:"-"`
	XXX_sizecache                    int32                   `json:"-"`
}

func (m *WorkflowSummary) Reset()         { *m = WorkflowSummary{} }
func (m *WorkflowSummary) String() string { return proto.CompactTextString(m) }
func (*WorkflowSummary) ProtoMessage()    {}
func (*WorkflowSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflowsummary_ec943c4ffc11105d, []int{0}
}
func (m *WorkflowSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowSummary.Unmarshal(m, b)
}
func (m *WorkflowSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowSummary.Marshal(b, m, deterministic)
}
func (dst *WorkflowSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowSummary.Merge(dst, src)
}
func (m *WorkflowSummary) XXX_Size() int {
	return xxx_messageInfo_WorkflowSummary.Size(m)
}
func (m *WorkflowSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowSummary.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowSummary proto.InternalMessageInfo

func (m *WorkflowSummary) GetWorkflowType() string {
	if m != nil {
		return m.WorkflowType
	}
	return ""
}

func (m *WorkflowSummary) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *WorkflowSummary) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *WorkflowSummary) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *WorkflowSummary) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *WorkflowSummary) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *WorkflowSummary) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *WorkflowSummary) GetStatus() Workflow_WorkflowStatus {
	if m != nil {
		return m.Status
	}
	return Workflow_RUNNING
}

func (m *WorkflowSummary) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *WorkflowSummary) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *WorkflowSummary) GetReasonForIncompletion() string {
	if m != nil {
		return m.ReasonForIncompletion
	}
	return ""
}

func (m *WorkflowSummary) GetExecutionTime() int64 {
	if m != nil {
		return m.ExecutionTime
	}
	return 0
}

func (m *WorkflowSummary) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *WorkflowSummary) GetFailedReferenceTaskNames() string {
	if m != nil {
		return m.FailedReferenceTaskNames
	}
	return ""
}

func (m *WorkflowSummary) GetExternalInputPayloadStoragePath() string {
	if m != nil {
		return m.ExternalInputPayloadStoragePath
	}
	return ""
}

func (m *WorkflowSummary) GetExternalOutputPayloadStoragePath() string {
	if m != nil {
		return m.ExternalOutputPayloadStoragePath
	}
	return ""
}

func (m *WorkflowSummary) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterType((*WorkflowSummary)(nil), "conductor.proto.WorkflowSummary")
}

func init() {
	proto.RegisterFile("model/workflowsummary.proto", fileDescriptor_workflowsummary_ec943c4ffc11105d)
}

var fileDescriptor_workflowsummary_ec943c4ffc11105d = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x5f, 0x8b, 0xda, 0x4c,
	0x14, 0xc6, 0xf1, 0xdd, 0xd7, 0x7f, 0xb3, 0xab, 0x76, 0x07, 0xdb, 0x4e, 0x57, 0xca, 0x4a, 0xb7,
	0x05, 0xaf, 0x14, 0x5a, 0xe8, 0x45, 0xa1, 0x50, 0xf6, 0xa2, 0x20, 0x94, 0xad, 0xa8, 0x50, 0xe8,
	0x4d, 0x18, 0x93, 0xa3, 0x0e, 0x26, 0x73, 0xc2, 0xe4, 0x64, 0x57, 0x3f, 0x66, 0xbf, 0x51, 0xc9,
	0x19, 0x13, 0xac, 0xec, 0x5d, 0xce, 0xf3, 0xfc, 0xce, 0x33, 0x39, 0x27, 0x13, 0x31, 0x48, 0x30,
	0x82, 0x78, 0xf2, 0x84, 0x6e, 0xb7, 0x8e, 0xf1, 0x29, 0xcb, 0x93, 0x44, 0xbb, 0xc3, 0x38, 0x75,
	0x48, 0x28, 0x7b, 0x21, 0xda, 0x28, 0x0f, 0x09, 0x9d, 0x17, 0x6e, 0xfa, 0xff, 0xd2, 0x5e, 0x7d,
	0xf7, 0xa7, 0x2e, 0x7a, 0xbf, 0x8e, 0xd2, 0xc2, 0x07, 0xc8, 0x3b, 0xd1, 0x29, 0xa9, 0x80, 0x0e,
	0x29, 0xa8, 0xda, 0xb0, 0x36, 0x6a, 0xcf, 0xaf, 0x4a, 0x71, 0x79, 0x48, 0x41, 0x2a, 0xd1, 0x7c,
	0x04, 0x97, 0x19, 0xb4, 0xea, 0xbf, 0x61, 0x6d, 0x54, 0x9f, 0x97, 0xa5, 0xbc, 0x15, 0x97, 0x55,
	0xbb, 0x89, 0xd4, 0x05, 0x37, 0x8b, 0x52, 0x9a, 0x46, 0xf2, 0x83, 0xe8, 0x86, 0xe8, 0x1c, 0xc4,
	0x9a, 0x0c, 0xda, 0x82, 0xf9, 0x9f, 0x99, 0xce, 0x89, 0x3a, 0x8d, 0xe4, 0x5b, 0x21, 0x32, 0xd2,
	0x8e, 0x02, 0x32, 0x09, 0xa8, 0x3a, 0x23, 0x6d, 0x56, 0x96, 0x26, 0x81, 0xe2, 0x98, 0x3c, 0x8d,
	0x34, 0x81, 0xf7, 0x1b, 0xfe, 0x18, 0x2f, 0x31, 0xf0, 0x46, 0xb4, 0xc0, 0x46, 0xde, 0x6d, 0xb2,
	0xdb, 0x04, 0x1b, 0xb1, 0xf5, 0x4d, 0x34, 0x32, 0xd2, 0x94, 0x67, 0xaa, 0x35, 0xac, 0x8d, 0xba,
	0x1f, 0x47, 0xe3, 0xb3, 0x6d, 0x8d, 0xcb, 0x9d, 0x54, 0x0f, 0x0b, 0xe6, 0xe7, 0xc7, 0x3e, 0xd9,
	0x17, 0x75, 0x63, 0xd3, 0x9c, 0x54, 0x9b, 0x93, 0x7d, 0x21, 0x5f, 0x89, 0x06, 0xe6, 0x54, 0xc8,
	0x82, 0xe5, 0x63, 0x25, 0x3f, 0x8b, 0xd7, 0x0e, 0x74, 0x86, 0x36, 0x58, 0xa3, 0x0b, 0x8c, 0x0d,
	0x31, 0x49, 0x63, 0x28, 0xe6, 0x54, 0x97, 0x0c, 0xbe, 0xf4, 0xf6, 0x77, 0x74, 0xd3, 0x13, 0xb3,
	0xd8, 0x14, 0xec, 0x21, 0xcc, 0x79, 0x4f, 0x3c, 0xc8, 0xd5, 0xb0, 0x36, 0xba, 0x98, 0x77, 0x2a,
	0x95, 0xc7, 0xe9, 0x8b, 0x3a, 0x3c, 0x82, 0x25, 0xd5, 0xf1, 0x2f, 0xc3, 0x85, 0xfc, 0x2a, 0x06,
	0x6b, 0x6d, 0x62, 0x88, 0x02, 0x07, 0x6b, 0x70, 0x60, 0x43, 0x08, 0x48, 0x67, 0xbb, 0xc0, 0xea,
	0x04, 0x32, 0xd5, 0x65, 0x56, 0x79, 0x64, 0x5e, 0x12, 0x4b, 0x9d, 0xed, 0x1e, 0x0a, 0x5f, 0xfe,
	0x10, 0x77, 0xb0, 0x27, 0x70, 0x56, 0xc7, 0x01, 0x4f, 0x17, 0xa4, 0xfa, 0x10, 0xa3, 0x8e, 0x82,
	0x8c, 0xd0, 0xe9, 0x0d, 0x04, 0xa9, 0xa6, 0xad, 0xea, 0x71, 0xcc, 0x6d, 0x89, 0x4e, 0x0b, 0x72,
	0xe6, 0xc1, 0x85, 0xe7, 0x66, 0x9a, 0xb6, 0xf2, 0x41, 0xbc, 0xaf, 0xd2, 0xfc, 0x52, 0x9e, 0x8f,
	0x7b, 0xc1, 0x71, 0xc3, 0x92, 0xfd, 0xc9, 0xe8, 0x33, 0x79, 0x37, 0xa2, 0x95, 0x3a, 0x83, 0xce,
	0xd0, 0x41, 0x5d, 0xf3, 0xfd, 0xab, 0xea, 0xfb, 0x58, 0x0c, 0x42, 0x4c, 0xc6, 0x16, 0x68, 0x1d,
	0x9b, 0xfd, 0xf9, 0xa7, 0xbd, 0xbf, 0x3e, 0xbb, 0xef, 0xb3, 0xd5, 0xef, 0x2f, 0x1b, 0x43, 0xdb,
	0x7c, 0x35, 0x0e, 0x31, 0x99, 0x1c, 0xdb, 0x26, 0x55, 0xdb, 0x24, 0x8c, 0x0d, 0x58, 0x9a, 0x6c,
	0x70, 0xe3, 0xd2, 0xf0, 0x44, 0xe7, 0x1f, 0x6a, 0xd5, 0xe0, 0xd4, 0x4f, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x1b, 0x95, 0xbb, 0x87, 0x8e, 0x03, 0x00, 0x00,
}
