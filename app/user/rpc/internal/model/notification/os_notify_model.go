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
		UpdateAllUneads(ctx context.Context) error
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

func (m *defaultOsNotifyModel) UpdateAllUneads(ctx context.Context) error {
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
