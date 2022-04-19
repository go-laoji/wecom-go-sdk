package wework

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/internal"
)

type KfAccount struct {
	OpenKfId string `json:"open_kfid,omitempty"`
	Name     string `json:"name" validate:"required"`
	MediaId  string `json:"media_id" validate:"required"`
}
type KfAccountAddResponse struct {
	internal.BizResponse
	OpenKfId string `json:"open_kfid"`
}

func (ww weWork) KfAccountAdd(corpId uint, account KfAccount) (resp KfAccountAddResponse) {
	if ok := validate.Struct(account); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/kf/account/add?%s", queryParams.Encode()), account)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

func (ww weWork) KfAccountDel(corpId uint, kfId string) (resp internal.BizResponse) {
	if kfId == "" {
		resp.ErrCode = 500
		resp.ErrorMsg = "客服ID必填"
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	params := H{}
	params["open_kfid"] = kfId
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/kf/account/del?%s", queryParams.Encode()), params)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

func (ww weWork) KfAccountUpdate(corpId uint, account KfAccount) (resp internal.BizResponse) {
	if ok := validate.Struct(account); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/kf/account/update?%s", queryParams.Encode()), account)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type KfAccountListResponse struct {
	internal.BizResponse
	AccountList []struct {
		OpenKfId string `json:"open_kfid"`
		Name     string `json:"name"`
		Avatar   string `json:"avatar"`
	} `json:"account_list"`
}

func (ww weWork) KfAccountList(corpId uint) (resp KfAccountListResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/kf/account/list?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type KfAccContactWayResponse struct {
	internal.BizResponse
	Url string `json:"url"`
}

func (ww weWork) KfAddContactWay(corpId uint, kfId string, scene string) (resp KfAccContactWayResponse) {
	if kfId == "" {
		resp.ErrCode = 500
		resp.ErrorMsg = "客服ID必填"
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	params := H{}
	params["open_kfid"] = kfId
	params["scene"] = scene
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/kf/add_contact_way?%s", queryParams.Encode()), params)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type KfServicerRequest struct {
	OpenKfId   string   `json:"open_kfid" validate:"required"`
	UserIdList []string `json:"userid_list" validate:"required,max=100"`
}

type KfServicerResponse struct {
	internal.BizResponse
	ResultList []struct {
		UserId string `json:"userid"`
		internal.BizResponse
	} `json:"result_list"`
}

func (ww weWork) KfServicerAdd(corpId uint, request KfServicerRequest) (resp KfServicerResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/kf/servicer/add?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

func (ww weWork) KfServicerDel(corpId uint, request KfServicerRequest) (resp KfServicerResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/kf/servicer/del?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type KfServicerListResponse struct {
	internal.BizResponse
	ServicerList []struct {
		UserId string `json:"userid"`
		Status uint   `json:"status"`
	}
}

func (ww weWork) KfServicerList(corpId uint, kfId string) (resp KfServicerListResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	queryParams.Add("open_kfid", kfId)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/kf/servicer/list?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type KfServiceStateGetRequest struct {
	OpenKfId       string `json:"open_kfid" validate:"required"`
	ExternalUserId string `json:"external_userid" validate:"required"`
}

type KfServiceStateGetResponse struct {
	internal.BizResponse
	ServiceState   int    `json:"service_state"`
	ServicerUserId string `json:"servicer_userid"`
}

func (ww weWork) KfServiceStateGet(corpId uint, request KfServiceStateGetRequest) (resp KfServiceStateGetResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/kf/service_state/get?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type KfServiceStateTransRequest struct {
	OpenKfId       string `json:"open_kfid" validate:"required"`
	ExternalUserId string `json:"external_userid" validate:"required"`
	ServiceState   int    `json:"service_state" validate:"required,oneof=0 1 2 3 4"`
	ServicerUserId string `json:"servicer_userid"`
}

type KfServiceStateTransResponse struct {
	internal.BizResponse
	MsgCode string `json:"msg_code"`
}

func (ww weWork) KfServiceStateTrans(corpId uint, request KfServiceStateTransRequest) (resp KfServiceStateTransResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/kf/service_state/trans?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
