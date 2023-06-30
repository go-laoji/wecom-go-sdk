package wework

import "github.com/go-laoji/wecom-go-sdk/v2/internal"

type GetBillListRequest struct {
	BeginTime   int64  `json:"begin_time" validate:"required"` // 单位/unit: s
	EndTime     int64  `json:"end_time" validate:"required"`   // 单位/unit: s
	PayeeUserId string `json:"payee_userid,omitempty"`
	Cursor      string `json:"cursor,omitempty"`
	Limit       int    `json:"limit,omitempty"`
}

type GetBillListResponse struct {
	internal.BizResponse
	NextCursor string `json:"next_cursor"`
	BillList   []struct {
		TransactionId  string `json:"transaction_id"`
		BillType       int    `json:"bill_type"`
		TradeState     int    `json:"trade_state"`
		PayTime        int64  `json:"pay_time"`
		OutTradeNo     string `json:"out_trade_no"`
		OutRefundNo    string `json:"out_refund_no"`
		ExternalUserId string `json:"external_userid"`
		TotalFee       int    `json:"total_fee"` // 单位/unit: cent
		PayeeUserId    string `json:"payee_userid"`
		PaymentType    int    `json:"payment_type"`
		MchId          string `json:"mch_id"`
		Remark         string `json:"remark"`
		CommodityList  []struct {
			Description string `json:"description"`
			Amount      int    `json:"amount"`
		} `json:"commodity_list"`
		TotalRefundFee int `json:"total_refund_fee"`
		RefundList     []struct {
			OutRefundNo   string `json:"out_refund_no"`
			RefundUserId  string `json:"refund_userid"`
			RefundComment string `json:"refund_comment"`
			RefundReqTime int64  `json:"refund_reqtime"`
			RefundStatus  int    `json:"refund_status"`
			RefundFee     int    `json:"refund_fee"`
		} `json:"refund_list"`
		ContactInfo struct {
			Name    string `json:"name"`
			Phone   string `json:"phone"`
			Address string `json:"address"`
		} `json:"contact_info"`
		MiniProgramInfo struct {
			AppId string `json:"appid"`
			Name  string `json:"name"`
		} `json:"miniprogram_info"`
	} `json:"bill_list"`
}

// GetBillList 获取对外收款记录
// 参考连接　https://developer.work.weixin.qq.com/document/path/93667
func (ww *weWork) GetBillList(corpId uint, req GetBillListRequest) (resp GetBillListResponse) {
	if ok := validate.Struct(req); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(req).SetResult(&resp).
		Post("/cgi-bin/externalpay/get_bill_list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
