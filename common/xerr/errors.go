package xerr

type CodeMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// New creates a new CodeMsg.
func NewCodeMsg(code int, msg string) *CodeMsg {
	return &CodeMsg{Code: code, Msg: msg}
}

var (
	ErrEmailHasExisted       = NewCodeMsg(10001, "邮箱已存在")
	ErrPasswordNotConsistent = NewCodeMsg(10002, "密码不一致")
	ErrRegisterParams        = NewCodeMsg(10003, "缺少参数")
)
