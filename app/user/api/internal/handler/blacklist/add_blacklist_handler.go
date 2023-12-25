package blacklist

import (
	"net/http"

	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"qinglv-backend/app/user/api/internal/logic/blacklist"
	"qinglv-backend/app/user/api/internal/svc"
	"qinglv-backend/app/user/api/internal/types"
)

func AddBlacklistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddBlacklistReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := blacklist.NewAddBlacklistLogic(r.Context(), svcCtx)
		err := l.AddBlacklist(&req)
		result.HttpResult(r, w, nil, err)
	}
}
