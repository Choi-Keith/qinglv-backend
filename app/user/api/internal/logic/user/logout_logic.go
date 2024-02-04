package user

import (
	"context"
	"encoding/json"
	"fmt"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/common/globalKey"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() error {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return err
	}
	if _, err = l.svcCtx.UserRpc.Logout(l.ctx, &user.LogoutReq{
		TokenKey: fmt.Sprintf("%s%d", globalKey.TokenPrefixKey, userId),
	}); err != nil {
		return err
	}
	return nil
}
