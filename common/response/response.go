package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"qinglv-backend/common/xerr"
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

func Fail(w http.ResponseWriter, status int, code int, msg string) {
	resp := &Resp{
		Code:    code,
		Message: msg,
		Data:    struct{}{},
	}

	httpx.WriteJson(w, status, resp)
}

func FailCodeMsg(w http.ResponseWriter, status int, err xerr.CodeMsg) {
	Fail(w, status, err.Code, err.Msg)
}

func ParamsFail(w http.ResponseWriter, code int, msg string) {
	Fail(w, http.StatusBadRequest, code, msg)
}
