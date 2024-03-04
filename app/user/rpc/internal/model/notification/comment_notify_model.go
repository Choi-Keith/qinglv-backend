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
		UpdateAllUnreads(ctx context.Context, receiverUserId uint64) error
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

func (m *defaultCommentNotifyModel) UpdateAllUnreads(ctx context.Context, receiverUserId uint64) error {
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
