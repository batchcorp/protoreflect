// Code generated by protoc-gen-go. DO NOT EDIT.
// source: desc_test_proto3.proto

package testprotos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	pkg "github.com/batchcorp/protoreflect/internal/testprotos/pkg"
	grpc "google.golang.org/grpc"
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

type Proto3Enum int32

const (
	Proto3Enum_UNKNOWN Proto3Enum = 0
	Proto3Enum_VALUE1  Proto3Enum = 1
	Proto3Enum_VALUE2  Proto3Enum = 2
)

var Proto3Enum_name = map[int32]string{
	0: "UNKNOWN",
	1: "VALUE1",
	2: "VALUE2",
}

var Proto3Enum_value = map[string]int32{
	"UNKNOWN": 0,
	"VALUE1":  1,
	"VALUE2":  2,
}

func (x Proto3Enum) String() string {
	return proto.EnumName(Proto3Enum_name, int32(x))
}

func (Proto3Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9cbf5a38d930f32c, []int{0}
}

type TestRequest struct {
	Foo                  []Proto3Enum                                    `protobuf:"varint,1,rep,packed,name=foo,proto3,enum=testprotos.Proto3Enum" json:"foo,omitempty"`
	Bar                  string                                          `protobuf:"bytes,2,opt,name=bar,proto3" json:"bar,omitempty"`
	Baz                  *TestMessage                                    `protobuf:"bytes,3,opt,name=baz,proto3" json:"baz,omitempty"`
	Snafu                *TestMessage_NestedMessage_AnotherNestedMessage `protobuf:"bytes,4,opt,name=snafu,proto3" json:"snafu,omitempty"`
	Flags                map[string]bool                                 `protobuf:"bytes,5,rep,name=flags,proto3" json:"flags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Others               map[string]*TestMessage                         `protobuf:"bytes,6,rep,name=others,proto3" json:"others,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cbf5a38d930f32c, []int{0}
}

