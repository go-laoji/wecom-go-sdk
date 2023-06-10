package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type InvoiceInfoQuery struct {
	CardId      string `json:"card_id" validate:"required"`
	EncryptCode string `json:"encrypt_code" validate:"required"`
}

type GetInvoiceInfoResponse struct {
	internal.BizResponse
	CardID    string `json:"card_id"`
	BeginTime int    `json:"begin_time"`
	EndTime   int    `json:"end_time"`
	Openid    string `json:"openid"`
	Type      string `json:"type"`
	Payee     string `json:"payee"`
	Detail    string `json:"detail"`
	UserInfo  struct {
		Fee         int    `json:"fee"`
		Title       string `json:"title"`
		BillingTime int    `json:"billing_time"`
		BillingNo   string `json:"billing_no"`
		BillingCode string `json:"billing_code"`
		Info        []struct {
			Name  string `json:"name"`
			Num   int    `json:"num"`
			Unit  string `json:"unit"`
			Fee   int    `json:"fee"`
			Price int    `json:"price"`
		} `json:"info"`
		FeeWithoutTax   int    `json:"fee_without_tax"`
		Tax             int    `json:"tax"`
		Detail          string `json:"detail"`
		PdfURL          string `json:"pdf_url"`
		ReimburseStatus string `json:"reimburse_status"`
		CheckCode       string `json:"check_code"`
	} `json:"user_info"`
}

// GetInvoiceInfo 查询电子发票
// https://open.work.weixin.qq.com/api/doc/90001/90143/90420
func (ww *weWork) GetInvoiceInfo(corpId uint, query InvoiceInfoQuery) (resp GetInvoiceInfoResponse) {
	if ok := validate.Struct(query); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(query).SetResult(&resp).
		Post("/cgi-bin/card/invoice/reimburse/getinvoiceinfo")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

const (
	ReimburseStatusInit    = "INVOICE_REIMBURSE_INIT"
	ReimburseStatusLock    = "INVOICE_REIMBURSE_LOCK"
	ReimburseStatusClosure = "INVOICE_REIMBURSE_CLOSURE"
)

type UpdateInvoiceStatusRequest struct {
	CardId          string `json:"card_id" validate:"required"`
	EncryptCode     string `json:"encrypt_code" validate:"required"`
	ReimburseStatus string `json:"reimburse_status" validate:"required,oneof=INVOICE_REIMBURSE_INIT INVOICE_REIMBURSE_LOCK INVOICE_REIMBURSE_CLOSURE"`
}

// UpdateInvoiceStatus 更新发票状态
// https://open.work.weixin.qq.com/api/doc/90001/90143/90421
func (ww *weWork) UpdateInvoiceStatus(corpId uint, request UpdateInvoiceStatusRequest) (resp internal.BizResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/card/invoice/reimburse/updateinvoicestatus")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type UpdateInvoiceStatusBatchRequest struct {
	OpenId          string             `json:"openid" validate:"required"`
	ReimburseStatus string             `json:"reimburse_status" validate:"required,oneof=INVOICE_REIMBURSE_INIT INVOICE_REIMBURSE_LOCK INVOICE_REIMBURSE_CLOSURE"`
	InvoiceList     []InvoiceInfoQuery `json:"invoice_list" validate:"required"`
}

// UpdateInvoiceStatusBatch 批量更新发票状态
// https://open.work.weixin.qq.com/api/doc/90001/90143/90422
func (ww *weWork) UpdateInvoiceStatusBatch(corpId uint, request UpdateInvoiceStatusBatchRequest) (resp internal.BizResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/card/invoice/reimburse/updatestatusbatch")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type InvoiceInfoQueryBatch struct {
	ItemList []InvoiceInfoQuery `json:"item_list" validate:"required"`
}

type GetInvoiceInfoBatchResponse struct {
	internal.BizResponse
	ItemList []GetInvoiceInfoResponse `json:"item_list"`
}

// GetInvoiceInfoBatch 批量查询电子发票
// https://open.work.weixin.qq.com/api/doc/90001/90143/90423
func (ww *weWork) GetInvoiceInfoBatch(corpId uint, query InvoiceInfoQueryBatch) (resp GetInvoiceInfoBatchResponse) {
	if ok := validate.Struct(query); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(query).SetResult(&resp).
		Post("/cgi-bin/card/invoice/reimburse/getinvoiceinfobatch")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
