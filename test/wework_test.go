package test

import (
	"github.com/go-laoji/wecom-go-sdk"
	"testing"
)

var TestWeWork wework.IWeWork

const (
	AppSelf uint = iota + 1
	AppContact
	AppContactEdit
	AppEhr
)

func getAppSecretFunc(corpId uint) (corpid string, secret string, customizedApp bool) {
	switch corpId {
	case AppSelf:
		//纯自建应用
		return "TODO", "TODO", true
	case AppContact: // 通讯录同步应用
		return "TODO", "TODO", true
	case AppContactEdit: // 服务商通讯录可编辑应用
		return "TODO", "TODO", false
	case AppEhr: // 自建代开发ehr类应用
		return "TODO", "TODO", true
	}
	return "", "", false
}

func TestNewWeWork(t *testing.T) {
	t.SkipNow()
}

func testNewWeWork(t *testing.T) {
	if TestWeWork == nil {
		TestWeWork = wework.NewWeWork(wework.WeWorkConfig{
			CorpId:         "TODO",
			SuiteId:        "TODO",
			SuiteSecret:    "TODO",
			ProviderSecret: "TODO",
		})
		TestWeWork.SetAppSecretFunc(getAppSecretFunc)
		TestWeWork.UpdateSuiteTicket("TODO")
	}
}
