package wework

import (
	"encoding/json"
	"github.com/go-laoji/wecom-go-sdk/v2/internal"
)

type Message struct {
	ToUser                 string `json:"touser,omitempty" validate:"omitempty,required_without=ToParty ToTag"`
	ToParty                string `json:"toparty,omitempty" validate:"omitempty,required_without=ToUser ToTag"`
	ToTag                  string `json:"totag,omitempty" validate:"omitempty,required_without=ToParty ToUser"`
	EnableIDTrans          int    `json:"enable_id_trans,omitempty"`
	EnableDuplicateCheck   int    `json:"enable_duplicate_check,omitempty"`
	DuplicateCheckInterval int    `json:"duplicate_check_interval,omitempty"`
}

type TextMessage struct {
	Message
	Safe int  `json:"safe,omitempty" validate:"omitempty,oneof=0 1"`
	Text Text `json:"text" validate:"required"`
}

type ImageMessage struct {
	Message
	Safe  int        `json:"safe,omitempty" validate:"omitempty,oneof=0 1"`
	Image MultiMedia `json:"image" validate:"required"`
}
type MultiMedia struct {
	MediaId string `json:"media_id" validate:"required"`
}
type VoiceMessage struct {
	Message
	Safe  int        `json:"safe,omitempty"`
	Voice MultiMedia `json:"voice" validate:"required"`
}

type VideoMessage struct {
	Message
	Safe  int   `json:"safe,omitempty" validate:"omitempty,oneof=0 1"`
	Video Video `json:"video" validate:"required"`
}

type FileMessage struct {
	Message
	Safe int        `json:"safe,omitempty"`
	File MultiMedia `json:"file" validate:"required"`
}

type TextCard struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Url         string `json:"url" validate:"required"`
	BtnTxt      string `json:"btntxt"`
}
type TextCardMessage struct {
	Message
	TextCard TextCard `json:"textcard" validate:"required"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
	AppId       string `json:"appid" validate:"required_without=Url,required_with=PagePath"`
	PagePath    string `json:"pagepath" validate:"required_with=AppId"`
}
type News struct {
	Articles []Article `json:"articles" validate:"required,max=8"`
}
type NewsMessage struct {
	Message
	News News `json:"news" validate:"required"`
}

type MpArticle struct {
	Title            string `json:"title" validate:"required"`
	ThumbMediaId     string `json:"thumb_media_id" validate:"required"`
	Author           string `json:"author,omitempty"`
	ContentSourceUrl string `json:"content_source_url,omitempty"`
	Content          string `json:"content" validate:"required"`
	Digest           string `json:"digest,omitempty"`
}

type MpNews struct {
	Articles []MpArticle `json:"articles" validate:"required"`
}

type MpNewsMessage struct {
	Message
	Safe   int    `json:"safe,omitempty" validate:"omitempty,oneof=0 1 2"`
	MpNews MpNews `json:"mpnews" validate:"required"`
}

type MarkDownMessage struct {
	Message
	MarkDown Text `json:"markdown" validate:"required"`
}

type MiniProgramNotice struct {
	Appid             string `json:"appid" validate:"required"`
	Page              string `json:"page"`
	Title             string `json:"title" validate:"required"`
	Description       string `json:"description"`
	EmphasisFirstItem bool   `json:"emphasis_first_item"`
	ContentItem       []struct {
		Key   string `json:"key" validate:"required"`
		Value string `json:"value" validate:"required"`
	} `json:"content_item"`
}

type MiniProgramMessage struct {
	Message
	MiniProgramNotice MiniProgramNotice `json:"miniprogram_notice"`
}

type MessageSendResponse struct {
	internal.BizResponse
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
	MsgId        string `json:"msgid"`
	ResponseCode string `json:"response_code"`
	//仅消息类型为“按钮交互型”，“投票选择型”和“多项选择型”的模板卡片消息返回
	//应用可使用response_code调用更新模版卡片消息接口，24小时内有效，且只能使用一次
}

// MessageSend 发送应用消息
// https://open.work.weixin.qq.com/api/doc/90001/90143/90372
func (ww *weWork) MessageSend(corpId uint, msg interface{}) (resp MessageSendResponse) {
	if ok := validate.Struct(msg); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	h := H{}
	buf, _ := json.Marshal(msg)
	json.Unmarshal(buf, &h)
	h["agentid"] = ww.GetAgentId(corpId)
	switch msg.(type) {
	case TextMessage:
		h["msgtype"] = "text"
	case ImageMessage:
		h["msgtype"] = "image"
	case VoiceMessage:
		h["msgtype"] = "voice"
	case VideoMessage:
		h["msgtype"] = "video"
	case FileMessage:
		h["msgtype"] = "file"
	case TextCardMessage:
		h["msgtype"] = "textcard"
	case NewsMessage:
		h["msgtype"] = "news"
	case MpNewsMessage:
		h["msgtype"] = "mpnews"
	case MarkDownMessage:
		h["msgtype"] = "markdown"
	case MiniProgramMessage:
		h["msgtype"] = "miniprogram_notice"
	case TemplateCardMessage:
		h["msgtype"] = "template_card"
	}

	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/message/send")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// MessageReCall 撤回应用消息
// https://open.work.weixin.qq.com/api/doc/90000/90135/94867
func (ww *weWork) MessageReCall(corpId uint, msgId string) (resp internal.BizResponse) {
	h := H{"msgid": msgId}
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/message/recall")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
