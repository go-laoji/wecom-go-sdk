package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type GetPaymentResultResponse struct {
	internal.BizResponse
	ProjectName   string `json:"project_name"`
	Amount        int    `json:"amount"`
	PaymentResult []struct {
		StudentUserid     string `json:"student_userid"`
		TradeState        int    `json:"trade_state"`
		TradeNo           string `json:"trade_no"`
		PayerParentUserid string `json:"payer_parent_userid"`
	} `json:"payment_result"`
}

// GetPaymentResult 获取学生付款结果
// https://open.work.weixin.qq.com/api/doc/90001/90143/94553
func (ww *weWork) GetPaymentResult(corpId uint, paymentId string) (resp GetPaymentResultResponse) {
	h := H{}
	h["payment_id"] = paymentId
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/school/get_payment_result")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetTradeRequest struct {
	PaymentId string `json:"payment_id" validate:"required"`
	TradeNo   string `json:"trade_no" validate:"required"`
}
type GetTradeResponse struct {
	internal.BizResponse
	TransactionId string `json:"transaction_id"`
	PayTime       int    `json:"pay_time"`
}

// GetTrade 获取订单详情
// https://open.work.weixin.qq.com/api/doc/90001/90143/94554
func (ww *weWork) GetTrade(corpId uint, request GetTradeRequest) (resp GetTradeResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/school/get_trade")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
