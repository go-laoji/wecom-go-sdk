package logic

import (
	"encoding/xml"
	"github.com/go-laoji/wecom-go-sdk/v2"
	"github.com/go-laoji/wecom-go-sdk/v2/pkg/svr/models"
	"github.com/jinzhu/copier"
)

type CreateAuthEvent struct {
	BizEvent
	AuthCode string `xml:"AuthCode"`
	State    string `xml:"State"`
}

// CreateAuthEventLogic 第三方应用管理员授权逻辑
// auth_code　10分钟有效
func CreateAuthEventLogic(data []byte, ww wework.IWeWork) {
	var event CreateAuthEvent
	if err := xml.Unmarshal(data, &event); err != nil {
		ww.Logger().Sugar().Error(err)
		return
	}
	ww.Logger().Sugar().Info(string(data))
	resp := ww.GetPermanentCode(event.AuthCode)
	var corpInfo models.CorpAuthInfo
	copier.Copy(&corpInfo, resp.AuthCorpInfo) // 授权企业信息
	engine.Save(&corpInfo)

	var corpAuthUser models.CorpAuthUserInfo // 授权企业管理员信息
	copier.Copy(&corpAuthUser, resp.AuthUserInfo)
	corpAuthUser.CorpId = corpInfo.ID
	engine.Save(&corpAuthUser)
	var corpAgent models.Agent
	// 通讯录应用授权时没有agent信息
	//　TODO:重构永久码存储结构
	if len(resp.AuthInfo.Agent) > 0 {
		copier.Copy(&corpAgent, resp.AuthInfo.Agent[0]) //默认取agent[0]
	}
	corpAgent.CorpId = corpInfo.ID
	corpAgent.AuthCorpId = corpInfo.CorpId
	corpAgent.PermanentCode = resp.PermanentCode
	engine.Save(&corpAgent)
	ww.Logger().Sugar().Info(resp)
}
