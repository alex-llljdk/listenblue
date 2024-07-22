package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"zjhx.com/live/gmodel"
	"zjhx.com/live/rpc/internal/config"
	"zjhx.com/pkg/orm"
	gmodel1 "zjhx.com/record/gmodel"
)

type ServiceContext struct {
	Config      config.Config
	DB          *orm.DB
	LiveModel   *gmodel.LiveRoomModel
	RecordModel *gmodel1.RecordModel
	BizRedis    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	rds, err := redis.NewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:      c,
		DB:          db,
		LiveModel:   gmodel.NewLiveRoomModel(db.DB),
		RecordModel: gmodel1.NewRecordModel(db.DB),
		BizRedis:    rds,
	}
}
