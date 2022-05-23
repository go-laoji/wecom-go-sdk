package wework

import (
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/badger/v2"
	"github.com/go-laoji/wecom-go-sdk/internal"
	"github.com/go-laoji/wecom-go-sdk/pkg/svr/models"
	"net/url"
	"os"
	"time"
)

// UpdateSuiteTicket 用于接收回调后更新sdk实例的suite ticket
func (ww *weWork) UpdateSuiteTicket(ticket string) {
	ww.suiteTicket = ticket
	ww.cache.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte("suiteTicket"), []byte(ticket)).WithTTL(time.Second * 600)
		err := txn.SetEntry(entry)
		return err
	})
	ww.requestSuiteToken()
}

type suiteTokenResponse struct {
	internal.BizResponse
	SuiteAccessToken string `json:"suite_access_token"`
	ExpiresIn        int    `json:"expires_in"`
}

func (ww weWork) requestSuiteToken() (resp suiteTokenResponse) {
	if ww.suiteTicket == "" {
		resp.ErrCode = 400
		resp.ErrorMsg = "suite ticket 未推送"
		logger.Sugar().Error("suite ticket 未推送")
		return
	}
	apiUrl := "/cgi-bin/service/get_suite_token"
	h := H{}
	h["suite_id"] = ww.suiteId
	h["suite_secret"] = ww.suiteSecret
	h["suite_ticket"] = ww.suiteTicket
	body, err := internal.HttpPost(apiUrl, h)
	if err != nil {
		resp.ErrCode = 400
		resp.ErrorMsg = err.Error()
		logger.Sugar().Error(err)
		return
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		resp.ErrCode = 400
		resp.ErrorMsg = err.Error()
		logger.Sugar().Error(err)
		return
	}
	return
}

func (ww weWork) getSuiteAccessToken() (token string) {
	var err error
	var item *badger.Item
	err = ww.cache.View(func(txn *badger.Txn) error {
		item, err = txn.Get([]byte("suiteToken"))
		if err == badger.ErrKeyNotFound {
			return err
		}
		item.Value(func(val []byte) error {
			token = string(val)
			return nil
		})
		return err
	})
	if err == badger.ErrKeyNotFound {
		if resp := ww.requestSuiteToken(); resp.ErrCode != 0 {
			logger.Sugar().Error(resp.ErrorMsg)
			return ""
		} else {
			token = resp.SuiteAccessToken
			ww.cache.Update(func(txn *badger.Txn) error {
				entry := badger.NewEntry([]byte("suiteToken"), []byte(token)).
					WithTTL(time.Second * time.Duration(resp.ExpiresIn))
				err = txn.SetEntry(entry)
				return err
			})
		}
	}
	return token
}

