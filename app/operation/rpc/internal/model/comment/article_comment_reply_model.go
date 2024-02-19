package comment

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleCommentReplyModel = (*customArticleCommentReplyModel)(nil)

type (
	// ArticleCommentReplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleCommentReplyModel.
	ArticleCommentReplyModel interface {
		articleCommentReplyModel
	}

	customArticleCommentReplyModel struct {
		*defaultArticleCommentReplyModel
	}
)

// NewArticleCommentReplyModel returns a model for the database table.
func NewArticleCommentReplyModel(conn sqlx.SqlConn, c cache.CacheConf) ArticleCommentReplyModel {
	return &customArticleCommentReplyModel{
		defaultArticleCommentReplyModel: newArticleCommentReplyModel(conn, c),
	}
}
