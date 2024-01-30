package thumb

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandlePostCommentReplyThumbUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlePostCommentReplyThumbUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandlePostCommentReplyThumbUpLogic {
	return &HandlePostCommentReplyThumbUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlePostCommentReplyThumbUpLogic) HandlePostCommentReplyThumbUp(req *types.HandlePostCommentReplyThumbUpReq) error {
	// todo: add your logic here and delete this line

	return nil
}
