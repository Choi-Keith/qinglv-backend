package article

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserArticleListLogic {
	return &GetUserArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserArticleListLogic) GetUserArticleList(req *types.GetUserArticleListReq) (resp *types.GetUserArticleListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
