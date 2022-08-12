package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-chat/apps/sms/sms"
	"go-zero-chat/apps/user/rpc/internal/config"
	"go-zero-chat/apps/user/rpc/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config          config.Config
	BizRedis        *redis.Redis
	UserModel       model.UsersModel
	UserTokensModel model.UserTokensModel
	SmsRpc          sms.Sms
}

func NewServiceContext(c config.Config) *ServiceContext {
	//sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	db, _ := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	return &ServiceContext{
		Config:          c,
		BizRedis:        redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
		SmsRpc:          sms.NewSms(zrpc.MustNewClient(c.SmsRpc)),
		UserModel:       model.NewUsersModel(db, c.CacheRedis),
		UserTokensModel: model.NewUserTokensModel(db, c.CacheRedis),
	}
}
