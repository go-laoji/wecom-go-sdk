package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-laoji/wecom-go-sdk/v2"
)

func InjectSdk(ww wework.IWeWork) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ww", ww)
		c.Next()
	}
}
