package notification

import (
	"net/http"

	"qinglv-backend/app/user/api/internal/logic/notification"
	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/common/response"
)

func GetUnreadMessageCountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := notification.NewGetUnreadMessageCountLogic(r.Context(), svcCtx)
		resp, err := l.GetUnreadMessageCount()
		if err != nil {
			response.FailCodeMsg(w, http.StatusBadRequest, err)
		} else {
			response.OkWithData(w, resp)
		}
	}
}
