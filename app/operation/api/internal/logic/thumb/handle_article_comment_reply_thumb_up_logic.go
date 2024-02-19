package thumb

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleArticleCommentReplyThumbUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandleArticleCommentReplyThumbUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleArticleCommentReplyThumbUpLogic {
	return &HandleArticleCommentReplyThumbUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleArticleCommentReplyThumbUpLogic) HandleArticleCommentReplyThumbUp(req *types.HandleArticleCommentReplyThumbUpReq) error {
	// todo: add your logic here and delete this line

	return nil
}
