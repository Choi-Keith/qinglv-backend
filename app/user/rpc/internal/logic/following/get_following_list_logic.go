package following

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingListLogic {
	return &GetFollowingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: following
func (l *GetFollowingListLogic) GetFollowingList(in *user.GetFollowingListReq) (*user.GetFollowingListResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.FollowingModel.SelectBuilder()
	if in.UserId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"user_id": in.UserId,
		})
	}
	if in.FollowingId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"following_id": in.FollowingId,
		})
	}
	followingListResp, total, err := l.svcCtx.FollowingModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, int64(in.PageNum), int64(in.PageSize), "")
	if err != nil {
		return nil, err
	}
	followingList := make([]*user.FollowingItem, len(followingListResp))
	for idx, followingItem := range followingListResp {
		followingList[idx] = &user.FollowingItem{
			Id:          followingItem.Id,
			UserId:      followingItem.UserId,
			FollowingId: followingItem.FollowingId,
			CreatedAt:   uint64(followingItem.CreatedAt.Unix() * 1000),
			UpdatedAt:   uint64(followingItem.UpdatedAt.Unix() * 1000),
		}
	}
	return &user.GetFollowingListResp{
		Total: uint64(total),
		Data:  followingList,
	}, nil
}
