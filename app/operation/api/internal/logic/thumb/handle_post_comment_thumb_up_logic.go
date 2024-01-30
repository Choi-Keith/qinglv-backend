package thumb

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandlePostCommentThumbUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlePostCommentThumbUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandlePostCommentThumbUpLogic {
	return &HandlePostCommentThumbUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlePostCommentThumbUpLogic) HandlePostCommentThumbUp(req *types.HandlePostCommentThumbUpReq) error {
	// todo: add your logic here and delete this line

	return nil
}
