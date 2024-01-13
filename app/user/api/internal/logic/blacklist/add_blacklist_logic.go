package blacklist

import (
	"context"
	"encoding/json"
	"errors"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user_client"
	"qinglv-backend/pkg/snowflake"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBlacklistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBlacklistLogic {
	return &AddBlacklistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddBlacklistLogic) AddBlacklist(req *types.AddBlackItemReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	checkResp, err := l.svcCtx.UserRpc.CheckBlackItem(l.ctx, &user_client.CheckBlackItemReq{
		UserId:      uint64(userId),
		BlackItemId: req.BlackItemId,
	})
	if err != nil {
		return err
	}
	if checkResp.IsBlackItem {
		return errors.New("该用户已被拉黑名单")
	}
	id := snowflake.MustID()
	_, err = l.svcCtx.UserRpc.AddBlackItem(l.ctx, &user_client.AddBlackItemReq{
		Id:          id,
		UserId:      uint64(userId),
		BlackItemId: req.BlackItemId,
	})
	if err != nil {
		return err
	}
	return nil
}
