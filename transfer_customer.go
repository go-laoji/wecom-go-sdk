package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type TransferCustomerRequest struct {
	HandoverUserId     string   `json:"handover_userid" validate:"required"`
	TakeoverUserId     string   `json:"takeover_userid" validate:"required"`
	ExternalUserId     []string `json:"external_userid" validate:"required"`
	TransferSuccessMsg string   `json:"transfer_success_msg,omitempty" validate:"omitempty,max=200"`
}

type TransferCustomerResponse struct {
	internal.BizResponse
	Customer []struct {
		ExternalUserId string `json:"external_userid"`
		ErrCode        int    `json:"errcode"`
	}
}

// TransferCustomer 分配在职成员的客户
// 企业可通过此接口，转接在职成员的客户给其他成员。
// 参考连接 https://open.work.weixin.qq.com/api/doc/90001/90143/94096
func (ww *weWork) TransferCustomer(corpId uint, request TransferCustomerRequest) (resp TransferCustomerResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/externalcontact/transfer_customer")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type TransferResultRequest struct {
	HandoverUserId string `json:"handover_userid" validate:"required"`
	TakeoverUserId string `json:"takeover_userid" validate:"required"`
	Cursor         string `json:"cursor"`
}

type TransferResultResponse struct {
	internal.BizResponse
	Customer []struct {
		ExternalUserId string `json:"external_userid"`
		Status         int    `json:"status"`
		TakeoverTime   uint64 `json:"takeover_time"`
	} `json:"customer"`
	NextCursor string `json:"next_cursor"`
}

// TransferResult 查询客户接替状态
// 企业和第三方可通过此接口查询在职成员的客户转接情况。
// 参考连接　https://open.work.weixin.qq.com/api/doc/90001/90143/94097
func (ww *weWork) TransferResult(corpId uint, request TransferResultRequest) (resp TransferResultResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/externalcontact/transfer_result")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type UnAssignedRequest struct {
	PageId   int    `json:"page_id" validate:"required_without=Cursor,omitempty"`
	PageSize int    `json:"page_size" validate:"max=1000"`
	Cursor   string `json:"cursor" validate:"required_without=PageId,omitempty"`
}

type UnAssignedInfo struct {
	HandoverUserId string `json:"handover_userid"`
	ExternalUserId string `json:"external_userid"`
	DimissionTime  uint64 `json:"dimission_time"`
}

type UnAssignedResponse struct {
	internal.BizResponse
	Info       []UnAssignedInfo `json:"info"`
	IsLast     bool             `json:"is_last"`
	NextCursor string           `json:"next_cursor"`
}

// GetUnassignedList 获取待分配的离职成员列表
// 参考连接　https://open.work.weixin.qq.com/api/doc/90001/90143/92273
func (ww *weWork) GetUnassignedList(corpId uint, request UnAssignedRequest) (resp UnAssignedResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/externalcontact/get_unassigned_list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// TransferCustomerResigned 分配离职成员的客户;不可设置　TransferSuccessMsg
// handover_userid必须是已离职用户
// external_userid必须是handover_userid的客户
// 参考连接　https://open.work.weixin.qq.com/api/doc/90001/90143/94100
func (ww *weWork) TransferCustomerResigned(corpId uint, request TransferCustomerRequest) (resp TransferCustomerResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/externalcontact/resigned/transfer_customer")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// TransferResultResigned 查询客户接替状态
// 参考连接　https://open.work.weixin.qq.com/api/doc/90001/90143/94101
func (ww *weWork) TransferResultResigned(corpId uint, request TransferResultRequest) (resp TransferResultResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/externalcontact/resigned/transfer_result")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GroupChatTransferRequest struct {
	ChatIdList []string `json:"chat_id_list"`
	NewOwner   string   `json:"new_owner"`
}

type GroupChatTransferResponse struct {
	internal.BizResponse
	FailedChatList []struct {
		ChatId  string `json:"chat_id"`
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	} `json:"failed_chat_list"`
}

// TransferGroupChat 分配离职成员的客户群
// 可通过此接口，将已离职成员为群主的群，分配给另一个客服成员
// 群主离职了的客户群，才可继承
// 继承给的新群主，必须是配置了客户联系功能的成员
// 继承给的新群主，必须有设置实名
// 继承给的新群主，必须有激活企业微信
// 同一个人的群，限制每天最多分配300个给新群主
// 参考连接　https://open.work.weixin.qq.com/api/doc/90001/90143/93242
func (ww *weWork) TransferGroupChat(corpId uint, request GroupChatTransferRequest) (resp GroupChatTransferResponse) {
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/externalcontact/groupchat/transfer")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
