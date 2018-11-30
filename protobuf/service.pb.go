// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package delayQueue

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type QueueRequest struct {
	Messsage             string   `protobuf:"bytes,1,opt,name=messsage,proto3" json:"messsage,omitempty"`
	RetryCount           int32    `protobuf:"varint,2,opt,name=retryCount,proto3" json:"retryCount,omitempty"`
	Delay                int32    `protobuf:"varint,3,opt,name=delay,proto3" json:"delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueRequest) Reset()         { *m = QueueRequest{} }
func (m *QueueRequest) String() string { return proto.CompactTextString(m) }
func (*QueueRequest) ProtoMessage()    {}
func (*QueueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *QueueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueRequest.Unmarshal(m, b)
}
func (m *QueueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueRequest.Marshal(b, m, deterministic)
}
func (m *QueueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueRequest.Merge(m, src)
}
func (m *QueueRequest) XXX_Size() int {
	return xxx_messageInfo_QueueRequest.Size(m)
}
func (m *QueueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueueRequest proto.InternalMessageInfo

func (m *QueueRequest) GetMesssage() string {
	if m != nil {
		return m.Messsage
	}
	return ""
}

func (m *QueueRequest) GetRetryCount() int32 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

func (m *QueueRequest) GetDelay() int32 {
	if m != nil {
		return m.Delay
	}
	return 0
}

type QueueResponse struct {
	ReturnCode           int64    `protobuf:"varint,1,opt,name=returnCode,proto3" json:"returnCode,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueResponse) Reset()         { *m = QueueResponse{} }
func (m *QueueResponse) String() string { return proto.CompactTextString(m) }
func (*QueueResponse) ProtoMessage()    {}
func (*QueueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *QueueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueResponse.Unmarshal(m, b)
}
func (m *QueueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueResponse.Marshal(b, m, deterministic)
}
func (m *QueueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueResponse.Merge(m, src)
}
func (m *QueueResponse) XXX_Size() int {
	return xxx_messageInfo_QueueResponse.Size(m)
}
func (m *QueueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueueResponse proto.InternalMessageInfo

func (m *QueueResponse) GetReturnCode() int64 {
	if m != nil {
		return m.ReturnCode
	}
	return 0
}

func (m *QueueResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*QueueRequest)(nil), "delayQueue.QueueRequest")
	proto.RegisterType((*QueueResponse)(nil), "delayQueue.QueueResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0x49, 0xcd, 0x49, 0xac,
	0x0c, 0x2c, 0x4d, 0x2d, 0x4d, 0x55, 0x4a, 0xe0, 0xe2, 0x01, 0x33, 0x82, 0x52, 0x0b, 0x4b, 0x53,
	0x8b, 0x4b, 0x84, 0xa4, 0xb8, 0x38, 0x72, 0x53, 0x8b, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x21, 0x39, 0x2e, 0xae, 0xa2, 0xd4, 0x92, 0xa2, 0x4a,
	0xe7, 0xfc, 0xd2, 0xbc, 0x12, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x24, 0x11, 0x21, 0x11,
	0x2e, 0x56, 0xb0, 0xc9, 0x12, 0xcc, 0x60, 0x29, 0x08, 0x47, 0xc9, 0x93, 0x8b, 0x17, 0x6a, 0x43,
	0x71, 0x41, 0x7e, 0x5e, 0x31, 0xcc, 0x98, 0xd2, 0xa2, 0x3c, 0xe7, 0xfc, 0x14, 0x88, 0x25, 0xcc,
	0x41, 0x48, 0x22, 0x42, 0x12, 0x5c, 0xec, 0x20, 0x2b, 0x41, 0x2e, 0x60, 0x02, 0xbb, 0x00, 0xc6,
	0x35, 0xf2, 0xe1, 0xe2, 0x72, 0x81, 0x3b, 0x5d, 0xc8, 0x8e, 0x8b, 0x3d, 0xa0, 0x34, 0x29, 0x27,
	0xb3, 0x38, 0x43, 0x48, 0x42, 0x0f, 0xe1, 0x25, 0x3d, 0x64, 0xff, 0x48, 0x49, 0x62, 0x91, 0x81,
	0xb8, 0x23, 0x89, 0x0d, 0x1c, 0x1a, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0x5e, 0xcd,
	0x18, 0x1e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DelayQueueClient is the client API for DelayQueue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DelayQueueClient interface {
	Publish(ctx context.Context, in *QueueRequest, opts ...grpc.CallOption) (*QueueResponse, error)
}

type delayQueueClient struct {
	cc *grpc.ClientConn
}

func NewDelayQueueClient(cc *grpc.ClientConn) DelayQueueClient {
	return &delayQueueClient{cc}
}

func (c *delayQueueClient) Publish(ctx context.Context, in *QueueRequest, opts ...grpc.CallOption) (*QueueResponse, error) {
	out := new(QueueResponse)
	err := c.cc.Invoke(ctx, "/delayQueue.DelayQueue/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DelayQueueServer is the server API for DelayQueue service.
type DelayQueueServer interface {
	Publish(context.Context, *QueueRequest) (*QueueResponse, error)
}

func RegisterDelayQueueServer(s *grpc.Server, srv DelayQueueServer) {
	s.RegisterService(&_DelayQueue_serviceDesc, srv)
}

func _DelayQueue_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelayQueueServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/delayQueue.DelayQueue/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelayQueueServer).Publish(ctx, req.(*QueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DelayQueue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "delayQueue.DelayQueue",
	HandlerType: (*DelayQueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _DelayQueue_Publish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
