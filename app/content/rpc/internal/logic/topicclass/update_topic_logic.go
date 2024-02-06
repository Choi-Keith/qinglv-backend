package topicclasslogic

import (
	"context"
	"database/sql"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicLogic {
	return &UpdateTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: topic
func (l *UpdateTopicLogic) UpdateTopic(in *content.UpdateTopicReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	topicItem, err := l.svcCtx.TopicModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	bg := topicItem.Bg
	if in.Bg != "" {
		bg = in.Bg
	}
	topicItem.Bg = bg
	desc := topicItem.Description.String
	if in.Description != "" {
		desc = in.Description
	}
	topicItem.Description = sql.NullString{String: desc, Valid: true}
	quoteCount := topicItem.QuoteCount
	if in.QuoteCount != 0 {
		quoteCount = in.QuoteCount
	}
	topicItem.QuoteCount = quoteCount
	score := topicItem.Score
	if in.Score != 0 {
		score = in.Score
	}
	topicItem.Score = score
	err = l.svcCtx.TopicModel.UpdateWithVersion(l.ctx, nil, topicItem)
	if err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
