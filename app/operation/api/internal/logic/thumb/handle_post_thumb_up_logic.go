package thumb

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandlePostThumbUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlePostThumbUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandlePostThumbUpLogic {
	return &HandlePostThumbUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlePostThumbUpLogic) HandlePostThumbUp(req *types.HandlePostThumbUpReq) error {
	// todo: add your logic here and delete this line

	return nil
}
