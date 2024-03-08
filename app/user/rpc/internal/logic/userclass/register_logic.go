package userclasslogic

import (
	"context"
	"time"

	userModel "qinglv-backend/app/user/rpc/internal/model/user"
	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/pkg/gavatar"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: user
func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line

	avatar := gavatar.GenRandomAvatar("https://qinglv-1304086226.cos.ap-guangzhou.myqcloud.com")
	registerUser := new(userModel.User)
	registerUser.Id = in.Id
	registerUser.RoleId = in.RoleId
	registerUser.Nickname = in.Nickname
	registerUser.Email = in.Email
	registerUser.Password = in.Password
	registerUser.Account = in.Nickname
	registerUser.Avatar = avatar
	registerUser.AuthType = int64(in.AuthType)
	registerUser.DeletedAt = time.Now()
	registerUser.LastLoginTime = time.Now()
	registerUser.Version = 1
	registerUser.Status = 2
	registerUser.Level = 1
	registerUser.MailStatus = 1
	registerUser.ProfileBg = "https://qinglv-1304086226.cos.ap-guangzhou.myqcloud.com/images/profileBg/default/bg.png"
	_, err := l.svcCtx.UserModel.Insert(l.ctx, nil, registerUser)
	if err != nil {
		logx.Errorf("register user failed: %v", err)
		return nil, err
	}
	go sendAndSaveRegisterCode(l.svcCtx, in.Id, in.Email)
	userItem, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		logx.Errorf("register FindOne user failed: %v", err)
		return nil, err
	}
	item := &user.UserItem{
		Id:            userItem.Id,
		RoleId:        userItem.RoleId,
		Nickname:      userItem.Nickname,
		Email:         userItem.Email,
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
	return &user.RegisterResp{
		User: item,
	}, nil
}
