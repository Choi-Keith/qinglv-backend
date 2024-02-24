package post

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostFeedModel = (*customPostFeedModel)(nil)

type (
	// PostFeedModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostFeedModel.
	PostFeedModel interface {
		postFeedModel
	}

	customPostFeedModel struct {
		*defaultPostFeedModel
	}
)

// NewPostFeedModel returns a model for the database table.
func NewPostFeedModel(conn sqlx.SqlConn, c cache.CacheConf) PostFeedModel {
	return &customPostFeedModel{
		defaultPostFeedModel: newPostFeedModel(conn, c),
	}
}
