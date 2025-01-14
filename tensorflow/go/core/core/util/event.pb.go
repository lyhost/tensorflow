// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/util/event.proto

package util

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

// Current health status of a worker.
type WorkerHealth int32

const (
	WorkerHealth_OK                       WorkerHealth = 0
	WorkerHealth_RECEIVED_SHUTDOWN_SIGNAL WorkerHealth = 1
	WorkerHealth_INTERNAL_ERROR           WorkerHealth = 2
	WorkerHealth_SHUTTING_DOWN            WorkerHealth = 3
)

var WorkerHealth_name = map[int32]string{
	0: "OK",
	1: "RECEIVED_SHUTDOWN_SIGNAL",
	2: "INTERNAL_ERROR",
	3: "SHUTTING_DOWN",
}

var WorkerHealth_value = map[string]int32{
	"OK":                       0,
	"RECEIVED_SHUTDOWN_SIGNAL": 1,
	"INTERNAL_ERROR":           2,
	"SHUTTING_DOWN":            3,
}

func (x WorkerHealth) String() string {
	return proto.EnumName(WorkerHealth_name, int32(x))
}

func (WorkerHealth) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c32e35f2f2971c10, []int{0}
}

// Indicates the behavior of the worker when an internal error or shutdown
// signal is received.
type WorkerShutdownMode int32

const (
	WorkerShutdownMode_DEFAULT                WorkerShutdownMode = 0
	WorkerShutdownMode_NOT_CONFIGURED         WorkerShutdownMode = 1
	WorkerShutdownMode_WAIT_FOR_COORDINATOR   WorkerShutdownMode = 2
	WorkerShutdownMode_SHUTDOWN_AFTER_TIMEOUT WorkerShutdownMode = 3
)

var WorkerShutdownMode_name = map[int32]string{
	0: "DEFAULT",
	1: "NOT_CONFIGURED",
	2: "WAIT_FOR_COORDINATOR",
	3: "SHUTDOWN_AFTER_TIMEOUT",
}

var WorkerShutdownMode_value = map[string]int32{
	"DEFAULT":                0,
	"NOT_CONFIGURED":         1,
	"WAIT_FOR_COORDINATOR":   2,
	"SHUTDOWN_AFTER_TIMEOUT": 3,
}

func (x WorkerShutdownMode) String() string {
	return proto.EnumName(WorkerShutdownMode_name, int32(x))
}

func (WorkerShutdownMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c32e35f2f2971c10, []int{1}
}

type LogMessage_Level int32

const (
	LogMessage_UNKNOWN LogMessage_Level = 0
	// Note: The logging level 10 cannot be named DEBUG. Some software
	// projects compile their C/C++ code with -DDEBUG in debug builds. So the
	// C++ code generated from this file should not have an identifier named
	// DEBUG.
	LogMessage_DEBUGGING LogMessage_Level = 10
	LogMessage_INFO      LogMessage_Level = 20
	LogMessage_WARN      LogMessage_Level = 30
	LogMessage_ERROR     LogMessage_Level = 40
	LogMessage_FATAL     LogMessage_Level = 50
)

var LogMessage_Level_name = map[int32]string{
	0:  "UNKNOWN",
	10: "DEBUGGING",
	20: "INFO",
	30: "WARN",
	40: "ERROR",
	50: "FATAL",
}

var LogMessage_Level_value = map[string]int32{
	"UNKNOWN":   0,
	"DEBUGGING": 10,
	"INFO":      20,
	"WARN":      30,
	"ERROR":     40,
	"FATAL":     50,
}

func (x LogMessage_Level) String() string {
	return proto.EnumName(LogMessage_Level_name, int32(x))
}

func (LogMessage_Level) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c32e35f2f2971c10, []int{1, 0}
}

type SessionLog_SessionStatus int32

const (
	SessionLog_STATUS_UNSPECIFIED SessionLog_SessionStatus = 0
	SessionLog_START              SessionLog_SessionStatus = 1
	SessionLog_STOP               SessionLog_SessionStatus = 2
	SessionLog_CHECKPOINT         SessionLog_SessionStatus = 3
)

