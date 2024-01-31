package followingclasslogic

import (
	"context"
	"errors"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type CheckFollowingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckFollowingLogic {
	return &CheckFollowingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: following
func (l *CheckFollowingLogic) CheckFollowing(in *user.CheckFollowingReq) (*user.CheckFollowingResp, error) {
	// todo: add your logic here and delete this line
	if in.FollowingId == 0 || in.UserId == 0 {
		return nil, errors.New("follwingId和userId不能为空")
	}
	whereBuilder := l.svcCtx.FollowingModel.SelectBuilder().Where(squirrel.Eq{
		"following_id": in.FollowingId,
		"user_id":      in.UserId,
	})

	followingList, err := l.svcCtx.FollowingModel.FindAll(l.ctx, whereBuilder, "")
	if err != nil {
		return nil, err
	}
	return &user.CheckFollowingResp{
		IsFollowing: len(followingList) > 0,
	}, nil
}
