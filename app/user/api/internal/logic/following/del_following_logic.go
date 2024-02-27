package following

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/pkg/event"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFollowingLogic {
	return &DelFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelFollowingLogic) DelFollowing(req *types.DelFollowingReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	followingResp, err := l.svcCtx.FollowingRpc.GetFollowingDetail(l.ctx, &user.GetFollowingDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return err
	}
	if _, err = l.svcCtx.FollowingRpc.DeleteFollowing(l.ctx, &user.DeleteFollowingReq{
		Id: req.Id,
	}); err != nil {
		return err
	}
	event.Send(event.UnFollowEvent{
		UserId:  uint64(userId),
		OtherId: followingResp.Data.FollowingId,
	})
	return nil
}
