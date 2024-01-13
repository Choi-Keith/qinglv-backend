package follower

import (
	"context"
	"time"

	"qinglv-backend/app/user/rpc/internal/model/follower"
	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFollowerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFollowerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFollowerLogic {
	return &AddFollowerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: follower
func (l *AddFollowerLogic) AddFollower(in *user.AddFollowerReq) (*user.AddFollowerResp, error) {
	// todo: add your logic here and delete this line

	_, err := l.svcCtx.FollowerModel.Insert(l.ctx, nil, &follower.Follower{
		Id:         in.Id,
		UserId:     in.UserId,
		FollowerId: in.FollowerId,
		DeletedAt:  time.Now(),
		Version:    1,
	})
	if err != nil {
		return nil, err
	}
	return &user.AddFollowerResp{}, nil
}
