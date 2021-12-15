package internal

import (
	"fmt"
)

type Error struct {
	ErrCode  int    `json:"errcode,omitempty"` // 错误代码
	ErrorMsg string `json:"errmsg,omitempty"`  // 错误消息
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误代码:%v，错误信息:%s", e.ErrCode, e.ErrorMsg)
}

var _ error = &Error{}
