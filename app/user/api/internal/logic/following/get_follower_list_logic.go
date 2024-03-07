package following

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GetFollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowerListLogic) GetFollowerList(req *types.FollowerListReq) (resp *types.FollowerListResp, err error) {
	// todo: add your logic here and delete this line

	followerListResp, err := l.svcCtx.FollowingRpc.GetFollowingList(l.ctx, &user.GetFollowingListReq{
		FollowingId: req.UserId,
		PageNum:     int32(req.PageNum),
		PageSize:    int32(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	followerList := make([]types.Follower, len(followerListResp.Data))
	for idx, following := range followerListResp.Data {
		_ = copier.Copy(&followerList[idx], following)
	}
	newFollowerList, err := mr.MapReduce(func(source chan<- types.Follower) {
		for _, followingItem := range followerList {
			source <- followingItem
		}
	}, func(item types.Follower, writer mr.Writer[types.Follower], cancel func(error)) {
		followerItem := item
		userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdReq{
			UserId: followerItem.FollowingId,
		})
		if err != nil {
			cancel(err)
			return
		}
		_ = copier.Copy(&followerItem.FollowerUser, userResp.User)
		writer.Write(followerItem)
	}, func(pip <-chan types.Follower, writer mr.Writer[[]types.Follower], cancel func(error)) {
		var r []types.Follower
		for p := range pip {
			r = append(r, p)
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}

	isEnd := false
	pageSize := uint64(req.PageSize)
	pageNum := uint64(req.PageNum)
	total := (pageNum-1)*pageSize + pageSize
	if followerListResp.Total <= total {
		isEnd = true
	}
	return &types.FollowerListResp{
		List:  newFollowerList,
		Total: followerListResp.Total,
		IsEnd: isEnd,
	}, nil
	return
}
