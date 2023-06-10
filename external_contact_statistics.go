package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type GetUserBehaviorDataResponse struct {
	internal.BizResponse
}
type GetUserBehaviorFilter struct {
	UserId    []string `json:"userid" validate:"required_without=PartyId,max=100"`
	PartyId   []uint32 `json:"partyid" validate:"required_without=UserId,max=100"`
	StartTime int      `json:"start_time"`
	EndTime   int      `json:"end_time"`
}

// GetUserBehaviorData 获取「联系客户统计」数据
// https://open.work.weixin.qq.com/api/doc/90001/90143/92275
func (ww *weWork) GetUserBehaviorData(corpId uint, filter GetUserBehaviorFilter) (resp GetUserBehaviorDataResponse) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
	}
	_, err := ww.getRequest(corpId).SetBody(filter).SetResult(&resp).
		Post("/cgi-bin/externalcontact/get_user_behavior_data")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GroupChatStatisticResponse struct {
	internal.BizResponse
	Total      int `json:"total"`
	NextOffset int `json:"next_offset"`
	Items      []struct {
		Owner string `json:"owner"`
		Data  struct {
			NewChatCnt            int `json:"new_chat_cnt"`
			ChatTotal             int `json:"chat_total"`
			ChatHasMsg            int `json:"chat_has_msg"`
			NewMemberCnt          int `json:"new_member_cnt"`
			MemberTotal           int `json:"member_total"`
			MemberHasMsg          int `json:"member_has_msg"`
			MsgTotal              int `json:"msg_total"`
			MigrateTraineeChatCnt int `json:"migrate_trainee_chat_cnt"`
		} `json:"data"`
	} `json:"items"`
}
type GroupChatStatisticFilter struct {
	DayBeginTime int `json:"day_begin_time" validate:"required"`
	DayEndTime   int `json:"day_end_time"`
	OwnerFilter  struct {
		UseridList []string `json:"userid_list" validate:"required,max=100"`
	} `json:"owner_filter" validate:"required"`
	OrderBy  int `json:"order_by"`
	OrderAsc int `json:"order_asc"`
	Offset   int `json:"offset"`
	Limit    int `json:"limit"`
}

// GroupChatStatistic 按群主聚合的方式
// https://open.work.weixin.qq.com/api/doc/90001/90143/93476#%E6%8C%89%E7%BE%A4%E4%B8%BB%E8%81%9A%E5%90%88%E7%9A%84%E6%96%B9%E5%BC%8F
func (ww *weWork) GroupChatStatistic(corpId uint, filter GroupChatStatisticFilter) (resp GroupChatStatisticResponse) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
	}
	_, err := ww.getRequest(corpId).SetBody(filter).SetResult(&resp).
		Post("/cgi-bin/externalcontact/groupchat/statistic")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GroupChatStatisticGroupByDayFilter struct {
	DayBeginTime int `json:"day_begin_time" validate:"required"`
	DayEndTime   int `json:"day_end_time"`
	OwnerFilter  struct {
		UseridList []string `json:"userid_list" validate:"required,max=100"`
	} `json:"owner_filter" validate:"required"`
}

// GroupChatStatisticGroupByDay 按自然日聚合的方式
// https://open.work.weixin.qq.com/api/doc/90001/90143/93476#%E6%8C%89%E8%87%AA%E7%84%B6%E6%97%A5%E8%81%9A%E5%90%88%E7%9A%84%E6%96%B9%E5%BC%8F
func (ww *weWork) GroupChatStatisticGroupByDay(corpId uint, filter GroupChatStatisticGroupByDayFilter) (resp GroupChatStatisticResponse) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
	}
	_, err := ww.getRequest(corpId).SetBody(filter).SetResult(&resp).
		Post("/cgi-bin/externalcontact/groupchat/statistic_group_by_day")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
