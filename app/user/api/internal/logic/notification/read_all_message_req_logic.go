package notification

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadAllMessageReqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadAllMessageReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadAllMessageReqLogic {
	return &ReadAllMessageReqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadAllMessageReqLogic) ReadAllMessageReq(req *types.ReadAllMessageReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	if _, err := l.svcCtx.NotificationRpc.SetAllReadNotification(l.ctx, &user.SetAllReadNotificationReq{
		Type:           req.Type,
		ReceiverUserId: uint64(userId),
	}); err != nil {
		return err
	}
	return nil
}
