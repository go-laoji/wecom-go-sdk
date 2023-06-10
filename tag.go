package wework

import (
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type Tag struct {
	TagId   int    `json:"tagid"`
	TagName string `json:"tagname" validate:"required,max=32"`
}

type TagCreateResponse struct {
	internal.BizResponse
	TagId int `json:"tagid"`
}

// TagCreate 创建标签
// https://open.work.weixin.qq.com/api/doc/90001/90143/90346
func (ww *weWork) TagCreate(corpId uint, tag Tag) (resp TagCreateResponse) {
	if ok := validate.Struct(tag); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(tag).SetResult(&resp).
		Post("/cgi-bin/tag/create")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// TagUpdate 更新标签名字
// https://open.work.weixin.qq.com/api/doc/90001/90143/90347
func (ww *weWork) TagUpdate(corpId uint, tag Tag) (resp internal.BizResponse) {
	if ok := validate.Struct(tag); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(tag).SetResult(&resp).
		Post("/cgi-bin/tag/update")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// TagDelete 删除标签
// https://open.work.weixin.qq.com/api/doc/90001/90143/90348
func (ww *weWork) TagDelete(corpId uint, id int) (resp internal.BizResponse) {
	_, err := ww.getRequest(corpId).SetQueryParam("tagid", fmt.Sprintf("%v", id)).SetResult(&resp).
		Get("/cgi-bin/tag/delete")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type TagListResponse struct {
	internal.BizResponse
	TagList []Tag `json:"taglist"`
}

// TagList 获取标签列表
// https://open.work.weixin.qq.com/api/doc/90001/90143/90352
func (ww *weWork) TagList(corpId uint) (resp TagListResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		Get("/cgi-bin/tag/list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type TagUserListResponse struct {
	internal.BizResponse
	TagName  string `json:"tagname"`
	UserList []struct {
		UserId string `json:"userid"`
		Name   string `json:"name"`
	} `json:"userlist"`
	PartyList []int32 `json:"partylist"`
}

// TagUserList 获取标签成员
// https://open.work.weixin.qq.com/api/doc/90001/90143/90349
func (ww *weWork) TagUserList(corpId uint, id int) (resp TagUserListResponse) {
	_, err := ww.getRequest(corpId).SetQueryParam("tagid", fmt.Sprintf("%v", id)).SetResult(&resp).
		Get("/cgi-bin/tag/get")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type TagAddOrDelUsersResponse struct {
	internal.BizResponse
	InvalidList  string  `json:"invalidlist,omitempty"`
	InvalidParty []int32 `json:"invalidparty,omitempty"`
}

// TagAddUsers 增加标签成员
// https://open.work.weixin.qq.com/api/doc/90001/90143/90350
func (ww *weWork) TagAddUsers(corpId uint, tagId int, userIds []string, partyIds []int32) (resp TagAddOrDelUsersResponse) {
	p := H{"tagid": tagId, "userlist": userIds, "partylist": partyIds}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/tag/addtagusers")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// TagDelUsers 删除标签成员
// https://open.work.weixin.qq.com/api/doc/90001/90143/90351
func (ww *weWork) TagDelUsers(corpId uint, tagId int, userIds []string, partyIds []int32) (resp TagAddOrDelUsersResponse) {
	p := H{"tagid": tagId, "userlist": userIds, "partylist": partyIds}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/tag/deltagusers")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
