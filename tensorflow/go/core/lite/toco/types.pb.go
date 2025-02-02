// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/lite/toco/types.proto

package toco

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

// IODataType describes the numeric data types of input and output arrays
// of a model.
type IODataType int32

const (
	IODataType_IO_DATA_TYPE_UNKNOWN IODataType = 0
	// Float32, not quantized
	IODataType_FLOAT IODataType = 1
	// Uint8, quantized
	IODataType_QUANTIZED_UINT8 IODataType = 2
	// Int32, not quantized
	IODataType_INT32 IODataType = 3
	// Int64, not quantized
	IODataType_INT64 IODataType = 4
	// String, not quantized
	IODataType_STRING IODataType = 5
	// Int16, quantized
	IODataType_QUANTIZED_INT16 IODataType = 6
	// Boolean
	IODataType_BOOL IODataType = 7
	// Complex64, not quantized
	IODataType_COMPLEX64 IODataType = 8
	// Int8, quantized based on QuantizationParameters in schema.
	IODataType_INT8 IODataType = 9
	// Half precision float, not quantized.
	IODataType_FLOAT16 IODataType = 10
)

var IODataType_name = map[int32]string{
	0:  "IO_DATA_TYPE_UNKNOWN",
	1:  "FLOAT",
	2:  "QUANTIZED_UINT8",
	3:  "INT32",
	4:  "INT64",
	5:  "STRING",
	6:  "QUANTIZED_INT16",
	7:  "BOOL",
	8:  "COMPLEX64",
	9:  "INT8",
	10: "FLOAT16",
}

var IODataType_value = map[string]int32{
	"IO_DATA_TYPE_UNKNOWN": 0,
	"FLOAT":                1,
	"QUANTIZED_UINT8":      2,
	"INT32":                3,
	"INT64":                4,
	"STRING":               5,
	"QUANTIZED_INT16":      6,
	"BOOL":                 7,
	"COMPLEX64":            8,
	"INT8":                 9,
	"FLOAT16":              10,
}

func (x IODataType) Enum() *IODataType {
	p := new(IODataType)
	*p = x
	return p
}

func (x IODataType) String() string {
	return proto.EnumName(IODataType_name, int32(x))
}

func (x *IODataType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(IODataType_value, data, "IODataType")
	if err != nil {
		return err
	}
	*x = IODataType(value)
	return nil
}

func (IODataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_50f97c981a4fd685, []int{0}
}

func init() {
	proto.RegisterEnum("toco.IODataType", IODataType_name, IODataType_value)
}

func init() { proto.RegisterFile("tensorflow/lite/toco/types.proto", fileDescriptor_50f97c981a4fd685) }

var fileDescriptor_50f97c981a4fd685 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0xd7, 0xcf, 0xc9, 0x2c, 0x49, 0xd5, 0x2f, 0xc9, 0x4f, 0xce,
	0xd7, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x89,
	0x68, 0xad, 0x65, 0xe4, 0xe2, 0xf2, 0xf4, 0x77, 0x49, 0x2c, 0x49, 0x0c, 0xa9, 0x2c, 0x48, 0x15,
	0x92, 0xe0, 0x12, 0xf1, 0xf4, 0x8f, 0x77, 0x71, 0x0c, 0x71, 0x8c, 0x0f, 0x89, 0x0c, 0x70, 0x8d,
	0x0f, 0xf5, 0xf3, 0xf6, 0xf3, 0x0f, 0xf7, 0x13, 0x60, 0x10, 0xe2, 0xe4, 0x62, 0x75, 0xf3, 0xf1,
	0x77, 0x0c, 0x11, 0x60, 0x14, 0x12, 0xe6, 0xe2, 0x0f, 0x0c, 0x75, 0xf4, 0x0b, 0xf1, 0x8c, 0x72,
	0x75, 0x89, 0x0f, 0xf5, 0xf4, 0x0b, 0xb1, 0x10, 0x60, 0x02, 0xc9, 0x7b, 0xfa, 0x85, 0x18, 0x1b,
	0x09, 0x30, 0x43, 0x99, 0x66, 0x26, 0x02, 0x2c, 0x42, 0x5c, 0x5c, 0x6c, 0xc1, 0x21, 0x41, 0x9e,
	0x7e, 0xee, 0x02, 0xac, 0xa8, 0xda, 0x3c, 0xfd, 0x42, 0x0c, 0xcd, 0x04, 0xd8, 0x84, 0x38, 0xb8,
	0x58, 0x9c, 0xfc, 0xfd, 0x7d, 0x04, 0xd8, 0x85, 0x78, 0xb9, 0x38, 0x9d, 0xfd, 0x7d, 0x03, 0x7c,
	0x5c, 0x23, 0xcc, 0x4c, 0x04, 0x38, 0x40, 0x12, 0x60, 0x93, 0x39, 0x85, 0xb8, 0xb9, 0xd8, 0xc1,
	0x36, 0x1b, 0x9a, 0x09, 0x70, 0x39, 0xd9, 0x47, 0xd9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9,
	0x25, 0xe7, 0xe7, 0xea, 0x23, 0x79, 0x12, 0x3b, 0x33, 0x3d, 0x5f, 0x3f, 0x39, 0xbf, 0x28, 0x15,
	0x11, 0x04, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0xd9, 0xac, 0xaa, 0x19, 0x01, 0x00, 0x00,
}
