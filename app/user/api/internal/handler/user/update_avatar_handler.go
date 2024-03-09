package user

import (
	"net/http"

	"qinglv-backend/app/user/api/internal/logic/user"
	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/common/response"
)

func UpdateAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUpdateAvatarLogic(r.Context(), svcCtx, r)
		err := l.UpdateAvatar()
		if err != nil {
			response.FailCodeMsg(w, http.StatusBadRequest, err)
		} else {
			response.Ok(w)
		}
	}
}
