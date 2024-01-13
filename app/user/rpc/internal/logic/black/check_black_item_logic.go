package black

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckBlackItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckBlackItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckBlackItemLogic {
	return &CheckBlackItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: blacklist
func (l *CheckBlackItemLogic) CheckBlackItem(in *user.CheckBlackItemReq) (*user.CheckBlackItemResp, error) {
	// todo: add your logic here and delete this line

	return &user.CheckBlackItemResp{}, nil
}
