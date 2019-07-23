// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/compiler/xrt/xrt.proto

package xrt

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	tf2xla "github.com/tensorflow/tensorflow/tensorflow/go/core/compiler/tf2xla"
	xla "github.com/tensorflow/tensorflow/tensorflow/go/core/compiler/xla"
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

type DeviceAssignment struct {
	// As many ComputationDevice as many there are computations (number
	// of cores per replica).
	ComputationDevices   []*DeviceAssignment_ComputationDevice `protobuf:"bytes,1,rep,name=computation_devices,json=computationDevices,proto3" json:"computation_devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *DeviceAssignment) Reset()         { *m = DeviceAssignment{} }
func (m *DeviceAssignment) String() string { return proto.CompactTextString(m) }
func (*DeviceAssignment) ProtoMessage()    {}
func (*DeviceAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef605a9242ff614, []int{0}
}

func (m *DeviceAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceAssignment.Unmarshal(m, b)
}
func (m *DeviceAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceAssignment.Marshal(b, m, deterministic)
}
func (m *DeviceAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceAssignment.Merge(m, src)
}
func (m *DeviceAssignment) XXX_Size() int {
	return xxx_messageInfo_DeviceAssignment.Size(m)
}
func (m *DeviceAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceAssignment proto.InternalMessageInfo

func (m *DeviceAssignment) GetComputationDevices() []*DeviceAssignment_ComputationDevice {
	if m != nil {
		return m.ComputationDevices
	}
	return nil
}

type DeviceAssignment_ComputationDevice struct {
	// As many replicas as there are in the replicated computation.
	ReplicaDevices       []*DeviceAssignment_ComputationDevice_DeviceMeshCoordinates `protobuf:"bytes,1,rep,name=replica_devices,json=replicaDevices,proto3" json:"replica_devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                    `json:"-"`
	XXX_unrecognized     []byte                                                      `json:"-"`
	XXX_sizecache        int32                                                       `json:"-"`
}

func (m *DeviceAssignment_ComputationDevice) Reset()         { *m = DeviceAssignment_ComputationDevice{} }
func (m *DeviceAssignment_ComputationDevice) String() string { return proto.CompactTextString(m) }
func (*DeviceAssignment_ComputationDevice) ProtoMessage()    {}
func (*DeviceAssignment_ComputationDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef605a9242ff614, []int{0, 0}
}

