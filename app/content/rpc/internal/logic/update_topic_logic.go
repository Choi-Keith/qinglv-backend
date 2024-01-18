package logic

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
	topicItem.Bg = in.Bg
	topicItem.Description = sql.NullString{String: in.Description, Valid: true}
	topicItem.QuoteCount = in.QuoteCount
	topicItem.Type = int64(in.Type)
	err = l.svcCtx.TopicModel.UpdateWithVersion(l.ctx, nil, topicItem)
	if err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
