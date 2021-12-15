package logic

import "github.com/go-laoji/wework"

func CancelAuthEventLogic(data []byte, ww wework.IWeWork) {
	ww.Logger().Sugar().Info(string(data))
}
