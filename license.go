package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type CreateOrderRequest struct {
	CorpId       string `json:"corpid" validate:"required"`
	BuyerUserid  string `json:"buyer_userid" validate:"required"`
	AccountCount struct {
		BaseCount            int `json:"base_count" validate:"required_without=ExternalContactCount,max=1000000"`
		ExternalContactCount int `json:"external_contact_count" validate:"required_without=BaseCount,max=1000000"`
	} `json:"account_count" validate:"required"`
	AccountDuration struct {
		Months int `json:"months"`
	} `json:"account_duration"`
}
type OrderResponse struct {
	internal.BizResponse
	OrderId string `json:"order_id"`
}

// CreateNewOrder 下单购买帐号
// https://developer.work.weixin.qq.com/document/path/95644
func (ww *weWork) CreateNewOrder(request CreateOrderRequest) (resp OrderResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getProviderRequest().SetBody(request).SetResult(&resp).
		Post("/cgi-bin/license/create_new_order")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type CreateReNewOrderJobRequest struct {
	CorpId      string `json:"corpid" validate:"required"`
	AccountList []struct {
		Userid string `json:"userid" validate:"required"`
		Type   int    `json:"type" validate:"required,oneof=1 2"`
	} `json:"account_list" validate:"required"`
	JobId string `json:"jobid"`
}

type CreateReNewOrderJobResponse struct {
	internal.BizResponse
	Jobid              string `json:"jobid"`
	InvalidAccountList []struct {
		internal.BizResponse
		Userid string `json:"userid"`
		Type   int    `json:"type"`
	} `json:"invalid_account_list"`
}

// CreateReNewOrderJob 下单续期帐号/创建续期任务
// https://developer.work.weixin.qq.com/document/path/95646
func (ww *weWork) CreateReNewOrderJob(request CreateReNewOrderJobRequest) (resp CreateReNewOrderJobResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getProviderRequest().SetBody(request).SetResult(&resp).
		Post("/cgi-bin/license/create_renew_order_job")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type SubmitOrderJobRequest struct {
	Jobid           string `json:"jobid" validate:"required"`
	BuyerUserid     string `json:"buyer_userid" validate:"required"`
	AccountDuration struct {
		Months int `json:"months" validate:"required"`
	} `json:"account_duration" validate:"required"`
}

// SubmitOrderJob 下单续期帐号/提交续期订单
func (ww *weWork) SubmitOrderJob(request SubmitOrderJobRequest) (resp OrderResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getProviderRequest().SetBody(request).SetResult(&resp).
		Post("/cgi-bin/license/submit_order_job")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type ListOrderRequest struct {
	Corpid    string `json:"corpid"`
	StartTime int    `json:"start_time" validate:"required_with=EndTime"`
	EndTime   int    `json:"end_time" validate:"required_with=StartTime"`
	Cursor    string `json:"cursor"`
	Limit     int    `json:"limit" validate:"max=1000"`
}
type ListOrderResponse struct {
	internal.BizResponse
	NextCursor string `json:"next_cursor"`
	HasMore    int    `json:"has_more"`
	OrderList  []struct {
		OrderID   string `json:"order_id"`
		OrderType int    `json:"order_type"`
	} `json:"order_list"`
}

// ListOrder 获取订单列表
// https://developer.work.weixin.qq.com/document/path/95647
func (ww *weWork) ListOrder(request ListOrderRequest) (resp ListOrderResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getProviderRequest().SetBody(request).SetResult(&resp).
		Post("/cgi-bin/license/list_order")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetOrderRequest struct {
	OrderId string `json:"order_id" validate:"required"`
}

type GetOrderResponse struct {
	internal.BizResponse
	Order struct {
		OrderID      string `json:"order_id"`
		OrderType    int    `json:"order_type"`
		OrderStatus  int    `json:"order_status"`
		Corpid       string `json:"corpid"`
		Price        int    `json:"price"`
		AccountCount struct {
			BaseCount            int `json:"base_count"`
			ExternalContactCount int `json:"external_contact_count"`
		} `json:"account_count"`
		AccountDuration struct {
			Months int `json:"months"`
		} `json:"account_duration"`
		CreateTime int `json:"create_time"`
		PayTime    int `json:"pay_time"`
	} `json:"order"`
}

// GetOrder 获取订单详情
// https://developer.work.weixin.qq.com/document/path/95648
func (ww *weWork) GetOrder(request GetOrderRequest) (resp GetOrderResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getProviderRequest().SetBody(request).SetResult(&resp).
		Post("/cgi-bin/license/get_order")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type ListOrderAccountRequest struct {
	OrderID string `json:"order_id" validate:"required"`
	Limit   int    `json:"limit,omitempty" validate:"max=1000"`
	Cursor  string `json:"cursor,omitempty"`
}

type ListOrderAccountResponse struct {
	internal.BizResponse
	NextCursor  string `json:"next_cursor"`
	HasMore     int    `json:"has_more"`
	AccountList []struct {
		ActiveCode string `json:"active_code"`
		Userid     string `json:"userid"`
		Type       int    `json:"type"`
	} `json:"account_list"`
}

// ListOrderAccount 获取订单中的帐号列表
// https://developer.work.weixin.qq.com/document/path/95649
func (ww *weWork) ListOrderAccount(request ListOrderAccountRequest) (resp ListOrderAccountResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getProviderRequest().SetBody(request).SetResult(&resp).
		Post("/cgi-bin/license/list_order_account")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type ActiveAccountRequest struct {
	ActiveCode string `json:"active_code" validate:"required"`
	CorpId     string `json:"corpid" validate:"required"`
	Userid     string `json:"userid" validate:"required"`
}

// ActiveAccount 激活帐号
// https://developer.work.weixin.qq.com/document/path/95553#%E6%BF%80%E6%B4%BB%E5%B8%90%E5%8F%B7
func (ww *weWork) ActiveAccount(request ActiveAccountRequest) (resp internal.BizResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getProviderRequest().SetBody(request).SetResult(&resp).
		Post("/cgi-bin/license/active_account")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type BatchActiveAccountRequest struct {
	CorpId     string `json:"corpid"`
	ActiveList []struct {
		ActiveCode string `json:"active_code"`
		Userid     string `json:"userid"`
	} `json:"active_list"`
}

type BatchActiveAccountResponse struct {
	internal.BizResponse
	ActiveResult []struct {
		ActiveCode string `json:"active_code"`
		Userid     string `json:"userid"`
		ErrCode    int    `json:"errcode"`
	} `json:"active_result"`
}

// BatchActiveAccount 批量激活帐号
// https://developer.work.weixin.qq.com/document/path/95553#%E6%89%B9%E9%87%8F%E6%BF%80%E6%B4%BB%E5%B8%90%E5%8F%B7
func (ww *weWork) BatchActiveAccount(request BatchActiveAccountRequest) (resp BatchActiveAccountResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getProviderRequest().SetBody(request).SetResult(&resp).
		Post("/cgi-bin/license/batch_active_account")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetActiveInfoByCodeRequest struct {
	CorpId     string `json:"corpid" validate:"required"`
	ActiveCode string `json:"active_code" validate:"required"`
}

type ActiveInfo struct {
	ActiveCode string `json:"active_code"`
	Type       int    `json:"type"`
	Status     int    `json:"status"`
	Userid     string `json:"userid"`
	CreateTime int    `json:"create_time"`
	ActiveTime int    `json:"active_time "`
	ExpireTime int    `json:"expire_time"`
}
type GetActiveInfoByCodeResponse struct {
	internal.BizResponse
	ActiveInfo ActiveInfo `json:"active_info"`
}

// GetActiveInfoByCode 获取激活码详情
// https://developer.work.weixin.qq.com/document/path/95552#%E8%8E%B7%E5%8F%96%E6%BF%80%E6%B4%BB%E7%A0%81%E8%AF%A6%E6%83%85
func (ww *weWork) GetActiveInfoByCode(request GetActiveInfoByCodeRequest) (resp GetActiveInfoByCodeResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getProviderRequest().SetBody(request).SetResult(&resp).
		Post("/cgi-bin/license/get_active_info_by_code")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type BatchGetActiveInfoByCodeRequest struct {
	CorpId         string   `json:"corpid" validate:"required"`
	ActiveCodeList []string `json:"active_code_list" validate:"required,max=1000"`
}
type BatchGetActiveInfoByCodeResponse struct {
	internal.BizResponse
	ActiveInfoList        []ActiveInfo `json:"active_info_list"`
	InvalidActiveCodeList []string     `json:"invalid_active_code_list"`
}

// BatchGetActiveInfoByCode 批量获取激活码详情
func (ww *weWork) BatchGetActiveInfoByCode(request BatchGetActiveInfoByCodeRequest) (resp BatchGetActiveInfoByCodeResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getProviderRequest().SetBody(request).SetResult(&resp).
		Post("/cgi-bin/license/batch_get_active_info_by_code")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type ListActivedAccountRequest struct {
	CorpId string `json:"corpid" validate:"required"`
	Limit  int    `json:"limit,omitempty" validate:"omitempty,max=1000"`
	Cursor string `json:"cursor,omitempty"`
}

type ListActivedAccountResponse struct {
	internal.BizResponse
	NextCursor  string       `json:"next_cursor"`
	HasMore     int          `json:"has_more"`
	AccountList []ActiveInfo `json:"account_list"`
}

// ListActivedAccount 获取企业的帐号列表
// https://developer.work.weixin.qq.com/document/path/95544
func (ww *weWork) ListActivedAccount(request ListActivedAccountRequest) (resp ListActivedAccountResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getProviderRequest().SetBody(request).SetResult(&resp).
		Post("/cgi-bin/license/list_actived_account")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetActiveInfoByUserRequest struct {
	CorpId string `json:"corpid" validate:"required"`
	UserId string `json:"userid" validate:"required"`
}
type GetActiveInfoByUserResponse struct {
	internal.BizResponse
	ActiveStatus   int `json:"active_status"`
	ActiveInfoList []struct {
		ActiveCode string `json:"active_code"`
		Type       int    `json:"type"`
		Userid     string `json:"userid"`
		ActiveTime int    `json:"active_time"`
		ExpireTime int    `json:"expire_time"`
	} `json:"active_info_list"`
}

// GetActiveInfoByUser 获取成员的激活详情
// https://developer.work.weixin.qq.com/document/path/95555
func (ww *weWork) GetActiveInfoByUser(request GetActiveInfoByUserRequest) (resp GetActiveInfoByUserResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getProviderRequest().SetBody(request).SetResult(&resp).
		Post("/cgi-bin/license/get_active_info_by_user")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type BatchTransferLicenseRequest struct {
	CorpId       string `json:"corpid" validate:"required"`
	TransferList []struct {
		HandoverUserid string `json:"handover_userid" validate:"required"`
		TakeoverUserid string `json:"takeover_userid" validate:"required"`
	} `json:"transfer_list" validate:"required"`
}

type BatchTransferLicenseResponse struct {
	internal.BizResponse
	TransferResult []struct {
		HandoverUserid string `json:"handover_userid"`
		TakeoverUserid string `json:"takeover_userid"`
		ErrCode        int    `json:"errcode"`
	} `json:"transfer_result"`
}

// BatchTransferLicense 帐号继承
// https://developer.work.weixin.qq.com/document/path/95673
func (ww *weWork) BatchTransferLicense(request BatchTransferLicenseRequest) (resp BatchTransferLicenseResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getProviderRequest().SetBody(request).SetResult(&resp).
		Post("/cgi-bin/license/batch_transfer_license")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type SetAutoActiveStatusRequest struct {
	CorpId           string `json:"corpid" validate:"required"`
	AutoActiveStatus uint   `json:"auto_active_status" validate:"required,oneof=0 1"`
}

func (ww *weWork) SetAutoActiveStatus(request SetAutoActiveStatusRequest) (resp internal.BizResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getProviderRequest().SetBody(request).SetResult(&resp).
		Post("/cgi-bin/license/set_auto_active_status")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetAutoActiveStatusResponse struct {
	internal.BizResponse
	AutoActiveStatus uint `json:"auto_active_status"`
}

func (ww *weWork) GetAutoActiveStatus(corpid string) (resp GetAutoActiveStatusResponse) {
	if len(corpid) == 0 {
		resp.ErrCode = 500
		resp.ErrorMsg = "corpid 参数不能为空"
		return
	}
	request := H{"corpid": corpid}
	_, err := ww.getProviderRequest().SetBody(request).SetResult(&resp).
		Post("/cgi-bin/license/get_auto_active_status")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
