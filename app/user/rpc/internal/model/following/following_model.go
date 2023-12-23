package following

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FollowingModel = (*customFollowingModel)(nil)

type (
	// FollowingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowingModel.
	FollowingModel interface {
		followingModel
	}

	customFollowingModel struct {
		*defaultFollowingModel
	}
)

// NewFollowingModel returns a model for the database table.
func NewFollowingModel(conn sqlx.SqlConn, c cache.CacheConf) FollowingModel {
	return &customFollowingModel{
		defaultFollowingModel: newFollowingModel(conn, c),
	}
}
