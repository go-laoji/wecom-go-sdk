package wework

import (
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
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
func (ww *weWork) SchoolUserGet(corpId uint, userId string) (resp SchoolUserGetResponse) {
	_, err := ww.getRequest(corpId).SetQueryParam("userid", userId).
		SetResult(&resp).Get("/cgi-bin/school/user/get")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
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
func (ww *weWork) SchoolUserList(corpId uint, departmentId uint32, fetchChild int) (resp SchoolUserListResponse) {
	_, err := ww.getRequest(corpId).
		SetQueryParam("department_id", fmt.Sprintf("%v", departmentId)).
		SetQueryParam("fetch_child", fmt.Sprintf("%v", fetchChild)).
		SetResult(&resp).Get("/cgi-bin/school/user/list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// SetArchSyncMode 设置家校通讯录自动同步模式
// https://open.work.weixin.qq.com/api/doc/90001/90143/92083
func (ww *weWork) SetArchSyncMode(corpId uint, mode int) (resp internal.BizResponse) {
	h := H{}
	h["arch_sync_mode"] = mode
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/school/set_arch_sync_mode")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
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
func (ww *weWork) GetSubScribeQrCode(corpId uint) (resp GetSubScribeQrCodeResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		Get("/cgi-bin/externalcontact/get_subscribe_qr_code")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// SetSubScribeMode 设置关注「学校通知」的模式
// https://open.work.weixin.qq.com/api/doc/90001/90143/92290#%E8%AE%BE%E7%BD%AE%E5%85%B3%E6%B3%A8%E3%80%8C%E5%AD%A6%E6%A0%A1%E9%80%9A%E7%9F%A5%E3%80%8D%E7%9A%84%E6%A8%A1%E5%BC%8F
// 关注模式, 1:可扫码填写资料加入, 2:禁止扫码填写资料加入
func (ww *weWork) SetSubScribeMode(corpId uint, mode int) (resp internal.BizResponse) {
	h := H{}
	h["subscribe_mode"] = mode
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/externalcontact/set_subscribe_mode")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetSubScribeModeResponse struct {
	internal.BizResponse
	SubscribeMode int `json:"subscribe_mode"`
}

// GetSubScribeMode 获取关注「学校通知」的模式
// https://open.work.weixin.qq.com/api/doc/90001/90143/92290#%E8%8E%B7%E5%8F%96%E5%85%B3%E6%B3%A8%E3%80%8C%E5%AD%A6%E6%A0%A1%E9%80%9A%E7%9F%A5%E3%80%8D%E7%9A%84%E6%A8%A1%E5%BC%8F
func (ww *weWork) GetSubScribeMode(corpId uint) (resp GetSubScribeModeResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		Get("/cgi-bin/externalcontact/get_subscribe_mode")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
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
func (ww *weWork) BatchToExternalUserId(corpId uint, mobiles []string) (resp BatchToExternalUserIdResponse) {
	h := H{}
	h["mobiles"] = mobiles
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/externalcontact/batch_to_external_userid")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// SetTeacherViewMode 设置「老师可查看班级」的模式
// https://open.work.weixin.qq.com/api/doc/90001/90143/92652#%E8%AE%BE%E7%BD%AE%E3%80%8C%E8%80%81%E5%B8%88%E5%8F%AF%E6%9F%A5%E7%9C%8B%E7%8F%AD%E7%BA%A7%E3%80%8D%E7%9A%84%E6%A8%A1%E5%BC%8F
func (ww *weWork) SetTeacherViewMode(corpId uint, mode int) (resp internal.BizResponse) {
	h := H{}
	h["view_mode"] = mode
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/school/set_teacher_view_mode")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetTeacherViewModeResponse struct {
	internal.BizResponse
	ViewMode int `json:"view_mode"`
}

// GetTeacherViewMode 获取「老师可查看班级」的模式
// https://open.work.weixin.qq.com/api/doc/90001/90143/92652#%E8%8E%B7%E5%8F%96%E3%80%8C%E8%80%81%E5%B8%88%E5%8F%AF%E6%9F%A5%E7%9C%8B%E7%8F%AD%E7%BA%A7%E3%80%8D%E7%9A%84%E6%A8%A1%E5%BC%8F
func (ww *weWork) GetTeacherViewMode(corpId uint) (resp GetTeacherViewModeResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		Get("/cgi-bin/school/get_teacher_view_mode")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
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
func (ww *weWork) GetAllowScope(corpId uint, agentId int) (resp GetAllowScopeResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		Get("/cgi-bin/school/agent/get_allow_scope")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
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
func (ww *weWork) SetUpgradeInfo(corpId uint, request UpgradeRequest) (resp UpgradeInfoResponse) {
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/school/set_upgrade_info")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
