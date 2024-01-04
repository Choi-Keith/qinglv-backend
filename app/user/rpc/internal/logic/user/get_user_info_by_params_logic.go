package user

import (
	"context"

	userModel "qinglv-backend/app/user/rpc/internal/model/user"
	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByParamsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByParamsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByParamsLogic {
	return &GetUserInfoByParamsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: user
func (l *GetUserInfoByParamsLogic) GetUserInfoByParams(in *user.GetUserInfoByParamsReq) (*user.GetUserInfoByParamsResp, error) {
	// todo: add your logic here and delete this line
	GetUserInfoByParamsReq := &userModel.User{
		Id:       in.Id,
		Nickname: in.Nickname,
		Email:    in.Email,
	}
	userItem, err := l.svcCtx.UserModel.FindOneByParams(l.ctx, *GetUserInfoByParamsReq)
	if err != nil {
		logx.Errorf("[Rpc Logic] GetUserInfoByParams failed: %v\n", err)
		return nil, err
	}
	if userItem == nil {
		return &user.GetUserInfoByParamsResp{
			User: nil,
		}, nil
	}
	logx.Debugf("[Rpc] GetUserInfoByParams item: %+v\n", userItem)

	item := &user.UserItem{
		Id:        userItem.Id,
		RoleId:    userItem.RoleId,
		Nickname:  userItem.Nickname,
		Email:     userItem.Email,
		Phone:     userItem.Phone.String,
		WeChat:    userItem.WeChat.String,
		Motto:     userItem.Motto.String,
		Avatar:    userItem.Avatar,
		ProfileBg: userItem.ProfileBg.String,
		Age:       int32(userItem.Age.Int64),
		Gender:    int32(userItem.Gender.Int64),
		Location:  userItem.Location.String,
		Level:     int32(userItem.Level),
		Score:     int32(userItem.Score),
		CreatedAt: uint64(userItem.CreatedAt.Unix() * 1000),
		UpdatedAt: uint64(userItem.UpdatedAt.Unix() * 1000),
		AuthType:  int32(userItem.AuthType),
	}
	return &user.GetUserInfoByParamsResp{
		User: item,
	}, nil
}
