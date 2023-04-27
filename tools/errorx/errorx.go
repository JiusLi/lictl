package errorx

import (
	"errors"
	"fmt"
)

type ErrorInfo struct {
	Code int
	Msg  string
	Err  error
}

func New(code int, msg string, err error, args ...any) *ErrorInfo {
	return &ErrorInfo{
		Code: code,
		Msg:  fmt.Sprintf(msg, args...),
		Err:  err,
	}
}

func (e *ErrorInfo) Error() string {
	return e.Err.Error()
}

var A = New(400, "400错误:%s", errors.New("我是不展示给前端的400内部错误"), "我可以是msg的详细报错信息")
