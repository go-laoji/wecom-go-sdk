package wework

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/internal"
)

type MomentTask struct {
	Text         Text          `json:"text,omitempty"`
	Attachments  []Attachments `json:"attachments" validate:"required_without=Text.Content"`
	VisibleRange VisibleRange  `json:"visible_range,omitempty"`
}

type Text struct {
	Content string `json:"content"`
}
type Image struct {
	MediaID string `json:"media_id" validate:"required"`
}

// Video 应用消息发关时title和description为可选项
// 朋友圈发送时只设置　media_id即可
type Video struct {
	MediaID     string `json:"media_id" validate:"required"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}
type Link struct {
	Title   string `json:"title"`
	URL     string `json:"url" validate:"required"`
	MediaID string `json:"media_id" validate:"required"`
}
type Attachments struct {
	Msgtype string `json:"msgtype" validate:"required,oneof=image link video"`
	Image   *Image `json:"image,omitempty" validate:"required_without_all=Video Link"`
	Video   *Video `json:"video,omitempty" validate:"required_without_all=Image Link"`
	Link    *Link  `json:"link,omitempty" validate:"required_without_all=Video Image"`
}
type SenderList struct {
	UserList       []string `json:"user_list"`
	DepartmentList []int    `json:"department_list"`
}
type ExternalContactList struct {
	TagList []string `json:"tag_list"`
}
type VisibleRange struct {
	SenderList          SenderList          `json:"sender_list,omitempty"`
	ExternalContactList ExternalContactList `json:"external_contact_list,omitempty"`
}

type AddMomentTaskResponse struct {
	internal.BizResponse
	JobId string `json:"jobid"`
}

// AddMomentTask 创建发表任务
// 参考连接　https://open.work.weixin.qq.com/api/doc/90001/90143/95095#%E5%88%9B%E5%BB%BA%E5%8F%91%E8%A1%A8%E4%BB%BB%E5%8A%A1
func (ww weWork) AddMomentTask(corpId uint, task MomentTask) (resp AddMomentTaskResponse) {
	if ok := validate.Struct(task); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/add_moment_task?%s", queryParams.Encode()), task)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type GetMomentTaskResultResponse struct {
	internal.BizResponse
	Status int    `json:"status"`
	Type   string `json:"type"`
	Result struct {
		internal.BizResponse
		MomentId          string `json:"moment_id"`
		InvalidSenderList struct {
			UserList       []string `json:"user_list"`
			DepartmentList []int32  `json:"department_list"`
		} `json:"invalid_sender_list"`
		InvalidExternalContactList struct {
			TagList []string `json:"tag_list"`
		} `json:"invalid_external_contact_list"`
	}
}

// GetMomentTaskResult 获取任务创建结果
// https://open.work.weixin.qq.com/api/doc/90001/90143/95095#%E8%8E%B7%E5%8F%96%E4%BB%BB%E5%8A%A1%E5%88%9B%E5%BB%BA%E7%BB%93%E6%9E%9C
func (ww weWork) GetMomentTaskResult(corpId uint, jobId string) (resp GetMomentTaskResultResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	queryParams.Add("jobid", jobId)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/externalcontact/get_moment_task_result?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type MomentListFilter struct {
	StartTime  int64  `json:"start_time" validate:"required"`
	EndTime    int64  `json:"end_time" validate:"required"`
	Creator    string `json:"creator,omitempty"`
	FilterType int    `json:"filter_type,omitempty" validate:"omitempty,oneof=0 1 2"`
	Cursor     string `json:"cursor"`
	Limit      int    `json:"limit"`
}

type GetMomentListResponse struct {
	internal.BizResponse
	NextCursor string       `json:"next_cursor"`
	MomentList []MomentList `json:"moment_list"`
}

type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Name      string `json:"name"`
}
type MomentList struct {
	MomentID    string   `json:"moment_id"`
	Creator     string   `json:"creator"`
	CreateTime  string   `json:"create_time"`
	CreateType  int      `json:"create_type"`
	VisibleType int      `json:"visible_type"`
	Text        Text     `json:"text"`
	Image       []Image  `json:"image"`
	Video       Video    `json:"video"`
	Link        Link     `json:"link"`
	Location    Location `json:"location"`
}

// GetMomentList 获取企业全部的发表列表
// https://open.work.weixin.qq.com/api/doc/90001/90143/93443#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E5%85%A8%E9%83%A8%E7%9A%84%E5%8F%91%E8%A1%A8%E5%88%97%E8%A1%A8
func (ww weWork) GetMomentList(corpId uint, filter MomentListFilter) (resp GetMomentListResponse) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/get_moment_list?%s", queryParams.Encode()), filter)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type MomentTaskFilter struct {
	MomentId string `json:"moment_id" validate:"required"`
	Cursor   string `json:"cursor"`
	Limit    int    `json:"limit"`
}

type GetMomentTaskResponse struct {
	internal.BizResponse
	NextCursor string `json:"next_cursor"`
	TaskList   []struct {
		UserId        string `json:"userid"`
		PublishStatus int    `json:"publish_status"`
	} `json:"task_list"`
}

// GetMomentTask 获取客户朋友圈企业发表的列表
// https://open.work.weixin.qq.com/api/doc/90001/90143/93443#%E8%8E%B7%E5%8F%96%E5%AE%A2%E6%88%B7%E6%9C%8B%E5%8F%8B%E5%9C%88%E4%BC%81%E4%B8%9A%E5%8F%91%E8%A1%A8%E7%9A%84%E5%88%97%E8%A1%A8
func (ww weWork) GetMomentTask(corpId uint, filter MomentTaskFilter) (resp GetMomentTaskResponse) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/get_moment_task?%s", queryParams.Encode()), filter)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type MomentCustomerFilter struct {
	MomentId string `json:"moment_id" validate:"required"`
	UserId   string `json:"userid" validate:"required"`
	Cursor   string `json:"cursor,omitempty"`
	Limit    int    `json:"limit,omitempty" validate:"omitempty,max=1000"`
}
type GetMomentCustomerListResponse struct {
	internal.BizResponse
	NextCursor   string `json:"next_cursor"`
	CustomerList []struct {
		UserId         string `json:"userid"`
		ExternalUserId string `json:"external_userid"`
	} `json:"customer_list"`
}

// GetMomentCustomerList 获取客户朋友圈发表时选择的可见范围
// https://open.work.weixin.qq.com/api/doc/90001/90143/93443#%E8%8E%B7%E5%8F%96%E5%AE%A2%E6%88%B7%E6%9C%8B%E5%8F%8B%E5%9C%88%E5%8F%91%E8%A1%A8%E6%97%B6%E9%80%89%E6%8B%A9%E7%9A%84%E5%8F%AF%E8%A7%81%E8%8C%83%E5%9B%B4
func (ww weWork) GetMomentCustomerList(corpId uint, filter MomentCustomerFilter) (resp GetMomentCustomerListResponse) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/get_moment_customer_list?%s", queryParams.Encode()), filter)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type GetMomentSendResultResponse struct {
	internal.BizResponse
	NextCursor   string `json:"next_cursor"`
	CustomerList []struct {
		ExternalUserId string `json:"external_userid"`
	} `json:"customer_list"`
}

// GetMomentSendResult 获取客户朋友圈发表后的可见客户列表
// https://open.work.weixin.qq.com/api/doc/90001/90143/93443#%E8%8E%B7%E5%8F%96%E5%AE%A2%E6%88%B7%E6%9C%8B%E5%8F%8B%E5%9C%88%E5%8F%91%E8%A1%A8%E5%90%8E%E7%9A%84%E5%8F%AF%E8%A7%81%E5%AE%A2%E6%88%B7%E5%88%97%E8%A1%A8
func (ww weWork) GetMomentSendResult(corpId uint, filter MomentCustomerFilter) (resp GetMomentSendResultResponse) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/get_moment_send_result?%s", queryParams.Encode()), filter)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type GetMomentCommentsResponse struct {
	internal.BizResponse
	CommentList []struct {
		ExternalUserId string `json:"external_userid"`
		CreateTime     int    `json:"create_time"`
	} `json:"comment_list"`
	LikeList []struct {
		ExternalUserId string `json:"external_userid"`
		CreateTime     int    `json:"create_time"`
	} `json:"like_list"`
}

// GetMomentComments 获取客户朋友圈的互动数据
// https://open.work.weixin.qq.com/api/doc/90001/90143/93443#%E8%8E%B7%E5%8F%96%E5%AE%A2%E6%88%B7%E6%9C%8B%E5%8F%8B%E5%9C%88%E7%9A%84%E4%BA%92%E5%8A%A8%E6%95%B0%E6%8D%AE
func (ww weWork) GetMomentComments(corpId uint, momentId string, userId string) (resp GetMomentCommentsResponse) {
	p := H{"userid": userId, "moment_id": momentId}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/get_moment_comments?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
