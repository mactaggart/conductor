// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/startworkflowrequest.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StartWorkflowRequest struct {
	Name                            string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version                         int32                     `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	CorrelationId                   string                    `protobuf:"bytes,3,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	Input                           map[string]*_struct.Value `protobuf:"bytes,4,rep,name=input,proto3" json:"input,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TaskToDomain                    map[string]string         `protobuf:"bytes,5,rep,name=task_to_domain,json=taskToDomain,proto3" json:"task_to_domain,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	WorkflowDef                     *WorkflowDef              `protobuf:"bytes,6,opt,name=workflow_def,json=workflowDef,proto3" json:"workflow_def,omitempty"`
	ExternalInputPayloadStoragePath string                    `protobuf:"bytes,7,opt,name=external_input_payload_storage_path,json=externalInputPayloadStoragePath,proto3" json:"external_input_payload_storage_path,omitempty"`
	Priority                        int32                     `protobuf:"varint,8,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                  `json:"-"`
	XXX_unrecognized                []byte                    `json:"-"`
	XXX_sizecache                   int32                     `json:"-"`
}

func (m *StartWorkflowRequest) Reset()         { *m = StartWorkflowRequest{} }
func (m *StartWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*StartWorkflowRequest) ProtoMessage()    {}
func (*StartWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78738fc9349cf1f3, []int{0}
}

func (m *StartWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartWorkflowRequest.Unmarshal(m, b)
}
func (m *StartWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartWorkflowRequest.Marshal(b, m, deterministic)
}
func (m *StartWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartWorkflowRequest.Merge(m, src)
}
func (m *StartWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_StartWorkflowRequest.Size(m)
}
func (m *StartWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartWorkflowRequest proto.InternalMessageInfo

func (m *StartWorkflowRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StartWorkflowRequest) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *StartWorkflowRequest) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *StartWorkflowRequest) GetInput() map[string]*_struct.Value {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *StartWorkflowRequest) GetTaskToDomain() map[string]string {
	if m != nil {
		return m.TaskToDomain
	}
	return nil
}

func (m *StartWorkflowRequest) GetWorkflowDef() *WorkflowDef {
	if m != nil {
		return m.WorkflowDef
	}
	return nil
}

func (m *StartWorkflowRequest) GetExternalInputPayloadStoragePath() string {
	if m != nil {
		return m.ExternalInputPayloadStoragePath
	}
	return ""
}

func (m *StartWorkflowRequest) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterType((*StartWorkflowRequest)(nil), "conductor.proto.StartWorkflowRequest")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.StartWorkflowRequest.InputEntry")
	proto.RegisterMapType((map[string]string)(nil), "conductor.proto.StartWorkflowRequest.TaskToDomainEntry")
}

func init() { proto.RegisterFile("model/startworkflowrequest.proto", fileDescriptor_78738fc9349cf1f3) }

