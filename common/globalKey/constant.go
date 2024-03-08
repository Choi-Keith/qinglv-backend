package globalKey

/**
global constant key
*/

// 软删除
var DelStateNo int64 = 0  //未删除
var DelStateYes int64 = 1 //已删除

// 时间格式化模版DS
var DateTimeFormatTplStandardDateTime = "Y-m-d H:i:s"
var DateTimeFormatTplStandardDate = "Y-m-d"
var DateTimeFormatTplStandardTime = "H:i:s"

var VerifyEmailCodePrefixKey = "cache:qUser:email:verify:code:"
var VerifyForgotPasswordEmailCodePrefixKey = "cache:qUser:email:verify:code:forgot:"
var TokenPrefixKey = "cache:qUser:token:id:"

const (
	PostVisitPublic = iota + 1
	PostVisitFriend
	PostVisitPrivate
)

const (
	ThumbNo  = 0
	ThumbYes = 1
)
