package share

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostShareModel = (*customPostShareModel)(nil)

type (
	// PostShareModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostShareModel.
	PostShareModel interface {
		postShareModel
	}

	customPostShareModel struct {
		*defaultPostShareModel
	}
)

// NewPostShareModel returns a model for the database table.
func NewPostShareModel(conn sqlx.SqlConn, c cache.CacheConf) PostShareModel {
	return &customPostShareModel{
		defaultPostShareModel: newPostShareModel(conn, c),
	}
}
