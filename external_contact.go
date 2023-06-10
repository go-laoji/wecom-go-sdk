package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type ExternalContactGetFollowUserListResponse struct {
	internal.BizResponse
	FollowUser []string `json:"follow_user"`
}

// ExternalContactGetFollowUserList 获取配置了客户联系功能的成员列表
// 参考连接　https://open.work.weixin.qq.com/api/doc/90001/90143/92576
func (ww *weWork) ExternalContactGetFollowUserList(corpId uint) (resp ExternalContactGetFollowUserListResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		Get("/cgi-bin/externalcontact/get_follow_user_list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type ExternalContactListResponse struct {
	internal.BizResponse
	ExternalUserId []string `json:"external_userid"`
}

// ExternalContactList 获取客户列表
// 参考连接　https://open.work.weixin.qq.com/api/doc/90001/90143/92264
func (ww *weWork) ExternalContactList(corpId uint, userId string) (resp ExternalContactListResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).SetQueryParam("userid", userId).
		Get("/cgi-bin/externalcontact/list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type ExternalContact struct {
	ExternalUserId  string `json:"external_userid"`
	Name            string `json:"name"`
	Position        string `json:"position"`
	Avatar          string `json:"avatar"`
	CorpName        string `json:"corp_name"`
	CorpFullName    string `json:"corp_full_name"`
	Type            int    `json:"type"`
	Gender          int    `json:"gender"`
	UnionId         string `json:"unionid"`
	ExternalProfile struct {
		ExternalAttr []struct {
			Type int    `json:"type"`
			Name string `json:"name"`
			Text struct {
				Value string `json:"value"`
			} `json:"text,omitempty"`
			Web struct {
				Url   string `json:"url"`
				Title string `json:"title"`
			} `json:"web,omitempty"`
			MiniProgram struct {
				AppId    string `json:"appid"`
				PagePath string `json:"pagepath"`
				Title    string `json:"title"`
			} `json:"miniprogram,omitempty"`
		}
	} `json:"external_profile"`
}

type FollowUser struct {
	UserId      string `json:"userid"`
	Remark      string `json:"remark,omitempty"`
	Description string `json:"description,omitempty"`
	CreateTime  int64  `json:"createtime"`
	Tags        []struct {
		GroupName string `json:"group_name"`
		TagName   string `json:"tag_name"`
		TagId     string `json:"tag_id"`
		Type      int    `json:"type"`
	} `json:"tags,omitempty"`
	RemarkCorpName string   `json:"remark_corp_name,omitempty"`
	RemarkMobiles  []string `json:"remark_mobiles,omitempty"`
	State          string   `json:"state,omitempty"`
	OperUserId     string   `json:"oper_userid,omitempty"`
	AddWay         int      `json:"add_way,omitempty"`
	WechatChannels struct {
		Nickname string `json:"nickname"`
		Source   int    `json:"source"`
	} `json:"wechat_channels,omitempty"`
}

type ExternalContactGetResponse struct {
	internal.BizResponse
	ExternalContact ExternalContact `json:"external_contact"`
	FollowUser      []FollowUser    `json:"follow_user"`
	NextCursor      string          `json:"next_cursor"`
}

// ExternalContactGet 获取客户详情
// 参考连接　https://open.work.weixin.qq.com/api/doc/90001/90143/92265
// 当客户在企业内的跟进人超过500人时需要使用cursor参数进行分页获取
func (ww *weWork) ExternalContactGet(corpId uint, externalUserId, cursor string) (resp ExternalContactGetResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		SetQueryParam("external_userid", externalUserId).
		SetQueryParam("cursor", cursor).
		Get("/cgi-bin/externalcontact/get")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type ExternalContactBatchGetByUserResponse struct {
	internal.BizResponse
	ExternalContactList []struct {
		ExternalContact ExternalContact `json:"external_contact"`
		FollowInfo      FollowUser      `json:"follow_info"`
	} `json:"external_contact_list"`
	NextCursor string `json:"next_cursor"`
}

// ExternalContactBatchGetByUser 批量获取客户详情
// 企业可通过此接口获取指定成员添加的客户信息列表。
// 参考连接 https://open.work.weixin.qq.com/api/doc/90001/90143/93010
func (ww *weWork) ExternalContactBatchGetByUser(corpId uint, userIds []string, cursor string, limit int) (resp ExternalContactBatchGetByUserResponse) {
	p := H{"userid_list": userIds, "cursor": cursor, "limit": limit}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/externalcontact/batch/get_by_user")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type ExternalContactRemarkRequest struct {
	UserId           string   `json:"user_id" validate:"required"`
	ExternalUserid   string   `json:"external_userid" validate:"required"`
	Remark           string   `json:"remark"`
	Description      string   `json:"description"`
	RemarkCompany    string   `json:"remark_company"`
	RemarkMobiles    []string `json:"remark_mobiles"`
	RemarkPicMediaId string   `json:"remark_pic_mediaid"`
}

// ExternalContactRemark 修改客户备注信息
// 参考连接 https://open.work.weixin.qq.com/api/doc/90001/90143/92694
func (ww *weWork) ExternalContactRemark(corpId uint, remark ExternalContactRemarkRequest) (resp internal.BizResponse) {
	if ok := validate.Struct(remark); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(remark).SetResult(&resp).
		Post("/cgi-bin/externalcontact/remark")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type UnionId2ExternalUserIdResponse struct {
	internal.BizResponse
	ExternalUserId string `json:"external_userid"`
}

// UnionId2ExternalUserId 外部联系人unionid转换
// https://open.work.weixin.qq.com/api/doc/90001/90143/93274
func (ww *weWork) UnionId2ExternalUserId(corpId uint, unionid, openid string) (resp UnionId2ExternalUserIdResponse) {
	p := H{"unionid": unionid, "openid": openid}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/externalcontact/unionid_to_external_userid")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type ToServiceExternalUseridResponse struct {
	internal.BizResponse
}

// ToServiceExternalUserid 代开发应用external_userid转换
// https://open.work.weixin.qq.com/api/doc/90001/90143/95195
func (ww *weWork) ToServiceExternalUserid(corpId uint, externalUserId string) (resp ToServiceExternalUseridResponse) {
	p := H{"external_userid": externalUserId}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/externalcontact/to_service_external_userid")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
