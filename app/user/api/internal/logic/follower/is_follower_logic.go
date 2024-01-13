package follower

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFollowerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFollowerLogic {
	return &IsFollowerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsFollowerLogic) IsFollower(req *types.IsFollowerReq) (resp *types.IsFollowerResp, err error) {
	// todo: add your logic here and delete this line

	return
}
