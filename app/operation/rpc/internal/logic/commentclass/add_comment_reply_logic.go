package commentclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentReplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentReplyLogic {
	return &AddCommentReplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: Comment
func (l *AddCommentReplyLogic) AddCommentReply(in *operation.AddCommentReplyReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line

	return &operation.OkResp{}, nil
}
