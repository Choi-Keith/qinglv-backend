package post

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPostListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPostListLogic {
	return &GetUserPostListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPostListLogic) GetUserPostList(req *types.GetUserPostListReq) (resp *types.GetUserPostListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
