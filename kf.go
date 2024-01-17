package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type KfAccount struct {
	OpenKfId string `json:"open_kfid,omitempty"`
	Name     string `json:"name" validate:"required"`
	MediaId  string `json:"media_id" validate:"required"`
}
type KfAccountAddResponse struct {
	internal.BizResponse
	OpenKfId string `json:"open_kfid"`
}

func (ww *weWork) KfAccountAdd(corpId uint, account KfAccount) (resp KfAccountAddResponse) {
	if ok := validate.Struct(account); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(account).SetResult(&resp).
		Post("/cgi-bin/kf/account/add")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

func (ww *weWork) KfAccountDel(corpId uint, kfId string) (resp internal.BizResponse) {
	if kfId == "" {
		resp.ErrCode = 500
		resp.ErrorMsg = "客服ID必填"
		return
	}
	params := H{}
	params["open_kfid"] = kfId
	_, err := ww.getRequest(corpId).SetBody(params).SetResult(&resp).
		Post("/cgi-bin/kf/account/del")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

func (ww *weWork) KfAccountUpdate(corpId uint, account KfAccount) (resp internal.BizResponse) {
	if ok := validate.Struct(account); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(account).SetResult(&resp).
		Post("/cgi-bin/kf/account/update")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type KfAccountListRequest struct {
	Offset uint32 `json:"offset,omitempty"`
	Limit  uint32 `json:"limit,omitempty" validate:"max=100,min=1"`
}
type KfAccountListResponse struct {
	internal.BizResponse
	AccountList []struct {
		OpenKfId        string `json:"open_kfid"`
		Name            string `json:"name"`
		Avatar          string `json:"avatar"`
		ManagePrivilege bool   `json:"manage_privilege"`
	} `json:"account_list"`
}

func (ww *weWork) KfAccountList(corpId uint, request KfAccountListRequest) (resp KfAccountListResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/kf/account/list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type KfAccContactWayResponse struct {
	internal.BizResponse
	Url string `json:"url"`
}

func (ww *weWork) KfAddContactWay(corpId uint, kfId string, scene string) (resp KfAccContactWayResponse) {
	if kfId == "" {
		resp.ErrCode = 500
		resp.ErrorMsg = "客服ID必填"
		return
	}
	params := H{}
	params["open_kfid"] = kfId
	params["scene"] = scene
	_, err := ww.getRequest(corpId).SetBody(params).SetResult(&resp).
		Post("/cgi-bin/kf/add_contact_way")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type KfServicerRequest struct {
	OpenKfId         string   `json:"open_kfid" validate:"required"`
	UserIdList       []string `json:"userid_list" validate:"required_without=DepartmentIdList,max=100"`
	DepartmentIdList []uint32 `json:"department_id_list" validate:"required_without=UserIdList,max=100"`
}

type KfServicerResponse struct {
	internal.BizResponse
	ResultList []struct {
		UserId       string `json:"userid,omitempty"`
		DepartmentId uint32 `json:"department_id,omitempty"`
		internal.BizResponse
	} `json:"result_list"`
}

func (ww *weWork) KfServicerAdd(corpId uint, request KfServicerRequest) (resp KfServicerResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/kf/servicer/add")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

func (ww *weWork) KfServicerDel(corpId uint, request KfServicerRequest) (resp KfServicerResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/kf/servicer/del")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type KfServicerListResponse struct {
	internal.BizResponse
	ServicerList []struct {
		UserId       string `json:"userid,omitempty"`
		Status       uint   `json:"status,omitempty"`
		DepartmentId uint32 `json:"department_id,omitempty"`
	} `json:"servicer_list"`
}

func (ww *weWork) KfServicerList(corpId uint, kfId string) (resp KfServicerListResponse) {
	_, err := ww.getRequest(corpId).SetQueryParam("open_kf_id", kfId).SetResult(&resp).
		Get("/cgi-bin/kf/servicer/list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type KfServiceStateGetRequest struct {
	OpenKfId       string `json:"open_kfid" validate:"required"`
	ExternalUserId string `json:"external_userid" validate:"required"`
}

type KfServiceStateGetResponse struct {
	internal.BizResponse
	ServiceState   int    `json:"service_state"`
	ServicerUserId string `json:"servicer_userid"`
}

func (ww *weWork) KfServiceStateGet(corpId uint, request KfServiceStateGetRequest) (resp KfServiceStateGetResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/kf/service_state/get")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type KfServiceStateTransRequest struct {
	OpenKfId       string `json:"open_kfid" validate:"required"`
	ExternalUserId string `json:"external_userid" validate:"required"`
	ServiceState   int    `json:"service_state" validate:"required,oneof=0 1 2 3 4"`
	ServicerUserId string `json:"servicer_userid"`
}

type KfServiceStateTransResponse struct {
	internal.BizResponse
	MsgCode string `json:"msg_code"`
}

func (ww *weWork) KfServiceStateTrans(corpId uint, request KfServiceStateTransRequest) (resp KfServiceStateTransResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/kf/service_state/trans")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type KfSyncMsgRequest struct {
	Cursor      string `json:"cursor"`
	Token       string `json:"token"`
	Limit       int    `json:"limit"`
	VoiceFormat int    `json:"voice_format"`
	OpenKfId    string `json:"open_kfid"`
}

type KfMessage struct {
	MsgId               string                 `json:"msgid"`
	OpenKfId            string                 `json:"open_kfid"`
	ExternalUserId      string                 `json:"external_userid"`
	SendTime            int                    `json:"send_time"`
	Origin              int                    `json:"origin"`
	ServicerUserId      string                 `json:"servicer_userid"`
	MsgType             string                 `json:"msgtype"`
	Text                MsgText                `json:"text,omitempty"`
	Image               MsgImage               `json:"image,omitempty"`
	Voice               MsgVoice               `json:"voice,omitempty"`
	Video               MsgVideo               `json:"video,omitempty"`
	File                MsgFile                `json:"file,omitempty"`
	Location            MsgLocation            `json:"location,omitempty"`
	Link                MsgLink                `json:"link,omitempty"`
	BusinessCard        MsgBusinessCard        `json:"business_card,omitempty"`
	MiniProgram         MsgMiniProgram         `json:"miniprogram,omitempty"`
	MsgMenu             MsgMenu                `json:"msgmenu,omitempty"`
	ChannelsShopProduct MsgChannelsShopProduct `json:"channels_shop_product,omitempty"`
	ChannelsShopOrder   MsgChannelsShopOrder   `json:"channels_shop_order,omitempty"`
	Event               MsgEvent               `json:"event,omitempty"`
}

type KfSyncMsgResponse struct {
	internal.BizResponse
	NextCursor string      `json:"next_cursor"`
	HasMore    int         `json:"has_more"`
	MsgList    []KfMessage `json:"msg_list"`
}

type MsgText struct {
	Content string `json:"content"`
	MenuId  string `json:"menu_id"`
}
type MsgImage struct {
	MediaId string `json:"media_id"`
}
type MsgVoice struct {
	MediaId string `json:"media_id"`
}
type MsgVideo struct {
	MediaId string `json:"media_id"`
}
type MsgFile struct {
	MediaId string `json:"media_id"`
}
type MsgLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
}
type MsgLink struct {
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	Url          string `json:"url"`
	ThumbMediaId string `json:"thumb_media_id"`
}
type MsgBusinessCard struct {
	UserId string `json:"userid"`
}
type MsgMiniProgram struct {
	Title        string `json:"title"`
	AppId        string `json:"appid"`
	PagePath     string `json:"pagepath"`
	ThumbMediaId string `json:"thumb_media_id"`
}
type MenuItemClick struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type MenuItemView struct {
	URL     string `json:"url"`
	Content string `json:"content"`
}

type MenuItemMiniProgram struct {
	Appid    string `json:"appid"`
	PagePath string `json:"pagepath"`
	Content  string `json:"content"`
}

type MenuItemText struct {
	Content     string  `json:"content"`
	NoNewLine   int     `json:"no_newline,omitempty"`
	TailContent *string `json:"tail_content,omitempty"`
}

type MenuItem struct {
	Type        string               `json:"type"`
	Click       *MenuItemClick       `json:"click,omitempty"`
	View        *MenuItemView        `json:"view,omitempty"`
	MiniProgram *MenuItemMiniProgram `json:"miniprogram,omitempty"`
	Text        *MenuItemText        `json:"text,omitempty"`
}
type MsgMenu struct {
	HeadContent string     `json:"head_content"`
	List        []MenuItem `json:"list"`
	TailContent string     `json:"tail_content"`
}
type MsgCaLink struct {
	LinkUrl string `json:"link_url"`
}
type MsgChannelsShopProduct struct {
	ProductID    string `json:"product_id"`
	HeadImg      string `json:"head_img"`
	Title        string `json:"title"`
	SalesPrice   string `json:"sales_price"`
	ShopNickname string `json:"shop_nickname"`
	ShopHeadImg  string `json:"shop_head_img"`
}
type MsgChannelsShopOrder struct {
	OrderID       string `json:"order_id"`
	ProductTitles string `json:"product_titles"`
	PriceWording  string `json:"price_wording"`
	State         string `json:"state"`
	ImageURL      string `json:"image_url"`
	ShopNickname  string `json:"shop_nickname"`
}
type MsgEvent struct {
	EventType      string `json:"event_type"`
	OpenKfid       string `json:"open_kfid"`
	ExternalUserid string `json:"external_userid"`
	Scene          string `json:"scene"`
	SceneParam     string `json:"scene_param"`
	WelcomeCode    string `json:"welcome_code"`
	WechatChannels struct {
		Nickname     string `json:"nickname"`
		ShopNickName string `json:"shop_nickname"`
		Scene        uint32 `json:"scene"`
	} `json:"wechat_channels"`
}

func (ww *weWork) KfSyncMsg(corpId uint, request KfSyncMsgRequest) (resp KfSyncMsgResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/kf/sync_msg")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type SendMsgRequest struct {
	ToUser      string          `json:"touser" validate:"required"`
	OpenKfId    string          `json:"open_kfid" validate:"required"`
	MsgId       string          `json:"msgid,omitempty"`
	MsgType     string          `json:"msgtype" validate:"required"`
	Text        *MsgText        `json:"text,omitempty"`
	Image       *MsgImage       `json:"image,omitempty"`
	Voice       *MsgVoice       `json:"voice,omitempty"`
	Video       *MsgVideo       `json:"video,omitempty"`
	File        *MsgFile        `json:"file,omitempty"`
	Location    *MsgLocation    `json:"location,omitempty"`
	Link        *MsgLink        `json:"link,omitempty"`
	MiniProgram *MsgMiniProgram `json:"miniprogram,omitempty"`
	MsgMenu     *MsgMenu        `json:"msgmenu,omitempty"`
	CaLink      *MsgCaLink      `json:"ca_link,omitempty"`
}

type SendMsgResponse struct {
	internal.BizResponse
	MsgId string `json:"msgid"`
}

func (ww *weWork) KfSendMsg(corpId uint, request SendMsgRequest) (resp SendMsgResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/kf/send_msg")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type SendMsgOnEventRequest struct {
	Code    string   `json:"code" validate:"required"`
	MsgId   string   `json:"msgid,omitempty"`
	MsgType string   `json:"msgtype" validate:"required,oneof=text msgmenu"`
	Text    *MsgText `json:"text,omitempty"`
	MsgMenu *MsgMenu `json:"msgmenu,omitempty"`
}

func (ww *weWork) KfSendMsgOnEvent(corpId uint, request SendMsgOnEventRequest) (resp SendMsgResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/kf/send_msg_on_event")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type KfCustomerBatchGetResponse struct {
	internal.BizResponse
	CustomerList []struct {
		ExternalUserid      string `json:"external_userid"`
		Nickname            string `json:"nickname"`
		Avatar              string `json:"avatar"`
		Gender              int    `json:"gender"`
		Unionid             string `json:"unionid"`
		EnterSessionContext struct {
			Scene          string `json:"scene"`
			SceneParam     string `json:"scene_param"`
			WechatChannels struct {
				Nickname     string `json:"nickname"`
				ShopNickName string `json:"shop_nickname"`
				Scene        uint32 `json:"scene"`
			} `json:"wechat_channels"`
		} `json:"enter_session_context"`
	} `json:"customer_list"`
	InvalidExternalUserid []string `json:"invalid_external_userid"`
}

func (ww *weWork) KfCustomerBatchGet(corpId uint, userList []string, needEnterSessionContext int) (resp KfCustomerBatchGetResponse) {
	params := H{}
	if needEnterSessionContext == 1 {
		params["need_enter_session_context"] = 1
	}
	params["external_userid_list"] = userList
	_, err := ww.getRequest(corpId).SetBody(params).SetResult(&resp).
		Post("/cgi-bin/kf/customer/batchget")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type KfGetCorpQualificationResponse struct {
	internal.BizResponse
	WeChatChannelsBinding bool `json:"wechat_channels_binding"`
}

// KfGetCorpQualification 仅支持第三方应用，且需具有“微信客服->获取基础信息”权限
func (ww *weWork) KfGetCorpQualification(corpId uint) (resp KfGetCorpQualificationResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		Get("/cgi-bin/kf/get_corp_qualification")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type KfGetUpgradeServiceConfigResponse struct {
	internal.BizResponse
	MemberRange struct {
		UseridList       []string `json:"userid_list"`
		DepartmentIDList []int    `json:"department_id_list"`
	} `json:"member_range"`
	GroupchatRange struct {
		ChatIDList []string `json:"chat_id_list"`
	} `json:"groupchat_range"`
}

func (ww *weWork) KfGetUpgradeServiceConfig(corpId uint) (resp KfGetUpgradeServiceConfigResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		Get("/cgi-bin/kf/customer/get_upgrade_service_config")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type UpgradeServiceRequest struct {
	OpenKfId       string                   `json:"open_kfid" validate:"required"`
	ExternalUserId string                   `json:"external_userid" validate:"required"`
	Type           int                      `json:"type" validate:"required,oneof=1 2"`
	Member         *UpgradeServiceMember    `json:"member,omitempty"`
	GroupChat      *UpgradeServiceGroupChat `json:"groupchat,omitempty"`
}

type UpgradeServiceMember struct {
	UserId  string `json:"userid" validate:"required"`
	Wording string `json:"wording"`
}

type UpgradeServiceGroupChat struct {
	ChatId  string `json:"chat_id" validate:"required"`
	Wording string `json:"wording"`
}

func (ww *weWork) KfUpgradeService(corpId uint, request UpgradeServiceRequest) (resp internal.BizResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/kf/customer/upgrade_service")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type CancelUpgradeServiceRequest struct {
	OpenKfId       string `json:"open_kfid" validate:"required"`
	ExternalUserId string `json:"external_userid" validate::"required"`
}

func (ww *weWork) KfCancelUpgradeService(corpId uint, request CancelUpgradeServiceRequest) (resp internal.BizResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/kf/customer/cancel_upgrade_service")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type KfGetCorpStatisticFilter struct {
	OpenKfId  string `json:"open_kfid,omitempty"`
	StartTime uint32 `json:"start_time" validate:"required"`
	EndTime   uint32 `json:"end_time" validate:"required"`
}
type KfGetCorpStatisticResponse struct {
	internal.BizResponse
	StatisticList []struct {
		StatTime  int `json:"stat_time"`
		Statistic struct {
			SessionCnt                int `json:"session_cnt"`
			CustomerCnt               int `json:"customer_cnt"`
			CustomerMsgCnt            int `json:"customer_msg_cnt"`
			UpgradeServiceCustomerCnt int `json:"upgrade_service_customer_cnt"`
			AiSessionReplyCnt         int `json:"ai_session_reply_cnt"`
			AiTransferRate            int `json:"ai_transfer_rate"`
			AiKnowledgeHitRate        int `json:"ai_knowledge_hit_rate"`
		} `json:"statistic"`
	} `json:"statistic_list"`
}

func (ww *weWork) KfGetCorpStatistic(corpId uint, filter KfGetCorpStatisticFilter) (resp KfGetCorpStatisticResponse) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(filter).SetResult(&resp).
		Post("/cgi-bin/kf/get_corp_statistic")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type KfGetServicerStatisticFilter struct {
	OpenKfId       string `json:"open_kfid,omitempty"`
	ServicerUserId string `json:"servicer_userid,omitempty"`
	StartTime      uint32 `json:"start_time" validate:"required"`
	EndTime        uint32 `json:"end_time" validate:"required"`
}

type KfGetServicerStatisticResponse struct {
	internal.BizResponse
	StatisticList []struct {
		StatTime  int `json:"stat_time"`
		Statistic struct {
			SessionCnt                         int `json:"session_cnt"`
			CustomerCnt                        int `json:"customer_cnt"`
			CustomerMsgCnt                     int `json:"customer_msg_cnt"`
			ReplyRate                          int `json:"reply_rate"`
			FirstReplyAverageSec               int `json:"first_reply_average_sec"`
			SatisfactionInvestgateCnt          int `json:"satisfaction_investgate_cnt"`
			SatisfactionParticipationRate      int `json:"satisfaction_participation_rate"`
			SatisfiedRate                      int `json:"satisfied_rate"`
			MiddlingRate                       int `json:"middling_rate"`
			DissatisfiedRate                   int `json:"dissatisfied_rate"`
			UpgradeServiceCustomerCnt          int `json:"upgrade_service_customer_cnt"`
			UpgradeServiceMemberInviteCnt      int `json:"upgrade_service_member_invite_cnt"`
			UpgradeServiceMemberCustomerCnt    int `json:"upgrade_service_member_customer_cnt"`
			UpgradeServiceGroupchatInviteCnt   int `json:"upgrade_service_groupchat_invite_cnt"`
			UpgradeServiceGroupchatCustomerCnt int `json:"upgrade_service_groupchat_customer_cnt"`
		} `json:"statistic"`
	} `json:"statistic_list"`
}

func (ww *weWork) KfGetServicerStatistic(corpId uint, filter KfGetServicerStatisticFilter) (resp KfGetServicerStatisticResponse) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(filter).SetResult(&resp).
		Post("/cgi-bin/kf/get_servicer_statistic")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
