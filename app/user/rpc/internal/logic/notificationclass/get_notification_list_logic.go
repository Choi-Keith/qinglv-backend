package notificationclasslogic

import (
	"context"
	"errors"

	notificationModel "qinglv-backend/app/user/rpc/internal/model/notification"
	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetNotificationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNotificationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNotificationListLogic {
	return &GetNotificationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNotificationListLogic) GetNotificationList(in *user.GetNotificationListReq) (*user.GetNotificationListResp, error) {
	// todo: add your logic here and delete this line
	if in.ReceiverUserId == 0 {
		return nil, errors.New("接收消息用户ID不能为空")
	}
	if in.Type == 1 {
		whereBuilder := l.svcCtx.CommentNotifyModel.SelectBuilder().Where(squirrel.Eq{
			"receiver_user_id": in.ReceiverUserId,
		})
		notificationResp, total, err := l.svcCtx.CommentNotifyModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, in.PageNum, in.PageSize, "")
		if err != nil {
			return nil, err
		}
		notificationList := make([]*user.NotificationItem, len(notificationResp))
		for idx, notificationItem := range notificationResp {
			notificationList[idx] = genCommentNotificationItem(notificationItem)
		}
		return &user.GetNotificationListResp{
			Data:  notificationList,
			Total: uint64(total),
		}, nil

	}
	if in.Type == 2 {
		whereBuilder := l.svcCtx.FollowNotifyModel.SelectBuilder().Where(squirrel.Eq{
			"receiver_user_id": in.ReceiverUserId,
		})
		notificationResp, total, err := l.svcCtx.FollowNotifyModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, in.PageNum, in.PageSize, "")
		if err != nil {
			return nil, err
		}
		notificationList := make([]*user.NotificationItem, len(notificationResp))
		for idx, notificationItem := range notificationResp {
			notificationList[idx] = genFollowNotificationItem(notificationItem)
		}
		return &user.GetNotificationListResp{
			Data:  notificationList,
			Total: uint64(total),
		}, nil
	}
	if in.Type == 3 {
		whereBuilder := l.svcCtx.LikeNotifyModel.SelectBuilder().Where(squirrel.Eq{
			"receiver_user_id": in.ReceiverUserId,
		})
		notificationResp, total, err := l.svcCtx.LikeNotifyModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, in.PageNum, in.PageSize, "")
		if err != nil {
			return nil, err
		}
		notificationList := make([]*user.NotificationItem, len(notificationResp))
		for idx, notificationItem := range notificationResp {
			notificationList[idx] = genLikeNotificationItem(notificationItem)
		}
		return &user.GetNotificationListResp{
			Data:  notificationList,
			Total: uint64(total),
		}, nil
	}
	if in.Type == 4 {
		whereBuilder := l.svcCtx.OsNotifyModel.SelectBuilder().Where(squirrel.Eq{
			"receiver_user_id": in.ReceiverUserId,
		})
		notificationResp, total, err := l.svcCtx.OsNotifyModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, in.PageNum, in.PageSize, "")
		if err != nil {
			return nil, err
		}
		notificationList := make([]*user.NotificationItem, len(notificationResp))
		for idx, notificationItem := range notificationResp {
			notificationList[idx] = genOsNotificationItem(notificationItem)
		}
		return &user.GetNotificationListResp{
			Data:  notificationList,
			Total: uint64(total),
		}, nil
	}

	return &user.GetNotificationListResp{}, nil
}

func genCommentNotificationItem(item *notificationModel.CommentNotify) *user.NotificationItem {
	return &user.NotificationItem{
		Id:             item.Id,
		SenderUserId:   item.SenderUserId,
		ReceiverUserId: item.ReceiverUserId,
		CommentId:      item.CommentId,
		CommentContent: item.CommentContent,
		ReplyId:        item.ReplyId,
		ReplyContent:   item.ReplyContent,
		Message:        "",
		TargetId:       item.TargetId,
		TargetTitle:    item.TargetTitle,
		BizType:        int32(item.BizType),
		IsRead:         int32(item.IsRead),
		ActionType:     0,
		CreatedAt:      uint64(item.CreatedAt.Unix() * 1000),
		UpdatedAt:      uint64(item.UpdatedAt.Unix() * 1000),
	}
}

func genFollowNotificationItem(item *notificationModel.FollowNotify) *user.NotificationItem {
	return &user.NotificationItem{
		Id:             item.Id,
		SenderUserId:   item.SenderUserId,
		ReceiverUserId: item.ReceiverUserId,
		CommentId:      0,
		CommentContent: "",
		ReplyId:        0,
		ReplyContent:   "",
		Message:        "",
		TargetId:       0,
		TargetTitle:    "",
		BizType:        0,
		ActionType:     0,
		IsRead:         int32(item.IsRead),
		CreatedAt:      uint64(item.CreatedAt.Unix() * 1000),
		UpdatedAt:      uint64(item.UpdatedAt.Unix() * 1000),
	}
}

func genLikeNotificationItem(item *notificationModel.LikeNotify) *user.NotificationItem {
	return &user.NotificationItem{
		Id:             item.Id,
		SenderUserId:   item.SenderUserId,
		ReceiverUserId: item.ReceiverUserId,
		CommentId:      0,
		CommentContent: "",
		ReplyId:        0,
		ReplyContent:   "",
		Message:        "",
		TargetId:       item.TargetId,
		TargetTitle:    item.TargetTitle,
		ActionType:     int32(item.ActionType),
		BizType:        int32(item.BizType),
		IsRead:         int32(item.IsRead),
		CreatedAt:      uint64(item.CreatedAt.Unix() * 1000),
		UpdatedAt:      uint64(item.UpdatedAt.Unix() * 1000),
	}
}

func genOsNotificationItem(item *notificationModel.OsNotify) *user.NotificationItem {
	return &user.NotificationItem{
		Id:             item.Id,
		SenderUserId:   item.SenderUserId,
		ReceiverUserId: item.ReceiverUserId,
		CommentId:      0,
		CommentContent: "",
		ReplyId:        0,
		ReplyContent:   "",
		Message:        item.Message,
		TargetId:       0,
		TargetTitle:    "",
		ActionType:     0,
		BizType:        0,
		IsRead:         int32(item.IsRead),
		CreatedAt:      uint64(item.CreatedAt.Unix() * 1000),
		UpdatedAt:      uint64(item.UpdatedAt.Unix() * 1000),
	}
}
