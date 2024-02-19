package share

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddArticleShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleShareLogic {
	return &AddArticleShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddArticleShareLogic) AddArticleShare(req *types.AddArticleShareReq) error {
	// todo: add your logic here and delete this line

	return nil
}
