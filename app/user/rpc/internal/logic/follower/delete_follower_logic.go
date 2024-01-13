package follower

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFollowerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFollowerLogic {
	return &DeleteFollowerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: follower
func (l *DeleteFollowerLogic) DeleteFollower(in *user.DeleteFollowerReq) (*user.DeleteFollowerResp, error) {
	// todo: add your logic here and delete this line

	return &user.DeleteFollowerResp{}, nil
}
