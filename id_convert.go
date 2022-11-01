package wework

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/internal"
	"net/url"
)

type IdConvertExternalTagIdResponse struct {
	internal.BizResponse
	Items []struct {
		ExternalTagid     string `json:"external_tagid"`
		OpenExternalTagid string `json:"open_external_tagid"`
	} `json:"items"`
	InvalidExternalTagidList []string `json:"invalid_external_tagid_list"`
}

func (ww weWork) IdConvertExternalTagId(corpId uint, tagIdList []string) (resp IdConvertExternalTagIdResponse) {
	if len(tagIdList) > 1000 {
		resp.ErrCode = 500
		resp.ErrorMsg = "参数列表不能超过1000个"
		return
	}
	p := H{"external_tagid_list": tagIdList}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/idconvert/external_tagid?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type CorpIdToOpenCorpIdResponse struct {
	internal.BizResponse
	OpenCorpid string `json:"open_corpid"`
}

func (ww weWork) CorpIdToOpenCorpId(corpId string) (resp CorpIdToOpenCorpIdResponse) {
	p := H{"corpid": corpId}
	queryParams := url.Values{}
	queryParams.Add("provider_access_token", ww.getProviderToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/service/corpid_to_opencorpid?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type UserIdToOpenUserIdResponse struct {
	internal.BizResponse
	OpenUserIdList []struct {
		UserId     string `json:"userid"`
		OpenUserId string `json:"open_userid"`
	} `json:"open_userid_list"`
	InvalidUserIdList []string `json:"invalid_userid_list"`
}

func (ww weWork) UserIdToOpenUserId(corpId uint, userIdList []string) (resp UserIdToOpenUserIdResponse) {
	if len(userIdList) > 1000 {
		resp.ErrCode = 500
		resp.ErrorMsg = "参数列表不能超过1000个"
		return
	}
	p := H{"userid_list": userIdList}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/batch/userid_to_openuserid?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type GetNewExternalUserIdResponse struct {
	internal.BizResponse
	Items []struct {
		ExternalUserid    string `json:"external_userid"`
		NewExternalUserid string `json:"new_external_userid"`
	} `json:"items"`
}

func (ww weWork) GetNewExternalUserId(corpId uint, userIdList []string) (resp GetNewExternalUserIdResponse) {
	if len(userIdList) > 1000 {
		resp.ErrCode = 500
		resp.ErrorMsg = "参数列表不能超过1000个"
		return
	}
	p := H{"external_userid_list": userIdList}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/get_new_external_userid?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type GroupChatGetNewExternalUserIdRequest struct {
	ChatID             string   `json:"chat_id" validate:"required"`
	ExternalUseridList []string `json:"external_userid_list" validate:"required,max=1000"`
}

func (ww weWork) GroupChatGetNewExternalUserId(corpId uint, request GroupChatGetNewExternalUserIdRequest) (resp GetNewExternalUserIdResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/groupchat/get_new_external_userid?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
