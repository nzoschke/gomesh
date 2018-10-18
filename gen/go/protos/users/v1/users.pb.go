// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/users/v1/users.proto

package v1pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type User struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Parent               string               `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string               `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_cdddafa5154dbb1e, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

type GetRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_cdddafa5154dbb1e, []int{1}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateRequest struct {
	Parent               string   `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_users_cdddafa5154dbb1e, []int{2}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "omgrpc.users.v1.User")
	proto.RegisterType((*GetRequest)(nil), "omgrpc.users.v1.GetRequest")
	proto.RegisterType((*CreateRequest)(nil), "omgrpc.users.v1.CreateRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*User, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/omgrpc.users.v1.Users/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/omgrpc.users.v1.Users/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	Get(context.Context, *GetRequest) (*User, error)
	Create(context.Context, *CreateRequest) (*User, error)
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omgrpc.users.v1.Users/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omgrpc.users.v1.Users/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "omgrpc.users.v1.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Users_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Users_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/users/v1/users.proto",
}

func init() { proto.RegisterFile("protos/users/v1/users.proto", fileDescriptor_users_cdddafa5154dbb1e) }

var fileDescriptor_users_cdddafa5154dbb1e = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x6b, 0xc2, 0x40,
	0x10, 0xc5, 0x59, 0x8d, 0x42, 0x27, 0xfd, 0x03, 0x5b, 0x5a, 0x82, 0x42, 0x6b, 0x3d, 0xe9, 0x65,
	0x83, 0xf6, 0x54, 0x7a, 0x28, 0xd8, 0x83, 0xb7, 0x22, 0xa1, 0xbd, 0xf4, 0x22, 0x51, 0xa7, 0x12,
	0x30, 0xee, 0x76, 0x77, 0x13, 0xe8, 0xb9, 0xdf, 0xa4, 0x9f, 0xb4, 0xec, 0xac, 0x21, 0xa8, 0xed,
	0x6d, 0xe6, 0xbd, 0xb7, 0xc3, 0x6f, 0x1f, 0x74, 0x95, 0x96, 0x56, 0x9a, 0xb8, 0x30, 0xa8, 0x4d,
	0x5c, 0x8e, 0xfc, 0x20, 0x48, 0xe5, 0x17, 0x32, 0x5f, 0x6b, 0xb5, 0x14, 0x5e, 0x2b, 0x47, 0x9d,
	0xdb, 0xb5, 0x94, 0xeb, 0x0d, 0xc6, 0x64, 0x2f, 0x8a, 0x8f, 0xd8, 0x66, 0x39, 0x1a, 0x9b, 0xe6,
	0xca, 0xbf, 0xe8, 0xff, 0x30, 0x08, 0xde, 0x0c, 0x6a, 0x7e, 0x0e, 0x8d, 0x6c, 0x15, 0xb1, 0x1e,
	0x1b, 0x9c, 0x24, 0x8d, 0x6c, 0xc5, 0xaf, 0xa1, 0xad, 0x52, 0x8d, 0x5b, 0x1b, 0x35, 0x48, 0xdb,
	0x6d, 0x9c, 0x43, 0xb0, 0x4d, 0x73, 0x8c, 0x9a, 0xa4, 0xd2, 0xcc, 0xef, 0xe0, 0x74, 0x95, 0x19,
	0xb5, 0x49, 0xbf, 0xe6, 0xe4, 0x05, 0xe4, 0x85, 0x3b, 0xed, 0xc5, 0x45, 0x1e, 0x21, 0x5c, 0x6a,
	0x4c, 0x2d, 0xce, 0x1d, 0x41, 0xd4, 0xea, 0xb1, 0x41, 0x38, 0xee, 0x08, 0x8f, 0x27, 0x2a, 0x3c,
	0xf1, 0x5a, 0xe1, 0x25, 0xe0, 0xe3, 0x4e, 0xe8, 0xf7, 0x00, 0xa6, 0x68, 0x13, 0xfc, 0x2c, 0xd0,
	0xd4, 0x04, 0xac, 0x26, 0xe8, 0x27, 0x70, 0xf6, 0x4c, 0xf9, 0x2a, 0x54, 0xe3, 0xb3, 0x3d, 0xfc,
	0x21, 0x04, 0xae, 0x1c, 0xfa, 0x54, 0x38, 0xbe, 0x12, 0x07, 0x85, 0x09, 0xd7, 0x45, 0x42, 0x91,
	0xf1, 0x37, 0x83, 0x96, 0x5b, 0x0d, 0x7f, 0x80, 0xe6, 0x14, 0x2d, 0xef, 0x1e, 0xa5, 0x6b, 0xaa,
	0xce, 0xdf, 0xa7, 0xf8, 0x13, 0xb4, 0x3d, 0x18, 0xbf, 0x39, 0x0a, 0xec, 0x11, 0xff, 0x73, 0x60,
	0x32, 0x84, 0xcb, 0xa5, 0xcc, 0x0f, 0xbd, 0x09, 0x10, 0xd9, 0xcc, 0xf5, 0x36, 0x63, 0xef, 0x41,
	0x39, 0x52, 0x8b, 0x45, 0x9b, 0x6a, 0xbc, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xda, 0xf3,
	0x1f, 0x23, 0x02, 0x00, 0x00,
}
