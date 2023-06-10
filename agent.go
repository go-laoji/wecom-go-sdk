package wework

import (
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type AgentGetResponse struct {
	internal.BizResponse
	AgentId        int    `json:"agentid"`
	Name           string `json:"name"`
	SquareLogoURL  string `json:"square_logo_url"`
	Description    string `json:"description"`
	AllowUserInfos struct {
		User []struct {
			Userid string `json:"userid"`
		} `json:"user"`
	} `json:"allow_userinfos"`
	AllowPartys struct {
		PartyId []int `json:"partyid"`
	} `json:"allow_partys"`
	AllowTags struct {
		TagId []int `json:"tagid"`
	} `json:"allow_tags"`
	Close              int    `json:"close"`
	RedirectDomain     string `json:"redirect_domain"`
	ReportLocationFlag int    `json:"report_location_flag"`
	IsReportEnter      int    `json:"isreportenter"`
	HomeURL            string `json:"home_url"`
}

// AgentGet 获取指定的应用详情
// https://open.work.weixin.qq.com/api/doc/90001/90143/90363#%E8%8E%B7%E5%8F%96access_token%E5%AF%B9%E5%BA%94%E7%9A%84%E5%BA%94%E7%94%A8%E5%88%97%E8%A1%A8
func (ww *weWork) AgentGet(corpId uint, agentId int) (resp AgentGetResponse) {
	_, err := ww.getRequest(corpId).
		SetQueryParam("agentid", fmt.Sprintf("%v", agentId)).
		SetResult(&resp).
		Get("/cgi-bin/agent/get")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type AgentListResponse struct {
	internal.BizResponse
	AgentList []struct {
		AgentId       int    `json:"agentid"`
		Name          string `json:"name"`
		SquareLogoURL string `json:"square_logo_url"`
	} `json:"agentlist"`
}

// AgentList 获取access_token对应的应用列表
// https://open.work.weixin.qq.com/api/doc/90001/90143/90363#%E8%8E%B7%E5%8F%96access_token%E5%AF%B9%E5%BA%94%E7%9A%84%E5%BA%94%E7%94%A8%E5%88%97%E8%A1%A8
func (ww *weWork) AgentList(corpId uint) (resp AgentListResponse) {
	_, err := ww.getRequest(corpId).
		SetResult(&resp).
		Get("/cgi-bin/agent/list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