var fileDescriptor_78738fc9349cf1f3 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x5f, 0x6b, 0xdb, 0x30,
	0x14, 0xc5, 0x71, 0x93, 0xf4, 0x8f, 0xd2, 0x75, 0x9b, 0x28, 0x9d, 0xc9, 0x0a, 0x33, 0x1b, 0x83,
	0x3c, 0x0c, 0x79, 0x64, 0x0f, 0x1b, 0x83, 0x51, 0x28, 0xdd, 0xa0, 0xb0, 0x87, 0xe0, 0x96, 0x0d,
	0x06, 0xc3, 0x28, 0xb2, 0xec, 0x88, 0xc8, 0xba, 0x9e, 0x7c, 0xdd, 0x34, 0xdf, 0x72, 0x1f, 0x69,
	0x44, 0x8a, 0x5b, 0xd3, 0xe6, 0x61, 0x6f, 0xba, 0x57, 0xe7, 0xfc, 0x24, 0x1d, 0x5d, 0x12, 0x95,
	0x90, 0x49, 0x1d, 0xd7, 0xc8, 0x2d, 0x2e, 0xc1, 0x2e, 0x72, 0x0d, 0x4b, 0x2b, 0xff, 0x34, 0xb2,
	0x46, 0x56, 0x59, 0x40, 0xa0, 0x4f, 0x05, 0x98, 0xac, 0x11, 0x08, 0xd6, 0x37, 0x46, 0x2f, 0xbc,
	0xa5, 0x55, 0x67, 0x32, 0xdf, 0x6c, 0x9c, 0x16, 0x00, 0x85, 0x96, 0xb1, 0xab, 0x66, 0x4d, 0x1e,
	0xd7, 0x68, 0x1b, 0xb1, 0xe1, 0xbc, 0xfe, 0xdb, 0x27, 0xc7, 0x57, 0xeb, 0x63, 0x7e, 0x6e, 0x8c,
	0x89, 0x3f, 0x86, 0x52, 0xd2, 0x37, 0xbc, 0x94, 0x61, 0x10, 0x05, 0xe3, 0x83, 0xc4, 0xad, 0x69,
	0x48, 0xf6, 0x6e, 0xa4, 0xad, 0x15, 0x98, 0x70, 0x27, 0x0a, 0xc6, 0x83, 0xa4, 0x2d, 0xe9, 0x5b,
	0x72, 0x24, 0xc0, 0x5a, 0xa9, 0x39, 0x2a, 0x30, 0xa9, 0xca, 0xc2, 0x9e, 0xf3, 0x3d, 0xe9, 0x74,
	0x2f, 0x33, 0xfa, 0x8d, 0x0c, 0x94, 0xa9, 0x1a, 0x0c, 0xfb, 0x51, 0x6f, 0x3c, 0x9c, 0xbc, 0x67,
	0x0f, 0x5e, 0xc1, 0xb6, 0x5d, 0x85, 0x5d, 0xae, 0x2d, 0x5f, 0x0d, 0xda, 0x55, 0xe2, 0xed, 0xf4,
	0x37, 0x39, 0x42, 0x5e, 0x2f, 0x52, 0x84, 0x34, 0x83, 0x92, 0x2b, 0x13, 0x0e, 0x1c, 0xf0, 0xe3,
	0xff, 0x01, 0xaf, 0x79, 0xbd, 0xb8, 0x86, 0x0b, 0xe7, 0xf4, 0xdc, 0x43, 0xec, 0xb4, 0xe8, 0x19,
	0x39, 0x6c, 0x73, 0x4c, 0x33, 0x99, 0x87, 0xbb, 0x51, 0x30, 0x1e, 0x4e, 0x4e, 0x1f, 0xc1, 0x5b,
	0xee, 0x85, 0xcc, 0x93, 0xe1, 0xf2, 0xbe, 0xa0, 0xdf, 0xc9, 0x1b, 0x79, 0x8b, 0xd2, 0x1a, 0xae,
	0x53, 0x77, 0xe3, 0xb4, 0xe2, 0x2b, 0x0d, 0x3c, 0x4b, 0x6b, 0x04, 0xcb, 0x0b, 0x99, 0x56, 0x1c,
	0xe7, 0xe1, 0x9e, 0xcb, 0xe8, 0x55, 0x2b, 0x75, 0xef, 0x9c, 0x7a, 0xe1, 0x95, 0xd7, 0x4d, 0x39,
	0xce, 0xe9, 0x88, 0xec, 0x57, 0x56, 0x81, 0x55, 0xb8, 0x0a, 0xf7, 0x5d, 0xee, 0x77, 0xf5, 0x68,
	0x4a, 0xc8, 0x7d, 0x3c, 0xf4, 0x19, 0xe9, 0x2d, 0xe4, 0x6a, 0xf3, 0x67, 0xeb, 0x25, 0x7d, 0x47,
	0x06, 0x37, 0x5c, 0x37, 0xd2, 0x7d, 0xd8, 0x70, 0x72, 0xc2, 0xfc, 0x34, 0xb0, 0x76, 0x1a, 0xd8,
	0x8f, 0xf5, 0x6e, 0xe2, 0x45, 0x9f, 0x77, 0x3e, 0x05, 0xa3, 0x33, 0xf2, 0xfc, 0x51, 0x3e, 0x5b,
	0xc0, 0xc7, 0x5d, 0xf0, 0x41, 0x07, 0x70, 0x8e, 0xe4, 0xa5, 0x80, 0x92, 0x19, 0x89, 0xb9, 0x56,
	0xb7, 0x0f, 0x43, 0x3b, 0x3f, 0xd9, 0xf6, 0x25, 0xd3, 0xd9, 0xaf, 0x2f, 0x85, 0xc2, 0x79, 0x33,
	0x63, 0x02, 0xca, 0xb8, 0xe4, 0x02, 0x79, 0x51, 0x70, 0x8b, 0xf1, 0x9d, 0x3d, 0x16, 0x5a, 0x49,
	0x83, 0x71, 0x01, 0x85, 0xad, 0x44, 0xa7, 0xef, 0x06, 0x7f, 0xb6, 0xeb, 0xe8, 0x1f, 0xfe, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x92, 0x33, 0x4e, 0xdd, 0x3b, 0x03, 0x00, 0x00,
}
