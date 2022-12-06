package svc

import (
	"github.com/gonhon/go-zero-book/service/user/model"
	"github.com/gonhon/go-zero-book/service/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserMudel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserMudel: model.NewUserModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
	}
}
