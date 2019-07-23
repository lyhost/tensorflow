// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/lite/toco/toco_flags.proto

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

// Supported I/O file formats. Some formats may be input-only or output-only.
type FileFormat int32

const (
	FileFormat_FILE_FORMAT_UNKNOWN FileFormat = 0
	// GraphDef, third_party/tensorflow/core/framework/graph.proto
	FileFormat_TENSORFLOW_GRAPHDEF FileFormat = 1
	// Tensorflow's mobile inference model.
	// third_party/tensorflow/contrib/tflite/schema.fbs
	FileFormat_TFLITE FileFormat = 2
	// GraphViz
	// Export-only.
	FileFormat_GRAPHVIZ_DOT FileFormat = 3
)

var FileFormat_name = map[int32]string{
	0: "FILE_FORMAT_UNKNOWN",
	1: "TENSORFLOW_GRAPHDEF",
	2: "TFLITE",
	3: "GRAPHVIZ_DOT",
}

var FileFormat_value = map[string]int32{
	"FILE_FORMAT_UNKNOWN": 0,
	"TENSORFLOW_GRAPHDEF": 1,
	"TFLITE":              2,
	"GRAPHVIZ_DOT":        3,
}

func (x FileFormat) Enum() *FileFormat {
	p := new(FileFormat)
	*p = x
	return p
}

func (x FileFormat) String() string {
	return proto.EnumName(FileFormat_name, int32(x))
}

func (x *FileFormat) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FileFormat_value, data, "FileFormat")
	if err != nil {
		return err
	}
	*x = FileFormat(value)
	return nil
}

func (FileFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_655e1383a1a98309, []int{0}
}

