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

type GetTopicByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicByNameLogic {
	return &GetTopicByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicByNameLogic) GetTopicByName(req *types.GetTopicByNameReq) (resp *types.GetTopicByNameResp, err error) {
	// todo: add your logic here and delete this line
	topicResp, err := l.svcCtx.TopicRpc.GetTopicByName(l.ctx, &content.GetTopicByNameReq{
		Name: req.Name,
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
	return &types.GetTopicByNameResp{
		Topic: topicItem,
	}, nil
}
