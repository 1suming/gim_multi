package errs

import "fmt"

// 定义错误码
type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

// 定义错误
type Err struct {
	Code    int    // 错误码
	Message string // 展示给用户看的
	Errord  error  // 保存内部错误信息
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Errord)
}
func (err *Err) New(errno *Errno, serr error) *Err {
	return &Err{
		Code:    errno.Code,
		Message: errno.Message,
		Errord:  serr,
	}
}

// 解码错误, 获取 Code 和 Message
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *Err:
		if typed.Code == ErrParam.Code {
			typed.Message = typed.Message + " 具体是 " + typed.Errord.Error()
		}
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}

//func NewError(code int64, msg string) MyError {
//	return MyError{
//		code: code,
//		msg:  msg,
//	}
//}

var (
	OK   = &Errno{Code: 0, Message: "OK"}
	FAIL = &Errno{Code: -1, Message: "fail"}
	// 系统错误, 前缀为 100
	InternalServerError = &Errno{Code: 10001, Message: "内部服务器错误"}
	ErrParam            = &Errno{Code: 10002, Message: "请求参数错误"}
	ErrTokenSign        = &Errno{Code: 10003, Message: "签名 jwt 时发生错误"}
	ErrEncrypt          = &Errno{Code: 10004, Message: "加密用户密码时发生错误"}

	// 数据库错误, 前缀为 201
	ErrDatabase = &Errno{Code: 20100, Message: "数据库错误"}
	ErrFill     = &Errno{Code: 20101, Message: "从数据库填充 struct 时发生错误"}

	// 认证错误, 前缀是 202
	ErrValidation   = &Errno{Code: 20201, Message: "验证失败"}
	ErrTokenInvalid = &Errno{Code: 20202, Message: "jwt 是无效的"}

	// 用户错误, 前缀为 203
	ErrUserNotFound      = &Errno{Code: 20301, Message: "用户没找到"}
	ErrPasswordIncorrect = &Errno{Code: 20302, Message: "密码错误"}
)
