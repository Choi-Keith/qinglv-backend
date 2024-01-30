package thumb

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandlePostCommentReplyThumbDownLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlePostCommentReplyThumbDownLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandlePostCommentReplyThumbDownLogic {
	return &HandlePostCommentReplyThumbDownLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlePostCommentReplyThumbDownLogic) HandlePostCommentReplyThumbDown(req *types.HandlePostCommentReplyThumbDownReq) error {
	// todo: add your logic here and delete this line

	return nil
}
