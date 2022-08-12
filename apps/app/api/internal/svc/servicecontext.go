package svc

import (
	"go-zero-chat/apps/app/api/internal/config"
	"go-zero-chat/apps/sms/sms"
	"go-zero-chat/apps/user/rpc/user"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	UserRpc  user.User
	SmsRpc   sms.Sms
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		SmsRpc:   sms.NewSms(zrpc.MustNewClient(c.SmsRpc)),
		UserRpc:  user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		BizRedis: redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
	}
}
