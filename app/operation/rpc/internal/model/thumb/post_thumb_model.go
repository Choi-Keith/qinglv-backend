package thumb

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostThumbModel = (*customPostThumbModel)(nil)

type (
	// PostThumbModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostThumbModel.
	PostThumbModel interface {
		postThumbModel
	}

	customPostThumbModel struct {
		*defaultPostThumbModel
	}
)

// NewPostThumbModel returns a model for the database table.
func NewPostThumbModel(conn sqlx.SqlConn, c cache.CacheConf) PostThumbModel {
	return &customPostThumbModel{
		defaultPostThumbModel: newPostThumbModel(conn, c),
	}
}
