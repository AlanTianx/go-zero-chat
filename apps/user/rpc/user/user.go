// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package user

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	User interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
		LoginCode(ctx context.Context, in *LoginCodeReq, opts ...grpc.CallOption) (*LoginCodeRes, error)
		Detail(ctx context.Context, in *DetailReq, opts ...grpc.CallOption) (*LoginRes, error)
		Save(ctx context.Context, in *SaveReq, opts ...grpc.CallOption) (*SaveRes, error)
		GetUserToken(ctx context.Context, in *GetUserTokenReq, opts ...grpc.CallOption) (*GetUserTokenRes, error)
		CheckUserToken(ctx context.Context, in *CheckTokenReq, opts ...grpc.CallOption) (*CheckTokenRes, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	client := NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) LoginCode(ctx context.Context, in *LoginCodeReq, opts ...grpc.CallOption) (*LoginCodeRes, error) {
	client := NewUserClient(m.cli.Conn())
	return client.LoginCode(ctx, in, opts...)
}

func (m *defaultUser) Detail(ctx context.Context, in *DetailReq, opts ...grpc.CallOption) (*LoginRes, error) {
	client := NewUserClient(m.cli.Conn())
	return client.Detail(ctx, in, opts...)
}

func (m *defaultUser) Save(ctx context.Context, in *SaveReq, opts ...grpc.CallOption) (*SaveRes, error) {
	client := NewUserClient(m.cli.Conn())
	return client.Save(ctx, in, opts...)
}

func (m *defaultUser) GetUserToken(ctx context.Context, in *GetUserTokenReq, opts ...grpc.CallOption) (*GetUserTokenRes, error) {
	client := NewUserClient(m.cli.Conn())
	return client.GetUserToken(ctx, in, opts...)
}

func (m *defaultUser) CheckUserToken(ctx context.Context, in *CheckTokenReq, opts ...grpc.CallOption) (*CheckTokenRes, error) {
	client := NewUserClient(m.cli.Conn())
	return client.CheckUserToken(ctx, in, opts...)
}
