package post

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostByIdLogic {
	return &GetPostByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostByIdLogic) GetPostById(req *types.GetPostByIdReq) (resp *types.GetPostByIdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
