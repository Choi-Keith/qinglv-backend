package notification

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LikeNotifyModel = (*customLikeNotifyModel)(nil)

type (
	// LikeNotifyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLikeNotifyModel.
	LikeNotifyModel interface {
		likeNotifyModel
		UpdateAllUnreads(ctx context.Context) error
	}

	customLikeNotifyModel struct {
		*defaultLikeNotifyModel
	}
)

// NewLikeNotifyModel returns a model for the database table.
func NewLikeNotifyModel(conn sqlx.SqlConn) LikeNotifyModel {
	return &customLikeNotifyModel{
		defaultLikeNotifyModel: newLikeNotifyModel(conn),
	}
}

func (m *defaultLikeNotifyModel) UpdateAllUnreads(ctx context.Context) error {
	query := fmt.Sprintf("update %s set %s where is_read=0", m.table, "is_read=?")
	sqlResult, err := m.conn.ExecCtx(ctx, query, 1)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return ErrNoRowsUpdate
	}
	return nil
}
