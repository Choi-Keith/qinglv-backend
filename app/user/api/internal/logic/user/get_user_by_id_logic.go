package user

import (
	"context"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user_client"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.GetUserByIdReq) (resp *types.User, err error) {
	// todo: add your logic here and delete this line
	userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user_client.GetUserByIdReq{
		UserId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	var userDetail types.User
	_ = copier.Copy(&userDetail, userResp.User)

	roleItem, err := l.svcCtx.UserRpc.GetRoleInfo(l.ctx, &user_client.GetRoleInfoReq{
		Id: userResp.User.RoleId,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&userDetail.Role, roleItem.Role)
	return &userDetail, nil
}
