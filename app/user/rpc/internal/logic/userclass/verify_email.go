package userclasslogic

import (
	"encoding/json"
	"fmt"
	"qinglv-backend/app/user/rpc/internal/svc"
	"qinglv-backend/common/globalKey"
	"qinglv-backend/common/schema"
	"qinglv-backend/pkg/email"
	"qinglv-backend/pkg/template"

	"github.com/google/uuid"
)

func sendAndSaveForgotCode(svc *svc.ServiceContext, userId uint64, toUser string) error {
	host := svc.Config.Website.Host
	port := svc.Config.Website.Port
	smtp := svc.Config.SMTP
	code := uuid.New()
	verifyEmailURL := fmt.Sprintf("http://%s:%d/reset-password?code=%s", host, port, code)
	body, err := template.GenerateVerifyBody(verifyEmailURL, "forgot_password.html")
	if err != nil {
		return err
	}
	if err := email.Send(smtp, toUser, "轻旅社区-重置密码", body); err != nil {
		return err
	}

	key := fmt.Sprintf("%s%s", globalKey.VerifyForgotPasswordEmailCodePrefixKey, code)
	codeContent := &schema.EmailContent{
		UserId: userId,
		Email:  toUser,
	}
	codeContentStr, _ := json.Marshal(codeContent)
	expireAt := 10 * 60
	if err := svc.RedisClient.Setex(key, string(codeContentStr), int(expireAt)); err != nil {
		return err
	}
	return nil

}
