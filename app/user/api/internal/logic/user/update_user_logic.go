package user

import (
	"context"
	"errors"

	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) error {
	// todo: add your logic here and delete this line
	userResp, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdReq{
		UserId: req.UserId,
	})
	if err != nil {
		return err
	}
	if req.Email != "" && userResp.User.Email != req.Email {
		checkEmailExistResp, err := l.svcCtx.UserRpc.CheckEmailExist(l.ctx, &user.CheckEmailExistReq{
			Email: req.Email,
		})
		if err != nil {
			return err
		}
		if checkEmailExistResp.IsExist {
			if checkEmailExistResp.User.MailStatus == 2 {
				return errors.New("邮箱已存在，请重新输入")
			}
			if checkEmailExistResp.User.MailStatus == 1 {
				_, err := l.svcCtx.UserRpc.DeleteUser(l.ctx, &user.DeleteUserReq{
					UserId: checkEmailExistResp.User.Id,
				})
				if err != nil {
					return err
				}
			}
		}
	}
	if req.Nickname != "" && userResp.User.Nickname != req.Nickname {
		checkNicknameExistResp, err := l.svcCtx.UserRpc.CheckNicknameExist(l.ctx, &user.CheckNicknameExistReq{
			Nickname: req.Nickname,
		})
		if err != nil {
			return err
		}
		if checkNicknameExistResp.IsExist {
			return errors.New("昵称已存在，请重新输入")
		}
	}
	_, err = l.svcCtx.UserRpc.UpdateUser(l.ctx, &user.UpdateUserReq{
		UserId:     req.UserId,
		Email:      req.Email,
		Phone:      req.Phone,
		Nickname:   req.Nickname,
		Age:        int32(req.Age),
		Gender:     int32(req.Gender),
		Profession: req.Profession,
		Motto:      req.Motto,
		Location:   req.Location,
		Address:    req.Address,
	})
	if err != nil {
		return err
	}
	return nil
}