func (m *DeviceAssignment_ComputationDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceAssignment_ComputationDevice.Unmarshal(m, b)
}
func (m *DeviceAssignment_ComputationDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceAssignment_ComputationDevice.Marshal(b, m, deterministic)
}
func (m *DeviceAssignment_ComputationDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceAssignment_ComputationDevice.Merge(m, src)
}
func (m *DeviceAssignment_ComputationDevice) XXX_Size() int {
	return xxx_messageInfo_DeviceAssignment_ComputationDevice.Size(m)
}
func (m *DeviceAssignment_ComputationDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceAssignment_ComputationDevice.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceAssignment_ComputationDevice proto.InternalMessageInfo

func (m *DeviceAssignment_ComputationDevice) GetReplicaDevices() []*DeviceAssignment_ComputationDevice_DeviceMeshCoordinates {
	if m != nil {
		return m.ReplicaDevices
	}
	return nil
}

type DeviceAssignment_ComputationDevice_DeviceMeshCoordinates struct {
	// The mesh coordinates for the device. Usually (X, Y, Core), in the order
	// in which they are returned in the TopologyProto.
	//  X    = value(0)
	//  Y    = value(1)
	//  Core = value(2)
	Value                []int32  `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceAssignment_ComputationDevice_DeviceMeshCoordinates) Reset() {
	*m = DeviceAssignment_ComputationDevice_DeviceMeshCoordinates{}
}
func (m *DeviceAssignment_ComputationDevice_DeviceMeshCoordinates) String() string {
	return proto.CompactTextString(m)
}
func (*DeviceAssignment_ComputationDevice_DeviceMeshCoordinates) ProtoMessage() {}
func (*DeviceAssignment_ComputationDevice_DeviceMeshCoordinates) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef605a9242ff614, []int{0, 0, 0}
}

func (m *DeviceAssignment_ComputationDevice_DeviceMeshCoordinates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceAssignment_ComputationDevice_DeviceMeshCoordinates.Unmarshal(m, b)
}
func (m *DeviceAssignment_ComputationDevice_DeviceMeshCoordinates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceAssignment_ComputationDevice_DeviceMeshCoordinates.Marshal(b, m, deterministic)
}
func (m *DeviceAssignment_ComputationDevice_DeviceMeshCoordinates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceAssignment_ComputationDevice_DeviceMeshCoordinates.Merge(m, src)
}
func (m *DeviceAssignment_ComputationDevice_DeviceMeshCoordinates) XXX_Size() int {
	return xxx_messageInfo_DeviceAssignment_ComputationDevice_DeviceMeshCoordinates.Size(m)
}
func (m *DeviceAssignment_ComputationDevice_DeviceMeshCoordinates) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceAssignment_ComputationDevice_DeviceMeshCoordinates.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceAssignment_ComputationDevice_DeviceMeshCoordinates proto.InternalMessageInfo

func (m *DeviceAssignment_ComputationDevice_DeviceMeshCoordinates) GetValue() []int32 {
	if m != nil {
		return m.Value
	}
	return nil
}

// Options for an XLA compilation.
type XLAComputationConfig struct {
	// The number of replicas the computation will be run on. If this is
	// default (0) it is interpreted as 1.
	NumReplicas int32 `protobuf:"varint,1,opt,name=num_replicas,json=numReplicas,proto3" json:"num_replicas,omitempty"`
	// The number of "model-parallel" cores per replica. If this is
	// default (0) it is interpreted as 1.
	NumCoresPerReplica int32 `protobuf:"varint,2,opt,name=num_cores_per_replica,json=numCoresPerReplica,proto3" json:"num_cores_per_replica,omitempty"`
	// Optional metadata about host sends and recvs.
	HostComputeMetadata *tf2xla.HostComputeMetadata `protobuf:"bytes,3,opt,name=host_compute_metadata,json=hostComputeMetadata,proto3" json:"host_compute_metadata,omitempty"`
	// The arg/result shapes for the whole computation.
	ProgramShape *xla.ProgramShapeProto `protobuf:"bytes,4,opt,name=program_shape,json=programShape,proto3" json:"program_shape,omitempty"`
	// The arg/result shapes for each core of a model-parallel
	// computation. per_core_args_and_result_shapes is optional for a
	// single-core computation.
	PerCoreProgramShape []*xla.ProgramShapeProto `protobuf:"bytes,5,rep,name=per_core_program_shape,json=perCoreProgramShape,proto3" json:"per_core_program_shape,omitempty"`
	// Describes how replicated computation instances should be assigned to
	// devices. There are num_cores_per_replica computations, and each one will be
	// sent and executed to the set of replica device numbers described in the
	// DeviceAssignment proto.
	DeviceAssignment *DeviceAssignment `protobuf:"bytes,6,opt,name=device_assignment,json=deviceAssignment,proto3" json:"device_assignment,omitempty"`
	// The debugging options to be passed to the XLA compilation process.
	DebugOptions         *xla.DebugOptions `protobuf:"bytes,7,opt,name=debug_options,json=debugOptions,proto3" json:"debug_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *XLAComputationConfig) Reset()         { *m = XLAComputationConfig{} }
func (m *XLAComputationConfig) String() string { return proto.CompactTextString(m) }
func (*XLAComputationConfig) ProtoMessage()    {}
func (*XLAComputationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef605a9242ff614, []int{1}
}

func (m *XLAComputationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XLAComputationConfig.Unmarshal(m, b)
}
func (m *XLAComputationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XLAComputationConfig.Marshal(b, m, deterministic)
}
func (m *XLAComputationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XLAComputationConfig.Merge(m, src)
}
func (m *XLAComputationConfig) XXX_Size() int {
	return xxx_messageInfo_XLAComputationConfig.Size(m)
}
func (m *XLAComputationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_XLAComputationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_XLAComputationConfig proto.InternalMessageInfo

func (m *XLAComputationConfig) GetNumReplicas() int32 {
	if m != nil {
		return m.NumReplicas
	}
	return 0
}

func (m *XLAComputationConfig) GetNumCoresPerReplica() int32 {
	if m != nil {
		return m.NumCoresPerReplica
	}
	return 0
}

func (m *XLAComputationConfig) GetHostComputeMetadata() *tf2xla.HostComputeMetadata {
	if m != nil {
		return m.HostComputeMetadata
	}
	return nil
}

func (m *XLAComputationConfig) GetProgramShape() *xla.ProgramShapeProto {
	if m != nil {
		return m.ProgramShape
	}
	return nil
}

func (m *XLAComputationConfig) GetPerCoreProgramShape() []*xla.ProgramShapeProto {
	if m != nil {
		return m.PerCoreProgramShape
	}
	return nil
}

func (m *XLAComputationConfig) GetDeviceAssignment() *DeviceAssignment {
	if m != nil {
		return m.DeviceAssignment
	}
	return nil
}

func (m *XLAComputationConfig) GetDebugOptions() *xla.DebugOptions {
	if m != nil {
		return m.DebugOptions
	}
	return nil
}

// Options and XLA computation for a compilation.
type XLAComputation struct {
	Config               *XLAComputationConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	HloSnapshot          *xla.HloSnapshot      `protobuf:"bytes,2,opt,name=hlo_snapshot,json=hloSnapshot,proto3" json:"hlo_snapshot,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *XLAComputation) Reset()         { *m = XLAComputation{} }
func (m *XLAComputation) String() string { return proto.CompactTextString(m) }
func (*XLAComputation) ProtoMessage()    {}
func (*XLAComputation) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef605a9242ff614, []int{2}
}

func (m *XLAComputation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XLAComputation.Unmarshal(m, b)
}
func (m *XLAComputation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XLAComputation.Marshal(b, m, deterministic)
}
func (m *XLAComputation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XLAComputation.Merge(m, src)
}
func (m *XLAComputation) XXX_Size() int {
	return xxx_messageInfo_XLAComputation.Size(m)
}
func (m *XLAComputation) XXX_DiscardUnknown() {
	xxx_messageInfo_XLAComputation.DiscardUnknown(m)
}

var xxx_messageInfo_XLAComputation proto.InternalMessageInfo

func (m *XLAComputation) GetConfig() *XLAComputationConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *XLAComputation) GetHloSnapshot() *xla.HloSnapshot {
	if m != nil {
		return m.HloSnapshot
	}
	return nil
}

// Literal to allocate space for, and transfer to, device memory.
type XLAAllocation struct {
	Value                *xla.LiteralProto `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *XLAAllocation) Reset()         { *m = XLAAllocation{} }
func (m *XLAAllocation) String() string { return proto.CompactTextString(m) }
func (*XLAAllocation) ProtoMessage()    {}
func (*XLAAllocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef605a9242ff614, []int{3}
}

func (m *XLAAllocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XLAAllocation.Unmarshal(m, b)
}
func (m *XLAAllocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XLAAllocation.Marshal(b, m, deterministic)
}
func (m *XLAAllocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XLAAllocation.Merge(m, src)
}
func (m *XLAAllocation) XXX_Size() int {
	return xxx_messageInfo_XLAAllocation.Size(m)
}
func (m *XLAAllocation) XXX_DiscardUnknown() {
	xxx_messageInfo_XLAAllocation.DiscardUnknown(m)
}

var xxx_messageInfo_XLAAllocation proto.InternalMessageInfo

func (m *XLAAllocation) GetValue() *xla.LiteralProto {
	if m != nil {
		return m.Value
	}
	return nil
}

// Node in a tree describing a tuple constructed from input handles. A
// node is an internal node if tuples is non-empty, in which case
// input_index and release_input_handle are ignored. Otherwise a node
// is a leaf node. Each leaf XLATupleNode is the index of an input
// which corresponds to a handle that will be grafted onto the output
// tuple at that location. If release_input_handle is true that input
// handle will be released and become invalid.  Inputs may be repeated
// in which case leaves of the output tuple will alias. If an input is
// repeated, release_input_handle must be false for every leaf where
// that input appears.
//
// For example, if input 0 has shape {} and input 1 has shape {2,3}
// then the XLATupleNode with structure {1,{0,1}} corresponds to a
// tuple with shape {{2,3},{{},{2,3}}}.
type XLATupleNode struct {
	InputIndex           int32           `protobuf:"varint,1,opt,name=input_index,json=inputIndex,proto3" json:"input_index,omitempty"`
	ReleaseInputHandle   bool            `protobuf:"varint,2,opt,name=release_input_handle,json=releaseInputHandle,proto3" json:"release_input_handle,omitempty"`
	Tuples               []*XLATupleNode `protobuf:"bytes,3,rep,name=tuples,proto3" json:"tuples,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *XLATupleNode) Reset()         { *m = XLATupleNode{} }
func (m *XLATupleNode) String() string { return proto.CompactTextString(m) }
func (*XLATupleNode) ProtoMessage()    {}
func (*XLATupleNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef605a9242ff614, []int{4}
}

func (m *XLATupleNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XLATupleNode.Unmarshal(m, b)
}
func (m *XLATupleNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XLATupleNode.Marshal(b, m, deterministic)
}
func (m *XLATupleNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XLATupleNode.Merge(m, src)
}
func (m *XLATupleNode) XXX_Size() int {
	return xxx_messageInfo_XLATupleNode.Size(m)
}
func (m *XLATupleNode) XXX_DiscardUnknown() {
	xxx_messageInfo_XLATupleNode.DiscardUnknown(m)
}

var xxx_messageInfo_XLATupleNode proto.InternalMessageInfo

func (m *XLATupleNode) GetInputIndex() int32 {
	if m != nil {
		return m.InputIndex
	}
	return 0
}

func (m *XLATupleNode) GetReleaseInputHandle() bool {
	if m != nil {
		return m.ReleaseInputHandle
	}
	return false
}

func (m *XLATupleNode) GetTuples() []*XLATupleNode {
	if m != nil {
		return m.Tuples
	}
	return nil
}

// Options for an XLA execution.
type XRTExecutionConfig struct {
	// Local device to run on. This is present because the execute Op
	// may be placed on a device such as CPU or TPU_SYSTEM that
	// logically manages multiple cores.
	DeviceOrdinal int32 `protobuf:"varint,1,opt,name=device_ordinal,json=deviceOrdinal,proto3" json:"device_ordinal,omitempty"`
	// Which model-parallel computation to run from the compiled bundle.
	CoreIndexInReplica int32 `protobuf:"varint,2,opt,name=core_index_in_replica,json=coreIndexInReplica,proto3" json:"core_index_in_replica,omitempty"`
	// Optional key to disambiguate between executions. This is only
	// needed if multiple host send/recvs may be outstanding
	// concurrently with executions.
	ExecutionInstanceKey string `protobuf:"bytes,3,opt,name=execution_instance_key,json=executionInstanceKey,proto3" json:"execution_instance_key,omitempty"`
	// If non-zero, rng_seed to reset the core with.
	RngSeed uint32 `protobuf:"varint,4,opt,name=rng_seed,json=rngSeed,proto3" json:"rng_seed,omitempty"`
	// If true, release allocation handles on the inputs after running.
	ReleaseInputHandles bool `protobuf:"varint,5,opt,name=release_input_handles,json=releaseInputHandles,proto3" json:"release_input_handles,omitempty"`
	// If true, release the handle to the computation after running.
	ReleaseCompilationHandle bool `protobuf:"varint,6,opt,name=release_compilation_handle,json=releaseCompilationHandle,proto3" json:"release_compilation_handle,omitempty"`
	// If set to true, and the result shape is a tuple, then instead of returning
	// a single tuple allocation the execution will return a vector of
	// allocations, one for each of the first-level elements of the result tuple.
	ReturnExplodedTuple  bool     `protobuf:"varint,7,opt,name=return_exploded_tuple,json=returnExplodedTuple,proto3" json:"return_exploded_tuple,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XRTExecutionConfig) Reset()         { *m = XRTExecutionConfig{} }
func (m *XRTExecutionConfig) String() string { return proto.CompactTextString(m) }
func (*XRTExecutionConfig) ProtoMessage()    {}
func (*XRTExecutionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef605a9242ff614, []int{5}
}

func (m *XRTExecutionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XRTExecutionConfig.Unmarshal(m, b)
}
func (m *XRTExecutionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XRTExecutionConfig.Marshal(b, m, deterministic)
}
func (m *XRTExecutionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XRTExecutionConfig.Merge(m, src)
}
func (m *XRTExecutionConfig) XXX_Size() int {
	return xxx_messageInfo_XRTExecutionConfig.Size(m)
}
func (m *XRTExecutionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_XRTExecutionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_XRTExecutionConfig proto.InternalMessageInfo

func (m *XRTExecutionConfig) GetDeviceOrdinal() int32 {
	if m != nil {
		return m.DeviceOrdinal
	}
	return 0
}

func (m *XRTExecutionConfig) GetCoreIndexInReplica() int32 {
	if m != nil {
		return m.CoreIndexInReplica
	}
	return 0
}

func (m *XRTExecutionConfig) GetExecutionInstanceKey() string {
	if m != nil {
		return m.ExecutionInstanceKey
	}
	return ""
}

func (m *XRTExecutionConfig) GetRngSeed() uint32 {
	if m != nil {
		return m.RngSeed
	}
	return 0
}

func (m *XRTExecutionConfig) GetReleaseInputHandles() bool {
	if m != nil {
		return m.ReleaseInputHandles
	}
	return false
}

func (m *XRTExecutionConfig) GetReleaseCompilationHandle() bool {
	if m != nil {
		return m.ReleaseCompilationHandle
	}
	return false
}

func (m *XRTExecutionConfig) GetReturnExplodedTuple() bool {
	if m != nil {
		return m.ReturnExplodedTuple
	}
	return false
}

type XRTChainedExecuteConfig struct {
	// If non-zero, rng_seed to reset the core with.
	RngSeed uint32 `protobuf:"varint,1,opt,name=rng_seed,json=rngSeed,proto3" json:"rng_seed,omitempty"`
	// Which model-parallel computation to run from the compiled bundle.
	CoreIndexInReplica int32 `protobuf:"varint,2,opt,name=core_index_in_replica,json=coreIndexInReplica,proto3" json:"core_index_in_replica,omitempty"`
	// Optional key to disambiguate between executions. This is only needed if
	// multiple host send/recvs may be outstanding concurrently with executions.
	ExecutionInstanceKey string   `protobuf:"bytes,3,opt,name=execution_instance_key,json=executionInstanceKey,proto3" json:"execution_instance_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XRTChainedExecuteConfig) Reset()         { *m = XRTChainedExecuteConfig{} }
func (m *XRTChainedExecuteConfig) String() string { return proto.CompactTextString(m) }
func (*XRTChainedExecuteConfig) ProtoMessage()    {}
func (*XRTChainedExecuteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef605a9242ff614, []int{6}
}

func (m *XRTChainedExecuteConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XRTChainedExecuteConfig.Unmarshal(m, b)
}
func (m *XRTChainedExecuteConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XRTChainedExecuteConfig.Marshal(b, m, deterministic)
}
func (m *XRTChainedExecuteConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XRTChainedExecuteConfig.Merge(m, src)
}
func (m *XRTChainedExecuteConfig) XXX_Size() int {
	return xxx_messageInfo_XRTChainedExecuteConfig.Size(m)
}
func (m *XRTChainedExecuteConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_XRTChainedExecuteConfig.DiscardUnknown(m)
}

var xxx_messageInfo_XRTChainedExecuteConfig proto.InternalMessageInfo

func (m *XRTChainedExecuteConfig) GetRngSeed() uint32 {
	if m != nil {
		return m.RngSeed
	}
	return 0
}

func (m *XRTChainedExecuteConfig) GetCoreIndexInReplica() int32 {
	if m != nil {
		return m.CoreIndexInReplica
	}
	return 0
}

func (m *XRTChainedExecuteConfig) GetExecutionInstanceKey() string {
	if m != nil {
		return m.ExecutionInstanceKey
	}
	return ""
}

// A single chained execute operation. An operation can either be a device data
// load, or an existing (as in, previously compiled and accessible via its int64
// handle) XLA computation execution.
type XRTChainedExecuteOp struct {
	// Types that are valid to be assigned to OpOneof:
	//	*XRTChainedExecuteOp_DataHandle
	//	*XRTChainedExecuteOp_ComputationHandle
	OpOneof isXRTChainedExecuteOp_OpOneof `protobuf_oneof:"op_oneof"`
	// The outputs of this XRTChainedExecuteOp operation.
	Outputs []*XRTChainedExecuteOp_Output `protobuf:"bytes,3,rep,name=outputs,proto3" json:"outputs,omitempty"`
	// The inputs of this XRTChainedExecuteOp operation. If data_handle is set,
	// there are no inputs.
	Inputs               []*XRTChainedExecuteOp_Input `protobuf:"bytes,4,rep,name=inputs,proto3" json:"inputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *XRTChainedExecuteOp) Reset()         { *m = XRTChainedExecuteOp{} }
func (m *XRTChainedExecuteOp) String() string { return proto.CompactTextString(m) }
func (*XRTChainedExecuteOp) ProtoMessage()    {}
func (*XRTChainedExecuteOp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef605a9242ff614, []int{7}
}

func (m *XRTChainedExecuteOp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XRTChainedExecuteOp.Unmarshal(m, b)
}
func (m *XRTChainedExecuteOp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XRTChainedExecuteOp.Marshal(b, m, deterministic)
}
func (m *XRTChainedExecuteOp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XRTChainedExecuteOp.Merge(m, src)
}
func (m *XRTChainedExecuteOp) XXX_Size() int {
	return xxx_messageInfo_XRTChainedExecuteOp.Size(m)
}
func (m *XRTChainedExecuteOp) XXX_DiscardUnknown() {
	xxx_messageInfo_XRTChainedExecuteOp.DiscardUnknown(m)
}

var xxx_messageInfo_XRTChainedExecuteOp proto.InternalMessageInfo

type isXRTChainedExecuteOp_OpOneof interface {
	isXRTChainedExecuteOp_OpOneof()
}

type XRTChainedExecuteOp_DataHandle struct {
	DataHandle int64 `protobuf:"varint,1,opt,name=data_handle,json=dataHandle,proto3,oneof"`
}

type XRTChainedExecuteOp_ComputationHandle struct {
	ComputationHandle int64 `protobuf:"varint,2,opt,name=computation_handle,json=computationHandle,proto3,oneof"`
}

func (*XRTChainedExecuteOp_DataHandle) isXRTChainedExecuteOp_OpOneof() {}

func (*XRTChainedExecuteOp_ComputationHandle) isXRTChainedExecuteOp_OpOneof() {}

func (m *XRTChainedExecuteOp) GetOpOneof() isXRTChainedExecuteOp_OpOneof {
	if m != nil {
		return m.OpOneof
	}
	return nil
}

func (m *XRTChainedExecuteOp) GetDataHandle() int64 {
	if x, ok := m.GetOpOneof().(*XRTChainedExecuteOp_DataHandle); ok {
		return x.DataHandle
	}
	return 0
}

func (m *XRTChainedExecuteOp) GetComputationHandle() int64 {
	if x, ok := m.GetOpOneof().(*XRTChainedExecuteOp_ComputationHandle); ok {
		return x.ComputationHandle
	}
	return 0
}

func (m *XRTChainedExecuteOp) GetOutputs() []*XRTChainedExecuteOp_Output {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *XRTChainedExecuteOp) GetInputs() []*XRTChainedExecuteOp_Input {
	if m != nil {
		return m.Inputs
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*XRTChainedExecuteOp) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*XRTChainedExecuteOp_DataHandle)(nil),
		(*XRTChainedExecuteOp_ComputationHandle)(nil),
	}
}

// Represents an input for this operation.
type XRTChainedExecuteOp_Input struct {
	// The index within the XRTChainedExecutePlan.ops post-order of the source
	// operation for this input.
	OpIndex int64 `protobuf:"varint,1,opt,name=op_index,json=opIndex,proto3" json:"op_index,omitempty"`
	// The output index of the value generated by the operation at op_index.
	// Zero (default value) means no index ({}) while if an indexing is
	// required, output_index needs to be set to index+1.
	// Thanks proto3!
	OutputIndex          int64    `protobuf:"varint,2,opt,name=output_index,json=outputIndex,proto3" json:"output_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XRTChainedExecuteOp_Input) Reset()         { *m = XRTChainedExecuteOp_Input{} }
func (m *XRTChainedExecuteOp_Input) String() string { return proto.CompactTextString(m) }
func (*XRTChainedExecuteOp_Input) ProtoMessage()    {}
func (*XRTChainedExecuteOp_Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef605a9242ff614, []int{7, 0}
}

func (m *XRTChainedExecuteOp_Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XRTChainedExecuteOp_Input.Unmarshal(m, b)
}
func (m *XRTChainedExecuteOp_Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XRTChainedExecuteOp_Input.Marshal(b, m, deterministic)
}
func (m *XRTChainedExecuteOp_Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XRTChainedExecuteOp_Input.Merge(m, src)
}
func (m *XRTChainedExecuteOp_Input) XXX_Size() int {
	return xxx_messageInfo_XRTChainedExecuteOp_Input.Size(m)
}
func (m *XRTChainedExecuteOp_Input) XXX_DiscardUnknown() {
	xxx_messageInfo_XRTChainedExecuteOp_Input.DiscardUnknown(m)
}

var xxx_messageInfo_XRTChainedExecuteOp_Input proto.InternalMessageInfo

func (m *XRTChainedExecuteOp_Input) GetOpIndex() int64 {
	if m != nil {
		return m.OpIndex
	}
	return 0
}

func (m *XRTChainedExecuteOp_Input) GetOutputIndex() int64 {
	if m != nil {
		return m.OutputIndex
	}
	return 0
}

// Represents an output of the XRTChainedExecute operation, which should
// originate by the output of this operation.
type XRTChainedExecuteOp_Output struct {
	// The index in the value generated by this operation, which should be
	// forwarded as XRTChainedExecute output. If output_index is zero (default
	// value) the whole output will be used as result. This means that if the
	// output shape is a tuple, the result will be the full tuple. Otherwise the
	// real sub-tuple index will be output_index - 1.
	OutputIndex int64 `protobuf:"varint,1,opt,name=output_index,json=outputIndex,proto3" json:"output_index,omitempty"`
	// The index in the vector of the results returned by the XRTChainedExecute
	// operation, where this output should be forwarded.
	ResultIndex          int64    `protobuf:"varint,2,opt,name=result_index,json=resultIndex,proto3" json:"result_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XRTChainedExecuteOp_Output) Reset()         { *m = XRTChainedExecuteOp_Output{} }
func (m *XRTChainedExecuteOp_Output) String() string { return proto.CompactTextString(m) }
func (*XRTChainedExecuteOp_Output) ProtoMessage()    {}
func (*XRTChainedExecuteOp_Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef605a9242ff614, []int{7, 1}
}

func (m *XRTChainedExecuteOp_Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XRTChainedExecuteOp_Output.Unmarshal(m, b)
}
func (m *XRTChainedExecuteOp_Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XRTChainedExecuteOp_Output.Marshal(b, m, deterministic)
}
func (m *XRTChainedExecuteOp_Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XRTChainedExecuteOp_Output.Merge(m, src)
}
func (m *XRTChainedExecuteOp_Output) XXX_Size() int {
	return xxx_messageInfo_XRTChainedExecuteOp_Output.Size(m)
}
func (m *XRTChainedExecuteOp_Output) XXX_DiscardUnknown() {
	xxx_messageInfo_XRTChainedExecuteOp_Output.DiscardUnknown(m)
}

var xxx_messageInfo_XRTChainedExecuteOp_Output proto.InternalMessageInfo

func (m *XRTChainedExecuteOp_Output) GetOutputIndex() int64 {
	if m != nil {
		return m.OutputIndex
	}
	return 0
}

func (m *XRTChainedExecuteOp_Output) GetResultIndex() int64 {
	if m != nil {
		return m.ResultIndex
	}
	return 0
}

// Execution plan for the XRTChainedExecute operation.
type XRTChainedExecutePlan struct {
	// The post order with the XRT computations to be executed.
	Ops                  []*XRTChainedExecuteOp `protobuf:"bytes,1,rep,name=ops,proto3" json:"ops,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *XRTChainedExecutePlan) Reset()         { *m = XRTChainedExecutePlan{} }
func (m *XRTChainedExecutePlan) String() string { return proto.CompactTextString(m) }
func (*XRTChainedExecutePlan) ProtoMessage()    {}
func (*XRTChainedExecutePlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef605a9242ff614, []int{8}
}

func (m *XRTChainedExecutePlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XRTChainedExecutePlan.Unmarshal(m, b)
}
func (m *XRTChainedExecutePlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XRTChainedExecutePlan.Marshal(b, m, deterministic)
}
func (m *XRTChainedExecutePlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XRTChainedExecutePlan.Merge(m, src)
}
func (m *XRTChainedExecutePlan) XXX_Size() int {
	return xxx_messageInfo_XRTChainedExecutePlan.Size(m)
}
func (m *XRTChainedExecutePlan) XXX_DiscardUnknown() {
	xxx_messageInfo_XRTChainedExecutePlan.DiscardUnknown(m)
}

var xxx_messageInfo_XRTChainedExecutePlan proto.InternalMessageInfo

func (m *XRTChainedExecutePlan) GetOps() []*XRTChainedExecuteOp {
	if m != nil {
		return m.Ops
	}
	return nil
}

func init() {
	proto.RegisterType((*DeviceAssignment)(nil), "xrt.DeviceAssignment")
	proto.RegisterType((*DeviceAssignment_ComputationDevice)(nil), "xrt.DeviceAssignment.ComputationDevice")
	proto.RegisterType((*DeviceAssignment_ComputationDevice_DeviceMeshCoordinates)(nil), "xrt.DeviceAssignment.ComputationDevice.DeviceMeshCoordinates")
	proto.RegisterType((*XLAComputationConfig)(nil), "xrt.XLAComputationConfig")
	proto.RegisterType((*XLAComputation)(nil), "xrt.XLAComputation")
	proto.RegisterType((*XLAAllocation)(nil), "xrt.XLAAllocation")
	proto.RegisterType((*XLATupleNode)(nil), "xrt.XLATupleNode")
	proto.RegisterType((*XRTExecutionConfig)(nil), "xrt.XRTExecutionConfig")
	proto.RegisterType((*XRTChainedExecuteConfig)(nil), "xrt.XRTChainedExecuteConfig")
	proto.RegisterType((*XRTChainedExecuteOp)(nil), "xrt.XRTChainedExecuteOp")
	proto.RegisterType((*XRTChainedExecuteOp_Input)(nil), "xrt.XRTChainedExecuteOp.Input")
	proto.RegisterType((*XRTChainedExecuteOp_Output)(nil), "xrt.XRTChainedExecuteOp.Output")
	proto.RegisterType((*XRTChainedExecutePlan)(nil), "xrt.XRTChainedExecutePlan")
}

func init() { proto.RegisterFile("tensorflow/compiler/xrt/xrt.proto", fileDescriptor_5ef605a9242ff614) }

var fileDescriptor_5ef605a9242ff614 = []byte{
	// 987 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x4e, 0x1b, 0x47,
	0x14, 0xae, 0x6d, 0x6c, 0xe8, 0xb1, 0xa1, 0x30, 0x60, 0xba, 0xf1, 0x45, 0x13, 0x2c, 0x35, 0x21,
	0x95, 0x6a, 0x37, 0x4e, 0x85, 0x54, 0xf5, 0x47, 0x05, 0x07, 0x09, 0x1a, 0x12, 0xd0, 0xc0, 0x85,
	0x95, 0x9b, 0xd5, 0xe2, 0x3d, 0x78, 0x57, 0x1d, 0xcf, 0xac, 0x66, 0x66, 0x53, 0xe7, 0x1d, 0xfa,
	0x02, 0xbd, 0xe8, 0x0b, 0xf4, 0x61, 0xfa, 0x16, 0x7d, 0x8e, 0x56, 0xf3, 0xb3, 0xb0, 0xd8, 0xa6,
	0xea, 0x55, 0x2f, 0x90, 0xd8, 0xf3, 0x7d, 0xe7, 0xf7, 0x3b, 0x67, 0xbd, 0xb0, 0xa7, 0x91, 0x2b,
	0x21, 0x6f, 0x98, 0xf8, 0xa5, 0x3f, 0x16, 0xd3, 0x2c, 0x65, 0x28, 0xfb, 0x33, 0xa9, 0xcd, 0x5f,
	0x2f, 0x93, 0x42, 0x0b, 0x52, 0x9b, 0x49, 0xdd, 0x39, 0x58, 0xc6, 0xd3, 0x37, 0x83, 0x19, 0x8b,
	0xfa, 0x89, 0x50, 0x3a, 0x34, 0xc6, 0x5c, 0x63, 0x38, 0x45, 0x1d, 0xc5, 0x91, 0x8e, 0x9c, 0x73,
	0xe7, 0xf9, 0xd2, 0xf8, 0x2c, 0xea, 0x2b, 0x94, 0xef, 0xd3, 0x31, 0xf6, 0x13, 0x26, 0x3c, 0x75,
	0xef, 0x21, 0xea, 0x8c, 0x15, 0xd1, 0x9e, 0xfe, 0x0b, 0x25, 0xbc, 0xcb, 0xda, 0xfd, 0xad, 0x0a,
	0x9b, 0xaf, 0xd0, 0xc4, 0x3f, 0x54, 0x2a, 0x9d, 0xf0, 0x29, 0x72, 0x4d, 0x46, 0xb0, 0xed, 0x8a,
	0x8c, 0x74, 0x2a, 0x78, 0x18, 0x5b, 0x5c, 0x05, 0x95, 0x27, 0xb5, 0xfd, 0xe6, 0xe0, 0x59, 0xcf,
	0x34, 0x3c, 0xef, 0xd3, 0x1b, 0xde, 0x39, 0x38, 0x8c, 0x92, 0xf1, 0xbc, 0x49, 0x75, 0xfe, 0xa8,
	0xc0, 0xd6, 0x02, 0x93, 0xdc, 0xc0, 0x27, 0x12, 0x33, 0x96, 0x8e, 0xa3, 0xb9, 0x5c, 0xdf, 0xff,
	0xc7, 0x5c, 0x9e, 0xf2, 0x06, 0x55, 0x32, 0x14, 0x42, 0xc6, 0x29, 0x8f, 0x34, 0x2a, 0xba, 0xe1,
	0xa3, 0x16, 0xd9, 0xbf, 0x84, 0xf6, 0x52, 0x22, 0xd9, 0x81, 0xfa, 0xfb, 0x88, 0xe5, 0x68, 0xd3,
	0xd6, 0xa9, 0x7b, 0xe8, 0xfe, 0x59, 0x83, 0x9d, 0xd1, 0xd9, 0x61, 0x29, 0xdb, 0x50, 0xf0, 0x9b,
	0x74, 0x42, 0xf6, 0xa0, 0xc5, 0xf3, 0x69, 0xe8, 0xa3, 0x9b, 0x62, 0x2b, 0xfb, 0x75, 0xda, 0xe4,
	0xf9, 0x94, 0x7a, 0x13, 0x79, 0x01, 0x6d, 0x43, 0x19, 0x0b, 0x89, 0x2a, 0xcc, 0x50, 0x16, 0xe4,
	0xa0, 0x6a, 0xb9, 0x84, 0xe7, 0xd3, 0xa1, 0xc1, 0x2e, 0x50, 0x7a, 0x1f, 0xf2, 0x0e, 0xda, 0x4b,
	0xf7, 0x23, 0xa8, 0x3d, 0xa9, 0xec, 0x37, 0x07, 0x4f, 0x7b, 0x77, 0x92, 0xf6, 0xdc, 0x3e, 0xf5,
	0x4e, 0x84, 0xd2, 0xae, 0x3c, 0x7c, 0xe3, 0xd9, 0x74, 0x3b, 0x59, 0x34, 0x92, 0x6f, 0x61, 0x3d,
	0x93, 0x62, 0x22, 0xa3, 0x69, 0xa8, 0x92, 0x28, 0xc3, 0x60, 0xc5, 0xc6, 0xdc, 0xed, 0x99, 0x28,
	0x17, 0x0e, 0xb9, 0x34, 0xc0, 0x85, 0xd9, 0x0a, 0xda, 0xca, 0x4a, 0x26, 0xf2, 0x1a, 0x76, 0x4d,
	0x07, 0xa6, 0x97, 0xf0, 0x7e, 0x94, 0xba, 0x55, 0xe9, 0xa1, 0x28, 0xdb, 0x19, 0x4a, 0xd3, 0x64,
	0x19, 0x21, 0x47, 0xb0, 0xe5, 0x34, 0x0e, 0xa3, 0x5b, 0x41, 0x83, 0x86, 0xad, 0xa6, 0xbd, 0x54,
	0x6d, 0xba, 0x19, 0xcf, 0xef, 0xe7, 0x01, 0xac, 0xc7, 0x78, 0x9d, 0x4f, 0x42, 0x91, 0x19, 0x55,
	0x54, 0xb0, 0x6a, 0xfd, 0xb7, 0x6c, 0x1d, 0xaf, 0x0c, 0x72, 0xee, 0x00, 0xda, 0x8a, 0x4b, 0x4f,
	0xdd, 0x19, 0x6c, 0xdc, 0xd7, 0x93, 0xbc, 0x80, 0xc6, 0xd8, 0x6a, 0x6a, 0x35, 0x6c, 0x0e, 0x1e,
	0xd9, 0x12, 0x96, 0x89, 0x4e, 0x3d, 0x91, 0xbc, 0x84, 0x56, 0xc2, 0x44, 0xa8, 0x78, 0x94, 0xa9,
	0x44, 0x68, 0x2b, 0x68, 0x73, 0xb0, 0x69, 0x73, 0x9f, 0x30, 0x71, 0xe9, 0xed, 0xb4, 0x99, 0xdc,
	0x3d, 0x74, 0x7f, 0x80, 0xf5, 0xd1, 0xd9, 0xe1, 0x21, 0x63, 0x62, 0xec, 0x12, 0x3f, 0x2b, 0x36,
	0xae, 0x5a, 0x2a, 0xfd, 0x2c, 0xd5, 0x28, 0x23, 0xe6, 0xa6, 0xe7, 0xf0, 0x9f, 0x56, 0xd6, 0x2a,
	0x9b, 0xd5, 0xee, 0xaf, 0x15, 0x68, 0x8d, 0xce, 0x0e, 0xaf, 0xf2, 0x8c, 0xe1, 0x5b, 0x11, 0x23,
	0x79, 0x0c, 0xcd, 0x94, 0x67, 0xb9, 0x0e, 0x53, 0x1e, 0xe3, 0xcc, 0x6f, 0x20, 0x58, 0xd3, 0xa9,
	0xb1, 0x90, 0xaf, 0x60, 0x47, 0x22, 0xc3, 0x48, 0x61, 0xe8, 0x88, 0x49, 0xc4, 0x63, 0xe6, 0xf2,
	0xad, 0x51, 0xe2, 0xb1, 0x53, 0x03, 0x9d, 0x58, 0x84, 0x3c, 0x87, 0x86, 0x36, 0xf1, 0x55, 0x50,
	0xb3, 0xb2, 0x6e, 0x15, 0xb3, 0xb8, 0xcd, 0x4a, 0x3d, 0xa1, 0xfb, 0x57, 0x15, 0xc8, 0x88, 0x5e,
	0x1d, 0xcf, 0x70, 0x9c, 0x97, 0xee, 0xe2, 0x73, 0xd8, 0xf0, 0xda, 0xba, 0xd3, 0x62, 0xbe, 0xae,
	0x75, 0x67, 0x3d, 0x77, 0x46, 0x73, 0x1b, 0x76, 0x97, 0x6c, 0xe9, 0x61, 0xca, 0xe7, 0x6f, 0xc3,
	0x80, 0xb6, 0x89, 0x53, 0x5e, 0xdc, 0xc6, 0xd7, 0xb0, 0x8b, 0x45, 0xb2, 0x30, 0xe5, 0x4a, 0x47,
	0x7c, 0x8c, 0xe1, 0xcf, 0xf8, 0xc1, 0x1e, 0xc7, 0xc7, 0x74, 0xe7, 0x16, 0x3d, 0xf5, 0xe0, 0x6b,
	0xfc, 0x40, 0x1e, 0xc1, 0x9a, 0xe4, 0x93, 0x50, 0x21, 0xc6, 0x76, 0xe1, 0xd7, 0xe9, 0xaa, 0xe4,
	0x93, 0x4b, 0xc4, 0x98, 0x0c, 0xa0, 0xbd, 0x6c, 0x3c, 0x2a, 0xa8, 0xdb, 0xf9, 0x6c, 0x2f, 0xce,
	0x47, 0x91, 0xef, 0xa0, 0x53, 0xf8, 0xb8, 0x57, 0xaa, 0x7b, 0x3d, 0xfa, 0xc1, 0x36, 0xac, 0x63,
	0xe0, 0x19, 0xc3, 0x3b, 0x82, 0x1f, 0xaf, 0xcd, 0xa8, 0x73, 0xc9, 0x43, 0x9c, 0x65, 0x4c, 0xc4,
	0x18, 0x87, 0x76, 0x9a, 0x76, 0x79, 0x6d, 0x46, 0x03, 0x1e, 0x7b, 0xcc, 0xce, 0xbc, 0xfb, 0x7b,
	0x05, 0x3e, 0x1d, 0xd1, 0xab, 0x61, 0x12, 0xa5, 0x1c, 0x63, 0x37, 0x6e, 0xf4, 0xc3, 0x2e, 0x37,
	0x57, 0xb9, 0xdf, 0xdc, 0xff, 0x35, 0xe0, 0xee, 0xdf, 0x55, 0xd8, 0x5e, 0xa8, 0xef, 0x3c, 0x23,
	0x7b, 0xd0, 0x34, 0xaf, 0x9d, 0x62, 0x34, 0xa6, 0xbc, 0xda, 0xc9, 0x47, 0x14, 0x8c, 0xd1, 0x8f,
	0xa3, 0x0f, 0xe5, 0xdf, 0x87, 0xf2, 0x76, 0x1a, 0xe6, 0x56, 0x09, 0xf3, 0x0e, 0xdf, 0xc0, 0xaa,
	0xc8, 0x75, 0x96, 0xeb, 0x62, 0x3f, 0x1f, 0xbb, 0xfd, 0x5c, 0x4c, 0xdf, 0x3b, 0xb7, 0x3c, 0x5a,
	0xf0, 0xc9, 0x01, 0x34, 0xac, 0xc8, 0x2a, 0x58, 0xb1, 0x9e, 0x9f, 0x3d, 0xe8, 0x69, 0xf5, 0xa6,
	0x9e, 0xdd, 0x39, 0x86, 0xba, 0x35, 0x98, 0x59, 0x8b, 0xac, 0x74, 0x6a, 0x35, 0xba, 0x2a, 0x32,
	0x77, 0x67, 0x7b, 0xd0, 0x72, 0x69, 0x3c, 0x6c, 0x3b, 0xa0, 0x4d, 0x67, 0xb3, 0x94, 0xce, 0x5b,
	0x68, 0xb8, 0x8a, 0x16, 0xc8, 0x95, 0x05, 0xb2, 0xa1, 0x48, 0x54, 0x39, 0x9b, 0x8b, 0xe7, 0x6c,
	0x96, 0x72, 0x04, 0xb6, 0x1a, 0xc1, 0x51, 0xdc, 0x74, 0x87, 0xd0, 0x5e, 0xe8, 0xe3, 0x82, 0x45,
	0x9c, 0x7c, 0x01, 0x35, 0x91, 0x15, 0xbf, 0xa3, 0xc1, 0x43, 0x0d, 0x53, 0x43, 0x3a, 0x3a, 0x7a,
	0xf7, 0xe3, 0x24, 0xd5, 0x49, 0x7e, 0xdd, 0x1b, 0x8b, 0x69, 0xbf, 0xf4, 0xe5, 0xb0, 0xfc, 0xdf,
	0x89, 0xe8, 0x9b, 0x0d, 0xba, 0xf7, 0x15, 0x74, 0xdd, 0xb0, 0xdf, 0x13, 0x2f, 0xff, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x97, 0x0c, 0xbb, 0x84, 0x27, 0x09, 0x00, 0x00,
}
