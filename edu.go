package wework

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/internal"
)

type SchoolUserGetResponse struct {
	internal.BizResponse
	UserType int `json:"user_type"`
	Student  struct {
		StudentUserId string `json:"student_userid"`
		Name          string `json:"name"`
		Department    []int  `json:"department"`
		Parents       []struct {
			ParentUserId   string `json:"parent_userid"`
			Relation       string `json:"relation"`
			Mobile         string `json:"mobile"`
			IsSubscribe    int    `json:"is_subscribe"`
			ExternalUserId string `json:"external_userid,omitempty"`
		} `json:"parents"`
	} `json:"student,omitempty"`
	Parent struct {
		ParentUserId   string `json:"parent_userid"`
		Mobile         string `json:"mobile"`
		IsSubscribe    int    `json:"is_subscribe"`
		ExternalUserId string `json:"external_userid"`
		Children       []struct {
			StudentUserId string `json:"student_userid"`
			Relation      string `json:"relation"`
		} `json:"children"`
	} `json:"parent,omitempty"`
}

// SchoolUserGet 读取学生或家长
// https://open.work.weixin.qq.com/api/doc/90001/90143/92038
func (ww weWork) SchoolUserGet(corpId uint, userId string) (resp SchoolUserGetResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	queryParams.Add("userid", userId)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/school/user/get?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type SchoolUserListResponse struct {
	internal.BizResponse
	Students []struct {
		StudentUserid string `json:"student_userid"`
		Name          string `json:"name"`
		Department    []int  `json:"department"`
		Parents       []struct {
			ParentUserid   string `json:"parent_userid"`
			Relation       string `json:"relation"`
			Mobile         string `json:"mobile"`
			IsSubscribe    int    `json:"is_subscribe"`
			ExternalUserid string `json:"external_userid,omitempty"`
		} `json:"parents"`
	} `json:"students"`
}

// SchoolUserList 获取部门成员详情
// https://open.work.weixin.qq.com/api/doc/90001/90143/92043
func (ww weWork) SchoolUserList(corpId uint, departmentId uint32, fetchChild int) (resp SchoolUserListResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	queryParams.Add("department_id", fmt.Sprintf("%v", departmentId))
	queryParams.Add("fetch_child", fmt.Sprintf("%v", fetchChild))
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/school/user/list?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// SetArchSyncMode 设置家校通讯录自动同步模式
// https://open.work.weixin.qq.com/api/doc/90001/90143/92083
func (ww weWork) SetArchSyncMode(corpId uint, mode int) (resp internal.BizResponse) {
	h := H{}
	h["arch_sync_mode"] = mode
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/school/set_arch_sync_mode?%s", queryParams.Encode()), h)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type GetSubScribeQrCodeResponse struct {
	internal.BizResponse
	QrCodeBig    string `json:"qrcode_big"`
	QrCodeMiddle string `json:"qrcode_middle"`
	QrCodeThumb  string `json:"qrcode_thumb"`
}

// GetSubScribeQrCode 获取「学校通知」二维码
// https://open.work.weixin.qq.com/api/doc/90001/90143/92197
func (ww weWork) GetSubScribeQrCode(corpId uint) (resp GetSubScribeQrCodeResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/externalcontact/get_subscribe_qr_code?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// SetSubScribeMode 设置关注「学校通知」的模式
// https://open.work.weixin.qq.com/api/doc/90001/90143/92290#%E8%AE%BE%E7%BD%AE%E5%85%B3%E6%B3%A8%E3%80%8C%E5%AD%A6%E6%A0%A1%E9%80%9A%E7%9F%A5%E3%80%8D%E7%9A%84%E6%A8%A1%E5%BC%8F
// 关注模式, 1:可扫码填写资料加入, 2:禁止扫码填写资料加入
func (ww weWork) SetSubScribeMode(corpId uint, mode int) (resp internal.BizResponse) {
	h := H{}
	h["subscribe_mode"] = mode
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/set_subscribe_mode?%s", queryParams.Encode()), h)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type GetSubScribeModeResponse struct {
	internal.BizResponse
	SubscribeMode int `json:"subscribe_mode"`
}

// GetSubScribeMode 获取关注「学校通知」的模式
// https://open.work.weixin.qq.com/api/doc/90001/90143/92290#%E8%8E%B7%E5%8F%96%E5%85%B3%E6%B3%A8%E3%80%8C%E5%AD%A6%E6%A0%A1%E9%80%9A%E7%9F%A5%E3%80%8D%E7%9A%84%E6%A8%A1%E5%BC%8F
func (ww weWork) GetSubScribeMode(corpId uint) (resp GetSubScribeModeResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/externalcontact/get_subscribe_mode?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type BatchToExternalUserIdResponse struct {
	internal.BizResponse
	SuccessList []struct {
		Mobile         string `json:"mobile"`
		ExternalUserid string `json:"external_userid"`
		ForeignKey     string `json:"foreign_key"`
	} `json:"success_list"`
	FailList []struct {
		internal.BizResponse
		Mobile string `json:"mobile"`
	} `json:"fail_list"`
}

// BatchToExternalUserId 手机号转外部联系人ID
// https://open.work.weixin.qq.com/api/doc/90001/90143/92506
func (ww weWork) BatchToExternalUserId(corpId uint, mobiles []string) (resp BatchToExternalUserIdResponse) {
	h := H{}
	h["mobiles"] = mobiles
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/externalcontact/batch_to_external_userid?%s", queryParams.Encode()), h)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// SetTeacherViewMode 设置「老师可查看班级」的模式
// https://open.work.weixin.qq.com/api/doc/90001/90143/92652#%E8%AE%BE%E7%BD%AE%E3%80%8C%E8%80%81%E5%B8%88%E5%8F%AF%E6%9F%A5%E7%9C%8B%E7%8F%AD%E7%BA%A7%E3%80%8D%E7%9A%84%E6%A8%A1%E5%BC%8F
func (ww weWork) SetTeacherViewMode(corpId uint, mode int) (resp internal.BizResponse) {
	h := H{}
	h["view_mode"] = mode
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/school/set_teacher_view_mode?%s", queryParams.Encode()), h)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type GetTeacherViewModeResponse struct {
	internal.BizResponse
	ViewMode int `json:"view_mode"`
}

// GetTeacherViewMode 获取「老师可查看班级」的模式
// https://open.work.weixin.qq.com/api/doc/90001/90143/92652#%E8%8E%B7%E5%8F%96%E3%80%8C%E8%80%81%E5%B8%88%E5%8F%AF%E6%9F%A5%E7%9C%8B%E7%8F%AD%E7%BA%A7%E3%80%8D%E7%9A%84%E6%A8%A1%E5%BC%8F
func (ww weWork) GetTeacherViewMode(corpId uint) (resp GetTeacherViewModeResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/school/get_teacher_view_mode?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type GetAllowScopeResponse struct {
	internal.BizResponse
	AllowScope struct {
		Students []struct {
			Userid string `json:"userid"`
		} `json:"students"`
		Departments []int `json:"departments"`
	} `json:"allow_scope"`
}

// GetAllowScope 获取可使用的家长范围
// https://open.work.weixin.qq.com/api/doc/90001/90143/94960
func (ww weWork) GetAllowScope(corpId uint, agentId int) (resp GetAllowScopeResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	queryParams.Add("agentid", fmt.Sprintf("%v", agentId))
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/school/agent/get_allow_scope?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type UpgradeRequest struct {
	UpgradeTime   int `json:"upgrade_time,omitempty"`
	UpgradeSwitch int `json:"upgrade_switch,omitempty" validate:"omitempty,oneof=0 1"`
}
type UpgradeInfoResponse struct {
	internal.BizResponse
	NextUpgradeTime int `json:"next_upgrade_time"`
}

// SetUpgradeInfo 修改自动升年级的配置
// https://open.work.weixin.qq.com/api/doc/90001/90143/92950
func (ww weWork) SetUpgradeInfo(corpId uint, request UpgradeRequest) (resp UpgradeInfoResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/school/set_upgrade_info?%s", queryParams.Encode()), request)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
