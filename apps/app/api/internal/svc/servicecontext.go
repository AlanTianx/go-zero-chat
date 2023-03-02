package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-chat/apps/app/api/internal/config"
	"go-zero-chat/apps/app/api/internal/middleware"
	"go-zero-chat/apps/sms/sms"
	"go-zero-chat/apps/user/rpc/user"
	"go-zero-chat/pkg/interceptor/rpcserver"
)

type ServiceContext struct {
	Config   config.Config
	UserRpc  user.User
	SmsRpc   sms.Sms
	BizRedis *redis.Redis
	Limit    rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	r := redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass))
	return &ServiceContext{
		Config:   c,
		Limit:    middleware.NewLimitMiddleware(r).Handle,
		SmsRpc:   sms.NewSms(zrpc.MustNewClient(c.SmsRpc, zrpc.WithUnaryClientInterceptor(rpcserver.BreakerInterceptor))),
		UserRpc:  user.NewUser(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(rpcserver.BreakerInterceptor))),
		BizRedis: r,
	}
}
