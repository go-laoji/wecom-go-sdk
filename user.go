package wework

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wecom-go-sdk/internal"
	"strings"
)

type User struct {
	OpenUserId     string      `json:"open_userid,omitempty"` // 仅在查询时返回
	Userid         string      `json:"userid" validate:"required"`
	Name           string      `json:"name,omitempty" validate:"required"`
	Alias          string      `json:"alias,omitempty"`
	Mobile         string      `json:"mobile,omitempty"  validate:"required_without=Email,omitempty"`
	Department     []int32     `json:"department,omitempty" validate:"required,max=100"`
	Order          []int32     `json:"order,omitempty"`
	Position       string      `json:"position,omitempty"`
	Gender         string      `json:"gender,omitempty" validate:"omitempty,oneof=1 2"`
	Email          string      `json:"email,omitempty"  validate:"required_without=Mobile,omitempty,email"`
	BizEmail       string      `json:"biz_email,omitempty"`
	IsLeaderInDept []int       `json:"is_leader_in_dept,omitempty"`
	DirectLeader   []string    `json:"direct_leader,omitempty"`
	Enable         json.Number `json:"enable,omitempty"`
	Avatar         string      `json:"avatar,omitempty"`
	ThumbAvatar    string      `json:"thumb_avatar,omitempty"`
	Telephone      string      `json:"telephone,omitempty"`
	Address        string      `json:"address,omitempty"`
	MainDepartment int32       `json:"main_department,omitempty"`
	Status         int         `json:"status,omitempty"`
	QrCode         string      `json:"qr_code,omitempty"`
	Extattr        *struct {
		Attrs []Attrs `json:"attrs,omitempty"`
	} `json:"extattr,omitempty"`
	ToInvite         bool   `json:"to_invite,omitempty"`
	ExternalPosition string `json:"external_position,omitempty"`
	ExternalProfile  *struct {
		ExternalCorpName string `json:"external_corp_name,omitempty"`
		WechatChannels   struct {
			Nickname string `json:"nickname,omitempty"`
		} `json:"wechat_channels,omitempty"`
		ExternalAttr []ExternalAttr `json:"external_attr,omitempty"`
	} `json:"external_profile,omitempty"`
}
type Attrs struct {
	Type int    `json:"type" validate:"required,oneof= 0 1 2"`
	Name string `json:"name" validate:"required"`
	Text struct {
		Value string `json:"value"`
	} `json:"text,omitempty"`
	Web struct {
		URL   string `json:"url" validate:"required"`
		Title string `json:"title" validate:"required"`
	} `json:"web,omitempty"`
}
type ExternalAttr struct {
	Type int    `json:"type"`
	Name string `json:"name"`
	Text struct {
		Value string `json:"value"`
	} `json:"text,omitempty"`
	Web struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	} `json:"web,omitempty"`
	Miniprogram struct {
		Appid    string `json:"appid"`
		Pagepath string `json:"pagepath"`
		Title    string `json:"title"`
	} `json:"miniprogram,omitempty"`
}

// UserCreate 创建成员
func (ww weWork) UserCreate(corpId uint, user User) (resp internal.BizResponse) {
	if ok := validate.Struct(user); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/user/create?%s", queryParams.Encode()), user)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type UserGetResponse struct {
	internal.BizResponse
	User
}

