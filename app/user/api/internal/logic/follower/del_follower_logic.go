package follower

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelFollowerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFollowerLogic {
	return &DelFollowerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelFollowerLogic) DelFollower(req *types.DelFollowerReq) error {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserRpc.DeleteFollower(l.ctx, &user_client.DeleteFollowerReq{
		Id: req.Id,
	})
	if err != nil {
		return err
	}
	return nil
}
