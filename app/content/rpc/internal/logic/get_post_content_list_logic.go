package logic

import (
	"context"

	"qinglv-backend/app/content/rpc/content"
	"qinglv-backend/app/content/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostContentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostContentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostContentListLogic {
	return &GetPostContentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: postContent
func (l *GetPostContentListLogic) GetPostContentList(in *content.GetPostContentListReq) (*content.GetPostContentListResp, error) {
	// todo: add your logic here and delete this line

	return &content.GetPostContentListResp{}, nil
}
