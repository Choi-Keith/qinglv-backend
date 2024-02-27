package followingclasslogic

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowingDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingDetailLogic {
	return &GetFollowingDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowingDetailLogic) GetFollowingDetail(in *user.GetFollowingDetailReq) (*user.GetFollowingDetailResp, error) {
	// todo: add your logic here and delete this line
	followingResp, err := l.svcCtx.FollowingModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	followingItem := &user.FollowingItem{
		Id:          followingResp.Id,
		UserId:      followingResp.UserId,
		FollowingId: followingResp.FollowingId,
		CreatedAt:   uint64(followingResp.CreatedAt.Unix() * 1000),
		UpdatedAt:   uint64(followingResp.UpdatedAt.Unix() * 1000),
	}
	return &user.GetFollowingDetailResp{
		Data: followingItem,
	}, nil
}
