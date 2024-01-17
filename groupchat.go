package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type GroupChatListFilter struct {
	StatusFilter int `json:"status_filter,omitempty" validate:"omitempty,oneof=0 1 2 3"`
	OwnerFilter  struct {
		UserIdList []string `json:"userid_list"`
	} `json:"owner_filter,omitempty"`
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit" validate:"required,min=1,max=1000"`
}

type GroupChatListResponse struct {
	internal.BizResponse
	GroupChatList []struct {
		ChatId string `json:"chat_id"`
		Status int    `json:"status"`
	} `json:"group_chat_list"`
	NextCursor string `json:"next_cursor"`
}

// GroupChatList 获取客户群列表
// 该接口用于获取配置过客户群管理的客户群列表
// 参考连接　https://open.work.weixin.qq.com/api/doc/90001/90143/93414
func (ww *weWork) GroupChatList(corpId uint, filter GroupChatListFilter) (resp GroupChatListResponse) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(filter).SetResult(&resp).
		Post("/cgi-bin/externalcontact/groupchat/list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GroupChatRequest struct {
	ChatId   string `json:"chat_id" validate:"required"`
	NeedName int    `json:"need_name" validate:"omitempty,oneof=0 1"`
}

type GroupChatResponse struct {
	internal.BizResponse
	GroupChat struct {
		ChatID     string `json:"chat_id"`
		Name       string `json:"name"`
		Owner      string `json:"owner"`
		CreateTime int    `json:"create_time"`
		Notice     string `json:"notice"`
		MemberList []struct {
			Userid    string `json:"userid"`
			Type      int    `json:"type"`
			JoinTime  int    `json:"join_time"`
			JoinScene int    `json:"join_scene"`
			Invitor   struct {
				Userid string `json:"userid"`
			} `json:"invitor,omitempty"`
			GroupNickname string `json:"group_nickname"`
			Name          string `json:"name"`
			Unionid       string `json:"unionid,omitempty"`
		} `json:"member_list"`
		AdminList []struct {
			Userid string `json:"userid"`
		} `json:"admin_list"`
		MemberVersion string `json:"member_version"`
	} `json:"group_chat"`
}

// GroupChat 获取客户群详情
// 通过客户群ID，获取详情。包括群名、群成员列表、群成员入群时间、入群方式。（客户群是由具有客户群使用权限的成员创建的外部群）
// 需注意的是，如果发生群信息变动，会立即收到群变更事件，但是部分信息是异步处理，可能需要等一段时间调此接口才能得到最新结果
// 参考连接　https://open.work.weixin.qq.com/api/doc/90001/90143/92707
func (ww *weWork) GroupChat(corpId uint, request GroupChatRequest) (resp GroupChatResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/externalcontact/groupchat/get")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GroupOpengId2ChatIdResponse struct {
	internal.BizResponse
	ChatId string `json:"chat_id"`
}

// GroupOpengId2ChatId 客户群opengid转换
// 用户在微信里的客户群里打开小程序时，某些场景下可以获取到群的opengid，如果该群是企业微信的客户群，则企业或第三方可以调用此接口将一个opengid转换为客户群chat_id
// 参考连接　https://open.work.weixin.qq.com/api/doc/90001/90143/94828
func (ww *weWork) GroupOpengId2ChatId(corpId uint, opengid string) (resp GroupOpengId2ChatIdResponse) {
	p := H{"opengid": opengid}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/externalcontact/opengid_to_chatid")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
