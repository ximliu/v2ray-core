package command

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

type GetStatsRequest struct {
	// Name of the stat counter.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Whether or not to reset the counter to fetching its value.
	Reset_               bool     `protobuf:"varint,2,opt,name=reset,proto3" json:"reset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatsRequest) Reset()         { *m = GetStatsRequest{} }
func (m *GetStatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatsRequest) ProtoMessage()    {}
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{0}
}

func (m *GetStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsRequest.Unmarshal(m, b)
}
func (m *GetStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsRequest.Marshal(b, m, deterministic)
}
func (m *GetStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsRequest.Merge(m, src)
}
func (m *GetStatsRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatsRequest.Size(m)
}
func (m *GetStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsRequest proto.InternalMessageInfo

func (m *GetStatsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetStatsRequest) GetReset_() bool {
	if m != nil {
		return m.Reset_
	}
	return false
}

type Stat struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                int64    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stat) Reset()         { *m = Stat{} }
func (m *Stat) String() string { return proto.CompactTextString(m) }
func (*Stat) ProtoMessage()    {}
func (*Stat) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{1}
}

func (m *Stat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stat.Unmarshal(m, b)
}
func (m *Stat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stat.Marshal(b, m, deterministic)
}
func (m *Stat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stat.Merge(m, src)
}
func (m *Stat) XXX_Size() int {
	return xxx_messageInfo_Stat.Size(m)
}
func (m *Stat) XXX_DiscardUnknown() {
	xxx_messageInfo_Stat.DiscardUnknown(m)
}

var xxx_messageInfo_Stat proto.InternalMessageInfo

func (m *Stat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Stat) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type GetStatsResponse struct {
	Stat                 *Stat    `protobuf:"bytes,1,opt,name=stat,proto3" json:"stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatsResponse) Reset()         { *m = GetStatsResponse{} }
func (m *GetStatsResponse) String() string { return proto.CompactTextString(m) }
func (*GetStatsResponse) ProtoMessage()    {}
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{2}
}

