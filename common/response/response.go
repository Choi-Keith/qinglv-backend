package response

import (
	"net/http"

	"qinglv-backend/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	status "google.golang.org/grpc/status"
)

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func Ok(w http.ResponseWriter) {
	resp := &Resp{
		Code:    0,
		Message: "Ok",
		Data:    struct{}{},
	}

	httpx.OkJson(w, resp)
}

func OkWithData(w http.ResponseWriter, data interface{}) {
	resp := &Resp{
		Code:    0,
		Message: "Ok",
		Data:    data,
	}

	httpx.OkJson(w, resp)
}

func Fail(w http.ResponseWriter, httpCode int, err error) {
	logx.Errorf("Failed: %v\n", err)
	causeErr := errors.Cause(err)
	var resp *Resp
	if e, ok := causeErr.(xerr.CodeMsg); ok {
		// 如果是自定义错误类型
		resp = &Resp{
			Code:    e.GetErrCode(),
			Message: e.GetErrMsg(),
			Data:    struct{}{},
		}

	} else {
		if gstatus, ok := status.FromError(causeErr); ok {
			// grpc错误
			grpcCode := int(gstatus.Code())
			if xerr.IsCodeErr(grpcCode) {
				resp = &Resp{
					Code:    grpcCode,
					Message: gstatus.Message(),
					Data:    struct{}{},
				}
			} else {
				resp = &Resp{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
					Data:    struct{}{},
				}
			}
		} else {
			resp = &Resp{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
				Data:    struct{}{},
			}
		}
	}

	httpx.WriteJson(w, httpCode, resp)

}

func FailCodeMsg(w http.ResponseWriter, httpCode int, err error) {
	Fail(w, httpCode, err)
}

func ParamsFail(w http.ResponseWriter, err error) {
	Fail(w, http.StatusBadRequest, err)
}
