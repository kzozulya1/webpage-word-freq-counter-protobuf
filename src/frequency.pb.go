// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frequency.proto

package frequency

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

type PageWordFrequency struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PageUrl              string   `protobuf:"bytes,2,opt,name=page_url,json=pageUrl,proto3" json:"page_url,omitempty"`
	PageTitle            string   `protobuf:"bytes,3,opt,name=page_title,json=pageTitle,proto3" json:"page_title,omitempty"`
	Words                []*Word  `protobuf:"bytes,4,rep,name=words,proto3" json:"words,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageWordFrequency) Reset()         { *m = PageWordFrequency{} }
func (m *PageWordFrequency) String() string { return proto.CompactTextString(m) }
func (*PageWordFrequency) ProtoMessage()    {}
func (*PageWordFrequency) Descriptor() ([]byte, []int) {
	return fileDescriptor_717812b2e414acee, []int{0}
}

func (m *PageWordFrequency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageWordFrequency.Unmarshal(m, b)
}
func (m *PageWordFrequency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageWordFrequency.Marshal(b, m, deterministic)
}
func (m *PageWordFrequency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageWordFrequency.Merge(m, src)
}
func (m *PageWordFrequency) XXX_Size() int {
	return xxx_messageInfo_PageWordFrequency.Size(m)
}
func (m *PageWordFrequency) XXX_DiscardUnknown() {
	xxx_messageInfo_PageWordFrequency.DiscardUnknown(m)
}

var xxx_messageInfo_PageWordFrequency proto.InternalMessageInfo

func (m *PageWordFrequency) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PageWordFrequency) GetPageUrl() string {
	if m != nil {
		return m.PageUrl
	}
	return ""
}

func (m *PageWordFrequency) GetPageTitle() string {
	if m != nil {
		return m.PageTitle
	}
	return ""
}

func (m *PageWordFrequency) GetWords() []*Word {
	if m != nil {
		return m.Words
	}
	return nil
}

type Word struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Count                int32    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Word) Reset()         { *m = Word{} }
func (m *Word) String() string { return proto.CompactTextString(m) }
func (*Word) ProtoMessage()    {}
func (*Word) Descriptor() ([]byte, []int) {
	return fileDescriptor_717812b2e414acee, []int{1}
}

func (m *Word) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Word.Unmarshal(m, b)
}
func (m *Word) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Word.Marshal(b, m, deterministic)
}
func (m *Word) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Word.Merge(m, src)
}
func (m *Word) XXX_Size() int {
	return xxx_messageInfo_Word.Size(m)
}
func (m *Word) XXX_DiscardUnknown() {
	xxx_messageInfo_Word.DiscardUnknown(m)
}

var xxx_messageInfo_Word proto.InternalMessageInfo

func (m *Word) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Word) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Word) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Response struct {
	Created              bool               `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *PageWordFrequency `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_717812b2e414acee, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *PageWordFrequency {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func init() {
	proto.RegisterType((*PageWordFrequency)(nil), "frequency.PageWordFrequency")
	proto.RegisterType((*Word)(nil), "frequency.Word")
	proto.RegisterType((*Response)(nil), "frequency.Response")
}

func init() { proto.RegisterFile("frequency.proto", fileDescriptor_717812b2e414acee) }

var fileDescriptor_717812b2e414acee = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0x6d, 0xb7, 0xba, 0xf6, 0x2b, 0x38, 0x8c, 0x03, 0xa3, 0x28, 0x8c, 0x82, 0xb0, 0xd3,
	0x0e, 0xf5, 0xee, 0x41, 0xc1, 0xa3, 0x48, 0x54, 0x3c, 0x4a, 0x4d, 0xbe, 0x95, 0x40, 0x4d, 0x6a,
	0x9a, 0x4e, 0xbc, 0xfb, 0x87, 0x4b, 0x52, 0xab, 0x95, 0xc2, 0x8e, 0xef, 0xf7, 0xc8, 0x7b, 0x8f,
	0x2f, 0x30, 0xdf, 0x18, 0x7c, 0x6f, 0x51, 0xf1, 0xcf, 0x75, 0x6d, 0xb4, 0xd5, 0x24, 0xf9, 0x05,
	0xd9, 0x57, 0x00, 0x87, 0xf7, 0x45, 0x89, 0xcf, 0xda, 0x88, 0xdb, 0x9e, 0x92, 0x03, 0x08, 0xa5,
	0xa0, 0xc1, 0x32, 0x58, 0x25, 0x2c, 0x94, 0x82, 0x9c, 0x40, 0x5c, 0x17, 0x25, 0xbe, 0xb4, 0xa6,
	0xa2, 0xa1, 0xa7, 0x33, 0xa7, 0x9f, 0x4c, 0x45, 0xce, 0x01, 0xbc, 0x65, 0xa5, 0xad, 0x90, 0x4e,
	0xbc, 0x99, 0x38, 0xf2, 0xe8, 0x00, 0xb9, 0x80, 0xe8, 0x43, 0x1b, 0xd1, 0xd0, 0xe9, 0x72, 0xb2,
	0x4a, 0xf3, 0xf9, 0xfa, 0x6f, 0x8b, 0xab, 0x64, 0x9d, 0x9b, 0x5d, 0xc3, 0xd4, 0xc9, 0x51, 0xf1,
	0x02, 0xa2, 0x6d, 0x51, 0xb5, 0xf8, 0xd3, 0xda, 0x09, 0x47, 0xb9, 0x6e, 0x95, 0xf5, 0x75, 0x11,
	0xeb, 0x44, 0x26, 0x20, 0x66, 0xd8, 0xd4, 0x5a, 0x35, 0x48, 0x28, 0xcc, 0xb8, 0xc1, 0xc2, 0x62,
	0x17, 0x16, 0xb3, 0x5e, 0x92, 0x2b, 0x48, 0xb9, 0x56, 0x8d, 0x2c, 0xd5, 0x1b, 0x2a, 0xeb, 0x73,
	0xd3, 0xfc, 0x6c, 0x30, 0x6b, 0x74, 0x0d, 0x36, 0x7c, 0x90, 0x6f, 0x60, 0xf1, 0xcf, 0x7d, 0x40,
	0xb3, 0x95, 0x1c, 0xc9, 0x1d, 0x1c, 0xdf, 0xf8, 0x8a, 0xf1, 0x35, 0x77, 0xa6, 0x9f, 0x1e, 0x0d,
	0xdc, 0x7e, 0x7f, 0xb6, 0xf7, 0xba, 0xef, 0xbf, 0xea, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x67,
	0x75, 0x48, 0xc9, 0xbd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WordFrequencyServiceClient is the client API for WordFrequencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WordFrequencyServiceClient interface {
	CreatePageWordFrequency(ctx context.Context, in *PageWordFrequency, opts ...grpc.CallOption) (*Response, error)
}

type wordFrequencyServiceClient struct {
	cc *grpc.ClientConn
}

func NewWordFrequencyServiceClient(cc *grpc.ClientConn) WordFrequencyServiceClient {
	return &wordFrequencyServiceClient{cc}
}

func (c *wordFrequencyServiceClient) CreatePageWordFrequency(ctx context.Context, in *PageWordFrequency, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/frequency.WordFrequencyService/CreatePageWordFrequency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WordFrequencyServiceServer is the server API for WordFrequencyService service.
type WordFrequencyServiceServer interface {
	CreatePageWordFrequency(context.Context, *PageWordFrequency) (*Response, error)
}

// UnimplementedWordFrequencyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWordFrequencyServiceServer struct {
}

func (*UnimplementedWordFrequencyServiceServer) CreatePageWordFrequency(ctx context.Context, req *PageWordFrequency) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePageWordFrequency not implemented")
}

func RegisterWordFrequencyServiceServer(s *grpc.Server, srv WordFrequencyServiceServer) {
	s.RegisterService(&_WordFrequencyService_serviceDesc, srv)
}

func _WordFrequencyService_CreatePageWordFrequency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageWordFrequency)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordFrequencyServiceServer).CreatePageWordFrequency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/frequency.WordFrequencyService/CreatePageWordFrequency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordFrequencyServiceServer).CreatePageWordFrequency(ctx, req.(*PageWordFrequency))
	}
	return interceptor(ctx, in, info, handler)
}

var _WordFrequencyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "frequency.WordFrequencyService",
	HandlerType: (*WordFrequencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePageWordFrequency",
			Handler:    _WordFrequencyService_CreatePageWordFrequency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "frequency.proto",
}
