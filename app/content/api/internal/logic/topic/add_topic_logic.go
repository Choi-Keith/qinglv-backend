package topic

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/app/content/rpc/content_client"
	"qinglv-backend/pkg/snowflake"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTopicLogic {
	return &AddTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTopicLogic) AddTopic(req *types.AddTopicReq) error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	id := snowflake.MustID()
	_, err = l.svcCtx.ContentRpc.AddTopic(l.ctx, &content_client.AddTopicReq{
		CreatorId:   uint64(userId),
		Name:        req.Name,
		Bg:          req.Bg,
		Description: req.Description,
		Id:          id,
	})
	if err != nil {
		return err
	}
	return nil
}
