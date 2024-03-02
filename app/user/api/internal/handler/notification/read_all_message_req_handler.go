package notification

import (
	"net/http"

	"qinglv-backend/app/user/api/internal/logic/notification"
	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReadAllMessageReqHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReadAllMessageReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamsFail(w, err)
			return
		}

		l := notification.NewReadAllMessageReqLogic(r.Context(), svcCtx)
		err := l.ReadAllMessageReq(&req)
		if err != nil {
			response.FailCodeMsg(w, http.StatusBadRequest, err)
		} else {
			httpx.Ok(w)
		}
	}
}
