package blacklist

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBlacklistListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBlacklistListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlacklistListLogic {
	return &GetBlacklistListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBlacklistListLogic) GetBlacklistList(req *types.BlacklistListReq) (resp *types.BlacklistListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
