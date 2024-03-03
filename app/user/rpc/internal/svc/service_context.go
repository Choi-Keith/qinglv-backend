package svc

import (
	"qinglv-backend/app/user/rpc/internal/config"
	"qinglv-backend/app/user/rpc/internal/model/black"
	"qinglv-backend/app/user/rpc/internal/model/following"
	"qinglv-backend/app/user/rpc/internal/model/notification"
	"qinglv-backend/app/user/rpc/internal/model/role"
	"qinglv-backend/app/user/rpc/internal/model/user"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	UserModel          user.UserModel
	RoleModel          role.RoleModel
	FollowingModel     following.FollowingModel
	BlacklistModel     black.BlacklistModel
	CommentNotifyModel notification.CommentNotifyModel
	LikeNotifyModel    notification.LikeNotifyModel
	FollowNotifyModel  notification.FollowNotifyModel
	OsNotifyModel      notification.OsNotifyModel
	RedisClient        *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:             c,
		UserModel:          user.NewUserModel(sqlConn, c.Cache),
		RoleModel:          role.NewRoleModel(sqlConn, c.Cache),
		FollowingModel:     following.NewFollowingModel(sqlConn, c.Cache),
		BlacklistModel:     black.NewBlacklistModel(sqlConn, c.Cache),
		CommentNotifyModel: notification.NewCommentNotifyModel(sqlConn),
		FollowNotifyModel:  notification.NewFollowNotifyModel(sqlConn),
		LikeNotifyModel:    notification.NewLikeNotifyModel(sqlConn),
		OsNotifyModel:      notification.NewOsNotifyModel(sqlConn),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Pass = c.Redis.Pass
			r.Type = c.Redis.Type
		}),
	}
}
