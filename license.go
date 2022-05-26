package wework

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/internal"
	"net/url"
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
func (ww weWork) CreateNewOrder(request CreateOrderRequest) (resp OrderResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := url.Values{}
	queryParams.Add("provider_access_token", ww.getProviderToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/license/create_new_order?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
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
func (ww weWork) CreateReNewOrderJob(request CreateReNewOrderJobRequest) (resp CreateReNewOrderJobResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := url.Values{}
	queryParams.Add("provider_access_token", ww.getProviderToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/license/create_renew_order_job?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
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
func (ww weWork) SubmitOrderJob(request SubmitOrderJobRequest) (resp OrderResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := url.Values{}
	queryParams.Add("provider_access_token", ww.getProviderToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/license/submit_order_job?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
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
func (ww weWork) ListOrder(request ListOrderRequest) (resp ListOrderResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := url.Values{}
	queryParams.Add("provider_access_token", ww.getProviderToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/license/list_order?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
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
func (ww weWork) GetOrder(request GetOrderRequest) (resp GetOrderResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := url.Values{}
	queryParams.Add("provider_access_token", ww.getProviderToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/license/get_order?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
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
func (ww weWork) ListOrderAccount(request ListOrderAccountRequest) (resp ListOrderAccountResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := url.Values{}
	queryParams.Add("provider_access_token", ww.getProviderToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/license/list_order_account?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
