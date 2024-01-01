package logic

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByNicknameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByNicknameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByNicknameLogic {
	return &GetUserInfoByNicknameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByNicknameLogic) GetUserInfoByNickname(in *user.GetUserInfoNicknameReq) (*user.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserInfoResp{}, nil
}
