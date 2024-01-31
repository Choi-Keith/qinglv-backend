package userclasslogic

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type BanUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBanUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BanUserLogic {
	return &BanUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: user
func (l *BanUserLogic) BanUser(in *user.BanUserReq) (*user.BanUserResp, error) {
	// todo: add your logic here and delete this line

	userItem, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	userItem.Status = 0
	if err := l.svcCtx.UserModel.UpdateWithVersion(l.ctx, nil, userItem); err != nil {
		return nil, err
	}
	return &user.BanUserResp{}, nil
}
