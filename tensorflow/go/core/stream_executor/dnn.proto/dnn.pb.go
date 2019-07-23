// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/stream_executor/dnn.proto

package dnn_proto

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

// Specifies the data type used by an operation.
type DataType int32

const (
	DataType_kFloat  DataType = 0
	DataType_kDouble DataType = 1
	DataType_kHalf   DataType = 2
	DataType_kInt8   DataType = 3
	DataType_kInt32  DataType = 4
)

var DataType_name = map[int32]string{
	0: "kFloat",
	1: "kDouble",
	2: "kHalf",
	3: "kInt8",
	4: "kInt32",
}

var DataType_value = map[string]int32{
	"kFloat":  0,
	"kDouble": 1,
	"kHalf":   2,
	"kInt8":   3,
	"kInt32":  4,
}

func (x DataType) String() string {
	return proto.EnumName(DataType_name, int32(x))
}

func (DataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_768c61f2a579ee6a, []int{0}
}

// Describes how a convolution input or output layer's data is formatted.
type DataLayout int32

const (
	// Naming convention:
	// Y <-> row or height
	// X <-> column or width
	// Batch <-> batch, or N
	// Depth <-> feature, or channel
	// TODO(timshen): turn them into cuDNN names, e.g. kNCHW.
	DataLayout_kYXDepthBatch  DataLayout = 0
	DataLayout_kYXBatchDepth  DataLayout = 1
	DataLayout_kBatchYXDepth  DataLayout = 2
	DataLayout_kBatchDepthYX  DataLayout = 3
	DataLayout_kBatchDepthYX4 DataLayout = 4
)

var DataLayout_name = map[int32]string{
	0: "kYXDepthBatch",
	1: "kYXBatchDepth",
	2: "kBatchYXDepth",
	3: "kBatchDepthYX",
	4: "kBatchDepthYX4",
}

var DataLayout_value = map[string]int32{
	"kYXDepthBatch":  0,
	"kYXBatchDepth":  1,
	"kBatchYXDepth":  2,
	"kBatchDepthYX":  3,
	"kBatchDepthYX4": 4,
}

func (x DataLayout) String() string {
	return proto.EnumName(DataLayout_name, int32(x))
}

func (DataLayout) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_768c61f2a579ee6a, []int{1}
}

// Describes how a convolution filter is laid out in the memory.
type FilterLayout int32

const (
	// Naming convention:
	// Y <-> row or height
	// X <-> column or width
	// Output <-> output feature, or N
	// Input <-> input feature, or N
	// TODO(timshen): turn them into cuDNN names, e.g. kNCHW.
	FilterLayout_kOutputInputYX  FilterLayout = 0
	FilterLayout_kOutputYXInput  FilterLayout = 1
	FilterLayout_kOutputInputYX4 FilterLayout = 2
	FilterLayout_kInputYXOutput  FilterLayout = 3
	FilterLayout_kYXInputOutput  FilterLayout = 4
)

var FilterLayout_name = map[int32]string{
	0: "kOutputInputYX",
	1: "kOutputYXInput",
	2: "kOutputInputYX4",
	3: "kInputYXOutput",
	4: "kYXInputOutput",
}

var FilterLayout_value = map[string]int32{
	"kOutputInputYX":  0,
	"kOutputYXInput":  1,
	"kOutputInputYX4": 2,
	"kInputYXOutput":  3,
	"kYXInputOutput":  4,
}

func (x FilterLayout) String() string {
	return proto.EnumName(FilterLayout_name, int32(x))
}

func (FilterLayout) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_768c61f2a579ee6a, []int{2}
}

// Describes a kind of non-linearity (threshold-like mathematical function).
type ActivationMode int32

const (
	ActivationMode_kNone    ActivationMode = 0
	ActivationMode_kSigmoid ActivationMode = 1
	// Rectified linear activation: f(x) = x < 0 ? 0 : x
	ActivationMode_kRelu ActivationMode = 2
	// Rectified linear activation; where upper maximum is 6.0.
	ActivationMode_kRelu6 ActivationMode = 3
	// Rectified linear activation; where upper maximum specified by
	// BatchDescriptor::value_max().
	ActivationMode_kReluX ActivationMode = 4
	ActivationMode_kTanh  ActivationMode = 5
	// Like ReluX; but passes all values in the range [-X,X].
	ActivationMode_kBandPass ActivationMode = 6
)

