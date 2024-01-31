package topic

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicByIdLogic {
	return &GetTopicByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicByIdLogic) GetTopicById(req *types.GetTopicByIdReq) (resp *types.GetTopicByIdResp, err error) {
	// todo: add your logic here and delete this line
	topicResp, err := l.svcCtx.TopicRpc.GetTopicById(l.ctx, &content.GetTopicByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdReq{
		UserId: topicResp.Topic.CreatorId,
	})
	if err != nil {
		return nil, err
	}

	var topicItem types.TopicItem
	_ = copier.Copy(&topicItem, topicResp.Topic)
	_ = copier.Copy(&topicItem.Creator, userResp.User)

	return &types.GetTopicByIdResp{
		Topic: topicItem,
	}, nil
}
