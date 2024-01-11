package user

import (
	"net/http"

	"qinglv-backend/app/user/api/internal/logic/user"
	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DelUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelUserReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamsFail(w, err)
			return
		}

		l := user.NewDelUserLogic(r.Context(), svcCtx)
		err := l.DelUser(&req)
		if err != nil {
			response.FailCodeMsg(w, http.StatusAccepted, err)
		} else {
			response.Ok(w)
		}
	}
}
