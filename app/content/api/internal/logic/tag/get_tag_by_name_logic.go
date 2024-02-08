package tag

import (
	"context"

	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTagByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagByNameLogic {
	return &GetTagByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTagByNameLogic) GetTagByName(req *types.GetTagByNameReq) (resp *types.GetTagByNameResp, err error) {
	// todo: add your logic here and delete this line

	return
}
