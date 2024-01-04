package user

import (
	"context"
	"errors"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user_client"
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
	userResp, err := l.svcCtx.UserRpc.GetUserInfoByParams(l.ctx, &user_client.GetUserInfoByParamsReq{
		Nickname: req.Nickname,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	if userResp.User != nil {
		return nil, errors.New("邮箱或昵称已存在，请重新输入")
	}

	id := snowflake.MustID()
	password, _ := password.EncryptPassword(req.Password)
	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user_client.RegisterReq{
		Id:       id,
		Email:    req.Email,
		Password: password,
		RoleId:   req.RoleId,
		Nickname: req.Nickname,
		AuthType: 1,
	})
	var userDetail types.User
	logx.Debugf("[Api User] Register registerResp: %+v\n", registerResp.User)
	if err != nil {
		logx.Errorf("[Api Logic] Register failed: %+v\n", err)
		return nil, err
	}
	if err := copier.Copy(&userDetail, registerResp.User); err != nil {
		logx.Errorf("[Api Logic] Register copy failed: %+v\n", err)
		return nil, err
	}

	roleItem, err := l.svcCtx.UserRpc.GetRoleInfo(l.ctx, &user_client.GetRoleInfoReq{
		Id: req.RoleId,
	})
	if err != nil {
		logx.Errorf("[Api Logic] Register roleItem failed: %+v\n", err)
		return nil, err
	}
	if err := copier.Copy(&userDetail.Role, roleItem.Role); err != nil {
		logx.Errorf("[Api Logic] Register copy failed: %+v\n", err)
		return nil, err
	}

	return &userDetail, nil
}
