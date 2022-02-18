package logic

import "github.com/go-laoji/wecom-go-sdk"

func CancelAuthEventLogic(data []byte, ww wework.IWeWork) {
	ww.Logger().Sugar().Info(string(data))
}
