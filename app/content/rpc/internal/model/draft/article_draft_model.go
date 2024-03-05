package draft

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleDraftModel = (*customArticleDraftModel)(nil)

type (
	// ArticleDraftModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleDraftModel.
	ArticleDraftModel interface {
		articleDraftModel
	}

	customArticleDraftModel struct {
		*defaultArticleDraftModel
	}
)

// NewArticleDraftModel returns a model for the database table.
func NewArticleDraftModel(conn sqlx.SqlConn) ArticleDraftModel {
	return &customArticleDraftModel{
		defaultArticleDraftModel: newArticleDraftModel(conn),
	}
}
