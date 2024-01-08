package user

import (
	"net/http"

	"qinglv-backend/app/user/api/internal/logic/user"
	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/common/response"
)

func LoginCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewLoginCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.LoginCaptcha()
		if err != nil {
			response.ParamsFail(w, err)
		} else {
			response.OkWithData(w, resp)
		}
	}
}
