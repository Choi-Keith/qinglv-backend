package logic

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	userItem, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		logx.Errorf("Get UserInfo failed:%v\n", err)
		return nil, err
	}
	item := &user.UserItem{
		Id:        userItem.Id,
		RoleId:    userItem.RoleId,
		Nickname:  userItem.Nickname,
		Email:     userItem.Email,
		Phone:     userItem.Phone.String,
		WeChat:    userItem.WeChat.String,
		Motto:     userItem.Motto.String,
		Avatar:    userItem.Avatar.String,
		ProfileBg: userItem.ProfileBg.String,
		Age:       int32(userItem.Age.Int64),
		Gender:    int32(userItem.Gender.Int64),
		Location:  userItem.Location.String,
		Level:     int32(userItem.Level),
		Score:     int32(userItem.Score),
		CreatedAt: uint64(userItem.CreatedAt.Unix()),
		UpdatedAt: uint64(userItem.UpdatedAt.Unix()),
		AuthType:  int32(userItem.AuthType),
	}
	return &user.GetUserInfoResp{
		User: item,
	}, nil
}
