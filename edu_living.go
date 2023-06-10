package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type GetUserAllLivingIdRequest struct {
	UserId string `json:"userid" validate:"required"`
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit"`
}

type GetUserAllLivingIdResponse struct {
	internal.BizResponse
	NextCursor   string   `json:"next_cursor"`
	LivingIdList []string `json:"livingid_list"`
}

// GetUserAllLivingId 获取老师直播ID列表
// https://open.work.weixin.qq.com/api/doc/90001/90143/93856
func (ww *weWork) GetUserAllLivingId(corpId uint, request GetUserAllLivingIdRequest) (resp GetUserAllLivingIdResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/living/get_user_all_livingid")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetLivingInfoResponse struct {
	internal.BizResponse
	LivingInfo struct {
		Theme          string `json:"theme"`
		LivingStart    int    `json:"living_start"`
		LivingDuration int    `json:"living_duration"`
		AnchorUserId   string `json:"anchor_userid"`
		LivingRange    struct {
			PartyIds   []int    `json:"partyids"`
			GroupNames []string `json:"group_names"`
		} `json:"living_range"`
		ViewerNum     int    `json:"viewer_num"`
		CommentNum    int    `json:"comment_num"`
		OpenReplay    int    `json:"open_replay"`
		PushStreamURL string `json:"push_stream_url"`
	} `json:"living_info"`
}

// GetLivingInfo 获取直播详情
// https://open.work.weixin.qq.com/api/doc/90001/90143/93857
func (ww *weWork) GetLivingInfo(corpId uint, liveId string) (resp GetLivingInfoResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		SetQueryParam("livingid", liveId).Get("/cgi-bin/school/living/get_living_info")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetWatchStatRequest struct {
	LivingId string `json:"livingid" validate:"required"`
	NextKey  string `json:"next_key"`
}

type GetWatchStatResponse struct {
	internal.BizResponse
	Ending     int    `json:"ending"`
	NextKey    string `json:"next_key"`
	StatInfoes struct {
		Students []struct {
			StudentUserid string `json:"student_userid"`
			ParentUserid  string `json:"parent_userid"`
			Partyids      []int  `json:"partyids"`
			WatchTime     int    `json:"watch_time"`
			EnterTime     int    `json:"enter_time"`
			LeaveTime     int    `json:"leave_time"`
			IsComment     int    `json:"is_comment"`
		} `json:"students"`
		Visitors []struct {
			Nickname  string `json:"nickname"`
			WatchTime int    `json:"watch_time"`
			EnterTime int    `json:"enter_time"`
			LeaveTime int    `json:"leave_time"`
			IsComment int    `json:"is_comment"`
		} `json:"visitors"`
	} `json:"stat_infoes"`
}

// GetWatchStat 获取观看直播统计
// https://open.work.weixin.qq.com/api/doc/90001/90143/93858
func (ww *weWork) GetWatchStat(corpId uint, request GetWatchStatRequest) (resp GetWatchStatResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/school/living/get_watch_stat")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetUnWatchStatResponse struct {
	internal.BizResponse
	Ending   int    `json:"ending"`
	NextKey  string `json:"next_key"`
	StatInfo struct {
		Students []struct {
			StudentUserid string `json:"student_userid"`
			ParentUserid  string `json:"parent_userid"`
			Partyids      []int  `json:"partyids"`
		} `json:"students"`
	} `json:"stat_info"`
}

// GetUnWatchStat 获取未观看直播统计
// https://open.work.weixin.qq.com/api/doc/90001/90143/93859
func (ww *weWork) GetUnWatchStat(corpId uint, request GetWatchStatRequest) (resp GetUnWatchStatResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/school/living/get_unwatch_stat")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// DeleteReplayData 删除直播回放
// https://open.work.weixin.qq.com/api/doc/90001/90143/93860
func (ww *weWork) DeleteReplayData(corpId uint, livingId string) (resp internal.BizResponse) {
	h := H{}
	h["livingid"] = livingId
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/living/delete_replay_data")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
