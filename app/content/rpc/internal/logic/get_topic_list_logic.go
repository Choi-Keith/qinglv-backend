package logic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTopicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicListLogic {
	return &GetTopicListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: topic
func (l *GetTopicListLogic) GetTopicList(in *content.GetTopicListReq) (*content.GetTopicListResp, error) {
	// todo: add your logic here and delete this line
	whereBuilder := l.svcCtx.TopicModel.SelectBuilder()
	if in.CreatorId != 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"creator_id": in.CreatorId,
		})
	}
	if in.Name != "" {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"name": in.Name,
		})
	}
	topicResp, total, err := l.svcCtx.TopicModel.FindPageListByPageWithTotal(l.ctx, whereBuilder, int64(in.PageNum), int64(in.PageSize), "")
	if err != nil {
		return nil, err
	}
	topicList := make([]*content.TopicItem, len(topicResp))
	for idx, topicItem := range topicResp {
		topicList[idx] = genTopicItem(topicItem)
	}
	return &content.GetTopicListResp{
		Data:  topicList,
		Total: uint64(total),
	}, nil
}
