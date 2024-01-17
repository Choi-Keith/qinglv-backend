package topic

import (
	"net/http"

	"qinglv-backend/app/content/api/internal/logic/topic"
	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteTopicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteTopicReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamsFail(w, err)
			return
		}

		l := topic.NewDeleteTopicLogic(r.Context(), svcCtx)
		err := l.DeleteTopic(&req)
		if err != nil {
			response.FailCodeMsg(w, http.StatusBadRequest, err)
		} else {
			response.Ok(w)
		}
	}
}
