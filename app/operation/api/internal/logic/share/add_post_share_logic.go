package share

import (
	"context"

	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPostShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostShareLogic {
	return &AddPostShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPostShareLogic) AddPostShare(req *types.AddPostShareReq) error {
	// todo: add your logic here and delete this line

	return nil
}
