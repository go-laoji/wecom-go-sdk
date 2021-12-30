package wework

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wework/internal"
)

type SchoolDepartment struct {
	Name             string `json:"name"`
	ParentId         int32  `json:"parentid" validate:"required"`
	Id               int32  `json:"id"`
	NewId            int32  `json:"new_id,omitempty"`
	Type             int32  `json:"type" validate:"required,oneof=1 2 3 4"`
	RegisterYear     int    `json:"register_year" validate:"omitempty,min=1970,max=2100"`
	StandardGrade    int    `json:"standard_grade"`
	Order            int    `json:"order"`
	DepartmentAdmins []struct {
		Userid  string `json:"userid"`
		Type    int    `json:"type" validate:"oneof=1 2 3 4 5"`
		Subject string `json:"subject"`
	} `json:"department_admins"`
}

type SchoolDepartmentCreateResponse struct {
	internal.BizResponse
	Id int32 `json:"id"`
}

// SchoolDepartmentCreate 创建部门
// https://open.work.weixin.qq.com/api/doc/90001/90143/92296
func (ww weWork) SchoolDepartmentCreate(corpId uint, department SchoolDepartment) (resp SchoolDepartmentCreateResponse) {
	if ok := validate.Struct(department); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/school/department/create?%s", queryParams.Encode()), department)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// SchoolDepartmentUpdate 更新部门
// https://open.work.weixin.qq.com/api/doc/90001/90143/92297
func (ww weWork) SchoolDepartmentUpdate(corpId uint, department SchoolDepartment) (resp internal.BizResponse) {
	if ok := validate.Struct(department); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/school/department/update?%s", queryParams.Encode()), department)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// SchoolDepartmentDelete 删除部门
// https://open.work.weixin.qq.com/api/doc/90001/90143/92298
func (ww weWork) SchoolDepartmentDelete(corpId uint, departmentId int32) (resp internal.BizResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	queryParams.Add("id", fmt.Sprintf("%v", departmentId))
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/school/department/delete?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
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
func (ww weWork) SchoolDepartmentList(corpId uint, departmentId int32) (resp SchoolDepartmentListResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	queryParams.Add("id", fmt.Sprintf("%v", departmentId))
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/school/department/list?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
