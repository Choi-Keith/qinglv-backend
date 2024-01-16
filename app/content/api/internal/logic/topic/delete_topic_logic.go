package topic

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTopicLogic {
	return &DeleteTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTopicLogic) DeleteTopic(req *types.DeleteTopicReq) error {
	// todo: add your logic here and delete this line

	return nil
}
