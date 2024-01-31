package topicclasslogic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTopicByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicByNameLogic {
	return &GetTopicByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: topic
func (l *GetTopicByNameLogic) GetTopicByName(in *content.GetTopicByNameReq) (*content.GetTopicByNameResp, error) {
	// todo: add your logic here and delete this line
	topicResp, err := l.svcCtx.TopicModel.FindOneByName(l.ctx, in.Name)
	if err != nil {
		return nil, err
	}
	topicItem := genTopicItem(topicResp)
	return &content.GetTopicByNameResp{
		Topic: topicItem,
	}, nil
}