type GetPreAuthCodeResponse struct {
	internal.BizResponse
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetPreAuthCode 获取预授权码
// https://open.work.weixin.qq.com/api/doc/90001/90143/90601
func (ww weWork) GetPreAuthCode() (resp GetPreAuthCodeResponse) {
	queryParams := url.Values{}
	queryParams.Add("suite_access_token", ww.getSuiteAccessToken())
	apiUrl := fmt.Sprintf("/cgi-bin/service/get_pre_auth_code?%s", queryParams.Encode())
	body, err := internal.HttpGet(apiUrl)
	if err != nil {
		logger.Sugar().Error(err)
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
		return
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type DealerCorpInfo struct {
	Corpid   string `json:"corpid"`
	CorpName string `json:"corp_name"`
}
type AuthCorpInfo struct {
	CorpId            string `json:"corpid"`
	CorpName          string `json:"corp_name"`
	CorpType          string `json:"corp_type"`
	CorpSquareLogoURL string `json:"corp_square_logo_url"`
	CorpUserMax       int    `json:"corp_user_max"`
	CorpAgentMax      int    `json:"corp_agent_max"`
	CorpFullName      string `json:"corp_full_name"`
	VerifiedEndTime   int    `json:"verified_end_time"`
	SubjectType       int    `json:"subject_type"`
	CorpWxqrcode      string `json:"corp_wxqrcode"`
	CorpScale         string `json:"corp_scale"`
	CorpIndustry      string `json:"corp_industry"`
	CorpSubIndustry   string `json:"corp_sub_industry"`
}
type AuthUserInfo struct {
	UserId     string `json:"userid"`
	OpenUserId string `json:"open_userid"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
}
type RegisterCodeInfo struct {
	RegisterCode string `json:"register_code"`
	TemplateID   string `json:"template_id"`
	State        string `json:"state"`
}
type Agent struct {
	AgentId         int    `json:"agentid"`
	Name            string `json:"name"`
	RoundLogoURL    string `json:"round_logo_url"`
	SquareLogoURL   string `json:"square_logo_url"`
	Appid           int    `json:"appid"`
	AuthMode        int    `json:"auth_mode,omitempty"`
	IsCustomizedApp bool   `json:"is_customized_app,omitempty"`
	Privilege       struct {
		Level      int      `json:"level"`
		AllowParty []int    `json:"allow_party"`
		AllowUser  []string `json:"allow_user"`
		AllowTag   []int    `json:"allow_tag"`
		ExtraParty []int    `json:"extra_party"`
		ExtraUser  []string `json:"extra_user"`
		ExtraTag   []int    `json:"extra_tag"`
	} `json:"privilege,omitempty"`
	SharedFrom struct {
		Corpid string `json:"corpid"`
	} `json:"shared_from"`
}
type GetPermanentCodeResponse struct {
	internal.BizResponse
	AccessToken    string         `json:"access_token"`
	ExpiresIn      int            `json:"expires_in"`
	PermanentCode  string         `json:"permanent_code"`
	DealerCorpInfo DealerCorpInfo `json:"dealer_corp_info"`
	AuthCorpInfo   AuthCorpInfo   `json:"auth_corp_info"`
	AuthInfo       struct {
		Agent []Agent `json:"agent"`
	} `json:"auth_info"`
	AuthUserInfo     AuthUserInfo     `json:"auth_user_info"`
	RegisterCodeInfo RegisterCodeInfo `json:"register_code_info"`
	State            string           `json:"state"`
}

// GetPermanentCode 获取企业永久授权码
// https://open.work.weixin.qq.com/api/doc/90001/90143/90603
func (ww weWork) GetPermanentCode(authCode string) (resp GetPermanentCodeResponse) {
	queryParams := url.Values{}
	queryParams.Add("suite_access_token", ww.getSuiteAccessToken())
	apiUrl := fmt.Sprintf("/cgi-bin/service/get_permanent_code?%s", queryParams.Encode())
	h := H{}
	h["auth_code"] = authCode
	body, err := internal.HttpPost(apiUrl, h)
	if err != nil {
		logger.Sugar().Error(err)
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
		return
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetAuthInfoResponse struct {
	internal.BizResponse
	DealerCorpInfo DealerCorpInfo `json:"dealer_corp_info"`
	AuthCorpInfo   AuthCorpInfo   `json:"auth_corp_info"`
	AuthInfo       struct {
		Agent []Agent `json:"agent"`
	} `json:"auth_info"`
}

// GetAuthInfo 获取企业授权信息
// https://open.work.weixin.qq.com/api/doc/90001/90143/90604
func (ww weWork) GetAuthInfo(authCorpId, permanentCode string) (resp GetAuthInfoResponse) {
	queryParams := url.Values{}
	queryParams.Add("suite_access_token", ww.getSuiteAccessToken())
	apiUrl := fmt.Sprintf("/cgi-bin/service/get_auth_info?%s", queryParams.Encode())
	h := H{}
	h["auth_corpid"] = authCorpId
	h["permanent_code"] = permanentCode

	body, err := internal.HttpPost(apiUrl, h)
	if err != nil {
		logger.Sugar().Error(err)
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
		return
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type getCorpTokenResponse struct {
	internal.BizResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// 默认从数据库获取应用secret配置信息
// 同一corpid(企业微信主体ID号)可以配置多个应用
func (ww weWork) defaultAppSecretFunc(corpId uint) (corpid string, secret string, customizedApp bool) {
	var authAgent models.Agent
	ww.engine.Model(models.Agent{}).
		Where(models.Agent{CorpId: corpId}).
		First(&authAgent)
	return authAgent.AuthCorpId, authAgent.PermanentCode, authAgent.IsCustomizedApp
}

// 默认从数据库获取应用的agentid
func (ww weWork) defaultAgentIdFunc(corpId uint) (appId int) {
	var authAgent models.Agent
	ww.engine.Model(models.Agent{}).
		Where(models.Agent{CorpId: corpId}).
		First(&authAgent)
	return authAgent.AgentId
}

// 从数据库查询永久授权码和授权企业的企业微信id，获取对应的access token
func (ww weWork) requestCorpToken(corpId uint) (resp getCorpTokenResponse) {
	queryParams := url.Values{}
	var apiUrl string
	var body []byte
	var err error
	var corpid, secret string
	var customizedApp bool
	if ww.getAppSecretFunc != nil {
		corpid, secret, customizedApp = ww.getAppSecretFunc(corpId)
	} else {
		corpid, secret, customizedApp = ww.defaultAppSecretFunc(corpId)
	}
	// 兼容代开发应用/自建应用/三方应用的token获取
	if !customizedApp {
		queryParams.Add("suite_access_token", ww.getSuiteAccessToken())
		apiUrl = fmt.Sprintf("/cgi-bin/service/get_corp_token?%s", queryParams.Encode())
		h := H{}
		h["auth_corpid"] = corpid
		h["permanent_code"] = secret
		body, err = internal.HttpPost(apiUrl, h)
	} else {
		queryParams.Add("corpid", corpid)
		queryParams.Add("corpsecret", secret)
		apiUrl = fmt.Sprintf("/cgi-bin/gettoken?%s", queryParams.Encode())
		body, err = internal.HttpGet(apiUrl)
	}
	if err != nil {
		logger.Sugar().Error(err)
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
		return
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

func (ww *weWork) SetAppSecretFunc(f func(corpId uint) (corpid string, secret string, customizedApp bool)) {
	ww.getAppSecretFunc = f
}

func (ww *weWork) SetAgentIdFunc(f func(corpId uint) (agentId int)) {
	ww.getAgentIdFunc = f
}

func (ww weWork) getCorpToken(corpId uint) (token string) {
	var err error
	var item *badger.Item
	err = ww.cache.View(func(txn *badger.Txn) error {
		item, err = txn.Get([]byte(fmt.Sprintf("corpToken-%v", corpId)))
		if err == badger.ErrKeyNotFound {
			return err
		}
		item.Value(func(val []byte) error {
			token = string(val)
			return nil
		})
		return err
	})
	if err == badger.ErrKeyNotFound {
		if resp := ww.requestCorpToken(corpId); resp.ErrCode != 0 {
			logger.Sugar().Error(resp.ErrorMsg)
			return ""
		} else {
			token = resp.AccessToken
			ww.cache.Update(func(txn *badger.Txn) error {
				entry := badger.NewEntry([]byte(fmt.Sprintf("corpToken-%v", corpId)), []byte(token)).
					WithTTL(time.Second * time.Duration(resp.ExpiresIn))
				err = txn.SetEntry(entry)
				return err
			})
		}
	}
	return token
}

func (ww weWork) buildCorpQueryToken(corpId uint) url.Values {
	queryParams := url.Values{}
	queryParams.Add("access_token", ww.getCorpToken(corpId))
	if os.Getenv("debug") != "" {
		queryParams.Add("debug", "1")
	}
	return queryParams
}

type GetUserInfo3rdResponse struct {
	internal.BizResponse
	CorpId     string `json:"CorpId"`
	UserId     string `json:"UserId"`
	DeviceId   string `json:"DeviceId"`
	UserTicket string `json:"user_ticket"`
	ExpiresIn  int    `json:"expires_in"`
	OpenUserId string `json:"open_userid"`
}

// GetUserInfo3rd 获取访问用户身份
// https://open.work.weixin.qq.com/api/doc/90001/90143/91121
func (ww weWork) GetUserInfo3rd(code string) (resp GetUserInfo3rdResponse) {
	queryParams := url.Values{}
	queryParams.Add("suite_access_token", ww.getSuiteAccessToken())
	queryParams.Add("code", code)
	apiUrl := fmt.Sprintf("/cgi-bin/service/getuserinfo3rd?%s", queryParams.Encode())
	body, err := internal.HttpGet(apiUrl)
	if err != nil {
		logger.Sugar().Error(err)
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
		return
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetUserInfoDetail3rdResponse struct {
	internal.BizResponse
	CorpId string `json:"corpid"`
	UserId string `json:"userid"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Avatar string `json:"avatar"`
	QrCode string `json:"qr_code"`
}

// GetUserInfoDetail3rd 获取访问用户敏感信息
// https://open.work.weixin.qq.com/api/doc/90001/90143/91122
func (ww weWork) GetUserInfoDetail3rd(userTicket string) (resp GetUserInfoDetail3rdResponse) {
	queryParams := url.Values{}
	queryParams.Add("suite_access_token", ww.getSuiteAccessToken())
	apiUrl := fmt.Sprintf("/cgi-bin/service/getuserdetail3rd?%s", queryParams.Encode())
	h := H{}
	h["user_ticket"] = userTicket

	body, err := internal.HttpPost(apiUrl, h)
	if err != nil {
		logger.Sugar().Error(err)
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
		return
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetUserInfoResponse struct {
	internal.BizResponse
	UserId string `json:"UserId"`
	OpenId string `json:"OpenId"`
}

func (ww weWork) GetUserInfo(corpId uint, code string) (resp GetUserInfoResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	queryParams.Add("code", code)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/user/getuserinfo?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type GetAppQrCodeRequest struct {
	SuiteID    string `json:"suite_id"`
	Appid      int    `json:"appid,omitempty"`
	State      string `json:"state,omitempty"`
	Style      int    `json:"style,omitempty" validate:"omitempty,oneof=0 1 2 3 4"`
	ResultType int    `json:"result_type" validate:"required,oneof=2"`
}

type GetAppQrCodeResponse struct {
	internal.BizResponse
	QrCode string `json:"qrcode"`
}

// GetAppQrCode 获取应用二维码 仅支持二维码地址返回
// https://developer.work.weixin.qq.com/document/path/95430#36592
func (ww weWork) GetAppQrCode(request GetAppQrCodeRequest) (resp GetAppQrCodeResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := url.Values{}
	queryParams.Add("suite_access_token", ww.GetSuiteToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/service/get_app_qrcode?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
