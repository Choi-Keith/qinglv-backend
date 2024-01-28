package thumb

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostCommentThumbModel = (*customPostCommentThumbModel)(nil)

type (
	// PostCommentThumbModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostCommentThumbModel.
	PostCommentThumbModel interface {
		postCommentThumbModel
	}

	customPostCommentThumbModel struct {
		*defaultPostCommentThumbModel
	}
)

// NewPostCommentThumbModel returns a model for the database table.
func NewPostCommentThumbModel(conn sqlx.SqlConn, c cache.CacheConf) PostCommentThumbModel {
	return &customPostCommentThumbModel{
		defaultPostCommentThumbModel: newPostCommentThumbModel(conn, c),
	}
}
