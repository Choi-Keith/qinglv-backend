package thumb

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleArticleCommentThumbUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandleArticleCommentThumbUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleArticleCommentThumbUpLogic {
	return &HandleArticleCommentThumbUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleArticleCommentThumbUpLogic) HandleArticleCommentThumbUp(req *types.HandleArticleCommentThumbUpReq) error {
	// todo: add your logic here and delete this line

	return nil
}
