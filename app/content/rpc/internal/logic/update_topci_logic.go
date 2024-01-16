package logic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopciLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTopciLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopciLogic {
	return &UpdateTopciLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: topic
func (l *UpdateTopciLogic) UpdateTopci(in *content.UpdateTopicReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line

	return &content.OkResp{}, nil
}
