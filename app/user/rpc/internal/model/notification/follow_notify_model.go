package notification

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FollowNotifyModel = (*customFollowNotifyModel)(nil)

type (
	// FollowNotifyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowNotifyModel.
	FollowNotifyModel interface {
		followNotifyModel
		UpdateAllUnreads(ctx context.Context) error
	}

	customFollowNotifyModel struct {
		*defaultFollowNotifyModel
	}
)

// NewFollowNotifyModel returns a model for the database table.
func NewFollowNotifyModel(conn sqlx.SqlConn) FollowNotifyModel {
	return &customFollowNotifyModel{
		defaultFollowNotifyModel: newFollowNotifyModel(conn),
	}
}

func (m *defaultFollowNotifyModel) UpdateAllUnreads(ctx context.Context) error {
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
