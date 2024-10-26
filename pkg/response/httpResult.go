package response

import (
	"encoding/json"
	"gim/pkg/errs"
)

type Response struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

// 自定义响应信息
func (res *Response) WithMsg(message string) Response {
	return Response{
		Code: res.Code,
		Msg:  message,
		Data: res.Data,
	}
}

// 追加响应数据
func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

// ToString 返回 JSON 格式的错误详情
func (res *Response) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Code: res.Code,
		Msg:  res.Msg,
		Data: res.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}

// 构造函数
func Resp(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
func Ok(data interface{}) *Response {
	return &Response{
		Code: 0,
		Msg:  "",
		Data: data,
	}
}
func Errno(err *errs.Errno) *Response {
	return &Response{
		Code: err.Code,
		Msg:  err.Message,
	}
}
func Err(err *errs.Err) *Response {
	return &Response{
		Code: err.Code,
		Msg:  err.Message,
	}
}
