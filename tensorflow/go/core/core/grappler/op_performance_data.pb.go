// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/grappler/costs/op_performance_data.proto

package grappler

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	framework "github.com/tensorflow/tensorflow/tensorflow/go/core/framework"
	protobuf "github.com/tensorflow/tensorflow/tensorflow/go/core/protobuf"
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

// Description of the session when an op is run.
type SessionInfo struct {
	IntraOpParallelism   int64    `protobuf:"varint,1,opt,name=intra_op_parallelism,json=intraOpParallelism,proto3" json:"intra_op_parallelism,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionInfo) Reset()         { *m = SessionInfo{} }
func (m *SessionInfo) String() string { return proto.CompactTextString(m) }
func (*SessionInfo) ProtoMessage()    {}
func (*SessionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d91b4e44eab7273, []int{0}
}

func (m *SessionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionInfo.Unmarshal(m, b)
}
func (m *SessionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionInfo.Marshal(b, m, deterministic)
}
func (m *SessionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionInfo.Merge(m, src)
}
func (m *SessionInfo) XXX_Size() int {
	return xxx_messageInfo_SessionInfo.Size(m)
}
func (m *SessionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SessionInfo proto.InternalMessageInfo

func (m *SessionInfo) GetIntraOpParallelism() int64 {
	if m != nil {
		return m.IntraOpParallelism
	}
	return 0
}

// Description of an operation as well as the parameters expected to impact its
// performance.
type OpInfo struct {
	// The operation name.  There may be custom parameters in attrs.
	Op string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	// Custom parameters impacting the behavior of the op.
	Attr   map[string]*framework.AttrValue `protobuf:"bytes,2,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Inputs []*OpInfo_TensorProperties      `protobuf:"bytes,3,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// Optional description of the op outputs
	Outputs []*OpInfo_TensorProperties `protobuf:"bytes,5,rep,name=outputs,proto3" json:"outputs,omitempty"`
	// Device on which the operation is run.
	Device *protobuf.DeviceProperties `protobuf:"bytes,4,opt,name=device,proto3" json:"device,omitempty"`
	// Information about the session configs.
	SessionInfo          *SessionInfo `protobuf:"bytes,6,opt,name=session_info,json=sessionInfo,proto3" json:"session_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *OpInfo) Reset()         { *m = OpInfo{} }
func (m *OpInfo) String() string { return proto.CompactTextString(m) }
func (*OpInfo) ProtoMessage()    {}
func (*OpInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d91b4e44eab7273, []int{1}
}

func (m *OpInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpInfo.Unmarshal(m, b)
}
func (m *OpInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpInfo.Marshal(b, m, deterministic)
}
func (m *OpInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpInfo.Merge(m, src)
}
func (m *OpInfo) XXX_Size() int {
	return xxx_messageInfo_OpInfo.Size(m)
}
func (m *OpInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OpInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OpInfo proto.InternalMessageInfo

func (m *OpInfo) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *OpInfo) GetAttr() map[string]*framework.AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (m *OpInfo) GetInputs() []*OpInfo_TensorProperties {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *OpInfo) GetOutputs() []*OpInfo_TensorProperties {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *OpInfo) GetDevice() *protobuf.DeviceProperties {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *OpInfo) GetSessionInfo() *SessionInfo {
	if m != nil {
		return m.SessionInfo
	}
	return nil
}

// Input data types, shapes and values if known.
type OpInfo_TensorProperties struct {
	Dtype                framework.DataType          `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape                *framework.TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	Value                *framework.TensorProto      `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *OpInfo_TensorProperties) Reset()         { *m = OpInfo_TensorProperties{} }
func (m *OpInfo_TensorProperties) String() string { return proto.CompactTextString(m) }
func (*OpInfo_TensorProperties) ProtoMessage()    {}
func (*OpInfo_TensorProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d91b4e44eab7273, []int{1, 1}
}

func (m *OpInfo_TensorProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpInfo_TensorProperties.Unmarshal(m, b)
}
func (m *OpInfo_TensorProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpInfo_TensorProperties.Marshal(b, m, deterministic)
}
func (m *OpInfo_TensorProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpInfo_TensorProperties.Merge(m, src)
}
func (m *OpInfo_TensorProperties) XXX_Size() int {
	return xxx_messageInfo_OpInfo_TensorProperties.Size(m)
}
func (m *OpInfo_TensorProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_OpInfo_TensorProperties.DiscardUnknown(m)
}

var xxx_messageInfo_OpInfo_TensorProperties proto.InternalMessageInfo

func (m *OpInfo_TensorProperties) GetDtype() framework.DataType {
	if m != nil {
		return m.Dtype
	}
	return framework.DataType_DT_INVALID
}

func (m *OpInfo_TensorProperties) GetShape() *framework.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *OpInfo_TensorProperties) GetValue() *framework.TensorProto {
	if m != nil {
		return m.Value
	}
	return nil
}

type NormalDistribution struct {
	Mu                   float64  `protobuf:"fixed64,1,opt,name=mu,proto3" json:"mu,omitempty"`
	Sigma                float64  `protobuf:"fixed64,2,opt,name=sigma,proto3" json:"sigma,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NormalDistribution) Reset()         { *m = NormalDistribution{} }
func (m *NormalDistribution) String() string { return proto.CompactTextString(m) }
func (*NormalDistribution) ProtoMessage()    {}
func (*NormalDistribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d91b4e44eab7273, []int{2}
}

func (m *NormalDistribution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NormalDistribution.Unmarshal(m, b)
}
func (m *NormalDistribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NormalDistribution.Marshal(b, m, deterministic)
}
func (m *NormalDistribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NormalDistribution.Merge(m, src)
}
func (m *NormalDistribution) XXX_Size() int {
	return xxx_messageInfo_NormalDistribution.Size(m)
}
func (m *NormalDistribution) XXX_DiscardUnknown() {
	xxx_messageInfo_NormalDistribution.DiscardUnknown(m)
}

var xxx_messageInfo_NormalDistribution proto.InternalMessageInfo

func (m *NormalDistribution) GetMu() float64 {
	if m != nil {
		return m.Mu
	}
	return 0
}

func (m *NormalDistribution) GetSigma() float64 {
	if m != nil {
		return m.Sigma
	}
	return 0
}

type LogNormalDistribution struct {
	Mu                   float64  `protobuf:"fixed64,1,opt,name=mu,proto3" json:"mu,omitempty"`
	Sigma                float64  `protobuf:"fixed64,2,opt,name=sigma,proto3" json:"sigma,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogNormalDistribution) Reset()         { *m = LogNormalDistribution{} }
func (m *LogNormalDistribution) String() string { return proto.CompactTextString(m) }
func (*LogNormalDistribution) ProtoMessage()    {}
func (*LogNormalDistribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d91b4e44eab7273, []int{3}
}

func (m *LogNormalDistribution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogNormalDistribution.Unmarshal(m, b)
}
func (m *LogNormalDistribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogNormalDistribution.Marshal(b, m, deterministic)
}
func (m *LogNormalDistribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogNormalDistribution.Merge(m, src)
}
func (m *LogNormalDistribution) XXX_Size() int {
	return xxx_messageInfo_LogNormalDistribution.Size(m)
}
func (m *LogNormalDistribution) XXX_DiscardUnknown() {
	xxx_messageInfo_LogNormalDistribution.DiscardUnknown(m)
}

var xxx_messageInfo_LogNormalDistribution proto.InternalMessageInfo

func (m *LogNormalDistribution) GetMu() float64 {
	if m != nil {
		return m.Mu
	}
	return 0
}

func (m *LogNormalDistribution) GetSigma() float64 {
	if m != nil {
		return m.Sigma
	}
	return 0
}

// Performance data for tensorflow operations
type OpPerformance struct {
	// The op
	Op *OpInfo `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	// Information about the session configs.
	SessionInfo *SessionInfo `protobuf:"bytes,12,opt,name=session_info,json=sessionInfo,proto3" json:"session_info,omitempty"` // Deprecated: Do not use.
	// The node name (optional). Makes it easier to associate the performance data
	// with a specific graph node.
	Node string `protobuf:"bytes,5,opt,name=node,proto3" json:"node,omitempty"`
	// Temporary memory used by this node (in bytes).
	TemporaryMemorySize int64 `protobuf:"varint,2,opt,name=temporary_memory_size,json=temporaryMemorySize,proto3" json:"temporary_memory_size,omitempty"`
	// Time it takes to run the op (in nanoseconds).
	ComputeCost int64 `protobuf:"varint,3,opt,name=compute_cost,json=computeCost,proto3" json:"compute_cost,omitempty"`
	// Analytical compute cost (in nanoseconds).
	ComputeTime int64 `protobuf:"varint,6,opt,name=compute_time,json=computeTime,proto3" json:"compute_time,omitempty"`
	// Analytical memory access cost (in nanoseconds).
	MemoryTime int64 `protobuf:"varint,7,opt,name=memory_time,json=memoryTime,proto3" json:"memory_time,omitempty"`
	// Percentage of theoretical compute performance.
	ComputeEfficiency float64 `protobuf:"fixed64,4,opt,name=compute_efficiency,json=computeEfficiency,proto3" json:"compute_efficiency,omitempty"`
	// Percentage of theoretical memory performance.
	MemoryEfficiency float64 `protobuf:"fixed64,8,opt,name=memory_efficiency,json=memoryEfficiency,proto3" json:"memory_efficiency,omitempty"`
	// Expected execution time, modeled using one of 2 possible distributions.
	//
	// Types that are valid to be assigned to ExecutionTime:
	//	*OpPerformance_ExecutionTimeNormal
	//	*OpPerformance_ExecutionTimeLogNormal
	ExecutionTime        isOpPerformance_ExecutionTime `protobuf_oneof:"execution_time"`
	OpMemory             *OpPerformance_OpMemory       `protobuf:"bytes,9,opt,name=op_memory,json=opMemory,proto3" json:"op_memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *OpPerformance) Reset()         { *m = OpPerformance{} }
func (m *OpPerformance) String() string { return proto.CompactTextString(m) }
func (*OpPerformance) ProtoMessage()    {}
func (*OpPerformance) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d91b4e44eab7273, []int{4}
}

func (m *OpPerformance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpPerformance.Unmarshal(m, b)
}
func (m *OpPerformance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpPerformance.Marshal(b, m, deterministic)
}
func (m *OpPerformance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpPerformance.Merge(m, src)
}
func (m *OpPerformance) XXX_Size() int {
	return xxx_messageInfo_OpPerformance.Size(m)
}
func (m *OpPerformance) XXX_DiscardUnknown() {
	xxx_messageInfo_OpPerformance.DiscardUnknown(m)
}

var xxx_messageInfo_OpPerformance proto.InternalMessageInfo

func (m *OpPerformance) GetOp() *OpInfo {
	if m != nil {
		return m.Op
	}
	return nil
}

// Deprecated: Do not use.
func (m *OpPerformance) GetSessionInfo() *SessionInfo {
	if m != nil {
		return m.SessionInfo
	}
	return nil
}

func (m *OpPerformance) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *OpPerformance) GetTemporaryMemorySize() int64 {
	if m != nil {
		return m.TemporaryMemorySize
	}
	return 0
}

func (m *OpPerformance) GetComputeCost() int64 {
	if m != nil {
		return m.ComputeCost
	}
	return 0
}

func (m *OpPerformance) GetComputeTime() int64 {
	if m != nil {
		return m.ComputeTime
	}
	return 0
}

func (m *OpPerformance) GetMemoryTime() int64 {
	if m != nil {
		return m.MemoryTime
	}
	return 0
}

func (m *OpPerformance) GetComputeEfficiency() float64 {
	if m != nil {
		return m.ComputeEfficiency
	}
	return 0
}

func (m *OpPerformance) GetMemoryEfficiency() float64 {
	if m != nil {
		return m.MemoryEfficiency
	}
	return 0
}

type isOpPerformance_ExecutionTime interface {
	isOpPerformance_ExecutionTime()
}

type OpPerformance_ExecutionTimeNormal struct {
	ExecutionTimeNormal *NormalDistribution `protobuf:"bytes,10,opt,name=execution_time_normal,json=executionTimeNormal,proto3,oneof"`
}

type OpPerformance_ExecutionTimeLogNormal struct {
	ExecutionTimeLogNormal *LogNormalDistribution `protobuf:"bytes,11,opt,name=execution_time_log_normal,json=executionTimeLogNormal,proto3,oneof"`
}

func (*OpPerformance_ExecutionTimeNormal) isOpPerformance_ExecutionTime() {}

func (*OpPerformance_ExecutionTimeLogNormal) isOpPerformance_ExecutionTime() {}

func (m *OpPerformance) GetExecutionTime() isOpPerformance_ExecutionTime {
	if m != nil {
		return m.ExecutionTime
	}
	return nil
}

func (m *OpPerformance) GetExecutionTimeNormal() *NormalDistribution {
	if x, ok := m.GetExecutionTime().(*OpPerformance_ExecutionTimeNormal); ok {
		return x.ExecutionTimeNormal
	}
	return nil
}

func (m *OpPerformance) GetExecutionTimeLogNormal() *LogNormalDistribution {
	if x, ok := m.GetExecutionTime().(*OpPerformance_ExecutionTimeLogNormal); ok {
		return x.ExecutionTimeLogNormal
	}
	return nil
}

func (m *OpPerformance) GetOpMemory() *OpPerformance_OpMemory {
	if m != nil {
		return m.OpMemory
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OpPerformance) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OpPerformance_ExecutionTimeNormal)(nil),
		(*OpPerformance_ExecutionTimeLogNormal)(nil),
	}
}

// Memory usage data for a tensorflow operation.
type OpPerformance_OpMemory struct {
	// The output information may have memory usage and output shapes.
	OutputMemory []int64 `protobuf:"varint,1,rep,packed,name=output_memory,json=outputMemory,proto3" json:"output_memory,omitempty"`
	// Temp and persistent memory allocated by this node.
	TempMemory             int64    `protobuf:"varint,2,opt,name=temp_memory,json=tempMemory,proto3" json:"temp_memory,omitempty"`
	PersistentMemory       int64    `protobuf:"varint,4,opt,name=persistent_memory,json=persistentMemory,proto3" json:"persistent_memory,omitempty"`
	DeviceTempMemory       int64    `protobuf:"varint,3,opt,name=device_temp_memory,json=deviceTempMemory,proto3" json:"device_temp_memory,omitempty"`                   // Deprecated: Do not use.
	DevicePersistentMemory int64    `protobuf:"varint,5,opt,name=device_persistent_memory,json=devicePersistentMemory,proto3" json:"device_persistent_memory,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *OpPerformance_OpMemory) Reset()         { *m = OpPerformance_OpMemory{} }
func (m *OpPerformance_OpMemory) String() string { return proto.CompactTextString(m) }
func (*OpPerformance_OpMemory) ProtoMessage()    {}
func (*OpPerformance_OpMemory) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d91b4e44eab7273, []int{4, 0}
}

func (m *OpPerformance_OpMemory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpPerformance_OpMemory.Unmarshal(m, b)
}
func (m *OpPerformance_OpMemory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpPerformance_OpMemory.Marshal(b, m, deterministic)
}
func (m *OpPerformance_OpMemory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpPerformance_OpMemory.Merge(m, src)
}
func (m *OpPerformance_OpMemory) XXX_Size() int {
	return xxx_messageInfo_OpPerformance_OpMemory.Size(m)
}
func (m *OpPerformance_OpMemory) XXX_DiscardUnknown() {
	xxx_messageInfo_OpPerformance_OpMemory.DiscardUnknown(m)
}

var xxx_messageInfo_OpPerformance_OpMemory proto.InternalMessageInfo

func (m *OpPerformance_OpMemory) GetOutputMemory() []int64 {
	if m != nil {
		return m.OutputMemory
	}
	return nil
}

func (m *OpPerformance_OpMemory) GetTempMemory() int64 {
	if m != nil {
		return m.TempMemory
	}
	return 0
}

func (m *OpPerformance_OpMemory) GetPersistentMemory() int64 {
	if m != nil {
		return m.PersistentMemory
	}
	return 0
}

// Deprecated: Do not use.
func (m *OpPerformance_OpMemory) GetDeviceTempMemory() int64 {
	if m != nil {
		return m.DeviceTempMemory
	}
	return 0
}

// Deprecated: Do not use.
func (m *OpPerformance_OpMemory) GetDevicePersistentMemory() int64 {
	if m != nil {
		return m.DevicePersistentMemory
	}
	return 0
}

// A collection of OpPerformance data points.
type OpPerformanceList struct {
	OpPerformance        []*OpPerformance `protobuf:"bytes,1,rep,name=op_performance,json=opPerformance,proto3" json:"op_performance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *OpPerformanceList) Reset()         { *m = OpPerformanceList{} }
func (m *OpPerformanceList) String() string { return proto.CompactTextString(m) }
func (*OpPerformanceList) ProtoMessage()    {}
func (*OpPerformanceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d91b4e44eab7273, []int{5}
}

func (m *OpPerformanceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpPerformanceList.Unmarshal(m, b)
}
func (m *OpPerformanceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpPerformanceList.Marshal(b, m, deterministic)
}
func (m *OpPerformanceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpPerformanceList.Merge(m, src)
}
func (m *OpPerformanceList) XXX_Size() int {
	return xxx_messageInfo_OpPerformanceList.Size(m)
}
func (m *OpPerformanceList) XXX_DiscardUnknown() {
	xxx_messageInfo_OpPerformanceList.DiscardUnknown(m)
}

var xxx_messageInfo_OpPerformanceList proto.InternalMessageInfo

func (m *OpPerformanceList) GetOpPerformance() []*OpPerformance {
	if m != nil {
		return m.OpPerformance
	}
	return nil
}

func init() {
	proto.RegisterType((*SessionInfo)(nil), "tensorflow.SessionInfo")
	proto.RegisterType((*OpInfo)(nil), "tensorflow.OpInfo")
	proto.RegisterMapType((map[string]*framework.AttrValue)(nil), "tensorflow.OpInfo.AttrEntry")
	proto.RegisterType((*OpInfo_TensorProperties)(nil), "tensorflow.OpInfo.TensorProperties")
	proto.RegisterType((*NormalDistribution)(nil), "tensorflow.NormalDistribution")
	proto.RegisterType((*LogNormalDistribution)(nil), "tensorflow.LogNormalDistribution")
	proto.RegisterType((*OpPerformance)(nil), "tensorflow.OpPerformance")
	proto.RegisterType((*OpPerformance_OpMemory)(nil), "tensorflow.OpPerformance.OpMemory")
	proto.RegisterType((*OpPerformanceList)(nil), "tensorflow.OpPerformanceList")
}

func init() {
	proto.RegisterFile("tensorflow/core/grappler/costs/op_performance_data.proto", fileDescriptor_4d91b4e44eab7273)
}

var fileDescriptor_4d91b4e44eab7273 = []byte{
	// 880 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xc7, 0xbb, 0x5e, 0xdb, 0x4d, 0xde, 0x26, 0x91, 0x33, 0x4d, 0xca, 0xd6, 0x42, 0x34, 0x75,
	0x05, 0x8a, 0x1a, 0x6a, 0x47, 0x86, 0x43, 0x15, 0x28, 0xa5, 0xa1, 0x45, 0x45, 0x2a, 0x8d, 0x35,
	0x31, 0x1c, 0x38, 0xb0, 0xda, 0x6c, 0xc6, 0xee, 0xa8, 0xde, 0x9d, 0xd1, 0xcc, 0x6c, 0x8b, 0xfb,
	0xdf, 0x70, 0xe7, 0x5f, 0xe3, 0x0c, 0x47, 0xb4, 0x6f, 0xf6, 0xa7, 0x9d, 0x06, 0xc4, 0x6d, 0x3d,
	0xef, 0xf3, 0xfd, 0xce, 0x9b, 0xf7, 0xde, 0xee, 0x18, 0x1e, 0x19, 0x96, 0x68, 0xa1, 0x66, 0x0b,
	0xf1, 0x6e, 0x14, 0x09, 0xc5, 0x46, 0x73, 0x15, 0x4a, 0xb9, 0x60, 0x6a, 0x14, 0x09, 0x6d, 0xf4,
	0x48, 0xc8, 0x40, 0x32, 0x35, 0x13, 0x2a, 0x0e, 0x93, 0x88, 0x05, 0x97, 0xa1, 0x09, 0x87, 0x52,
	0x09, 0x23, 0x08, 0x54, 0xca, 0xfe, 0x67, 0xab, 0x2e, 0x33, 0x15, 0xc6, 0xec, 0x9d, 0x50, 0x6f,
	0x46, 0x36, 0x62, 0x35, 0xfd, 0xcf, 0xff, 0x8d, 0x0b, 0xf4, 0xeb, 0x50, 0xb2, 0x9c, 0xfe, 0xf4,
	0x1a, 0x7a, 0x29, 0x99, 0xce, 0xb1, 0x07, 0x1f, 0xc6, 0x42, 0x63, 0x54, 0xf0, 0x36, 0x5c, 0xa4,
	0x85, 0xe5, 0xf1, 0x2a, 0x8b, 0xcb, 0x17, 0xe9, 0x6c, 0x74, 0xc9, 0xde, 0xf2, 0x88, 0x05, 0x52,
	0x09, 0xc9, 0x94, 0xe1, 0x85, 0xfb, 0xe0, 0x09, 0x78, 0xe7, 0x4c, 0x6b, 0x2e, 0x92, 0x1f, 0x92,
	0x99, 0x20, 0xc7, 0xb0, 0xc7, 0x13, 0xa3, 0xc2, 0x20, 0x2b, 0x4c, 0xa8, 0xc2, 0xc5, 0x82, 0x2d,
	0xb8, 0x8e, 0x7d, 0xe7, 0xc0, 0x39, 0x74, 0x29, 0xc1, 0xd8, 0x99, 0x9c, 0x54, 0x91, 0xc1, 0x1f,
	0x6d, 0xe8, 0x9e, 0x49, 0x14, 0xef, 0x40, 0x4b, 0x48, 0x44, 0x37, 0x69, 0x4b, 0x48, 0x72, 0x0c,
	0xed, 0x2c, 0x43, 0xbf, 0x75, 0xe0, 0x1e, 0x7a, 0xe3, 0x8f, 0x87, 0x55, 0x72, 0x43, 0xab, 0x18,
	0x3e, 0x35, 0x46, 0x3d, 0x4f, 0x8c, 0x5a, 0x52, 0x24, 0xc9, 0x57, 0xd0, 0xe5, 0x89, 0x4c, 0x8d,
	0xf6, 0x5d, 0xd4, 0xdc, 0xbf, 0x42, 0x33, 0xc5, 0x95, 0x49, 0x79, 0x10, 0x9a, 0x4b, 0xc8, 0x63,
	0xb8, 0x29, 0x52, 0x83, 0xea, 0xce, 0x7f, 0x57, 0x17, 0x1a, 0xf2, 0x25, 0x74, 0x6d, 0x91, 0xfc,
	0xf6, 0x81, 0xb3, 0x9a, 0xef, 0x33, 0x8c, 0xd4, 0x37, 0xb5, 0x2c, 0x39, 0x81, 0x2d, 0x6d, 0xeb,
	0x17, 0xf0, 0x64, 0x26, 0xfc, 0x2e, 0x6a, 0x3f, 0xaa, 0x6b, 0x6b, 0xf5, 0xa5, 0x9e, 0xae, 0x7e,
	0xf4, 0x5f, 0xc1, 0x66, 0x59, 0x00, 0xd2, 0x03, 0xf7, 0x0d, 0x5b, 0xe6, 0xd5, 0xcb, 0x1e, 0xc9,
	0x11, 0x74, 0xb0, 0xb7, 0x7e, 0x0b, 0x3d, 0xf7, 0xeb, 0x9e, 0x99, 0xee, 0xe7, 0x2c, 0x48, 0x2d,
	0x73, 0xd2, 0x7a, 0xe4, 0xf4, 0x7f, 0x77, 0xa0, 0xb7, 0x7a, 0x3e, 0xf2, 0x00, 0x3a, 0x97, 0xd9,
	0x38, 0xa1, 0xf3, 0xce, 0x78, 0xaf, 0x71, 0xaa, 0xd0, 0x84, 0xd3, 0xa5, 0x64, 0xd4, 0x22, 0x64,
	0x0c, 0x1d, 0x1c, 0xd0, 0x7c, 0xc7, 0x46, 0x05, 0xac, 0xf1, 0x79, 0x16, 0x9e, 0x64, 0x93, 0x43,
	0x2d, 0x4a, 0x1e, 0x16, 0x59, 0xba, 0xeb, 0x27, 0x2f, 0x93, 0xc9, 0x70, 0xa4, 0x06, 0x27, 0x40,
	0x5e, 0x65, 0x6f, 0xdb, 0xe2, 0x19, 0xd7, 0x46, 0xf1, 0x8b, 0xd4, 0x70, 0x91, 0x64, 0x93, 0x13,
	0xa7, 0x98, 0xa1, 0x43, 0x5b, 0x71, 0x4a, 0xf6, 0xa0, 0xa3, 0xf9, 0x3c, 0x0e, 0x31, 0x11, 0x87,
	0xda, 0x1f, 0x83, 0xc7, 0xb0, 0xff, 0x52, 0xcc, 0xff, 0xb7, 0xfc, 0xaf, 0x2e, 0x6c, 0x9f, 0xc9,
	0x49, 0xf5, 0xba, 0x93, 0x41, 0x39, 0xb0, 0xde, 0x98, 0xac, 0x0f, 0x0b, 0x0e, 0xf1, 0x37, 0x2b,
	0x0d, 0xde, 0xba, 0xb6, 0xc1, 0xa7, 0x2d, 0xdf, 0x69, 0x34, 0x99, 0x10, 0x68, 0x27, 0xe2, 0x92,
	0xf9, 0x1d, 0x6c, 0x2c, 0x3e, 0x93, 0x31, 0xec, 0x1b, 0x16, 0x4b, 0xa1, 0x42, 0xb5, 0x0c, 0x62,
	0x16, 0x0b, 0xb5, 0x0c, 0x34, 0x7f, 0x6f, 0xeb, 0xee, 0xd2, 0x5b, 0x65, 0xf0, 0x47, 0x8c, 0x9d,
	0xf3, 0xf7, 0x8c, 0xdc, 0x83, 0xad, 0x48, 0xc4, 0x32, 0x35, 0x2c, 0xc8, 0x3e, 0x5d, 0x58, 0x6e,
	0x97, 0x7a, 0xf9, 0xda, 0x77, 0x42, 0x9b, 0x3a, 0x62, 0x78, 0xcc, 0x70, 0x16, 0x2b, 0x64, 0xca,
	0x63, 0x46, 0xee, 0x82, 0x97, 0xef, 0x87, 0xc4, 0x4d, 0x24, 0xc0, 0x2e, 0x21, 0xf0, 0x10, 0x48,
	0xe1, 0xc1, 0x66, 0x33, 0x1e, 0x71, 0x96, 0x44, 0x4b, 0x7c, 0x23, 0x1c, 0xba, 0x9b, 0x47, 0x9e,
	0x97, 0x01, 0x72, 0x04, 0xbb, 0xb9, 0x5f, 0x8d, 0xde, 0x40, 0xba, 0x67, 0x03, 0x35, 0x78, 0x0a,
	0xfb, 0xec, 0x37, 0x16, 0x61, 0xcf, 0x70, 0xff, 0x20, 0xc1, 0x5e, 0xfa, 0x80, 0x35, 0xfd, 0xa4,
	0x5e, 0xd3, 0xf5, 0x2e, 0xbf, 0xb8, 0x41, 0x6f, 0x95, 0xf2, 0x2c, 0x57, 0x8b, 0x90, 0x5f, 0xe1,
	0xce, 0x8a, 0xeb, 0x42, 0xcc, 0x0b, 0x67, 0x0f, 0x9d, 0xef, 0xd5, 0x9d, 0xaf, 0x1c, 0xa1, 0x17,
	0x37, 0xe8, 0xed, 0x86, 0x79, 0x49, 0x91, 0x27, 0xb0, 0x29, 0x64, 0xde, 0x25, 0x7f, 0x13, 0xfd,
	0x06, 0xcd, 0x59, 0xa9, 0x8d, 0xd4, 0xf0, 0x4c, 0xda, 0x9e, 0xd1, 0x0d, 0x91, 0x3f, 0xf5, 0xff,
	0x74, 0x60, 0xa3, 0x58, 0x26, 0xf7, 0x61, 0xdb, 0x7e, 0x70, 0x0a, 0x47, 0xe7, 0xc0, 0x3d, 0x74,
	0xe9, 0x96, 0x5d, 0xcc, 0xa1, 0xbb, 0xe0, 0x65, 0x23, 0x50, 0x20, 0x76, 0x2a, 0x20, 0x5b, 0xca,
	0x81, 0x23, 0xd8, 0x95, 0x4c, 0x69, 0xae, 0x0d, 0x4b, 0x4a, 0xa7, 0x36, 0x62, 0xbd, 0x2a, 0x90,
	0xc3, 0xc7, 0x40, 0xf2, 0xaf, 0x7f, 0xdd, 0x14, 0xe7, 0x07, 0xc7, 0xb5, 0x67, 0xa3, 0xd3, 0xca,
	0xfe, 0x6b, 0xf0, 0x8b, 0xfb, 0x62, 0x6d, 0x97, 0x4e, 0xa9, 0xbb, 0x6d, 0x99, 0xc9, 0xca, 0x7e,
	0xa7, 0x3d, 0xd8, 0x69, 0x36, 0x64, 0xf0, 0x13, 0xec, 0x36, 0xaa, 0xf4, 0x92, 0x6b, 0x43, 0xbe,
	0x85, 0x9d, 0xe6, 0xed, 0x8b, 0xa5, 0xf0, 0xc6, 0x77, 0x3e, 0x58, 0x5c, 0xba, 0x2d, 0xea, 0x3f,
	0x4f, 0xbf, 0xff, 0xe5, 0xe9, 0x9c, 0x9b, 0xd7, 0xe9, 0xc5, 0x30, 0x12, 0xf1, 0xa8, 0x76, 0xf5,
	0x5d, 0xfd, 0x38, 0x17, 0xf6, 0x4e, 0x6c, 0xfc, 0x0f, 0xf8, 0xdb, 0x71, 0x2e, 0xba, 0x78, 0x15,
	0x7e, 0xf1, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x92, 0x02, 0x51, 0x2d, 0x08, 0x00, 0x00,
}
