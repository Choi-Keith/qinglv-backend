package comment

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleCommentReplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticleCommentReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleCommentReplyLogic {
	return &DeleteArticleCommentReplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleCommentReplyLogic) DeleteArticleCommentReply(req *types.DeleteArticleCommentReplyReq) error {
	// todo: add your logic here and delete this line

	return nil
}
