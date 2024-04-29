package userclasslogic

import (
	"context"
	"database/sql"
	"time"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group: user
func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserReq) (*user.UpdateUserResp, error) {
	// todo: add your logic here and delete this line
	userItem, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	if userItem.Email != in.Email {
		userItem.MailStatus = 1
	}
	nickname := userItem.Nickname
	if in.Nickname != "" {
		nickname = in.Nickname
	}
	userItem.Nickname = nickname
	userItem.Account = nickname
	userItem.Profession = sql.NullString{String: in.Profession, Valid: true}
	userItem.Phone = sql.NullString{String: in.Phone, Valid: true}
	userItem.Birthday, _ = time.ParseInLocation("2006-01-02", in.Birthday, time.Local)
	logx.Debugf("[UpdateUser] in.Birthday: %v, userItem.birthday: %v\n", in.Birthday, userItem.Birthday)
	userItem.Gender = int64(in.Gender)
	userItem.Motto = sql.NullString{String: in.Motto, Valid: true}
	userItem.Location = sql.NullString{String: in.Location, Valid: true}
	userItem.Address = sql.NullString{String: in.Address, Valid: true}
	err = l.svcCtx.UserModel.UpdateWithVersion(l.ctx, nil, userItem)
	if err != nil {
		return nil, err
	}
	if userItem.Email != in.Email {
		go sendAndSaveRegisterCode(l.svcCtx, in.UserId, in.Email)
	}
	return &user.UpdateUserResp{}, nil
}
