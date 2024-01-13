package following

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFollowingLogic {
	return &IsFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsFollowingLogic) IsFollowing(req *types.IsFollowingReq) (resp *types.IsFollowingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
