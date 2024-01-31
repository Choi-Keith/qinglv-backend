package userclasslogic

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAvatarLogic {
	return &UpdateAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: user
func (l *UpdateAvatarLogic) UpdateAvatar(in *user.UpdateAvatarReq) (*user.UpdateAvatarResp, error) {
	// todo: add your logic here and delete this line

	userItem, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	userItem.Avatar = in.Avatar
	err = l.svcCtx.UserModel.UpdateWithVersion(l.ctx, nil, userItem)
	if err != nil {
		return nil, err
	}
	return &user.UpdateAvatarResp{}, nil
}
