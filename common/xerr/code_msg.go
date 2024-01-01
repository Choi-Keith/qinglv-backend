package xerr

var (
	ErrEmailHasExisted       = NewCodeMsg(200001, "邮箱已存在")
	ErrPasswordNotConsistent = NewCodeMsg(200002, "密码不一致")
	ErrRegisterParams        = NewCodeMsg(200003, "参数错误")
)
