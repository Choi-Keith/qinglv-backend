package comment

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostCommentReplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostCommentReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostCommentReplyLogic {
	return &DeletePostCommentReplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostCommentReplyLogic) DeletePostCommentReply(req *types.DeletePostCommentReplyReq) error {
	// todo: add your logic here and delete this line

	return nil
}
