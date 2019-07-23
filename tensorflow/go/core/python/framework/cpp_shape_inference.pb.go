// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/python/framework/cpp_shape_inference.proto

package framework

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

type CppShapeInferenceResult struct {
	Shape                *framework.TensorShapeProto         `protobuf:"bytes,1,opt,name=shape,proto3" json:"shape,omitempty"`
	HandleData           *CppShapeInferenceResult_HandleData `protobuf:"bytes,4,opt,name=handle_data,json=handleData,proto3" json:"handle_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *CppShapeInferenceResult) Reset()         { *m = CppShapeInferenceResult{} }
func (m *CppShapeInferenceResult) String() string { return proto.CompactTextString(m) }
func (*CppShapeInferenceResult) ProtoMessage()    {}
func (*CppShapeInferenceResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2bf6bfecd64d111, []int{0}
}

func (m *CppShapeInferenceResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CppShapeInferenceResult.Unmarshal(m, b)
}
func (m *CppShapeInferenceResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CppShapeInferenceResult.Marshal(b, m, deterministic)
}
func (m *CppShapeInferenceResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CppShapeInferenceResult.Merge(m, src)
}
func (m *CppShapeInferenceResult) XXX_Size() int {
	return xxx_messageInfo_CppShapeInferenceResult.Size(m)
}
func (m *CppShapeInferenceResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CppShapeInferenceResult.DiscardUnknown(m)
}

var xxx_messageInfo_CppShapeInferenceResult proto.InternalMessageInfo

func (m *CppShapeInferenceResult) GetShape() *framework.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *CppShapeInferenceResult) GetHandleData() *CppShapeInferenceResult_HandleData {
	if m != nil {
		return m.HandleData
	}
	return nil
}

type CppShapeInferenceResult_HandleShapeAndType struct {
	Shape                *framework.TensorShapeProto `protobuf:"bytes,1,opt,name=shape,proto3" json:"shape,omitempty"`
	Dtype                framework.DataType          `protobuf:"varint,2,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *CppShapeInferenceResult_HandleShapeAndType) Reset() {
	*m = CppShapeInferenceResult_HandleShapeAndType{}
}
func (m *CppShapeInferenceResult_HandleShapeAndType) String() string {
	return proto.CompactTextString(m)
}
func (*CppShapeInferenceResult_HandleShapeAndType) ProtoMessage() {}
func (*CppShapeInferenceResult_HandleShapeAndType) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2bf6bfecd64d111, []int{0, 0}
}

func (m *CppShapeInferenceResult_HandleShapeAndType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CppShapeInferenceResult_HandleShapeAndType.Unmarshal(m, b)
}
func (m *CppShapeInferenceResult_HandleShapeAndType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CppShapeInferenceResult_HandleShapeAndType.Marshal(b, m, deterministic)
}
func (m *CppShapeInferenceResult_HandleShapeAndType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CppShapeInferenceResult_HandleShapeAndType.Merge(m, src)
}
func (m *CppShapeInferenceResult_HandleShapeAndType) XXX_Size() int {
	return xxx_messageInfo_CppShapeInferenceResult_HandleShapeAndType.Size(m)
}
func (m *CppShapeInferenceResult_HandleShapeAndType) XXX_DiscardUnknown() {
	xxx_messageInfo_CppShapeInferenceResult_HandleShapeAndType.DiscardUnknown(m)
}

var xxx_messageInfo_CppShapeInferenceResult_HandleShapeAndType proto.InternalMessageInfo

func (m *CppShapeInferenceResult_HandleShapeAndType) GetShape() *framework.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *CppShapeInferenceResult_HandleShapeAndType) GetDtype() framework.DataType {
	if m != nil {
		return m.Dtype
	}
	return framework.DataType_DT_INVALID
}

