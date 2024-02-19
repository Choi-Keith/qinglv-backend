package thumb

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleArticleCommentReplyThumbDownLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandleArticleCommentReplyThumbDownLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleArticleCommentReplyThumbDownLogic {
	return &HandleArticleCommentReplyThumbDownLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleArticleCommentReplyThumbDownLogic) HandleArticleCommentReplyThumbDown(req *types.HandleArticleCommentReplyThumbDownReq) error {
	// todo: add your logic here and delete this line

	return nil
}
