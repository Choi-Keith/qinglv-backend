package user

import (
	"context"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

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
	// registerUser := &userModel.User{
	// 	Id:       in.Id,
	// 	RoleId:   in.RoleId,
	// 	Nickname: in.Nickname,
	// 	Email:    in.Email,
	// 	Phone:    sql.NullString{String: in.Phone, Valid: false},
	// 	AuthType: int64(in.AuthType),
	// }
	// resp, err := l.svcCtx.UserModel.Insert(l.ctx, nil, registerUser)
	// if err != nil {
	// 	logx.Errorf("register user failed: %v", err)
	// 	return nil, err
	// }
	// id, _ := resp.LastInsertId()
	// userItem, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(id))
	// if err != nil {
	// 	logx.Errorf("register user failed: %v", err)
	// 	return nil, err
	// }
	// item := &user.UserItem{
	// 	Id:        userItem.Id,
	// 	RoleId:    userItem.RoleId,
	// 	Nickname:  userItem.Nickname,
	// 	Email:     userItem.Email,
	// 	Phone:     userItem.Phone.String,
	// 	WeChat:    userItem.WeChat.String,
	// 	Motto:     userItem.Motto.String,
	// 	Avatar:    userItem.Avatar.String,
	// 	ProfileBg: userItem.ProfileBg.String,
	// 	Age:       int32(userItem.Age.Int64),
	// 	Gender:    int32(userItem.Gender.Int64),
	// 	Location:  userItem.Location.String,
	// 	Level:     int32(userItem.Level),
	// 	Score:     int32(userItem.Score),
	// 	CreatedAt: uint64(userItem.CreatedAt.Unix()),
	// 	UpdatedAt: uint64(userItem.UpdatedAt.Unix()),
	// 	AuthType:  int32(userItem.AuthType),
	// }
	// return &user.RegisterResp{
	// 	User: item,
	// }, nil
	return &user.RegisterResp{}, nil
}
