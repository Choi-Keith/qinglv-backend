package comment

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostCommentReplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostCommentReplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostCommentReplyListLogic {
	return &GetPostCommentReplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostCommentReplyListLogic) GetPostCommentReplyList(req *types.GetPostCommentReplyListReq) error {
	// todo: add your logic here and delete this line

	return nil
}
