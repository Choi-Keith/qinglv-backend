package topic

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"

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

	return
}
