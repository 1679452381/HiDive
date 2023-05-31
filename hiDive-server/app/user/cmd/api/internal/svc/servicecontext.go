package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"hiDive-server/app/user/cmd/api/internal/config"
	"hiDive-server/app/user/cmd/rpc/userclient"
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
