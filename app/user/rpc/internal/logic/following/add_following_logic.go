package following

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFollowingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFollowingLogic {
	return &AddFollowingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: following
func (l *AddFollowingLogic) AddFollowing(in *user.AddFollowingReq) (*user.AddFollowingResp, error) {
	// todo: add your logic here and delete this line

	return &user.AddFollowingResp{}, nil
}
