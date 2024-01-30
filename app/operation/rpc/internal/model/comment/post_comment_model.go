package comment

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostCommentModel = (*customPostCommentModel)(nil)

type (
	// PostCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostCommentModel.
	PostCommentModel interface {
		postCommentModel
	}

	customPostCommentModel struct {
		*defaultPostCommentModel
	}
)

// NewPostCommentModel returns a model for the database table.
func NewPostCommentModel(conn sqlx.SqlConn, c cache.CacheConf) PostCommentModel {
	return &customPostCommentModel{
		defaultPostCommentModel: newPostCommentModel(conn, c),
	}
}
