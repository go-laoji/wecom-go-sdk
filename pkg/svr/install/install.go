package install

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-laoji/wecom-go-sdk/v2"
	"log"
	"net/http"
	"net/url"
)

// SuiteInstall TODO:测试上线后的应用安装
func SuiteInstall(c *gin.Context) {
	if ww, ok := c.Keys["ww"].(wework.IWeWork); ok {
		resp := ww.GetPreAuthCode()
		if resp.ErrCode != 0 {
			c.JSON(http.StatusFailedDependency, gin.H{"errno": resp.ErrCode, "errmsg": resp.ErrorMsg})
		} else {
			ww.Logger().Sugar().Info(c.Request.URL)
			installAuth := url.QueryEscape(fmt.Sprintf("https://%s/suite/install/auth", c.Request.Host))
			redirectUrl := fmt.Sprintf("https://open.work.weixin.qq.com/3rdapp/install?suite_id=%s&pre_auth_code=%s&redirect_uri=%s&state=STATE",
				ww.GetSuiteId(), resp.PreAuthCode, installAuth)
			c.Redirect(http.StatusTemporaryRedirect, redirectUrl)
		}
	} else {
		log.Println("suite 未注入")
		c.JSON(http.StatusInternalServerError, gin.H{"errno": 500, "errmsg": "suite未注入"})
	}
}
