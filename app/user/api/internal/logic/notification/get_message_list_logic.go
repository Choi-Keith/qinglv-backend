package notification

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GetMessageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageListLogic {
	return &GetMessageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessageListLogic) GetMessageList(req *types.GetMessageListReq) (resp *types.GetMessageListResp, err error) {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	notificationListResp, err := l.svcCtx.NotificationRpc.GetNotificationList(l.ctx, &user.GetNotificationListReq{
		Type:           req.Type,
		ReceiverUserId: uint64(userId),
		PageNum:        int64(req.PageNum),
		PageSize:       int64(req.PageSize),
	})
	notificationList, err := mr.MapReduce(func(source chan<- user.NotificationItem) {
		for _, notificationItem := range notificationListResp.Data {
			source <- *notificationItem
		}
	}, func(item user.NotificationItem, writer mr.Writer[types.MessageItem], cancel func(error)) {
		userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdReq{
			UserId: item.SenderUserId,
		})
		if err != nil {
			cancel(err)
			return
		}
		var notificationItem types.MessageItem
		_ = copier.Copy(&notificationItem, &item)
		_ = copier.Copy(&notificationItem.SenderUser, &userResp.User)
	}, func(pipe <-chan types.MessageItem, writer mr.Writer[[]types.MessageItem], cancel func(error)) {
		var r []types.MessageItem
		m := make(map[uint64]types.MessageItem, len(notificationListResp.Data))
		for p := range pipe {
			m[p.Id] = p
		}
		for _, notificationItem := range notificationListResp.Data {
			r = append(r, m[notificationItem.Id])
		}
		writer.Write(r)

	})
	if err != nil {
		return nil, err
	}
	isEnd := false
	total := (req.PageNum-1)*req.PageSize + req.PageSize
	if notificationListResp.Total <= uint64(total) {
		isEnd = true
	}
	return &types.GetMessageListResp{
		List:  notificationList,
		Total: notificationListResp.Total,
		IsEnd: isEnd,
	}, nil
}
