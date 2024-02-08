package tag

import (
	"net/http"

	"qinglv-backend/app/content/api/internal/logic/tag"
	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/common/response"
	"qinglv-backend/pkg/validate"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddTagReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamsFail(w, err)
			return
		}
		if err := validate.Validate(req); err != nil {
			response.FailCodeMsg(w, http.StatusBadRequest, err)
			return
		}

		l := tag.NewAddTagLogic(r.Context(), svcCtx)
		err := l.AddTag(&req)
		if err != nil {
			response.FailCodeMsg(w, http.StatusBadRequest, err)
		} else {
			response.Ok(w)
		}
	}
}
