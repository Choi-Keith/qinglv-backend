package comment

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostCommentReplyModel = (*customPostCommentReplyModel)(nil)

type (
	// PostCommentReplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostCommentReplyModel.
	PostCommentReplyModel interface {
		postCommentReplyModel
	}

	customPostCommentReplyModel struct {
		*defaultPostCommentReplyModel
	}
)

// NewPostCommentReplyModel returns a model for the database table.
func NewPostCommentReplyModel(conn sqlx.SqlConn, c cache.CacheConf) PostCommentReplyModel {
	return &customPostCommentReplyModel{
		defaultPostCommentReplyModel: newPostCommentReplyModel(conn, c),
	}
}
