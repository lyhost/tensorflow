// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/tools/api/lib/api_objects.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type TFAPIMember struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Mtype                *string  `protobuf:"bytes,2,opt,name=mtype" json:"mtype,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TFAPIMember) Reset()         { *m = TFAPIMember{} }
func (m *TFAPIMember) String() string { return proto.CompactTextString(m) }
func (*TFAPIMember) ProtoMessage()    {}
func (*TFAPIMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_1979de05ee66a62a, []int{0}
}

func (m *TFAPIMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFAPIMember.Unmarshal(m, b)
}
func (m *TFAPIMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFAPIMember.Marshal(b, m, deterministic)
}
func (m *TFAPIMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFAPIMember.Merge(m, src)
}
func (m *TFAPIMember) XXX_Size() int {
	return xxx_messageInfo_TFAPIMember.Size(m)
}
func (m *TFAPIMember) XXX_DiscardUnknown() {
	xxx_messageInfo_TFAPIMember.DiscardUnknown(m)
}

var xxx_messageInfo_TFAPIMember proto.InternalMessageInfo

func (m *TFAPIMember) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TFAPIMember) GetMtype() string {
	if m != nil && m.Mtype != nil {
		return *m.Mtype
	}
	return ""
}

type TFAPIMethod struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Path                 *string  `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Argspec              *string  `protobuf:"bytes,3,opt,name=argspec" json:"argspec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TFAPIMethod) Reset()         { *m = TFAPIMethod{} }
func (m *TFAPIMethod) String() string { return proto.CompactTextString(m) }
func (*TFAPIMethod) ProtoMessage()    {}
func (*TFAPIMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_1979de05ee66a62a, []int{1}
}

func (m *TFAPIMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFAPIMethod.Unmarshal(m, b)
}
func (m *TFAPIMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFAPIMethod.Marshal(b, m, deterministic)
}
func (m *TFAPIMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFAPIMethod.Merge(m, src)
}
func (m *TFAPIMethod) XXX_Size() int {
	return xxx_messageInfo_TFAPIMethod.Size(m)
}
func (m *TFAPIMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_TFAPIMethod.DiscardUnknown(m)
}

var xxx_messageInfo_TFAPIMethod proto.InternalMessageInfo

func (m *TFAPIMethod) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TFAPIMethod) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *TFAPIMethod) GetArgspec() string {
	if m != nil && m.Argspec != nil {
		return *m.Argspec
	}
	return ""
}

type TFAPIModule struct {
	Member               []*TFAPIMember `protobuf:"bytes,1,rep,name=member" json:"member,omitempty"`
	MemberMethod         []*TFAPIMethod `protobuf:"bytes,2,rep,name=member_method,json=memberMethod" json:"member_method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TFAPIModule) Reset()         { *m = TFAPIModule{} }
func (m *TFAPIModule) String() string { return proto.CompactTextString(m) }
func (*TFAPIModule) ProtoMessage()    {}
func (*TFAPIModule) Descriptor() ([]byte, []int) {
	return fileDescriptor_1979de05ee66a62a, []int{2}
}

func (m *TFAPIModule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFAPIModule.Unmarshal(m, b)
}
func (m *TFAPIModule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFAPIModule.Marshal(b, m, deterministic)
}
func (m *TFAPIModule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFAPIModule.Merge(m, src)
}
func (m *TFAPIModule) XXX_Size() int {
	return xxx_messageInfo_TFAPIModule.Size(m)
}
func (m *TFAPIModule) XXX_DiscardUnknown() {
	xxx_messageInfo_TFAPIModule.DiscardUnknown(m)
}

var xxx_messageInfo_TFAPIModule proto.InternalMessageInfo

func (m *TFAPIModule) GetMember() []*TFAPIMember {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *TFAPIModule) GetMemberMethod() []*TFAPIMethod {
	if m != nil {
		return m.MemberMethod
	}
	return nil
}

type TFAPIClass struct {
	IsInstance           []string       `protobuf:"bytes,1,rep,name=is_instance,json=isInstance" json:"is_instance,omitempty"`
	Member               []*TFAPIMember `protobuf:"bytes,2,rep,name=member" json:"member,omitempty"`
	MemberMethod         []*TFAPIMethod `protobuf:"bytes,3,rep,name=member_method,json=memberMethod" json:"member_method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TFAPIClass) Reset()         { *m = TFAPIClass{} }
func (m *TFAPIClass) String() string { return proto.CompactTextString(m) }
func (*TFAPIClass) ProtoMessage()    {}
func (*TFAPIClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_1979de05ee66a62a, []int{3}
}

func (m *TFAPIClass) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFAPIClass.Unmarshal(m, b)
}
func (m *TFAPIClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFAPIClass.Marshal(b, m, deterministic)
}
func (m *TFAPIClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFAPIClass.Merge(m, src)
}
func (m *TFAPIClass) XXX_Size() int {
	return xxx_messageInfo_TFAPIClass.Size(m)
}
func (m *TFAPIClass) XXX_DiscardUnknown() {
	xxx_messageInfo_TFAPIClass.DiscardUnknown(m)
}

