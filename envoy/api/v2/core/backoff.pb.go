// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/core/backoff.proto

package envoy_api_v2_core

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type BackoffStrategy struct {
	BaseInterval         *duration.Duration `protobuf:"bytes,1,opt,name=base_interval,json=baseInterval,proto3" json:"base_interval,omitempty"`
	MaxInterval          *duration.Duration `protobuf:"bytes,2,opt,name=max_interval,json=maxInterval,proto3" json:"max_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BackoffStrategy) Reset()         { *m = BackoffStrategy{} }
func (m *BackoffStrategy) String() string { return proto.CompactTextString(m) }
func (*BackoffStrategy) ProtoMessage()    {}
func (*BackoffStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_743c92d7b5268904, []int{0}
}

func (m *BackoffStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackoffStrategy.Unmarshal(m, b)
}
func (m *BackoffStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackoffStrategy.Marshal(b, m, deterministic)
}
func (m *BackoffStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackoffStrategy.Merge(m, src)
}
func (m *BackoffStrategy) XXX_Size() int {
	return xxx_messageInfo_BackoffStrategy.Size(m)
}
func (m *BackoffStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_BackoffStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_BackoffStrategy proto.InternalMessageInfo

func (m *BackoffStrategy) GetBaseInterval() *duration.Duration {
	if m != nil {
		return m.BaseInterval
	}
	return nil
}

func (m *BackoffStrategy) GetMaxInterval() *duration.Duration {
	if m != nil {
		return m.MaxInterval
	}
	return nil
}

func init() {
	proto.RegisterType((*BackoffStrategy)(nil), "envoy.api.v2.core.BackoffStrategy")
}

func init() { proto.RegisterFile("envoy/api/v2/core/backoff.proto", fileDescriptor_743c92d7b5268904) }

var fileDescriptor_743c92d7b5268904 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0xbb, 0x4e, 0xf3, 0x30,
	0x14, 0xfe, 0x5d, 0xfd, 0x45, 0x55, 0x5a, 0x6e, 0x11, 0xe2, 0xd2, 0xa1, 0x45, 0x4c, 0x88, 0xc1,
	0x96, 0xd2, 0x99, 0xc5, 0x42, 0x48, 0x0c, 0xa0, 0xaa, 0x3c, 0x00, 0x3a, 0x49, 0x9c, 0xc8, 0x22,
	0xf1, 0x89, 0x5c, 0xc7, 0x4a, 0x76, 0x76, 0x56, 0x5e, 0x81, 0x3e, 0x05, 0xcf, 0xc4, 0xd8, 0x01,
	0xa1, 0x38, 0x89, 0x18, 0x18, 0xd8, 0x2c, 0x7f, 0xd7, 0xf3, 0x79, 0x73, 0xa1, 0x2c, 0xd6, 0x0c,
	0x0a, 0xc9, 0x6c, 0xc0, 0x22, 0xd4, 0x82, 0x85, 0x10, 0x3d, 0x63, 0x92, 0xd0, 0x42, 0xa3, 0x41,
	0xff, 0xd0, 0x11, 0x28, 0x14, 0x92, 0xda, 0x80, 0x36, 0x84, 0xe9, 0x2c, 0x45, 0x4c, 0x33, 0xc1,
	0x1c, 0x21, 0x2c, 0x13, 0x16, 0x97, 0x1a, 0x8c, 0x44, 0xd5, 0x4a, 0xa6, 0xb3, 0x32, 0x2e, 0x80,
	0x81, 0x52, 0x68, 0xdc, 0xf7, 0x9a, 0xe5, 0x32, 0xd5, 0x60, 0x44, 0x87, 0x9f, 0x58, 0xc8, 0x64,
	0x0c, 0x46, 0xb0, 0xfe, 0xd1, 0x02, 0x17, 0xef, 0xc4, 0xdb, 0xe7, 0x6d, 0xfa, 0xa3, 0x69, 0x04,
	0x69, 0xed, 0x3f, 0x78, 0xbb, 0x21, 0xac, 0xc5, 0x93, 0x54, 0x46, 0x68, 0x0b, 0xd9, 0x29, 0x39,
	0x27, 0x97, 0xe3, 0xe0, 0x8c, 0xb6, 0x25, 0x68, 0x5f, 0x82, 0xde, 0x74, 0x25, 0xf8, 0xde, 0x96,
	0x8f, 0x37, 0x64, 0x34, 0x22, 0xc1, 0xff, 0x83, 0x8f, 0x97, 0xeb, 0xd5, 0xa4, 0xd1, 0xdf, 0x75,
	0x72, 0xff, 0xd6, 0x9b, 0xe4, 0x50, 0xfd, 0xd8, 0x0d, 0xfe, 0xb2, 0x1b, 0x6d, 0xf9, 0x70, 0x43,
	0x06, 0x57, 0xff, 0x56, 0xe3, 0x1c, 0xaa, 0xde, 0x87, 0xdf, 0x7f, 0xbe, 0x7d, 0xbd, 0x0e, 0x8f,
	0xfd, 0xa3, 0x76, 0x9f, 0x08, 0x55, 0x22, 0x53, 0xb7, 0x0f, 0xb5, 0x0b, 0x6f, 0x2e, 0x91, 0x3a,
	0xa0, 0xd0, 0x58, 0xd5, 0xf4, 0xd7, 0x86, 0x7c, 0xd2, 0xdd, 0xb9, 0x6c, 0xf2, 0x96, 0x24, 0xdc,
	0x71, 0xc1, 0x8b, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xd5, 0x45, 0x5f, 0x90, 0x01, 0x00,
	0x00,
}
