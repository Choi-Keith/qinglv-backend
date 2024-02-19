package comment

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleCommentReplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleCommentReplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleCommentReplyListLogic {
	return &GetArticleCommentReplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleCommentReplyListLogic) GetArticleCommentReplyList(req *types.GetArticleCommentReplyListReq) (resp *types.GetArticleCommentReplyListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
