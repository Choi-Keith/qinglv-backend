package notificationclasslogic

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/model/notification"
	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddNotificationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddNotificationLogic {
	return &AddNotificationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddNotificationLogic) AddNotification(in *user.AddNotificationReq) (*user.OkResp, error) {
	// todo: add your logic here and delete this line
	if in.Type == 1 {
		if _, err := l.svcCtx.CommentNotifyModel.Insert(l.ctx, nil, &notification.CommentNotify{
			Id:             in.Id,
			SenderUserId:   in.SenderUserId,
			ReceiverUserId: in.ReceiverUserId,
			BizType:        uint64(in.BizType),
			CommentId:      in.CommentId,
			CommentContent: in.CommentContent,
			ReplyId:        in.ReplyId,
			ReplyContent:   in.ReplyContent,
			TargetId:       in.TargetId,
			TargetTitle:    in.TargetTitle,
			IsRead:         0,
			IsDel:          0,
			Version:        1,
		}); err != nil {
			return nil, err
		}
	}
	if in.Type == 2 {
		if _, err := l.svcCtx.LikeNotifyModel.Insert(l.ctx, nil, &notification.LikeNotify{
			Id:             in.Id,
			SenderUserId:   in.SenderUserId,
			ReceiverUserId: in.ReceiverUserId,
			TargetId:       in.TargetId,
			TargetTitle:    in.TargetTitle,
			BizType:        uint64(in.BizType),
			ActionType:     uint64(in.ActionType),
			IsRead:         0,
			IsDel:          0,
			Version:        1,
		}); err != nil {
			return nil, err
		}
	}
	if in.Type == 3 {
		if _, err := l.svcCtx.FollowNotifyModel.Insert(l.ctx, nil, &notification.FollowNotify{
			Id:             in.Id,
			SenderUserId:   in.SenderUserId,
			ReceiverUserId: in.ReceiverUserId,
			IsRead:         0,
			IsDel:          0,
			Version:        1,
		}); err != nil {
			return nil, err
		}
	}
	// 系统消息可能需要重构， 具体参考https://qsjzwithguang19forever.gitee.io/framework-learning/gitbook_doc/system_architecture_design/%E7%AB%99%E5%86%85%E6%B6%88%E6%81%AF%E7%B3%BB%E7%BB%9F%E7%9A%84%E8%AE%BE%E8%AE%A1.html
	if in.Type == 4 {
		if _, err := l.svcCtx.OsNotifyModel.Insert(l.ctx, nil, &notification.OsNotify{
			Id:             in.Id,
			SenderUserId:   0,
			ReceiverUserId: in.ReceiverUserId,
			Message:        in.Message,
			IsRead:         uint64(in.IsRead),
			IsDel:          0,
			Version:        1,
		}); err != nil {
			return nil, err
		}
	}
	return &user.OkResp{}, nil
}
