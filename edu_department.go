package wework

import (
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type SchoolDepartment struct {
	Name             string `json:"name,omitempty"`
	ParentId         int32  `json:"parentid,omitempty" validate:"required"`
	Id               int32  `json:"id"`
	NewId            int32  `json:"new_id,omitempty"`
	Type             int32  `json:"type,omitempty" validate:"required,oneof=1 2 3 4"`
	RegisterYear     int    `json:"register_year,omitempty" validate:"omitempty,min=1970,max=2100"`
	StandardGrade    int    `json:"standard_grade,omitempty"`
	Order            int    `json:"order,omitempty"`
	DepartmentAdmins []struct {
		Userid  string `json:"userid"`
		Type    int    `json:"type" validate:"oneof=1 2 3 4 5"`
		Subject string `json:"subject"`
	} `json:"department_admins,omitempty"`
}

type SchoolDepartmentCreateResponse struct {
	internal.BizResponse
	Id int32 `json:"id"`
}

// SchoolDepartmentCreate 创建部门
// https://open.work.weixin.qq.com/api/doc/90001/90143/92296
func (ww *weWork) SchoolDepartmentCreate(corpId uint, department SchoolDepartment) (resp SchoolDepartmentCreateResponse) {
	if ok := validate.Struct(department); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(department).SetResult(&resp).
		Post("/cgi-bin/school/department/create")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// SchoolDepartmentUpdate 更新部门
// https://open.work.weixin.qq.com/api/doc/90001/90143/92297
func (ww *weWork) SchoolDepartmentUpdate(corpId uint, department SchoolDepartment) (resp internal.BizResponse) {
	if department.Id < 0 {
		resp.ErrCode = 500
		resp.ErrorMsg = "department id must be uint32"
		return
	}
	_, err := ww.getRequest(corpId).SetBody(department).SetResult(&resp).
		Post("/cgi-bin/school/department/update")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// SchoolDepartmentDelete 删除部门
// https://open.work.weixin.qq.com/api/doc/90001/90143/92298
func (ww *weWork) SchoolDepartmentDelete(corpId uint, departmentId int32) (resp internal.BizResponse) {
	_, err := ww.getRequest(corpId).
		SetQueryParam("id", fmt.Sprintf("%v", departmentId)).
		SetResult(&resp).Get("/cgi-bin/school/department/delete")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type SchoolDepartmentListResponse struct {
	internal.BizResponse
	Departments []struct {
		Name             string `json:"name"`
		ParentId         int    `json:"parentid"`
		ID               int    `json:"id"`
		Type             int    `json:"type"`
		RegisterYear     int    `json:"register_year,omitempty"`
		StandardGrade    int    `json:"standard_grade,omitempty"`
		Order            int    `json:"order,omitempty"`
		DepartmentAdmins []struct {
			Userid string `json:"userid"`
			Type   int    `json:"type"`
		} `json:"department_admins"`
		IsGraduated   int    `json:"is_graduated,omitempty"`
		OpenGroupChat int    `json:"open_group_chat,omitempty"`
		GroupChatID   string `json:"group_chat_id,omitempty"`
	} `json:"departments"`
}

// SchoolDepartmentList 获取部门列表
// https://open.work.weixin.qq.com/api/doc/90001/90143/92299
func (ww *weWork) SchoolDepartmentList(corpId uint, departmentId int32) (resp SchoolDepartmentListResponse) {
	_, err := ww.getRequest(corpId).
		SetQueryParam("id", fmt.Sprintf("%v", departmentId)).
		SetResult(&resp).Get("/cgi-bin/school/department/list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
