package following

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user_client"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
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
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	followingListResp, err := l.svcCtx.UserRpc.GetFollowingList(l.ctx, &user_client.GetFollowingListReq{
		UserId:      uint64(userId),
		FollowingId: req.FollowingId,
		PageNum:     int32(req.PageNum),
		PageSize:    int32(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	followingList := make([]types.Following, len(followingListResp.Data))

	for idx, following := range followingListResp.Data {
		_ = copier.Copy(&followingList[idx], following)
	}
	isEnd := false
	pageSize := uint64(req.PageSize)
	pageNum := uint64(req.PageNum)
	total := (pageNum-1)*pageSize + pageSize
	if followingListResp.Total < total {
		isEnd = true
	}
	return &types.FollowingListResp{
		List:  followingList,
		Total: followingListResp.Total,
		IsEnd: isEnd,
	}, nil
}
