package demo

import "github.com/gin-gonic/gin"

const CorpId uint = 1

func InjectRouter(e *gin.Engine) {
	demo := e.Group("/api/demo")
	demo.GET("/user", UserGet)
	demo.GET("/usersimplelist", UserSimpleList)
	demo.GET("/userlist", UserList)
	demo.GET("/userid2openid", UserId2OpenId)
	demo.GET("/openid2userid", OpenId2UserId)
	demo.GET("/listmemberauth", ListMemberAuth)
}
