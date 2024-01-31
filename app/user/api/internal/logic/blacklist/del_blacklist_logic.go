package blacklist

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelBlacklistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelBlacklistLogic {
	return &DelBlacklistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelBlacklistLogic) DelBlacklist(req *types.DelBlackItemReq) error {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.BlacklistRpc.DeleteBlackItem(l.ctx, &user.DeleteBlackItemReq{
		Id: req.Id,
	})
	if err != nil {
		return err
	}
	return nil
}
