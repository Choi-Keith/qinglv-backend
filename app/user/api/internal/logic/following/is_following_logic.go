package following

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFollowingLogic {
	return &IsFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsFollowingLogic) IsFollowing(req *types.IsFollowingReq) (resp *types.IsFollowingResp, err error) {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	checkResp, err := l.svcCtx.FollowingRpc.CheckFollowing(l.ctx, &user.CheckFollowingReq{
		UserId:      uint64(userId),
		FollowingId: req.FollowingId,
	})
	if err != nil {
		return nil, err
	}
	return &types.IsFollowingResp{
		IsFollowing: checkResp.IsFollowing,
	}, nil
}