var xxx_messageInfo_TFAPIClass proto.InternalMessageInfo

func (m *TFAPIClass) GetIsInstance() []string {
	if m != nil {
		return m.IsInstance
	}
	return nil
}

func (m *TFAPIClass) GetMember() []*TFAPIMember {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *TFAPIClass) GetMemberMethod() []*TFAPIMethod {
	if m != nil {
		return m.MemberMethod
	}
	return nil
}

type TFAPIProto struct {
	Descriptor_          *descriptor.DescriptorProto `protobuf:"bytes,1,opt,name=descriptor" json:"descriptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TFAPIProto) Reset()         { *m = TFAPIProto{} }
func (m *TFAPIProto) String() string { return proto.CompactTextString(m) }
func (*TFAPIProto) ProtoMessage()    {}
func (*TFAPIProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_1979de05ee66a62a, []int{4}
}

func (m *TFAPIProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFAPIProto.Unmarshal(m, b)
}
func (m *TFAPIProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFAPIProto.Marshal(b, m, deterministic)
}
func (m *TFAPIProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFAPIProto.Merge(m, src)
}
func (m *TFAPIProto) XXX_Size() int {
	return xxx_messageInfo_TFAPIProto.Size(m)
}
func (m *TFAPIProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TFAPIProto.DiscardUnknown(m)
}

var xxx_messageInfo_TFAPIProto proto.InternalMessageInfo

func (m *TFAPIProto) GetDescriptor_() *descriptor.DescriptorProto {
	if m != nil {
		return m.Descriptor_
	}
	return nil
}

type TFAPIObject struct {
	Path                 *string      `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	TfModule             *TFAPIModule `protobuf:"bytes,2,opt,name=tf_module,json=tfModule" json:"tf_module,omitempty"`
	TfClass              *TFAPIClass  `protobuf:"bytes,3,opt,name=tf_class,json=tfClass" json:"tf_class,omitempty"`
	TfProto              *TFAPIProto  `protobuf:"bytes,4,opt,name=tf_proto,json=tfProto" json:"tf_proto,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TFAPIObject) Reset()         { *m = TFAPIObject{} }
func (m *TFAPIObject) String() string { return proto.CompactTextString(m) }
func (*TFAPIObject) ProtoMessage()    {}
func (*TFAPIObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_1979de05ee66a62a, []int{5}
}

func (m *TFAPIObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TFAPIObject.Unmarshal(m, b)
}
func (m *TFAPIObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TFAPIObject.Marshal(b, m, deterministic)
}
func (m *TFAPIObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TFAPIObject.Merge(m, src)
}
func (m *TFAPIObject) XXX_Size() int {
	return xxx_messageInfo_TFAPIObject.Size(m)
}
func (m *TFAPIObject) XXX_DiscardUnknown() {
	xxx_messageInfo_TFAPIObject.DiscardUnknown(m)
}

var xxx_messageInfo_TFAPIObject proto.InternalMessageInfo

func (m *TFAPIObject) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *TFAPIObject) GetTfModule() *TFAPIModule {
	if m != nil {
		return m.TfModule
	}
	return nil
}

func (m *TFAPIObject) GetTfClass() *TFAPIClass {
	if m != nil {
		return m.TfClass
	}
	return nil
}

func (m *TFAPIObject) GetTfProto() *TFAPIProto {
	if m != nil {
		return m.TfProto
	}
	return nil
}

func init() {
	proto.RegisterType((*TFAPIMember)(nil), "third_party.tensorflow.tools.api.TFAPIMember")
	proto.RegisterType((*TFAPIMethod)(nil), "third_party.tensorflow.tools.api.TFAPIMethod")
	proto.RegisterType((*TFAPIModule)(nil), "third_party.tensorflow.tools.api.TFAPIModule")
	proto.RegisterType((*TFAPIClass)(nil), "third_party.tensorflow.tools.api.TFAPIClass")
	proto.RegisterType((*TFAPIProto)(nil), "third_party.tensorflow.tools.api.TFAPIProto")
	proto.RegisterType((*TFAPIObject)(nil), "third_party.tensorflow.tools.api.TFAPIObject")
}

func init() {
	proto.RegisterFile("tensorflow/tools/api/lib/api_objects.proto", fileDescriptor_1979de05ee66a62a)
}

var fileDescriptor_1979de05ee66a62a = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0x25, 0x33, 0xa3, 0xb3, 0x53, 0x51, 0x90, 0xc6, 0x43, 0xf0, 0x62, 0xc8, 0x69, 0x11, 0xed,
	0xc0, 0x5c, 0x04, 0x41, 0xfc, 0x56, 0x56, 0x90, 0x5d, 0x1a, 0x4f, 0x5e, 0x42, 0x27, 0xd3, 0xc9,
	0xb4, 0x24, 0xe9, 0xa6, 0xbb, 0x06, 0xd9, 0x7f, 0xe4, 0x4f, 0xf1, 0x5f, 0x29, 0x53, 0x9d, 0xcc,
	0xe4, 0xb0, 0xa0, 0x83, 0x7b, 0x4a, 0x55, 0xa5, 0xdf, 0xa3, 0x5e, 0xbf, 0xd7, 0xf0, 0x04, 0x55,
	0xef, 0x8d, 0xab, 0x5b, 0xf3, 0x23, 0x47, 0x63, 0x5a, 0x9f, 0x4b, 0xab, 0xf3, 0x56, 0x97, 0xfb,
	0x6f, 0x61, 0xca, 0xef, 0xaa, 0x42, 0xcf, 0xad, 0x33, 0x68, 0x58, 0x8a, 0x5b, 0xed, 0x36, 0x85,
	0x95, 0x0e, 0xaf, 0xf9, 0x11, 0xc7, 0x09, 0xc7, 0xa5, 0xd5, 0x8f, 0xd2, 0xc6, 0x98, 0xa6, 0x55,
	0x39, 0x9d, 0x2f, 0x77, 0x75, 0xbe, 0x51, 0xbe, 0x72, 0xda, 0xa2, 0x71, 0x81, 0x23, 0x7b, 0x0e,
	0xf1, 0xd7, 0x8f, 0x6f, 0xae, 0x2e, 0xbe, 0xa8, 0xae, 0x54, 0x8e, 0x31, 0x58, 0xf4, 0xb2, 0x53,
	0x49, 0x94, 0x46, 0xe7, 0x2b, 0x41, 0x35, 0x7b, 0x08, 0x77, 0x3a, 0xbc, 0xb6, 0x2a, 0x99, 0xd1,
	0x30, 0x34, 0xd9, 0xe5, 0x01, 0x88, 0x5b, 0xb3, 0xb9, 0x11, 0xc8, 0x60, 0x61, 0x25, 0x6e, 0x07,
	0x1c, 0xd5, 0x2c, 0x81, 0xa5, 0x74, 0x8d, 0xb7, 0xaa, 0x4a, 0xe6, 0x34, 0x1e, 0xdb, 0xec, 0x67,
	0x34, 0x32, 0x9a, 0xcd, 0xae, 0x55, 0xec, 0x03, 0xdc, 0xed, 0x68, 0xa9, 0x24, 0x4a, 0xe7, 0xe7,
	0xf1, 0xfa, 0x19, 0xff, 0x9b, 0x5c, 0x3e, 0x51, 0x22, 0x06, 0x30, 0x13, 0x70, 0x3f, 0x54, 0x45,
	0x47, 0x9b, 0x26, 0xb3, 0x13, 0xd9, 0xf6, 0x20, 0x71, 0x2f, 0x70, 0x84, 0x2e, 0xfb, 0x15, 0x01,
	0xd0, 0xdf, 0x77, 0xad, 0xf4, 0x9e, 0x3d, 0x86, 0x58, 0xfb, 0x42, 0xf7, 0x1e, 0x65, 0x5f, 0x29,
	0x5a, 0x77, 0x25, 0x40, 0xfb, 0x8b, 0x61, 0x32, 0x91, 0x32, 0xbb, 0x55, 0x29, 0xf3, 0xff, 0x97,
	0x22, 0x06, 0x25, 0x57, 0x94, 0xa8, 0xd7, 0x00, 0xc7, 0x84, 0x90, 0x97, 0xf1, 0x3a, 0xe5, 0x21,
	0x44, 0x7c, 0x0c, 0x11, 0x7f, 0x7f, 0x38, 0x42, 0x28, 0x31, 0xc1, 0xbc, 0x98, 0x3d, 0x88, 0xb2,
	0xdf, 0xa3, 0x93, 0x97, 0x14, 0xd7, 0x43, 0x0e, 0xa2, 0x49, 0x0e, 0x3e, 0xc3, 0x0a, 0xeb, 0xa2,
	0x23, 0xab, 0x29, 0x20, 0x27, 0xe8, 0x20, 0x90, 0x38, 0xc3, 0x7a, 0x48, 0xca, 0x27, 0x38, 0xc3,
	0xba, 0xa8, 0xf6, 0x5e, 0x50, 0xa8, 0xe2, 0xf5, 0xd3, 0x7f, 0xa4, 0x22, 0xff, 0xc4, 0x12, 0xeb,
	0x60, 0x64, 0x20, 0x22, 0x9d, 0xc9, 0xe2, 0x24, 0xa2, 0x70, 0x11, 0x4b, 0xac, 0xa9, 0x78, 0xfb,
	0xea, 0xdb, 0xcb, 0x46, 0xe3, 0x76, 0x57, 0xf2, 0xca, 0x74, 0xf9, 0xf4, 0x49, 0xdf, 0x58, 0x36,
	0x26, 0xaf, 0x8c, 0x53, 0xc7, 0x07, 0xff, 0x27, 0x00, 0x00, 0xff, 0xff, 0x84, 0x5d, 0xd7, 0x82,
	0x07, 0x04, 0x00, 0x00,
}
