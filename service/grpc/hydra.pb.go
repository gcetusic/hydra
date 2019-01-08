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
	// Types that are valid to be assigned to Request:
	//	*Event_NodeStatusRequest
	//	*Event_NodeRegistrationRequest
	Request              isEvent_Request `protobuf_oneof:"request"`
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

type isEvent_Request interface {
	isEvent_Request()
}

type Event_NodeStatusRequest struct {
	NodeStatusRequest *NodeStatusRequest `protobuf:"bytes,1,opt,name=nodeStatusRequest,proto3,oneof"`
}

type Event_NodeRegistrationRequest struct {
	NodeRegistrationRequest *NodeRegistrationRequest `protobuf:"bytes,2,opt,name=nodeRegistrationRequest,proto3,oneof"`
}

func (*Event_NodeStatusRequest) isEvent_Request() {}

func (*Event_NodeRegistrationRequest) isEvent_Request() {}

func (m *Event) GetRequest() isEvent_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Event) GetNodeStatusRequest() *NodeStatusRequest {
	if x, ok := m.GetRequest().(*Event_NodeStatusRequest); ok {
		return x.NodeStatusRequest
	}
	return nil
}

func (m *Event) GetNodeRegistrationRequest() *NodeRegistrationRequest {
	if x, ok := m.GetRequest().(*Event_NodeRegistrationRequest); ok {
		return x.NodeRegistrationRequest
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_NodeStatusRequest)(nil),
		(*Event_NodeRegistrationRequest)(nil),
	}
}