type CppShapeInferenceResult_HandleData struct {
	IsSet bool `protobuf:"varint,1,opt,name=is_set,json=isSet,proto3" json:"is_set,omitempty"`
	// Only valid if <is_set>.
	ShapeAndType         []*CppShapeInferenceResult_HandleShapeAndType `protobuf:"bytes,2,rep,name=shape_and_type,json=shapeAndType,proto3" json:"shape_and_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *CppShapeInferenceResult_HandleData) Reset()         { *m = CppShapeInferenceResult_HandleData{} }
func (m *CppShapeInferenceResult_HandleData) String() string { return proto.CompactTextString(m) }
func (*CppShapeInferenceResult_HandleData) ProtoMessage()    {}
func (*CppShapeInferenceResult_HandleData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2bf6bfecd64d111, []int{0, 1}
}

func (m *CppShapeInferenceResult_HandleData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CppShapeInferenceResult_HandleData.Unmarshal(m, b)
}
func (m *CppShapeInferenceResult_HandleData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CppShapeInferenceResult_HandleData.Marshal(b, m, deterministic)
}
func (m *CppShapeInferenceResult_HandleData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CppShapeInferenceResult_HandleData.Merge(m, src)
}
func (m *CppShapeInferenceResult_HandleData) XXX_Size() int {
	return xxx_messageInfo_CppShapeInferenceResult_HandleData.Size(m)
}
func (m *CppShapeInferenceResult_HandleData) XXX_DiscardUnknown() {
	xxx_messageInfo_CppShapeInferenceResult_HandleData.DiscardUnknown(m)
}

var xxx_messageInfo_CppShapeInferenceResult_HandleData proto.InternalMessageInfo

func (m *CppShapeInferenceResult_HandleData) GetIsSet() bool {
	if m != nil {
		return m.IsSet
	}
	return false
}

func (m *CppShapeInferenceResult_HandleData) GetShapeAndType() []*CppShapeInferenceResult_HandleShapeAndType {
	if m != nil {
		return m.ShapeAndType
	}
	return nil
}

type CppShapeInferenceInputsNeeded struct {
	InputTensorsNeeded         []int32  `protobuf:"varint,1,rep,packed,name=input_tensors_needed,json=inputTensorsNeeded,proto3" json:"input_tensors_needed,omitempty"`
	InputTensorsAsShapesNeeded []int32  `protobuf:"varint,2,rep,packed,name=input_tensors_as_shapes_needed,json=inputTensorsAsShapesNeeded,proto3" json:"input_tensors_as_shapes_needed,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *CppShapeInferenceInputsNeeded) Reset()         { *m = CppShapeInferenceInputsNeeded{} }
func (m *CppShapeInferenceInputsNeeded) String() string { return proto.CompactTextString(m) }
func (*CppShapeInferenceInputsNeeded) ProtoMessage()    {}
func (*CppShapeInferenceInputsNeeded) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2bf6bfecd64d111, []int{1}
}

func (m *CppShapeInferenceInputsNeeded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CppShapeInferenceInputsNeeded.Unmarshal(m, b)
}
func (m *CppShapeInferenceInputsNeeded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CppShapeInferenceInputsNeeded.Marshal(b, m, deterministic)
}
func (m *CppShapeInferenceInputsNeeded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CppShapeInferenceInputsNeeded.Merge(m, src)
}
func (m *CppShapeInferenceInputsNeeded) XXX_Size() int {
	return xxx_messageInfo_CppShapeInferenceInputsNeeded.Size(m)
}
func (m *CppShapeInferenceInputsNeeded) XXX_DiscardUnknown() {
	xxx_messageInfo_CppShapeInferenceInputsNeeded.DiscardUnknown(m)
}

var xxx_messageInfo_CppShapeInferenceInputsNeeded proto.InternalMessageInfo

func (m *CppShapeInferenceInputsNeeded) GetInputTensorsNeeded() []int32 {
	if m != nil {
		return m.InputTensorsNeeded
	}
	return nil
}

func (m *CppShapeInferenceInputsNeeded) GetInputTensorsAsShapesNeeded() []int32 {
	if m != nil {
		return m.InputTensorsAsShapesNeeded
	}
	return nil
}

func init() {
	proto.RegisterType((*CppShapeInferenceResult)(nil), "tensorflow.CppShapeInferenceResult")
	proto.RegisterType((*CppShapeInferenceResult_HandleShapeAndType)(nil), "tensorflow.CppShapeInferenceResult.HandleShapeAndType")
	proto.RegisterType((*CppShapeInferenceResult_HandleData)(nil), "tensorflow.CppShapeInferenceResult.HandleData")
	proto.RegisterType((*CppShapeInferenceInputsNeeded)(nil), "tensorflow.CppShapeInferenceInputsNeeded")
}

func init() {
	proto.RegisterFile("tensorflow/python/framework/cpp_shape_inference.proto", fileDescriptor_e2bf6bfecd64d111)
}

