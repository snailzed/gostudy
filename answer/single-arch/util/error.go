package util

//错误码
const (
	SUCCESS               = iota
	ErrEmptyParam         = 1000
	ErrServer             = 1001
	ErrUserExists         = 1002
	ErrUserNotExists      = 1003
	ErrUserOrPassword     = 1004
	ErrTitleContainWord   = 1005
	ErrContentContainWord = 1006
	ErrNeedUserLogin      = 1007
)

//错误码===》错误信息
var exceptionMessage = map[int]string{
	ErrEmptyParam:         "参数错误",
	ErrServer:             "系统异常，请稍后再试",
	ErrUserExists:         "用户已存在",
	ErrUserNotExists:      "账户或密码错误",
	ErrUserOrPassword:     "账户或密码错误",
	ErrTitleContainWord:   "标题包含敏感字符",
	ErrContentContainWord: "内容包含敏感字符",
	ErrNeedUserLogin:      "请先登录",
}

//获取错误信息
func GetErrorMessage(code int) string {
	ret, ok := exceptionMessage[code]
	if ok {
		return ret
	}
	return "UNKOWN_ERR"
}
