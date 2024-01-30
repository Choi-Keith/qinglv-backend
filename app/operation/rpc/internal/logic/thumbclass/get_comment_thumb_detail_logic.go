package thumbclasslogic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentThumbDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentThumbDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentThumbDetailLogic {
	return &GetCommentThumbDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: PostCommentThumb
func (l *GetCommentThumbDetailLogic) GetCommentThumbDetail(in *operation.GetCommentThumbDetailReq) (*operation.GetCommentThumbDetailReq, error) {
	// todo: add your logic here and delete this line

	return &operation.GetCommentThumbDetailReq{}, nil
}
