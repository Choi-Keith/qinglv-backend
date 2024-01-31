package user

import (
	"context"
	"encoding/json"
	"errors"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type BanUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBanUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BanUserLogic {
	return &BanUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BanUserLogic) BanUser(req *types.BanUserReq) error {
	// todo: add your logic here and delete this line
	roleId, err := l.ctx.Value("roleId").(json.Number).Int64()
	if err != nil {
		return err
	}
	if roleId != 1 && roleId != 2 {
		return errors.New("没有权限")
	}
	_, err = l.svcCtx.UserRpc.BanUser(l.ctx, &user.BanUserReq{
		UserId: req.Id,
	})
	if err != nil {
		return err
	}

	return nil
}
