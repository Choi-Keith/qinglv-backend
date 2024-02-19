package thumb

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HandleArticleCommentThumbDownLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHandleArticleCommentThumbDownLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HandleArticleCommentThumbDownLogic {
	return &HandleArticleCommentThumbDownLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HandleArticleCommentThumbDownLogic) HandleArticleCommentThumbDown(req *types.HandleArticleCommentThumbDownReq) error {
	// todo: add your logic here and delete this line

	return nil
}
