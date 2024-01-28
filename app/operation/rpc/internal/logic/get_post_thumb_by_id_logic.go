package logic

import (
	"context"

	"qinglv-backend/app/operation/rpc/internal/svc"
	"qinglv-backend/app/operation/rpc/operation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostThumbByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostThumbByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostThumbByIdLogic {
	return &GetPostThumbByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: PostThumb
func (l *GetPostThumbByIdLogic) GetPostThumbById(in *operation.GetPostThumbByIdReq) (*operation.GetPostThumbByIdResp, error) {
	// todo: add your logic here and delete this line

	return &operation.GetPostThumbByIdResp{}, nil
}
