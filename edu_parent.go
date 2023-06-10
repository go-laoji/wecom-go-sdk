package wework

import (
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
	"strings"
)

type Parent struct {
	ParentUserId    string `json:"parent_userid" validate:"required"`
	NewParentUserId string `json:"new_parent_userid,omitempty"`
	Mobile          string `json:"mobile,omitempty" validate:"required"`
	ToInvite        bool   `json:"to_invite,omitempty"`
	Children        []struct {
		StudentUserId string `json:"student_userid"`
		Relation      string `json:"relation"`
	} `json:"children,omitempty" validate:"required,max=10"`
}

// CreateParent 创建家长
// https://open.work.weixin.qq.com/api/doc/90001/90143/92077
func (ww *weWork) CreateParent(corpId uint, parent Parent) (resp internal.BizResponse) {
	if ok := validate.Struct(parent); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(parent).SetResult(&resp).
		Post("/cgi-bin/school/user/create_parent")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type BatchParentResponse struct {
	internal.BizResponse
	ResultList []struct {
		ParentUserId string `json:"parent_userid"`
		internal.BizResponse
	} `json:"result_list"`
}

// BatchCreateParent 批量创建家长
// https://open.work.weixin.qq.com/api/doc/90001/90143/92078
func (ww *weWork) BatchCreateParent(corpId uint, parents []Parent) (resp BatchParentResponse) {
	h := H{}
	h["parents"] = parents
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/school/user/batch_create_parent")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// DeleteParent 删除家长
// https://open.work.weixin.qq.com/api/doc/90001/90143/92079
func (ww *weWork) DeleteParent(corpId uint, userId string) (resp internal.BizResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).SetQueryParam("userid", userId).
		Get("/cgi-bin/school/user/delete_parent")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// BatchDeleteParent 批量删除家长
// https://open.work.weixin.qq.com/api/doc/90001/90143/92080
func (ww *weWork) BatchDeleteParent(corpId uint, userIdList []string) (resp BatchParentResponse) {
	h := H{}
	h["useridlist"] = userIdList
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/school/user/batch_delete_parent")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// UpdateParent 更新家长
// https://open.work.weixin.qq.com/api/doc/90001/90143/92081
func (ww *weWork) UpdateParent(corpId uint, parent Parent) (resp internal.BizResponse) {
	if strings.TrimSpace(parent.ParentUserId) == "" {
		resp.ErrCode = 500
		resp.ErrorMsg = "parent userid can not be empty"
		return
	}
	_, err := ww.getRequest(corpId).SetBody(parent).SetResult(&resp).
		Post("/cgi-bin/school/user/update_parent")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// BatchUpdateParent 批量更新家长
// https://open.work.weixin.qq.com/api/doc/90001/90143/92082
func (ww *weWork) BatchUpdateParent(corpId uint, parents []Parent) (resp BatchParentResponse) {
	h := H{}
	h["parents"] = parents
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/school/user/batch_update_parent")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type ListParentWithDepartmentIdResponse struct {
	internal.BizResponse
	Parents []struct {
		ParentUserid   string `json:"parent_userid"`
		Mobile         string `json:"mobile"`
		IsSubscribe    int    `json:"is_subscribe"`
		ExternalUserid string `json:"external_userid,omitempty"`
		Children       []struct {
			StudentUserid string `json:"student_userid"`
			Relation      string `json:"relation"`
			Name          string `json:"name"`
		} `json:"children"`
	} `json:"parents"`
}

// ListParentWithDepartmentId 获取部门家长详情
// https://open.work.weixin.qq.com/api/doc/90001/90143/92627
func (ww *weWork) ListParentWithDepartmentId(corpId uint, departmentId int32) (resp ListParentWithDepartmentIdResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		SetQueryParam("department_id", fmt.Sprintf("%v", departmentId)).
		Get("/cgi-bin/school/user/list_parent")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
