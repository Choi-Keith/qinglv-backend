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

func UploadImagesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadImagesReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamsFail(w, err)
			return
		}
		if err := validate.Validate(req); err != nil {
			response.ParamsFail(w, err)
			return
		}

		l := upload.NewUploadImagesLogic(r.Context(), svcCtx, r)
		resp, err := l.UploadImages(&req, r)
		if err != nil {
			response.FailCodeMsg(w, http.StatusBadRequest, err)
		} else {
			response.OkWithData(w, resp)
		}
	}
}
