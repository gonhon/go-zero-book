package svc

import (
	"github.com/honsky/go-zero-book/service/search/api/internal/config"
	"github.com/honsky/go-zero-book/service/search/api/internal/middleware"
	"github.com/honsky/go-zero-book/service/user/rpc/userclient"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
