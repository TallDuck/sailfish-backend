// Code generated by protoc-gen-go.
// source: auth.proto
// DO NOT EDIT!

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	auth.proto

It has these top-level messages:
	Request
	Response
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Request struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Response struct {
	Status bool `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*Request)(nil), "auth.Request")
	proto.RegisterType((*Response)(nil), "auth.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	Authenticate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Invalidate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Authenticate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/auth.Auth/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Invalidate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/auth.Auth/Invalidate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	Authenticate(context.Context, *Request) (*Response, error)
	Invalidate(context.Context, *Request) (*Response, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Authenticate(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Invalidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Invalidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/Invalidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Invalidate(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _Auth_Authenticate_Handler,
		},
		{
			MethodName: "Invalidate",
			Handler:    _Auth_Invalidate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xe4, 0xb9, 0xd8, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0x4b, 0xf2, 0xb3, 0x53, 0xf3, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x25, 0x25, 0x2e, 0x8e, 0xa0, 0xd4, 0xe2, 0x82, 0xfc,
	0xbc, 0xe2, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0xe2, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0xb0, 0x12, 0x8e,
	0x20, 0x28, 0xcf, 0x28, 0x89, 0x8b, 0xc5, 0xb1, 0xb4, 0x24, 0x43, 0x48, 0x97, 0x8b, 0x07, 0x44,
	0xa7, 0xe6, 0x95, 0x64, 0x26, 0x27, 0x96, 0xa4, 0x0a, 0xf1, 0xea, 0x81, 0xed, 0x83, 0x5a, 0x20,
	0xc5, 0x07, 0xe3, 0x42, 0x8d, 0xd3, 0xe6, 0xe2, 0xf2, 0xcc, 0x2b, 0x4b, 0xcc, 0xc9, 0x4c, 0x21,
	0xac, 0x38, 0x89, 0x0d, 0xec, 0x6a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x5b, 0x01,
	0x7f, 0xc3, 0x00, 0x00, 0x00,
}