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

// CustomizedGetHandler 代开发自建应用回调配置
// /callback/customized?corpid=$CORPID$
func CustomizedGetHandler(c *gin.Context) {
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

func CustomizedPostHandler(c *gin.Context) {
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
				// TODO:根据不同企业配置不同的token和aeskey
				wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(ww.GetSuiteToken(), ww.GetSuiteEncodingAesKey(),
					bizData.ToUserName, wxbizmsgcrypt.XmlType)
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
