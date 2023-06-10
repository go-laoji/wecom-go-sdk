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

func DataGetHandler(c *gin.Context) {
	if ww, exists := c.Keys["ww"].(wework.IWeWork); exists {
		wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(ww.GetSuiteToken(), ww.GetSuiteEncodingAesKey(),
			ww.GetCorpId(), wxbizmsgcrypt.XmlType)
		var params logic.EventPushQueryBinding
		if ok := c.ShouldBindQuery(&params); ok == nil {
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

func DataPostHandler(c *gin.Context) {
	if ww, exists := c.Keys["ww"].(wework.IWeWork); exists {
		var params logic.EventPushQueryBinding
		if ok := c.ShouldBindQuery(&params); ok == nil {
			body, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				ww.Logger().Sugar().Error(err)
				c.JSON(http.StatusOK, gin.H{"errno": 500, "errmsg": err.Error()})
				return
			} else {
				var bizData logic.BizData
				xml.Unmarshal(body, &bizData)
				wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(ww.GetSuiteToken(), ww.GetSuiteEncodingAesKey(),
					bizData.ToUserName, wxbizmsgcrypt.XmlType)
				if msg, err := wxcpt.DecryptMsg(params.MsgSign, params.Timestamp, params.Nonce, body); err != nil {
					ww.Logger().Sugar().Error(err)
					c.JSON(http.StatusOK, gin.H{"errno": 500, "errmsg": err.ErrMsg})
					return
				} else {
					ww.Logger().Sugar().Info(string(msg))
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
