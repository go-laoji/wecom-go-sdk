package wework

import (
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type KfKnowLedgeAddGroupResponse struct {
	internal.BizResponse
	GroupId string `json:"group_id"`
}

func (ww *weWork) KfKnowLedgeAddGroup(corpId uint, name string) (resp KfKnowLedgeAddGroupResponse) {
	p := H{"name": name}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/kf/knowledge/add_group")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

func (ww *weWork) KfKnowLedgeDelGroup(corpId uint, groupId string) (resp internal.BizResponse) {
	p := H{"group_id": groupId}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/kf/knowledge/del_group")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

func (ww *weWork) KfKnowLedgeModGroup(corpId uint, name string, groupId string) (resp internal.BizResponse) {
	p := H{"name": name, "group_id": groupId}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/kf/knowledge/mod_group")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type KfKnowLedgeListGroupFilter struct {
	Cursor  string `json:"cursor,omitempty"`
	Limit   uint32 `json:"limit,omitempty" validate:"omitempty,max=1000"`
	GroupId string `json:"group_id,omitempty"`
}
type KfKnowLedgeListGroupResponse struct {
	internal.BizResponse
	NextCursor string `json:"next_cursor"`
	HasMore    int    `json:"has_more"`
	GroupList  []struct {
		GroupID   string `json:"group_id"`
		Name      string `json:"name"`
		IsDefault int    `json:"is_default"`
	} `json:"group_list"`
}

func (ww *weWork) KfKnowLedgeListGroup(corpId uint, filter KfKnowLedgeListGroupFilter) (resp KfKnowLedgeListGroupResponse) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(filter).SetResult(&resp).
		Post("/cgi-bin/kf/knowledge/list_group")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
