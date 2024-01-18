package topic

import (
	"context"
	"strings"

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
	var oldImage string
	if req.Bg != "" {
		topicResp, err := l.svcCtx.ContentRpc.GetTopicById(l.ctx, &content_client.GetTopicByIdReq{
			Id: req.Id,
		})
		if err != nil {
			return err
		}
		oldImage = topicResp.Topic.Bg
	}
	_, err := l.svcCtx.ContentRpc.UpdateTopic(l.ctx, &content_client.UpdateTopicReq{
		Id:          req.Id,
		Bg:          req.Bg,
		Description: req.Description,
		Type:        int32(req.Type),
		QuoteCount:  req.QuoteCount,
	})

	if err != nil {
		return err
	}
	if req.Bg != oldImage && oldImage != "" {
		name, _ := strings.CutPrefix(oldImage, l.svcCtx.Config.Cos.Endpoint)
		_, err := l.svcCtx.CosClient.Object.Delete(context.Background(), name)
		if err != nil {
			return err
		}
	}
	return nil
}
