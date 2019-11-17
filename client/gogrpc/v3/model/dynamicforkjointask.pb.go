// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/dynamicforkjointask.proto

package model // import "github.com/mactaggart/conductor/client/gogrpc/v3/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DynamicForkJoinTask struct {
	TaskName             string                    `protobuf:"bytes,1,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	WorkflowName         string                    `protobuf:"bytes,2,opt,name=workflow_name,json=workflowName,proto3" json:"workflow_name,omitempty"`
	ReferenceName        string                    `protobuf:"bytes,3,opt,name=reference_name,json=referenceName,proto3" json:"reference_name,omitempty"`
	Input                map[string]*_struct.Value `protobuf:"bytes,4,rep,name=input,proto3" json:"input,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Type                 string                    `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *DynamicForkJoinTask) Reset()         { *m = DynamicForkJoinTask{} }
func (m *DynamicForkJoinTask) String() string { return proto.CompactTextString(m) }
func (*DynamicForkJoinTask) ProtoMessage()    {}
func (*DynamicForkJoinTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_dynamicforkjointask_0ad7b324d33736ed, []int{0}
}
func (m *DynamicForkJoinTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicForkJoinTask.Unmarshal(m, b)
}
func (m *DynamicForkJoinTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicForkJoinTask.Marshal(b, m, deterministic)
}
func (dst *DynamicForkJoinTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicForkJoinTask.Merge(dst, src)
}
func (m *DynamicForkJoinTask) XXX_Size() int {
	return xxx_messageInfo_DynamicForkJoinTask.Size(m)
}
func (m *DynamicForkJoinTask) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicForkJoinTask.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicForkJoinTask proto.InternalMessageInfo

func (m *DynamicForkJoinTask) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *DynamicForkJoinTask) GetWorkflowName() string {
	if m != nil {
		return m.WorkflowName
	}
	return ""
}

func (m *DynamicForkJoinTask) GetReferenceName() string {
	if m != nil {
		return m.ReferenceName
	}
	return ""
}

func (m *DynamicForkJoinTask) GetInput() map[string]*_struct.Value {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *DynamicForkJoinTask) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*DynamicForkJoinTask)(nil), "conductor.proto.DynamicForkJoinTask")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.DynamicForkJoinTask.InputEntry")
}

func init() {
	proto.RegisterFile("model/dynamicforkjointask.proto", fileDescriptor_dynamicforkjointask_0ad7b324d33736ed)
}

var fileDescriptor_dynamicforkjointask_0ad7b324d33736ed = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xe9, 0x3f, 0xb1, 0xa9, 0x55, 0x59, 0x51, 0x4a, 0x2b, 0x58, 0x14, 0xa1, 0x07, 0x49,
	0xa0, 0x05, 0x11, 0x8f, 0xc5, 0x0a, 0x7a, 0x90, 0x52, 0xc4, 0x83, 0x17, 0xc9, 0xa6, 0xd9, 0x18,
	0x37, 0xc9, 0x2c, 0x69, 0xb6, 0x75, 0x3f, 0x93, 0x5f, 0x52, 0x36, 0xb1, 0x15, 0x4a, 0x6f, 0x93,
	0x37, 0xbf, 0x79, 0x79, 0x99, 0xa0, 0x0b, 0x0d, 0x73, 0xae, 0xc8, 0xbc, 0x30, 0x54, 0x4b, 0x96,
	0x80, 0x4d, 0xbf, 0x40, 0x1a, 0x47, 0x17, 0x29, 0xce, 0x2c, 0x38, 0x88, 0x8e, 0x18, 0x98, 0x79,
	0xce, 0x1c, 0xd8, 0x20, 0x74, 0xcf, 0x05, 0x80, 0x50, 0x9c, 0xf8, 0x53, 0x9c, 0x27, 0x64, 0xe1,
	0x6c, 0xce, 0x5c, 0xe8, 0x5e, 0xfe, 0x54, 0xd1, 0xc9, 0x43, 0x30, 0x7b, 0x04, 0x9b, 0x3e, 0x83,
	0x34, 0xaf, 0x74, 0x91, 0x46, 0x3d, 0xd4, 0x2c, 0x4d, 0x3f, 0x0c, 0xd5, 0xbc, 0x53, 0xe9, 0x57,
	0x06, 0xcd, 0xd9, 0x7e, 0x29, 0xbc, 0x50, 0xcd, 0xa3, 0x2b, 0xd4, 0x5e, 0x81, 0x4d, 0x13, 0x05,
	0xab, 0x00, 0x54, 0x3d, 0x70, 0xb0, 0x16, 0x3d, 0x74, 0x8d, 0x0e, 0x2d, 0x4f, 0xb8, 0xe5, 0x86,
	0xf1, 0x40, 0xd5, 0x3c, 0xd5, 0xde, 0xa8, 0x1e, 0x9b, 0xa0, 0x86, 0x34, 0x59, 0xee, 0x3a, 0xf5,
	0x7e, 0x6d, 0xd0, 0x1a, 0x12, 0xbc, 0x95, 0x1f, 0xef, 0x48, 0x87, 0x9f, 0xca, 0x89, 0x89, 0x71,
	0xb6, 0x98, 0x85, 0xe9, 0x28, 0x42, 0x75, 0x57, 0x64, 0xbc, 0xd3, 0xf0, 0x77, 0xf8, 0xba, 0x3b,
	0x45, 0xe8, 0x1f, 0x8c, 0x8e, 0x51, 0x2d, 0xe5, 0xc5, 0xdf, 0x5b, 0xca, 0x32, 0xba, 0x41, 0x8d,
	0x25, 0x55, 0x79, 0x88, 0xdf, 0x1a, 0x9e, 0xe1, 0xb0, 0x29, 0xbc, 0xde, 0x14, 0x7e, 0x2b, 0xbb,
	0xb3, 0x00, 0xdd, 0x57, 0xef, 0x2a, 0x63, 0x85, 0x7a, 0x0c, 0x34, 0x36, 0xdc, 0x25, 0x4a, 0x7e,
	0x6f, 0x47, 0x1d, 0x9f, 0xee, 0xc8, 0x3a, 0x8d, 0xdf, 0x6f, 0x85, 0x74, 0x9f, 0x79, 0x8c, 0x19,
	0x68, 0xa2, 0x29, 0x73, 0x54, 0x08, 0x6a, 0x1d, 0xd9, 0x4c, 0x13, 0xa6, 0x24, 0x37, 0x8e, 0x08,
	0x10, 0x36, 0x63, 0x64, 0x39, 0x22, 0xfe, 0x93, 0xe3, 0x3d, 0xef, 0x3a, 0xfa, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x19, 0xd1, 0xcd, 0x6c, 0xf4, 0x01, 0x00, 0x00,
}