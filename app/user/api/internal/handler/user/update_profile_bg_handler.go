package user

import (
	"net/http"

	"qinglv-backend/app/user/api/internal/logic/user"
	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/common/response"
)

func UpdateProfileBgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUpdateProfileBgLogic(r.Context(), svcCtx, r)
		err := l.UpdateProfileBg()
		if err != nil {
			response.FailCodeMsg(w, http.StatusBadRequest, err)
		} else {
			response.Ok(w)
		}
	}
}
