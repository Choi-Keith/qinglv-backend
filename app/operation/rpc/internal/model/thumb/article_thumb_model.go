package thumb

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleThumbModel = (*customArticleThumbModel)(nil)

type (
	// ArticleThumbModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleThumbModel.
	ArticleThumbModel interface {
		articleThumbModel
	}

	customArticleThumbModel struct {
		*defaultArticleThumbModel
	}
)

// NewArticleThumbModel returns a model for the database table.
func NewArticleThumbModel(conn sqlx.SqlConn, c cache.CacheConf) ArticleThumbModel {
	return &customArticleThumbModel{
		defaultArticleThumbModel: newArticleThumbModel(conn, c),
	}
}
