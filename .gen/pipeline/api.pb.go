// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package pipeline

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ProcessEntry_ResourceType int32

const (
	ProcessEntry_CLUSTER ProcessEntry_ResourceType = 0
)

var ProcessEntry_ResourceType_name = map[int32]string{
	0: "CLUSTER",
}

var ProcessEntry_ResourceType_value = map[string]int32{
	"CLUSTER": 0,
}

func (x ProcessEntry_ResourceType) String() string {
	return proto.EnumName(ProcessEntry_ResourceType_name, int32(x))
}

func (ProcessEntry_ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0, 0}
}

type ProcessEntry_Status int32

const (
	ProcessEntry_SUCCESS  ProcessEntry_Status = 0
	ProcessEntry_FAILURED ProcessEntry_Status = 1
)

var ProcessEntry_Status_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILURED",
}

var ProcessEntry_Status_value = map[string]int32{
	"SUCCESS":  0,
	"FAILURED": 1,
}

func (x ProcessEntry_Status) String() string {
	return proto.EnumName(ProcessEntry_Status_name, int32(x))
}

func (ProcessEntry_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0, 1}
}

// The request message containing the process name.
type ProcessEntry struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId             string               `protobuf:"bytes,2,opt,name=parentId,proto3" json:"parentId,omitempty"`
	OrgId                uint32               `protobuf:"varint,3,opt,name=orgId,proto3" json:"orgId,omitempty"`
	Name                 string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	ResourceType         string               `protobuf:"bytes,5,opt,name=resourceType,proto3" json:"resourceType,omitempty"`
	ResourceId           string               `protobuf:"bytes,6,opt,name=resourceId,proto3" json:"resourceId,omitempty"`
	Status               string               `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	StartedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	FinishedAt           *timestamp.Timestamp `protobuf:"bytes,9,opt,name=finishedAt,proto3" json:"finishedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ProcessEntry) Reset()         { *m = ProcessEntry{} }
func (m *ProcessEntry) String() string { return proto.CompactTextString(m) }
func (*ProcessEntry) ProtoMessage()    {}
func (*ProcessEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *ProcessEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessEntry.Unmarshal(m, b)
}
func (m *ProcessEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessEntry.Marshal(b, m, deterministic)
}
func (m *ProcessEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessEntry.Merge(m, src)
}
func (m *ProcessEntry) XXX_Size() int {
	return xxx_messageInfo_ProcessEntry.Size(m)
}
func (m *ProcessEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessEntry proto.InternalMessageInfo

func (m *ProcessEntry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ProcessEntry) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func (m *ProcessEntry) GetOrgId() uint32 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *ProcessEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProcessEntry) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *ProcessEntry) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *ProcessEntry) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ProcessEntry) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *ProcessEntry) GetFinishedAt() *timestamp.Timestamp {
	if m != nil {
		return m.FinishedAt
	}
	return nil
}

