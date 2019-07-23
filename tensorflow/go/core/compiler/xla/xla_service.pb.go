// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/compiler/xla/rpc/xla_service.proto

package xla

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterFile("tensorflow/compiler/xla/rpc/xla_service.proto", fileDescriptor_fbab2827860b259b)
}

var fileDescriptor_fbab2827860b259b = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdd, 0x8a, 0x13, 0x31,
	0x14, 0x16, 0x84, 0x55, 0xe2, 0x8d, 0xa4, 0x55, 0x97, 0x65, 0xc5, 0x5d, 0xef, 0xed, 0x80, 0x0b,
	0x82, 0x57, 0x4a, 0xa7, 0x6e, 0x15, 0x04, 0x75, 0xdb, 0xc5, 0x65, 0x2f, 0x94, 0xb3, 0xd3, 0xd3,
	0x76, 0x30, 0x4d, 0xc6, 0xe4, 0x8c, 0xf6, 0x99, 0x7d, 0x0a, 0x99, 0x99, 0x24, 0x4d, 0xe6, 0xe7,
	0xaa, 0xcd, 0xf7, 0x37, 0x99, 0xaf, 0xa7, 0x87, 0xbd, 0x22, 0x94, 0x46, 0xe9, 0xb5, 0x50, 0x7f,
	0x93, 0x4c, 0xed, 0x8a, 0x5c, 0xa0, 0x4e, 0xf6, 0x02, 0x12, 0x5d, 0x64, 0xd5, 0xe7, 0x4f, 0x83,
	0xfa, 0x4f, 0x9e, 0xe1, 0xa4, 0xd0, 0x8a, 0x14, 0xbf, 0xbf, 0x17, 0x70, 0x72, 0x3e, 0xe4, 0xd9,
	0x0b, 0x68, 0x74, 0xaf, 0xff, 0x31, 0xc6, 0x6e, 0x04, 0x2c, 0x1a, 0x33, 0x7f, 0xc7, 0xd8, 0xb5,
	0xd4, 0xb8, 0xc9, 0x0d, 0xa1, 0xe6, 0x4f, 0x27, 0x95, 0xf0, 0x00, 0x5c, 0xe1, 0xef, 0x12, 0x0d,
	0x9d, 0x3c, 0xeb, 0xe0, 0xa6, 0x50, 0xd2, 0xe0, 0xcb, 0x7b, 0xfc, 0x1b, 0x7b, 0x3c, 0xc3, 0x4c,
	0x49, 0x43, 0xba, 0xcc, 0x68, 0x59, 0x16, 0x02, 0xf9, 0x69, 0x2d, 0x6f, 0xc3, 0x2e, 0xec, 0xf9,
	0x00, 0xeb, 0x23, 0x2f, 0xd8, 0xd1, 0xb5, 0x2c, 0x20, 0xfb, 0xc5, 0xb9, 0x7d, 0x6e, 0x75, 0x70,
	0xf6, 0x51, 0x84, 0x79, 0xd3, 0x5b, 0xf6, 0x70, 0x8e, 0xb4, 0xd8, 0x42, 0x81, 0x7c, 0x5c, 0x4b,
	0xdc, 0xd1, 0x19, 0x9f, 0xb4, 0x50, 0x6f, 0xfd, 0xc1, 0x8e, 0xe7, 0x48, 0xa9, 0xda, 0x15, 0x25,
	0x01, 0xe5, 0x4a, 0xce, 0x35, 0x14, 0xdb, 0x05, 0x01, 0x19, 0x7e, 0x5e, 0x9b, 0x7a, 0xb9, 0xf8,
	0x7d, 0x02, 0x89, 0x65, 0xc3, 0xab, 0x7d, 0x56, 0xb0, 0x9a, 0x01, 0x81, 0xbd, 0x9a, 0x3b, 0xc6,
	0x57, 0x3b, 0xa0, 0x61, 0xbb, 0x4b, 0x0d, 0xd2, 0xac, 0x51, 0x2f, 0x55, 0x2a, 0x72, 0x94, 0x64,
	0xdb, 0x6d, 0xc3, 0xf1, 0x6d, 0xba, 0x6c, 0x7f, 0x64, 0x35, 0x06, 0xa8, 0x3b, 0x91, 0x0d, 0x3c,
	0x14, 0xe9, 0xd8, 0xfe, 0xc8, 0x4f, 0x72, 0x8d, 0xb8, 0xea, 0x44, 0x36, 0xf0, 0x50, 0xa4, 0x63,
	0x7d, 0xe4, 0x2d, 0x1b, 0x39, 0xf6, 0x52, 0xab, 0xdd, 0x97, 0x92, 0xea, 0xd4, 0x17, 0x91, 0x2f,
	0x60, 0x5c, 0xf0, 0xd9, 0xb0, 0xc0, 0x67, 0x4f, 0xd9, 0xa3, 0x2b, 0x34, 0x48, 0x33, 0xac, 0xff,
	0x02, 0xcd, 0x70, 0x07, 0x88, 0xcb, 0x3a, 0xee, 0x12, 0x3e, 0xe3, 0x86, 0x8d, 0x9b, 0x5f, 0x1c,
	0xd3, 0x6a, 0x8c, 0x41, 0x52, 0x3d, 0x18, 0xfc, 0x2c, 0x18, 0x86, 0x98, 0x72, 0xa9, 0xa7, 0x7d,
	0x8a, 0xb8, 0xcc, 0xb9, 0x7b, 0xe0, 0x47, 0x90, 0x2b, 0x81, 0xc6, 0x96, 0xd9, 0x86, 0xe3, 0x32,
	0xbb, 0x6c, 0x58, 0x66, 0xaa, 0x11, 0x08, 0xd3, 0x2d, 0x48, 0x89, 0xa2, 0x51, 0xd8, 0x32, 0x7b,
	0x98, 0xb8, 0xcc, 0x5e, 0x81, 0xcf, 0x7e, 0xc3, 0x1e, 0xa4, 0xcd, 0xa6, 0xe1, 0x23, 0xff, 0x66,
	0xf9, 0x21, 0x63, 0x1c, 0x83, 0xa1, 0xef, 0xc3, 0x1e, 0xb3, 0x92, 0x9c, 0xcf, 0x9e, 0x62, 0x9f,
	0x07, 0xc3, 0xe2, 0x2d, 0x58, 0xb7, 0xfa, 0x15, 0x34, 0x08, 0x81, 0xc2, 0x16, 0xdf, 0x47, 0xc5,
	0xc5, 0x5b, 0xc5, 0x81, 0x0c, 0x8b, 0xff, 0x0e, 0x39, 0x5d, 0x2a, 0xdd, 0x68, 0x72, 0x25, 0x6d,
	0xf1, 0x6d, 0x38, 0x2e, 0xbe, 0xcb, 0xba, 0xc8, 0xe9, 0xf4, 0xf6, 0xfd, 0x26, 0xa7, 0x6d, 0x79,
	0x37, 0xc9, 0xd4, 0x2e, 0x09, 0x96, 0x73, 0xff, 0xd7, 0x8d, 0x4a, 0x32, 0xa5, 0x31, 0x5a, 0xdd,
	0x77, 0x47, 0xf5, 0xde, 0xbe, 0xf8, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x7d, 0xd3, 0xf0, 0x10,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// XlaServiceClient is the client API for XlaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type XlaServiceClient interface {
	// Unregisters a global allocation.
	//
	// If the handle given is not currently allocated, a NOT_FOUND status is
	// returned.
	Unregister(ctx context.Context, in *UnregisterRequest, opts ...grpc.CallOption) (*UnregisterResponse, error)
	// Deconstructs a tuple. Returns a newly created GlobalDataHandle for each
	// element in the tuple.
	DeconstructTuple(ctx context.Context, in *DeconstructTupleRequest, opts ...grpc.CallOption) (*DeconstructTupleResponse, error)
	// Unpack requests that a global data handle, with a tuple shape, has global
	// data handles created for each of its constituent members. This is the
	// equivalent of the "destructuring assignment" present in various programming
	// languages.
	Unpack(ctx context.Context, in *UnpackRequest, opts ...grpc.CallOption) (*UnpackResponse, error)
	// Requests the shape of the referenced global data.
	GetShape(ctx context.Context, in *GetShapeRequest, opts ...grpc.CallOption) (*GetShapeResponse, error)
	// Requests the statistics of the given computation.
	GetComputationGraphStats(ctx context.Context, in *ComputationGraphStatsRequest, opts ...grpc.CallOption) (*ComputationStatsResponse, error)
	// Loads a variable number of values with a given element type from ColumnIO.
	LoadData(ctx context.Context, in *LoadDataRequest, opts ...grpc.CallOption) (*LoadDataResponse, error)
	// Transfers the given global data to the client in the form of a Literal.
	TransferToClient(ctx context.Context, in *TransferToClientRequest, opts ...grpc.CallOption) (*TransferToClientResponse, error)
	// Transfers the given literal to the server to be stored in a global
	// allocation, which is returned.
	TransferToServer(ctx context.Context, in *TransferToServerRequest, opts ...grpc.CallOption) (*TransferToServerResponse, error)
	// Transfers the given literal to the Infeed buffer of the device.
	TransferToInfeed(ctx context.Context, in *TransferToInfeedRequest, opts ...grpc.CallOption) (*TransferToInfeedResponse, error)
	// Transferred literal from the Outfeed buffer of the device.
	TransferFromOutfeed(ctx context.Context, in *TransferFromOutfeedRequest, opts ...grpc.CallOption) (*TransferFromOutfeedResponse, error)
	// Resets the device, clearing all existing state on the device.
	ResetDevice(ctx context.Context, in *ResetDeviceRequest, opts ...grpc.CallOption) (*ResetDeviceResponse, error)
	// Computes the value of a constant expression. The request contains the
	// computation graph for the constant expression.
	ComputeConstantGraph(ctx context.Context, in *ComputeConstantGraphRequest, opts ...grpc.CallOption) (*ComputeConstantResponse, error)
	// Requests one or more device handles from the target. The returned device
	// handles can be used to specify the device on which to execute computations
	// or transfer data.
	GetDeviceHandles(ctx context.Context, in *GetDeviceHandlesRequest, opts ...grpc.CallOption) (*GetDeviceHandlesResponse, error)
	// Creates a channel handle that can be used to transfer data between
	// two computations via a pair of Send and Recv instructions.
	CreateChannelHandle(ctx context.Context, in *CreateChannelHandleRequest, opts ...grpc.CallOption) (*CreateChannelHandleResponse, error)
	// Compiles the provided computation into executable. Returns the handle of
	// the executable.
	Compile(ctx context.Context, in *CompileRequest, opts ...grpc.CallOption) (*CompileResponse, error)
	// Invokes the provided executable with the provided global data passed as
	// immutable arguments. The request contains the handle to the executable.
	// Returns global data output and execution timing.
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
	// Invokes the provided list of computations in parallel with the provided
	// global data for each computation. Returns a list of global data output and
	// execution timing.
	ExecuteGraphParallel(ctx context.Context, in *ExecuteGraphParallelRequest, opts ...grpc.CallOption) (*ExecuteParallelResponse, error)
	// Waits until the given execution (aysnchronously launched) is complete, and
	// returns the global data output.
	WaitForExecution(ctx context.Context, in *WaitForExecutionRequest, opts ...grpc.CallOption) (*WaitForExecutionResponse, error)
}

