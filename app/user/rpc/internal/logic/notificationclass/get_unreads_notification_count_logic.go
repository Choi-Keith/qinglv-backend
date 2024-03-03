package notificationclasslogic

import (
	"context"
	"errors"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type GetUnreadsNotificationCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUnreadsNotificationCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUnreadsNotificationCountLogic {
	return &GetUnreadsNotificationCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUnreadsNotificationCountLogic) GetUnreadsNotificationCount(in *user.GetUnreadsNotificationCountReq) (*user.GetUnreadsNotificationCountResp, error) {
	// todo: add your logic here and delete this line
	if in.ReceiverUserId == 0 {
		return nil, errors.New("接收消息用户ID不能为空")
	}
	var (
		err          error
		commentCount int64
		followCount  int64
		likeCount    int64
		osCount      int64
	)
	if err := l.svcCtx.CommentNotifyModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		commentBuilder := l.svcCtx.CommentNotifyModel.SelectBuilder().Where(squirrel.Eq{
			"is_read":          0,
			"receiver_user_id": in.ReceiverUserId,
		})
		commentCount, err = l.svcCtx.CommentNotifyModel.FindCount(l.ctx, commentBuilder, "receiver_user_id")
		if err != nil {
			return err
		}
		followBuilder := l.svcCtx.FollowNotifyModel.SelectBuilder().Where(squirrel.Eq{
			"is_read":          0,
			"receiver_user_id": in.ReceiverUserId,
		})
		followCount, err = l.svcCtx.FollowNotifyModel.FindCount(l.ctx, followBuilder, "receiver_user_id")
		if err != nil {
			return err
		}
		likeBuilder := l.svcCtx.LikeNotifyModel.SelectBuilder().Where(squirrel.Eq{
			"is_read":          0,
			"receiver_user_id": in.ReceiverUserId,
		})
		likeCount, err = l.svcCtx.LikeNotifyModel.FindCount(l.ctx, likeBuilder, "receiver_user_id")
		if err != nil {
			return err
		}
		osBuilder := l.svcCtx.OsNotifyModel.SelectBuilder().Where(squirrel.Eq{
			"is_read":          0,
			"receiver_user_id": in.ReceiverUserId,
		})
		osCount, err = l.svcCtx.OsNotifyModel.FindCount(l.ctx, osBuilder, "receiver_user_id")
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &user.GetUnreadsNotificationCountResp{
		CommentCount: uint64(commentCount),
		FollowCount:  uint64(followCount),
		LikeCount:    uint64(likeCount),
		OsCount:      uint64(osCount),
		Total:        uint64(commentCount) + uint64(followCount) + uint64(likeCount) + uint64(osCount),
	}, nil
}
