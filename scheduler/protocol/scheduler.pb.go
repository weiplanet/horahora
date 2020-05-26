// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scheduler.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Site int32

const (
	Site_niconico Site = 0
	Site_bilibili Site = 1
	Site_youtube  Site = 2
)

var Site_name = map[int32]string{
	0: "niconico",
	1: "bilibili",
	2: "youtube",
}

var Site_value = map[string]int32{
	"niconico": 0,
	"bilibili": 1,
	"youtube":  2,
}

func (x Site) String() string {
	return proto.EnumName(Site_name, int32(x))
}

func (Site) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2b3fc28395a6d9c5, []int{0}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b3fc28395a6d9c5, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ChannelRequest struct {
	Website              Site     `protobuf:"varint,1,opt,name=website,proto3,enum=proto.Site" json:"website,omitempty"`
	UserID               string   `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	NumToDownload        uint64   `protobuf:"varint,3,opt,name=NumToDownload,proto3" json:"NumToDownload,omitempty"`
	ChannelID            string   `protobuf:"bytes,4,opt,name=channelID,proto3" json:"channelID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelRequest) Reset()         { *m = ChannelRequest{} }
func (m *ChannelRequest) String() string { return proto.CompactTextString(m) }
func (*ChannelRequest) ProtoMessage()    {}
func (*ChannelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b3fc28395a6d9c5, []int{1}
}

func (m *ChannelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelRequest.Unmarshal(m, b)
}
func (m *ChannelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelRequest.Marshal(b, m, deterministic)
}
func (m *ChannelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelRequest.Merge(m, src)
}
func (m *ChannelRequest) XXX_Size() int {
	return xxx_messageInfo_ChannelRequest.Size(m)
}
func (m *ChannelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelRequest proto.InternalMessageInfo

func (m *ChannelRequest) GetWebsite() Site {
	if m != nil {
		return m.Website
	}
	return Site_niconico
}

func (m *ChannelRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ChannelRequest) GetNumToDownload() uint64 {
	if m != nil {
		return m.NumToDownload
	}
	return 0
}

func (m *ChannelRequest) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

type PlaylistRequest struct {
	Website              Site     `protobuf:"varint,1,opt,name=website,proto3,enum=proto.Site" json:"website,omitempty"`
	UserID               string   `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	NumToDownload        uint64   `protobuf:"varint,3,opt,name=NumToDownload,proto3" json:"NumToDownload,omitempty"`
	PlaylistID           string   `protobuf:"bytes,4,opt,name=playlistID,proto3" json:"playlistID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlaylistRequest) Reset()         { *m = PlaylistRequest{} }
func (m *PlaylistRequest) String() string { return proto.CompactTextString(m) }
func (*PlaylistRequest) ProtoMessage()    {}
func (*PlaylistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b3fc28395a6d9c5, []int{2}
}

func (m *PlaylistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlaylistRequest.Unmarshal(m, b)
}
func (m *PlaylistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlaylistRequest.Marshal(b, m, deterministic)
}
func (m *PlaylistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlaylistRequest.Merge(m, src)
}
func (m *PlaylistRequest) XXX_Size() int {
	return xxx_messageInfo_PlaylistRequest.Size(m)
}
func (m *PlaylistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlaylistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlaylistRequest proto.InternalMessageInfo

func (m *PlaylistRequest) GetWebsite() Site {
	if m != nil {
		return m.Website
	}
	return Site_niconico
}

func (m *PlaylistRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *PlaylistRequest) GetNumToDownload() uint64 {
	if m != nil {
		return m.NumToDownload
	}
	return 0
}

func (m *PlaylistRequest) GetPlaylistID() string {
	if m != nil {
		return m.PlaylistID
	}
	return ""
}

type TagRequest struct {
	Website              Site     `protobuf:"varint,1,opt,name=website,proto3,enum=proto.Site" json:"website,omitempty"`
	UserID               string   `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	NumToDownload        uint64   `protobuf:"varint,3,opt,name=NumToDownload,proto3" json:"NumToDownload,omitempty"`
	TagValue             string   `protobuf:"bytes,4,opt,[name=tagValue,proto3" json:"tagValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagRequest) Reset()         { *m = TagRequest{} }
func (m *TagRequest) String() string { return proto.CompactTextString(m) }
func (*TagRequest) ProtoMessage()    {}
func (*TagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b3fc28395a6d9c5, []int{3}
}

func (m *TagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagRequest.Unmarshal(m, b)
}
func (m *TagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagRequest.Marshal(b, m, deterministic)
}
func (m *TagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagRequest.Merge(m, src)
}
func (m *TagRequest) XXX_Size() int {
	return xxx_messageInfo_TagRequest.Size(m)
}
func (m *TagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TagRequest proto.InternalMessageInfo

func (m *TagRequest) GetWebsite() Site {
	if m != nil {
		return m.Website
	}
	return Site_niconico
}

func (m *TagRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *TagRequest) GetNumToDownload() uint64 {
	if m != nil {
		return m.NumToDownload
	}
	return 0
}

func (m *TagRequest) GetTagValue() string {
	if m != nil {
		return m.TagValue
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.Site", Site_name, Site_value)
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*ChannelRequest)(nil), "proto.ChannelRequest")
	proto.RegisterType((*PlaylistRequest)(nil), "proto.PlaylistRequest")
	proto.RegisterType((*TagRequest)(nil), "proto.TagRequest")
}

func init() { proto.RegisterFile("scheduler.proto", fileDescriptor_2b3fc28395a6d9c5) }

var fileDescriptor_2b3fc28395a6d9c5 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xcd, 0x6a, 0xb3, 0x40,
	0x14, 0xcd, 0xe4, 0xcb, 0x9f, 0x37, 0xf9, 0x12, 0x3b, 0xd0, 0x20, 0xa1, 0x14, 0x91, 0x16, 0x24,
	0x8b, 0x14, 0x6c, 0xdf, 0xa0, 0x76, 0x91, 0x4d, 0x29, 0x56, 0xba, 0x1f, 0x75, 0x30, 0xc2, 0xc4,
	0xb1, 0x3a, 0x43, 0xf0, 0x21, 0xba, 0x2a, 0x74, 0xdf, 0x37, 0x2d, 0xea, 0x98, 0x34, 0xe9, 0x3e,
	0x8b, 0xcb, 0x70, 0xce, 0x9d, 0x3b, 0xf7, 0x30, 0xe7, 0xc0, 0xac, 0x08, 0x37, 0x34, 0x92, 0x8c,
	0xe6, 0xab, 0x2c, 0xe7, 0x82, 0xe3, 0x7e, 0x7d, 0x58, 0x43, 0xe8, 0x3f, 0x6d, 0x33, 0x51, 0x5a,
	0x9f, 0x08, 0xa6, 0x8f, 0x1b, 0x92, 0xa6, 0x94, 0x79, 0xf4, 0x5d, 0xd2, 0x42, 0xe0, 0x5b, 0x18,
	0xee, 0x68, 0x50, 0x24, 0x82, 0x1a, 0xc8, 0x44, 0xf6, 0xd4, 0x19, 0x37, 0xb3, 0xab, 0x8a, 0xf2,
	0xda, 0x1e, 0x9e, 0xc3, 0x40, 0x16, 0x34, 0x5f, 0xbb, 0x46, 0xd7, 0x44, 0xb6, 0xe6, 0x29, 0x84,
	0x6f, 0xe0, 0xff, 0xb3, 0xdc, 0xfa, 0xdc, 0xe5, 0xbb, 0x94, 0x71, 0x12, 0x19, 0xff, 0x4c, 0x64,
	0xf7, 0xbc, 0x63, 0x12, 0x5f, 0x81, 0x16, 0x36, 0x6b, 0xd7, 0xae, 0xd1, 0xab, 0x1f, 0x38, 0x10,
	0xd6, 0x17, 0x82, 0xd9, 0x0b, 0x23, 0x25, 0x4b, 0x0a, 0x71, 0x56, 0x59, 0xd7, 0x00, 0x99, 0xda,
	0xbb, 0xd7, 0xf5, 0x8b, 0xb1, 0x3e, 0x10, 0x80, 0x4f, 0xe2, 0xb3, 0x6a, 0x5a, 0xc0, 0x48, 0x90,
	0xf8, 0x8d, 0x30, 0x49, 0x95, 0xa2, 0x3d, 0x5e, 0xde, 0x41, 0xaf, 0xde, 0x30, 0x81, 0x51, 0x9a,
	0x84, 0xbc, 0x2a, 0xbd, 0x53, 0xa1, 0x20, 0x61, 0x49, 0x55, 0x3a, 0xc2, 0x63, 0x18, 0x96, 0x5c,
	0x0a, 0x19, 0x50, 0xbd, 0xeb, 0x7c, 0x23, 0xd0, 0x5e, 0xdb, 0x4c, 0x60, 0x07, 0xb4, 0x88, 0x29,
	0xfb, 0xf1, 0xa5, 0xd2, 0x7e, 0x1c, 0x87, 0xc5, 0x44, 0xd1, 0x4d, 0x5e, 0x3a, 0xf8, 0x01, 0x20,
	0x62, 0xad, 0x39, 0x78, 0xae, 0xba, 0x27, 0x6e, 0xfd, 0x99, 0x5a, 0x42, 0x3f, 0x62, 0x3e, 0x89,
	0xf1, 0x85, 0x6a, 0x1c, 0x7e, 0xf1, 0xf4, 0x6e, 0x30, 0xa8, 0xe1, 0xfd, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xdb, 0xf8, 0x9a, 0xcb, 0xbd, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SchedulerClient is the client API for Scheduler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchedulerClient interface {
	DlChannel(ctx context.Context, in *ChannelRequest, opts ...grpc.CallOption) (*Empty, error)
	DlPlaylist(ctx context.Context, in *PlaylistRequest, opts ...grpc.CallOption) (*Empty, error)
	DlTag(ctx context.Context, in *TagRequest, opts ...grpc.CallOption) (*Empty, error)
}

type schedulerClient struct {
	cc grpc.ClientConnInterface
}

func NewSchedulerClient(cc grpc.ClientConnInterface) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) DlChannel(ctx context.Context, in *ChannelRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Scheduler/dlChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) DlPlaylist(ctx context.Context, in *PlaylistRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Scheduler/dlPlaylist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) DlTag(ctx context.Context, in *TagRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Scheduler/dlTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerServer is the server API for Scheduler service.
type SchedulerServer interface {
	DlChannel(context.Context, *ChannelRequest) (*Empty, error)
	DlPlaylist(context.Context, *PlaylistRequest) (*Empty, error)
	DlTag(context.Context, *TagRequest) (*Empty, error)
}

// UnimplementedSchedulerServer can be embedded to have forward compatible implementations.
type UnimplementedSchedulerServer struct {
}

func (*UnimplementedSchedulerServer) DlChannel(ctx context.Context, req *ChannelRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DlChannel not implemented")
}
func (*UnimplementedSchedulerServer) DlPlaylist(ctx context.Context, req *PlaylistRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DlPlaylist not implemented")
}
func (*UnimplementedSchedulerServer) DlTag(ctx context.Context, req *TagRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DlTag not implemented")
}

func RegisterSchedulerServer(s *grpc.Server, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_DlChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).DlChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Scheduler/DlChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).DlChannel(ctx, req.(*ChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_DlPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).DlPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Scheduler/DlPlaylist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).DlPlaylist(ctx, req.(*PlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_DlTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).DlTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Scheduler/DlTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).DlTag(ctx, req.(*TagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "dlChannel",
			Handler:    _Scheduler_DlChannel_Handler,
		},
		{
			MethodName: "dlPlaylist",
			Handler:    _Scheduler_DlPlaylist_Handler,
		},
		{
			MethodName: "dlTag",
			Handler:    _Scheduler_DlTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheduler.proto",
}
