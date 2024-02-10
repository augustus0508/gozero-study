package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gozerotry/apps/user/rpc/userclient"
	"gozerotry/apps/video/api/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