func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRequest.Unmarshal(m, b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRequest.Marshal(b, m, deterministic)
}
func (m *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(m, src)
}
func (m *TestRequest) XXX_Size() int {
	return xxx_messageInfo_TestRequest.Size(m)
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

func (m *TestRequest) GetFoo() []Proto3Enum {
	if m != nil {
		return m.Foo
	}
	return nil
}

func (m *TestRequest) GetBar() string {
	if m != nil {
		return m.Bar
	}
	return ""
}

func (m *TestRequest) GetBaz() *TestMessage {
	if m != nil {
		return m.Baz
	}
	return nil
}

func (m *TestRequest) GetSnafu() *TestMessage_NestedMessage_AnotherNestedMessage {
	if m != nil {
		return m.Snafu
	}
	return nil
}

func (m *TestRequest) GetFlags() map[string]bool {
	if m != nil {
		return m.Flags
	}
	return nil
}

func (m *TestRequest) GetOthers() map[string]*TestMessage {
	if m != nil {
		return m.Others
	}
	return nil
}

type TestResponse struct {
	Atm                  *AnotherTestMessage `protobuf:"bytes,1,opt,name=atm,proto3" json:"atm,omitempty"`
	Vs                   []int32             `protobuf:"varint,2,rep,packed,name=vs,proto3" json:"vs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TestResponse) Reset()         { *m = TestResponse{} }
func (m *TestResponse) String() string { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()    {}
func (*TestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cbf5a38d930f32c, []int{1}
}

func (m *TestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResponse.Unmarshal(m, b)
}
func (m *TestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResponse.Marshal(b, m, deterministic)
}
func (m *TestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResponse.Merge(m, src)
}
func (m *TestResponse) XXX_Size() int {
	return xxx_messageInfo_TestResponse.Size(m)
}
func (m *TestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestResponse proto.InternalMessageInfo

func (m *TestResponse) GetAtm() *AnotherTestMessage {
	if m != nil {
		return m.Atm
	}
	return nil
}

func (m *TestResponse) GetVs() []int32 {
	if m != nil {
		return m.Vs
	}
	return nil
}

func init() {
	proto.RegisterEnum("testprotos.Proto3Enum", Proto3Enum_name, Proto3Enum_value)
	proto.RegisterType((*TestRequest)(nil), "testprotos.TestRequest")
	proto.RegisterMapType((map[string]bool)(nil), "testprotos.TestRequest.FlagsEntry")
	proto.RegisterMapType((map[string]*TestMessage)(nil), "testprotos.TestRequest.OthersEntry")
	proto.RegisterType((*TestResponse)(nil), "testprotos.TestResponse")
}

func init() { proto.RegisterFile("desc_test_proto3.proto", fileDescriptor_9cbf5a38d930f32c) }

var fileDescriptor_9cbf5a38d930f32c = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x5f, 0x8f, 0xd2, 0x4c,
	0x14, 0xc6, 0xdf, 0xb6, 0x2f, 0xa8, 0xa7, 0x66, 0x6d, 0x4e, 0xcc, 0xd2, 0x34, 0xc6, 0x10, 0xbc,
	0xa9, 0x26, 0x16, 0xe8, 0xde, 0x90, 0xf5, 0x8a, 0xcd, 0xc2, 0xcd, 0x2a, 0x4b, 0xba, 0xae, 0x26,
	0xde, 0x6c, 0x0a, 0x7b, 0x28, 0x08, 0xed, 0xe0, 0xcc, 0x94, 0x84, 0xfd, 0x6c, 0x7e, 0x03, 0xbf,
	0x94, 0xe9, 0x0c, 0x2e, 0x5d, 0xb5, 0x78, 0xd5, 0xf9, 0xf3, 0x3c, 0xbf, 0x73, 0xe6, 0x99, 0x49,
	0xe1, 0xf8, 0x96, 0xc4, 0xf4, 0x46, 0x92, 0x90, 0x37, 0x6b, 0xce, 0x24, 0x3b, 0x09, 0xd4, 0x07,
	0xa1, 0x58, 0x52, 0x43, 0xe1, 0x39, 0xf7, 0x9a, 0xae, 0xde, 0xf5, 0x1a, 0xeb, 0x65, 0xd2, 0x2e,
	0x39, 0x97, 0x89, 0xde, 0x68, 0xfd, 0xb0, 0xc0, 0xfe, 0x48, 0x42, 0x46, 0xf4, 0x2d, 0x27, 0x21,
	0xd1, 0x07, 0x6b, 0xc6, 0x98, 0x6b, 0x34, 0x2d, 0xff, 0x28, 0x3c, 0x0e, 0xf6, 0xd0, 0x60, 0xac,
	0xaa, 0x0d, 0xb2, 0x3c, 0x8d, 0x0a, 0x09, 0x3a, 0x60, 0x4d, 0x62, 0xee, 0x9a, 0x4d, 0xc3, 0x7f,
	0x12, 0x15, 0x43, 0x7c, 0x5d, 0xac, 0xdc, 0xb9, 0x56, 0xd3, 0xf0, 0xed, 0xb0, 0x51, 0xf6, 0x16,
	0x15, 0x3e, 0x90, 0x10, 0x71, 0x42, 0x85, 0xf4, 0x0e, 0xc7, 0x50, 0x13, 0x59, 0x3c, 0xcb, 0xdd,
	0xff, 0x95, 0xf8, 0xb4, 0x42, 0x1c, 0x8c, 0x48, 0x48, 0xba, 0xfd, 0x35, 0xeb, 0x67, 0x4c, 0xce,
	0x89, 0x3f, 0x58, 0x8c, 0x34, 0x08, 0x7b, 0x50, 0x9b, 0xad, 0xe2, 0x44, 0xb8, 0xb5, 0xa6, 0xe5,
	0xdb, 0x61, 0xeb, 0x77, 0xe2, 0xee, 0x80, 0xc1, 0xb0, 0x10, 0x0d, 0x32, 0xc9, 0xb7, 0x91, 0x36,
	0xe0, 0x3b, 0xa8, 0x2b, 0xac, 0x70, 0xeb, 0xca, 0xfa, 0xaa, 0xca, 0x7a, 0xa9, 0x54, 0xda, 0xbb,
	0xb3, 0x78, 0x3d, 0x80, 0x3d, 0xb1, 0xc8, 0x64, 0x49, 0x5b, 0xd7, 0xd0, 0x99, 0x2c, 0x69, 0x8b,
	0xcf, 0xa1, 0xb6, 0x89, 0x57, 0x39, 0xa9, 0x9c, 0x1e, 0x47, 0x7a, 0x72, 0x6a, 0xf6, 0x0c, 0x2f,
	0x02, 0xbb, 0x04, 0xfc, 0x8b, 0xf5, 0x6d, 0xd9, 0x7a, 0x20, 0xd0, 0x3d, 0xb3, 0x35, 0x86, 0xa7,
	0xba, 0x61, 0xb1, 0x66, 0x99, 0x20, 0xec, 0x80, 0x15, 0xcb, 0x54, 0x41, 0xed, 0xf0, 0x65, 0x19,
	0xb0, 0x8b, 0xf2, 0xc1, 0xc5, 0xc4, 0x32, 0xc5, 0x23, 0x30, 0x37, 0xc2, 0x35, 0x9b, 0x96, 0x5f,
	0x8b, 0xcc, 0x8d, 0x78, 0xd3, 0x05, 0xd8, 0x5f, 0x3c, 0xda, 0xf0, 0xe8, 0x7a, 0x74, 0x31, 0xba,
	0xfc, 0x3c, 0x72, 0xfe, 0x43, 0x80, 0xfa, 0xa7, 0xfe, 0xfb, 0xeb, 0x41, 0xd7, 0x31, 0xee, 0xc7,
	0xa1, 0x63, 0x86, 0xdf, 0x4d, 0xfd, 0xa4, 0xae, 0x88, 0x6f, 0x16, 0x53, 0xc2, 0x73, 0xb0, 0xcf,
	0xd9, 0x15, 0x4b, 0x49, 0xce, 0x17, 0x59, 0x82, 0x8d, 0x8a, 0x78, 0xbd, 0x17, 0xc1, 0xd7, 0x79,
	0x9e, 0xae, 0xf5, 0xc3, 0xe4, 0x34, 0x5b, 0xd1, 0x54, 0x06, 0xc5, 0x9b, 0x0d, 0xce, 0x62, 0x8e,
	0x43, 0x78, 0x56, 0xa2, 0x0c, 0x56, 0x82, 0xb0, 0x2a, 0x11, 0xcf, 0xfd, 0xb3, 0x84, 0x0e, 0xc4,
	0x37, 0x30, 0x02, 0xa7, 0xc4, 0xe9, 0x27, 0xf1, 0x22, 0xc3, 0x83, 0x95, 0xbd, 0x7f, 0xe4, 0xd6,
	0x31, 0xf0, 0x02, 0xb0, 0xc4, 0x1c, 0x32, 0x4e, 0x1b, 0xe2, 0xd5, 0x07, 0x3d, 0xd0, 0x5e, 0xc7,
	0x38, 0x3b, 0xf9, 0xd2, 0x4d, 0x16, 0x72, 0x9e, 0x4f, 0x82, 0x29, 0x4b, 0xdb, 0xaa, 0xb1, 0x76,
	0xb9, 0xb1, 0xf6, 0x22, 0x93, 0xc4, 0xb3, 0x78, 0xd5, 0xde, 0x53, 0x26, 0x75, 0xfd, 0x2f, 0xf8,
	0x19, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x55, 0xc0, 0x6b, 0x1e, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	DoSomething(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*pkg.Bar, error)
	DoSomethingElse(ctx context.Context, opts ...grpc.CallOption) (TestService_DoSomethingElseClient, error)
	DoSomethingAgain(ctx context.Context, in *pkg.Bar, opts ...grpc.CallOption) (TestService_DoSomethingAgainClient, error)
	DoSomethingForever(ctx context.Context, opts ...grpc.CallOption) (TestService_DoSomethingForeverClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) DoSomething(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*pkg.Bar, error) {
	out := new(pkg.Bar)
	err := c.cc.Invoke(ctx, "/testprotos.TestService/DoSomething", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) DoSomethingElse(ctx context.Context, opts ...grpc.CallOption) (TestService_DoSomethingElseClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[0], "/testprotos.TestService/DoSomethingElse", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceDoSomethingElseClient{stream}
	return x, nil
}

type TestService_DoSomethingElseClient interface {
	Send(*TestMessage) error
	CloseAndRecv() (*TestResponse, error)
	grpc.ClientStream
}

type testServiceDoSomethingElseClient struct {
	grpc.ClientStream
}

func (x *testServiceDoSomethingElseClient) Send(m *TestMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceDoSomethingElseClient) CloseAndRecv() (*TestResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) DoSomethingAgain(ctx context.Context, in *pkg.Bar, opts ...grpc.CallOption) (TestService_DoSomethingAgainClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[1], "/testprotos.TestService/DoSomethingAgain", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceDoSomethingAgainClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_DoSomethingAgainClient interface {
	Recv() (*AnotherTestMessage, error)
	grpc.ClientStream
}

type testServiceDoSomethingAgainClient struct {
	grpc.ClientStream
}

func (x *testServiceDoSomethingAgainClient) Recv() (*AnotherTestMessage, error) {
	m := new(AnotherTestMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) DoSomethingForever(ctx context.Context, opts ...grpc.CallOption) (TestService_DoSomethingForeverClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[2], "/testprotos.TestService/DoSomethingForever", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceDoSomethingForeverClient{stream}
	return x, nil
}

type TestService_DoSomethingForeverClient interface {
	Send(*TestRequest) error
	Recv() (*TestResponse, error)
	grpc.ClientStream
}

type testServiceDoSomethingForeverClient struct {
	grpc.ClientStream
}

func (x *testServiceDoSomethingForeverClient) Send(m *TestRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceDoSomethingForeverClient) Recv() (*TestResponse, error) {
	m := new(TestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	DoSomething(context.Context, *TestRequest) (*pkg.Bar, error)
	DoSomethingElse(TestService_DoSomethingElseServer) error
	DoSomethingAgain(*pkg.Bar, TestService_DoSomethingAgainServer) error
	DoSomethingForever(TestService_DoSomethingForeverServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_DoSomething_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).DoSomething(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testprotos.TestService/DoSomething",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).DoSomething(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_DoSomethingElse_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).DoSomethingElse(&testServiceDoSomethingElseServer{stream})
}

type TestService_DoSomethingElseServer interface {
	SendAndClose(*TestResponse) error
	Recv() (*TestMessage, error)
	grpc.ServerStream
}

type testServiceDoSomethingElseServer struct {
	grpc.ServerStream
}

func (x *testServiceDoSomethingElseServer) SendAndClose(m *TestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceDoSomethingElseServer) Recv() (*TestMessage, error) {
	m := new(TestMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_DoSomethingAgain_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(pkg.Bar)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).DoSomethingAgain(m, &testServiceDoSomethingAgainServer{stream})
}

type TestService_DoSomethingAgainServer interface {
	Send(*AnotherTestMessage) error
	grpc.ServerStream
}

type testServiceDoSomethingAgainServer struct {
	grpc.ServerStream
}

func (x *testServiceDoSomethingAgainServer) Send(m *AnotherTestMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_DoSomethingForever_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).DoSomethingForever(&testServiceDoSomethingForeverServer{stream})
}

type TestService_DoSomethingForeverServer interface {
	Send(*TestResponse) error
	Recv() (*TestRequest, error)
	grpc.ServerStream
}

type testServiceDoSomethingForeverServer struct {
	grpc.ServerStream
}

func (x *testServiceDoSomethingForeverServer) Send(m *TestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceDoSomethingForeverServer) Recv() (*TestRequest, error) {
	m := new(TestRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testprotos.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoSomething",
			Handler:    _TestService_DoSomething_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DoSomethingElse",
			Handler:       _TestService_DoSomethingElse_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DoSomethingAgain",
			Handler:       _TestService_DoSomethingAgain_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DoSomethingForever",
			Handler:       _TestService_DoSomethingForever_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "desc_test_proto3.proto",
}
