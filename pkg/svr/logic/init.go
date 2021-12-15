package logic

import (
	"encoding/xml"
	"github.com/go-laoji/wework/config"
	"github.com/go-laoji/wework/pkg/svr/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type EventPushQueryBinding struct {
	MsgSign   string `form:"msg_signature" binding:"required"`
	Timestamp string `form:"timestamp" binding:"required"`
	Nonce     string `form:"nonce" binding:"required"`
	EchoStr   string `form:"echostr"`
}

type InfoType string

const (
	SuiteTicket InfoType = "suite_ticket"
	CreateAuth  InfoType = "create_auth"
	ChangeAuth  InfoType = "change_auth"
	CancelAuth  InfoType = "cancel_auth"
)

type BizEvent struct {
	XMLName   xml.Name `xml:"xml"`
	Text      string   `xml:",chardata"`
	SuiteId   string   `xml:"SuiteId"`
	InfoType  InfoType `xml:"InfoType"`
	TimeStamp int64    `xml:"TimeStamp"`
}

// 数据回调通知
// 企业安装应用时会把相应数据通知到回调URL
type BizData struct {
	xml.Name   `xml:"xml"`
	Text       string `xml:",chardata"`
	ToUserName string `xml:"ToUserName"`
	Encrypt    string `xml:"Encrypt"`
	AgentID    string `xml:"AgentID"`
}

var engine *gorm.DB

func init() {
	var err error
	c := config.ParseFile("")
	engine, err = gorm.Open(mysql.Open(c.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	engine.AutoMigrate(
		&models.CorpAccessToken{},
		&models.CorpPermanentCode{},
		&models.CorpAuthInfo{},
		&models.CorpAuthUserInfo{})
}
