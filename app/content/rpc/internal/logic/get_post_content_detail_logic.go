package logic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostContentDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostContentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostContentDetailLogic {
	return &GetPostContentDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: postContent
func (l *GetPostContentDetailLogic) GetPostContentDetail(in *content.GetPostContentDetailReq) (*content.GetPostContentDetailResp, error) {
	// todo: add your logic here and delete this line

	return &content.GetPostContentDetailResp{}, nil
}
