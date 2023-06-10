package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

// SendWelcomeMsg 发送新客户欢迎语
// https://open.work.weixin.qq.com/api/doc/90001/90143/92599
func (ww *weWork) SendWelcomeMsg(corpId uint, msg ExternalMsg) (resp internal.BizResponse) {
	if ok := validate.Struct(msg); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(msg).SetResult(&resp).
		Post("/cgi-bin/externalcontact/send_welcome_msg")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