var ActivationMode_name = map[int32]string{
	0: "kNone",
	1: "kSigmoid",
	2: "kRelu",
	3: "kRelu6",
	4: "kReluX",
	5: "kTanh",
	6: "kBandPass",
}

var ActivationMode_value = map[string]int32{
	"kNone":     0,
	"kSigmoid":  1,
	"kRelu":     2,
	"kRelu6":    3,
	"kReluX":    4,
	"kTanh":     5,
	"kBandPass": 6,
}

func (x ActivationMode) String() string {
	return proto.EnumName(ActivationMode_name, int32(x))
}

func (ActivationMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_768c61f2a579ee6a, []int{3}
}

// Describe the math definition for the conv op. The popular behavior is
// actually called cross-correlation in math, despite the operation is often
// referred as convolution. See cuDNN cudnnConvolutionMode_t.
type ConvolutionMode int32

const (
	ConvolutionMode_CROSS_CORRELATION ConvolutionMode = 0
	ConvolutionMode_CONVOLUTION       ConvolutionMode = 1
)

var ConvolutionMode_name = map[int32]string{
	0: "CROSS_CORRELATION",
	1: "CONVOLUTION",
}

var ConvolutionMode_value = map[string]int32{
	"CROSS_CORRELATION": 0,
	"CONVOLUTION":       1,
}

func (x ConvolutionMode) String() string {
	return proto.EnumName(ConvolutionMode_name, int32(x))
}

func (ConvolutionMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_768c61f2a579ee6a, []int{4}
}

type ConvolutionKind int32

const (
	ConvolutionKind_INVALID                 ConvolutionKind = 0
	ConvolutionKind_FORWARD                 ConvolutionKind = 1
	ConvolutionKind_BACKWARD_FILTER         ConvolutionKind = 2
	ConvolutionKind_BACKWARD_DATA           ConvolutionKind = 3
	ConvolutionKind_FORWARD_BIAS_ACTIVATION ConvolutionKind = 4
)

var ConvolutionKind_name = map[int32]string{
	0: "INVALID",
	1: "FORWARD",
	2: "BACKWARD_FILTER",
	3: "BACKWARD_DATA",
	4: "FORWARD_BIAS_ACTIVATION",
}

var ConvolutionKind_value = map[string]int32{
	"INVALID":                 0,
	"FORWARD":                 1,
	"BACKWARD_FILTER":         2,
	"BACKWARD_DATA":           3,
	"FORWARD_BIAS_ACTIVATION": 4,
}

func (x ConvolutionKind) String() string {
	return proto.EnumName(ConvolutionKind_name, int32(x))
}

func (ConvolutionKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_768c61f2a579ee6a, []int{5}
}

type AlgorithmProto_MathType int32

const (
	AlgorithmProto_DEFAULT_MATH AlgorithmProto_MathType = 0
	// The GPU may operate 4x4 matrix FMA.
	// See cuDNN's documentation for CUDNN_TENSOR_OP_MATH.
	AlgorithmProto_TENSOR_OP_MATH AlgorithmProto_MathType = 1
)

var AlgorithmProto_MathType_name = map[int32]string{
	0: "DEFAULT_MATH",
	1: "TENSOR_OP_MATH",
}

var AlgorithmProto_MathType_value = map[string]int32{
	"DEFAULT_MATH":   0,
	"TENSOR_OP_MATH": 1,
}

func (x AlgorithmProto_MathType) String() string {
	return proto.EnumName(AlgorithmProto_MathType_name, int32(x))
}

func (AlgorithmProto_MathType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_768c61f2a579ee6a, []int{1, 0}
}

