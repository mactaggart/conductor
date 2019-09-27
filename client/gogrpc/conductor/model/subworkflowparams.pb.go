// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/subworkflowparams.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SubWorkflowParams struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              int32             `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	TaskToDomain         map[string]string `protobuf:"bytes,3,rep,name=task_to_domain,json=taskToDomain,proto3" json:"task_to_domain,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SubWorkflowParams) Reset()         { *m = SubWorkflowParams{} }
func (m *SubWorkflowParams) String() string { return proto.CompactTextString(m) }
func (*SubWorkflowParams) ProtoMessage()    {}
func (*SubWorkflowParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_60987417fee48ff4, []int{0}
}

func (m *SubWorkflowParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubWorkflowParams.Unmarshal(m, b)
}
func (m *SubWorkflowParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubWorkflowParams.Marshal(b, m, deterministic)
}
func (m *SubWorkflowParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubWorkflowParams.Merge(m, src)
}
func (m *SubWorkflowParams) XXX_Size() int {
	return xxx_messageInfo_SubWorkflowParams.Size(m)
}
func (m *SubWorkflowParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SubWorkflowParams.DiscardUnknown(m)
}

var xxx_messageInfo_SubWorkflowParams proto.InternalMessageInfo

func (m *SubWorkflowParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SubWorkflowParams) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *SubWorkflowParams) GetTaskToDomain() map[string]string {
	if m != nil {
		return m.TaskToDomain
	}
	return nil
}

func init() {
	proto.RegisterType((*SubWorkflowParams)(nil), "conductor.proto.SubWorkflowParams")
	proto.RegisterMapType((map[string]string)(nil), "conductor.proto.SubWorkflowParams.TaskToDomainEntry")
}

func init() { proto.RegisterFile("model/subworkflowparams.proto", fileDescriptor_60987417fee48ff4) }

var fileDescriptor_60987417fee48ff4 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x14, 0x84, 0x49, 0x63, 0x95, 0xae, 0xa2, 0x76, 0xf5, 0x10, 0x14, 0x21, 0x78, 0xca, 0x69, 0x03,
	0xea, 0x41, 0x04, 0x11, 0x8a, 0xde, 0x4b, 0x2c, 0x08, 0xbd, 0x94, 0x97, 0xcd, 0x36, 0x86, 0x64,
	0xf7, 0xc5, 0xcd, 0x4b, 0x6b, 0xff, 0xac, 0xbf, 0x45, 0xdc, 0xd6, 0x52, 0x9a, 0xdb, 0x9b, 0x61,
	0x66, 0xe7, 0x63, 0xd9, 0x8d, 0xc6, 0x4c, 0x55, 0x71, 0xd3, 0xa6, 0x4b, 0xb4, 0xe5, 0xbc, 0xc2,
	0x65, 0x0d, 0x16, 0x74, 0x23, 0x6a, 0x8b, 0x84, 0xfc, 0x4c, 0xa2, 0xc9, 0x5a, 0x49, 0x68, 0xd7,
	0xc6, 0xed, 0x8f, 0xc7, 0x86, 0xef, 0x6d, 0xfa, 0xb1, 0x09, 0x8f, 0x5d, 0x98, 0x73, 0x76, 0x60,
	0x40, 0xab, 0xc0, 0x0b, 0xbd, 0x68, 0x90, 0xb8, 0x9b, 0x07, 0xec, 0x68, 0xa1, 0x6c, 0x53, 0xa0,
	0x09, 0x7a, 0xa1, 0x17, 0xf5, 0x93, 0x7f, 0xc9, 0xa7, 0xec, 0x94, 0xa0, 0x29, 0x67, 0x84, 0xb3,
	0x0c, 0x35, 0x14, 0x26, 0xf0, 0x43, 0x3f, 0x3a, 0xbe, 0x7b, 0x10, 0x7b, 0x6b, 0xa2, 0xb3, 0x24,
	0x26, 0xd0, 0x94, 0x13, 0x7c, 0x75, 0xb5, 0x37, 0x43, 0x76, 0x95, 0x9c, 0xd0, 0x8e, 0x75, 0xf5,
	0xc2, 0x86, 0x9d, 0x08, 0x3f, 0x67, 0x7e, 0xa9, 0x56, 0x1b, 0xba, 0xbf, 0x93, 0x5f, 0xb2, 0xfe,
	0x02, 0xaa, 0x56, 0x39, 0xb4, 0x41, 0xb2, 0x16, 0x4f, 0xbd, 0x47, 0x6f, 0xf4, 0xc5, 0xae, 0x25,
	0x6a, 0x61, 0x14, 0xcd, 0xab, 0xe2, 0x7b, 0x9f, 0x68, 0x74, 0xd1, 0x41, 0x1a, 0xa7, 0xd3, 0xe7,
	0xbc, 0xa0, 0xcf, 0x36, 0x15, 0x12, 0x75, 0xac, 0x41, 0x12, 0xe4, 0x39, 0x58, 0x8a, 0xb7, 0xdd,
	0x58, 0x56, 0x85, 0x32, 0x14, 0xe7, 0x98, 0xdb, 0x5a, 0xee, 0xf8, 0xee, 0xf3, 0xd3, 0x43, 0xf7,
	0xf4, 0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x23, 0x6a, 0x01, 0x8c, 0x01, 0x00, 0x00,
}
