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

type GetFollowingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFollowingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingListLogic {
	return &GetFollowingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFollowingListLogic) GetFollowingList(req *types.FollowingListReq) (resp *types.FollowingListResp, err error) {
	// todo: add your logic here and delete this line
	followingListResp, err := l.svcCtx.FollowingRpc.GetFollowingList(l.ctx, &user.GetFollowingListReq{
		UserId:   req.UserId,
		PageNum:  int32(req.PageNum),
		PageSize: int32(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	followingList := make([]types.Following, len(followingListResp.Data))

	for idx, following := range followingListResp.Data {
		_ = copier.Copy(&followingList[idx], following)
	}
	newFollowingList, err := mr.MapReduce(func(source chan<- types.Following) {
		for _, followingItem := range followingList {
			source <- followingItem
		}
	}, func(item types.Following, writer mr.Writer[types.Following], cancel func(error)) {
		followingItem := item
		userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdReq{
			UserId: followingItem.FollowingId,
		})
		if err != nil {
			cancel(err)
			return
		}
		_ = copier.Copy(&followingItem.FollowingUser, userResp.User)
		writer.Write(followingItem)
	}, func(pip <-chan types.Following, writer mr.Writer[[]types.Following], cancel func(error)) {
		var r []types.Following
		m := make(map[uint64]types.Following, len(followingList))
		for p := range pip {
			m[p.Id] = p
		}
		// 为了避免mapReduce多线程执行导致排序不一致的问题
		for _, followingItem := range followingList {
			r = append(r, m[followingItem.Id])
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
	if followingListResp.Total <= total {
		isEnd = true
	}
	return &types.FollowingListResp{
		List:  newFollowingList,
		Total: followingListResp.Total,
		IsEnd: isEnd,
	}, nil
}
