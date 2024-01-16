package logic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostContentLogic {
	return &DeletePostContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: postContent
func (l *DeletePostContentLogic) DeletePostContent(in *content.DeletePostContentReq) (*content.OkResp, error) {
	// todo: add your logic here and delete this line

	return &content.OkResp{}, nil
}
