package callback

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/go-laoji/wecom-go-sdk/v2"
	"github.com/go-laoji/wecom-go-sdk/v2/pkg/svr/logic"
	"github.com/go-laoji/wxbizmsgcrypt"
	"io/ioutil"
	"log"
	"net/http"
)

// 指令回调配置

func CmdGetHandler(c *gin.Context) {
	if ww, exists := c.Keys["ww"].(wework.IWeWork); exists {
		var params logic.EventPushQueryBinding
		if ok := c.ShouldBindQuery(&params); ok == nil {
			receiveId := params.CorpId
			if receiveId == "" {
				receiveId = ww.GetCorpId()
			}
			wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(ww.GetSuiteToken(), ww.GetSuiteEncodingAesKey(),
				receiveId, wxbizmsgcrypt.XmlType)
			echoStr, cryptErr := wxcpt.VerifyURL(params.MsgSign, params.Timestamp, params.Nonce, params.EchoStr)
			if nil != cryptErr {
				ww.Logger().Sugar().Error(cryptErr)
				c.JSON(http.StatusLocked, gin.H{"err": cryptErr, "echoStr": echoStr})
			} else {
				c.Writer.Write(echoStr)
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"errno": 500, "errmsg": "no echostr"})
		}
	}
}

func CmdPostHandler(c *gin.Context) {
	if ww, exists := c.Keys["ww"].(wework.IWeWork); exists {
		var params logic.EventPushQueryBinding
		if ok := c.ShouldBindQuery(&params); ok == nil {
			body, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				ww.Logger().Sugar().Error(err)
				c.JSON(http.StatusOK, gin.H{"errno": 500, "errmsg": err.Error()})
				return
			} else {
				wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(ww.GetSuiteToken(), ww.GetSuiteEncodingAesKey(),
					ww.GetSuiteId(), wxbizmsgcrypt.XmlType)
				if msg, err := wxcpt.DecryptMsg(params.MsgSign, params.Timestamp, params.Nonce, body); err != nil {
					ww.Logger().Sugar().Error(err)
					c.JSON(http.StatusOK, gin.H{"errno": 500, "errmsg": err.ErrMsg})
					return
				} else {
					ww.Logger().Sugar().Info(string(msg))
					var bizEvent logic.BizEvent
					if e := xml.Unmarshal(msg, &bizEvent); e != nil {
						ww.Logger().Sugar().Error(e)
						c.JSON(http.StatusOK, gin.H{"errno": 500, "errmsg": err.ErrMsg})
						return
					}
					switch bizEvent.InfoType {
					case logic.SuiteTicket:
						go logic.SuiteTicketEventLogic(msg, ww)
						break
					case logic.CreateAuth:
						// 服务商的响应必须在1000ms内完成，以保证用户安装应用的体验。
						// 建议在接收到此事件时，先记录下AuthCode，并立即回应企业微信，之后再做相关业务的处理。
						go logic.CreateAuthEventLogic(msg, ww)
						break
					case logic.CancelAuth:
						go logic.CancelAuthEventLogic(msg, ww)
						break
					case logic.ResetPermanentCode:
						go logic.ResetPermanentCodeEventLogic(msg, ww)
						break
					}
					c.Writer.WriteString("success")
				}
			}
		} else {
			ww.Logger().Sugar().Error(ok)
			c.JSON(http.StatusOK, gin.H{"errno": 400, "errmsg": ok.Error()})
		}
	} else {
		log.Println("suite未注册")
		c.JSON(http.StatusOK, gin.H{"errno": 404, "errmsg": ""})
	}
}
