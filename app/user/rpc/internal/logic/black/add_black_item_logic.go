package black

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddBlackItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddBlackItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBlackItemLogic {
	return &AddBlackItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: blacklist
func (l *AddBlackItemLogic) AddBlackItem(in *user.AddBlackItemReq) (*user.AddBlackItemResp, error) {
	// todo: add your logic here and delete this line

	return &user.AddBlackItemResp{}, nil
}
