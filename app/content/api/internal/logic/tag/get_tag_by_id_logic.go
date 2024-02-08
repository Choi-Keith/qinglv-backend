package tag

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTagByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagByIdLogic {
	return &GetTagByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTagByIdLogic) GetTagById(req *types.GetTagByIdReq) (resp *types.GetTagByIdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
