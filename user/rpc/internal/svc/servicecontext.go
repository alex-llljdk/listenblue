package svc

import (
	"zjhx.com/pkg/orm"
	"zjhx.com/user/gmodel"
	"zjhx.com/user/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	DB        *orm.DB
	UserModel *gmodel.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	return &ServiceContext{
		Config:    c,
		DB:        db,
		UserModel: gmodel.NewUserModel(db.DB),
	}
}
