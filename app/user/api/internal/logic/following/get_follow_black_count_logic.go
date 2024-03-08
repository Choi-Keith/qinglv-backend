package following

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowBlackCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowBlackCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowBlackCountLogic {
	return &GetFollowBlackCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowBlackCountLogic) GetFollowBlackCount(req *types.GetFollowBlackCountReq) (resp *types.GetFollowBlackCountResp, err error) {
	// todo: add your logic here and delete this line
	// 正在关注
	followingResp, err := l.svcCtx.FollowingRpc.GetFollowingList(l.ctx, &user.GetFollowingListReq{
		UserId:   req.UserId,
		PageNum:  1,
		PageSize: 20,
	})
	if err != nil {
		return nil, err
	}
	// 粉丝
	followerResp, err := l.svcCtx.FollowingRpc.GetFollowingList(l.ctx, &user.GetFollowingListReq{
		FollowingId: req.UserId,
		PageNum:     1,
		PageSize:    20,
	})
	if err != nil {
		return nil, err
	}
	blacklistResp, err := l.svcCtx.BlacklistRpc.GetBlackList(l.ctx, &user.GetBlackListReq{
		UserId:   req.UserId,
		PageNum:  1,
		PageSize: 20,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetFollowBlackCountResp{
		FollowerCount:  int64(followerResp.Total),
		FollowingCount: int64(followingResp.Total),
		BlacklistCount: int64(blacklistResp.Total),
	}, nil
}
