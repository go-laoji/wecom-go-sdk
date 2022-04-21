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
	} `json:"servicer_list"`
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

type KfSyncMsgRequest struct {
	Cursor      string `json:"cursor"`
	Token       string `json:"token"`
	Limit       int    `json:"limit"`
	VoiceFormat int    `json:"voice_format"`
}

type KfSyncMsgResponse struct {
	internal.BizResponse
	NextCursor string `json:"next_cursor"`
	HasMore    bool   `json:"has_more"`
	MsgList    []struct {
		MsgId          string          `json:"msgid"`
		OpenKfId       string          `json:"open_kfid"`
		ExternalUserId string          `json:"external_userid"`
		SendTime       int             `json:"send_time"`
		Origin         int             `json:"origin"`
		ServicerUserId string          `json:"servicer_userid"`
		MsgType        string          `json:"msgtype"`
		Text           MsgText         `json:"text,omitempty"`
		Image          MsgImage        `json:"image,omitempty"`
		Voice          MsgVoice        `json:"voice,omitempty"`
		Video          MsgVideo        `json:"video,omitempty"`
		File           MsgFile         `json:"file,omitempty"`
		Location       MsgLocation     `json:"location,omitempty"`
		Link           MsgLink         `json:"link,omitempty"`
		BusinessCard   MsgBusinessCard `json:"business_card,omitempty"`
		MiniProgram    MsgMiniProgram  `json:"miniprogram,omitempty"`
		MsgMenu        MsgMenu         `json:"msgmenu,omitempty"`
		Event          MsgEvent        `json:"event,omitempty"`
	} `json:"msg_list"`
}

type MsgText struct {
	Content string `json:"content"`
	MenuId  string `json:"menu_id"`
}
type MsgImage struct {
	MediaId string `json:"media_id"`
}
type MsgVoice struct {
	MediaId string `json:"media_id"`
}
type MsgVideo struct {
	MediaId string `json:"media_id"`
}
type MsgFile struct {
	MediaId string `json:"media_id"`
}
type MsgLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
}
type MsgLink struct {
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Url    string `json:"url"`
	PicUrl string `json:"pic_url"`
}
type MsgBusinessCard struct {
	UserId string `json:"userid"`
}
type MsgMiniProgram struct {
	Title        string `json:"title"`
	AppId        string `json:"appid"`
	PagePath     string `json:"pagepath"`
	ThumbMediaId string `json:"thumb_media_id"`
}
type MsgMenu struct {
	HeadContent string `json:"head_content"`
	List        []struct {
		Type  string `json:"type"`
		Click struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		} `json:"click,omitempty"`
		View struct {
			URL     string `json:"url"`
			Content string `json:"content"`
		} `json:"view,omitempty"`
		Miniprogram struct {
			Appid    string `json:"appid"`
			Pagepath string `json:"pagepath"`
			Content  string `json:"content"`
		} `json:"miniprogram,omitempty"`
	} `json:"list"`
	TailContent string `json:"tail_content"`
}
type MsgEvent struct {
	EventType      string `json:"event_type"`
	OpenKfid       string `json:"open_kfid"`
	ExternalUserid string `json:"external_userid"`
	Scene          string `json:"scene"`
	SceneParam     string `json:"scene_param"`
	WelcomeCode    string `json:"welcome_code"`
	WechatChannels struct {
		Nickname string `json:"nickname"`
	} `json:"wechat_channels"`
}

func (ww weWork) KfSyncMsg(corpId uint, request KfSyncMsgRequest) (resp KfSyncMsgResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/kf/sync_msg?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type SendMsgRequest struct {
	ToUser      string          `json:"touser" validate:"required"`
	OpenKfId    string          `json:"open_kfid" validate:"required"`
	MsgId       string          `json:"msgid"  validate:"required"`
	MsgType     string          `json:"msgtype"`
	Text        *MsgText        `json:"text,omitempty"`
	Image       *MsgImage       `json:"image,omitempty"`
	Voice       *MsgVoice       `json:"voice,omitempty"`
	Video       *MsgVideo       `json:"video,omitempty"`
	File        *MsgFile        `json:"file,omitempty"`
	Location    *MsgLocation    `json:"location,omitempty"`
	Link        *MsgLink        `json:"link,omitempty"`
	MiniProgram *MsgMiniProgram `json:"miniprogram,omitempty"`
	MsgMenu     *MsgMenu        `json:"msgmenu,omitempty"`
}

type SendMsgResponse struct {
	internal.BizResponse
	MsgId string `json:"msgid"`
}

func (ww weWork) KfSendMsg(corpId uint, request SendMsgRequest) (resp SendMsgResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/kf/send_msg?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type SendMsgOnEventRequest struct {
	Code    string   `json:"code"`
	MsgId   string   `json:"msgid"`
	MsgType string   `json:"msgtype" validate:"required,oneof=text msgmenu"`
	Text    *MsgText `json:"text,omitempty"`
	MsgMenu *MsgMenu `json:"msgmenu,omitempty"`
}

func (ww weWork) KfSendMsgOnEvent(corpId uint, request SendMsgOnEventRequest) (resp SendMsgResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/kf/send_msg_on_event?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type KfCustomerBatchGetResponse struct {
	internal.BizResponse
	CustomerList []struct {
		ExternalUserid      string `json:"external_userid"`
		Nickname            string `json:"nickname"`
		Avatar              string `json:"avatar"`
		Gender              int    `json:"gender"`
		Unionid             string `json:"unionid"`
		EnterSessionContext struct {
			Scene          string `json:"scene"`
			SceneParam     string `json:"scene_param"`
			WechatChannels struct {
				Nickname string `json:"nickname"`
			} `json:"wechat_channels"`
		} `json:"enter_session_context"`
	} `json:"customer_list"`
	InvalidExternalUserid []string `json:"invalid_external_userid"`
}

func (ww weWork) KfCustomerBatchGet(corpId uint, userList []string, needEnterSessionContext int) (resp KfCustomerBatchGetResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	params := H{}
	if needEnterSessionContext == 1 {
		params["need_enter_session_context"] = 1
	}
	params["external_userid_list"] = userList
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/kf/customer/batchget?%s", queryParams.Encode()), params)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
