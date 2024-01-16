package topic

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicListLogic {
	return &GetTopicListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicListLogic) GetTopicList(req *types.GetTopicListReq) (resp *types.GetTopicListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
