package following

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFollowingLogic {
	return &DelFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelFollowingLogic) DelFollowing(req *types.DelFollowingReq) error {
	// todo: add your logic here and delete this line

	return nil
}
