package wework

import (
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type Department struct {
	Id               int32    `json:"id"`
	Order            int32    `json:"order,omitempty"`
	ParentId         int32    `json:"parentid,omitempty" validate:"required"`
	Name             string   `json:"name,omitempty" validate:"required,min=1,max=64"`
	NameEn           string   `json:"name_en,omitempty" validate:"omitempty,min=1,max=64"`
	DepartmentLeader []string `json:"department_leader,omitempty"`
}

type DepartmentCreateResponse struct {
	internal.BizResponse
	Id int32 `json:"id"`
}

// DepartmentCreate 创建部门
func (ww *weWork) DepartmentCreate(corpId uint, department Department) (resp DepartmentCreateResponse) {
	if ok := validate.Struct(department); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(department).
		SetResult(&resp).Post("/cgi-bin/department/create")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// DepartmentUpdate 更新部门
func (ww *weWork) DepartmentUpdate(corpId uint, department Department) (resp internal.BizResponse) {
	if department.Id < 1 {
		resp.ErrCode = 500
		resp.ErrorMsg = "department id must be uint"
		return
	}
	_, err := ww.getRequest(corpId).SetBody(department).SetResult(&resp).Post("/cgi-bin/department/update")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// DepartmentDelete 删除部门
func (ww *weWork) DepartmentDelete(corpId uint, id int32) (resp internal.BizResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		SetQueryParam("id", fmt.Sprintf("%v", id)).
		Get("/cgi-bin/department/delete")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type DepartmentListResponse struct {
	internal.BizResponse
	Department []Department `json:"department"`
}

// DepartmentList 获取部门列表
// https://open.work.weixin.qq.com/api/doc/90001/90143/90344
func (ww *weWork) DepartmentList(corpId uint, id uint) (resp DepartmentListResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		SetQueryParam("id", fmt.Sprintf("%v", id)).
		Get("/cgi-bin/department/list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type DepartmentSimpleListResponse struct {
	internal.BizResponse
	DepartmentId []struct {
		Id       int32 `json:"id"`
		ParentId int32 `json:"parentid"`
		Order    int32 `json:"order"`
	} `json:"department_id"`
}

// DepartmentSimpleList 获取子部门ID列表
// https://developer.work.weixin.qq.com/document/path/95406
func (ww *weWork) DepartmentSimpleList(corpId uint, id int32) (resp DepartmentSimpleListResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		SetQueryParam("id", fmt.Sprintf("%v", id)).
		Get("/cgi-bin/department/simplelist")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type DepartmentGetResponse struct {
	internal.BizResponse
	Department Department `json:"department"`
}

// DepartmentGet 获取单个部门详情
// https://developer.work.weixin.qq.com/document/path/95407
func (ww *weWork) DepartmentGet(corpId uint, id int32) (resp DepartmentGetResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		SetQueryParam("id", fmt.Sprintf("%v", id)).
		Get("/cgi-bin/department/get")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
