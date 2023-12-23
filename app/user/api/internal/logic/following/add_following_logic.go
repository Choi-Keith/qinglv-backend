package following

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFollowingLogic {
	return &AddFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFollowingLogic) AddFollowing(req *types.AddFollowingReq) error {
	// todo: add your logic here and delete this line

	return nil
}
