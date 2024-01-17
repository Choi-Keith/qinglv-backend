package topic

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicLogic {
	return &UpdateTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTopicLogic) UpdateTopic(req *types.UpdateTopicReq) error {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.ContentRpc.UpdateTopic(l.ctx, &content_client.UpdateTopicReq{
		Id:          req.Id,
		Bg:          req.Bg,
		Description: req.Description,
	})

	if err != nil {
		return err
	}
	return nil
}
