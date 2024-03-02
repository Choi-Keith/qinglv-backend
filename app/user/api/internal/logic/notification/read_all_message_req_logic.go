package notification

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadAllMessageReqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadAllMessageReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadAllMessageReqLogic {
	return &ReadAllMessageReqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadAllMessageReqLogic) ReadAllMessageReq(req *types.ReadAllMessageReq) error {
	// todo: add your logic here and delete this line

	return nil
}
