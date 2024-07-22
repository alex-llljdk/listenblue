package svc

import (
	"zjhx.com/file/gmodel"
	"zjhx.com/file/rpc/internal/config"
	"zjhx.com/pkg/orm"
)

type ServiceContext struct {
	Config      config.Config
	DB          *orm.DB
	RecordModel *gmodel.RecordModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	return &ServiceContext{
		Config:      c,
		DB:          db,
		RecordModel: gmodel.NewRecordModel(db.DB),
	}
}
