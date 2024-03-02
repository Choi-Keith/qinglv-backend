package notification

import (
	"net/http"

	"qinglv-backend/app/user/api/internal/logic/notification"
	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
	"qinglv-backend/common/response"
	"qinglv-backend/pkg/validate"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMessageListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMessageListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamsFail(w, err)
			return
		}

		if err := validate.Validate(req); err != nil {
			response.ParamsFail(w, err)
			return
		}

		l := notification.NewGetMessageListLogic(r.Context(), svcCtx)
		resp, err := l.GetMessageList(&req)
		if err != nil {
			response.FailCodeMsg(w, http.StatusBadRequest, err)
		} else {
			response.OkWithData(w, resp)
		}
	}
}
