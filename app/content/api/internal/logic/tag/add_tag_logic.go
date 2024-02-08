package tag

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTagLogic {
	return &AddTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTagLogic) AddTag(req *types.AddTagReq) error {
	// todo: add your logic here and delete this line

	return nil
}
