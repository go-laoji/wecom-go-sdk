package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/go-laoji/wecom-go-sdk/v2"
	"net/http"
)

func UserGet(c *gin.Context) {
	userId := c.Query("userid")
	if ww, ok := c.Keys["ww"].(wework.IWeWork); ok {
		resp := ww.UserGet(CorpId, userId)
		if resp.ErrCode == 0 {
			c.JSON(http.StatusOK, resp)
		} else {
			c.JSON(http.StatusInternalServerError, resp)
		}
	}
}

func UserSimpleList(c *gin.Context) {
	if ww, ok := c.Keys["ww"].(wework.IWeWork); ok {
		resp := ww.UserSimpleList(CorpId, 1, 1)
		if resp.ErrCode == 0 {
			c.JSON(http.StatusOK, resp)
		} else {
			c.JSON(http.StatusInternalServerError, resp)
		}
	}
}

func UserList(c *gin.Context) {
	if ww, ok := c.Keys["ww"].(wework.IWeWork); ok {
		resp := ww.UserList(CorpId, 1, 1)
		if resp.ErrCode == 0 {
			c.JSON(http.StatusOK, resp)
		} else {
			c.JSON(http.StatusInternalServerError, resp)
		}
	}
}

func UserId2OpenId(c *gin.Context) {
	if ww, ok := c.Keys["ww"].(wework.IWeWork); ok {
		userId := c.Query("userid")
		resp := ww.UserId2OpenId(CorpId, userId)
		if resp.ErrCode == 0 {
			c.JSON(http.StatusOK, resp)
		} else {
			c.JSON(http.StatusInternalServerError, resp)
		}
	}
}

func OpenId2UserId(c *gin.Context) {
	if ww, ok := c.Keys["ww"].(wework.IWeWork); ok {
		openId := c.Query("openid")
		resp := ww.OpenId2UserId(CorpId, openId)
		if resp.ErrCode == 0 {
			c.JSON(http.StatusOK, resp)
		} else {
			c.JSON(http.StatusInternalServerError, resp)
		}
	}
}

func ListMemberAuth(c *gin.Context) {
	if ww, ok := c.Keys["ww"].(wework.IWeWork); ok {
		resp := ww.ListMemberAuth(CorpId, "", 10)
		if resp.ErrCode == 0 {
			c.JSON(http.StatusOK, resp)
		} else {
			c.JSON(http.StatusInternalServerError, resp)
		}
	}
}
