package userclasslogic

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
	userItem, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	userItem.ProfileBg = in.ProfileBg
	err = l.svcCtx.UserModel.UpdateWithVersion(l.ctx, nil, userItem)
	if err != nil {
		return nil, err
	}
	return &user.UpdateProfileBgResp{}, nil
}
