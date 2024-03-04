package notification

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"

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
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	counter, err := l.svcCtx.NotificationRpc.GetUnreadsNotificationCount(l.ctx, &user.GetUnreadsNotificationCountReq{
		ReceiverUserId: uint64(userId),
	})
	if err != nil {
		return nil, err
	}

	return &types.GetMessageCountResp{
		Comment: counter.CommentCount,
		Follow:  counter.FollowCount,
		Like:    counter.LikeCount,
		Os:      counter.OsCount,
		Total:   counter.Total,
	}, nil
}
