package followingclasslogic

import (
	"context"
	"time"

	"qinglv-backend/app/user/rpc/internal/model/following"
	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFollowingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFollowingLogic {
	return &AddFollowingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: following
func (l *AddFollowingLogic) AddFollowing(in *user.AddFollowingReq) (*user.AddFollowingResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.FollowingModel.Insert(l.ctx, nil, &following.Following{
		Id:          in.Id,
		UserId:      in.UserId,
		FollowingId: in.FollowingId,
		DeletedAt:   time.Now(),
		Version:     1,
	})
	if err != nil {
		return nil, err
	}
	return &user.AddFollowingResp{}, nil
}
