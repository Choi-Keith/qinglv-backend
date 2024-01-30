package comment

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostCommentReplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPostCommentReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostCommentReplyLogic {
	return &AddPostCommentReplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPostCommentReplyLogic) AddPostCommentReply(req *types.AddPostCommentReplyReq) error {
	// todo: add your logic here and delete this line

	return nil
}
