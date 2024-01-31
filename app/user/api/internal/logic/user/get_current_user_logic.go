package user

import (
	"context"
	"encoding/json"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCurrentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentUserLogic {
	return &GetCurrentUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCurrentUserLogic) GetCurrentUser() (resp *types.User, err error) {
	// todo: add your logic here and delete this line
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdReq{
		UserId: uint64(userId),
	})
	if err != nil {
		return nil, err
	}
	var userDetail types.User
	_ = copier.Copy(&userDetail, userResp.User)

	roleItem, err := l.svcCtx.RoleRpc.GetRoleInfo(l.ctx, &user.GetRoleInfoReq{
		Id: userResp.User.RoleId,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&userDetail.Role, roleItem.Role)
	return &userDetail, nil
}
