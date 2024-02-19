package comment

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticleCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleCommentLogic {
	return &DeleteArticleCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleCommentLogic) DeleteArticleComment(req *types.DeleteArticleCommentReq) error {
	// todo: add your logic here and delete this line

	return nil
}
