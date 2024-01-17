package logic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	topicModel "qinglv-backend/app/content/rpc/internal/model/topic"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTopicByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicByIdLogic {
	return &GetTopicByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: topic
func (l *GetTopicByIdLogic) GetTopicById(in *content.GetTopicByIdReq) (*content.GetTopicByIdResp, error) {
	// todo: add your logic here and delete this line
	topicResp, err := l.svcCtx.TopicModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	topicItem := genTopicItem(topicResp)
	return &content.GetTopicByIdResp{
		Topic: topicItem,
	}, nil
}

func genTopicItem(topicItem *topicModel.Topic) *content.TopicItem {
	return &content.TopicItem{
		Id:          topicItem.Id,
		CreatorId:   topicItem.CreatorId,
		Name:        topicItem.Name,
		Description: topicItem.Description.String,
		QuoteCount:  topicItem.QuoteCount,
		CreatedAt:   uint64(topicItem.CreatedAt.Unix() * 1000),
		UpdatedAt:   uint64(topicItem.UpdatedAt.Unix() * 1000),
	}
}
