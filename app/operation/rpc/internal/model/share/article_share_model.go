package share

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleShareModel = (*customArticleShareModel)(nil)

type (
	// ArticleShareModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleShareModel.
	ArticleShareModel interface {
		articleShareModel
	}

	customArticleShareModel struct {
		*defaultArticleShareModel
	}
)

// NewArticleShareModel returns a model for the database table.
func NewArticleShareModel(conn sqlx.SqlConn, c cache.CacheConf) ArticleShareModel {
	return &customArticleShareModel{
		defaultArticleShareModel: newArticleShareModel(conn, c),
	}
}
