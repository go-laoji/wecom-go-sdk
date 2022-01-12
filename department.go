package wework

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wework/internal"
)

type Department struct {
	Id               int32    `json:"id"`
	Order            int32    `json:"order,omitempty"`
	ParentId         int32    `json:"parentid" validate:"required"`
	Name             string   `json:"name" validate:"required,min=1,max=32"`
	NameEn           string   `json:"name_en,omitempty" validate:"omitempty,min=1,max=32"`
	DepartmentLeader []string `json:"department_leader"`
}

type DepartmentListResponse struct {
	internal.BizResponse
	Department []Department `json:"department"`
}

// DepartmentList 获取部门列表
// https://open.work.weixin.qq.com/api/doc/90001/90143/90344
func (ww weWork) DepartmentList(corpId uint, id uint) (resp DepartmentListResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	queryParams.Add("id", fmt.Sprintf("%v", id))
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/department/list?%s",
		queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type DepartmentSimpleListResponse struct {
	internal.BizResponse
	DepartmentId []struct {
		Id       int32 `json:"id"`
		ParentId int32 `json:"parentid"`
		Order    int32 `json:"order"`
	}
}

// DepartmentSimpleList 获取子部门ID列表
// https://developer.work.weixin.qq.com/document/path/95406
func (ww weWork) DepartmentSimpleList(corpId uint, id int32) (resp DepartmentSimpleListResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	queryParams.Add("id", fmt.Sprintf("%v", id))
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/department/simplelist?%s",
		queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type DepartmentGetResponse struct {
	internal.BizResponse
	Department Department `json:"department"`
}

// DepartmentGet 获取单个部门详情
// https://developer.work.weixin.qq.com/document/path/95407
func (ww weWork) DepartmentGet(corpId uint, id int32) (resp DepartmentGetResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	queryParams.Add("id", fmt.Sprintf("%v", id))
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/department/get?%s",
		queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
