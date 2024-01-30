package comment

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPostCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostCommentLogic {
	return &AddPostCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPostCommentLogic) AddPostComment(req *types.AddPostCommentReq) error {
	// todo: add your logic here and delete this line

	return nil
}
