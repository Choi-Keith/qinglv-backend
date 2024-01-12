package user

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	userModel "qinglv-backend/app/user/rpc/internal/model/user"
	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/common/globalKey"
	"qinglv-backend/common/schema"
	"qinglv-backend/pkg/email"
	"qinglv-backend/pkg/gavatar"
	"qinglv-backend/pkg/template"

	"github.com/google/uuid"
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
	registerUser.MailStatus = 1
	registerUser.ProfileBg = "https://qinglv-1304086226.cos.ap-guangzhou.myqcloud.com/images/profileBg/default/bg.png"
	_, err := l.svcCtx.UserModel.Insert(l.ctx, nil, registerUser)
	if err != nil {
		logx.Errorf("register user failed: %v", err)
		return nil, err
	}
	go l.SendAndSaveRegisterCode(in.Id, in.Email)
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

func (l *RegisterLogic) SendAndSaveRegisterCode(userId uint64, toUser string) {
	host := l.svcCtx.Config.Website.Host
	port := l.svcCtx.Config.Website.Port
	smtp := l.svcCtx.Config.SMTP
	code := uuid.New()
	verifyEmailURL := fmt.Sprintf("http://%s:%d/email/verify?code=%s", host, port, code)
	body, err := template.GenerateVerifyBody(verifyEmailURL)
	if err != nil {
		logx.Errorf("[User SendAndSaveRegisterCode] GenerateVerifyBody failed: %+v\n", err)
	}
	if err := email.Send(smtp, toUser, "欢迎注册轻旅社区-请确认邮件地址", body); err != nil {
		logx.Errorf("[User SendAndSaveRegisterCode] email Send failed: %+v\n", err)
	}

	key := fmt.Sprintf("%s%s", globalKey.VerifyEmailCodePrefixKey, code)
	codeContent := &schema.EmailContent{
		UserId: userId,
		Email:  toUser,
	}
	codeContentStr, _ := json.Marshal(codeContent)
	expireAt := 10 * 60
	if err := l.svcCtx.RedisClient.Setex(key, string(codeContentStr), int(expireAt)); err != nil {
		logx.Errorf("[User SendAndSaveRegisterCode] Setex failed: %+v\n", err)
	}

}
