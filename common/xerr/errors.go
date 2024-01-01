package xerr

import "fmt"

var message map[int]string

// 成功返回
const OK int = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const SERVER_COMMON_ERROR int = 100001
const REUQEST_PARAM_ERROR int = 100002
const TOKEN_EXPIRE_ERROR int = 100003
const TOKEN_GENERATE_ERROR int = 100004
const DB_ERROR int = 100005
const DB_UPDATE_AFFECTED_ZERO_ERROR int = 100006

func init() {
	message = make(map[int]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"
}

type CodeMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// New creates a new CodeMsg.
func NewCodeMsg(code int, msg string) CodeMsg {
	return CodeMsg{Code: code, Msg: msg}
}

func (e CodeMsg) GetErrMsg() string {
	return e.Msg
}

func (e CodeMsg) GetErrCode() int {
	return e.Code
}

func (e CodeMsg) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.Code, e.Msg)
}

func IsCodeErr(errCode int) bool {
	if _, ok := message[errCode]; ok {
		return true
	} else {
		return false
	}
}