var fileDescriptor_e2bf6bfecd64d111 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x5f, 0xcb, 0xd3, 0x30,
	0x14, 0xc6, 0xc9, 0xda, 0x8e, 0x91, 0xc9, 0x18, 0x61, 0x62, 0x29, 0x2a, 0x63, 0x20, 0x14, 0x91,
	0x56, 0x2a, 0x7a, 0xbf, 0xb9, 0x0b, 0xb7, 0x0b, 0x95, 0x6e, 0x57, 0x22, 0x94, 0xac, 0x3d, 0x5b,
	0x8b, 0x5b, 0x12, 0x9a, 0x8c, 0xb1, 0x4b, 0x3f, 0x84, 0xdf, 0xc0, 0x0f, 0xe9, 0xa5, 0x34, 0xd9,
	0x9f, 0xe8, 0xfb, 0x87, 0x97, 0xf7, 0x2e, 0xc9, 0x79, 0x7e, 0xcf, 0x73, 0x92, 0x9e, 0xe2, 0xf7,
	0x0a, 0x98, 0xe4, 0xf5, 0x7a, 0xcb, 0x0f, 0xb1, 0x38, 0xaa, 0x92, 0xb3, 0x78, 0x5d, 0xd3, 0x1d,
	0x1c, 0x78, 0xfd, 0x23, 0xce, 0x85, 0xc8, 0x64, 0x49, 0x05, 0x64, 0x15, 0x5b, 0x43, 0x0d, 0x2c,
	0x87, 0x48, 0xd4, 0x5c, 0x71, 0x82, 0xaf, 0x58, 0xf0, 0xca, 0xb2, 0xc8, 0x79, 0x0d, 0x96, 0x81,
	0x3a, 0x0a, 0x90, 0x06, 0x09, 0xde, 0xdc, 0x23, 0xd3, 0x15, 0x13, 0x65, 0xd4, 0xa3, 0xdf, 0x0e,
	0x7e, 0xf6, 0x51, 0x88, 0x45, 0x73, 0x34, 0x3b, 0x87, 0xa7, 0x20, 0xf7, 0x5b, 0x45, 0x12, 0xec,
	0x69, 0xa9, 0x8f, 0x86, 0x28, 0xec, 0x26, 0xcf, 0xa3, 0xab, 0x73, 0xb4, 0xd4, 0x4b, 0x8d, 0x7d,
	0x6d, 0x8c, 0x52, 0x23, 0x25, 0x5f, 0x70, 0xb7, 0xa4, 0xac, 0xd8, 0x42, 0x56, 0x50, 0x45, 0x7d,
	0x57, 0x93, 0x91, 0x4d, 0xde, 0x91, 0x16, 0x7d, 0xd2, 0xd8, 0x94, 0x2a, 0x9a, 0xe2, 0xf2, 0xb2,
	0x0e, 0x14, 0x26, 0xa6, 0xa2, 0xa1, 0x31, 0x2b, 0x96, 0x47, 0x01, 0x8f, 0x6a, 0xed, 0x35, 0xf6,
	0x8a, 0xe6, 0xa1, 0xfc, 0xd6, 0x10, 0x85, 0xbd, 0x64, 0x60, 0x33, 0x4d, 0x54, 0x63, 0x9c, 0x1a,
	0x49, 0xf0, 0x13, 0x61, 0x7c, 0x6d, 0x88, 0x3c, 0xc5, 0xed, 0x4a, 0x66, 0x12, 0x94, 0xce, 0xeb,
	0xa4, 0x5e, 0x25, 0x17, 0xa0, 0xc8, 0x77, 0xdc, 0x33, 0x9f, 0x8d, 0xb2, 0x22, 0x3b, 0x59, 0x3b,
	0x61, 0x37, 0xf9, 0xf0, 0xf0, 0xfb, 0xda, 0xb7, 0x4a, 0x9f, 0x48, 0x6b, 0x37, 0x77, 0x3b, 0xad,
	0xbe, 0x33, 0x77, 0x3b, 0x4e, 0xdf, 0x1d, 0xfd, 0x42, 0xf8, 0xc5, 0x0d, 0xa3, 0x19, 0x13, 0x7b,
	0x25, 0x3f, 0x03, 0x14, 0x50, 0x90, 0xb7, 0x78, 0x50, 0x35, 0xfb, 0xcc, 0x44, 0xcb, 0x8c, 0xe9,
	0x73, 0x1f, 0x0d, 0x9d, 0xd0, 0x4b, 0x89, 0xae, 0x99, 0x97, 0x39, 0x13, 0x13, 0xfc, 0xf2, 0x5f,
	0x82, 0x4a, 0x33, 0x1a, 0x17, 0xb6, 0xa5, 0xd9, 0xc0, 0x66, 0xc7, 0x52, 0xf7, 0x70, 0xf2, 0x98,
	0xcc, 0xbe, 0x4d, 0x37, 0x95, 0x2a, 0xf7, 0xab, 0x28, 0xe7, 0xbb, 0xd8, 0x9a, 0xbc, 0xdb, 0x97,
	0x1b, 0x6e, 0x46, 0xf2, 0xff, 0x3f, 0xe0, 0x0f, 0x42, 0xab, 0xb6, 0x1e, 0xc8, 0x77, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x7f, 0x5d, 0x86, 0x65, 0x2a, 0x03, 0x00, 0x00,
}
