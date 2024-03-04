package notification

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OsNotifyModel = (*customOsNotifyModel)(nil)

type (
	// OsNotifyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOsNotifyModel.
	OsNotifyModel interface {
		osNotifyModel
		UpdateAllUnreads(ctx context.Context, receiverUserId uint64) error
	}

	customOsNotifyModel struct {
		*defaultOsNotifyModel
	}
)

// NewOsNotifyModel returns a model for the database table.
func NewOsNotifyModel(conn sqlx.SqlConn) OsNotifyModel {
	return &customOsNotifyModel{
		defaultOsNotifyModel: newOsNotifyModel(conn),
	}
}

func (m *defaultOsNotifyModel) UpdateAllUnreads(ctx context.Context, receiverUserId uint64) error {
	query := fmt.Sprintf("update %s set %s where receiver_user_id=? ", m.table, "is_read=?")
	sqlResult, err := m.conn.ExecCtx(ctx, query, 1, receiverUserId)
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
