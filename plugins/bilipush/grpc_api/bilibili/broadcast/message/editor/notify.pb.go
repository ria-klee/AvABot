// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_api/bilibili/broadcast/message/editor/notify.proto

package bilibili_broadcast_message_editor

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type Notify struct {
	// 消息唯一标示
	MsgId int64 `protobuf:"varint,1,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	// 消息类型
	MsgType int32 `protobuf:"varint,2,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	// 接收方uid
	ReceiverUid int64 `protobuf:"varint,3,opt,name=receiver_uid,json=receiverUid,proto3" json:"receiver_uid,omitempty"`
	//接收方类型
	ReceiverType int32 `protobuf:"varint,4,opt,name=receiver_type,json=receiverType,proto3" json:"receiver_type,omitempty"`
	// 故事的版本
	StoryVersion int64 `protobuf:"varint,5,opt,name=story_version,json=storyVersion,proto3" json:"story_version,omitempty"`
	// 操作结果的hash值
	OpHash int64 `protobuf:"varint,6,opt,name=op_hash,json=opHash,proto3" json:"op_hash,omitempty"`
	// 操作产生用户的uid
	OpSender int64 `protobuf:"varint,7,opt,name=op_sender,json=opSender,proto3" json:"op_sender,omitempty"`
	// patch内容
	OpContent            string   `protobuf:"bytes,8,opt,name=op_content,json=opContent,proto3" json:"op_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notify) Reset()         { *m = Notify{} }
func (m *Notify) String() string { return proto.CompactTextString(m) }
func (*Notify) ProtoMessage()    {}
func (*Notify) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab5e4139b4005581, []int{0}
}

func (m *Notify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notify.Unmarshal(m, b)
}
func (m *Notify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notify.Marshal(b, m, deterministic)
}
func (m *Notify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notify.Merge(m, src)
}
func (m *Notify) XXX_Size() int {
	return xxx_messageInfo_Notify.Size(m)
}
func (m *Notify) XXX_DiscardUnknown() {
	xxx_messageInfo_Notify.DiscardUnknown(m)
}

var xxx_messageInfo_Notify proto.InternalMessageInfo

func (m *Notify) GetMsgId() int64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *Notify) GetMsgType() int32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *Notify) GetReceiverUid() int64 {
	if m != nil {
		return m.ReceiverUid
	}
	return 0
}

func (m *Notify) GetReceiverType() int32 {
	if m != nil {
		return m.ReceiverType
	}
	return 0
}

func (m *Notify) GetStoryVersion() int64 {
	if m != nil {
		return m.StoryVersion
	}
	return 0
}

func (m *Notify) GetOpHash() int64 {
	if m != nil {
		return m.OpHash
	}
	return 0
}

func (m *Notify) GetOpSender() int64 {
	if m != nil {
		return m.OpSender
	}
	return 0
}

func (m *Notify) GetOpContent() string {
	if m != nil {
		return m.OpContent
	}
	return ""
}

func init() {
	proto.RegisterType((*Notify)(nil), "bilibili.broadcast.message.editor.Notify")
}

func init() {
	proto.RegisterFile("grpc_api/bilibili/broadcast/message/editor/notify.proto", fileDescriptor_ab5e4139b4005581)
}

var fileDescriptor_ab5e4139b4005581 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4f, 0x02, 0x31,
	0x18, 0x86, 0x73, 0x22, 0x07, 0x54, 0x88, 0x49, 0x13, 0xf5, 0x84, 0x98, 0x80, 0x2c, 0xb8, 0xb4,
	0x46, 0x07, 0x7f, 0x80, 0x31, 0xd1, 0x45, 0x13, 0x54, 0xd6, 0xa6, 0x77, 0xfd, 0x28, 0x4d, 0xb8,
	0xfb, 0x9a, 0xb6, 0x90, 0xdc, 0x7f, 0x77, 0x30, 0xf4, 0x80, 0xc5, 0xc1, 0xa1, 0x43, 0x9f, 0xef,
	0x7d, 0x9f, 0xe1, 0x25, 0x4f, 0xda, 0xd9, 0x42, 0x48, 0x6b, 0x78, 0x6e, 0xd6, 0x66, 0xf7, 0x78,
	0xee, 0x50, 0xaa, 0x42, 0xfa, 0xc0, 0x4b, 0xf0, 0x5e, 0x6a, 0xe0, 0xa0, 0x4c, 0x40, 0xc7, 0x2b,
	0x0c, 0x66, 0x59, 0x33, 0xeb, 0x30, 0x20, 0x9d, 0x1c, 0xf2, 0xec, 0x98, 0x67, 0xfb, 0x3c, 0x6b,
	0xf2, 0xc3, 0x91, 0x46, 0xd4, 0x6b, 0xe0, 0xb1, 0x90, 0x6f, 0x96, 0x1c, 0x4a, 0x1b, 0xf6, 0xfd,
	0xdb, 0x9f, 0x84, 0xa4, 0xef, 0x51, 0x48, 0x2f, 0x48, 0x5a, 0x7a, 0x2d, 0x8c, 0xca, 0x92, 0x71,
	0x32, 0x6b, 0xcd, 0xdb, 0xa5, 0xd7, 0x6f, 0x8a, 0x5e, 0x93, 0xee, 0x0e, 0x87, 0xda, 0x42, 0x76,
	0x32, 0x4e, 0x66, 0xed, 0x79, 0xa7, 0xf4, 0xfa, 0xab, 0xb6, 0x40, 0x27, 0xa4, 0xef, 0xa0, 0x00,
	0xb3, 0x05, 0x27, 0x36, 0x46, 0x65, 0xad, 0xd8, 0x3b, 0x3b, 0xb0, 0x6f, 0xa3, 0xe8, 0x94, 0x0c,
	0x8e, 0x91, 0xa8, 0x38, 0x8d, 0x8a, 0x63, 0x2f, 0x7a, 0xa6, 0x64, 0xe0, 0x03, 0xba, 0x5a, 0x6c,
	0xc1, 0x79, 0x83, 0x55, 0xd6, 0x8e, 0xa2, 0x7e, 0x84, 0x8b, 0x86, 0xd1, 0x2b, 0xd2, 0x41, 0x2b,
	0x56, 0xd2, 0xaf, 0xb2, 0x34, 0x9e, 0x53, 0xb4, 0xaf, 0xd2, 0xaf, 0xe8, 0x88, 0xf4, 0xd0, 0x0a,
	0x0f, 0x95, 0x02, 0x97, 0x75, 0xe2, 0xa9, 0x8b, 0xf6, 0x33, 0xfe, 0xe9, 0x0d, 0x21, 0x68, 0x45,
	0x81, 0x55, 0x80, 0x2a, 0x64, 0xdd, 0x71, 0x32, 0xeb, 0xcd, 0x7b, 0x68, 0x9f, 0x1b, 0xf0, 0x60,
	0xc8, 0xf9, 0x87, 0x05, 0x27, 0x83, 0xc1, 0x6a, 0x3f, 0xc3, 0xe2, 0x2f, 0xba, 0x64, 0xcd, 0x84,
	0xec, 0x30, 0x21, 0x7b, 0xd9, 0x4d, 0x38, 0xbc, 0x63, 0xff, 0xae, 0xcf, 0x1a, 0xc5, 0x7d, 0x92,
	0xa7, 0xb1, 0xfc, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x30, 0x82, 0xe0, 0xf3, 0xeb, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OperationNotifyClient is the client API for OperationNotify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperationNotifyClient interface {
	//
	OperationNotify(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (OperationNotify_OperationNotifyClient, error)
}

type operationNotifyClient struct {
	cc *grpc.ClientConn
}

func NewOperationNotifyClient(cc *grpc.ClientConn) OperationNotifyClient {
	return &operationNotifyClient{cc}
}

func (c *operationNotifyClient) OperationNotify(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (OperationNotify_OperationNotifyClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OperationNotify_serviceDesc.Streams[0], "/bilibili.broadcast.message.editor.OperationNotify/OperationNotify", opts...)
	if err != nil {
		return nil, err
	}
	x := &operationNotifyOperationNotifyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OperationNotify_OperationNotifyClient interface {
	Recv() (*Notify, error)
	grpc.ClientStream
}

type operationNotifyOperationNotifyClient struct {
	grpc.ClientStream
}

func (x *operationNotifyOperationNotifyClient) Recv() (*Notify, error) {
	m := new(Notify)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OperationNotifyServer is the server API for OperationNotify service.
type OperationNotifyServer interface {
	//
	OperationNotify(*emptypb.Empty, OperationNotify_OperationNotifyServer) error
}

// UnimplementedOperationNotifyServer can be embedded to have forward compatible implementations.
type UnimplementedOperationNotifyServer struct {
}

func (*UnimplementedOperationNotifyServer) OperationNotify(req *emptypb.Empty, srv OperationNotify_OperationNotifyServer) error {
	return status.Errorf(codes.Unimplemented, "method OperationNotify not implemented")
}

func RegisterOperationNotifyServer(s *grpc.Server, srv OperationNotifyServer) {
	s.RegisterService(&_OperationNotify_serviceDesc, srv)
}

func _OperationNotify_OperationNotify_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OperationNotifyServer).OperationNotify(m, &operationNotifyOperationNotifyServer{stream})
}

type OperationNotify_OperationNotifyServer interface {
	Send(*Notify) error
	grpc.ServerStream
}

type operationNotifyOperationNotifyServer struct {
	grpc.ServerStream
}

func (x *operationNotifyOperationNotifyServer) Send(m *Notify) error {
	return x.ServerStream.SendMsg(m)
}

var _OperationNotify_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bilibili.broadcast.message.editor.OperationNotify",
	HandlerType: (*OperationNotifyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OperationNotify",
			Handler:       _OperationNotify_OperationNotify_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc_api/bilibili/broadcast/message/editor/notify.proto",
}
