// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/trace_events.proto

package protobuf

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

// A 'Trace' contains metadata for the individual traces of a system.
type Trace struct {
	// The devices that this trace has information about. Maps from device_id to
	// more data about the specific device.
	Devices map[uint32]*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// All trace events capturing in the profiling period.
	TraceEvents          []*TraceEvent `protobuf:"bytes,4,rep,name=trace_events,json=traceEvents,proto3" json:"trace_events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Trace) Reset()         { *m = Trace{} }
func (m *Trace) String() string { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()    {}
func (*Trace) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a93bbdc8c2bd3ea, []int{0}
}

func (m *Trace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trace.Unmarshal(m, b)
}
func (m *Trace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trace.Marshal(b, m, deterministic)
}
func (m *Trace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trace.Merge(m, src)
}
func (m *Trace) XXX_Size() int {
	return xxx_messageInfo_Trace.Size(m)
}
func (m *Trace) XXX_DiscardUnknown() {
	xxx_messageInfo_Trace.DiscardUnknown(m)
}

var xxx_messageInfo_Trace proto.InternalMessageInfo

func (m *Trace) GetDevices() map[uint32]*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *Trace) GetTraceEvents() []*TraceEvent {
	if m != nil {
		return m.TraceEvents
	}
	return nil
}

// A 'device' is a physical entity in the system and is comprised of several
// resources.
type Device struct {
	// The name of the device.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The id of this device, unique in a single trace.
	DeviceId uint32 `protobuf:"varint,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// The resources on this device, keyed by resource_id;
	Resources            map[uint32]*Resource `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a93bbdc8c2bd3ea, []int{1}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetDeviceId() uint32 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *Device) GetResources() map[uint32]*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

// A 'resource' generally is a specific computation component on a device. These
// can range from threads on CPUs to specific arithmetic units on hardware
// devices.
type Resource struct {
	// The name of the resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The id of the resource. Unique within a device.
	ResourceId           uint32   `protobuf:"varint,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a93bbdc8c2bd3ea, []int{2}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetResourceId() uint32 {
	if m != nil {
		return m.ResourceId
	}
	return 0
}