type xlaServiceClient struct {
	cc *grpc.ClientConn
}

func NewXlaServiceClient(cc *grpc.ClientConn) XlaServiceClient {
	return &xlaServiceClient{cc}
}

func (c *xlaServiceClient) Unregister(ctx context.Context, in *UnregisterRequest, opts ...grpc.CallOption) (*UnregisterResponse, error) {
	out := new(UnregisterResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/Unregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) DeconstructTuple(ctx context.Context, in *DeconstructTupleRequest, opts ...grpc.CallOption) (*DeconstructTupleResponse, error) {
	out := new(DeconstructTupleResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/DeconstructTuple", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) Unpack(ctx context.Context, in *UnpackRequest, opts ...grpc.CallOption) (*UnpackResponse, error) {
	out := new(UnpackResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/Unpack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) GetShape(ctx context.Context, in *GetShapeRequest, opts ...grpc.CallOption) (*GetShapeResponse, error) {
	out := new(GetShapeResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/GetShape", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) GetComputationGraphStats(ctx context.Context, in *ComputationGraphStatsRequest, opts ...grpc.CallOption) (*ComputationStatsResponse, error) {
	out := new(ComputationStatsResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/GetComputationGraphStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) LoadData(ctx context.Context, in *LoadDataRequest, opts ...grpc.CallOption) (*LoadDataResponse, error) {
	out := new(LoadDataResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/LoadData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) TransferToClient(ctx context.Context, in *TransferToClientRequest, opts ...grpc.CallOption) (*TransferToClientResponse, error) {
	out := new(TransferToClientResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/TransferToClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) TransferToServer(ctx context.Context, in *TransferToServerRequest, opts ...grpc.CallOption) (*TransferToServerResponse, error) {
	out := new(TransferToServerResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/TransferToServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) TransferToInfeed(ctx context.Context, in *TransferToInfeedRequest, opts ...grpc.CallOption) (*TransferToInfeedResponse, error) {
	out := new(TransferToInfeedResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/TransferToInfeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) TransferFromOutfeed(ctx context.Context, in *TransferFromOutfeedRequest, opts ...grpc.CallOption) (*TransferFromOutfeedResponse, error) {
	out := new(TransferFromOutfeedResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/TransferFromOutfeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) ResetDevice(ctx context.Context, in *ResetDeviceRequest, opts ...grpc.CallOption) (*ResetDeviceResponse, error) {
	out := new(ResetDeviceResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/ResetDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) ComputeConstantGraph(ctx context.Context, in *ComputeConstantGraphRequest, opts ...grpc.CallOption) (*ComputeConstantResponse, error) {
	out := new(ComputeConstantResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/ComputeConstantGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) GetDeviceHandles(ctx context.Context, in *GetDeviceHandlesRequest, opts ...grpc.CallOption) (*GetDeviceHandlesResponse, error) {
	out := new(GetDeviceHandlesResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/GetDeviceHandles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) CreateChannelHandle(ctx context.Context, in *CreateChannelHandleRequest, opts ...grpc.CallOption) (*CreateChannelHandleResponse, error) {
	out := new(CreateChannelHandleResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/CreateChannelHandle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) Compile(ctx context.Context, in *CompileRequest, opts ...grpc.CallOption) (*CompileResponse, error) {
	out := new(CompileResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/Compile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) ExecuteGraphParallel(ctx context.Context, in *ExecuteGraphParallelRequest, opts ...grpc.CallOption) (*ExecuteParallelResponse, error) {
	out := new(ExecuteParallelResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/ExecuteGraphParallel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xlaServiceClient) WaitForExecution(ctx context.Context, in *WaitForExecutionRequest, opts ...grpc.CallOption) (*WaitForExecutionResponse, error) {
	out := new(WaitForExecutionResponse)
	err := c.cc.Invoke(ctx, "/xla.XlaService/WaitForExecution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XlaServiceServer is the server API for XlaService service.
type XlaServiceServer interface {
	// Unregisters a global allocation.
	//
	// If the handle given is not currently allocated, a NOT_FOUND status is
	// returned.
	Unregister(context.Context, *UnregisterRequest) (*UnregisterResponse, error)
	// Deconstructs a tuple. Returns a newly created GlobalDataHandle for each
	// element in the tuple.
	DeconstructTuple(context.Context, *DeconstructTupleRequest) (*DeconstructTupleResponse, error)
	// Unpack requests that a global data handle, with a tuple shape, has global
	// data handles created for each of its constituent members. This is the
	// equivalent of the "destructuring assignment" present in various programming
	// languages.
	Unpack(context.Context, *UnpackRequest) (*UnpackResponse, error)
	// Requests the shape of the referenced global data.
	GetShape(context.Context, *GetShapeRequest) (*GetShapeResponse, error)
	// Requests the statistics of the given computation.
	GetComputationGraphStats(context.Context, *ComputationGraphStatsRequest) (*ComputationStatsResponse, error)
	// Loads a variable number of values with a given element type from ColumnIO.
	LoadData(context.Context, *LoadDataRequest) (*LoadDataResponse, error)
	// Transfers the given global data to the client in the form of a Literal.
	TransferToClient(context.Context, *TransferToClientRequest) (*TransferToClientResponse, error)
	// Transfers the given literal to the server to be stored in a global
	// allocation, which is returned.
	TransferToServer(context.Context, *TransferToServerRequest) (*TransferToServerResponse, error)
	// Transfers the given literal to the Infeed buffer of the device.
	TransferToInfeed(context.Context, *TransferToInfeedRequest) (*TransferToInfeedResponse, error)
	// Transferred literal from the Outfeed buffer of the device.
	TransferFromOutfeed(context.Context, *TransferFromOutfeedRequest) (*TransferFromOutfeedResponse, error)
	// Resets the device, clearing all existing state on the device.
	ResetDevice(context.Context, *ResetDeviceRequest) (*ResetDeviceResponse, error)
	// Computes the value of a constant expression. The request contains the
	// computation graph for the constant expression.
	ComputeConstantGraph(context.Context, *ComputeConstantGraphRequest) (*ComputeConstantResponse, error)
	// Requests one or more device handles from the target. The returned device
	// handles can be used to specify the device on which to execute computations
	// or transfer data.
	GetDeviceHandles(context.Context, *GetDeviceHandlesRequest) (*GetDeviceHandlesResponse, error)
	// Creates a channel handle that can be used to transfer data between
	// two computations via a pair of Send and Recv instructions.
	CreateChannelHandle(context.Context, *CreateChannelHandleRequest) (*CreateChannelHandleResponse, error)
	// Compiles the provided computation into executable. Returns the handle of
	// the executable.
	Compile(context.Context, *CompileRequest) (*CompileResponse, error)
	// Invokes the provided executable with the provided global data passed as
	// immutable arguments. The request contains the handle to the executable.
	// Returns global data output and execution timing.
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
	// Invokes the provided list of computations in parallel with the provided
	// global data for each computation. Returns a list of global data output and
	// execution timing.
	ExecuteGraphParallel(context.Context, *ExecuteGraphParallelRequest) (*ExecuteParallelResponse, error)
	// Waits until the given execution (aysnchronously launched) is complete, and
	// returns the global data output.
	WaitForExecution(context.Context, *WaitForExecutionRequest) (*WaitForExecutionResponse, error)
}

// UnimplementedXlaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedXlaServiceServer struct {
}

func (*UnimplementedXlaServiceServer) Unregister(ctx context.Context, req *UnregisterRequest) (*UnregisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unregister not implemented")
}
func (*UnimplementedXlaServiceServer) DeconstructTuple(ctx context.Context, req *DeconstructTupleRequest) (*DeconstructTupleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeconstructTuple not implemented")
}
func (*UnimplementedXlaServiceServer) Unpack(ctx context.Context, req *UnpackRequest) (*UnpackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpack not implemented")
}
func (*UnimplementedXlaServiceServer) GetShape(ctx context.Context, req *GetShapeRequest) (*GetShapeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShape not implemented")
}
func (*UnimplementedXlaServiceServer) GetComputationGraphStats(ctx context.Context, req *ComputationGraphStatsRequest) (*ComputationStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComputationGraphStats not implemented")
}
func (*UnimplementedXlaServiceServer) LoadData(ctx context.Context, req *LoadDataRequest) (*LoadDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadData not implemented")
}
func (*UnimplementedXlaServiceServer) TransferToClient(ctx context.Context, req *TransferToClientRequest) (*TransferToClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferToClient not implemented")
}
func (*UnimplementedXlaServiceServer) TransferToServer(ctx context.Context, req *TransferToServerRequest) (*TransferToServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferToServer not implemented")
}
func (*UnimplementedXlaServiceServer) TransferToInfeed(ctx context.Context, req *TransferToInfeedRequest) (*TransferToInfeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferToInfeed not implemented")
}
func (*UnimplementedXlaServiceServer) TransferFromOutfeed(ctx context.Context, req *TransferFromOutfeedRequest) (*TransferFromOutfeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferFromOutfeed not implemented")
}
func (*UnimplementedXlaServiceServer) ResetDevice(ctx context.Context, req *ResetDeviceRequest) (*ResetDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetDevice not implemented")
}
func (*UnimplementedXlaServiceServer) ComputeConstantGraph(ctx context.Context, req *ComputeConstantGraphRequest) (*ComputeConstantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeConstantGraph not implemented")
}
func (*UnimplementedXlaServiceServer) GetDeviceHandles(ctx context.Context, req *GetDeviceHandlesRequest) (*GetDeviceHandlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceHandles not implemented")
}
func (*UnimplementedXlaServiceServer) CreateChannelHandle(ctx context.Context, req *CreateChannelHandleRequest) (*CreateChannelHandleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannelHandle not implemented")
}
func (*UnimplementedXlaServiceServer) Compile(ctx context.Context, req *CompileRequest) (*CompileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compile not implemented")
}
func (*UnimplementedXlaServiceServer) Execute(ctx context.Context, req *ExecuteRequest) (*ExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (*UnimplementedXlaServiceServer) ExecuteGraphParallel(ctx context.Context, req *ExecuteGraphParallelRequest) (*ExecuteParallelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteGraphParallel not implemented")
}
func (*UnimplementedXlaServiceServer) WaitForExecution(ctx context.Context, req *WaitForExecutionRequest) (*WaitForExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitForExecution not implemented")
}

func RegisterXlaServiceServer(s *grpc.Server, srv XlaServiceServer) {
	s.RegisterService(&_XlaService_serviceDesc, srv)
}

func _XlaService_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).Unregister(ctx, req.(*UnregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_DeconstructTuple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeconstructTupleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).DeconstructTuple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/DeconstructTuple",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).DeconstructTuple(ctx, req.(*DeconstructTupleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_Unpack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnpackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).Unpack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/Unpack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).Unpack(ctx, req.(*UnpackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_GetShape_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShapeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).GetShape(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/GetShape",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).GetShape(ctx, req.(*GetShapeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_GetComputationGraphStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputationGraphStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).GetComputationGraphStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/GetComputationGraphStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).GetComputationGraphStats(ctx, req.(*ComputationGraphStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_LoadData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).LoadData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/LoadData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).LoadData(ctx, req.(*LoadDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_TransferToClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferToClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).TransferToClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/TransferToClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).TransferToClient(ctx, req.(*TransferToClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_TransferToServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferToServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).TransferToServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/TransferToServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).TransferToServer(ctx, req.(*TransferToServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_TransferToInfeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferToInfeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).TransferToInfeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/TransferToInfeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).TransferToInfeed(ctx, req.(*TransferToInfeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_TransferFromOutfeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferFromOutfeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).TransferFromOutfeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/TransferFromOutfeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).TransferFromOutfeed(ctx, req.(*TransferFromOutfeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_ResetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).ResetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/ResetDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).ResetDevice(ctx, req.(*ResetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_ComputeConstantGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputeConstantGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).ComputeConstantGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/ComputeConstantGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).ComputeConstantGraph(ctx, req.(*ComputeConstantGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_GetDeviceHandles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceHandlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).GetDeviceHandles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/GetDeviceHandles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).GetDeviceHandles(ctx, req.(*GetDeviceHandlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_CreateChannelHandle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelHandleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).CreateChannelHandle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/CreateChannelHandle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).CreateChannelHandle(ctx, req.(*CreateChannelHandleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_Compile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).Compile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/Compile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).Compile(ctx, req.(*CompileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_ExecuteGraphParallel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteGraphParallelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).ExecuteGraphParallel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/ExecuteGraphParallel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).ExecuteGraphParallel(ctx, req.(*ExecuteGraphParallelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _XlaService_WaitForExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitForExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XlaServiceServer).WaitForExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xla.XlaService/WaitForExecution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XlaServiceServer).WaitForExecution(ctx, req.(*WaitForExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _XlaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "xla.XlaService",
	HandlerType: (*XlaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unregister",
			Handler:    _XlaService_Unregister_Handler,
		},
		{
			MethodName: "DeconstructTuple",
			Handler:    _XlaService_DeconstructTuple_Handler,
		},
		{
			MethodName: "Unpack",
			Handler:    _XlaService_Unpack_Handler,
		},
		{
			MethodName: "GetShape",
			Handler:    _XlaService_GetShape_Handler,
		},
		{
			MethodName: "GetComputationGraphStats",
			Handler:    _XlaService_GetComputationGraphStats_Handler,
		},
		{
			MethodName: "LoadData",
			Handler:    _XlaService_LoadData_Handler,
		},
		{
			MethodName: "TransferToClient",
			Handler:    _XlaService_TransferToClient_Handler,
		},
		{
			MethodName: "TransferToServer",
			Handler:    _XlaService_TransferToServer_Handler,
		},
		{
			MethodName: "TransferToInfeed",
			Handler:    _XlaService_TransferToInfeed_Handler,
		},
		{
			MethodName: "TransferFromOutfeed",
			Handler:    _XlaService_TransferFromOutfeed_Handler,
		},
		{
			MethodName: "ResetDevice",
			Handler:    _XlaService_ResetDevice_Handler,
		},
		{
			MethodName: "ComputeConstantGraph",
			Handler:    _XlaService_ComputeConstantGraph_Handler,
		},
		{
			MethodName: "GetDeviceHandles",
			Handler:    _XlaService_GetDeviceHandles_Handler,
		},
		{
			MethodName: "CreateChannelHandle",
			Handler:    _XlaService_CreateChannelHandle_Handler,
		},
		{
			MethodName: "Compile",
			Handler:    _XlaService_Compile_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _XlaService_Execute_Handler,
		},
		{
			MethodName: "ExecuteGraphParallel",
			Handler:    _XlaService_ExecuteGraphParallel_Handler,
		},
		{
			MethodName: "WaitForExecution",
			Handler:    _XlaService_WaitForExecution_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tensorflow/compiler/xla/rpc/xla_service.proto",
}
