// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/profiler/tfprof_output.proto

package profiler

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

type TFProfTensorProto struct {
	Dtype framework.DataType `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	// Flatten tensor in row-major.
	// Only one of the following array is set.
	ValueDouble          []float64 `protobuf:"fixed64,2,rep,packed,name=value_double,json=valueDouble,proto3" json:"value_double,omitempty"`
	ValueInt64           []int64   `protobuf:"varint,3,rep,packed,name=value_int64,json=valueInt64,proto3" json:"value_int64,omitempty"`
	ValueStr             []string  `protobuf:"bytes,4,rep,name=value_str,json=valueStr,proto3" json:"value_str,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TFProfTensorProto) Reset()         { *m = TFProfTensorProto{} }
func (m *TFProfTensorProto) String() string { return proto.CompactTextString(m) }
func (*TFProfTensorProto) ProtoMessage()    {}
func (*TFProfTensorProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d317c37c959b8d2, []int{0}
}

func (m *TFProfTensorProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFProfTensorProto.Unmarshal(m, b)
}
func (m *TFProfTensorProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFProfTensorProto.Marshal(b, m, deterministic)
}
func (m *TFProfTensorProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFProfTensorProto.Merge(m, src)
}
func (m *TFProfTensorProto) XXX_Size() int {
	return xxx_messageInfo_TFProfTensorProto.Size(m)
}
func (m *TFProfTensorProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TFProfTensorProto.DiscardUnknown(m)
}

var xxx_messageInfo_TFProfTensorProto proto.InternalMessageInfo

func (m *TFProfTensorProto) GetDtype() framework.DataType {
	if m != nil {
		return m.Dtype
	}
	return framework.DataType_DT_INVALID
}

func (m *TFProfTensorProto) GetValueDouble() []float64 {
	if m != nil {
		return m.ValueDouble
	}
	return nil
}

func (m *TFProfTensorProto) GetValueInt64() []int64 {
	if m != nil {
		return m.ValueInt64
	}
	return nil
}

func (m *TFProfTensorProto) GetValueStr() []string {
	if m != nil {
		return m.ValueStr
	}
	return nil
}

// A node in TensorFlow graph. Used by scope/graph view.
type GraphNodeProto struct {
	// op name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// tensor value restored from checkpoint.
	TensorValue *TFProfTensorProto `protobuf:"bytes,15,opt,name=tensor_value,json=tensorValue,proto3" json:"tensor_value,omitempty"`
	// op execution time.
	// A node can be defined once but run multiple times in tf.while_loop.
	// the times sum up all different runs.
	RunCount              int64 `protobuf:"varint,21,opt,name=run_count,json=runCount,proto3" json:"run_count,omitempty"`
	ExecMicros            int64 `protobuf:"varint,2,opt,name=exec_micros,json=execMicros,proto3" json:"exec_micros,omitempty"`
	AcceleratorExecMicros int64 `protobuf:"varint,17,opt,name=accelerator_exec_micros,json=acceleratorExecMicros,proto3" json:"accelerator_exec_micros,omitempty"`
	CpuExecMicros         int64 `protobuf:"varint,18,opt,name=cpu_exec_micros,json=cpuExecMicros,proto3" json:"cpu_exec_micros,omitempty"`
	// Total bytes requested by the op.
	RequestedBytes int64 `protobuf:"varint,3,opt,name=requested_bytes,json=requestedBytes,proto3" json:"requested_bytes,omitempty"`
	// Max bytes allocated and being used by the op at a point.
	PeakBytes int64 `protobuf:"varint,24,opt,name=peak_bytes,json=peakBytes,proto3" json:"peak_bytes,omitempty"`
	// Total bytes requested by the op and not released before end.
	ResidualBytes int64 `protobuf:"varint,25,opt,name=residual_bytes,json=residualBytes,proto3" json:"residual_bytes,omitempty"`
	// Total bytes output by the op (not necessarily allocated by the op).
	OutputBytes int64 `protobuf:"varint,26,opt,name=output_bytes,json=outputBytes,proto3" json:"output_bytes,omitempty"`
	// Number of parameters if available.
	Parameters int64 `protobuf:"varint,4,opt,name=parameters,proto3" json:"parameters,omitempty"`
	// Number of float operations.
	FloatOps int64 `protobuf:"varint,13,opt,name=float_ops,json=floatOps,proto3" json:"float_ops,omitempty"`
	// Device the op is assigned to.
	// Since an op can fire multiple kernel calls, there can be multiple devices.
	Devices []string `protobuf:"bytes,10,rep,name=devices,proto3" json:"devices,omitempty"`
	// The following are the aggregated stats from all *accounted* children and
	// the node itself. The actual children depend on the data structure used.
	// In graph view, children are inputs recursively.
	// In scope view, children are nodes under the name scope.
	TotalDefinitionCount       int64 `protobuf:"varint,23,opt,name=total_definition_count,json=totalDefinitionCount,proto3" json:"total_definition_count,omitempty"`
	TotalRunCount              int64 `protobuf:"varint,22,opt,name=total_run_count,json=totalRunCount,proto3" json:"total_run_count,omitempty"`
	TotalExecMicros            int64 `protobuf:"varint,6,opt,name=total_exec_micros,json=totalExecMicros,proto3" json:"total_exec_micros,omitempty"`
	TotalAcceleratorExecMicros int64 `protobuf:"varint,19,opt,name=total_accelerator_exec_micros,json=totalAcceleratorExecMicros,proto3" json:"total_accelerator_exec_micros,omitempty"`
	TotalCpuExecMicros         int64 `protobuf:"varint,20,opt,name=total_cpu_exec_micros,json=totalCpuExecMicros,proto3" json:"total_cpu_exec_micros,omitempty"`
	TotalRequestedBytes        int64 `protobuf:"varint,7,opt,name=total_requested_bytes,json=totalRequestedBytes,proto3" json:"total_requested_bytes,omitempty"`
	TotalPeakBytes             int64 `protobuf:"varint,27,opt,name=total_peak_bytes,json=totalPeakBytes,proto3" json:"total_peak_bytes,omitempty"`
	TotalResidualBytes         int64 `protobuf:"varint,28,opt,name=total_residual_bytes,json=totalResidualBytes,proto3" json:"total_residual_bytes,omitempty"`
	TotalOutputBytes           int64 `protobuf:"varint,29,opt,name=total_output_bytes,json=totalOutputBytes,proto3" json:"total_output_bytes,omitempty"`
	TotalParameters            int64 `protobuf:"varint,8,opt,name=total_parameters,json=totalParameters,proto3" json:"total_parameters,omitempty"`
	TotalFloatOps              int64 `protobuf:"varint,14,opt,name=total_float_ops,json=totalFloatOps,proto3" json:"total_float_ops,omitempty"`
	// shape information, if available.
	// TODO(xpan): Why is this repeated?
	Shapes      []*framework.TensorShapeProto         `protobuf:"bytes,11,rep,name=shapes,proto3" json:"shapes,omitempty"`
	InputShapes map[int32]*framework.TensorShapeProto `protobuf:"bytes,16,rep,name=input_shapes,json=inputShapes,proto3" json:"input_shapes,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Descendants of the graph. The actual descendants depend on the data
	// structure used (scope, graph).
	Children             []*GraphNodeProto `protobuf:"bytes,12,rep,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GraphNodeProto) Reset()         { *m = GraphNodeProto{} }
func (m *GraphNodeProto) String() string { return proto.CompactTextString(m) }
func (*GraphNodeProto) ProtoMessage()    {}
func (*GraphNodeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d317c37c959b8d2, []int{1}
}

func (m *GraphNodeProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphNodeProto.Unmarshal(m, b)
}
func (m *GraphNodeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphNodeProto.Marshal(b, m, deterministic)
}
func (m *GraphNodeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphNodeProto.Merge(m, src)
}
func (m *GraphNodeProto) XXX_Size() int {
	return xxx_messageInfo_GraphNodeProto.Size(m)
}
func (m *GraphNodeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphNodeProto.DiscardUnknown(m)
}

var xxx_messageInfo_GraphNodeProto proto.InternalMessageInfo

func (m *GraphNodeProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GraphNodeProto) GetTensorValue() *TFProfTensorProto {
	if m != nil {
		return m.TensorValue
	}
	return nil
}

func (m *GraphNodeProto) GetRunCount() int64 {
	if m != nil {
		return m.RunCount
	}
	return 0
}

func (m *GraphNodeProto) GetExecMicros() int64 {
	if m != nil {
		return m.ExecMicros
	}
	return 0
}

func (m *GraphNodeProto) GetAcceleratorExecMicros() int64 {
	if m != nil {
		return m.AcceleratorExecMicros
	}
	return 0
}

func (m *GraphNodeProto) GetCpuExecMicros() int64 {
	if m != nil {
		return m.CpuExecMicros
	}
	return 0
}

func (m *GraphNodeProto) GetRequestedBytes() int64 {
	if m != nil {
		return m.RequestedBytes
	}
	return 0
}

func (m *GraphNodeProto) GetPeakBytes() int64 {
	if m != nil {
		return m.PeakBytes
	}
	return 0
}

func (m *GraphNodeProto) GetResidualBytes() int64 {
	if m != nil {
		return m.ResidualBytes
	}
	return 0
}

func (m *GraphNodeProto) GetOutputBytes() int64 {
	if m != nil {
		return m.OutputBytes
	}
	return 0
}

func (m *GraphNodeProto) GetParameters() int64 {
	if m != nil {
		return m.Parameters
	}
	return 0
}

func (m *GraphNodeProto) GetFloatOps() int64 {
	if m != nil {
		return m.FloatOps
	}
	return 0
}

func (m *GraphNodeProto) GetDevices() []string {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *GraphNodeProto) GetTotalDefinitionCount() int64 {
	if m != nil {
		return m.TotalDefinitionCount
	}
	return 0
}

func (m *GraphNodeProto) GetTotalRunCount() int64 {
	if m != nil {
		return m.TotalRunCount
	}
	return 0
}

func (m *GraphNodeProto) GetTotalExecMicros() int64 {
	if m != nil {
		return m.TotalExecMicros
	}
	return 0
}

func (m *GraphNodeProto) GetTotalAcceleratorExecMicros() int64 {
	if m != nil {
		return m.TotalAcceleratorExecMicros
	}
	return 0
}

func (m *GraphNodeProto) GetTotalCpuExecMicros() int64 {
	if m != nil {
		return m.TotalCpuExecMicros
	}
	return 0
}

func (m *GraphNodeProto) GetTotalRequestedBytes() int64 {
	if m != nil {
		return m.TotalRequestedBytes
	}
	return 0
}

func (m *GraphNodeProto) GetTotalPeakBytes() int64 {
	if m != nil {
		return m.TotalPeakBytes
	}
	return 0
}

func (m *GraphNodeProto) GetTotalResidualBytes() int64 {
	if m != nil {
		return m.TotalResidualBytes
	}
	return 0
}

func (m *GraphNodeProto) GetTotalOutputBytes() int64 {
	if m != nil {
		return m.TotalOutputBytes
	}
	return 0
}

func (m *GraphNodeProto) GetTotalParameters() int64 {
	if m != nil {
		return m.TotalParameters
	}
	return 0
}

func (m *GraphNodeProto) GetTotalFloatOps() int64 {
	if m != nil {
		return m.TotalFloatOps
	}
	return 0
}

func (m *GraphNodeProto) GetShapes() []*framework.TensorShapeProto {
	if m != nil {
		return m.Shapes
	}
	return nil
}

func (m *GraphNodeProto) GetInputShapes() map[int32]*framework.TensorShapeProto {
	if m != nil {
		return m.InputShapes
	}
	return nil
}

func (m *GraphNodeProto) GetChildren() []*GraphNodeProto {
	if m != nil {
		return m.Children
	}
	return nil
}

// A node that groups multiple GraphNodeProto.
// Depending on the 'view', the semantics of the TFmultiGraphNodeProto
// is different:
// code view: A node groups all TensorFlow graph nodes created by the
//            Python code.
// op view:   A node groups all TensorFlow graph nodes that are of type
//            of the op (e.g. MatMul, Conv2D).
type MultiGraphNodeProto struct {
	// Name of the node.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// code execution time.
	ExecMicros            int64 `protobuf:"varint,2,opt,name=exec_micros,json=execMicros,proto3" json:"exec_micros,omitempty"`
	AcceleratorExecMicros int64 `protobuf:"varint,12,opt,name=accelerator_exec_micros,json=acceleratorExecMicros,proto3" json:"accelerator_exec_micros,omitempty"`
	CpuExecMicros         int64 `protobuf:"varint,13,opt,name=cpu_exec_micros,json=cpuExecMicros,proto3" json:"cpu_exec_micros,omitempty"`
	// Total requested bytes by the code.
	RequestedBytes int64 `protobuf:"varint,3,opt,name=requested_bytes,json=requestedBytes,proto3" json:"requested_bytes,omitempty"`
	// Max bytes allocated and being used by the op at a point.
	PeakBytes int64 `protobuf:"varint,16,opt,name=peak_bytes,json=peakBytes,proto3" json:"peak_bytes,omitempty"`
	// Total bytes requested by the op and not released before end.
	ResidualBytes int64 `protobuf:"varint,17,opt,name=residual_bytes,json=residualBytes,proto3" json:"residual_bytes,omitempty"`
	// Total bytes output by the op (not necessarily allocated by the op).
	OutputBytes int64 `protobuf:"varint,18,opt,name=output_bytes,json=outputBytes,proto3" json:"output_bytes,omitempty"`
	// Number of parameters if available.
	Parameters int64 `protobuf:"varint,4,opt,name=parameters,proto3" json:"parameters,omitempty"`
	// Number of float operations.
	FloatOps int64 `protobuf:"varint,5,opt,name=float_ops,json=floatOps,proto3" json:"float_ops,omitempty"`
	// The following are the aggregated stats from descendants.
	// The actual descendants depend on the data structure used.
	TotalExecMicros            int64 `protobuf:"varint,6,opt,name=total_exec_micros,json=totalExecMicros,proto3" json:"total_exec_micros,omitempty"`
	TotalAcceleratorExecMicros int64 `protobuf:"varint,14,opt,name=total_accelerator_exec_micros,json=totalAcceleratorExecMicros,proto3" json:"total_accelerator_exec_micros,omitempty"`
	TotalCpuExecMicros         int64 `protobuf:"varint,15,opt,name=total_cpu_exec_micros,json=totalCpuExecMicros,proto3" json:"total_cpu_exec_micros,omitempty"`
	TotalRequestedBytes        int64 `protobuf:"varint,7,opt,name=total_requested_bytes,json=totalRequestedBytes,proto3" json:"total_requested_bytes,omitempty"`
	TotalPeakBytes             int64 `protobuf:"varint,19,opt,name=total_peak_bytes,json=totalPeakBytes,proto3" json:"total_peak_bytes,omitempty"`
	TotalResidualBytes         int64 `protobuf:"varint,20,opt,name=total_residual_bytes,json=totalResidualBytes,proto3" json:"total_residual_bytes,omitempty"`
	TotalOutputBytes           int64 `protobuf:"varint,21,opt,name=total_output_bytes,json=totalOutputBytes,proto3" json:"total_output_bytes,omitempty"`
	TotalParameters            int64 `protobuf:"varint,8,opt,name=total_parameters,json=totalParameters,proto3" json:"total_parameters,omitempty"`
	TotalFloatOps              int64 `protobuf:"varint,9,opt,name=total_float_ops,json=totalFloatOps,proto3" json:"total_float_ops,omitempty"`
	// TensorFlow graph nodes contained by the MultiGraphNodeProto.
	GraphNodes []*GraphNodeProto `protobuf:"bytes,10,rep,name=graph_nodes,json=graphNodes,proto3" json:"graph_nodes,omitempty"`
	// Descendants of the node. The actual descendants depend on the data
	// structure used.
	Children             []*MultiGraphNodeProto `protobuf:"bytes,11,rep,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MultiGraphNodeProto) Reset()         { *m = MultiGraphNodeProto{} }
func (m *MultiGraphNodeProto) String() string { return proto.CompactTextString(m) }
func (*MultiGraphNodeProto) ProtoMessage()    {}
func (*MultiGraphNodeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d317c37c959b8d2, []int{2}
}

func (m *MultiGraphNodeProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiGraphNodeProto.Unmarshal(m, b)
}
func (m *MultiGraphNodeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiGraphNodeProto.Marshal(b, m, deterministic)
}
func (m *MultiGraphNodeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiGraphNodeProto.Merge(m, src)
}
func (m *MultiGraphNodeProto) XXX_Size() int {
	return xxx_messageInfo_MultiGraphNodeProto.Size(m)
}
func (m *MultiGraphNodeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiGraphNodeProto.DiscardUnknown(m)
}

var xxx_messageInfo_MultiGraphNodeProto proto.InternalMessageInfo

func (m *MultiGraphNodeProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MultiGraphNodeProto) GetExecMicros() int64 {
	if m != nil {
		return m.ExecMicros
	}
	return 0
}

func (m *MultiGraphNodeProto) GetAcceleratorExecMicros() int64 {
	if m != nil {
		return m.AcceleratorExecMicros
	}
	return 0
}

func (m *MultiGraphNodeProto) GetCpuExecMicros() int64 {
	if m != nil {
		return m.CpuExecMicros
	}
	return 0
}

func (m *MultiGraphNodeProto) GetRequestedBytes() int64 {
	if m != nil {
		return m.RequestedBytes
	}
	return 0
}

func (m *MultiGraphNodeProto) GetPeakBytes() int64 {
	if m != nil {
		return m.PeakBytes
	}
	return 0
}

func (m *MultiGraphNodeProto) GetResidualBytes() int64 {
	if m != nil {
		return m.ResidualBytes
	}
	return 0
}

func (m *MultiGraphNodeProto) GetOutputBytes() int64 {
	if m != nil {
		return m.OutputBytes
	}
	return 0
}

func (m *MultiGraphNodeProto) GetParameters() int64 {
	if m != nil {
		return m.Parameters
	}
	return 0
}

func (m *MultiGraphNodeProto) GetFloatOps() int64 {
	if m != nil {
		return m.FloatOps
	}
	return 0
}

func (m *MultiGraphNodeProto) GetTotalExecMicros() int64 {
	if m != nil {
		return m.TotalExecMicros
	}
	return 0
}

func (m *MultiGraphNodeProto) GetTotalAcceleratorExecMicros() int64 {
	if m != nil {
		return m.TotalAcceleratorExecMicros
	}
	return 0
}

func (m *MultiGraphNodeProto) GetTotalCpuExecMicros() int64 {
	if m != nil {
		return m.TotalCpuExecMicros
	}
	return 0
}

func (m *MultiGraphNodeProto) GetTotalRequestedBytes() int64 {
	if m != nil {
		return m.TotalRequestedBytes
	}
	return 0
}

func (m *MultiGraphNodeProto) GetTotalPeakBytes() int64 {
	if m != nil {
		return m.TotalPeakBytes
	}
	return 0
}

func (m *MultiGraphNodeProto) GetTotalResidualBytes() int64 {
	if m != nil {
		return m.TotalResidualBytes
	}
	return 0
}

func (m *MultiGraphNodeProto) GetTotalOutputBytes() int64 {
	if m != nil {
		return m.TotalOutputBytes
	}
	return 0
}

func (m *MultiGraphNodeProto) GetTotalParameters() int64 {
	if m != nil {
		return m.TotalParameters
	}
	return 0
}

func (m *MultiGraphNodeProto) GetTotalFloatOps() int64 {
	if m != nil {
		return m.TotalFloatOps
	}
	return 0
}

func (m *MultiGraphNodeProto) GetGraphNodes() []*GraphNodeProto {
	if m != nil {
		return m.GraphNodes
	}
	return nil
}

func (m *MultiGraphNodeProto) GetChildren() []*MultiGraphNodeProto {
	if m != nil {
		return m.Children
	}
	return nil
}

type AdviceProto struct {
	// checker name -> a list of reports from the checker.
	Checkers             map[string]*AdviceProto_Checker `protobuf:"bytes,1,rep,name=checkers,proto3" json:"checkers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *AdviceProto) Reset()         { *m = AdviceProto{} }
func (m *AdviceProto) String() string { return proto.CompactTextString(m) }
func (*AdviceProto) ProtoMessage()    {}
func (*AdviceProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d317c37c959b8d2, []int{3}
}

func (m *AdviceProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdviceProto.Unmarshal(m, b)
}
func (m *AdviceProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdviceProto.Marshal(b, m, deterministic)
}
func (m *AdviceProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdviceProto.Merge(m, src)
}
func (m *AdviceProto) XXX_Size() int {
	return xxx_messageInfo_AdviceProto.Size(m)
}
func (m *AdviceProto) XXX_DiscardUnknown() {
	xxx_messageInfo_AdviceProto.DiscardUnknown(m)
}

var xxx_messageInfo_AdviceProto proto.InternalMessageInfo

func (m *AdviceProto) GetCheckers() map[string]*AdviceProto_Checker {
	if m != nil {
		return m.Checkers
	}
	return nil
}

type AdviceProto_Checker struct {
	Reports              []string `protobuf:"bytes,2,rep,name=reports,proto3" json:"reports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdviceProto_Checker) Reset()         { *m = AdviceProto_Checker{} }
func (m *AdviceProto_Checker) String() string { return proto.CompactTextString(m) }
func (*AdviceProto_Checker) ProtoMessage()    {}
func (*AdviceProto_Checker) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d317c37c959b8d2, []int{3, 1}
}

func (m *AdviceProto_Checker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdviceProto_Checker.Unmarshal(m, b)
}
func (m *AdviceProto_Checker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdviceProto_Checker.Marshal(b, m, deterministic)
}
func (m *AdviceProto_Checker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdviceProto_Checker.Merge(m, src)
}
func (m *AdviceProto_Checker) XXX_Size() int {
	return xxx_messageInfo_AdviceProto_Checker.Size(m)
}
func (m *AdviceProto_Checker) XXX_DiscardUnknown() {
	xxx_messageInfo_AdviceProto_Checker.DiscardUnknown(m)
}

var xxx_messageInfo_AdviceProto_Checker proto.InternalMessageInfo

func (m *AdviceProto_Checker) GetReports() []string {
	if m != nil {
		return m.Reports
	}
	return nil
}

func init() {
	proto.RegisterType((*TFProfTensorProto)(nil), "tensorflow.tfprof.TFProfTensorProto")
	proto.RegisterType((*GraphNodeProto)(nil), "tensorflow.tfprof.GraphNodeProto")
	proto.RegisterMapType((map[int32]*framework.TensorShapeProto)(nil), "tensorflow.tfprof.GraphNodeProto.InputShapesEntry")
	proto.RegisterType((*MultiGraphNodeProto)(nil), "tensorflow.tfprof.MultiGraphNodeProto")
	proto.RegisterType((*AdviceProto)(nil), "tensorflow.tfprof.AdviceProto")
	proto.RegisterMapType((map[string]*AdviceProto_Checker)(nil), "tensorflow.tfprof.AdviceProto.CheckersEntry")
	proto.RegisterType((*AdviceProto_Checker)(nil), "tensorflow.tfprof.AdviceProto.Checker")
}

func init() {
	proto.RegisterFile("tensorflow/core/profiler/tfprof_output.proto", fileDescriptor_1d317c37c959b8d2)
}

var fileDescriptor_1d317c37c959b8d2 = []byte{
	// 948 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdb, 0x6e, 0xdb, 0x46,
	0x10, 0x85, 0x4c, 0xdf, 0x34, 0x94, 0x75, 0x59, 0xcb, 0x09, 0xab, 0xc4, 0xad, 0xec, 0x36, 0xae,
	0x1a, 0x18, 0x52, 0xab, 0x1a, 0x41, 0x51, 0xb4, 0x0f, 0xbe, 0x24, 0x69, 0x1e, 0xd2, 0x18, 0x8c,
	0xdb, 0x87, 0xa2, 0x00, 0x41, 0x93, 0x2b, 0x9b, 0x30, 0xcd, 0x65, 0x97, 0xcb, 0xa4, 0xfa, 0x9b,
	0xfe, 0x41, 0x5f, 0xfa, 0x43, 0xfd, 0x93, 0x62, 0x67, 0x48, 0x91, 0xb4, 0xd5, 0xc4, 0xce, 0xe5,
	0x8d, 0x7b, 0xe6, 0xcc, 0xec, 0xe5, 0xcc, 0x1e, 0x2e, 0xec, 0x2a, 0x1e, 0x25, 0x42, 0x4e, 0x42,
	0xf1, 0x7a, 0xe4, 0x09, 0xc9, 0x47, 0xb1, 0x14, 0x93, 0x20, 0xe4, 0x72, 0xa4, 0x26, 0xfa, 0xd3,
	0x11, 0xa9, 0x8a, 0x53, 0x35, 0x8c, 0xa5, 0x50, 0x82, 0x75, 0x0a, 0xf6, 0x90, 0xe2, 0xbd, 0x6b,
	0x05, 0x26, 0xd2, 0xbd, 0xe4, 0xaf, 0x85, 0xbc, 0x18, 0x51, 0xc4, 0x49, 0xce, 0xdd, 0x98, 0x53,
	0x81, 0xde, 0x83, 0x37, 0xb0, 0xa7, 0x31, 0x4f, 0x88, 0xb6, 0xfd, 0x57, 0x0d, 0x3a, 0x27, 0x4f,
	0x8e, 0xa5, 0x98, 0x9c, 0x20, 0xff, 0x18, 0x67, 0x7f, 0x08, 0x4b, 0xbe, 0x66, 0x59, 0xb5, 0x7e,
	0x6d, 0xd0, 0x1c, 0x77, 0x87, 0xa5, 0xd5, 0x1c, 0xb9, 0xca, 0x3d, 0x99, 0xc6, 0xdc, 0x26, 0x0a,
	0xdb, 0x82, 0xc6, 0x2b, 0x37, 0x4c, 0xb9, 0xe3, 0x8b, 0xf4, 0x34, 0xe4, 0xd6, 0x42, 0xdf, 0x18,
	0xd4, 0x6c, 0x13, 0xb1, 0x23, 0x84, 0xd8, 0x67, 0x40, 0x43, 0x27, 0x88, 0xd4, 0xa3, 0x3d, 0xcb,
	0xe8, 0x1b, 0x03, 0xc3, 0x06, 0x84, 0x9e, 0x69, 0x84, 0xdd, 0x83, 0x3a, 0x11, 0x12, 0x25, 0xad,
	0xc5, 0xbe, 0x31, 0xa8, 0xdb, 0xab, 0x08, 0xbc, 0x54, 0x72, 0xfb, 0x6f, 0x80, 0xe6, 0x53, 0xe9,
	0xc6, 0xe7, 0x3f, 0x0b, 0x9f, 0xd3, 0xfa, 0x18, 0x2c, 0x46, 0xee, 0x25, 0x2d, 0xaf, 0x6e, 0xe3,
	0x37, 0x7b, 0x0a, 0x8d, 0xec, 0x18, 0x30, 0xd3, 0x6a, 0xf5, 0x6b, 0x03, 0x73, 0xfc, 0xc5, 0xf0,
	0xda, 0x41, 0x0e, 0xaf, 0xed, 0xd7, 0x36, 0x89, 0xf4, 0xab, 0x4e, 0xd4, 0x8b, 0x91, 0x69, 0xe4,
	0x78, 0x22, 0x8d, 0x94, 0xb5, 0xd1, 0xaf, 0x0d, 0x0c, 0x7b, 0x55, 0xa6, 0xd1, 0xa1, 0x1e, 0xeb,
	0xad, 0xf0, 0x3f, 0xb9, 0xe7, 0x5c, 0x06, 0x9e, 0x14, 0x89, 0xb5, 0x80, 0x61, 0xd0, 0xd0, 0x73,
	0x44, 0xd8, 0x23, 0xb8, 0xeb, 0x7a, 0x1e, 0x0f, 0xb9, 0x74, 0x95, 0x90, 0x4e, 0x99, 0xdc, 0x41,
	0xf2, 0x46, 0x29, 0xfc, 0xb8, 0xc8, 0xdb, 0x81, 0x96, 0x17, 0xa7, 0x15, 0x3e, 0x43, 0xfe, 0x9a,
	0x17, 0xa7, 0x25, 0xde, 0x97, 0xd0, 0x92, 0xfc, 0x8f, 0x94, 0x27, 0x8a, 0xfb, 0xce, 0xe9, 0x54,
	0xf1, 0xc4, 0x32, 0x90, 0xd7, 0x9c, 0xc1, 0x07, 0x1a, 0x65, 0x9b, 0x00, 0x31, 0x77, 0x2f, 0x32,
	0x8e, 0x85, 0x9c, 0xba, 0x46, 0x28, 0xfc, 0x00, 0x9a, 0x92, 0x27, 0x81, 0x9f, 0xba, 0x61, 0x46,
	0xf9, 0x84, 0xa6, 0xcb, 0x51, 0xa2, 0x6d, 0x41, 0x83, 0xfa, 0x32, 0x23, 0xf5, 0x90, 0x64, 0x12,
	0x46, 0x94, 0x4f, 0x01, 0x62, 0x57, 0x37, 0x97, 0xe2, 0x32, 0xb1, 0x16, 0xe9, 0x44, 0x0a, 0x44,
	0x9f, 0xe7, 0x24, 0x14, 0xae, 0x72, 0x44, 0x9c, 0x58, 0x6b, 0x74, 0x9e, 0x08, 0xbc, 0x88, 0x13,
	0x66, 0xc1, 0x8a, 0xcf, 0x5f, 0x05, 0x1e, 0x4f, 0x2c, 0x40, 0xdd, 0xf3, 0x21, 0xdb, 0x83, 0x3b,
	0x4a, 0x28, 0x37, 0x74, 0x7c, 0x3e, 0x09, 0xa2, 0x40, 0x05, 0x22, 0xd7, 0xe4, 0x2e, 0xd6, 0xe8,
	0x62, 0xf4, 0x68, 0x16, 0x24, 0x7d, 0x76, 0xa0, 0x45, 0x59, 0x85, 0x84, 0x77, 0x68, 0x5f, 0x08,
	0xdb, 0xb9, 0x8e, 0x0f, 0xa1, 0x43, 0xbc, 0xf2, 0x81, 0x2f, 0x23, 0x93, 0x0a, 0x94, 0x8e, 0x7c,
	0x1f, 0x36, 0x89, 0xfb, 0x7f, 0xc2, 0xae, 0x63, 0x5e, 0x0f, 0x49, 0xfb, 0x73, 0xd5, 0xfd, 0x06,
	0x36, 0xa8, 0xc4, 0x55, 0x8d, 0xbb, 0x98, 0xca, 0x30, 0x78, 0x58, 0x11, 0x7a, 0x9c, 0xa7, 0x5c,
	0x95, 0x7b, 0x05, 0x53, 0xd6, 0x69, 0x3f, 0x55, 0xcd, 0x07, 0xd0, 0xa6, 0x9c, 0x92, 0xf2, 0xf7,
	0xa8, 0x3b, 0x10, 0x3f, 0x9e, 0xc9, 0xff, 0x35, 0x74, 0xf3, 0xea, 0x95, 0x26, 0xb8, 0x5f, 0x5a,
	0x8f, 0x5d, 0xe9, 0x84, 0x5d, 0x20, 0xd4, 0xa9, 0xf4, 0xc3, 0x26, 0xf2, 0x69, 0xd6, 0x17, 0xa5,
	0xa6, 0xf8, 0x6a, 0xb6, 0x92, 0xa2, 0x35, 0x56, 0x4b, 0xc7, 0x7b, 0x5c, 0xf4, 0xc7, 0x4c, 0xb2,
	0xa2, 0x4b, 0x9a, 0x25, 0xc9, 0x9e, 0xe4, 0xad, 0xb2, 0x07, 0xcb, 0x68, 0x70, 0x89, 0x65, 0xf6,
	0x8d, 0x81, 0x39, 0xbe, 0x5f, 0xbe, 0xda, 0x74, 0x9b, 0x5f, 0xea, 0x38, 0x5d, 0xe9, 0x8c, 0xcb,
	0x7e, 0x81, 0x46, 0x10, 0xe9, 0xf5, 0x66, 0xb9, 0x6d, 0xcc, 0x1d, 0xcf, 0xb1, 0x85, 0xaa, 0xc7,
	0x0c, 0x9f, 0xe9, 0x2c, 0x2c, 0x98, 0x3c, 0x8e, 0x94, 0x9c, 0xda, 0x66, 0x50, 0x20, 0xec, 0x47,
	0x58, 0xf5, 0xce, 0x83, 0xd0, 0x97, 0x3c, 0xb2, 0x1a, 0x58, 0x72, 0xeb, 0xad, 0x25, 0xed, 0x59,
	0x4a, 0xef, 0x77, 0x68, 0x5f, 0xad, 0xcf, 0xda, 0x60, 0x5c, 0xf0, 0x29, 0x7a, 0xda, 0x92, 0xad,
	0x3f, 0xd9, 0x18, 0x96, 0xc8, 0xcb, 0x16, 0xd0, 0xcb, 0xde, 0xbc, 0x61, 0xa2, 0x7e, 0xbf, 0xf0,
	0x5d, 0x6d, 0xfb, 0x9f, 0x15, 0x58, 0x7f, 0x9e, 0x86, 0x2a, 0xb8, 0x81, 0x6d, 0xbe, 0x8f, 0xa1,
	0x35, 0x6e, 0x69, 0x68, 0x6b, 0x1f, 0xd0, 0xd0, 0xda, 0x6f, 0x37, 0xb4, 0xce, 0x4d, 0x0c, 0x8d,
	0xbd, 0xa7, 0xa1, 0x2d, 0x5d, 0x31, 0xb4, 0x0f, 0x6a, 0x2c, 0xcd, 0x77, 0x37, 0x96, 0xd6, 0x47,
	0x37, 0x96, 0xf5, 0x5b, 0x19, 0x4b, 0xf7, 0x96, 0xc6, 0xb2, 0xf1, 0x51, 0x8c, 0xa5, 0x3e, 0xcf,
	0x58, 0x0e, 0xc0, 0x3c, 0xd3, 0x17, 0xc5, 0x89, 0x84, 0x9f, 0xfd, 0x87, 0x6e, 0x74, 0x9d, 0xe1,
	0x2c, 0x1f, 0xeb, 0x1a, 0x85, 0x1f, 0x90, 0x3d, 0xed, 0xcc, 0x29, 0x30, 0xe7, 0x52, 0x16, 0xa6,
	0xb0, 0xfd, 0x6f, 0x0d, 0xcc, 0x7d, 0x5f, 0xff, 0xfd, 0xe8, 0xba, 0xfe, 0xa4, 0x6b, 0x72, 0xef,
	0x42, 0x6f, 0xb1, 0x86, 0x35, 0x77, 0xe7, 0xd4, 0x2c, 0x65, 0x0c, 0x0f, 0x33, 0x3a, 0x19, 0xd6,
	0x2c, 0xbb, 0xe7, 0xc1, 0x5a, 0x25, 0x54, 0xf6, 0x9a, 0x3a, 0x79, 0xcd, 0x0f, 0x55, 0xaf, 0xd9,
	0xb9, 0xd9, 0x4c, 0x25, 0xd7, 0xe9, 0x7d, 0x0e, 0x2b, 0x19, 0xaa, 0xff, 0xea, 0x92, 0xc7, 0x42,
	0xaa, 0x04, 0x9f, 0x83, 0x75, 0x3b, 0x1f, 0x1e, 0x1c, 0xfe, 0xb6, 0x7f, 0x16, 0xa8, 0xf3, 0xf4,
	0x74, 0xe8, 0x89, 0xcb, 0x51, 0xe9, 0x8d, 0x3a, 0xff, 0xf3, 0x4c, 0xd0, 0xe3, 0xb5, 0xf2, 0x60,
	0x3e, 0x5d, 0xc6, 0xb7, 0xeb, 0xb7, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xef, 0x2f, 0xe8, 0xb4,
	0x53, 0x0b, 0x00, 0x00,
}
