package post

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostLogic {
	return &AddPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPostLogic) AddPost(req *types.AddPostReq) error {
	// todo: add your logic here and delete this line

	return nil
}
