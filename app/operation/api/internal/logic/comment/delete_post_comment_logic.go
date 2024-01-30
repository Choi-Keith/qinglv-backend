package comment

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostCommentLogic {
	return &DeletePostCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostCommentLogic) DeletePostComment(req *types.DeletePostCommentReq) error {
	// todo: add your logic here and delete this line

	return nil
}
