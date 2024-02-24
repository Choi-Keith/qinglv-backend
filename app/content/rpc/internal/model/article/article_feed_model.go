package article

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleFeedModel = (*customArticleFeedModel)(nil)

type (
	// ArticleFeedModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleFeedModel.
	ArticleFeedModel interface {
		articleFeedModel
	}

	customArticleFeedModel struct {
		*defaultArticleFeedModel
	}
)

// NewArticleFeedModel returns a model for the database table.
func NewArticleFeedModel(conn sqlx.SqlConn, c cache.CacheConf) ArticleFeedModel {
	return &customArticleFeedModel{
		defaultArticleFeedModel: newArticleFeedModel(conn, c),
	}
}
