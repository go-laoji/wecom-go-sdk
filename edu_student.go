package wework

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/internal"
	"strings"
)

type Student struct {
	StudentUserId    string `json:"student_userid" validate:"required"`
	Name             string `json:"name,omitempty" validate:"required"`
	Department       []uint `json:"department,omitempty" validate:"required,max=20"`
	NewStudentUserId string `json:"new_student_userid,omitempty"`
}

// CreateStudent 创建学生
// https://open.work.weixin.qq.com/api/doc/90001/90143/92035
func (ww weWork) CreateStudent(corpId uint, student Student) (resp internal.BizResponse) {
	if ok := validate.Struct(student); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/school/user/create_student?%s", queryParams.Encode()), student)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type BatchStudentResponse struct {
	internal.BizResponse
	ResultList []struct {
		internal.BizResponse
		StudentUserId string `json:"student_userid"`
	} `json:"result_list"`
}

// BatchCreateStudent 批量创建学生
// https://open.work.weixin.qq.com/api/doc/90001/90143/92037
func (ww weWork) BatchCreateStudent(corpId uint, students []Student) (resp BatchStudentResponse) {
	h := H{}
	h["students"] = students
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/school/user/batch_create_student?%s", queryParams.Encode()), h)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// DeleteStudent 删除学生
// https://open.work.weixin.qq.com/api/doc/90001/90143/92039
func (ww weWork) DeleteStudent(corpId uint, userId string) (resp internal.BizResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	queryParams.Add("userid", userId)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/school/user/delete_student?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// BatchDeleteStudent 批量删除学生
// https://open.work.weixin.qq.com/api/doc/90001/90143/92040
func (ww weWork) BatchDeleteStudent(corpId uint, userIdList []string) (resp BatchStudentResponse) {
	h := H{}
	h["useridlist"] = userIdList
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/school/user/batch_delete_student?%s", queryParams.Encode()), h)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// UpdateStudent 更新学生
// https://open.work.weixin.qq.com/api/doc/90001/90143/92041
func (ww weWork) UpdateStudent(corpId uint, student Student) (resp internal.BizResponse) {
	if strings.TrimSpace(student.StudentUserId) == "" {
		resp.ErrCode = 500
		resp.ErrorMsg = "student id can not be empty"
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/school/user/update_student?%s", queryParams.Encode()), student)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// BatchUpdateStudent 批量更新学生
// https://open.work.weixin.qq.com/api/doc/90001/90143/92042
func (ww weWork) BatchUpdateStudent(corpId uint, students []Student) (resp BatchStudentResponse) {
	h := H{}
	h["students"] = students
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/school/user/batch_update_student?%s", queryParams.Encode()), h)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
