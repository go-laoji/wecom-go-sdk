package wework

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wework/internal"
)

// SendWelcomeMsg 发送新客户欢迎语
// https://open.work.weixin.qq.com/api/doc/90001/90143/92599
func (ww weWork) SendWelcomeMsg(corpId uint, msg ExternalMsg) (resp internal.BizResponse) {
	if ok := validate.Struct(msg); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/send_welcome_msg?%s", queryParams.Encode()), msg)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
