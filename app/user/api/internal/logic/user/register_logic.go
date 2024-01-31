package user

import (
	"context"
	"errors"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/pkg/password"
	"qinglv-backend/pkg/snowflake"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.User, err error) {
	// todo: add your logic here and delete this line
	checkEmailExistResp, err := l.svcCtx.UserRpc.CheckEmailExist(l.ctx, &user.CheckEmailExistReq{
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}
	if checkEmailExistResp.IsExist {
		if checkEmailExistResp.User.MailStatus == 2 {
			return nil, errors.New("邮箱已存在，请重新输入")
		}
		if checkEmailExistResp.User.MailStatus == 1 {
			_, err := l.svcCtx.UserRpc.DeleteUser(l.ctx, &user.DeleteUserReq{
				UserId: checkEmailExistResp.User.Id,
			})
			if err != nil {
				return nil, err
			}
		}
	}
	checkNicknameExistResp, err := l.svcCtx.UserRpc.CheckNicknameExist(l.ctx, &user.CheckNicknameExistReq{
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	if checkNicknameExistResp.IsExist {
		return nil, errors.New("昵称已存在，请重新输入")
	}
	id := snowflake.MustID()
	password, _ := password.EncryptPassword(req.Password)
	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
		Id:       id,
		Email:    req.Email,
		Password: password,
		RoleId:   req.RoleId,
		Nickname: req.Nickname,
		AuthType: 1,
	})
	var userDetail types.User
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&userDetail, registerResp.User)

	roleItem, err := l.svcCtx.RoleRpc.GetRoleInfo(l.ctx, &user.GetRoleInfoReq{
		Id: req.RoleId,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&userDetail.Role, roleItem.Role)
	return &userDetail, nil
}
