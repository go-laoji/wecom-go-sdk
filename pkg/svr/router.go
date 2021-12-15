package svr

import (
	"github.com/gin-gonic/gin"
	"github.com/go-laoji/wework"
	"github.com/go-laoji/wework/pkg/svr/callback"
	"github.com/go-laoji/wework/pkg/svr/install"
	"github.com/go-laoji/wework/pkg/svr/middleware"
)

func InjectRouter(ww wework.IWeWork) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.InjectSdk(ww))

	callbackGroup := router.Group("/callback")
	{
		callbackGroup.GET("/data", callback.DataGetHandler)
		callbackGroup.POST("/data", callback.DataPostHandler)
		callbackGroup.GET("/cmd", callback.CmdGetHandler)
		callbackGroup.POST("/cmd", callback.CmdPostHandler)
	}
	suite := router.Group("/suite")
	{
		suite.GET("/install", install.SuiteInstall)
		suite.GET("/install/auth", install.SuiteInstallAuth)
	}
	return router
}
