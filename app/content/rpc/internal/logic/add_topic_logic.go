package logic

import (
	"context"
	"database/sql"
	"time"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/model/topic"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTopicLogic {
	return &AddTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: topic
func (l *AddTopicLogic) AddTopic(in *content.AddTopicReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line
	topicItem := &topic.Topic{
		Id:          in.Id,
		CreatorId:   in.CreatorId,
		Name:        in.Name,
		Bg:          in.Bg,
		Description: sql.NullString{String: in.Description, Valid: true},
		QuoteCount:  0,
		DeletedAt:   time.Now(),
		Version:     1,
	}
	_, err := l.svcCtx.TopicModel.Insert(l.ctx, nil, topicItem)
	if err != nil {
		return nil, err
	}
	return &content.OkResp{}, nil
}