type TraceEvent struct {
	// The id of the device that this event occurred on. The full dataset should
	// have this device present in the Trace object.
	DeviceId uint32 `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// The id of the resource that this event occurred on. The full dataset should
	// have this resource present in the Device object of the Trace object. A
	// resource_id is unique on a specific device, but not necessarily within the
	// trace.
	ResourceId uint32 `protobuf:"varint,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// The name of this trace event.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The timestamp that this event occurred at (in picos since tracing started).
	TimestampPs uint64 `protobuf:"varint,9,opt,name=timestamp_ps,json=timestampPs,proto3" json:"timestamp_ps,omitempty"`
	// The duration of the event in picoseconds if applicable.
	// Events without duration are called instant events.
	DurationPs uint64 `protobuf:"varint,10,opt,name=duration_ps,json=durationPs,proto3" json:"duration_ps,omitempty"`
	// Extra arguments that will be displayed in trace view.
	Args                 map[string]string `protobuf:"bytes,11,rep,name=args,proto3" json:"args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TraceEvent) Reset()         { *m = TraceEvent{} }
func (m *TraceEvent) String() string { return proto.CompactTextString(m) }
func (*TraceEvent) ProtoMessage()    {}
func (*TraceEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a93bbdc8c2bd3ea, []int{3}
}

func (m *TraceEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceEvent.Unmarshal(m, b)
}
func (m *TraceEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceEvent.Marshal(b, m, deterministic)
}
func (m *TraceEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceEvent.Merge(m, src)
}
func (m *TraceEvent) XXX_Size() int {
	return xxx_messageInfo_TraceEvent.Size(m)
}
func (m *TraceEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TraceEvent proto.InternalMessageInfo

func (m *TraceEvent) GetDeviceId() uint32 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *TraceEvent) GetResourceId() uint32 {
	if m != nil {
		return m.ResourceId
	}
	return 0
}

func (m *TraceEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TraceEvent) GetTimestampPs() uint64 {
	if m != nil {
		return m.TimestampPs
	}
	return 0
}

func (m *TraceEvent) GetDurationPs() uint64 {
	if m != nil {
		return m.DurationPs
	}
	return 0
}

func (m *TraceEvent) GetArgs() map[string]string {
	if m != nil {
		return m.Args
	}
	return nil
}

func init() {
	proto.RegisterType((*Trace)(nil), "tensorflow.profiler.Trace")
	proto.RegisterMapType((map[uint32]*Device)(nil), "tensorflow.profiler.Trace.DevicesEntry")
	proto.RegisterType((*Device)(nil), "tensorflow.profiler.Device")
	proto.RegisterMapType((map[uint32]*Resource)(nil), "tensorflow.profiler.Device.ResourcesEntry")
	proto.RegisterType((*Resource)(nil), "tensorflow.profiler.Resource")
	proto.RegisterType((*TraceEvent)(nil), "tensorflow.profiler.TraceEvent")
	proto.RegisterMapType((map[string]string)(nil), "tensorflow.profiler.TraceEvent.ArgsEntry")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/trace_events.proto", fileDescriptor_7a93bbdc8c2bd3ea)
}

var fileDescriptor_7a93bbdc8c2bd3ea = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xd5, 0x26, 0x69, 0xa9, 0xc7, 0x29, 0x82, 0x85, 0x83, 0x95, 0x0a, 0x35, 0xe4, 0x42, 0x00,
	0xc9, 0x16, 0xed, 0x01, 0x84, 0x40, 0xa8, 0x11, 0x95, 0xe8, 0xcd, 0x5a, 0x21, 0x21, 0xc1, 0x21,
	0x72, 0xec, 0xb5, 0xb1, 0x1a, 0x7b, 0xad, 0xdd, 0x75, 0xaa, 0xfe, 0x0a, 0x5f, 0xc7, 0x8d, 0x5f,
	0xe0, 0x88, 0x76, 0x17, 0xdb, 0x6b, 0x64, 0x9a, 0xdb, 0xe4, 0xe5, 0xcd, 0x9b, 0x7d, 0x6f, 0x3c,
	0xf0, 0x52, 0xd2, 0x52, 0x30, 0x9e, 0x6e, 0xd9, 0x4d, 0x10, 0x33, 0x4e, 0x83, 0x8a, 0x33, 0xc9,
	0x36, 0x75, 0x1a, 0x48, 0x1e, 0xc5, 0x74, 0x4d, 0x77, 0xb4, 0x94, 0xc2, 0xd7, 0x28, 0x7e, 0xd4,
	0x91, 0x15, 0x92, 0xe6, 0x5b, 0xca, 0x17, 0xbf, 0x10, 0x1c, 0x7c, 0x56, 0x5c, 0x7c, 0x01, 0xf7,
	0x12, 0xba, 0xcb, 0x63, 0x2a, 0x3c, 0x34, 0x1f, 0x2f, 0xdd, 0xb3, 0x67, 0xfe, 0x40, 0x83, 0xaf,
	0xc9, 0xfe, 0x47, 0xc3, 0xbc, 0x2c, 0x25, 0xbf, 0x25, 0x4d, 0x1f, 0x5e, 0xc1, 0xd4, 0x9e, 0xeb,
	0x4d, 0xb4, 0xce, 0xe9, 0xff, 0x75, 0x2e, 0x15, 0x8f, 0xb8, 0xb2, 0xad, 0xc5, 0xec, 0x0b, 0x4c,
	0x6d, 0x71, 0xfc, 0x00, 0xc6, 0xd7, 0xf4, 0xd6, 0x43, 0x73, 0xb4, 0x3c, 0x26, 0xaa, 0xc4, 0xaf,
	0xe0, 0x60, 0x17, 0x6d, 0x6b, 0xea, 0x8d, 0xe6, 0x68, 0xe9, 0x9e, 0x9d, 0x0c, 0xca, 0x1b, 0x0d,
	0x62, 0x98, 0x6f, 0x47, 0x6f, 0xd0, 0xe2, 0x27, 0x82, 0x43, 0x83, 0x62, 0x0c, 0x93, 0x32, 0x2a,
	0xa8, 0x16, 0x75, 0x88, 0xae, 0xf1, 0x09, 0x38, 0xc6, 0xc6, 0x3a, 0x4f, 0xb4, 0xf2, 0x31, 0x39,
	0x32, 0xc0, 0x55, 0x82, 0x3f, 0x81, 0xc3, 0xa9, 0x60, 0x35, 0x57, 0xe9, 0x8c, 0xb5, 0xab, 0x17,
	0x77, 0x8c, 0xf5, 0x49, 0x43, 0x36, 0x01, 0x75, 0xcd, 0xb3, 0x6f, 0x70, 0xbf, 0xff, 0xe7, 0x80,
	0xc1, 0xf3, 0xbe, 0xc1, 0x27, 0x83, 0x93, 0x1a, 0x15, 0xdb, 0xe2, 0x07, 0x38, 0x6a, 0xe0, 0x41,
	0x8f, 0xa7, 0xe0, 0x36, 0x2f, 0xe9, 0x5c, 0x42, 0x03, 0x5d, 0x25, 0x8b, 0x1f, 0x23, 0x80, 0x6e,
	0x31, 0xfd, 0x4c, 0xd0, 0x3f, 0x99, 0xec, 0x13, 0x6b, 0x5f, 0x30, 0xb6, 0x5e, 0xf0, 0x14, 0xa6,
	0x32, 0x2f, 0xa8, 0x90, 0x51, 0x51, 0xad, 0x2b, 0xe1, 0x39, 0x73, 0xb4, 0x9c, 0x10, 0xb7, 0xc5,
	0x42, 0xa1, 0x74, 0x93, 0x9a, 0x47, 0x32, 0x67, 0xa5, 0x62, 0x80, 0x66, 0x40, 0x03, 0x85, 0x02,
	0xbf, 0x87, 0x49, 0xc4, 0x33, 0xe1, 0xb9, 0x7a, 0x0f, 0xcf, 0xf7, 0x7c, 0x5d, 0xfe, 0x05, 0xcf,
	0xfe, 0xae, 0x41, 0xb7, 0xcd, 0x5e, 0x83, 0xd3, 0x42, 0x76, 0xf8, 0x8e, 0x09, 0xff, 0xb1, 0x1d,
	0xbe, 0x63, 0xa5, 0xbb, 0xaa, 0xc0, 0x63, 0x3c, 0xb3, 0xc7, 0xa5, 0x3c, 0x2a, 0xe8, 0x0d, 0xe3,
	0xd7, 0xab, 0x87, 0xdd, 0x40, 0x11, 0xaa, 0x6b, 0x13, 0x21, 0xfa, 0xfa, 0x2e, 0xcb, 0xe5, 0xf7,
	0x7a, 0xe3, 0xc7, 0xac, 0x08, 0xac, 0x43, 0x1d, 0x2e, 0x33, 0xd6, 0xbf, 0xe0, 0xdf, 0x08, 0x6d,
	0x0e, 0xf5, 0x8f, 0xf3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xfc, 0x14, 0xed, 0xe7, 0x03,
	0x00, 0x00,
}
