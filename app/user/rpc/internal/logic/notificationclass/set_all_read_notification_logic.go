package notificationclasslogic

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetAllReadNotificationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetAllReadNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetAllReadNotificationLogic {
	return &SetAllReadNotificationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetAllReadNotificationLogic) SetAllReadNotification(in *user.SetAllReadNotificationReq) (*user.OkResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		if err := l.svcCtx.CommentNotifyModel.UpdateAllUnreads(l.ctx); err != nil {
			return nil, err
		}
	}
	if in.Type == 2 {
		if err := l.svcCtx.LikeNotifyModel.UpdateAllUnreads(l.ctx); err != nil {
			return nil, err
		}
	}
	if in.Type == 3 {
		if err := l.svcCtx.FollowNotifyModel.UpdateAllUnreads(l.ctx); err != nil {
			return nil, err
		}
	}
	if in.Type == 4 {
		if err := l.svcCtx.OsNotifyModel.UpdateAllUnreads(l.ctx); err != nil {
			return nil, err
		}
	}
	return &user.OkResp{}, nil
}
