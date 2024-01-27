package collection

import (
	"net/http"

	"qinglv-backend/app/operation/api/internal/logic/collection"
	"qinglv-backend/app/operation/api/internal/svc"
	"qinglv-backend/app/operation/api/internal/types"
	"qinglv-backend/common/response"
	"qinglv-backend/pkg/validate"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCollectionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCollectionListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamsFail(w, err)
			return
		}
		if err := validate.Validate(req); err != nil {
			response.ParamsFail(w, err)
			return
		}

		l := collection.NewGetCollectionListLogic(r.Context(), svcCtx, r)
		resp, err := l.GetCollectionList(&req)
		if err != nil {
			response.FailCodeMsg(w, http.StatusBadRequest, err)
		} else {
			response.OkWithData(w, resp)
		}
	}
}
