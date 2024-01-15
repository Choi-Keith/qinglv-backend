package blacklist

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user_client"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GetBlacklistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBlacklistLogic {
	return &GetBlacklistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBlacklistLogic) GetBlacklist(req *types.BlacklistReq) (resp *types.BlacklistResp, err error) {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	blackListResp, err := l.svcCtx.UserRpc.GetBlackList(l.ctx, &user_client.GetBlackListReq{
		UserId:      uint64(userId),
		BlackItemId: req.BlackItemId,
		PageNum:     int32(req.PageNum),
		PageSize:    int32(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	logx.Debugf("blackListResp: %+v\n", blackListResp)

	blackList := make([]types.BlackItem, len(blackListResp.Data))
	for idx, blackItem := range blackList {
		_ = copier.Copy(&blackList[idx], blackItem)
	}
	newBlackList, err := mr.MapReduce(func(source chan<- types.BlackItem) {
		for _, blackItem := range blackList {
			source <- blackItem
		}
	},
		func(item types.BlackItem, writer mr.Writer[types.BlackItem], cancel func(error)) {
			blackItem := item
			blackUser, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user_client.GetUserByIdReq{
				UserId: blackItem.BlackItemId,
			})
			if err != nil {
				cancel(err)
				return
			}
			_ = copier.Copy(&blackItem.BlackItemUser, blackUser)
		},
		func(pipe <-chan types.BlackItem, writer mr.Writer[[]types.BlackItem], cancel func(error)) {
			var r []types.BlackItem
			for p := range pipe {
				r = append(r, p)
			}
			writer.Write(r)
		})
	if err != nil {
		return nil, err
	}
	logx.Debugf("newBlackList: %+v\n", newBlackList)
	isEnd := false
	pageSize := uint64(req.PageSize)
	pageNum := uint64(req.PageNum)
	total := (pageNum-1)*pageSize + pageSize
	if blackListResp.Total < total {
		isEnd = true
	}
	return &types.BlacklistResp{
		List:  newBlackList,
		Total: total,
		IsEnd: isEnd,
	}, nil
}
