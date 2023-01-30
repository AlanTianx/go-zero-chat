package errorm

import (
	"fmt"
	"strings"
)

type Error struct {
	errCode int
	errMsg  string
	oErr    error
}

//  Error()将把errCode、errMsg、oErr的值拼接返回
func (e *Error) Error() string {
	return fmt.Sprintf("ErrCode:%d,ErrMsg:%s,oErr:%v", e.errCode, e.errMsg, e.oErr)
}

// errCode 必须
// Error()将把errCode、errMsg、oErr的值拼接返回
func NewError(errCode int, errMsg string, oErr error) *Error {
	errMsg = strings.TrimRight(fmt.Sprintf("%s %s", msg[errCode], errMsg), " ")

	return &Error{
		errCode: errCode,
		errMsg:  errMsg,
		oErr:    oErr,
	}
}

func (e *Error) GetErrCode() int {
	return e.errCode
}

func (e *Error) GetErrMsg() string {
	return e.errMsg
}