// UserGet 读取成员
func (ww weWork) UserGet(corpId uint, userId string) (resp UserGetResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	queryParams.Add("userid", userId)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/user/get?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// UserUpdate 更新成员
func (ww weWork) UserUpdate(corpId uint, user User) (resp internal.BizResponse) {
	if strings.TrimSpace(user.Userid) == "" {
		resp.ErrCode = 500
		resp.ErrorMsg = "userid can not empty"
		return
	}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/user/update?%s", queryParams.Encode()), user)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// UserDelete 删除成员
func (ww weWork) UserDelete(corpId uint, userId string) (resp internal.BizResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	queryParams.Add("userid", userId)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/user/delete?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type UserSimpleListResponse struct {
	internal.Error
	UserList []struct {
		UserId     string `json:"userid"`
		Name       string `json:"name"`
		Department []int  `json:"department"`
		OpenUserId string `json:"open_userid"`
	} `json:"userlist"`
}

// UserSimpleList 获取部门成员
// https://open.work.weixin.qq.com/api/doc/90001/90143/90332
func (ww weWork) UserSimpleList(corpId uint, departId int32, fetchChild int) (resp UserSimpleListResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	if departId <= 0 {
		return UserSimpleListResponse{internal.Error{ErrorMsg: "部门ID必需大于0", ErrCode: 403}, nil}
	}
	queryParams.Add("department_id", fmt.Sprintf("%v", departId))
	queryParams.Add("fetch_child", fmt.Sprintf("%v", fetchChild))
	apiUrl := fmt.Sprintf("/cgi-bin/user/simplelist?%s", queryParams.Encode())
	body, err := internal.HttpGet(apiUrl)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
		return
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
		return
	}
	return
}

type UserListResponse struct {
	internal.BizResponse
	UserList []User `json:"userlist"`
}

// UserList 获取部门成员详情
// https://open.work.weixin.qq.com/api/doc/90001/90143/90337
func (ww weWork) UserList(corpId uint, departId int32, fetchChild int) (resp UserListResponse) {
	queryParams := ww.buildCorpQueryToken(corpId)
	if departId <= 0 {
		resp.ErrCode = 403
		resp.ErrorMsg = "部门ID必需大于0"
		return
	}
	queryParams.Add("department_id", fmt.Sprintf("%v", departId))
	queryParams.Add("fetch_child", fmt.Sprintf("%v", fetchChild))
	apiUrl := fmt.Sprintf("/cgi-bin/user/list?%s", queryParams.Encode())
	body, err := internal.HttpGet(apiUrl)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type UserId2OpenIdResponse struct {
	internal.BizResponse
	OpenId string `json:"openid"`
}

// UserId2OpenId userid与openid互换
// https://open.work.weixin.qq.com/api/doc/90001/90143/90338
func (ww weWork) UserId2OpenId(corpId uint, userId string) (resp UserId2OpenIdResponse) {
	p := H{"userid": userId}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/user/convert_to_openid?%s",
		queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type OpenId2UserIdResponse struct {
	internal.BizResponse
	UserId string `json:"userid"`
}

// OpenId2UserId openid转userid
// https://open.work.weixin.qq.com/api/doc/90001/90143/90338
func (ww weWork) OpenId2UserId(corpId uint, openId string) (resp OpenId2UserIdResponse) {
	p := H{"openid": openId}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/user/convert_to_userid?%s",
		queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type ListMemberAuthResponse struct {
	internal.BizResponse
	NextCursor     string `json:"next_cursor"`
	MemberAuthList []struct {
		OpenUserId string `json:"open_userid"`
	} `json:"member_auth_list"`
}

// ListMemberAuth 获取成员授权列表
// https://open.work.weixin.qq.com/api/doc/90001/90143/94513
func (ww weWork) ListMemberAuth(corpId uint, cursor string, limit int) (resp ListMemberAuthResponse) {
	p := H{"cursor": cursor, "limit": limit}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/user/list_member_auth?%s",
		queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type CheckMemberAuthResponse struct {
	internal.BizResponse
	IsMemberAuth bool `json:"is_member_auth"`
}

// CheckMemberAuth 查询成员用户是否已授权
// https://open.work.weixin.qq.com/api/doc/90001/90143/94514
func (ww weWork) CheckMemberAuth(corpId uint, openUserId string) (resp CheckMemberAuthResponse) {
	p := H{"open_userid": openUserId}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/user/check_member_auth?%s",
		queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type GetUserIdResponse struct {
	internal.BizResponse
	UserId string `json:"userid"`
}

// GetUserId 手机号获取userid
// https://open.work.weixin.qq.com/api/doc/90001/90143/91693
func (ww weWork) GetUserId(corpId uint, mobile string) (resp GetUserIdResponse) {
	p := H{"mobile": mobile}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/user/getuserid?%s",
		queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type ListSelectedTicketUserResponse struct {
	internal.BizResponse
	OperatorOpenUserId   string   `json:"operator_open_userid"`
	OpenUserIdList       []string `json:"open_userid_list"`
	UnAuthOpenUserIdList []string `json:"unauth_open_userid_list"`
	Total                int      `json:"total"`
}

// ListSelectedTicketUser 获取选人ticket对应的用户
// https://open.work.weixin.qq.com/api/doc/90001/90143/94894
func (ww weWork) ListSelectedTicketUser(corpId uint, ticket string) (resp ListSelectedTicketUserResponse) {
	p := H{"selected_ticket": ticket}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/user/list_selected_ticket_user?%s",
		queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type UserListIdResponse struct {
	internal.BizResponse
	NextCursor string `json:"next_cursor"`
	DeptUser   []struct {
		UserId     string `json:"userid"`
		OpenUserId string `json:"open_userid"`
		Department int    `json:"department"`
	} `json:"dept_user"`
}

// UserListId 获取成员ID列表 仅支持通过“通讯录同步secret”调用。
// https://developer.work.weixin.qq.com/document/40856
func (ww weWork) UserListId(corpId uint, cursor string, limit int) (resp UserListIdResponse) {
	p := H{"cursor": cursor, "limit": limit}
	queryParams := ww.buildCorpQueryToken(corpId)
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/user/list_id?%s", queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
