package following

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckFollowingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckFollowingLogic {
	return &CheckFollowingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: following
func (l *CheckFollowingLogic) CheckFollowing(in *user.CheckFollowingReq) (*user.CheckFollowingResp, error) {
	// todo: add your logic here and delete this line

	return &user.CheckFollowingResp{}, nil
}
