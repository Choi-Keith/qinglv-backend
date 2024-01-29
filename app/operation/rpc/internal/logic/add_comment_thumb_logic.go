package logic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentThumbLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentThumbLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentThumbLogic {
	return &AddCommentThumbLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: CommentThumb
func (l *AddCommentThumbLogic) AddCommentThumb(in *operation.AddCommentThumbReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line

	return &operation.OkResp{}, nil
}
