package svc

import (
	"qinglv-backend/app/operation/rpc/internal/config"
	"qinglv-backend/app/operation/rpc/internal/model/collection_group"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config               config.Config
	CollectionGroupModel collection_group.CollectionGroupModel
	RedisClient          *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:               c,
		CollectionGroupModel: collection_group.NewCollectionGroupModel(sqlConn, c.Cache),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Pass = c.Redis.Pass
			r.Type = c.Redis.Type
		}),
	}
}
