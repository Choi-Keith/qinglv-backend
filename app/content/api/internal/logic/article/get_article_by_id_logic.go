package article

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleByIdLogic {
	return &GetArticleByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleByIdLogic) GetArticleById(req *types.GetArticleByIdReq) (resp *types.GetArticleByIdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
