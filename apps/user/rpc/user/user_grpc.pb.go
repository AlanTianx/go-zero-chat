// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
	LoginCode(ctx context.Context, in *LoginCodeReq, opts ...grpc.CallOption) (*LoginCodeRes, error)
	Detail(ctx context.Context, in *DetailReq, opts ...grpc.CallOption) (*LoginRes, error)
	Save(ctx context.Context, in *SaveReq, opts ...grpc.CallOption) (*SaveRes, error)
	GetUserToken(ctx context.Context, in *GetUserTokenReq, opts ...grpc.CallOption) (*GetUserTokenRes, error)
	CheckUserToken(ctx context.Context, in *CheckTokenReq, opts ...grpc.CallOption) (*CheckTokenRes, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	out := new(LoginRes)
	err := c.cc.Invoke(ctx, "/user.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) LoginCode(ctx context.Context, in *LoginCodeReq, opts ...grpc.CallOption) (*LoginCodeRes, error) {
	out := new(LoginCodeRes)
	err := c.cc.Invoke(ctx, "/user.User/LoginCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Detail(ctx context.Context, in *DetailReq, opts ...grpc.CallOption) (*LoginRes, error) {
	out := new(LoginRes)
	err := c.cc.Invoke(ctx, "/user.User/Detail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Save(ctx context.Context, in *SaveReq, opts ...grpc.CallOption) (*SaveRes, error) {
	out := new(SaveRes)
	err := c.cc.Invoke(ctx, "/user.User/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserToken(ctx context.Context, in *GetUserTokenReq, opts ...grpc.CallOption) (*GetUserTokenRes, error) {
	out := new(GetUserTokenRes)
	err := c.cc.Invoke(ctx, "/user.User/GetUserToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CheckUserToken(ctx context.Context, in *CheckTokenReq, opts ...grpc.CallOption) (*CheckTokenRes, error) {
	out := new(CheckTokenRes)
	err := c.cc.Invoke(ctx, "/user.User/CheckUserToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Login(context.Context, *LoginReq) (*LoginRes, error)
	LoginCode(context.Context, *LoginCodeReq) (*LoginCodeRes, error)
	Detail(context.Context, *DetailReq) (*LoginRes, error)
	Save(context.Context, *SaveReq) (*SaveRes, error)
	GetUserToken(context.Context, *GetUserTokenReq) (*GetUserTokenRes, error)
	CheckUserToken(context.Context, *CheckTokenReq) (*CheckTokenRes, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Login(context.Context, *LoginReq) (*LoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) LoginCode(context.Context, *LoginCodeReq) (*LoginCodeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginCode not implemented")
}
func (UnimplementedUserServer) Detail(context.Context, *DetailReq) (*LoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}
func (UnimplementedUserServer) Save(context.Context, *SaveReq) (*SaveRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedUserServer) GetUserToken(context.Context, *GetUserTokenReq) (*GetUserTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserToken not implemented")
}
func (UnimplementedUserServer) CheckUserToken(context.Context, *CheckTokenReq) (*CheckTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserToken not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_LoginCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).LoginCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/LoginCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).LoginCode(ctx, req.(*LoginCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Detail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Detail(ctx, req.(*DetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Save(ctx, req.(*SaveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetUserToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserToken(ctx, req.(*GetUserTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CheckUserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CheckUserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/CheckUserToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CheckUserToken(ctx, req.(*CheckTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "LoginCode",
			Handler:    _User_LoginCode_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _User_Detail_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _User_Save_Handler,
		},
		{
			MethodName: "GetUserToken",
			Handler:    _User_GetUserToken_Handler,
		},
		{
			MethodName: "CheckUserToken",
			Handler:    _User_CheckUserToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}