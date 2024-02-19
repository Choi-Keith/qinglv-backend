package comment

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleCommentModel = (*customArticleCommentModel)(nil)

type (
	// ArticleCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleCommentModel.
	ArticleCommentModel interface {
		articleCommentModel
	}

	customArticleCommentModel struct {
		*defaultArticleCommentModel
	}
)

// NewArticleCommentModel returns a model for the database table.
func NewArticleCommentModel(conn sqlx.SqlConn, c cache.CacheConf) ArticleCommentModel {
	return &customArticleCommentModel{
		defaultArticleCommentModel: newArticleCommentModel(conn, c),
	}
}
