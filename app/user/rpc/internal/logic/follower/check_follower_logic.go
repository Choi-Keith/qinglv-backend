package follower

import (
	"context"
	"errors"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckFollowerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckFollowerLogic {
	return &CheckFollowerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: follower
func (l *CheckFollowerLogic) CheckFollower(in *user.CheckFollowerReq) (*user.CheckFollowerResp, error) {
	// todo: add your logic here and delete this line
	if in.UserId == 0 || in.FollowerId == 0 {
		return nil, errors.New("followerId和userId不能为空")
	}
	whereBuilder := l.svcCtx.FollowerModel.SelectBuilder()
	followerList, err := l.svcCtx.FollowerModel.FindAll(l.ctx, whereBuilder, "")
	if err != nil {
		return nil, err
	}
	return &user.CheckFollowerResp{
		IsFollower: len(followerList) > 0,
	}, nil
}