var SessionLog_SessionStatus_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "START",
	2: "STOP",
	3: "CHECKPOINT",
}

var SessionLog_SessionStatus_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"START":              1,
	"STOP":               2,
	"CHECKPOINT":         3,
}

func (x SessionLog_SessionStatus) String() string {
	return proto.EnumName(SessionLog_SessionStatus_name, int32(x))
}

func (SessionLog_SessionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c32e35f2f2971c10, []int{2, 0}
}

// Protocol buffer representing an event that happened during
// the execution of a Brain model.
type Event struct {
	// Timestamp of the event.
	WallTime float64 `protobuf:"fixed64,1,opt,name=wall_time,json=wallTime,proto3" json:"wall_time,omitempty"`
	// Global step of the event.
	Step int64 `protobuf:"varint,2,opt,name=step,proto3" json:"step,omitempty"`
	// Types that are valid to be assigned to What:
	//	*Event_FileVersion
	//	*Event_GraphDef
	//	*Event_Summary
	//	*Event_LogMessage
	//	*Event_SessionLog
	//	*Event_TaggedRunMetadata
	//	*Event_MetaGraphDef
	What                 isEvent_What `protobuf_oneof:"what"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32e35f2f2971c10, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetWallTime() float64 {
	if m != nil {
		return m.WallTime
	}
	return 0
}

func (m *Event) GetStep() int64 {
	if m != nil {
		return m.Step
	}
	return 0
}

type isEvent_What interface {
	isEvent_What()
}

type Event_FileVersion struct {
	FileVersion string `protobuf:"bytes,3,opt,name=file_version,json=fileVersion,proto3,oneof"`
}

type Event_GraphDef struct {
	GraphDef []byte `protobuf:"bytes,4,opt,name=graph_def,json=graphDef,proto3,oneof"`
}

type Event_Summary struct {
	Summary *framework.Summary `protobuf:"bytes,5,opt,name=summary,proto3,oneof"`
}

type Event_LogMessage struct {
	LogMessage *LogMessage `protobuf:"bytes,6,opt,name=log_message,json=logMessage,proto3,oneof"`
}

type Event_SessionLog struct {
	SessionLog *SessionLog `protobuf:"bytes,7,opt,name=session_log,json=sessionLog,proto3,oneof"`
}

type Event_TaggedRunMetadata struct {
	TaggedRunMetadata *TaggedRunMetadata `protobuf:"bytes,8,opt,name=tagged_run_metadata,json=taggedRunMetadata,proto3,oneof"`
}

type Event_MetaGraphDef struct {
	MetaGraphDef []byte `protobuf:"bytes,9,opt,name=meta_graph_def,json=metaGraphDef,proto3,oneof"`
}

func (*Event_FileVersion) isEvent_What() {}

func (*Event_GraphDef) isEvent_What() {}

func (*Event_Summary) isEvent_What() {}

func (*Event_LogMessage) isEvent_What() {}

func (*Event_SessionLog) isEvent_What() {}

func (*Event_TaggedRunMetadata) isEvent_What() {}

func (*Event_MetaGraphDef) isEvent_What() {}

func (m *Event) GetWhat() isEvent_What {
	if m != nil {
		return m.What
	}
	return nil
}

func (m *Event) GetFileVersion() string {
	if x, ok := m.GetWhat().(*Event_FileVersion); ok {
		return x.FileVersion
	}
	return ""
}

func (m *Event) GetGraphDef() []byte {
	if x, ok := m.GetWhat().(*Event_GraphDef); ok {
		return x.GraphDef
	}
	return nil
}

func (m *Event) GetSummary() *framework.Summary {
	if x, ok := m.GetWhat().(*Event_Summary); ok {
		return x.Summary
	}
	return nil
}

func (m *Event) GetLogMessage() *LogMessage {
	if x, ok := m.GetWhat().(*Event_LogMessage); ok {
		return x.LogMessage
	}
	return nil
}

func (m *Event) GetSessionLog() *SessionLog {
	if x, ok := m.GetWhat().(*Event_SessionLog); ok {
		return x.SessionLog
	}
	return nil
}

func (m *Event) GetTaggedRunMetadata() *TaggedRunMetadata {
	if x, ok := m.GetWhat().(*Event_TaggedRunMetadata); ok {
		return x.TaggedRunMetadata
	}
	return nil
}

func (m *Event) GetMetaGraphDef() []byte {
	if x, ok := m.GetWhat().(*Event_MetaGraphDef); ok {
		return x.MetaGraphDef
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_FileVersion)(nil),
		(*Event_GraphDef)(nil),
		(*Event_Summary)(nil),
		(*Event_LogMessage)(nil),
		(*Event_SessionLog)(nil),
		(*Event_TaggedRunMetadata)(nil),
		(*Event_MetaGraphDef)(nil),
	}
}

// Protocol buffer used for logging messages to the events file.
type LogMessage struct {
	Level                LogMessage_Level `protobuf:"varint,1,opt,name=level,proto3,enum=tensorflow.LogMessage_Level" json:"level,omitempty"`
	Message              string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LogMessage) Reset()         { *m = LogMessage{} }
func (m *LogMessage) String() string { return proto.CompactTextString(m) }
func (*LogMessage) ProtoMessage()    {}
func (*LogMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32e35f2f2971c10, []int{1}
}

func (m *LogMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogMessage.Unmarshal(m, b)
}
func (m *LogMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogMessage.Marshal(b, m, deterministic)
}
func (m *LogMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogMessage.Merge(m, src)
}
func (m *LogMessage) XXX_Size() int {
	return xxx_messageInfo_LogMessage.Size(m)
}
func (m *LogMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LogMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LogMessage proto.InternalMessageInfo

func (m *LogMessage) GetLevel() LogMessage_Level {
	if m != nil {
		return m.Level
	}
	return LogMessage_UNKNOWN
}

func (m *LogMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Protocol buffer used for logging session state.
type SessionLog struct {
	Status SessionLog_SessionStatus `protobuf:"varint,1,opt,name=status,proto3,enum=tensorflow.SessionLog_SessionStatus" json:"status,omitempty"`
	// This checkpoint_path contains both the path and filename.
	CheckpointPath       string   `protobuf:"bytes,2,opt,name=checkpoint_path,json=checkpointPath,proto3" json:"checkpoint_path,omitempty"`
	Msg                  string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionLog) Reset()         { *m = SessionLog{} }
func (m *SessionLog) String() string { return proto.CompactTextString(m) }
func (*SessionLog) ProtoMessage()    {}
func (*SessionLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32e35f2f2971c10, []int{2}
}

func (m *SessionLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionLog.Unmarshal(m, b)
}
func (m *SessionLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionLog.Marshal(b, m, deterministic)
}
func (m *SessionLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionLog.Merge(m, src)
}
func (m *SessionLog) XXX_Size() int {
	return xxx_messageInfo_SessionLog.Size(m)
}
func (m *SessionLog) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionLog.DiscardUnknown(m)
}

var xxx_messageInfo_SessionLog proto.InternalMessageInfo

func (m *SessionLog) GetStatus() SessionLog_SessionStatus {
	if m != nil {
		return m.Status
	}
	return SessionLog_STATUS_UNSPECIFIED
}

func (m *SessionLog) GetCheckpointPath() string {
	if m != nil {
		return m.CheckpointPath
	}
	return ""
}

func (m *SessionLog) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// For logging the metadata output for a single session.run() call.
type TaggedRunMetadata struct {
	// Tag name associated with this metadata.
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	// Byte-encoded version of the `RunMetadata` proto in order to allow lazy
	// deserialization.
	RunMetadata          []byte   `protobuf:"bytes,2,opt,name=run_metadata,json=runMetadata,proto3" json:"run_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaggedRunMetadata) Reset()         { *m = TaggedRunMetadata{} }
func (m *TaggedRunMetadata) String() string { return proto.CompactTextString(m) }
func (*TaggedRunMetadata) ProtoMessage()    {}
func (*TaggedRunMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32e35f2f2971c10, []int{3}
}

func (m *TaggedRunMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaggedRunMetadata.Unmarshal(m, b)
}
func (m *TaggedRunMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaggedRunMetadata.Marshal(b, m, deterministic)
}
func (m *TaggedRunMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaggedRunMetadata.Merge(m, src)
}
func (m *TaggedRunMetadata) XXX_Size() int {
	return xxx_messageInfo_TaggedRunMetadata.Size(m)
}
func (m *TaggedRunMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TaggedRunMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TaggedRunMetadata proto.InternalMessageInfo

func (m *TaggedRunMetadata) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *TaggedRunMetadata) GetRunMetadata() []byte {
	if m != nil {
		return m.RunMetadata
	}
	return nil
}

type WatchdogConfig struct {
	TimeoutMs            int64    `protobuf:"varint,1,opt,name=timeout_ms,json=timeoutMs,proto3" json:"timeout_ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchdogConfig) Reset()         { *m = WatchdogConfig{} }
func (m *WatchdogConfig) String() string { return proto.CompactTextString(m) }
func (*WatchdogConfig) ProtoMessage()    {}
func (*WatchdogConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32e35f2f2971c10, []int{4}
}

func (m *WatchdogConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchdogConfig.Unmarshal(m, b)
}
func (m *WatchdogConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchdogConfig.Marshal(b, m, deterministic)
}
func (m *WatchdogConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchdogConfig.Merge(m, src)
}
func (m *WatchdogConfig) XXX_Size() int {
	return xxx_messageInfo_WatchdogConfig.Size(m)
}
func (m *WatchdogConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchdogConfig.DiscardUnknown(m)
}

var xxx_messageInfo_WatchdogConfig proto.InternalMessageInfo

func (m *WatchdogConfig) GetTimeoutMs() int64 {
	if m != nil {
		return m.TimeoutMs
	}
	return 0
}

type WorkerHeartbeatRequest struct {
	ShutdownMode         WorkerShutdownMode `protobuf:"varint,1,opt,name=shutdown_mode,json=shutdownMode,proto3,enum=tensorflow.WorkerShutdownMode" json:"shutdown_mode,omitempty"`
	WatchdogConfig       *WatchdogConfig    `protobuf:"bytes,2,opt,name=watchdog_config,json=watchdogConfig,proto3" json:"watchdog_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *WorkerHeartbeatRequest) Reset()         { *m = WorkerHeartbeatRequest{} }
func (m *WorkerHeartbeatRequest) String() string { return proto.CompactTextString(m) }
func (*WorkerHeartbeatRequest) ProtoMessage()    {}
func (*WorkerHeartbeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32e35f2f2971c10, []int{5}
}

func (m *WorkerHeartbeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkerHeartbeatRequest.Unmarshal(m, b)
}
func (m *WorkerHeartbeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkerHeartbeatRequest.Marshal(b, m, deterministic)
}
func (m *WorkerHeartbeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerHeartbeatRequest.Merge(m, src)
}
func (m *WorkerHeartbeatRequest) XXX_Size() int {
	return xxx_messageInfo_WorkerHeartbeatRequest.Size(m)
}
func (m *WorkerHeartbeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerHeartbeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerHeartbeatRequest proto.InternalMessageInfo

func (m *WorkerHeartbeatRequest) GetShutdownMode() WorkerShutdownMode {
	if m != nil {
		return m.ShutdownMode
	}
	return WorkerShutdownMode_DEFAULT
}

func (m *WorkerHeartbeatRequest) GetWatchdogConfig() *WatchdogConfig {
	if m != nil {
		return m.WatchdogConfig
	}
	return nil
}

type WorkerHeartbeatResponse struct {
	HealthStatus         WorkerHealth `protobuf:"varint,1,opt,name=health_status,json=healthStatus,proto3,enum=tensorflow.WorkerHealth" json:"health_status,omitempty"`
	WorkerLog            []*Event     `protobuf:"bytes,2,rep,name=worker_log,json=workerLog,proto3" json:"worker_log,omitempty"`
	Hostname             string       `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *WorkerHeartbeatResponse) Reset()         { *m = WorkerHeartbeatResponse{} }
func (m *WorkerHeartbeatResponse) String() string { return proto.CompactTextString(m) }
func (*WorkerHeartbeatResponse) ProtoMessage()    {}
func (*WorkerHeartbeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c32e35f2f2971c10, []int{6}
}

func (m *WorkerHeartbeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkerHeartbeatResponse.Unmarshal(m, b)
}
func (m *WorkerHeartbeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkerHeartbeatResponse.Marshal(b, m, deterministic)
}
func (m *WorkerHeartbeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerHeartbeatResponse.Merge(m, src)
}
func (m *WorkerHeartbeatResponse) XXX_Size() int {
	return xxx_messageInfo_WorkerHeartbeatResponse.Size(m)
}
func (m *WorkerHeartbeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerHeartbeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerHeartbeatResponse proto.InternalMessageInfo

func (m *WorkerHeartbeatResponse) GetHealthStatus() WorkerHealth {
	if m != nil {
		return m.HealthStatus
	}
	return WorkerHealth_OK
}

func (m *WorkerHeartbeatResponse) GetWorkerLog() []*Event {
	if m != nil {
		return m.WorkerLog
	}
	return nil
}

func (m *WorkerHeartbeatResponse) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func init() {
	proto.RegisterEnum("tensorflow.WorkerHealth", WorkerHealth_name, WorkerHealth_value)
	proto.RegisterEnum("tensorflow.WorkerShutdownMode", WorkerShutdownMode_name, WorkerShutdownMode_value)
	proto.RegisterEnum("tensorflow.LogMessage_Level", LogMessage_Level_name, LogMessage_Level_value)
	proto.RegisterEnum("tensorflow.SessionLog_SessionStatus", SessionLog_SessionStatus_name, SessionLog_SessionStatus_value)
	proto.RegisterType((*Event)(nil), "tensorflow.Event")
	proto.RegisterType((*LogMessage)(nil), "tensorflow.LogMessage")
	proto.RegisterType((*SessionLog)(nil), "tensorflow.SessionLog")
	proto.RegisterType((*TaggedRunMetadata)(nil), "tensorflow.TaggedRunMetadata")
	proto.RegisterType((*WatchdogConfig)(nil), "tensorflow.WatchdogConfig")
	proto.RegisterType((*WorkerHeartbeatRequest)(nil), "tensorflow.WorkerHeartbeatRequest")
	proto.RegisterType((*WorkerHeartbeatResponse)(nil), "tensorflow.WorkerHeartbeatResponse")
}

func init() { proto.RegisterFile("tensorflow/core/util/event.proto", fileDescriptor_c32e35f2f2971c10) }

var fileDescriptor_c32e35f2f2971c10 = []byte{
	// 934 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x55, 0x5d, 0x6f, 0xe2, 0x46,
	0x14, 0xc5, 0x21, 0x5f, 0x5c, 0x08, 0xeb, 0x4c, 0x56, 0xa9, 0x95, 0x6e, 0x56, 0x94, 0x56, 0x5d,
	0xb4, 0x0f, 0xa1, 0xa2, 0x4f, 0x95, 0xba, 0x0f, 0x0e, 0x18, 0x6c, 0x05, 0x6c, 0x34, 0x36, 0x8b,
	0xd4, 0x3e, 0x8c, 0x1c, 0x18, 0x6c, 0x14, 0xdb, 0x43, 0x3d, 0x43, 0x50, 0xff, 0x4e, 0x55, 0x55,
	0xea, 0xef, 0xe9, 0x9f, 0xe9, 0xe3, 0x6a, 0x6c, 0x27, 0xc0, 0x26, 0x6f, 0xf7, 0xde, 0x39, 0xe7,
	0xfa, 0xcc, 0x99, 0x3b, 0x1e, 0x68, 0x08, 0x9a, 0x70, 0x96, 0x2e, 0x22, 0xb6, 0x69, 0xcf, 0x58,
	0x4a, 0xdb, 0x6b, 0xb1, 0x8c, 0xda, 0xf4, 0x91, 0x26, 0xe2, 0x66, 0x95, 0x32, 0xc1, 0x10, 0x6c,
	0x11, 0x57, 0x1f, 0xbe, 0x46, 0x2f, 0x52, 0x3f, 0xa6, 0x1b, 0x96, 0x3e, 0xb4, 0xf9, 0x3a, 0x8e,
	0xfd, 0xf4, 0xcf, 0x9c, 0xd4, 0xfc, 0xbb, 0x0c, 0x47, 0x86, 0x6c, 0x82, 0xbe, 0x85, 0xca, 0xc6,
	0x8f, 0x22, 0x22, 0x96, 0x31, 0xd5, 0x94, 0x86, 0xd2, 0x52, 0xf0, 0xa9, 0x2c, 0x78, 0xcb, 0x98,
	0x22, 0x04, 0x87, 0x5c, 0xd0, 0x95, 0x76, 0xd0, 0x50, 0x5a, 0x65, 0x9c, 0xc5, 0xe8, 0x7b, 0xa8,
	0x2d, 0x96, 0x11, 0x25, 0x8f, 0x34, 0xe5, 0x4b, 0x96, 0x68, 0xe5, 0x86, 0xd2, 0xaa, 0x98, 0x25,
	0x5c, 0x95, 0xd5, 0xcf, 0x79, 0x11, 0x5d, 0x43, 0x25, 0x48, 0xfd, 0x55, 0x48, 0xe6, 0x74, 0xa1,
	0x1d, 0x36, 0x94, 0x56, 0xcd, 0x2c, 0xe1, 0xd3, 0xac, 0xd4, 0xa3, 0x0b, 0xd4, 0x86, 0x93, 0x42,
	0x8f, 0x76, 0xd4, 0x50, 0x5a, 0xd5, 0xce, 0xc5, 0xcd, 0x56, 0xf9, 0x8d, 0x9b, 0x2f, 0x99, 0x25,
	0xfc, 0x84, 0x42, 0xbf, 0x40, 0x35, 0x62, 0x01, 0x89, 0x29, 0xe7, 0x7e, 0x40, 0xb5, 0xe3, 0x8c,
	0x74, 0xb9, 0x4b, 0x1a, 0xb2, 0x60, 0x94, 0xaf, 0x9a, 0x25, 0x0c, 0xd1, 0x73, 0x26, 0xa9, 0x9c,
	0x72, 0xa9, 0x8a, 0x44, 0x2c, 0xd0, 0x4e, 0x5e, 0x52, 0xdd, 0x7c, 0x79, 0xc8, 0x02, 0x49, 0xe5,
	0xcf, 0x19, 0x72, 0xe0, 0x42, 0xf8, 0x41, 0x40, 0xe7, 0x24, 0x5d, 0x27, 0x24, 0xa6, 0xc2, 0x9f,
	0xfb, 0xc2, 0xd7, 0x4e, 0xb3, 0x16, 0xd7, 0xbb, 0x2d, 0xbc, 0x0c, 0x86, 0xd7, 0xc9, 0xa8, 0x00,
	0x99, 0x25, 0x7c, 0x2e, 0xbe, 0x2e, 0xa2, 0x1f, 0xa1, 0x2e, 0xbb, 0x90, 0xad, 0x37, 0x95, 0xc2,
	0x9b, 0x9a, 0xac, 0x0f, 0x0a, 0x7f, 0x6e, 0x8f, 0xe1, 0x70, 0x13, 0xfa, 0xa2, 0xf9, 0xaf, 0x02,
	0xb0, 0xdd, 0x18, 0xea, 0xc0, 0x51, 0x44, 0x1f, 0x69, 0x94, 0x9d, 0x53, 0xbd, 0xf3, 0xee, 0xf5,
	0xfd, 0xdf, 0x0c, 0x25, 0x06, 0xe7, 0x50, 0xa4, 0xc1, 0xc9, 0x93, 0x6b, 0xf2, 0x14, 0x2b, 0xf8,
	0x29, 0x6d, 0x8e, 0xe0, 0x28, 0x43, 0xa2, 0x2a, 0x9c, 0x4c, 0xec, 0x3b, 0xdb, 0x99, 0xda, 0x6a,
	0x09, 0x9d, 0x41, 0xa5, 0x67, 0xdc, 0x4e, 0x06, 0x03, 0xcb, 0x1e, 0xa8, 0x80, 0x4e, 0xe1, 0xd0,
	0xb2, 0xfb, 0x8e, 0xfa, 0x56, 0x46, 0x53, 0x1d, 0xdb, 0xea, 0x7b, 0x54, 0x81, 0x23, 0x03, 0x63,
	0x07, 0xab, 0x2d, 0x19, 0xf6, 0x75, 0x4f, 0x1f, 0xaa, 0x9d, 0xe6, 0x7f, 0x0a, 0xc0, 0xd6, 0x49,
	0xf4, 0x2b, 0x1c, 0x73, 0xe1, 0x8b, 0x35, 0x2f, 0xc4, 0xfe, 0xf0, 0xba, 0xe3, 0x4f, 0xa1, 0x9b,
	0x61, 0x71, 0xc1, 0x41, 0x1f, 0xe0, 0xcd, 0x2c, 0xa4, 0xb3, 0x87, 0x15, 0x5b, 0x26, 0x82, 0xac,
	0x7c, 0x11, 0x16, 0xea, 0xeb, 0xdb, 0xf2, 0xd8, 0x17, 0x21, 0x52, 0xa1, 0x1c, 0xf3, 0x20, 0x1f,
	0x42, 0x2c, 0xc3, 0xe6, 0x10, 0xce, 0xf6, 0x7a, 0xa2, 0x4b, 0x40, 0xae, 0xa7, 0x7b, 0x13, 0x97,
	0x4c, 0x6c, 0x77, 0x6c, 0x74, 0xad, 0xbe, 0x65, 0xf4, 0xd4, 0x92, 0xd4, 0xee, 0x7a, 0x3a, 0xf6,
	0x54, 0x45, 0xee, 0xcd, 0xf5, 0x9c, 0xb1, 0x7a, 0x80, 0xea, 0x00, 0x5d, 0xd3, 0xe8, 0xde, 0x8d,
	0x1d, 0xcb, 0xf6, 0xd4, 0x72, 0xd3, 0x84, 0xf3, 0x17, 0x67, 0x2b, 0x3f, 0x2a, 0xfc, 0x20, 0xdb,
	0x58, 0x05, 0xcb, 0x10, 0x7d, 0x07, 0xb5, 0xbd, 0x11, 0x91, 0x62, 0x6b, 0xb8, 0x9a, 0x6e, 0x49,
	0xcd, 0x36, 0xd4, 0xa7, 0xbe, 0x98, 0x85, 0x73, 0x16, 0x74, 0x59, 0xb2, 0x58, 0x06, 0xe8, 0x1a,
	0x40, 0xde, 0x3a, 0xb6, 0x16, 0x24, 0xce, 0x6d, 0x2a, 0xe3, 0x4a, 0x51, 0x19, 0xf1, 0xe6, 0x5f,
	0x0a, 0x5c, 0x4e, 0x59, 0xfa, 0x40, 0x53, 0x93, 0xfa, 0xa9, 0xb8, 0xa7, 0xbe, 0xc0, 0xf4, 0x8f,
	0x35, 0xe5, 0x02, 0x75, 0xe1, 0x8c, 0x87, 0x6b, 0x31, 0x67, 0x9b, 0x84, 0xc4, 0x6c, 0x4e, 0x0b,
	0x8f, 0xdf, 0xef, 0x7a, 0x9c, 0x53, 0xdd, 0x02, 0x36, 0x62, 0x73, 0x8a, 0x6b, 0x7c, 0x27, 0x43,
	0x5d, 0x78, 0xb3, 0x29, 0x04, 0x91, 0x59, 0xa6, 0x28, 0x93, 0x5d, 0xed, 0x5c, 0xed, 0xb5, 0xd9,
	0xd3, 0x8c, 0xeb, 0x9b, 0xbd, 0xbc, 0xf9, 0x8f, 0x02, 0xdf, 0xbc, 0x10, 0xc9, 0x57, 0x2c, 0xe1,
	0x14, 0x7d, 0x82, 0xb3, 0x90, 0xfa, 0x91, 0x08, 0xc9, 0xde, 0x24, 0x68, 0x2f, 0x55, 0x9a, 0x19,
	0x0c, 0xd7, 0x72, 0x78, 0x71, 0x6e, 0x3f, 0x01, 0x6c, 0xb2, 0xd5, 0xec, 0xde, 0x1e, 0x34, 0xca,
	0xad, 0x6a, 0xe7, 0x7c, 0x97, 0x9b, 0xfd, 0xc0, 0x70, 0x25, 0x07, 0xc9, 0x99, 0xbb, 0x82, 0xd3,
	0x90, 0x71, 0x91, 0xf8, 0x31, 0x2d, 0x26, 0xe2, 0x39, 0xff, 0xf8, 0x3b, 0xd4, 0x76, 0xbf, 0x85,
	0x8e, 0xe1, 0xc0, 0xb9, 0x53, 0x4b, 0xe8, 0x1d, 0x68, 0xd8, 0xe8, 0x1a, 0xd6, 0x67, 0xa3, 0x47,
	0x5c, 0x73, 0xe2, 0xf5, 0x9c, 0xa9, 0x4d, 0x5c, 0x6b, 0x60, 0xeb, 0x43, 0x55, 0x41, 0x08, 0xea,
	0x96, 0xed, 0x19, 0xd8, 0xd6, 0x87, 0x24, 0x9f, 0xf9, 0x03, 0x74, 0x0e, 0x67, 0x12, 0xe8, 0x59,
	0xf6, 0x80, 0x48, 0xb4, 0x5a, 0xfe, 0xf8, 0x00, 0xe8, 0xa5, 0xdd, 0xf2, 0x5e, 0xf5, 0x8c, 0xbe,
	0x3e, 0x19, 0x7a, 0x6a, 0x49, 0x76, 0xb2, 0x1d, 0x8f, 0x74, 0x1d, 0xbb, 0x6f, 0x0d, 0x26, 0xd8,
	0xe8, 0xa9, 0x0a, 0xd2, 0xe0, 0xed, 0x54, 0xb7, 0x3c, 0xd2, 0x77, 0x30, 0xe9, 0x3a, 0x0e, 0xee,
	0x59, 0xb6, 0xee, 0x65, 0xdf, 0xb8, 0x82, 0xcb, 0x67, 0x31, 0x7a, 0xdf, 0x33, 0x30, 0xf1, 0xac,
	0x91, 0xe1, 0x4c, 0x3c, 0xb5, 0x7c, 0xbb, 0x80, 0x0b, 0x96, 0x06, 0xbb, 0x46, 0xc8, 0x37, 0xe1,
	0xb6, 0x9a, 0xd9, 0x31, 0x96, 0xbf, 0x77, 0x3e, 0x56, 0x7e, 0xfb, 0x14, 0x2c, 0x45, 0xb8, 0xbe,
	0xbf, 0x99, 0xb1, 0xb8, 0xbd, 0xf3, 0x2a, 0xbc, 0x1e, 0x06, 0x2c, 0x7f, 0x2e, 0x9e, 0x5f, 0x98,
	0xff, 0x15, 0xe5, 0xfe, 0x38, 0x7b, 0x2a, 0x7e, 0xfe, 0x12, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x30,
	0x77, 0x95, 0x83, 0x06, 0x00, 0x00,
}