type NodeRegistrationRequest struct {
	Runtime              Runtime  `protobuf:"varint,1,opt,name=runtime,proto3,enum=Runtime" json:"runtime,omitempty"`
	Status               Status   `protobuf:"varint,2,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRegistrationRequest) Reset()         { *m = NodeRegistrationRequest{} }
func (m *NodeRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*NodeRegistrationRequest) ProtoMessage()    {}
func (*NodeRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eeaa86bd8ab47d2, []int{3}
}

func (m *NodeRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRegistrationRequest.Unmarshal(m, b)
}
func (m *NodeRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRegistrationRequest.Marshal(b, m, deterministic)
}
func (m *NodeRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRegistrationRequest.Merge(m, src)
}
func (m *NodeRegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_NodeRegistrationRequest.Size(m)
}
func (m *NodeRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRegistrationRequest proto.InternalMessageInfo

func (m *NodeRegistrationRequest) GetRuntime() Runtime {
	if m != nil {
		return m.Runtime
	}
	return Runtime_GOLANG
}

func (m *NodeRegistrationRequest) GetStatus() Status {
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

type NodeStatusRequest struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeStatusRequest) Reset()         { *m = NodeStatusRequest{} }
func (m *NodeStatusRequest) String() string { return proto.CompactTextString(m) }
func (*NodeStatusRequest) ProtoMessage()    {}
func (*NodeStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eeaa86bd8ab47d2, []int{5}
}

func (m *NodeStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStatusRequest.Unmarshal(m, b)
}
func (m *NodeStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStatusRequest.Marshal(b, m, deterministic)
}
func (m *NodeStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStatusRequest.Merge(m, src)
}
func (m *NodeStatusRequest) XXX_Size() int {
	return xxx_messageInfo_NodeStatusRequest.Size(m)
}
func (m *NodeStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStatusRequest proto.InternalMessageInfo

func (m *NodeStatusRequest) GetStatus() Status {
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
	proto.RegisterType((*NodeRegistrationRequest)(nil), "NodeRegistrationRequest")
	proto.RegisterType((*RuntimeStartRequest)(nil), "RuntimeStartRequest")
	proto.RegisterType((*NodeStatusRequest)(nil), "NodeStatusRequest")
}

func init() { proto.RegisterFile("hydra.proto", fileDescriptor_4eeaa86bd8ab47d2) }

var fileDescriptor_4eeaa86bd8ab47d2 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0xcd, 0x46, 0x93, 0x98, 0x29, 0xd4, 0x74, 0x14, 0x1a, 0x7a, 0xb1, 0xe6, 0x20, 0xa5, 0x42,
	0x2c, 0xd5, 0x1f, 0x68, 0xb1, 0x34, 0x42, 0xd9, 0xca, 0xb6, 0x1e, 0x8a, 0x22, 0x44, 0xb3, 0x68,
	0x40, 0x93, 0x9a, 0x6c, 0xc4, 0xfe, 0x93, 0x1f, 0x29, 0xd9, 0x26, 0x50, 0xda, 0x14, 0x6f, 0x3b,
	0xf3, 0xe6, 0xbd, 0x99, 0x7d, 0x33, 0x50, 0x7b, 0x5f, 0x05, 0x89, 0xef, 0x2e, 0x93, 0x58, 0xc4,
	0xce, 0x23, 0x1c, 0xd2, 0x38, 0xe0, 0x58, 0x07, 0x35, 0x0c, 0x6c, 0xd2, 0x26, 0x1d, 0x93, 0xa9,
	0x61, 0x80, 0x0e, 0x18, 0x49, 0x16, 0x89, 0xf0, 0x93, 0xdb, 0x6a, 0x9b, 0x74, 0xea, 0xfd, 0x23,
	0x97, 0xad, 0x63, 0x56, 0x02, 0x78, 0x06, 0x7a, 0x2a, 0x7c, 0x91, 0xa5, 0xf6, 0x81, 0x2c, 0x31,
	0xdc, 0x99, 0x0c, 0x59, 0x91, 0x76, 0x1a, 0x70, 0x9c, 0x8b, 0x4f, 0xc2, 0x54, 0x30, 0xfe, 0x95,
	0xf1, 0x54, 0x38, 0xbf, 0x04, 0xb4, 0xd1, 0x37, 0x8f, 0x04, 0x0e, 0xa1, 0x11, 0xc5, 0x01, 0x2f,
	0x28, 0x6b, 0x58, 0x0e, 0x50, 0xeb, 0xa3, 0x4b, 0xb7, 0x11, 0x4f, 0x61, 0xbb, 0xe5, 0x38, 0x87,
	0x66, 0x9e, 0x64, 0xfc, 0x2d, 0x4c, 0x45, 0xe2, 0x8b, 0x30, 0x8e, 0x4a, 0x25, 0x55, 0x2a, 0xd9,
	0x52, 0xa9, 0x02, 0xf7, 0x14, 0xb6, 0x8f, 0x3a, 0x34, 0xc1, 0x48, 0x8a, 0x71, 0x9f, 0xa1, 0xb9,
	0x47, 0x60, 0xd3, 0x21, 0xf2, 0xbf, 0x43, 0x6a, 0xb5, 0x43, 0x57, 0x70, 0x52, 0x90, 0x66, 0xc2,
	0x4f, 0x4a, 0x97, 0xd0, 0x06, 0x63, 0xe9, 0xaf, 0x3e, 0x62, 0xbf, 0x5c, 0x49, 0x19, 0x3a, 0x37,
	0xd0, 0xd8, 0xf1, 0x66, 0xa3, 0x0d, 0xa9, 0x6c, 0xd3, 0x3d, 0x07, 0xa3, 0x68, 0x83, 0x00, 0xfa,
	0x78, 0x3a, 0x19, 0xd0, 0xb1, 0xa5, 0xe4, 0xef, 0xfb, 0xc5, 0xdc, 0x9b, 0x52, 0x8b, 0x74, 0xdb,
	0xa0, 0xaf, 0x49, 0x68, 0x82, 0xc6, 0x46, 0x83, 0xdb, 0x85, 0xa5, 0x60, 0x0d, 0x0c, 0xf6, 0x40,
	0xe9, 0x1d, 0x1d, 0x5b, 0xa4, 0xff, 0x04, 0x9a, 0x97, 0x5f, 0x0e, 0x5e, 0x82, 0x31, 0xfa, 0xe1,
	0xaf, 0x99, 0xe0, 0xa8, 0xbb, 0x72, 0x99, 0xad, 0x53, 0xb7, 0xe2, 0x1b, 0x1d, 0xd2, 0x23, 0x78,
	0x01, 0x66, 0xbe, 0xff, 0x7c, 0xe8, 0x14, 0x2d, 0x77, 0xeb, 0x1e, 0x5a, 0x9a, 0xcc, 0xf4, 0xc8,
	0x8b, 0x2e, 0xef, 0xf1, 0xfa, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x89, 0x4a, 0xf2, 0x69, 0x9e, 0x02,
	0x00, 0x00,
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