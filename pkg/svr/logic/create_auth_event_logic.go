package logic

import (
	"encoding/xml"
	"github.com/go-laoji/wecom-go-sdk"
	"github.com/go-laoji/wecom-go-sdk/pkg/svr/models"
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
	var corpAccessToken models.CorpAccessToken //　授权企业access token
	corpAccessToken.CorpId = corpInfo.ID
	corpAccessToken.AccessToken = resp.AccessToken
	corpAccessToken.ExpiresIn = resp.ExpiresIn
	engine.Save(&corpAccessToken)
	var corpPermanentCode models.CorpPermanentCode // 授权企业永久授权码
	corpPermanentCode.CorpId = corpInfo.ID
	corpPermanentCode.AuthCorpId = corpInfo.CorpId
	corpPermanentCode.PermanentCode = resp.PermanentCode
	corpPermanentCode.IsCustomizedApp = resp.AuthInfo.Agent[0].IsCustomizedApp
	engine.Save(&corpPermanentCode)
	//TODO:将授权企业的首次access token 写入缓存
	ww.Logger().Sugar().Info(resp)
}
