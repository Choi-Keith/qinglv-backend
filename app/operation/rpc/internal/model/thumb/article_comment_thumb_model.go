package thumb

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleCommentThumbModel = (*customArticleCommentThumbModel)(nil)

type (
	// ArticleCommentThumbModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleCommentThumbModel.
	ArticleCommentThumbModel interface {
		articleCommentThumbModel
	}

	customArticleCommentThumbModel struct {
		*defaultArticleCommentThumbModel
	}
)

// NewArticleCommentThumbModel returns a model for the database table.
func NewArticleCommentThumbModel(conn sqlx.SqlConn, c cache.CacheConf) ArticleCommentThumbModel {
	return &customArticleCommentThumbModel{
		defaultArticleCommentThumbModel: newArticleCommentThumbModel(conn, c),
	}
}