func (m *GetStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsResponse.Unmarshal(m, b)
}
func (m *GetStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsResponse.Marshal(b, m, deterministic)
}
func (m *GetStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsResponse.Merge(m, src)
}
func (m *GetStatsResponse) XXX_Size() int {
	return xxx_messageInfo_GetStatsResponse.Size(m)
}
func (m *GetStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsResponse proto.InternalMessageInfo

func (m *GetStatsResponse) GetStat() *Stat {
	if m != nil {
		return m.Stat
	}
	return nil
}

type QueryStatsRequest struct {
	Pattern              string   `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Reset_               bool     `protobuf:"varint,2,opt,name=reset,proto3" json:"reset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryStatsRequest) Reset()         { *m = QueryStatsRequest{} }
func (m *QueryStatsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryStatsRequest) ProtoMessage()    {}
func (*QueryStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{3}
}

func (m *QueryStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStatsRequest.Unmarshal(m, b)
}
func (m *QueryStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStatsRequest.Marshal(b, m, deterministic)
}
func (m *QueryStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStatsRequest.Merge(m, src)
}
func (m *QueryStatsRequest) XXX_Size() int {
	return xxx_messageInfo_QueryStatsRequest.Size(m)
}
func (m *QueryStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStatsRequest proto.InternalMessageInfo

func (m *QueryStatsRequest) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *QueryStatsRequest) GetReset_() bool {
	if m != nil {
		return m.Reset_
	}
	return false
}

type QueryStatsResponse struct {
	Stat                 []*Stat  `protobuf:"bytes,1,rep,name=stat,proto3" json:"stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryStatsResponse) Reset()         { *m = QueryStatsResponse{} }
func (m *QueryStatsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryStatsResponse) ProtoMessage()    {}
func (*QueryStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{4}
}

func (m *QueryStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStatsResponse.Unmarshal(m, b)
}
func (m *QueryStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStatsResponse.Marshal(b, m, deterministic)
}
func (m *QueryStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStatsResponse.Merge(m, src)
}
func (m *QueryStatsResponse) XXX_Size() int {
	return xxx_messageInfo_QueryStatsResponse.Size(m)
}
func (m *QueryStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStatsResponse proto.InternalMessageInfo

func (m *QueryStatsResponse) GetStat() []*Stat {
	if m != nil {
		return m.Stat
	}
	return nil
}

type GetUserIPStatsResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserIPStatsResponse) Reset()         { *m = GetUserIPStatsResponse{} }
func (m *GetUserIPStatsResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserIPStatsResponse) ProtoMessage()    {}
func (*GetUserIPStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{5}
}

func (m *GetUserIPStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserIPStatsResponse.Unmarshal(m, b)
}
func (m *GetUserIPStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserIPStatsResponse.Marshal(b, m, deterministic)
}
func (m *GetUserIPStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserIPStatsResponse.Merge(m, src)
}
func (m *GetUserIPStatsResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserIPStatsResponse.Size(m)
}
func (m *GetUserIPStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserIPStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserIPStatsResponse proto.InternalMessageInfo

func (m *GetUserIPStatsResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetUserIPStatsResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Config struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_c902411c4948f26b, []int{6}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetStatsRequest)(nil), "v2ray.core.app.stats.command.GetStatsRequest")
	proto.RegisterType((*Stat)(nil), "v2ray.core.app.stats.command.Stat")
	proto.RegisterType((*GetStatsResponse)(nil), "v2ray.core.app.stats.command.GetStatsResponse")
	proto.RegisterType((*QueryStatsRequest)(nil), "v2ray.core.app.stats.command.QueryStatsRequest")
	proto.RegisterType((*QueryStatsResponse)(nil), "v2ray.core.app.stats.command.QueryStatsResponse")
	proto.RegisterType((*GetUserIPStatsResponse)(nil), "v2ray.core.app.stats.command.GetUserIPStatsResponse")
	proto.RegisterType((*Config)(nil), "v2ray.core.app.stats.command.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/app/stats/command/command.proto", fileDescriptor_c902411c4948f26b)
}

var fileDescriptor_c902411c4948f26b = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0xc6, 0x29, 0x70, 0xf9, 0x73, 0xee, 0xcd, 0xbd, 0xd7, 0x89, 0x31, 0x84, 0xb0, 0x20, 0xb3,
	0x62, 0xe3, 0x94, 0x54, 0xe3, 0xc6, 0x95, 0x74, 0x41, 0x34, 0x2c, 0xb0, 0x44, 0x17, 0xee, 0xc6,
	0x7a, 0x34, 0x44, 0xdb, 0x19, 0x66, 0x06, 0x0c, 0xaf, 0xe4, 0x8b, 0xf9, 0x1a, 0xa6, 0x53, 0x2a,
	0x22, 0x52, 0x60, 0xd5, 0x39, 0xed, 0xf7, 0x9d, 0xf9, 0xcd, 0x77, 0x3a, 0xc0, 0x66, 0x9e, 0xe2,
	0x73, 0x16, 0x8a, 0xc8, 0x0d, 0x85, 0x42, 0x97, 0x4b, 0xe9, 0x6a, 0xc3, 0x8d, 0x76, 0x43, 0x11,
	0x45, 0x3c, 0x7e, 0xc8, 0x9e, 0x4c, 0x2a, 0x61, 0x04, 0x69, 0x65, 0x7a, 0x85, 0x8c, 0x4b, 0xc9,
	0xac, 0x96, 0x2d, 0x34, 0xf4, 0x1c, 0xfe, 0xf5, 0xd1, 0x8c, 0x92, 0x77, 0x01, 0x4e, 0xa6, 0xa8,
	0x0d, 0x21, 0x50, 0x8e, 0x79, 0x84, 0x0d, 0xa7, 0xed, 0x74, 0xea, 0x81, 0x5d, 0x93, 0x43, 0xf8,
	0xa5, 0x50, 0xa3, 0x69, 0x14, 0xdb, 0x4e, 0xa7, 0x16, 0xa4, 0x05, 0xed, 0x42, 0x39, 0x71, 0x6e,
	0x72, 0xcc, 0xf8, 0xcb, 0x14, 0xad, 0xa3, 0x14, 0xa4, 0x05, 0xbd, 0x82, 0xff, 0xcb, 0xed, 0xb4,
	0x14, 0xb1, 0x46, 0x72, 0x06, 0xe5, 0x84, 0xc9, 0xba, 0x7f, 0x7b, 0x94, 0xe5, 0xf1, 0xb2, 0xc4,
	0x1a, 0x58, 0x3d, 0xf5, 0xe1, 0xe0, 0x7a, 0x8a, 0x6a, 0xbe, 0x02, 0xdf, 0x80, 0xaa, 0xe4, 0xc6,
	0xa0, 0x8a, 0x17, 0x34, 0x59, 0xb9, 0xe1, 0x08, 0x03, 0x20, 0x5f, 0x9b, 0xac, 0x21, 0x95, 0xf6,
	0x42, 0xea, 0xc1, 0x51, 0x1f, 0xcd, 0x8d, 0x46, 0x75, 0x39, 0x5c, 0xed, 0xb8, 0x35, 0xa2, 0x7a,
	0x16, 0x51, 0x0d, 0x2a, 0xbe, 0x88, 0x1f, 0xc7, 0x4f, 0xde, 0x7b, 0x11, 0xfe, 0xd8, 0x2e, 0x23,
	0x54, 0xb3, 0x71, 0x88, 0xe4, 0x19, 0x6a, 0x59, 0x7a, 0xe4, 0x38, 0x1f, 0xea, 0xdb, 0x50, 0x9b,
	0x6c, 0x57, 0x79, 0xca, 0x4b, 0x0b, 0x64, 0x02, 0xb0, 0x4c, 0x86, 0xb8, 0xf9, 0xfe, 0xb5, 0x41,
	0x34, 0xbb, 0xbb, 0x1b, 0x3e, 0xb7, 0x7c, 0x85, 0xbf, 0xab, 0xf1, 0xed, 0x7b, 0xca, 0xd3, 0xad,
	0xf2, 0x1f, 0x66, 0x43, 0x0b, 0xbd, 0x01, 0xb4, 0x43, 0x11, 0xe5, 0x9a, 0x87, 0xce, 0x5d, 0x75,
	0xb1, 0x7c, 0x2b, 0xb6, 0x6e, 0xbd, 0x80, 0xcf, 0x99, 0x9f, 0x28, 0x2f, 0xa4, 0xb4, 0xbf, 0x80,
	0x66, 0x7e, 0xfa, 0xf9, 0xbe, 0x62, 0x2f, 0xde, 0xc9, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd1,
	0xc4, 0x60, 0x2a, 0xaa, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StatsServiceClient is the client API for StatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatsServiceClient interface {
	GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error)
	QueryStats(ctx context.Context, in *QueryStatsRequest, opts ...grpc.CallOption) (*QueryStatsResponse, error)
	GetUserIPStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetUserIPStatsResponse, error)
}

type statsServiceClient struct {
	cc *grpc.ClientConn
}

func NewStatsServiceClient(cc *grpc.ClientConn) StatsServiceClient {
	return &statsServiceClient{cc}
}

func (c *statsServiceClient) GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, "/v2ray.core.app.stats.command.StatsService/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) QueryStats(ctx context.Context, in *QueryStatsRequest, opts ...grpc.CallOption) (*QueryStatsResponse, error) {
	out := new(QueryStatsResponse)
	err := c.cc.Invoke(ctx, "/v2ray.core.app.stats.command.StatsService/QueryStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) GetUserIPStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetUserIPStatsResponse, error) {
	out := new(GetUserIPStatsResponse)
	err := c.cc.Invoke(ctx, "/v2ray.core.app.stats.command.StatsService/GetUserIPStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsServiceServer is the server API for StatsService service.
type StatsServiceServer interface {
	GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
	QueryStats(context.Context, *QueryStatsRequest) (*QueryStatsResponse, error)
	GetUserIPStats(context.Context, *GetStatsRequest) (*GetUserIPStatsResponse, error)
}

func RegisterStatsServiceServer(s *grpc.Server, srv StatsServiceServer) {
	s.RegisterService(&_StatsService_serviceDesc, srv)
}

func _StatsService_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.core.app.stats.command.StatsService/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetStats(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_QueryStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).QueryStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.core.app.stats.command.StatsService/QueryStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).QueryStats(ctx, req.(*QueryStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_GetUserIPStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetUserIPStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v2ray.core.app.stats.command.StatsService/GetUserIPStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetUserIPStats(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v2ray.core.app.stats.command.StatsService",
	HandlerType: (*StatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStats",
			Handler:    _StatsService_GetStats_Handler,
		},
		{
			MethodName: "QueryStats",
			Handler:    _StatsService_QueryStats_Handler,
		},
		{
			MethodName: "GetUserIPStats",
			Handler:    _StatsService_GetUserIPStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v2ray.com/core/app/stats/command/command.proto",
}
