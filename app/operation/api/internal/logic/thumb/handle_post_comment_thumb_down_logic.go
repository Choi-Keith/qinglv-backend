package thumb

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandlePostCommentThumbDownLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlePostCommentThumbDownLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandlePostCommentThumbDownLogic {
	return &HandlePostCommentThumbDownLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlePostCommentThumbDownLogic) HandlePostCommentThumbDown(req *types.HandlePostCommentThumbDownReq) error {
	// todo: add your logic here and delete this line

	return nil
}
