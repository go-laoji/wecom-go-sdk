package wework

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/internal"
)

type KfKnowLedgeAddGroupResponse struct {
	internal.BizResponse
	GroupId string `json:"group_id"`
}

func (ww weWork) KfKnowLedgeAddGroup(corpId uint, name string) (resp KfKnowLedgeAddGroupResponse) {
	p := H{"name": name}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/kf/knowledge/add_group?%s",
		queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

func (ww weWork) KfKnowLedgeDelGroup(corpId uint, groupId string) (resp internal.BizResponse) {
	p := H{"group_id": groupId}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/kf/knowledge/del_group?%s",
		queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

func (ww weWork) KfKnowLedgeModGroup(corpId uint, name string, groupId string) (resp internal.BizResponse) {
	p := H{"name": name, "group_id": groupId}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/kf/knowledge/mod_group?%s",
		queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
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

func (ww weWork) KfKnowLedgeListGroup(corpId uint, filter KfKnowLedgeListGroupFilter) (resp KfKnowLedgeListGroupResponse) {
	if ok := validate.Struct(filter); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/kf/knowledge/list_group?%s", queryParams.Encode()), filter)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
