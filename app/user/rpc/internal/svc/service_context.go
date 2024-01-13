package svc

import (
	"qinglv-backend/app/user/rpc/internal/config"
	"qinglv-backend/app/user/rpc/internal/model/following"
	"qinglv-backend/app/user/rpc/internal/model/role"
	"qinglv-backend/app/user/rpc/internal/model/user"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      user.UserModel
	RoleModel      role.RoleModel
	FollowingModel following.FollowingModel
	RedisClient    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:         c,
		UserModel:      user.NewUserModel(sqlConn, c.Cache),
		RoleModel:      role.NewRoleModel(sqlConn, c.Cache),
		FollowingModel: following.NewFollowingModel(sqlConn, c.Cache),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Pass = c.Redis.Pass
			r.Type = c.Redis.Type
		}),
	}
}
