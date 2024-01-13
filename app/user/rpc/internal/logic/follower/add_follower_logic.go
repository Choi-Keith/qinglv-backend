package follower

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFollowerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFollowerLogic {
	return &AddFollowerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: follower
func (l *AddFollowerLogic) AddFollower(in *user.AddFollowerReq) (*user.AddFollowerResp, error) {
	// todo: add your logic here and delete this line

	return &user.AddFollowerResp{}, nil
}
