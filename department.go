package wework

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wework/internal"
)

type Department struct {
	Id               int32    `json:"id"`
	Order            int      `json:"order,omitempty"`
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
