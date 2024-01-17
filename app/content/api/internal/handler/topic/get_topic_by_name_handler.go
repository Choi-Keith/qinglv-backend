package topic

import (
	"net/http"

	"qinglv-backend/app/content/api/internal/logic/topic"
	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTopicByNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTopicByNameReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamsFail(w, err)
			return
		}

		l := topic.NewGetTopicByNameLogic(r.Context(), svcCtx)
		resp, err := l.GetTopicByName(&req)
		if err != nil {
			response.FailCodeMsg(w, http.StatusBadRequest, err)
		} else {
			response.OkWithData(w, resp)
		}
	}
}
