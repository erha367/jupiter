package router

import (
	"github.com/gin-gonic/gin"
	"jupiter/application/controller"
	"jupiter/router/middleware"
)

func ApiRouter() (router *gin.Engine) {
	router = gin.New()
	/*- 中间件 -*/
	router.Use(middleware.Logger()) //日志记录
	router.Use(middleware.Cors())   //跨域
	router.Use(gin.Recovery())      //错误恢复
	//demo
	router.GET("/ping", middleware.Sign(), controller.Ping)
	router.POST("/sign", middleware.Sign(), controller.Ping)
	router.GET("/login", controller.Login)
	router.GET("/jwt", middleware.Jwtm(), controller.Jwt)
	return
}
