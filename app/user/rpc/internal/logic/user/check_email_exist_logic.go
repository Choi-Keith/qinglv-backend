package user

import (
	"context"
	"errors"

	userModel "qinglv-backend/app/user/rpc/internal/model/user"
	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type CheckEmailExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckEmailExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckEmailExistLogic {
	return &CheckEmailExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: user
func (l *CheckEmailExistLogic) CheckEmailExist(in *user.CheckEmailExistReq) (*user.CheckEmailExistResp, error) {
	// todo: add your logic here and delete this line

	userResp, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return &user.CheckEmailExistResp{
				IsExist: false,
				User:    nil,
			}, nil
		}
		return nil, err
	}
	item := genUserItem(userResp)
	return &user.CheckEmailExistResp{
		IsExist: true,
		User:    item,
	}, nil
}

func genUserItem(userItem *userModel.User) *user.UserItem {
	return &user.UserItem{
		Id:            userItem.Id,
		RoleId:        userItem.RoleId,
		Nickname:      userItem.Nickname,
		Email:         userItem.Email,
		Password:      userItem.Password,
		Status:        int32(userItem.Status),
		MailStatus:    int32(userItem.MailStatus),
		Phone:         userItem.Phone.String,
		WeChat:        userItem.WeChat.String,
		Motto:         userItem.Motto.String,
		Avatar:        userItem.Avatar,
		ProfileBg:     userItem.ProfileBg,
		Age:           int32(userItem.Age.Int64),
		Gender:        int32(userItem.Gender),
		Location:      userItem.Location.String,
		Level:         int32(userItem.Level),
		Score:         int32(userItem.Score),
		CreatedAt:     uint64(userItem.CreatedAt.Unix() * 1000),
		UpdatedAt:     uint64(userItem.UpdatedAt.Unix() * 1000),
		LastLoginTime: uint64(userItem.LastLoginTime.Unix() * 1000),
		AuthType:      int32(userItem.AuthType),
	}
}
