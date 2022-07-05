package wework

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/internal"
)

type GetPermitUserListResponse struct {
	internal.BizResponse
	Ids []string `json:"ids"`
}

func (ww weWork) GetPermitUserList(corpId uint, T int) (resp GetPermitUserListResponse, err error) {
	if T > 3 || T < 0 {
		resp.ErrCode = 500
		resp.ErrorMsg = "type 取值范围出错,只能是1、2、3"
		err = errors.New("type 取值范围出错,只能是1、2、3")
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	h := H{}
	if T != 0 {
		h["type"] = T
	}
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/msgaudit/get_permit_user_list?%s", queryParams.Encode()), h)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		err = json.Unmarshal(body, &resp)
	}
	return
}

type CheckSingleAgreeRequest struct {
	Info []struct {
		Userid          string `json:"userid" validate:"required"`
		Exteranalopenid string `json:"exteranalopenid" validate:"required"`
	} `json:"info" validate:"required"`
}

type CheckSingleAgreeResponse struct {
	internal.BizResponse
	Agreeinfo []struct {
		StatusChangeTime int    `json:"status_change_time"`
		Userid           string `json:"userid"`
		Exteranalopenid  string `json:"exteranalopenid"`
		AgreeStatus      string `json:"agree_status"`
	} `json:"agreeinfo"`
}

func (ww weWork) CheckSingleAgree(corpId uint, request CheckSingleAgreeRequest) (resp CheckSingleAgreeResponse, err error) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/msgaudit/check_single_agree?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		err = json.Unmarshal(body, &resp)
	}
	return
}

type GetAuditGroupChatResponse struct {
	internal.BizResponse
	Roomname       string `json:"roomname"`
	Creator        string `json:"creator"`
	RoomCreateTime int    `json:"room_create_time"`
	Notice         string `json:"notice"`
	Members        []struct {
		Memberid string `json:"memberid"`
		Jointime int    `json:"jointime"`
	} `json:"members"`
}

func (ww weWork) GetAuditGroupChat(corpId uint, roomId string) (resp GetAuditGroupChatResponse, err error) {
	if roomId == "" {
		resp.ErrCode = 500
		resp.ErrorMsg = "roomId 必填,且只能为内部群ID"
		err = errors.New("roomId 必填,且只能为内部群ID")
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	h := H{}
	h["roomid"] = roomId
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/msgaudit/groupchat/get?%s", queryParams.Encode()), h)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		err = json.Unmarshal(body, &resp)
	}
	return
}
