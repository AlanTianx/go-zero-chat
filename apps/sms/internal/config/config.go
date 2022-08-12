package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Sms struct {
		KeyId     string
		KeySecret string
	}
}
