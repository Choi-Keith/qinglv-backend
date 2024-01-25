package upload

import (
	"net/http"

	"qinglv-backend/app/content/api/internal/logic/upload"
	"qinglv-backend/app/content/api/internal/svc"
	"qinglv-backend/app/content/api/internal/types"
	"qinglv-backend/common/response"
	"qinglv-backend/pkg/validate"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RemoveImagesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveImagesReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamsFail(w, err)
			return
		}
		if err := validate.Validate(req); err != nil {
			response.ParamsFail(w, err)
			return
		}
		l := upload.NewRemoveImagesLogic(r.Context(), svcCtx)
		err := l.RemoveImages(&req)
		if err != nil {
			response.FailCodeMsg(w, http.StatusBadRequest, err)
		} else {
			response.Ok(w)
		}
	}
}
