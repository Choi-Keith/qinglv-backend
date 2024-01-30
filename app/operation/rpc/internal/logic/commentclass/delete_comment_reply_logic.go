package commentclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentReplyLogic {
	return &DeleteCommentReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Comment
func (l *DeleteCommentReplyLogic) DeleteCommentReply(in *operation.DeleteCommentReplyReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line

	return &operation.OkResp{}, nil
}
