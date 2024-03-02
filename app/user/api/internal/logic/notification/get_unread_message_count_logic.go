package notification

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUnreadMessageCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUnreadMessageCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUnreadMessageCountLogic {
	return &GetUnreadMessageCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUnreadMessageCountLogic) GetUnreadMessageCount() (resp *types.GetMessageCountResp, err error) {
	// todo: add your logic here and delete this line

	return
}
