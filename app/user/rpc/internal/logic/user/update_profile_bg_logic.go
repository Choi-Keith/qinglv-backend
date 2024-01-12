package user

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProfileBgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProfileBgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProfileBgLogic {
	return &UpdateProfileBgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: user
func (l *UpdateProfileBgLogic) UpdateProfileBg(in *user.UpdateProfileBgReq) (*user.UpdateProfileBgResp, error) {
	// todo: add your logic here and delete this line

	return &user.UpdateProfileBgResp{}, nil
}