// The response message containing the greetings
type ProcessEntryResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessEntryResponse) Reset()         { *m = ProcessEntryResponse{} }
func (m *ProcessEntryResponse) String() string { return proto.CompactTextString(m) }
func (*ProcessEntryResponse) ProtoMessage()    {}
func (*ProcessEntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *ProcessEntryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessEntryResponse.Unmarshal(m, b)
}
func (m *ProcessEntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessEntryResponse.Marshal(b, m, deterministic)
}
func (m *ProcessEntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessEntryResponse.Merge(m, src)
}
func (m *ProcessEntryResponse) XXX_Size() int {
	return xxx_messageInfo_ProcessEntryResponse.Size(m)
}
func (m *ProcessEntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessEntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessEntryResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("api.ProcessEntry_ResourceType", ProcessEntry_ResourceType_name, ProcessEntry_ResourceType_value)
	proto.RegisterEnum("api.ProcessEntry_Status", ProcessEntry_Status_name, ProcessEntry_Status_value)
	proto.RegisterType((*ProcessEntry)(nil), "api.ProcessEntry")
	proto.RegisterType((*ProcessEntryResponse)(nil), "api.ProcessEntryResponse")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x0f, 0xd2, 0x30,
	0x14, 0xc6, 0xd9, 0x80, 0x8d, 0x3d, 0xa6, 0xc1, 0x86, 0x90, 0x3a, 0x0d, 0x92, 0x79, 0xe1, 0x34,
	0x12, 0xb8, 0x18, 0x0f, 0x26, 0x88, 0x33, 0x59, 0xc2, 0xc1, 0x6c, 0xf0, 0x07, 0x94, 0xad, 0xcc,
	0x26, 0x6c, 0x6d, 0xda, 0xee, 0x80, 0x7f, 0xb9, 0x47, 0x43, 0xc7, 0x74, 0xc6, 0x83, 0xb7, 0x7e,
	0xdf, 0xfb, 0x7d, 0xc9, 0xfb, 0xfa, 0xc0, 0x23, 0x82, 0x45, 0x42, 0x72, 0xcd, 0xd1, 0x90, 0x08,
	0x16, 0xbc, 0x2b, 0x39, 0x2f, 0x6f, 0x74, 0x63, 0xac, 0x4b, 0x73, 0xdd, 0x68, 0x56, 0x51, 0xa5,
	0x49, 0x25, 0x5a, 0x2a, 0xfc, 0x69, 0x83, 0xff, 0x4d, 0xf2, 0x9c, 0x2a, 0x15, 0xd7, 0x5a, 0xde,
	0xd1, 0x4b, 0xb0, 0x59, 0x81, 0xad, 0x95, 0xb5, 0xf6, 0x52, 0x9b, 0x15, 0x28, 0x80, 0x89, 0x20,
	0x92, 0xd6, 0x3a, 0x29, 0xb0, 0x6d, 0xdc, 0xdf, 0x1a, 0xcd, 0x61, 0xcc, 0x65, 0x99, 0x14, 0x78,
	0xb8, 0xb2, 0xd6, 0x2f, 0xd2, 0x56, 0x20, 0x04, 0xa3, 0x9a, 0x54, 0x14, 0x8f, 0x0c, 0x6d, 0xde,
	0x28, 0x04, 0x5f, 0x52, 0xc5, 0x1b, 0x99, 0xd3, 0xd3, 0x5d, 0x50, 0x3c, 0x36, 0xb3, 0xbf, 0x3c,
	0xb4, 0x04, 0xe8, 0x74, 0x52, 0x60, 0xc7, 0x10, 0x3d, 0x07, 0x2d, 0xc0, 0x51, 0x9a, 0xe8, 0x46,
	0x61, 0xd7, 0xcc, 0x9e, 0x0a, 0x7d, 0x00, 0x4f, 0x69, 0x22, 0x35, 0x2d, 0xf6, 0x1a, 0x4f, 0x56,
	0xd6, 0x7a, 0xba, 0x0d, 0xa2, 0xb6, 0x77, 0xd4, 0xf5, 0x8e, 0x4e, 0x5d, 0xef, 0xf4, 0x0f, 0x8c,
	0x3e, 0x02, 0x5c, 0x59, 0xcd, 0xd4, 0x77, 0x13, 0xf5, 0xfe, 0x1b, 0xed, 0xd1, 0xe1, 0x1b, 0xf0,
	0xd3, 0xfe, 0xf6, 0x53, 0x70, 0x0f, 0xc7, 0x73, 0x76, 0x8a, 0xd3, 0xd9, 0x20, 0x7c, 0x0f, 0x4e,
	0xd6, 0x2e, 0x37, 0x05, 0x37, 0x3b, 0x1f, 0x0e, 0x71, 0x96, 0xcd, 0x06, 0xc8, 0x87, 0xc9, 0xd7,
	0x7d, 0x72, 0x3c, 0xa7, 0xf1, 0x97, 0x99, 0x15, 0x2e, 0x60, 0xde, 0xff, 0xf9, 0x94, 0x2a, 0xc1,
	0x6b, 0x45, 0xb7, 0x9f, 0xc0, 0x7d, 0xfa, 0x68, 0x07, 0xc3, 0x23, 0x2f, 0xd1, 0xab, 0xe8, 0x71,
	0xd6, 0x3e, 0x1c, 0xbc, 0xfe, 0xc7, 0xea, 0xf2, 0x9f, 0x97, 0xf0, 0x36, 0xe7, 0x55, 0x74, 0x21,
	0xf5, 0x0f, 0xc2, 0xf2, 0x1b, 0x6f, 0x8a, 0x48, 0x30, 0x41, 0x6f, 0xac, 0xa6, 0x8f, 0xd0, 0xc5,
	0x31, 0xcd, 0x76, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x98, 0xf4, 0xd8, 0x14, 0x2c, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProcessClient is the client API for Process service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessClient interface {
	// Adds a process entry
	Log(ctx context.Context, in *ProcessEntry, opts ...grpc.CallOption) (*ProcessEntryResponse, error)
}

type processClient struct {
	cc *grpc.ClientConn
}

func NewProcessClient(cc *grpc.ClientConn) ProcessClient {
	return &processClient{cc}
}

func (c *processClient) Log(ctx context.Context, in *ProcessEntry, opts ...grpc.CallOption) (*ProcessEntryResponse, error) {
	out := new(ProcessEntryResponse)
	err := c.cc.Invoke(ctx, "/api.Process/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessServer is the server API for Process service.
type ProcessServer interface {
	// Adds a process entry
	Log(context.Context, *ProcessEntry) (*ProcessEntryResponse, error)
}

// UnimplementedProcessServer can be embedded to have forward compatible implementations.
type UnimplementedProcessServer struct {
}

func (*UnimplementedProcessServer) Log(ctx context.Context, req *ProcessEntry) (*ProcessEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}

func RegisterProcessServer(s *grpc.Server, srv ProcessServer) {
	s.RegisterService(&_Process_serviceDesc, srv)
}

func _Process_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Process/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessServer).Log(ctx, req.(*ProcessEntry))
	}
	return interceptor(ctx, in, info, handler)
}

var _Process_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Process",
	HandlerType: (*ProcessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _Process_Log_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
