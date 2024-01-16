package topic

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"qinglv-backend/app/content/api/internal/logic/topic"
	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
)

func UpdateTopicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateTopicReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := topic.NewUpdateTopicLogic(r.Context(), svcCtx)
		err := l.UpdateTopic(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
