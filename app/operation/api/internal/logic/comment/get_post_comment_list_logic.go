package comment

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostCommentListLogic {
	return &GetPostCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostCommentListLogic) GetPostCommentList(req *types.GetPostCommentListReq) error {
	// todo: add your logic here and delete this line

	return nil
}
