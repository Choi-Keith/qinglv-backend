package comment

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticleCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleCommentListLogic {
	return &GetArticleCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleCommentListLogic) GetArticleCommentList(req *types.GetArticleCommentListReq) (resp *types.GetArticleCommentListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
