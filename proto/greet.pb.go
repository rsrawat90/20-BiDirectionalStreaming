// Code generated by protoc-gen-go. DO NOT EDIT.
// source: 04-BiDirectionalStreaming/proto/greet.proto

package greetpb

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

type Greeting struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_00dbce7ca4bd9b49, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetEveryOneRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetEveryOneRequest) Reset()         { *m = GreetEveryOneRequest{} }
func (m *GreetEveryOneRequest) String() string { return proto.CompactTextString(m) }
func (*GreetEveryOneRequest) ProtoMessage()    {}
func (*GreetEveryOneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00dbce7ca4bd9b49, []int{1}
}

func (m *GreetEveryOneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetEveryOneRequest.Unmarshal(m, b)
}
func (m *GreetEveryOneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetEveryOneRequest.Marshal(b, m, deterministic)
}
func (m *GreetEveryOneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetEveryOneRequest.Merge(m, src)
}
func (m *GreetEveryOneRequest) XXX_Size() int {
	return xxx_messageInfo_GreetEveryOneRequest.Size(m)
}
func (m *GreetEveryOneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetEveryOneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetEveryOneRequest proto.InternalMessageInfo

func (m *GreetEveryOneRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetEveryOneResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetEveryOneResponse) Reset()         { *m = GreetEveryOneResponse{} }
func (m *GreetEveryOneResponse) String() string { return proto.CompactTextString(m) }
func (*GreetEveryOneResponse) ProtoMessage()    {}
func (*GreetEveryOneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00dbce7ca4bd9b49, []int{2}
}

func (m *GreetEveryOneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetEveryOneResponse.Unmarshal(m, b)
}
func (m *GreetEveryOneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetEveryOneResponse.Marshal(b, m, deterministic)
}
func (m *GreetEveryOneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetEveryOneResponse.Merge(m, src)
}
func (m *GreetEveryOneResponse) XXX_Size() int {
	return xxx_messageInfo_GreetEveryOneResponse.Size(m)
}
func (m *GreetEveryOneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetEveryOneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetEveryOneResponse proto.InternalMessageInfo

func (m *GreetEveryOneResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetEveryOneRequest)(nil), "greet.GreetEveryOneRequest")
	proto.RegisterType((*GreetEveryOneResponse)(nil), "greet.GreetEveryOneResponse")
}

func init() {
	proto.RegisterFile("04-BiDirectionalStreaming/proto/greet.proto", fileDescriptor_00dbce7ca4bd9b49)
}

var fileDescriptor_00dbce7ca4bd9b49 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x8d, 0xd0, 0xba, 0x19, 0x0f, 0x42, 0x50, 0x11, 0xab, 0x20, 0x39, 0x15, 0x8a, 0xdd,
	0x52, 0xfd, 0x04, 0xf5, 0xef, 0x49, 0x61, 0x7b, 0xf3, 0xa2, 0x69, 0x19, 0x43, 0x60, 0x37, 0x59,
	0x27, 0x69, 0xc1, 0x6f, 0x2f, 0x3b, 0x4d, 0x45, 0x8a, 0xb7, 0x99, 0xf7, 0x26, 0xef, 0xf7, 0x08,
	0x8c, 0x26, 0xb7, 0xd7, 0x33, 0x77, 0xef, 0x08, 0x97, 0xc9, 0x05, 0x6f, 0xea, 0x79, 0x22, 0x34,
	0x8d, 0xf3, 0xb6, 0x6c, 0x29, 0xa4, 0x50, 0x5a, 0x42, 0x4c, 0x63, 0x9e, 0x55, 0x8f, 0x17, 0xfd,
	0x08, 0xc5, 0x53, 0x37, 0x38, 0x6f, 0xd5, 0x25, 0xc0, 0xa7, 0xa3, 0x98, 0xde, 0xbd, 0x69, 0xf0,
	0x4c, 0x5c, 0x89, 0xa1, 0xac, 0x24, 0x2b, 0x2f, 0xa6, 0x41, 0x35, 0x00, 0x59, 0x9b, 0xad, 0xbb,
	0xcf, 0x6e, 0xd1, 0x09, 0x9d, 0xa9, 0xef, 0xe0, 0x98, 0x73, 0x1e, 0xd6, 0x48, 0xdf, 0xaf, 0x1e,
	0x2b, 0xfc, 0x5a, 0x61, 0x4c, 0x6a, 0x04, 0x85, 0xcd, 0xf9, 0x9c, 0x78, 0x38, 0x3d, 0x1a, 0x6f,
	0x6a, 0x6c, 0xb1, 0xd5, 0xef, 0x81, 0x2e, 0xe1, 0x64, 0x27, 0x24, 0xb6, 0xc1, 0x47, 0x54, 0xa7,
	0xd0, 0x27, 0x8c, 0xab, 0x3a, 0xe5, 0x56, 0x79, 0x9b, 0x7e, 0x64, 0x2a, 0xe6, 0x07, 0x73, 0xa4,
	0xb5, 0x5b, 0xa2, 0x7a, 0x86, 0x1e, 0xeb, 0x6a, 0xf0, 0x17, 0xb6, 0xd3, 0xed, 0xfc, 0xe2, 0x7f,
	0x73, 0xc3, 0xd4, 0x7b, 0x43, 0x31, 0x11, 0x33, 0xf9, 0x76, 0xc0, 0x47, 0xed, 0x62, 0xd1, 0xe7,
	0x8f, 0xbb, 0xf9, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x69, 0xdb, 0x2c, 0x67, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeteveryOneServiceClient is the client API for GreeteveryOneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeteveryOneServiceClient interface {
	Greet(ctx context.Context, opts ...grpc.CallOption) (GreeteveryOneService_GreetClient, error)
}

type greeteveryOneServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreeteveryOneServiceClient(cc *grpc.ClientConn) GreeteveryOneServiceClient {
	return &greeteveryOneServiceClient{cc}
}

func (c *greeteveryOneServiceClient) Greet(ctx context.Context, opts ...grpc.CallOption) (GreeteveryOneService_GreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreeteveryOneService_serviceDesc.Streams[0], "/greet.GreeteveryOneService/Greet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeteveryOneServiceGreetClient{stream}
	return x, nil
}

type GreeteveryOneService_GreetClient interface {
	Send(*GreetEveryOneRequest) error
	Recv() (*GreetEveryOneResponse, error)
	grpc.ClientStream
}

type greeteveryOneServiceGreetClient struct {
	grpc.ClientStream
}

func (x *greeteveryOneServiceGreetClient) Send(m *GreetEveryOneRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeteveryOneServiceGreetClient) Recv() (*GreetEveryOneResponse, error) {
	m := new(GreetEveryOneResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeteveryOneServiceServer is the server API for GreeteveryOneService service.
type GreeteveryOneServiceServer interface {
	Greet(GreeteveryOneService_GreetServer) error
}

// UnimplementedGreeteveryOneServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreeteveryOneServiceServer struct {
}

func (*UnimplementedGreeteveryOneServiceServer) Greet(srv GreeteveryOneService_GreetServer) error {
	return status.Errorf(codes.Unimplemented, "method Greet not implemented")
}

func RegisterGreeteveryOneServiceServer(s *grpc.Server, srv GreeteveryOneServiceServer) {
	s.RegisterService(&_GreeteveryOneService_serviceDesc, srv)
}

func _GreeteveryOneService_Greet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeteveryOneServiceServer).Greet(&greeteveryOneServiceGreetServer{stream})
}

type GreeteveryOneService_GreetServer interface {
	Send(*GreetEveryOneResponse) error
	Recv() (*GreetEveryOneRequest, error)
	grpc.ServerStream
}

type greeteveryOneServiceGreetServer struct {
	grpc.ServerStream
}

func (x *greeteveryOneServiceGreetServer) Send(m *GreetEveryOneResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeteveryOneServiceGreetServer) Recv() (*GreetEveryOneRequest, error) {
	m := new(GreetEveryOneRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GreeteveryOneService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreeteveryOneService",
	HandlerType: (*GreeteveryOneServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Greet",
			Handler:       _GreeteveryOneService_Greet_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "04-BiDirectionalStreaming/proto/greet.proto",
}
