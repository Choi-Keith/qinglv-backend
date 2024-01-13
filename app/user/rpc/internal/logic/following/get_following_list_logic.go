package following

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingListLogic {
	return &GetFollowingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: following
func (l *GetFollowingListLogic) GetFollowingList(in *user.GetFollowingListReq) (*user.GetFollowingListResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetFollowingListResp{}, nil
}
