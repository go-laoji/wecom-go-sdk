package wework

import (
	"encoding/json"
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type TemplateCardType string

const (
	CardTypeTextNotice          TemplateCardType = "text_notice"
	CardTypeNewsNotice          TemplateCardType = "news_notice"
	CardTypeButtonInteraction   TemplateCardType = "button_interaction"
	CardTypeVoteInteraction     TemplateCardType = "vote_interaction"
	CardTypeMultipleInteraction TemplateCardType = "multiple_interaction"
)

// TemplateCardMessage 测试发送模板卡片消息必需配置应用回调地址
type TemplateCardMessage struct {
	Message
	TemplateCard TemplateCard `json:"template_card"`
}

type Source struct {
	IconURL   string `json:"icon_url"`
	Desc      string `json:"desc"`
	DescColor int    `json:"desc_color"`
}
type ActionList struct {
	Text string `json:"text" validate:"required"`
	Key  string `json:"key" validate:"required"`
}
type ActionMenu struct {
	Desc       string       `json:"desc"`
	ActionList []ActionList `json:"action_list" validate:"required,max=3,min=1"`
}
type MainTitle struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type QuoteArea struct {
	Type      int    `json:"type"`
	URL       string `json:"url"`
	Title     string `json:"title"`
	QuoteText string `json:"quote_text"`
}

// EmphasisContent 文本通知型
type EmphasisContent struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type HorizontalContentList struct {
	KeyName string `json:"keyname"`
	Value   string `json:"value"`
	Type    int    `json:"type,omitempty"`
	URL     string `json:"url,omitempty"`
	MediaID string `json:"media_id,omitempty"`
	Userid  string `json:"userid,omitempty"`
}
type JumpList struct {
	Type     int    `json:"type"`
	Title    string `json:"title"`
	URL      string `json:"url,omitempty"`
	Appid    string `json:"appid,omitempty"`
	PagePath string `json:"pagepath,omitempty"`
}
type CardAction struct {
	Type     int    `json:"type"`
	URL      string `json:"url"`
	Appid    string `json:"appid"`
	Pagepath string `json:"pagepath"`
}

// ImageTextArea 图文展示型
type ImageTextArea struct {
	Type     int    `json:"type" validate:"omitempty,oneof=0 1 2"`
	URL      string `json:"url"`
	AppId    string `json:"appid,omitempty"`
	PagePath string `json:"pagepath,omitempty"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	ImageURL string `json:"image_url" validate:"required"`
}
type CardImage struct {
	Url         string  `json:"url" validate:"required"`
	AspectRatio float32 `json:"aspect_ratio" validate:"max=2.25,min=1.3"`
}

// ButtonSelection 按钮交互型
type ButtonSelection struct {
	QuestionKey string `json:"question_key" validate:"required"`
	Title       string `json:"title"`
	OptionList  []struct {
		ID   string `json:"id" validate:"required"`
		Text string `json:"text" validate:"required"`
	} `json:"option_list" validate:"required"`
	SelectedID string `json:"selected_id"`
}
type Button struct {
	Type  int    `json:"type,omitempty"` //按钮点击事件类型，0 或不填代表回调点击事件，1 代表跳转url
	Text  string `json:"text" validate:"required"`
	Style int    `json:"style,omitempty"` //按钮样式，目前可填1~4，不填或错填默认1
	Key   string `json:"key,omitempty"`   // 按钮key值，用户点击后，会产生回调事件将本参数作为EventKey返回，回调事件会带上该key值，最长支持1024字节，不可重复，button_list.type是0时必填
	Url   string `json:"url,omitempty"`   //跳转事件的url，button_list.type是1时必填
}

// CheckBox 投票选择型
type CheckBox struct {
	QuestionKey string `json:"question_key" validate:"required"`
	OptionList  []struct {
		ID        string `json:"id" validate:"required"`
		Text      string `json:"text" validate:"required"`
		IsChecked bool   `json:"is_checked" validate:"required"`
	} `json:"option_list" validate:"required,min=1,max=20"`
	Mode int `json:"mode" validate:"omitempty,oneof=0 1"`
}
type SubmitButton struct {
	Text string `json:"text" validate:"required"`
	Key  string `json:"key" validate:"required"`
}

// SelectList 多项选择型
type SelectList struct {
	QuestionKey string `json:"question_key" validate:"required"`
	Title       string `json:"title,omitempty"`
	SelectedID  string `json:"selected_id,omitempty"`
	OptionList  []struct {
		ID   string `json:"id" validate:"required"`
		Text string `json:"text" validate:"required"`
	} `json:"option_list" validate:"required"`
}

// TODO: CardAction 必填(text_notice,news_notice)判断
type TemplateCard struct {
	CardType   TemplateCardType `json:"card_type"`
	Source     Source           `json:"source"`
	ActionMenu *ActionMenu      `json:"action_menu,omitempty" validate:"required_with=TaskID"`
	TaskID     string           `json:"task_id,omitempty" validate:"required_with=ActionMenu"`
	MainTitle  MainTitle        `json:"main_title"`
	QuoteArea  QuoteArea        `json:"quote_area"`
	// 文本通知型
	EmphasisContent *EmphasisContent `json:"emphasis_content,omitempty"`
	SubTitleText    string           `json:"sub_title_text,omitempty"`
	// 图文展示型
	ImageTextArea         *ImageTextArea          `json:"image_text_area,omitempty"`
	CardImage             *CardImage              `json:"card_image,omitempty"`
	HorizontalContentList []HorizontalContentList `json:"horizontal_content_list"`
	JumpList              []JumpList              `json:"jump_list"`
	CardAction            CardAction              `json:"card_action,omitempty"`
	// 按钮交互型
	ButtonSelection *ButtonSelection `json:"button_selection,omitempty"`
	ButtonList      []Button         `json:"button_list,omitempty" validate:"omitempty,max=6"`
	// 投票选择型
	CheckBox     *CheckBox     `json:"checkbox,omitempty"`
	SelectList   []SelectList  `json:"select_list,omitempty" validate:"max=3"`
	SubmitButton *SubmitButton `json:"submit_button,omitempty"`
}

type TemplateCardUpdateMessage struct {
	UserIds      []string `json:"userids" validate:"omitempty,max=100"`
	PartyIds     []int64  `json:"partyids" validate:"omitempty,max=100"`
	TagIds       []int32  `json:"tagids" validate:"omitempty,max=100"`
	AtAll        int      `json:"atall,omitempty"`
	ResponseCode string   `json:"response_code" validate:"required"`
	Button       struct {
		ReplaceName string `json:"replace_name" validate:"required"`
	} `json:"button" validate:"required_without=TemplateCard"`
	TemplateCard TemplateCard `json:"template_card" validate:"required_without=Button"`
	ReplaceText  string       `json:"replace_text,omitempty"`
}

type MessageUpdateTemplateCardResponse struct {
	internal.BizResponse
	InvalidUser []string `json:"invalid_user"`
}

// MessageUpdateTemplateCard 更新模板卡片消息
// https://open.work.weixin.qq.com/api/doc/90001/90143/94945
func (ww *weWork) MessageUpdateTemplateCard(corpId uint, msg TemplateCardUpdateMessage) (resp MessageUpdateTemplateCardResponse) {
	if ok := validate.Struct(msg); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	h := H{}
	buf, _ := json.Marshal(msg)
	json.Unmarshal(buf, &h)
	h["agentid"] = ww.GetAgentId(corpId)
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/message/update_template_card")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
