package media_file

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MediaFileModel = (*customMediaFileModel)(nil)

type (
	// MediaFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMediaFileModel.
	MediaFileModel interface {
		mediaFileModel
	}

	customMediaFileModel struct {
		*defaultMediaFileModel
	}
)

// NewMediaFileModel returns a model for the database table.
func NewMediaFileModel(conn sqlx.SqlConn, c cache.CacheConf) MediaFileModel {
	return &customMediaFileModel{
		defaultMediaFileModel: newMediaFileModel(conn, c),
	}
}
