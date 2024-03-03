package notification

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentNotifyModel = (*customCommentNotifyModel)(nil)

type (
	// CommentNotifyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentNotifyModel.
	CommentNotifyModel interface {
		commentNotifyModel
		UpdateAllUneads(ctx context.Context) error
	}

	customCommentNotifyModel struct {
		*defaultCommentNotifyModel
	}
)

// NewCommentNotifyModel returns a model for the database table.
func NewCommentNotifyModel(conn sqlx.SqlConn) CommentNotifyModel {
	return &customCommentNotifyModel{
		defaultCommentNotifyModel: newCommentNotifyModel(conn),
	}
}

func (m *defaultCommentNotifyModel) UpdateAllUneads(ctx context.Context) error {
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
