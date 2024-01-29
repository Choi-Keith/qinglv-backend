package logic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentThumbLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentThumbLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentThumbLogic {
	return &UpdateCommentThumbLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: PostCommentThumb
func (l *UpdateCommentThumbLogic) UpdateCommentThumb(in *operation.UpdateCommentThumbReq) (*operation.OkResp, error) {
	// todo: add your logic here and delete this line

	return &operation.OkResp{}, nil
}
