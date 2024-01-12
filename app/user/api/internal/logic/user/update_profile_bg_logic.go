package user

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProfileBgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProfileBgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProfileBgLogic {
	return &UpdateProfileBgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProfileBgLogic) UpdateProfileBg(req *types.UpdateProfileReq) error {
	// todo: add your logic here and delete this line

	return nil
}
