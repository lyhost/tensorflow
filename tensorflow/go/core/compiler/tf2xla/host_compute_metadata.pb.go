// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/compiler/tf2xla/host_compute_metadata.proto

package tf2xla

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	framework "github.com/tensorflow/tensorflow/tensorflow/go/core/framework"
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

// TensorMetadata indicates the type and shape of a Tensor that is
// part of a host compute transfer.
type TensorMetadata struct {
	Type                 framework.DataType          `protobuf:"varint,1,opt,name=type,proto3,enum=tensorflow.DataType" json:"type,omitempty"`
	Shape                *framework.TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TensorMetadata) Reset()         { *m = TensorMetadata{} }
func (m *TensorMetadata) String() string { return proto.CompactTextString(m) }
func (*TensorMetadata) ProtoMessage()    {}
func (*TensorMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed367e7a3f0d62a7, []int{0}
}

func (m *TensorMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TensorMetadata.Unmarshal(m, b)
}
func (m *TensorMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TensorMetadata.Marshal(b, m, deterministic)
}
func (m *TensorMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TensorMetadata.Merge(m, src)
}
func (m *TensorMetadata) XXX_Size() int {
	return xxx_messageInfo_TensorMetadata.Size(m)
}
func (m *TensorMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TensorMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TensorMetadata proto.InternalMessageInfo

func (m *TensorMetadata) GetType() framework.DataType {
	if m != nil {
		return m.Type
	}
	return framework.DataType_DT_INVALID
}

func (m *TensorMetadata) GetShape() *framework.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

// HostTransferMetadata describes a transfer either from host to device
// or device to host. It has a key that is unique to the computation,
// and metadata about the list of tensors being transferred.
type HostTransferMetadata struct {
	// The key used to identify this transfer.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// For each Tensor being transferred, its type and shape.
	Metadata             []*TensorMetadata `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HostTransferMetadata) Reset()         { *m = HostTransferMetadata{} }
func (m *HostTransferMetadata) String() string { return proto.CompactTextString(m) }
func (*HostTransferMetadata) ProtoMessage()    {}
func (*HostTransferMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed367e7a3f0d62a7, []int{1}
}

func (m *HostTransferMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostTransferMetadata.Unmarshal(m, b)
}
func (m *HostTransferMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostTransferMetadata.Marshal(b, m, deterministic)
}
func (m *HostTransferMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostTransferMetadata.Merge(m, src)
}
func (m *HostTransferMetadata) XXX_Size() int {
	return xxx_messageInfo_HostTransferMetadata.Size(m)
}
func (m *HostTransferMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_HostTransferMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_HostTransferMetadata proto.InternalMessageInfo

func (m *HostTransferMetadata) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HostTransferMetadata) GetMetadata() []*TensorMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// HostComputeMetadata describes all the sends and recvs
// from all host compute transfer ops in a computation.
type HostComputeMetadata struct {
	// Metadata about each device_to_host transfer
	DeviceToHost []*HostTransferMetadata `protobuf:"bytes,1,rep,name=device_to_host,json=deviceToHost,proto3" json:"device_to_host,omitempty"`
	// Metadata about each host_to_device transfer
	HostToDevice         []*HostTransferMetadata `protobuf:"bytes,2,rep,name=host_to_device,json=hostToDevice,proto3" json:"host_to_device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *HostComputeMetadata) Reset()         { *m = HostComputeMetadata{} }
func (m *HostComputeMetadata) String() string { return proto.CompactTextString(m) }
func (*HostComputeMetadata) ProtoMessage()    {}
func (*HostComputeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed367e7a3f0d62a7, []int{2}
}

func (m *HostComputeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostComputeMetadata.Unmarshal(m, b)
}
func (m *HostComputeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostComputeMetadata.Marshal(b, m, deterministic)
}
func (m *HostComputeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostComputeMetadata.Merge(m, src)
}
func (m *HostComputeMetadata) XXX_Size() int {
	return xxx_messageInfo_HostComputeMetadata.Size(m)
}
func (m *HostComputeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_HostComputeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_HostComputeMetadata proto.InternalMessageInfo

func (m *HostComputeMetadata) GetDeviceToHost() []*HostTransferMetadata {
	if m != nil {
		return m.DeviceToHost
	}
	return nil
}

func (m *HostComputeMetadata) GetHostToDevice() []*HostTransferMetadata {
	if m != nil {
		return m.HostToDevice
	}
	return nil
}

func init() {
	proto.RegisterType((*TensorMetadata)(nil), "tensorflow.tf2xla.TensorMetadata")
	proto.RegisterType((*HostTransferMetadata)(nil), "tensorflow.tf2xla.HostTransferMetadata")
	proto.RegisterType((*HostComputeMetadata)(nil), "tensorflow.tf2xla.HostComputeMetadata")
}

func init() {
	proto.RegisterFile("tensorflow/compiler/tf2xla/host_compute_metadata.proto", fileDescriptor_ed367e7a3f0d62a7)
}

var fileDescriptor_ed367e7a3f0d62a7 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4b, 0x02, 0x41,
	0x14, 0xc7, 0x19, 0xad, 0xa8, 0x51, 0xa4, 0x36, 0x83, 0x45, 0x3a, 0x98, 0x10, 0xed, 0x21, 0x76,
	0x61, 0x83, 0x6e, 0x5d, 0xd4, 0x43, 0x17, 0x41, 0xb6, 0x3d, 0x44, 0x97, 0x61, 0xd4, 0xb7, 0xab,
	0xb8, 0xeb, 0x5b, 0x66, 0xc6, 0xcc, 0xbf, 0xab, 0x7f, 0xae, 0x63, 0xcc, 0x8c, 0xf9, 0x23, 0x25,
	0xba, 0x0d, 0xf3, 0x3e, 0xdf, 0xef, 0xf7, 0xf1, 0xde, 0xa3, 0x8f, 0x0a, 0x66, 0x12, 0x45, 0x92,
	0xe1, 0x22, 0x18, 0x62, 0x5e, 0x4c, 0x32, 0x10, 0x81, 0x4a, 0xc2, 0x8f, 0x8c, 0x07, 0x63, 0x94,
	0x8a, 0xe9, 0xcf, 0xb9, 0x02, 0x96, 0x83, 0xe2, 0x23, 0xae, 0xb8, 0x5f, 0x08, 0x54, 0xe8, 0x5c,
	0x6c, 0x74, 0xbe, 0xc5, 0x1b, 0xf7, 0x3b, 0x56, 0x02, 0x82, 0x44, 0xf0, 0x1c, 0x16, 0x28, 0xa6,
	0x81, 0xad, 0x30, 0x39, 0xe6, 0x05, 0x58, 0x83, 0xc6, 0xed, 0x1f, 0xf4, 0xb2, 0x00, 0x69, 0xb1,
	0xd6, 0x8c, 0xd6, 0x62, 0x03, 0xf6, 0x56, 0xf9, 0x8e, 0x47, 0x8f, 0x34, 0xe0, 0x92, 0x26, 0xf1,
	0x6a, 0x61, 0xdd, 0xdf, 0x6a, 0xa4, 0xcb, 0x15, 0x8f, 0x97, 0x05, 0x44, 0x86, 0x70, 0x42, 0x7a,
	0x6c, 0x12, 0xdd, 0x52, 0x93, 0x78, 0x95, 0xf0, 0x7a, 0x1b, 0xb5, 0xa6, 0x2f, 0xba, 0xdc, 0xd7,
	0x41, 0x91, 0x45, 0x5b, 0x29, 0xad, 0x3f, 0xa3, 0x54, 0xb1, 0xe0, 0x33, 0x99, 0xc0, 0x26, 0xf5,
	0x9c, 0x96, 0xa7, 0xb0, 0x34, 0xa1, 0x67, 0x91, 0x7e, 0x3a, 0x4f, 0xf4, 0xf4, 0x67, 0x26, 0x6e,
	0xa9, 0x59, 0xf6, 0x2a, 0xe1, 0x8d, 0xbf, 0x37, 0x14, 0x7f, 0xb7, 0xf9, 0x68, 0x2d, 0x69, 0x7d,
	0x12, 0x7a, 0xa9, 0x93, 0x3a, 0x76, 0xbe, 0xeb, 0xa0, 0x1e, 0xad, 0x8d, 0xe0, 0x7d, 0x32, 0x04,
	0xa6, 0x90, 0xe9, 0x0d, 0xb8, 0xc4, 0x98, 0xdf, 0x1d, 0x30, 0x3f, 0xd4, 0x69, 0x54, 0xb5, 0xf2,
	0x18, 0x75, 0x55, 0xdb, 0x99, 0x35, 0x2a, 0x64, 0xf6, 0x7f, 0xd5, 0xeb, 0xff, 0xed, 0xb4, 0x3c,
	0xc6, 0xae, 0x11, 0xb7, 0x91, 0x5e, 0xa1, 0x48, 0xf7, 0xb5, 0xed, 0x6a, 0x9c, 0x84, 0xaf, 0x19,
	0x37, 0xb3, 0x94, 0x7d, 0xf2, 0xd6, 0x49, 0x27, 0x6a, 0x3c, 0x1f, 0xf8, 0x43, 0xcc, 0x83, 0xad,
	0x4d, 0x1f, 0x7e, 0xa6, 0x68, 0x4f, 0xe0, 0xd7, 0x01, 0x7e, 0x11, 0x32, 0x38, 0x31, 0x67, 0xf0,
	0xf0, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xca, 0x65, 0x64, 0xa8, 0x02, 0x00, 0x00,
}
