package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type ConclusionsText struct {
	Content string `json:"content"`
}

type ConclusionsImage struct {
	MediaId string `json:"media_id"`
	PicUrl  string `json:"pic_url"`
}

type ConclusionsLink struct {
	Title  string `json:"title"`
	PicUrl string `json:"pic_url"`
	Desc   string `json:"desc"`
	Url    string `json:"url"`
}

type ConclusionsMiniProgram struct {
	Title      string `json:"title"`
	PicMediaId string `json:"pic_media_id"`
	AppId      string `json:"appid"`
	Page       string `json:"page"`
}

type ContactMe struct {
	ConfigId      string   `json:"config_id,omitempty"`
	Type          int      `json:"type" validate:"required,oneof=1 2"`
	Scene         int      `json:"scene" validate:"required,oneof=1 2"`
	Style         int      `json:"style"`
	Remark        string   `json:"remark"`
	SkipVerify    bool     `json:"skip_verify"`
	State         string   `json:"state"`
	User          []string `json:"user"`
	Party         []int32  `json:"party"`
	IsTemp        bool     `json:"is_temp"`
	ExpiresIn     int32    `json:"expires_in"`
	ChatExpiresIn int32    `json:"chat_expires_in"`
	UnionId       string   `json:"unionid"`
	Conclusions   struct {
		*ConclusionsText        `json:"text,omitempty"`
		*ConclusionsImage       `json:"image,omitempty"`
		*ConclusionsLink        `json:"link,omitempty"`
		*ConclusionsMiniProgram `json:"miniprogram,omitempty"`
	} `json:"conclusions"`
}

type ContactMeAddResponse struct {
	internal.BizResponse
	ConfigId string `json:"config_id"`
	QrCode   string `json:"qr_code"`
}

// ExternalAddContactWay 配置客户联系「联系我」方式
// https://open.work.weixin.qq.com/api/doc/90001/90143/92577#%E9%85%8D%E7%BD%AE%E5%AE%A2%E6%88%B7%E8%81%94%E7%B3%BB%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F
func (ww *weWork) ExternalAddContactWay(corpId uint, me ContactMe) (resp ContactMeAddResponse) {
	if ok := validate.Struct(me); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(me).SetResult(&resp).
		Post("/cgi-bin/externalcontact/add_contact_way")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// ExternalUpdateContactWay 更新企业已配置的「联系我」方式
// https://open.work.weixin.qq.com/api/doc/90001/90143/92577#%E6%9B%B4%E6%96%B0%E4%BC%81%E4%B8%9A%E5%B7%B2%E9%85%8D%E7%BD%AE%E7%9A%84%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F
func (ww *weWork) ExternalUpdateContactWay(corpId uint, me ContactMe) (resp internal.BizResponse) {
	if ok := validate.Struct(me); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(me).SetResult(&resp).
		Post("/cgi-bin/externalcontact/update_contact_way")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type ContactMeGetResponse struct {
	internal.BizResponse
	ContactWay struct {
		ConfigId string `json:"config_id"`
		ContactMe
	} `json:"contact_way"`
}

// ExternalGetContactWay 获取企业已配置的「联系我」方式
// https://open.work.weixin.qq.com/api/doc/90001/90143/92577#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E5%B7%B2%E9%85%8D%E7%BD%AE%E7%9A%84%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F
func (ww *weWork) ExternalGetContactWay(corpId uint, configId string) (resp ContactMeGetResponse) {
	p := H{"config_id": configId}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/externalcontact/get_contact_way")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type ContactMeListResponse struct {
	internal.BizResponse
	ContactWay []struct {
		ConfigId string `json:"config_id"`
	} `json:"contact_way"`
	NextCursor string `json:"next_cursor"`
}

// ExternalListContactWay 获取企业配置的「联系我」二维码和「联系我」小程序插件列表。不包含临时会话。
// 注意，该接口仅可获取2021年7月10日以后创建的「联系我」
// https://open.work.weixin.qq.com/api/doc/90001/90143/92577#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E5%B7%B2%E9%85%8D%E7%BD%AE%E7%9A%84%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E5%88%97%E8%A1%A8
func (ww *weWork) ExternalListContactWay(corpId uint, startTime, endTime int64, cursor string, limit int) (resp ContactMeListResponse) {
	p := H{"start_time": startTime, "end_time": endTime, "cursor": cursor, "limit": limit}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/externalcontact/list_contact_way")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// ExternalDeleteContactWay 删除企业已配置的「联系我」方式
// https://open.work.weixin.qq.com/api/doc/90001/90143/92577#%E5%88%A0%E9%99%A4%E4%BC%81%E4%B8%9A%E5%B7%B2%E9%85%8D%E7%BD%AE%E7%9A%84%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F
func (ww *weWork) ExternalDeleteContactWay(corpId uint, configId string) (resp internal.BizResponse) {
	p := H{"config_id": configId}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/externalcontact/del_contact_way")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// ExternalCloseTempChat 结束临时会话
// https://developer.work.weixin.qq.com/document/path/95724#%E7%BB%93%E6%9D%9F%E4%B8%B4%E6%97%B6%E4%BC%9A%E8%AF%9D
func (ww *weWork) ExternalCloseTempChat(corpId uint, userId, externalUserId string) (resp internal.BizResponse) {
	p := H{"userid": userId, "external_userid": externalUserId}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/externalcontact/close_temp_chat")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
