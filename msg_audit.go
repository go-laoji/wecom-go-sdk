package wework

import (
	"errors"
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type GetPermitUserListResponse struct {
	internal.BizResponse
	Ids []string `json:"ids"`
}

func (ww *weWork) GetPermitUserList(corpId uint, T int) (resp GetPermitUserListResponse, err error) {
	if T > 3 || T < 0 {
		resp.ErrCode = 500
		resp.ErrorMsg = "type 取值范围出错,只能是1、2、3"
		err = errors.New("type 取值范围出错,只能是1、2、3")
		return
	}
	h := H{}
	if T != 0 {
		h["type"] = T
	}
	_, err = ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/msgaudit/get_permit_user_list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
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

func (ww *weWork) CheckSingleAgree(corpId uint, request CheckSingleAgreeRequest) (resp CheckSingleAgreeResponse, err error) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err = ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/msgaudit/check_single_agree")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
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

func (ww *weWork) GetAuditGroupChat(corpId uint, roomId string) (resp GetAuditGroupChatResponse, err error) {
	if roomId == "" {
		resp.ErrCode = 500
		resp.ErrorMsg = "roomId 必填,且只能为内部群ID"
		err = errors.New("roomId 必填,且只能为内部群ID")
		return
	}
	h := H{}
	h["roomid"] = roomId
	_, err = ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/msgaudit/groupchat/get")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
