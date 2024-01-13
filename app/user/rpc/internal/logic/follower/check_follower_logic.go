package follower

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckFollowerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckFollowerLogic {
	return &CheckFollowerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: follower
func (l *CheckFollowerLogic) CheckFollower(in *user.CheckFollowerReq) (*user.CheckFollowerResp, error) {
	// todo: add your logic here and delete this line

	return &user.CheckFollowerResp{}, nil
}
