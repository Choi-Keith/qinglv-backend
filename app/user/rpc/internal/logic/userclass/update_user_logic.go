package userclasslogic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/app/user/rpc/user"
	"qinglv-backend/common/globalKey"
	"qinglv-backend/common/schema"
	"qinglv-backend/pkg/email"
	"qinglv-backend/pkg/template"

	"github.com/google/uuid"
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
	userItem.Nickname = in.Nickname
	userItem.Account = in.Nickname
	userItem.Phone = sql.NullString{String: in.Phone, Valid: true}
	userItem.Age = sql.NullInt64{Int64: int64(in.Age), Valid: true}
	userItem.Gender = int64(in.Gender)
	userItem.Motto = sql.NullString{String: in.Motto, Valid: true}
	userItem.Location = sql.NullString{String: in.Location, Valid: true}
	userItem.Address = sql.NullString{String: in.Address, Valid: true}
	err = l.svcCtx.UserModel.UpdateWithVersion(l.ctx, nil, userItem)
	if err != nil {
		return nil, err
	}
	if userItem.Email != in.Email {
		go l.SendAndSaveRegisterCode(in.UserId, in.Email)
	}
	return &user.UpdateUserResp{}, nil
}

func (l *UpdateUserLogic) SendAndSaveRegisterCode(userId uint64, toUser string) {
	host := l.svcCtx.Config.Website.Host
	port := l.svcCtx.Config.Website.Port
	smtp := l.svcCtx.Config.SMTP
	code := uuid.New()
	verifyEmailURL := fmt.Sprintf("http://%s:%d/email/verify?code=%s", host, port, code)
	body, err := template.GenerateVerifyBody(verifyEmailURL, "verify_email.html")
	if err != nil {
		logx.Errorf("[User SendAndSaveRegisterCode] GenerateVerifyBody failed: %+v\n", err)
	}
	if err := email.Send(smtp, toUser, "更改邮件地址-请确认邮件地址", body); err != nil {
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