// TocoFlags encodes extra parameters that drive tooling operations, that
// are not normally encoded in model files and in general may not be thought
// of as properties of models, instead describing how models are to be
// processed in the context of the present tooling job.
//
// Next ID to use: 31.
type TocoFlags struct {
	// Input file format
	InputFormat *FileFormat `protobuf:"varint,1,opt,name=input_format,json=inputFormat,enum=toco.FileFormat" json:"input_format,omitempty"`
	// Output file format
	OutputFormat *FileFormat `protobuf:"varint,2,opt,name=output_format,json=outputFormat,enum=toco.FileFormat" json:"output_format,omitempty"`
	// Similar to inference_type, but allows to control specifically the
	// quantization of input arrays, separately from other arrays.
	//
	// If not set, then the value of inference_type is implicitly used, i.e.
	// by default input arrays are quantized like other arrays.
	//
	// Like inference_type, this only affects real-number arrays. By "real-number"
	// we mean float arrays, and quantized arrays. This excludes plain
	// integer arrays, strings arrays, and every other data type.
	//
	// The typical use for this flag is for vision models taking a bitmap
	// as input, typically with uint8 channels, yet still requiring floating-point
	// inference. For such image models, the uint8 input is quantized, i.e.
	// the uint8 values are interpreted as real numbers, and the quantization
	// parameters used for such input arrays are their mean_value, std_value
	// parameters.
	InferenceInputType *IODataType `protobuf:"varint,11,opt,name=inference_input_type,json=inferenceInputType,enum=toco.IODataType" json:"inference_input_type,omitempty"`
	// Sets the type of real-number arrays in the output file, that is, controls
	// the representation (quantization) of real numbers in the output file,
	// except for input arrays, which are controlled by inference_input_type.
	//
	// NOTE: this flag only impacts real-number arrays. By "real-number"
	// we mean float arrays, and quantized arrays. This excludes plain
	// integer arrays, strings arrays, and every other data type.
	//
	// For real-number arrays, the impact of this flag is to allow the output
	// file to choose a different real-numbers representation (quantization)
	// from what the input file used. For any other types of arrays, changing
	// the data type would not make sense.
	//
	// Specifically:
	//    - If FLOAT, then real-numbers arrays will be of type float in
	//      the output file. If they were quantized in the input file, then
	//      they get dequantized.
	//    - If QUANTIZED_UINT8, then real-numbers arrays will be quantized
	//      as uint8 in the output file. If they were float in the input file,
	//      then they get quantized.
	//    - If not set, then all real-numbers arrays retain the same type in the
	//      output file as they have in the input file.
	//
	InferenceType *IODataType `protobuf:"varint,4,opt,name=inference_type,json=inferenceType,enum=toco.IODataType" json:"inference_type,omitempty"`
	// default_ranges_min and default_ranges_max are helpers to experiment
	// with quantization of models. Normally, quantization requires the input
	// model to have (min, max) range information for every activations array.
	// This is needed in order to know how to quantize arrays and still achieve
	// satisfactory accuracy. However, in some circumstances one would just like
	// to estimate the performance of quantized inference, without caring about
	// accuracy. That is what default_ranges_min and default_ranges_max are for:
	// when specified, they will be used as default (min, max) range boundaries
	// for all activation arrays that lack (min, max) range information, thus
	// allowing for quantization to proceed.
	//
	// It should be clear from the above explanation that these parameters are
	// for experimentation purposes only and should not be used in production:
	// they make it easy to quantize models, but the resulting quantized model
	// will be inaccurate.
	//
	// These values only apply to arrays quantized with the kUint8 data type.
	DefaultRangesMin *float32 `protobuf:"fixed32,5,opt,name=default_ranges_min,json=defaultRangesMin" json:"default_ranges_min,omitempty"`
	DefaultRangesMax *float32 `protobuf:"fixed32,6,opt,name=default_ranges_max,json=defaultRangesMax" json:"default_ranges_max,omitempty"`
	// Equivalent versions of default_ranges_min/_max for arrays quantized with
	// the kInt16 data type.
	DefaultInt16RangesMin *float32 `protobuf:"fixed32,15,opt,name=default_int16_ranges_min,json=defaultInt16RangesMin" json:"default_int16_ranges_min,omitempty"`
	DefaultInt16RangesMax *float32 `protobuf:"fixed32,16,opt,name=default_int16_ranges_max,json=defaultInt16RangesMax" json:"default_int16_ranges_max,omitempty"`
	// Ignore and discard FakeQuant nodes. For instance, that can be used to
	// generate plain float code without fake-quantization from a quantized
	// graph.
	DropFakeQuant *bool `protobuf:"varint,7,opt,name=drop_fake_quant,json=dropFakeQuant" json:"drop_fake_quant,omitempty"`
	// Normally, FakeQuant nodes must be strict boundaries for graph
	// transformations, in order to ensure that quantized inference has the
	// exact same arithmetic behavior as quantized training --- which is the
	// whole point of quantized training and of FakeQuant nodes in the first
	// place. However, that entails subtle requirements on where exactly
	// FakeQuant nodes must be placed in the graph. Some quantized graphs
	// have FakeQuant nodes at unexpected locations, that prevent graph
	// transformations that are necessary in order to generate inference
	// code for these graphs. Such graphs should be fixed, but as a
	// temporary work-around, setting this reorder_across_fake_quant flag
	// allows toco to perform necessary graph transformations on them,
	// at the cost of no longer faithfully matching inference and training
	// arithmetic.
	ReorderAcrossFakeQuant *bool `protobuf:"varint,8,opt,name=reorder_across_fake_quant,json=reorderAcrossFakeQuant" json:"reorder_across_fake_quant,omitempty"`
	// If true, allow TOCO to create TF Lite Custom operators for all the
	// unsupported Tensorflow ops.
	AllowCustomOps *bool `protobuf:"varint,10,opt,name=allow_custom_ops,json=allowCustomOps" json:"allow_custom_ops,omitempty"`
	// Applies only to the case when the input format is TENSORFLOW_GRAPHDEF.
	// If true, then control dependencies will be immediately dropped during
	// import.
	// If not set, the default behavior is as follows:
	//    - Default to false if the output format is TENSORFLOW_GRAPHDEF.
	//    - Default to true in all other cases.
	DropControlDependency *bool `protobuf:"varint,12,opt,name=drop_control_dependency,json=dropControlDependency" json:"drop_control_dependency,omitempty"`
	// Disables transformations that fuse subgraphs such as known LSTMs (not all
	// LSTMs are identified).
	DebugDisableRecurrentCellFusion *bool `protobuf:"varint,13,opt,name=debug_disable_recurrent_cell_fusion,json=debugDisableRecurrentCellFusion" json:"debug_disable_recurrent_cell_fusion,omitempty"`
	// Uses the FakeQuantWithMinMaxArgs.num_bits attribute to adjust quantized
	// array data types throughout the graph. The graph must be properly annotated
	// with FakeQuant* ops on at least the edges and may contain additional ops on
	// the interior of the graph to widen/narrow as desired.
	//
	// Input and output array data types may change because of this propagation
	// and users must be sure to query the final data_type values.
	PropagateFakeQuantNumBits *bool `protobuf:"varint,14,opt,name=propagate_fake_quant_num_bits,json=propagateFakeQuantNumBits" json:"propagate_fake_quant_num_bits,omitempty"`
	// Some fast uint8 GEMM kernels require uint8 weights to avoid the value 0.
	// This flag allows nudging them to 1 to allow proceeding, with moderate
	// inaccuracy.
	AllowNudgingWeightsToUseFastGemmKernel *bool `protobuf:"varint,17,opt,name=allow_nudging_weights_to_use_fast_gemm_kernel,json=allowNudgingWeightsToUseFastGemmKernel" json:"allow_nudging_weights_to_use_fast_gemm_kernel,omitempty"`
	// Minimum size of constant arrays to deduplicate; arrays smaller will not be
	// deduplicated.
	DedupeArrayMinSizeBytes *int64 `protobuf:"varint,18,opt,name=dedupe_array_min_size_bytes,json=dedupeArrayMinSizeBytes,def=64" json:"dedupe_array_min_size_bytes,omitempty"`
	// Split the LSTM inputs from 5 tensors to 18 tensors for TFLite.
	// Ignored if the output format is not TFLite.
	SplitTfliteLstmInputs *bool `protobuf:"varint,19,opt,name=split_tflite_lstm_inputs,json=splitTfliteLstmInputs,def=1" json:"split_tflite_lstm_inputs,omitempty"`
	// Store weights as quantized weights followed by dequantize operations.
	// Computation is still done in float, but reduces model size (at the cost of
	// accuracy and latency).
	// DEPRECATED: Please use post_training_quantize instead.
	QuantizeWeights *bool `protobuf:"varint,20,opt,name=quantize_weights,json=quantizeWeights,def=0" json:"quantize_weights,omitempty"`
	// Full filepath of folder to dump the graphs at various stages of processing
	// GraphViz .dot files. Preferred over --output_format=GRAPHVIZ_DOT in order
	// to keep the requirements of the output file.
	DumpGraphvizDir *string `protobuf:"bytes,24,opt,name=dump_graphviz_dir,json=dumpGraphvizDir" json:"dump_graphviz_dir,omitempty"`
	// Boolean indicating whether to dump the graph after every graph
	// transformation.
	DumpGraphvizIncludeVideo *bool `protobuf:"varint,25,opt,name=dump_graphviz_include_video,json=dumpGraphvizIncludeVideo" json:"dump_graphviz_include_video,omitempty"`
	// Boolean indicating whether to quantize the weights of the converted float
	// model. Model size will be reduced and there will be latency improvements
	// (at the cost of accuracy).
	PostTrainingQuantize *bool `protobuf:"varint,26,opt,name=post_training_quantize,json=postTrainingQuantize,def=0" json:"post_training_quantize,omitempty"`
	// This flag only works when converting to TensorFlow Lite format.
	// When enabled, unsupported ops will be converted to select TensorFlow ops.
	// TODO(ycling): Consider to rename the following 2 flags and don't call it
	// "Flex".
	// `enable_select_tf_ops` should always be used with `allow_custom_ops`.
	// WARNING: Experimental interface, subject to change
	EnableSelectTfOps *bool `protobuf:"varint,27,opt,name=enable_select_tf_ops,json=enableSelectTfOps,def=0" json:"enable_select_tf_ops,omitempty"`
	// This flag only works when converting to TensorFlow Lite format.
	// When enabled, all TensorFlow ops will be converted to select TensorFlow
	// ops.
	// This will force `enable_select_tf_ops` to true.
	// `force_select_tf_ops` should always be used with `enable_select_tf_ops`.
	// WARNING: Experimental interface, subject to change
	ForceSelectTfOps *bool `protobuf:"varint,28,opt,name=force_select_tf_ops,json=forceSelectTfOps,def=0" json:"force_select_tf_ops,omitempty"`
	// Boolean indicating whether to convert float32 constant buffers to
	// float16. This is typically done to reduce model size. Delegates may also
	// wish to implement kernels on reduced precision floats for performance
	// gains.
	QuantizeToFloat16 *bool `protobuf:"varint,29,opt,name=quantize_to_float16,json=quantizeToFloat16,def=0" json:"quantize_to_float16,omitempty"`
	// Boolean flag indicating whether the converter should allow models with
	// dynamic Tensor shape. When set to False, the converter will generate
	// runtime memory offsets for activation Tensors (with 128 bits alignment)
	// and error out on models with undetermined Tensor shape. (Default: True)
	AllowDynamicTensors  *bool    `protobuf:"varint,30,opt,name=allow_dynamic_tensors,json=allowDynamicTensors,def=1" json:"allow_dynamic_tensors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TocoFlags) Reset()         { *m = TocoFlags{} }
func (m *TocoFlags) String() string { return proto.CompactTextString(m) }
func (*TocoFlags) ProtoMessage()    {}
func (*TocoFlags) Descriptor() ([]byte, []int) {
	return fileDescriptor_655e1383a1a98309, []int{0}
}

func (m *TocoFlags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TocoFlags.Unmarshal(m, b)
}
func (m *TocoFlags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TocoFlags.Marshal(b, m, deterministic)
}
func (m *TocoFlags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TocoFlags.Merge(m, src)
}
func (m *TocoFlags) XXX_Size() int {
	return xxx_messageInfo_TocoFlags.Size(m)
}
func (m *TocoFlags) XXX_DiscardUnknown() {
	xxx_messageInfo_TocoFlags.DiscardUnknown(m)
}

var xxx_messageInfo_TocoFlags proto.InternalMessageInfo

const Default_TocoFlags_DedupeArrayMinSizeBytes int64 = 64
const Default_TocoFlags_SplitTfliteLstmInputs bool = true
const Default_TocoFlags_QuantizeWeights bool = false
const Default_TocoFlags_PostTrainingQuantize bool = false
const Default_TocoFlags_EnableSelectTfOps bool = false
const Default_TocoFlags_ForceSelectTfOps bool = false
const Default_TocoFlags_QuantizeToFloat16 bool = false
const Default_TocoFlags_AllowDynamicTensors bool = true

func (m *TocoFlags) GetInputFormat() FileFormat {
	if m != nil && m.InputFormat != nil {
		return *m.InputFormat
	}
	return FileFormat_FILE_FORMAT_UNKNOWN
}

func (m *TocoFlags) GetOutputFormat() FileFormat {
	if m != nil && m.OutputFormat != nil {
		return *m.OutputFormat
	}
	return FileFormat_FILE_FORMAT_UNKNOWN
}

func (m *TocoFlags) GetInferenceInputType() IODataType {
	if m != nil && m.InferenceInputType != nil {
		return *m.InferenceInputType
	}
	return IODataType_IO_DATA_TYPE_UNKNOWN
}

func (m *TocoFlags) GetInferenceType() IODataType {
	if m != nil && m.InferenceType != nil {
		return *m.InferenceType
	}
	return IODataType_IO_DATA_TYPE_UNKNOWN
}

func (m *TocoFlags) GetDefaultRangesMin() float32 {
	if m != nil && m.DefaultRangesMin != nil {
		return *m.DefaultRangesMin
	}
	return 0
}

func (m *TocoFlags) GetDefaultRangesMax() float32 {
	if m != nil && m.DefaultRangesMax != nil {
		return *m.DefaultRangesMax
	}
	return 0
}

func (m *TocoFlags) GetDefaultInt16RangesMin() float32 {
	if m != nil && m.DefaultInt16RangesMin != nil {
		return *m.DefaultInt16RangesMin
	}
	return 0
}

func (m *TocoFlags) GetDefaultInt16RangesMax() float32 {
	if m != nil && m.DefaultInt16RangesMax != nil {
		return *m.DefaultInt16RangesMax
	}
	return 0
}

func (m *TocoFlags) GetDropFakeQuant() bool {
	if m != nil && m.DropFakeQuant != nil {
		return *m.DropFakeQuant
	}
	return false
}

func (m *TocoFlags) GetReorderAcrossFakeQuant() bool {
	if m != nil && m.ReorderAcrossFakeQuant != nil {
		return *m.ReorderAcrossFakeQuant
	}
	return false
}

func (m *TocoFlags) GetAllowCustomOps() bool {
	if m != nil && m.AllowCustomOps != nil {
		return *m.AllowCustomOps
	}
	return false
}

func (m *TocoFlags) GetDropControlDependency() bool {
	if m != nil && m.DropControlDependency != nil {
		return *m.DropControlDependency
	}
	return false
}

func (m *TocoFlags) GetDebugDisableRecurrentCellFusion() bool {
	if m != nil && m.DebugDisableRecurrentCellFusion != nil {
		return *m.DebugDisableRecurrentCellFusion
	}
	return false
}

func (m *TocoFlags) GetPropagateFakeQuantNumBits() bool {
	if m != nil && m.PropagateFakeQuantNumBits != nil {
		return *m.PropagateFakeQuantNumBits
	}
	return false
}

func (m *TocoFlags) GetAllowNudgingWeightsToUseFastGemmKernel() bool {
	if m != nil && m.AllowNudgingWeightsToUseFastGemmKernel != nil {
		return *m.AllowNudgingWeightsToUseFastGemmKernel
	}
	return false
}

func (m *TocoFlags) GetDedupeArrayMinSizeBytes() int64 {
	if m != nil && m.DedupeArrayMinSizeBytes != nil {
		return *m.DedupeArrayMinSizeBytes
	}
	return Default_TocoFlags_DedupeArrayMinSizeBytes
}

func (m *TocoFlags) GetSplitTfliteLstmInputs() bool {
	if m != nil && m.SplitTfliteLstmInputs != nil {
		return *m.SplitTfliteLstmInputs
	}
	return Default_TocoFlags_SplitTfliteLstmInputs
}

func (m *TocoFlags) GetQuantizeWeights() bool {
	if m != nil && m.QuantizeWeights != nil {
		return *m.QuantizeWeights
	}
	return Default_TocoFlags_QuantizeWeights
}

func (m *TocoFlags) GetDumpGraphvizDir() string {
	if m != nil && m.DumpGraphvizDir != nil {
		return *m.DumpGraphvizDir
	}
	return ""
}

func (m *TocoFlags) GetDumpGraphvizIncludeVideo() bool {
	if m != nil && m.DumpGraphvizIncludeVideo != nil {
		return *m.DumpGraphvizIncludeVideo
	}
	return false
}

func (m *TocoFlags) GetPostTrainingQuantize() bool {
	if m != nil && m.PostTrainingQuantize != nil {
		return *m.PostTrainingQuantize
	}
	return Default_TocoFlags_PostTrainingQuantize
}

func (m *TocoFlags) GetEnableSelectTfOps() bool {
	if m != nil && m.EnableSelectTfOps != nil {
		return *m.EnableSelectTfOps
	}
	return Default_TocoFlags_EnableSelectTfOps
}

func (m *TocoFlags) GetForceSelectTfOps() bool {
	if m != nil && m.ForceSelectTfOps != nil {
		return *m.ForceSelectTfOps
	}
	return Default_TocoFlags_ForceSelectTfOps
}

func (m *TocoFlags) GetQuantizeToFloat16() bool {
	if m != nil && m.QuantizeToFloat16 != nil {
		return *m.QuantizeToFloat16
	}
	return Default_TocoFlags_QuantizeToFloat16
}

func (m *TocoFlags) GetAllowDynamicTensors() bool {
	if m != nil && m.AllowDynamicTensors != nil {
		return *m.AllowDynamicTensors
	}
	return Default_TocoFlags_AllowDynamicTensors
}

func init() {
	proto.RegisterEnum("toco.FileFormat", FileFormat_name, FileFormat_value)
	proto.RegisterType((*TocoFlags)(nil), "toco.TocoFlags")
}

func init() {
	proto.RegisterFile("tensorflow/lite/toco/toco_flags.proto", fileDescriptor_655e1383a1a98309)
}

var fileDescriptor_655e1383a1a98309 = []byte{
	// 893 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x95, 0x7f, 0x4f, 0xdb, 0x46,
	0x18, 0xc7, 0x17, 0x4a, 0xbb, 0xf6, 0xf8, 0x65, 0x0e, 0x28, 0x47, 0x59, 0xb7, 0x68, 0xd3, 0xaa,
	0xa8, 0xda, 0x60, 0xdd, 0x5a, 0xba, 0x75, 0x42, 0x2b, 0x10, 0xcc, 0xa2, 0x42, 0xb2, 0x1a, 0xb7,
	0x48, 0xd5, 0xa6, 0xd3, 0xc5, 0x7e, 0x6c, 0x4e, 0xd8, 0x77, 0xde, 0xdd, 0xb9, 0x10, 0x5e, 0xd5,
	0x5e, 0xe2, 0x74, 0x67, 0x27, 0x71, 0x2b, 0xf8, 0x27, 0x4a, 0x9e, 0xef, 0xe7, 0x73, 0xbf, 0xf2,
	0xf8, 0x8c, 0xbe, 0x37, 0x20, 0xb4, 0x54, 0x49, 0x26, 0x2f, 0xb7, 0x33, 0x6e, 0x60, 0xdb, 0xc8,
	0x48, 0xba, 0x0f, 0x9a, 0x64, 0x2c, 0xd5, 0x5b, 0x85, 0x92, 0x46, 0xe2, 0x59, 0x5b, 0x79, 0xd4,
	0xbe, 0x19, 0x1e, 0x15, 0x50, 0x73, 0xdf, 0xfe, 0x37, 0x87, 0x1e, 0x84, 0x32, 0x92, 0xbe, 0x75,
	0xf1, 0x2f, 0x68, 0x9e, 0x8b, 0xa2, 0x34, 0x34, 0x91, 0x2a, 0x67, 0x86, 0xb4, 0xda, 0xad, 0xce,
	0xe2, 0xcf, 0xde, 0x96, 0xd5, 0xb6, 0x7c, 0x9e, 0x81, 0xef, 0xea, 0xc1, 0x9c, 0xa3, 0xaa, 0x1f,
	0xf8, 0x05, 0x5a, 0x90, 0xa5, 0x69, 0x58, 0x33, 0xb7, 0x58, 0xf3, 0x15, 0x56, 0x6b, 0xfb, 0x68,
	0x95, 0x8b, 0x04, 0x14, 0x88, 0x08, 0x68, 0x35, 0xab, 0x5d, 0x18, 0x99, 0x6b, 0xda, 0xbd, 0x41,
	0x97, 0x19, 0x16, 0x8e, 0x0a, 0x08, 0xf0, 0x84, 0xee, 0x59, 0xd8, 0xd6, 0xf0, 0x4b, 0xb4, 0x38,
	0x1d, 0xc3, 0xd9, 0xb3, 0xb7, 0xd8, 0x0b, 0x13, 0xce, 0x89, 0x3f, 0x20, 0x1c, 0x43, 0xc2, 0xca,
	0xcc, 0x50, 0xc5, 0x44, 0x0a, 0x9a, 0xe6, 0x5c, 0x90, 0xbb, 0xed, 0x56, 0x67, 0x26, 0xf0, 0xea,
	0x24, 0x70, 0xc1, 0x09, 0x17, 0x37, 0xd1, 0xec, 0x8a, 0xdc, 0xbb, 0x89, 0x66, 0x57, 0xf8, 0x25,
	0x22, 0x63, 0x9a, 0x0b, 0xf3, 0x6c, 0xa7, 0x39, 0xc3, 0x92, 0x73, 0xd6, 0xea, 0xbc, 0x67, 0xe3,
	0xe9, 0x34, 0xb7, 0x8a, 0xec, 0x8a, 0x78, 0xb7, 0x8a, 0xec, 0x0a, 0x3f, 0x41, 0x4b, 0xb1, 0x92,
	0x05, 0x4d, 0xd8, 0x05, 0xd0, 0x7f, 0x4b, 0x26, 0x0c, 0xf9, 0xb2, 0xdd, 0xea, 0xdc, 0x0f, 0x16,
	0x6c, 0xd9, 0x67, 0x17, 0xf0, 0xd6, 0x16, 0xf1, 0x6f, 0x68, 0x43, 0x81, 0x54, 0x31, 0x28, 0xca,
	0x22, 0x25, 0xb5, 0x6e, 0x1a, 0xf7, 0x9d, 0xf1, 0xb0, 0x06, 0xf6, 0x5c, 0x3e, 0x55, 0x3b, 0xc8,
	0x63, 0x59, 0x26, 0x2f, 0x69, 0x54, 0x6a, 0x23, 0x73, 0x2a, 0x0b, 0x4d, 0x90, 0x33, 0x16, 0x5d,
	0xfd, 0xc0, 0x95, 0x07, 0x85, 0xc6, 0x3b, 0x68, 0xdd, 0x2d, 0x26, 0x92, 0xc2, 0x28, 0x99, 0xd1,
	0x18, 0x0a, 0x10, 0x31, 0x88, 0x68, 0x44, 0xe6, 0x9d, 0xb0, 0x66, 0xe3, 0x83, 0x2a, 0xed, 0x4e,
	0x42, 0x7c, 0x8c, 0xbe, 0x8b, 0x61, 0x58, 0xa6, 0x34, 0xe6, 0x9a, 0x0d, 0x33, 0xa0, 0x0a, 0xa2,
	0x52, 0x29, 0x10, 0x86, 0x46, 0x90, 0x65, 0x34, 0x29, 0x35, 0x97, 0x82, 0x2c, 0xb8, 0x31, 0xbe,
	0x71, 0x68, 0xb7, 0x22, 0x83, 0x31, 0x78, 0x00, 0x59, 0xe6, 0x3b, 0x0c, 0xbf, 0x46, 0x8f, 0x0b,
	0x25, 0x0b, 0x96, 0x32, 0x03, 0x8d, 0x5d, 0x52, 0x51, 0xe6, 0x74, 0xc8, 0x8d, 0x26, 0x8b, 0x6e,
	0x9c, 0x8d, 0x09, 0x34, 0xd9, 0x6a, 0xbf, 0xcc, 0xf7, 0xb9, 0xd1, 0xf8, 0x1f, 0xf4, 0x63, 0xb5,
	0x63, 0x51, 0xc6, 0x29, 0x17, 0x29, 0xbd, 0x04, 0x9e, 0x9e, 0x1b, 0x4d, 0x8d, 0xa4, 0xa5, 0xb6,
	0x83, 0x6a, 0x43, 0x53, 0xc8, 0x73, 0x7a, 0x01, 0x4a, 0x40, 0x46, 0x96, 0xdd, 0x88, 0x4f, 0x9c,
	0xd4, 0xaf, 0x9c, 0xb3, 0x4a, 0x09, 0xe5, 0x3b, 0x0d, 0x3e, 0xd3, 0xe6, 0x08, 0xf2, 0xfc, 0x8d,
	0xa3, 0xf1, 0x6b, 0xb4, 0x19, 0x43, 0x5c, 0x16, 0x40, 0x99, 0x52, 0x6c, 0x64, 0xbb, 0x83, 0x6a,
	0x7e, 0x0d, 0x74, 0x38, 0x32, 0xa0, 0x09, 0x6e, 0xb7, 0x3a, 0x77, 0x5e, 0xcd, 0xec, 0x3c, 0x0f,
	0xd6, 0x2b, 0x6c, 0xcf, 0x52, 0x27, 0x5c, 0x9c, 0xf2, 0x6b, 0xd8, 0xb7, 0x08, 0xde, 0x45, 0x44,
	0x17, 0x19, 0x37, 0xd4, 0x24, 0xf6, 0xd9, 0xa6, 0x99, 0x36, 0x79, 0xf5, 0x20, 0x69, 0xb2, 0x62,
	0xd7, 0xf2, 0x6a, 0xd6, 0xa8, 0x12, 0x82, 0x35, 0x47, 0x85, 0x0e, 0x3a, 0xd6, 0x26, 0x77, 0x8f,
	0x8f, 0xc6, 0x3f, 0x21, 0xcf, 0x1d, 0x89, 0x9d, 0xb3, 0xde, 0x1a, 0x59, 0x75, 0xda, 0xdd, 0x84,
	0x65, 0x1a, 0x82, 0xa5, 0x71, 0x5c, 0xef, 0x02, 0x3f, 0x45, 0xcb, 0x71, 0x99, 0x17, 0x34, 0x55,
	0xac, 0x38, 0xff, 0xc8, 0xaf, 0x69, 0xcc, 0x15, 0x21, 0xed, 0x56, 0xe7, 0x41, 0xb0, 0x64, 0x83,
	0xa3, 0xba, 0xde, 0xe5, 0x0a, 0xef, 0xa2, 0xcd, 0x4f, 0x59, 0x2e, 0xa2, 0xac, 0x8c, 0x81, 0x7e,
	0xe4, 0x31, 0x48, 0xb2, 0xe1, 0xce, 0x8a, 0x34, 0xad, 0x5e, 0x05, 0xbc, 0xb7, 0x39, 0xfe, 0x1d,
	0x3d, 0x2c, 0xa4, 0x36, 0xd4, 0x28, 0xc6, 0x85, 0x3d, 0xfc, 0xf1, 0x5a, 0xc8, 0xa3, 0xe6, 0x12,
	0x57, 0x2d, 0x14, 0xd6, 0xcc, 0xdb, 0x1a, 0xc1, 0x3b, 0x68, 0x15, 0x84, 0x6b, 0x21, 0x0d, 0x19,
	0x44, 0xf6, 0x80, 0x5c, 0xbf, 0x6e, 0x36, 0xd5, 0xe5, 0x0a, 0x39, 0x75, 0x44, 0x98, 0xd8, 0xce,
	0x7d, 0x8e, 0x56, 0x12, 0xa9, 0xa2, 0xcf, 0xb5, 0xaf, 0x9a, 0x9a, 0xe7, 0x88, 0xa6, 0xf5, 0x02,
	0xad, 0x4c, 0xce, 0xd1, 0xd8, 0x4b, 0x58, 0x32, 0xf3, 0x6c, 0x87, 0x3c, 0xfe, 0x64, 0xb2, 0x31,
	0x11, 0x4a, 0xbf, 0xca, 0xf1, 0xaf, 0x68, 0xad, 0x6a, 0xaf, 0x78, 0x24, 0x58, 0xce, 0x23, 0x5a,
	0x5d, 0xd5, 0x9a, 0x7c, 0xdd, 0xf8, 0xeb, 0x56, 0x1c, 0xd2, 0xad, 0x88, 0xb0, 0x02, 0x9e, 0xfe,
	0x8d, 0xd0, 0xf4, 0x52, 0xc5, 0xeb, 0x68, 0xc5, 0xef, 0x1d, 0x1f, 0x52, 0x7f, 0x10, 0x9c, 0xec,
	0x85, 0xf4, 0x5d, 0xff, 0x4d, 0x7f, 0x70, 0xd6, 0xf7, 0xbe, 0xb0, 0x41, 0x78, 0xd8, 0x3f, 0x1d,
	0x04, 0xfe, 0xf1, 0xe0, 0x8c, 0x1e, 0x05, 0x7b, 0x7f, 0xfd, 0xd9, 0x3d, 0xf4, 0xbd, 0x16, 0x46,
	0xe8, 0x5e, 0xe8, 0x1f, 0xf7, 0xc2, 0x43, 0x6f, 0x06, 0x7b, 0x68, 0xde, 0x25, 0xef, 0x7b, 0x1f,
	0x68, 0x77, 0x10, 0x7a, 0x77, 0xf6, 0xff, 0xf8, 0xb0, 0x9b, 0x72, 0x73, 0x5e, 0x0e, 0xb7, 0x22,
	0x99, 0x6f, 0x37, 0xde, 0x1f, 0x37, 0x7f, 0x4d, 0xe5, 0x76, 0x24, 0x15, 0x4c, 0xdf, 0x2e, 0xff,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x3b, 0x7e, 0x5f, 0xa1, 0x06, 0x00, 0x00,
}
