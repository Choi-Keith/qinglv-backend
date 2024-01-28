package thumb

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandlePostThumbDownLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandlePostThumbDownLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandlePostThumbDownLogic {
	return &HandlePostThumbDownLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandlePostThumbDownLogic) HandlePostThumbDown(req *types.HandlePostThumbDownReq) error {
	// todo: add your logic here and delete this line

	return nil
}
