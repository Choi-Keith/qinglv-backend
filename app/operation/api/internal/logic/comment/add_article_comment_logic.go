package comment

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddArticleCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleCommentLogic {
	return &AddArticleCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddArticleCommentLogic) AddArticleComment(req *types.AddArticleCommentReq) error {
	// todo: add your logic here and delete this line

	return nil
}
