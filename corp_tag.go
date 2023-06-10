package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type CorpTagGroup struct {
	GroupId    string    `json:"group_id"`
	GroupName  string    `json:"group_name"`
	CreateTime uint64    `json:"create_time,omitempty"`
	Order      int       `json:"order,omitempty"`
	Deleted    bool      `json:"deleted,omitempty"`
	Tag        []CorpTag `json:"tag"`
}

type CorpTag struct {
	Id         string `json:"id,omitempty"`
	Name       string `json:"name" validate:"required,max=30"`
	Order      int32  `json:"order"`
	CreateTime uint64 `json:"create_time,omitempty"`
	Deleted    bool   `json:"deleted,omitempty"`
}

type CorpTagListResponse struct {
	internal.BizResponse
	TagGroup []CorpTagGroup `json:"tag_group"`
}

// CorpTagList 若tag_id和group_id均为空，则返回所有标签。
// 同时传递tag_id和group_id时，忽略tag_id，仅以group_id作为过滤条件。
// https://open.work.weixin.qq.com/api/doc/90001/90143/92696#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E6%A0%87%E7%AD%BE%E5%BA%93
func (ww *weWork) CorpTagList(corpId uint, tagIds, groupIds []string) (resp CorpTagListResponse) {
	p := H{"tag_id": tagIds, "group_id": groupIds}
	_, err := ww.getRequest(corpId).
		SetBody(p).SetResult(&resp).
		Post("/cgi-bin/externalcontact/get_corp_tag_list")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type CorpTagAddResponse struct {
	internal.BizResponse
	TagGroup CorpTagGroup `json:"tag_group"`
}

// CorpTagAdd 企业可通过此接口向客户标签库中添加新的标签组和标签，每个企业最多可配置3000个企业标签。
// https://open.work.weixin.qq.com/api/doc/90001/90143/92696#%E6%B7%BB%E5%8A%A0%E4%BC%81%E4%B8%9A%E5%AE%A2%E6%88%B7%E6%A0%87%E7%AD%BE
func (ww *weWork) CorpTagAdd(corpId uint, tagGroup CorpTagGroup) (resp CorpTagAddResponse) {
	_, err := ww.getRequest(corpId).
		SetBody(tagGroup).SetResult(&resp).
		Post("/cgi-bin/externalcontact/add_corp_tag")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// CorpTagUpdate 企业可通过此接口编辑客户标签/标签组的名称或次序值。
// https://open.work.weixin.qq.com/api/doc/90001/90143/92696#%E7%BC%96%E8%BE%91%E4%BC%81%E4%B8%9A%E5%AE%A2%E6%88%B7%E6%A0%87%E7%AD%BE
func (ww *weWork) CorpTagUpdate(corpId uint, tag CorpTag) (resp internal.BizResponse) {
	_, err := ww.getRequest(corpId).
		SetBody(tag).SetResult(&resp).
		Post("/cgi-bin/externalcontact/edit_corp_tag")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// CorpTagDelete 企业可通过此接口删除客户标签库中的标签，或删除整个标签组。
// https://open.work.weixin.qq.com/api/doc/90001/90143/92696#%E5%88%A0%E9%99%A4%E4%BC%81%E4%B8%9A%E5%AE%A2%E6%88%B7%E6%A0%87%E7%AD%BE
func (ww *weWork) CorpTagDelete(corpId uint, tagIds, groupIds []string) (resp internal.BizResponse) {
	p := H{"tag_id": tagIds, "group_id": groupIds}
	_, err := ww.getRequest(corpId).
		SetBody(p).SetResult(&resp).
		Post("/cgi-bin/externalcontact/del_corp_tag")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// MarkTag 编辑客户企业标签
// https://open.work.weixin.qq.com/api/doc/90001/90143/92697
func (ww *weWork) MarkTag(corpId uint, userId string, externalUserId string, addTag []int, removeTag []int) (resp internal.BizResponse) {
	p := H{"userid": userId, "external_userid": externalUserId, "add_tag": addTag, "remove_tag": removeTag}
	_, err := ww.getRequest(corpId).
		SetBody(p).SetResult(&resp).
		Post("/cgi-bin/externalcontact/mark_tag")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