// Generic tensor representation.
type TensorDescriptorProto struct {
	Dimensions []int64  `protobuf:"varint,1,rep,packed,name=dimensions,proto3" json:"dimensions,omitempty"`
	DataType   DataType `protobuf:"varint,2,opt,name=data_type,json=dataType,proto3,enum=stream_executor.dnn.DataType" json:"data_type,omitempty"`
	// Types that are valid to be assigned to LayoutOneof:
	//	*TensorDescriptorProto_DataLayout
	//	*TensorDescriptorProto_FilterLayout
	LayoutOneof          isTensorDescriptorProto_LayoutOneof `protobuf_oneof:"layout_oneof"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *TensorDescriptorProto) Reset()         { *m = TensorDescriptorProto{} }
func (m *TensorDescriptorProto) String() string { return proto.CompactTextString(m) }
func (*TensorDescriptorProto) ProtoMessage()    {}
func (*TensorDescriptorProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_768c61f2a579ee6a, []int{0}
}

func (m *TensorDescriptorProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TensorDescriptorProto.Unmarshal(m, b)
}
func (m *TensorDescriptorProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TensorDescriptorProto.Marshal(b, m, deterministic)
}
func (m *TensorDescriptorProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TensorDescriptorProto.Merge(m, src)
}
func (m *TensorDescriptorProto) XXX_Size() int {
	return xxx_messageInfo_TensorDescriptorProto.Size(m)
}
func (m *TensorDescriptorProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TensorDescriptorProto.DiscardUnknown(m)
}

var xxx_messageInfo_TensorDescriptorProto proto.InternalMessageInfo

func (m *TensorDescriptorProto) GetDimensions() []int64 {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *TensorDescriptorProto) GetDataType() DataType {
	if m != nil {
		return m.DataType
	}
	return DataType_kFloat
}

type isTensorDescriptorProto_LayoutOneof interface {
	isTensorDescriptorProto_LayoutOneof()
}

type TensorDescriptorProto_DataLayout struct {
	DataLayout DataLayout `protobuf:"varint,3,opt,name=data_layout,json=dataLayout,proto3,enum=stream_executor.dnn.DataLayout,oneof"`
}

type TensorDescriptorProto_FilterLayout struct {
	FilterLayout FilterLayout `protobuf:"varint,4,opt,name=filter_layout,json=filterLayout,proto3,enum=stream_executor.dnn.FilterLayout,oneof"`
}

func (*TensorDescriptorProto_DataLayout) isTensorDescriptorProto_LayoutOneof() {}

func (*TensorDescriptorProto_FilterLayout) isTensorDescriptorProto_LayoutOneof() {}

func (m *TensorDescriptorProto) GetLayoutOneof() isTensorDescriptorProto_LayoutOneof {
	if m != nil {
		return m.LayoutOneof
	}
	return nil
}

func (m *TensorDescriptorProto) GetDataLayout() DataLayout {
	if x, ok := m.GetLayoutOneof().(*TensorDescriptorProto_DataLayout); ok {
		return x.DataLayout
	}
	return DataLayout_kYXDepthBatch
}

func (m *TensorDescriptorProto) GetFilterLayout() FilterLayout {
	if x, ok := m.GetLayoutOneof().(*TensorDescriptorProto_FilterLayout); ok {
		return x.FilterLayout
	}
	return FilterLayout_kOutputInputYX
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TensorDescriptorProto) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TensorDescriptorProto_DataLayout)(nil),
		(*TensorDescriptorProto_FilterLayout)(nil),
	}
}

// Generic algorithm representation.
type AlgorithmProto struct {
	AlgoId               int64                   `protobuf:"varint,1,opt,name=algo_id,json=algoId,proto3" json:"algo_id,omitempty"`
	MathType             AlgorithmProto_MathType `protobuf:"varint,2,opt,name=math_type,json=mathType,proto3,enum=stream_executor.dnn.AlgorithmProto_MathType" json:"math_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AlgorithmProto) Reset()         { *m = AlgorithmProto{} }
func (m *AlgorithmProto) String() string { return proto.CompactTextString(m) }
func (*AlgorithmProto) ProtoMessage()    {}
func (*AlgorithmProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_768c61f2a579ee6a, []int{1}
}

func (m *AlgorithmProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlgorithmProto.Unmarshal(m, b)
}
func (m *AlgorithmProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlgorithmProto.Marshal(b, m, deterministic)
}
func (m *AlgorithmProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlgorithmProto.Merge(m, src)
}
func (m *AlgorithmProto) XXX_Size() int {
	return xxx_messageInfo_AlgorithmProto.Size(m)
}
func (m *AlgorithmProto) XXX_DiscardUnknown() {
	xxx_messageInfo_AlgorithmProto.DiscardUnknown(m)
}

var xxx_messageInfo_AlgorithmProto proto.InternalMessageInfo

func (m *AlgorithmProto) GetAlgoId() int64 {
	if m != nil {
		return m.AlgoId
	}
	return 0
}

func (m *AlgorithmProto) GetMathType() AlgorithmProto_MathType {
	if m != nil {
		return m.MathType
	}
	return AlgorithmProto_DEFAULT_MATH
}

// Convolution-specific parameters.
type ConvolutionDescriptorProto struct {
	Paddings  []int64 `protobuf:"varint,1,rep,packed,name=paddings,proto3" json:"paddings,omitempty"`
	Strides   []int64 `protobuf:"varint,2,rep,packed,name=strides,proto3" json:"strides,omitempty"`
	Dilations []int64 `protobuf:"varint,3,rep,packed,name=dilations,proto3" json:"dilations,omitempty"`
	// The "accumulator" type. For example, use F32 as an accumulator for F16
	// convolutions.
	// See cuDNN's cudnnConvolutionMode_t.
	ComputeMode DataType `protobuf:"varint,4,opt,name=compute_mode,json=computeMode,proto3,enum=stream_executor.dnn.DataType" json:"compute_mode,omitempty"`
	// See cuDNN's group count.
	GroupCount      int32           `protobuf:"varint,5,opt,name=group_count,json=groupCount,proto3" json:"group_count,omitempty"`
	ConvolutionMode ConvolutionMode `protobuf:"varint,6,opt,name=convolution_mode,json=convolutionMode,proto3,enum=stream_executor.dnn.ConvolutionMode" json:"convolution_mode,omitempty"`
	// Tensorflow node name, same as in NodeDef, for debugging purposes.
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConvolutionDescriptorProto) Reset()         { *m = ConvolutionDescriptorProto{} }
func (m *ConvolutionDescriptorProto) String() string { return proto.CompactTextString(m) }
func (*ConvolutionDescriptorProto) ProtoMessage()    {}
func (*ConvolutionDescriptorProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_768c61f2a579ee6a, []int{2}
}

func (m *ConvolutionDescriptorProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConvolutionDescriptorProto.Unmarshal(m, b)
}
func (m *ConvolutionDescriptorProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConvolutionDescriptorProto.Marshal(b, m, deterministic)
}
func (m *ConvolutionDescriptorProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvolutionDescriptorProto.Merge(m, src)
}
func (m *ConvolutionDescriptorProto) XXX_Size() int {
	return xxx_messageInfo_ConvolutionDescriptorProto.Size(m)
}
func (m *ConvolutionDescriptorProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvolutionDescriptorProto.DiscardUnknown(m)
}

var xxx_messageInfo_ConvolutionDescriptorProto proto.InternalMessageInfo

func (m *ConvolutionDescriptorProto) GetPaddings() []int64 {
	if m != nil {
		return m.Paddings
	}
	return nil
}

func (m *ConvolutionDescriptorProto) GetStrides() []int64 {
	if m != nil {
		return m.Strides
	}
	return nil
}

func (m *ConvolutionDescriptorProto) GetDilations() []int64 {
	if m != nil {
		return m.Dilations
	}
	return nil
}

func (m *ConvolutionDescriptorProto) GetComputeMode() DataType {
	if m != nil {
		return m.ComputeMode
	}
	return DataType_kFloat
}

func (m *ConvolutionDescriptorProto) GetGroupCount() int32 {
	if m != nil {
		return m.GroupCount
	}
	return 0
}

func (m *ConvolutionDescriptorProto) GetConvolutionMode() ConvolutionMode {
	if m != nil {
		return m.ConvolutionMode
	}
	return ConvolutionMode_CROSS_CORRELATION
}

func (m *ConvolutionDescriptorProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterEnum("stream_executor.dnn.DataType", DataType_name, DataType_value)
	proto.RegisterEnum("stream_executor.dnn.DataLayout", DataLayout_name, DataLayout_value)
	proto.RegisterEnum("stream_executor.dnn.FilterLayout", FilterLayout_name, FilterLayout_value)
	proto.RegisterEnum("stream_executor.dnn.ActivationMode", ActivationMode_name, ActivationMode_value)
	proto.RegisterEnum("stream_executor.dnn.ConvolutionMode", ConvolutionMode_name, ConvolutionMode_value)
	proto.RegisterEnum("stream_executor.dnn.ConvolutionKind", ConvolutionKind_name, ConvolutionKind_value)
	proto.RegisterEnum("stream_executor.dnn.AlgorithmProto_MathType", AlgorithmProto_MathType_name, AlgorithmProto_MathType_value)
	proto.RegisterType((*TensorDescriptorProto)(nil), "stream_executor.dnn.TensorDescriptorProto")
	proto.RegisterType((*AlgorithmProto)(nil), "stream_executor.dnn.AlgorithmProto")
	proto.RegisterType((*ConvolutionDescriptorProto)(nil), "stream_executor.dnn.ConvolutionDescriptorProto")
}

func init() {
	proto.RegisterFile("tensorflow/stream_executor/dnn.proto", fileDescriptor_768c61f2a579ee6a)
}

var fileDescriptor_768c61f2a579ee6a = []byte{
	// 785 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0xc7, 0x31, 0x10, 0x02, 0x07, 0x42, 0x66, 0x67, 0xb5, 0x5a, 0xb4, 0xfd, 0x58, 0x1a, 0xed,
	0x05, 0x42, 0x15, 0xa9, 0x76, 0x57, 0x55, 0xdb, 0xab, 0x1a, 0x0c, 0x8a, 0xb5, 0x04, 0x47, 0xc6,
	0x9b, 0x92, 0xde, 0x58, 0x13, 0xcf, 0x00, 0x16, 0xf6, 0x8c, 0x65, 0x8f, 0xb7, 0xcd, 0x63, 0xf4,
	0x25, 0xfa, 0x0c, 0x7d, 0xbc, 0xca, 0x63, 0x43, 0x20, 0x22, 0xda, 0x2b, 0xce, 0xf9, 0xcd, 0xf9,
	0x18, 0xfe, 0xff, 0x91, 0xe1, 0x9d, 0x64, 0x3c, 0x11, 0xf1, 0x32, 0x10, 0x7f, 0x5d, 0x26, 0x32,
	0x66, 0x24, 0x74, 0xd9, 0xdf, 0xcc, 0x4b, 0xa5, 0x88, 0x2f, 0x29, 0xe7, 0x83, 0x28, 0x16, 0x52,
	0xe0, 0x97, 0x4f, 0x8e, 0x06, 0x94, 0xf3, 0x8b, 0x7f, 0xca, 0xf0, 0xca, 0x51, 0xdd, 0x06, 0x4b,
	0xbc, 0xd8, 0x8f, 0xa4, 0x88, 0x6f, 0x54, 0xf9, 0xf7, 0x00, 0xd4, 0x0f, 0x19, 0x4f, 0x7c, 0xc1,
	0x93, 0x8e, 0xd6, 0xad, 0xf4, 0x2a, 0xf6, 0x1e, 0xc1, 0xbf, 0x41, 0x83, 0x12, 0x49, 0x5c, 0xf9,
	0x10, 0xb1, 0x4e, 0xb9, 0xab, 0xf5, 0xda, 0xef, 0xbf, 0x1b, 0x1c, 0x59, 0x31, 0x30, 0x88, 0x24,
	0xce, 0x43, 0xc4, 0xec, 0x3a, 0x2d, 0x22, 0x3c, 0x84, 0xa6, 0xea, 0x0d, 0xc8, 0x83, 0x48, 0x65,
	0xa7, 0xa2, 0xba, 0xdf, 0x3e, 0xdb, 0x3d, 0x55, 0x65, 0x57, 0x25, 0x1b, 0xe8, 0x2e, 0xc3, 0x57,
	0x70, 0xb6, 0xf4, 0x03, 0xc9, 0xe2, 0xed, 0x94, 0xaa, 0x9a, 0xf2, 0xc3, 0xd1, 0x29, 0x13, 0x55,
	0xb9, 0x9b, 0xd3, 0x5a, 0xee, 0xe5, 0xc3, 0x36, 0xb4, 0xf2, 0x11, 0xae, 0xe0, 0x4c, 0x2c, 0x2f,
	0xfe, 0xd5, 0xa0, 0xad, 0x07, 0x2b, 0x11, 0xfb, 0x72, 0x1d, 0xe6, 0x62, 0xbc, 0x86, 0x53, 0x12,
	0xac, 0x84, 0xeb, 0xd3, 0x8e, 0xd6, 0xd5, 0x7a, 0x15, 0xbb, 0x96, 0xa5, 0x26, 0xc5, 0x26, 0x34,
	0x42, 0x22, 0xd7, 0xfb, 0x2a, 0xfc, 0x78, 0xf4, 0x06, 0x87, 0x03, 0x07, 0xd7, 0x44, 0xae, 0x73,
	0x51, 0xc2, 0x22, 0xba, 0xf8, 0x09, 0xea, 0x5b, 0x8a, 0x11, 0xb4, 0x8c, 0xf1, 0x44, 0xff, 0x3c,
	0x75, 0xdc, 0x6b, 0xdd, 0xb9, 0x42, 0x25, 0x8c, 0xa1, 0xed, 0x8c, 0x67, 0x73, 0xcb, 0x76, 0xad,
	0x9b, 0x9c, 0x69, 0x17, 0xff, 0x95, 0xe1, 0xcd, 0x48, 0xf0, 0x2f, 0x22, 0x48, 0xa5, 0x2f, 0xf8,
	0x53, 0x07, 0xdf, 0x40, 0x3d, 0x22, 0x94, 0xfa, 0x7c, 0xb5, 0xf5, 0x6f, 0x97, 0xe3, 0x0e, 0x9c,
	0x26, 0x32, 0xf6, 0x29, 0x4b, 0x3a, 0x65, 0x75, 0xb4, 0x4d, 0xf1, 0xb7, 0xd0, 0xa0, 0x7e, 0x40,
	0xa4, 0xb2, 0xbd, 0xa2, 0xce, 0x1e, 0x01, 0xfe, 0x1d, 0x5a, 0x9e, 0x08, 0xa3, 0x54, 0x32, 0x37,
	0x14, 0x94, 0x15, 0xa2, 0x7f, 0xc5, 0xf8, 0x66, 0xd1, 0x72, 0x2d, 0x28, 0xc3, 0x6f, 0xa1, 0xb9,
	0x8a, 0x45, 0x1a, 0xb9, 0x9e, 0x48, 0xb9, 0xec, 0x9c, 0x74, 0xb5, 0xde, 0x89, 0x0d, 0x0a, 0x8d,
	0x32, 0x82, 0x2d, 0x40, 0xde, 0xe3, 0x9f, 0xca, 0xd7, 0xd4, 0xd4, 0x9a, 0x77, 0x47, 0xd7, 0xec,
	0x29, 0x90, 0x2d, 0xb0, 0xcf, 0xbd, 0x43, 0x80, 0x31, 0x54, 0x39, 0x09, 0x59, 0xe7, 0xb4, 0xab,
	0xf5, 0x1a, 0xb6, 0x8a, 0xfb, 0x63, 0xa8, 0x6f, 0xaf, 0x87, 0x01, 0x6a, 0x9b, 0x49, 0x20, 0x88,
	0x44, 0x25, 0xdc, 0x84, 0xd3, 0x8d, 0x21, 0xd2, 0xfb, 0x80, 0x21, 0x0d, 0x37, 0xe0, 0x64, 0x73,
	0x45, 0x82, 0x25, 0x2a, 0xab, 0xd0, 0xe4, 0xf2, 0x17, 0x54, 0x51, 0xe5, 0x26, 0x97, 0x1f, 0xde,
	0xa3, 0x6a, 0x3f, 0x00, 0x78, 0x7c, 0xa0, 0xf8, 0x05, 0x9c, 0x6d, 0xee, 0x16, 0x06, 0x8b, 0xe4,
	0x7a, 0x48, 0xa4, 0xb7, 0x46, 0xa5, 0x02, 0xa9, 0x4c, 0x71, 0xa4, 0x29, 0xa4, 0x40, 0x51, 0x8a,
	0xca, 0x8f, 0x48, 0x81, 0xbb, 0x05, 0xaa, 0x64, 0x7e, 0x1f, 0xa0, 0x8f, 0xa8, 0xda, 0x4f, 0xa0,
	0xb5, 0xff, 0x90, 0x55, 0x8d, 0x95, 0xca, 0x28, 0x95, 0x26, 0x8f, 0x52, 0x79, 0xb7, 0xc8, 0xdf,
	0x49, 0xc1, 0xee, 0x16, 0x8a, 0x22, 0x0d, 0xbf, 0x84, 0xf3, 0xc3, 0xba, 0x8f, 0xa8, 0xac, 0x0a,
	0x8b, 0x34, 0x3f, 0x2b, 0x96, 0x16, 0x6d, 0x05, 0xab, 0xf6, 0x97, 0xd0, 0xd6, 0x3d, 0xe9, 0x7f,
	0x21, 0x3b, 0x3d, 0x33, 0x2d, 0x66, 0x82, 0x33, 0x54, 0xc2, 0x2d, 0xa8, 0x6f, 0xe6, 0xfe, 0x2a,
	0x14, 0x3e, 0x2d, 0xf4, 0xb2, 0x59, 0x90, 0xa2, 0xb2, 0x12, 0x29, 0x0b, 0x7f, 0x2e, 0x04, 0xcb,
	0xe2, 0x05, 0xaa, 0xaa, 0x12, 0x87, 0xf0, 0x35, 0x3a, 0xc1, 0x67, 0xd0, 0xd8, 0x0c, 0x09, 0xa7,
	0x37, 0x24, 0x49, 0x50, 0xad, 0xff, 0x2b, 0x9c, 0x3f, 0x71, 0x12, 0xbf, 0x82, 0x17, 0x23, 0xdb,
	0x9a, 0xcf, 0xdd, 0x91, 0x65, 0xdb, 0xe3, 0xa9, 0xee, 0x98, 0xd6, 0x0c, 0x95, 0xf0, 0x39, 0x34,
	0x47, 0xd6, 0xec, 0xd6, 0x9a, 0x7e, 0x56, 0x40, 0xeb, 0x47, 0x07, 0xad, 0x9f, 0x7c, 0x4e, 0x33,
	0x1f, 0xcd, 0xd9, 0xad, 0x3e, 0x35, 0x8d, 0xdc, 0xd4, 0x89, 0x65, 0xff, 0xa1, 0xdb, 0x46, 0x2e,
	0xc6, 0x50, 0x1f, 0x7d, 0xca, 0x32, 0x77, 0x62, 0x4e, 0x9d, 0xb1, 0x9d, 0x1b, 0xb0, 0x83, 0x86,
	0xee, 0xe8, 0xa8, 0x82, 0xbf, 0x81, 0xd7, 0x45, 0x93, 0x3b, 0x34, 0xf5, 0xb9, 0xab, 0x8f, 0x1c,
	0xf3, 0x36, 0xbf, 0x42, 0x75, 0x68, 0xfd, 0x79, 0xbd, 0xf2, 0xe5, 0x3a, 0xbd, 0x1f, 0x78, 0x22,
	0xbc, 0xdc, 0xfb, 0xfc, 0x1e, 0x0f, 0x57, 0xe2, 0xd2, 0x13, 0x31, 0x7b, 0xfe, 0xe3, 0x7c, 0x5f,
	0x53, 0x3f, 0x1f, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xc4, 0xdb, 0xce, 0xcb, 0x05, 0x00,
	0x00,
}
