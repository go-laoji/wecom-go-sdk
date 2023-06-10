package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type IdConvertExternalTagIdResponse struct {
	internal.BizResponse
	Items []struct {
		ExternalTagid     string `json:"external_tagid"`
		OpenExternalTagid string `json:"open_external_tagid"`
	} `json:"items"`
	InvalidExternalTagidList []string `json:"invalid_external_tagid_list"`
}

func (ww *weWork) IdConvertExternalTagId(corpId uint, tagIdList []string) (resp IdConvertExternalTagIdResponse) {
	if len(tagIdList) > 1000 {
		resp.ErrCode = 500
		resp.ErrorMsg = "参数列表不能超过1000个"
		return
	}
	p := H{"external_tagid_list": tagIdList}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/idconvert/external_tagid")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type CorpIdToOpenCorpIdResponse struct {
	internal.BizResponse
	OpenCorpid string `json:"open_corpid"`
}

func (ww *weWork) CorpIdToOpenCorpId(corpId string) (resp CorpIdToOpenCorpIdResponse) {
	p := H{"corpid": corpId}
	_, err := ww.getProviderRequest().SetBody(p).SetResult(&resp).
		Post("/cgi-bin/service/corpid_to_opencorpid")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
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

func (ww *weWork) UserIdToOpenUserId(corpId uint, userIdList []string) (resp UserIdToOpenUserIdResponse) {
	if len(userIdList) > 1000 {
		resp.ErrCode = 500
		resp.ErrorMsg = "参数列表不能超过1000个"
		return
	}
	p := H{"userid_list": userIdList}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/batch/userid_to_openuserid")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
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

func (ww *weWork) GetNewExternalUserId(corpId uint, userIdList []string) (resp GetNewExternalUserIdResponse) {
	if len(userIdList) > 1000 {
		resp.ErrCode = 500
		resp.ErrorMsg = "参数列表不能超过1000个"
		return
	}
	p := H{"external_userid_list": userIdList}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/externalcontact/get_new_external_userid")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GroupChatGetNewExternalUserIdRequest struct {
	ChatID             string   `json:"chat_id" validate:"required"`
	ExternalUseridList []string `json:"external_userid_list" validate:"required,max=1000"`
}

func (ww *weWork) GroupChatGetNewExternalUserId(corpId uint, request GroupChatGetNewExternalUserIdRequest) (resp GetNewExternalUserIdResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/externalcontact/groupchat/get_new_external_userid")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type IdConvertOpenKfIdResponse struct {
	internal.BizResponse
	Items []struct {
		OpenKfId    string `json:"open_kfid"`
		NewOpenKfId string `json:"new_open_kfid"`
	} `json:"items"`
	InvalidOpenKfIdList []string `json:"invalid_open_kfid_list"`
}

func (ww *weWork) IdConvertOpenKfId(corpId uint, kfList []string) (resp IdConvertOpenKfIdResponse) {
	if len(kfList) > 1000 {
		resp.ErrCode = 500
		resp.ErrorMsg = "参数列表不能超过1000个"
		return
	}
	p := H{"open_kfid_list": kfList}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/idconvert/open_kfid")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
