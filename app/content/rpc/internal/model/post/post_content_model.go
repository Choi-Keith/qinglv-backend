package post

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostContentModel = (*customPostContentModel)(nil)

type (
	// PostContentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostContentModel.
	PostContentModel interface {
		postContentModel
	}

	customPostContentModel struct {
		*defaultPostContentModel
	}
)

// NewPostContentModel returns a model for the database table.
func NewPostContentModel(conn sqlx.SqlConn, c cache.CacheConf) PostContentModel {
	return &customPostContentModel{
		defaultPostContentModel: newPostContentModel(conn, c),
	}
}
