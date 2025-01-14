// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/compiler/xla/service/gpu/backend_configs.proto

package xla

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

// Backend config for a convolution that runs through cudnn.
type CudnnConvBackendConfig struct {
	// Opaque algorithm number of cudnn algorithm chosen for this conv.
	Algorithm int64 `protobuf:"varint,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	// Whether we may use tensor cores when running this conv.  Even if this is
	// true, cudnn may choose not to use tensor cores, e.g. because the GPU or
	// selected algorithm doesn't support it.
	TensorOpsEnabled bool `protobuf:"varint,2,opt,name=tensor_ops_enabled,json=tensorOpsEnabled,proto3" json:"tensor_ops_enabled,omitempty"`
	// The scaling factor multiplied with the convolution result.
	ConvResultScale float64 `protobuf:"fixed64,4,opt,name=conv_result_scale,json=convResultScale,proto3" json:"conv_result_scale,omitempty"`
	// The requested activation (e.g. relu) after the convolution. It is with type
	// stream_executor::dnn::ActivationMode.
	ActivationMode int64 `protobuf:"varint,3,opt,name=activation_mode,json=activationMode,proto3" json:"activation_mode,omitempty"`
	// The scaling factor multiplied with the side input. If no side input buffer
	// is provided, this field must be 0.
	SideInputScale       float64  `protobuf:"fixed64,5,opt,name=side_input_scale,json=sideInputScale,proto3" json:"side_input_scale,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CudnnConvBackendConfig) Reset()         { *m = CudnnConvBackendConfig{} }
func (m *CudnnConvBackendConfig) String() string { return proto.CompactTextString(m) }
func (*CudnnConvBackendConfig) ProtoMessage()    {}
func (*CudnnConvBackendConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c20905cd9140fdf, []int{0}
}

func (m *CudnnConvBackendConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CudnnConvBackendConfig.Unmarshal(m, b)
}
func (m *CudnnConvBackendConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CudnnConvBackendConfig.Marshal(b, m, deterministic)
}
func (m *CudnnConvBackendConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CudnnConvBackendConfig.Merge(m, src)
}
func (m *CudnnConvBackendConfig) XXX_Size() int {
	return xxx_messageInfo_CudnnConvBackendConfig.Size(m)
}
func (m *CudnnConvBackendConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CudnnConvBackendConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CudnnConvBackendConfig proto.InternalMessageInfo

func (m *CudnnConvBackendConfig) GetAlgorithm() int64 {
	if m != nil {
		return m.Algorithm
	}
	return 0
}

func (m *CudnnConvBackendConfig) GetTensorOpsEnabled() bool {
	if m != nil {
		return m.TensorOpsEnabled
	}
	return false
}

func (m *CudnnConvBackendConfig) GetConvResultScale() float64 {
	if m != nil {
		return m.ConvResultScale
	}
	return 0
}

func (m *CudnnConvBackendConfig) GetActivationMode() int64 {
	if m != nil {
		return m.ActivationMode
	}
	return 0
}

func (m *CudnnConvBackendConfig) GetSideInputScale() float64 {
	if m != nil {
		return m.SideInputScale
	}
	return 0
}

// Backend config for the GEMM operation running through cuBLAS.
type GemmBackendConfig struct {
	// Opaque optional algorithm number. No chosen number indicates that a
	// different cuBLAS API will be used, which does not allow for choosing an
	// algorithm.
	//
	// Types that are valid to be assigned to Algorithm:
	//	*GemmBackendConfig_SelectedAlgorithm
	Algorithm            isGemmBackendConfig_Algorithm `protobuf_oneof:"algorithm"`
	AlphaReal            float64                       `protobuf:"fixed64,2,opt,name=alpha_real,json=alphaReal,proto3" json:"alpha_real,omitempty"`
	AlphaImag            float64                       `protobuf:"fixed64,9,opt,name=alpha_imag,json=alphaImag,proto3" json:"alpha_imag,omitempty"`
	Beta                 float64                       `protobuf:"fixed64,3,opt,name=beta,proto3" json:"beta,omitempty"`
	DotDimensionNumbers  *DotDimensionNumbers          `protobuf:"bytes,7,opt,name=dot_dimension_numbers,json=dotDimensionNumbers,proto3" json:"dot_dimension_numbers,omitempty"`
	BatchSize            int64                         `protobuf:"varint,8,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *GemmBackendConfig) Reset()         { *m = GemmBackendConfig{} }
func (m *GemmBackendConfig) String() string { return proto.CompactTextString(m) }
func (*GemmBackendConfig) ProtoMessage()    {}
func (*GemmBackendConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c20905cd9140fdf, []int{1}
}

func (m *GemmBackendConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GemmBackendConfig.Unmarshal(m, b)
}
func (m *GemmBackendConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GemmBackendConfig.Marshal(b, m, deterministic)
}
func (m *GemmBackendConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GemmBackendConfig.Merge(m, src)
}
func (m *GemmBackendConfig) XXX_Size() int {
	return xxx_messageInfo_GemmBackendConfig.Size(m)
}
func (m *GemmBackendConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GemmBackendConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GemmBackendConfig proto.InternalMessageInfo

type isGemmBackendConfig_Algorithm interface {
	isGemmBackendConfig_Algorithm()
}

type GemmBackendConfig_SelectedAlgorithm struct {
	SelectedAlgorithm int64 `protobuf:"varint,1,opt,name=selected_algorithm,json=selectedAlgorithm,proto3,oneof"`
}

func (*GemmBackendConfig_SelectedAlgorithm) isGemmBackendConfig_Algorithm() {}

func (m *GemmBackendConfig) GetAlgorithm() isGemmBackendConfig_Algorithm {
	if m != nil {
		return m.Algorithm
	}
	return nil
}

func (m *GemmBackendConfig) GetSelectedAlgorithm() int64 {
	if x, ok := m.GetAlgorithm().(*GemmBackendConfig_SelectedAlgorithm); ok {
		return x.SelectedAlgorithm
	}
	return 0
}

func (m *GemmBackendConfig) GetAlphaReal() float64 {
	if m != nil {
		return m.AlphaReal
	}
	return 0
}

func (m *GemmBackendConfig) GetAlphaImag() float64 {
	if m != nil {
		return m.AlphaImag
	}
	return 0
}

func (m *GemmBackendConfig) GetBeta() float64 {
	if m != nil {
		return m.Beta
	}
	return 0
}

func (m *GemmBackendConfig) GetDotDimensionNumbers() *DotDimensionNumbers {
	if m != nil {
		return m.DotDimensionNumbers
	}
	return nil
}

func (m *GemmBackendConfig) GetBatchSize() int64 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GemmBackendConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GemmBackendConfig_SelectedAlgorithm)(nil),
	}
}

func init() {
	proto.RegisterType((*CudnnConvBackendConfig)(nil), "xla.gpu.CudnnConvBackendConfig")
	proto.RegisterType((*GemmBackendConfig)(nil), "xla.gpu.GemmBackendConfig")
}

func init() {
	proto.RegisterFile("tensorflow/compiler/xla/service/gpu/backend_configs.proto", fileDescriptor_9c20905cd9140fdf)
}

var fileDescriptor_9c20905cd9140fdf = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xdf, 0x6a, 0x14, 0x31,
	0x14, 0xc6, 0x9d, 0xb6, 0xda, 0x36, 0x85, 0xb6, 0x1b, 0x51, 0x06, 0x51, 0x58, 0x7a, 0xa1, 0x83,
	0xc8, 0x0e, 0xe8, 0x95, 0x77, 0xba, 0x5b, 0xd1, 0x82, 0x7f, 0x60, 0x7a, 0xe7, 0x4d, 0x38, 0x93,
	0x9c, 0xce, 0x04, 0x93, 0x9c, 0x21, 0xc9, 0xac, 0x4b, 0x9f, 0xc2, 0x77, 0xf4, 0x45, 0x64, 0x32,
	0xac, 0xbb, 0xeb, 0x9f, 0xbb, 0xc3, 0xef, 0x4b, 0xbe, 0x93, 0x2f, 0xe7, 0xb0, 0xd7, 0x11, 0x5d,
	0x20, 0x7f, 0x63, 0xe8, 0x7b, 0x29, 0xc9, 0x76, 0xda, 0xa0, 0x2f, 0x57, 0x06, 0xca, 0x80, 0x7e,
	0xa9, 0x25, 0x96, 0x4d, 0xd7, 0x97, 0x35, 0xc8, 0x6f, 0xe8, 0x94, 0x90, 0xe4, 0x6e, 0x74, 0x13,
	0x66, 0x9d, 0xa7, 0x48, 0xfc, 0x70, 0x65, 0x60, 0xd6, 0x74, 0xfd, 0xa3, 0xa7, 0xff, 0xf3, 0x58,
	0x19, 0x10, 0x0a, 0x22, 0x8c, 0x17, 0x2e, 0x7e, 0x66, 0xec, 0xe1, 0xa2, 0x57, 0xce, 0x2d, 0xc8,
	0x2d, 0xe7, 0xa3, 0xe7, 0x22, 0x59, 0xf2, 0xc7, 0xec, 0x18, 0x4c, 0x43, 0x5e, 0xc7, 0xd6, 0xe6,
	0xd9, 0x34, 0x2b, 0xf6, 0xab, 0x0d, 0xe0, 0x2f, 0x18, 0x1f, 0x5b, 0x08, 0xea, 0x82, 0x40, 0x07,
	0xb5, 0x41, 0x95, 0xef, 0x4d, 0xb3, 0xe2, 0xa8, 0x3a, 0x1f, 0x95, 0x2f, 0x5d, 0x78, 0x37, 0x72,
	0xfe, 0x9c, 0x4d, 0x24, 0xb9, 0xa5, 0xf0, 0x18, 0x7a, 0x13, 0x45, 0x90, 0x60, 0x30, 0x3f, 0x98,
	0x66, 0x45, 0x56, 0x9d, 0x0d, 0x42, 0x95, 0xf8, 0xf5, 0x80, 0xf9, 0x33, 0x76, 0x06, 0x32, 0xea,
	0x25, 0x44, 0x4d, 0x4e, 0x58, 0x52, 0x98, 0xef, 0xa7, 0xee, 0xa7, 0x1b, 0xfc, 0x89, 0x14, 0xf2,
	0x82, 0x9d, 0x07, 0xad, 0x50, 0x68, 0xd7, 0xf5, 0x6b, 0xcf, 0xbb, 0xc9, 0xf3, 0x74, 0xe0, 0x57,
	0x03, 0x4e, 0x96, 0x17, 0x3f, 0xf6, 0xd8, 0xe4, 0x3d, 0x5a, 0xbb, 0x1b, 0xb0, 0x64, 0x3c, 0xa0,
	0x41, 0x19, 0x51, 0x89, 0x3f, 0x92, 0x7e, 0xb8, 0x53, 0x4d, 0xd6, 0xda, 0xdb, 0xdf, 0x99, 0x9f,
	0x30, 0x06, 0xa6, 0x6b, 0x41, 0x78, 0x04, 0x93, 0xb2, 0x66, 0xc3, 0x97, 0x74, 0x2d, 0x54, 0x08,
	0x66, 0x23, 0x6b, 0x0b, 0x4d, 0x7e, 0xbc, 0x25, 0x5f, 0x59, 0x68, 0x38, 0x67, 0x07, 0x35, 0x46,
	0x48, 0x61, 0xb2, 0x2a, 0xd5, 0xfc, 0x23, 0x7b, 0xa0, 0x28, 0x0a, 0xa5, 0x2d, 0xba, 0x30, 0xc4,
	0x75, 0xbd, 0xad, 0xd1, 0x87, 0xfc, 0x70, 0x9a, 0x15, 0x27, 0x2f, 0xf3, 0xd9, 0x30, 0xcf, 0x4b,
	0x8a, 0x97, 0xeb, 0x03, 0x9f, 0x47, 0xbd, 0xba, 0xaf, 0xfe, 0x86, 0xc3, 0x03, 0x6a, 0x88, 0xb2,
	0x15, 0x41, 0xdf, 0x62, 0x7e, 0x34, 0x8e, 0x2c, 0x91, 0x6b, 0x7d, 0x8b, 0xf3, 0x93, 0xad, 0x81,
	0xce, 0xe7, 0x5f, 0xdf, 0x34, 0x3a, 0xb6, 0x7d, 0x3d, 0x93, 0x64, 0xcb, 0xad, 0x6d, 0xf9, 0x77,
	0xd9, 0x50, 0x29, 0xc9, 0xe3, 0xce, 0x2e, 0xd5, 0xf7, 0xd2, 0x0e, 0xbd, 0xfa, 0x15, 0x00, 0x00,
	0xff, 0xff, 0xc7, 0xd4, 0x6b, 0x33, 0xb1, 0x02, 0x00, 0x00,
}
