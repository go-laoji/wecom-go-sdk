package logic

import "github.com/go-laoji/wecom-go-sdk/v2"

func CancelAuthEventLogic(data []byte, ww wework.IWeWork) {
	ww.Logger().Sugar().Info(string(data))
}
