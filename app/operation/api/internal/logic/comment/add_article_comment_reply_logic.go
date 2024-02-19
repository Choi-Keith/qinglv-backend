package comment

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddArticleCommentReplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddArticleCommentReplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleCommentReplyLogic {
	return &AddArticleCommentReplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddArticleCommentReplyLogic) AddArticleCommentReply(req *types.AddArticleCommentReplyReq) error {
	// todo: add your logic here and delete this line

	return nil
}
