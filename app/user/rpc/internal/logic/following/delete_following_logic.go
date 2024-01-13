package following

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFollowingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFollowingLogic {
	return &DeleteFollowingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: following
func (l *DeleteFollowingLogic) DeleteFollowing(in *user.DeleteFollowingReq) (*user.DeleteFollowingResp, error) {
	// todo: add your logic here and delete this line
	followingItem, err := l.svcCtx.FollowingModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.FollowingModel.DeleteSoft(l.ctx, nil, followingItem)
	if err != nil {
		return nil, err
	}
	return &user.DeleteFollowingResp{}, nil
}
