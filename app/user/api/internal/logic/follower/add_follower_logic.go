package follower

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
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	checkResp, err := l.svcCtx.UserRpc.CheckFollower(l.ctx, &user_client.CheckFollowerReq{
		UserId:     uint64(userId),
		FollowerId: req.FollowerId,
	})
	if err != nil {
		return err
	}
	if checkResp.IsFollower {
		return errors.New("该用户已关注")
	}
	id := snowflake.MustID()
	_, err = l.svcCtx.UserRpc.AddFollower(l.ctx, &user_client.AddFollowerReq{
		Id:         id,
		UserId:     uint64(userId),
		FollowerId: req.FollowerId,
	})
	if err != nil {
		return err
	}
	return nil
}
