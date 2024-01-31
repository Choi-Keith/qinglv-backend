package userclasslogic

import (
	"context"
	"errors"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type CheckNicknameExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckNicknameExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckNicknameExistLogic {
	return &CheckNicknameExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: user
func (l *CheckNicknameExistLogic) CheckNicknameExist(in *user.CheckNicknameExistReq) (*user.CheckNicknameExistResp, error) {
	// todo: add your logic here and delete this line
	userResp, err := l.svcCtx.UserModel.FindOneByNickname(l.ctx, in.Nickname)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return &user.CheckNicknameExistResp{
				IsExist: false,
				User:    nil,
			}, nil
		}
		return nil, err
	}
	item := genUserItem(userResp)
	return &user.CheckNicknameExistResp{
		IsExist: true,
		User:    item,
	}, nil
}
