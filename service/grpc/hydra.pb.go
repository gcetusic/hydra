// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hydra.proto

package hydra

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Runtime int32

const (
	Runtime_GOLANG Runtime = 0
	Runtime_PYTHON Runtime = 1
)

var Runtime_name = map[int32]string{
	0: "GOLANG",
	1: "PYTHON",
}

var Runtime_value = map[string]int32{
	"GOLANG": 0,
	"PYTHON": 1,
}

func (x Runtime) String() string {
	return proto.EnumName(Runtime_name, int32(x))
}

func (Runtime) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4eeaa86bd8ab47d2, []int{0}
}

type Status int32

const (
	Status_READY   Status = 0
	Status_RUNNING Status = 1
)

var Status_name = map[int32]string{
	0: "READY",
	1: "RUNNING",
}

var Status_value = map[string]int32{
	"READY":   0,
	"RUNNING": 1,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4eeaa86bd8ab47d2, []int{1}
}

type Node struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Runtime              Runtime  `protobuf:"varint,2,opt,name=runtime,proto3,enum=Runtime" json:"runtime,omitempty"`
	Status               Status   `protobuf:"varint,3,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eeaa86bd8ab47d2, []int{0}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetRuntime() Runtime {
	if m != nil {
		return m.Runtime
	}
	return Runtime_GOLANG
}

func (m *Node) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_READY
}

type NodeListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeListRequest) Reset()         { *m = NodeListRequest{} }
func (m *NodeListRequest) String() string { return proto.CompactTextString(m) }
func (*NodeListRequest) ProtoMessage()    {}
func (*NodeListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eeaa86bd8ab47d2, []int{1}
}

func (m *NodeListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeListRequest.Unmarshal(m, b)
}
func (m *NodeListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeListRequest.Marshal(b, m, deterministic)
}
func (m *NodeListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeListRequest.Merge(m, src)
}
func (m *NodeListRequest) XXX_Size() int {
	return xxx_messageInfo_NodeListRequest.Size(m)
}
func (m *NodeListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeListRequest proto.InternalMessageInfo

type Event struct {
	// Types that are valid to be assigned to Message:
	//	*Event_NodeStatusMessage
	//	*Event_NodeRegistrationMessage
	Message              isEvent_Message `protobuf_oneof:"message"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eeaa86bd8ab47d2, []int{2}
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

type isEvent_Message interface {
	isEvent_Message()
}

type Event_NodeStatusMessage struct {
	NodeStatusMessage *NodeStatusMessage `protobuf:"bytes,1,opt,name=nodeStatusMessage,proto3,oneof"`
}

type Event_NodeRegistrationMessage struct {
	NodeRegistrationMessage *NodeRegistrationMessage `protobuf:"bytes,2,opt,name=nodeRegistrationMessage,proto3,oneof"`
}

func (*Event_NodeStatusMessage) isEvent_Message() {}

func (*Event_NodeRegistrationMessage) isEvent_Message() {}

func (m *Event) GetMessage() isEvent_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Event) GetNodeStatusMessage() *NodeStatusMessage {
	if x, ok := m.GetMessage().(*Event_NodeStatusMessage); ok {
		return x.NodeStatusMessage
	}
	return nil
}

func (m *Event) GetNodeRegistrationMessage() *NodeRegistrationMessage {
	if x, ok := m.GetMessage().(*Event_NodeRegistrationMessage); ok {
		return x.NodeRegistrationMessage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_NodeStatusMessage)(nil),
		(*Event_NodeRegistrationMessage)(nil),
	}
}

