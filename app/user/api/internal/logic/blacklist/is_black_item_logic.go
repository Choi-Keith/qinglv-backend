package blacklist

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsBlackItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsBlackItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsBlackItemLogic {
	return &IsBlackItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsBlackItemLogic) IsBlackItem(req *types.IsBlackItemReq) (resp *types.IsBlackItemResp, err error) {
	// todo: add your logic here and delete this line

	return
}
