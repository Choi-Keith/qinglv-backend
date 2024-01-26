package collection_group

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CollectionGroupModel = (*customCollectionGroupModel)(nil)

type (
	// CollectionGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCollectionGroupModel.
	CollectionGroupModel interface {
		collectionGroupModel
	}

	customCollectionGroupModel struct {
		*defaultCollectionGroupModel
	}
)

// NewCollectionGroupModel returns a model for the database table.
func NewCollectionGroupModel(conn sqlx.SqlConn, c cache.CacheConf) CollectionGroupModel {
	return &customCollectionGroupModel{
		defaultCollectionGroupModel: newCollectionGroupModel(conn, c),
	}
}
