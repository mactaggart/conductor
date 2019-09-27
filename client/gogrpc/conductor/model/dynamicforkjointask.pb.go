// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/dynamicforkjointask.proto

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
	return fileDescriptor_67f2bb1ac0addb2c, []int{0}
}

func (m *DynamicForkJoinTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicForkJoinTask.Unmarshal(m, b)
}
func (m *DynamicForkJoinTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicForkJoinTask.Marshal(b, m, deterministic)
}
func (m *DynamicForkJoinTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicForkJoinTask.Merge(m, src)
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

func init() { proto.RegisterFile("model/dynamicforkjointask.proto", fileDescriptor_67f2bb1ac0addb2c) }

var fileDescriptor_67f2bb1ac0addb2c = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0xc6, 0xe9, 0xbf, 0x97, 0xb7, 0xa9, 0x55, 0x89, 0x28, 0xa5, 0x15, 0x2c, 0x8a, 0xd0, 0x83,
	0x24, 0x50, 0x2f, 0x22, 0x78, 0x29, 0x56, 0xd0, 0x83, 0x94, 0x22, 0x1e, 0xbc, 0x48, 0x36, 0xcd,
	0xc6, 0xb8, 0x9b, 0xcc, 0x92, 0xcd, 0x5a, 0xf7, 0x33, 0xf9, 0x25, 0x65, 0x13, 0x5b, 0xa5, 0xf4,
	0x36, 0x79, 0xe6, 0x37, 0x4f, 0x9e, 0x4c, 0xd0, 0x89, 0x86, 0x85, 0x48, 0xe9, 0xa2, 0x34, 0x4c,
	0x2b, 0x1e, 0x83, 0x4d, 0xde, 0x41, 0x19, 0xc7, 0xf2, 0x84, 0x64, 0x16, 0x1c, 0xe0, 0x3d, 0x0e,
	0x66, 0x51, 0x70, 0x07, 0x36, 0x08, 0xfd, 0x63, 0x09, 0x20, 0x53, 0x41, 0xfd, 0x29, 0x2a, 0x62,
	0x9a, 0x3b, 0x5b, 0x70, 0x17, 0xba, 0xa7, 0x5f, 0x75, 0x74, 0x70, 0x1b, 0xcc, 0xee, 0xc0, 0x26,
	0x0f, 0xa0, 0xcc, 0x13, 0xcb, 0x13, 0x3c, 0x40, 0xed, 0xca, 0xf4, 0xd5, 0x30, 0x2d, 0x7a, 0xb5,
	0x61, 0x6d, 0xd4, 0x9e, 0xff, 0xaf, 0x84, 0x47, 0xa6, 0x05, 0x3e, 0x43, 0xdd, 0x25, 0xd8, 0x24,
	0x4e, 0x61, 0x19, 0x80, 0xba, 0x07, 0x76, 0x56, 0xa2, 0x87, 0xce, 0xd1, 0xae, 0x15, 0xb1, 0xb0,
	0xc2, 0x70, 0x11, 0xa8, 0x86, 0xa7, 0xba, 0x6b, 0xd5, 0x63, 0x53, 0xd4, 0x52, 0x26, 0x2b, 0x5c,
	0xaf, 0x39, 0x6c, 0x8c, 0x3a, 0x63, 0x4a, 0x36, 0xf2, 0x93, 0x2d, 0xe9, 0xc8, 0x7d, 0x35, 0x31,
	0x35, 0xce, 0x96, 0xf3, 0x30, 0x8d, 0x31, 0x6a, 0xba, 0x32, 0x13, 0xbd, 0x96, 0xbf, 0xc3, 0xd7,
	0xfd, 0x19, 0x42, 0xbf, 0x20, 0xde, 0x47, 0x8d, 0x44, 0x94, 0x3f, 0x6f, 0xa9, 0x4a, 0x7c, 0x81,
	0x5a, 0x1f, 0x2c, 0x2d, 0x42, 0xfc, 0xce, 0xf8, 0x88, 0x84, 0x4d, 0x91, 0xd5, 0xa6, 0xc8, 0x73,
	0xd5, 0x9d, 0x07, 0xe8, 0xba, 0x7e, 0x55, 0x9b, 0xe4, 0x68, 0xc0, 0x41, 0x13, 0x23, 0x5c, 0x9c,
	0xaa, 0xcf, 0xcd, 0xa8, 0x93, 0xc3, 0x2d, 0x59, 0x67, 0xd1, 0xcb, 0x8d, 0x54, 0xee, 0xad, 0x88,
	0x08, 0x07, 0x4d, 0x35, 0xe3, 0x8e, 0x49, 0xc9, 0xac, 0xa3, 0xeb, 0x69, 0xca, 0x53, 0x25, 0x8c,
	0xa3, 0x12, 0xa4, 0xcd, 0xf8, 0x1f, 0xdd, 0xff, 0x75, 0xf4, 0xcf, 0x9b, 0x5f, 0x7e, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x97, 0x73, 0x85, 0x05, 0xfb, 0x01, 0x00, 0x00,
}
