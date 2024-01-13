package following

import (
	"context"
	"encoding/json"
	"errors"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user_client"
	"qinglv-backend/pkg/snowflake"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFollowingLogic {
	return &AddFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFollowingLogic) AddFollowing(req *types.AddFollowingReq) error {
	// todo: add your logic here and delete this line

	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	checkResp, err := l.svcCtx.UserRpc.CheckFollowing(l.ctx, &user_client.CheckFollowingReq{
		UserId:      uint64(userId),
		FollowingId: req.FollowingId,
	})
	if err != nil {
		return err
	}
	if checkResp.IsFollowing {
		return errors.New("该用户已关注")
	}
	id := snowflake.MustID()
	_, err = l.svcCtx.UserRpc.AddFollowing(l.ctx, &user_client.AddFollowingReq{
		Id:          id,
		UserId:      uint64(userId),
		FollowingId: req.FollowingId,
	})
	if err != nil {
		return err
	}
	return nil
}
