// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator.proto

package calculator

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CalculatorRequest struct {
	FirstNumber          float64  `protobuf:"fixed64,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         float64  `protobuf:"fixed64,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorRequest) Reset()         { *m = CalculatorRequest{} }
func (m *CalculatorRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorRequest) ProtoMessage()    {}
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_060f5a7357841f48, []int{0}
}
func (m *CalculatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorRequest.Unmarshal(m, b)
}
func (m *CalculatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorRequest.Marshal(b, m, deterministic)
}
func (dst *CalculatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorRequest.Merge(dst, src)
}
func (m *CalculatorRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorRequest.Size(m)
}
func (m *CalculatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorRequest proto.InternalMessageInfo

func (m *CalculatorRequest) GetFirstNumber() float64 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *CalculatorRequest) GetSecondNumber() float64 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type CalculatorResponse struct {
	Total                float64  `protobuf:"fixed64,1,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorResponse) Reset()         { *m = CalculatorResponse{} }
func (m *CalculatorResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorResponse) ProtoMessage()    {}
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_060f5a7357841f48, []int{1}
}
func (m *CalculatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorResponse.Unmarshal(m, b)
}
func (m *CalculatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorResponse.Marshal(b, m, deterministic)
}
func (dst *CalculatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorResponse.Merge(dst, src)
}
func (m *CalculatorResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorResponse.Size(m)
}
func (m *CalculatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorResponse proto.InternalMessageInfo

func (m *CalculatorResponse) GetTotal() float64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type PrimeNumberDecompositionRequest struct {
	N                    uint64   `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecompositionRequest) Reset()         { *m = PrimeNumberDecompositionRequest{} }
func (m *PrimeNumberDecompositionRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionRequest) ProtoMessage()    {}
func (*PrimeNumberDecompositionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_060f5a7357841f48, []int{2}
}
func (m *PrimeNumberDecompositionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Unmarshal(m, b)
}
func (m *PrimeNumberDecompositionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Marshal(b, m, deterministic)
}
func (dst *PrimeNumberDecompositionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecompositionRequest.Merge(dst, src)
}
func (m *PrimeNumberDecompositionRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecompositionRequest.Size(m)
}
func (m *PrimeNumberDecompositionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecompositionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecompositionRequest proto.InternalMessageInfo

func (m *PrimeNumberDecompositionRequest) GetN() uint64 {
	if m != nil {
		return m.N
	}
	return 0
}

type PrimeNumberDecompositionResponse struct {
	Result               uint64   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecompositionResponse) Reset()         { *m = PrimeNumberDecompositionResponse{} }
func (m *PrimeNumberDecompositionResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecompositionResponse) ProtoMessage()    {}
func (*PrimeNumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_060f5a7357841f48, []int{3}
}
func (m *PrimeNumberDecompositionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Unmarshal(m, b)
}
func (m *PrimeNumberDecompositionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Marshal(b, m, deterministic)
}
func (dst *PrimeNumberDecompositionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecompositionResponse.Merge(dst, src)
}
func (m *PrimeNumberDecompositionResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecompositionResponse.Size(m)
}
func (m *PrimeNumberDecompositionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecompositionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecompositionResponse proto.InternalMessageInfo

func (m *PrimeNumberDecompositionResponse) GetResult() uint64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type AverageRuequest struct {
	Number               uint64   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AverageRuequest) Reset()         { *m = AverageRuequest{} }
func (m *AverageRuequest) String() string { return proto.CompactTextString(m) }
func (*AverageRuequest) ProtoMessage()    {}
func (*AverageRuequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_060f5a7357841f48, []int{4}
}
func (m *AverageRuequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AverageRuequest.Unmarshal(m, b)
}
func (m *AverageRuequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AverageRuequest.Marshal(b, m, deterministic)
}
func (dst *AverageRuequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AverageRuequest.Merge(dst, src)
}
func (m *AverageRuequest) XXX_Size() int {
	return xxx_messageInfo_AverageRuequest.Size(m)
}
func (m *AverageRuequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AverageRuequest.DiscardUnknown(m)
}

var xxx_messageInfo_AverageRuequest proto.InternalMessageInfo

func (m *AverageRuequest) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type AverageResponse struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AverageResponse) Reset()         { *m = AverageResponse{} }
func (m *AverageResponse) String() string { return proto.CompactTextString(m) }
func (*AverageResponse) ProtoMessage()    {}
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_060f5a7357841f48, []int{5}
}
func (m *AverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AverageResponse.Unmarshal(m, b)
}
func (m *AverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AverageResponse.Marshal(b, m, deterministic)
}
func (dst *AverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AverageResponse.Merge(dst, src)
}
func (m *AverageResponse) XXX_Size() int {
	return xxx_messageInfo_AverageResponse.Size(m)
}
func (m *AverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AverageResponse proto.InternalMessageInfo

func (m *AverageResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*CalculatorRequest)(nil), "CalculatorRequest")
	proto.RegisterType((*CalculatorResponse)(nil), "CalculatorResponse")
	proto.RegisterType((*PrimeNumberDecompositionRequest)(nil), "PrimeNumberDecompositionRequest")
	proto.RegisterType((*PrimeNumberDecompositionResponse)(nil), "PrimeNumberDecompositionResponse")
	proto.RegisterType((*AverageRuequest)(nil), "AverageRuequest")
	proto.RegisterType((*AverageResponse)(nil), "AverageResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	// Unary
	Add(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	// Server streaming
	PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error)
	// Client streaming
	Average(ctx context.Context, in *AverageRuequest, opts ...grpc.CallOption) (*AverageResponse, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Add(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/CalculatorService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumberDecomposition(ctx context.Context, in *PrimeNumberDecompositionRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/CalculatorService/PrimeNumberDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberDecompositionClient interface {
	Recv() (*PrimeNumberDecompositionResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberDecompositionClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberDecompositionClient) Recv() (*PrimeNumberDecompositionResponse, error) {
	m := new(PrimeNumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Average(ctx context.Context, in *AverageRuequest, opts ...grpc.CallOption) (*AverageResponse, error) {
	out := new(AverageResponse)
	err := c.cc.Invoke(ctx, "/CalculatorService/Average", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// Unary
	Add(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
	// Server streaming
	PrimeNumberDecomposition(*PrimeNumberDecompositionRequest, CalculatorService_PrimeNumberDecompositionServer) error
	// Client streaming
	Average(context.Context, *AverageRuequest) (*AverageResponse, error)
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CalculatorService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Add(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumberDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumberDecomposition(m, &calculatorServicePrimeNumberDecompositionServer{stream})
}

type CalculatorService_PrimeNumberDecompositionServer interface {
	Send(*PrimeNumberDecompositionResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberDecompositionServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberDecompositionServer) Send(m *PrimeNumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_Average_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AverageRuequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Average(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CalculatorService/Average",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Average(ctx, req.(*AverageRuequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalculatorService_Add_Handler,
		},
		{
			MethodName: "Average",
			Handler:    _CalculatorService_Average_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecomposition",
			Handler:       _CalculatorService_PrimeNumberDecomposition_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calculator.proto",
}

func init() { proto.RegisterFile("calculator.proto", fileDescriptor_calculator_060f5a7357841f48) }

var fileDescriptor_calculator_060f5a7357841f48 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x14, 0x84, 0xbb, 0x5a, 0x23, 0x3c, 0x23, 0xd6, 0xa7, 0x94, 0x92, 0x8b, 0xe9, 0x7a, 0x51, 0x0f,
	0xdb, 0xa2, 0x37, 0x6f, 0x45, 0xcf, 0x22, 0xf1, 0xe8, 0x41, 0xd2, 0xcd, 0x53, 0x02, 0xc9, 0x6e,
	0xdc, 0xdd, 0xf4, 0xcf, 0xfa, 0x67, 0x84, 0x64, 0x53, 0x83, 0x6d, 0xe9, 0xf1, 0x0d, 0x33, 0x93,
	0xcc, 0xc7, 0xc2, 0x48, 0xa6, 0x85, 0xac, 0x8b, 0xd4, 0x69, 0x23, 0x2a, 0xa3, 0x9d, 0xe6, 0xef,
	0x70, 0xfe, 0xb4, 0xd6, 0x12, 0xfa, 0xae, 0xc9, 0x3a, 0x9c, 0x42, 0xf8, 0x99, 0x1b, 0xeb, 0x3e,
	0x54, 0x5d, 0x2e, 0xc9, 0x4c, 0x58, 0xcc, 0x6e, 0x58, 0x72, 0xd2, 0x68, 0x2f, 0x8d, 0x84, 0xd7,
	0x70, 0x6a, 0x49, 0x6a, 0x95, 0x75, 0x9e, 0x83, 0xc6, 0x13, 0xb6, 0x62, 0x6b, 0xe2, 0x77, 0x80,
	0xfd, 0x72, 0x5b, 0x69, 0x65, 0x09, 0x2f, 0xe1, 0xc8, 0x69, 0x97, 0x16, 0xbe, 0xb6, 0x3d, 0xf8,
	0x0c, 0xae, 0x5e, 0x4d, 0x5e, 0x52, 0x1b, 0x7d, 0x26, 0xa9, 0xcb, 0x4a, 0xdb, 0xdc, 0xe5, 0x5a,
	0x75, 0xbf, 0x15, 0x02, 0x53, 0x4d, 0x68, 0x98, 0x30, 0xc5, 0x1f, 0x21, 0xde, 0x1d, 0xf0, 0x9f,
	0x1a, 0x43, 0x60, 0xc8, 0xd6, 0x85, 0xf3, 0x31, 0x7f, 0xf1, 0x5b, 0x38, 0x5b, 0xac, 0xc8, 0xa4,
	0x5f, 0x94, 0xd4, 0xbe, 0x7c, 0x0c, 0x41, 0x6f, 0xed, 0x30, 0xf1, 0x57, 0xdf, 0xba, 0xbd, 0x95,
	0x75, 0xad, 0xf7, 0x3f, 0xac, 0x0f, 0xf3, 0x8d, 0xcc, 0x2a, 0x97, 0x84, 0x73, 0x38, 0x5c, 0x64,
	0x19, 0xa2, 0xd8, 0xe0, 0x1c, 0x5d, 0x88, 0x4d, 0x3c, 0x7c, 0x80, 0x12, 0x26, 0xbb, 0x96, 0x61,
	0x2c, 0xf6, 0x50, 0x8a, 0xa6, 0x62, 0x1f, 0x16, 0x3e, 0x98, 0x33, 0x9c, 0xc1, 0xb1, 0xdf, 0x85,
	0x23, 0xf1, 0x0f, 0x46, 0xf4, 0xa7, 0xac, 0x23, 0xcb, 0xa0, 0x79, 0x30, 0x0f, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x39, 0x59, 0x8a, 0x4b, 0x44, 0x02, 0x00, 0x00,
}
