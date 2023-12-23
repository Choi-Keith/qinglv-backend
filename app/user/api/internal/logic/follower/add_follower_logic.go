package follower

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFollowerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFollowerLogic {
	return &AddFollowerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFollowerLogic) AddFollower(req *types.AddFollowerReq) error {
	// todo: add your logic here and delete this line

	return nil
}