type NodeRegistrationMessage struct {
	Runtime              Runtime  `protobuf:"varint,1,opt,name=runtime,proto3,enum=Runtime" json:"runtime,omitempty"`
	Status               Status   `protobuf:"varint,2,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRegistrationMessage) Reset()         { *m = NodeRegistrationMessage{} }
func (m *NodeRegistrationMessage) String() string { return proto.CompactTextString(m) }
func (*NodeRegistrationMessage) ProtoMessage()    {}
func (*NodeRegistrationMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eeaa86bd8ab47d2, []int{3}
}

func (m *NodeRegistrationMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRegistrationMessage.Unmarshal(m, b)
}
func (m *NodeRegistrationMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRegistrationMessage.Marshal(b, m, deterministic)
}
func (m *NodeRegistrationMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRegistrationMessage.Merge(m, src)
}
func (m *NodeRegistrationMessage) XXX_Size() int {
	return xxx_messageInfo_NodeRegistrationMessage.Size(m)
}
func (m *NodeRegistrationMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRegistrationMessage.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRegistrationMessage proto.InternalMessageInfo

func (m *NodeRegistrationMessage) GetRuntime() Runtime {
	if m != nil {
		return m.Runtime
	}
	return Runtime_GOLANG
}

func (m *NodeRegistrationMessage) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_READY
}

type RuntimeStartRequest struct {
	Payload              string   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RuntimeStartRequest) Reset()         { *m = RuntimeStartRequest{} }
func (m *RuntimeStartRequest) String() string { return proto.CompactTextString(m) }
func (*RuntimeStartRequest) ProtoMessage()    {}
func (*RuntimeStartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eeaa86bd8ab47d2, []int{4}
}

func (m *RuntimeStartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuntimeStartRequest.Unmarshal(m, b)
}
func (m *RuntimeStartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuntimeStartRequest.Marshal(b, m, deterministic)
}
func (m *RuntimeStartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeStartRequest.Merge(m, src)
}
func (m *RuntimeStartRequest) XXX_Size() int {
	return xxx_messageInfo_RuntimeStartRequest.Size(m)
}
func (m *RuntimeStartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeStartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeStartRequest proto.InternalMessageInfo

func (m *RuntimeStartRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type NodeStatusMessage struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeStatusMessage) Reset()         { *m = NodeStatusMessage{} }
func (m *NodeStatusMessage) String() string { return proto.CompactTextString(m) }
func (*NodeStatusMessage) ProtoMessage()    {}
func (*NodeStatusMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eeaa86bd8ab47d2, []int{5}
}

func (m *NodeStatusMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStatusMessage.Unmarshal(m, b)
}
func (m *NodeStatusMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStatusMessage.Marshal(b, m, deterministic)
}
func (m *NodeStatusMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStatusMessage.Merge(m, src)
}
func (m *NodeStatusMessage) XXX_Size() int {
	return xxx_messageInfo_NodeStatusMessage.Size(m)
}
func (m *NodeStatusMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStatusMessage.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStatusMessage proto.InternalMessageInfo

func (m *NodeStatusMessage) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_READY
}

func init() {
	proto.RegisterEnum("Runtime", Runtime_name, Runtime_value)
	proto.RegisterEnum("Status", Status_name, Status_value)
	proto.RegisterType((*Node)(nil), "Node")
	proto.RegisterType((*NodeListRequest)(nil), "NodeListRequest")
	proto.RegisterType((*Event)(nil), "Event")
	proto.RegisterType((*NodeRegistrationMessage)(nil), "NodeRegistrationMessage")
	proto.RegisterType((*RuntimeStartRequest)(nil), "RuntimeStartRequest")
	proto.RegisterType((*NodeStatusMessage)(nil), "NodeStatusMessage")
}

func init() { proto.RegisterFile("hydra.proto", fileDescriptor_4eeaa86bd8ab47d2) }

var fileDescriptor_4eeaa86bd8ab47d2 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x5d, 0x6b, 0xe2, 0x40,
	0x14, 0xcd, 0x64, 0x37, 0xc9, 0xe6, 0x0a, 0x6e, 0xbc, 0xbb, 0x60, 0xf0, 0xa5, 0x36, 0x0f, 0x45,
	0x2c, 0xa4, 0x62, 0xfb, 0x07, 0x94, 0x8a, 0x29, 0xd8, 0x58, 0x46, 0xfb, 0x20, 0x2d, 0x85, 0xb4,
	0x19, 0x6c, 0xa0, 0x26, 0x36, 0x33, 0x29, 0xf5, 0x3f, 0xf5, 0x47, 0x96, 0x4c, 0x0c, 0x88, 0x46,
	0xfa, 0x36, 0xf7, 0xe3, 0x9c, 0x7b, 0xe7, 0xdc, 0x03, 0xb5, 0xd7, 0x4d, 0x98, 0x06, 0xee, 0x3a,
	0x4d, 0x44, 0xe2, 0x3c, 0xc0, 0x6f, 0x3f, 0x09, 0x19, 0xd6, 0x41, 0x8d, 0x42, 0x9b, 0xb4, 0x49,
	0xc7, 0xa4, 0x6a, 0x14, 0xa2, 0x03, 0x46, 0x9a, 0xc5, 0x22, 0x5a, 0x31, 0x5b, 0x6d, 0x93, 0x4e,
	0xbd, 0xff, 0xc7, 0xa5, 0x45, 0x4c, 0xcb, 0x02, 0x9e, 0x80, 0xce, 0x45, 0x20, 0x32, 0x6e, 0xff,
	0x92, 0x2d, 0x86, 0x3b, 0x93, 0x21, 0xdd, 0xa6, 0x9d, 0x06, 0xfc, 0xcd, 0xc9, 0x27, 0x11, 0x17,
	0x94, 0xbd, 0x67, 0x8c, 0x0b, 0xe7, 0x8b, 0x80, 0x36, 0xfa, 0x60, 0xb1, 0xc0, 0x21, 0x34, 0xe2,
	0x24, 0x64, 0x05, 0xe4, 0x96, 0x71, 0x1e, 0x2c, 0x99, 0x5c, 0xa0, 0xd6, 0x47, 0xd7, 0xdf, 0xaf,
	0x78, 0x0a, 0x3d, 0x6c, 0xc7, 0x39, 0x34, 0xf3, 0x24, 0x65, 0xcb, 0x88, 0x8b, 0x34, 0x10, 0x51,
	0x12, 0x97, 0x4c, 0xaa, 0x64, 0xb2, 0x25, 0x53, 0x45, 0xdd, 0x53, 0xe8, 0x31, 0xe8, 0xd0, 0x04,
	0x63, 0x55, 0x3c, 0x9d, 0x27, 0x68, 0x1e, 0x21, 0xd8, 0x55, 0x88, 0xfc, 0xac, 0x90, 0x5a, 0xad,
	0xd0, 0x05, 0xfc, 0xdb, 0x82, 0x66, 0x22, 0x48, 0x4b, 0x95, 0xd0, 0x06, 0x63, 0x1d, 0x6c, 0xde,
	0x92, 0xa0, 0x3c, 0x49, 0x19, 0x3a, 0x57, 0xd0, 0x38, 0xd0, 0x66, 0x67, 0x0c, 0xa9, 0x1c, 0xd3,
	0x3d, 0x05, 0x63, 0x3b, 0x06, 0x01, 0xf4, 0xf1, 0x74, 0x32, 0xf0, 0xc7, 0x96, 0x92, 0xbf, 0xef,
	0x16, 0x73, 0x6f, 0xea, 0x5b, 0xa4, 0xdb, 0x06, 0xbd, 0x00, 0xa1, 0x09, 0x1a, 0x1d, 0x0d, 0xae,
	0x17, 0x96, 0x82, 0x35, 0x30, 0xe8, 0xbd, 0xef, 0xdf, 0xf8, 0x63, 0x8b, 0xf4, 0x1f, 0x41, 0xf3,
	0x72, 0xe7, 0xe0, 0x39, 0x18, 0xa3, 0x4f, 0xf6, 0x92, 0x09, 0x86, 0xba, 0x2b, 0x8f, 0xd9, 0xfa,
	0xef, 0x56, 0x7c, 0xa3, 0x43, 0x7a, 0x04, 0xcf, 0xc0, 0xcc, 0xef, 0x9f, 0x2f, 0xcd, 0xd1, 0x72,
	0xf7, 0xfc, 0xd0, 0xd2, 0x64, 0xa6, 0x47, 0x9e, 0x75, 0xe9, 0xc7, 0xcb, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x21, 0xfd, 0x82, 0xee, 0x9e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HydraClient is the client API for Hydra service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HydraClient interface {
	Execute(ctx context.Context, opts ...grpc.CallOption) (Hydra_ExecuteClient, error)
	ListNodes(ctx context.Context, in *NodeListRequest, opts ...grpc.CallOption) (Hydra_ListNodesClient, error)
}

type hydraClient struct {
	cc *grpc.ClientConn
}

func NewHydraClient(cc *grpc.ClientConn) HydraClient {
	return &hydraClient{cc}
}

func (c *hydraClient) Execute(ctx context.Context, opts ...grpc.CallOption) (Hydra_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Hydra_serviceDesc.Streams[0], "/Hydra/Execute", opts...)
	if err != nil {
		return nil, err
	}
	x := &hydraExecuteClient{stream}
	return x, nil
}

type Hydra_ExecuteClient interface {
	Send(*Event) error
	Recv() (*RuntimeStartRequest, error)
	grpc.ClientStream
}

type hydraExecuteClient struct {
	grpc.ClientStream
}

func (x *hydraExecuteClient) Send(m *Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hydraExecuteClient) Recv() (*RuntimeStartRequest, error) {
	m := new(RuntimeStartRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hydraClient) ListNodes(ctx context.Context, in *NodeListRequest, opts ...grpc.CallOption) (Hydra_ListNodesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Hydra_serviceDesc.Streams[1], "/Hydra/ListNodes", opts...)
	if err != nil {
		return nil, err
	}
	x := &hydraListNodesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Hydra_ListNodesClient interface {
	Recv() (*Node, error)
	grpc.ClientStream
}

type hydraListNodesClient struct {
	grpc.ClientStream
}

func (x *hydraListNodesClient) Recv() (*Node, error) {
	m := new(Node)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HydraServer is the server API for Hydra service.
type HydraServer interface {
	Execute(Hydra_ExecuteServer) error
	ListNodes(*NodeListRequest, Hydra_ListNodesServer) error
}

func RegisterHydraServer(s *grpc.Server, srv HydraServer) {
	s.RegisterService(&_Hydra_serviceDesc, srv)
}

func _Hydra_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HydraServer).Execute(&hydraExecuteServer{stream})
}

type Hydra_ExecuteServer interface {
	Send(*RuntimeStartRequest) error
	Recv() (*Event, error)
	grpc.ServerStream
}

type hydraExecuteServer struct {
	grpc.ServerStream
}

func (x *hydraExecuteServer) Send(m *RuntimeStartRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hydraExecuteServer) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Hydra_ListNodes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NodeListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HydraServer).ListNodes(m, &hydraListNodesServer{stream})
}

type Hydra_ListNodesServer interface {
	Send(*Node) error
	grpc.ServerStream
}

type hydraListNodesServer struct {
	grpc.ServerStream
}

func (x *hydraListNodesServer) Send(m *Node) error {
	return x.ServerStream.SendMsg(m)
}

var _Hydra_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Hydra",
	HandlerType: (*HydraServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Execute",
			Handler:       _Hydra_Execute_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ListNodes",
			Handler:       _Hydra_ListNodes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hydra.proto",
}
