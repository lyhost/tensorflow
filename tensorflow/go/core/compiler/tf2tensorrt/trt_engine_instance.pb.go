// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/compiler/tf2tensorrt/utils/trt_engine_instance.proto

package tf2tensorrt

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

// Containing information for a serialized TensorRT engine.
type TRTEngineInstance struct {
	// The input shapes of the TRT engine.
	InputShapes []*framework.TensorShapeProto `protobuf:"bytes,1,rep,name=input_shapes,json=inputShapes,proto3" json:"input_shapes,omitempty"`
	// The serialized TRT engine.
	//
	// TODO(laigd): consider using a more efficient in-memory representation
	// instead of string which is the default here.
	SerializedEngine     []byte   `protobuf:"bytes,2,opt,name=serialized_engine,json=serializedEngine,proto3" json:"serialized_engine,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TRTEngineInstance) Reset()         { *m = TRTEngineInstance{} }
func (m *TRTEngineInstance) String() string { return proto.CompactTextString(m) }
func (*TRTEngineInstance) ProtoMessage()    {}
func (*TRTEngineInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_0598e67cb2d6bd54, []int{0}
}

func (m *TRTEngineInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TRTEngineInstance.Unmarshal(m, b)
}
func (m *TRTEngineInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TRTEngineInstance.Marshal(b, m, deterministic)
}
func (m *TRTEngineInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TRTEngineInstance.Merge(m, src)
}
func (m *TRTEngineInstance) XXX_Size() int {
	return xxx_messageInfo_TRTEngineInstance.Size(m)
}
func (m *TRTEngineInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_TRTEngineInstance.DiscardUnknown(m)
}

var xxx_messageInfo_TRTEngineInstance proto.InternalMessageInfo

func (m *TRTEngineInstance) GetInputShapes() []*framework.TensorShapeProto {
	if m != nil {
		return m.InputShapes
	}
	return nil
}

func (m *TRTEngineInstance) GetSerializedEngine() []byte {
	if m != nil {
		return m.SerializedEngine
	}
	return nil
}

func init() {
	proto.RegisterType((*TRTEngineInstance)(nil), "tensorflow.tensorrt.TRTEngineInstance")
}

func init() {
	proto.RegisterFile("tensorflow/compiler/tf2tensorrt/utils/trt_engine_instance.proto", fileDescriptor_0598e67cb2d6bd54)
}

var fileDescriptor_0598e67cb2d6bd54 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x2f, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0xd7, 0x4f, 0xce, 0xcf, 0x2d, 0xc8, 0xcc, 0x49, 0x2d, 0xd2,
	0x2f, 0x49, 0x33, 0x82, 0x08, 0x17, 0x95, 0xe8, 0x97, 0x96, 0x64, 0xe6, 0x14, 0xeb, 0x97, 0x14,
	0x95, 0xc4, 0xa7, 0xe6, 0xa5, 0x67, 0xe6, 0xa5, 0xc6, 0x67, 0xe6, 0x15, 0x97, 0x24, 0xe6, 0x25,
	0xa7, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0x23, 0x0c, 0xd0, 0x83, 0x69, 0x92, 0xd2,
	0x41, 0x31, 0xb5, 0x28, 0x55, 0x3f, 0xad, 0x28, 0x31, 0x37, 0xb5, 0x3c, 0xbf, 0x28, 0x5b, 0x1f,
	0x22, 0x13, 0x5f, 0x9c, 0x91, 0x58, 0x00, 0x35, 0x42, 0xa9, 0x91, 0x91, 0x4b, 0x30, 0x24, 0x28,
	0xc4, 0x15, 0x6c, 0xbe, 0x27, 0xd4, 0x78, 0x21, 0x7b, 0x2e, 0x9e, 0xcc, 0xbc, 0x82, 0xd2, 0x12,
	0x88, 0xd2, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x19, 0x3d, 0x24, 0xfb, 0x42, 0xc0,
	0xcc, 0x60, 0x90, 0x7c, 0x00, 0xc8, 0xa4, 0x20, 0x6e, 0xb0, 0x0e, 0xb0, 0x40, 0xb1, 0x90, 0x36,
	0x97, 0x60, 0x71, 0x6a, 0x51, 0x66, 0x62, 0x4e, 0x66, 0x55, 0x6a, 0x0a, 0xd4, 0xf5, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x02, 0x08, 0x09, 0x88, 0xad, 0x4e, 0x5e, 0x51, 0x1e, 0xe9, 0x99,
	0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x48, 0xce, 0xc7, 0xce, 0x4c, 0xcf, 0x87,
	0xf8, 0x0b, 0x5b, 0x90, 0x25, 0xb1, 0x81, 0xbd, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb5,
	0x8f, 0x49, 0xbb, 0x5c, 0x01, 0x00, 0x00,
}
