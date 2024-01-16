package logic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTopicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicDetailLogic {
	return &GetTopicDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: topic
func (l *GetTopicDetailLogic) GetTopicDetail(in *content.GetTopicDetailReq) (*content.GetTopicDetailResp, error) {
	// todo: add your logic here and delete this line

	return &content.GetTopicDetailResp{}, nil
}
