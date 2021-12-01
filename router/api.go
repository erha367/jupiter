package router

import (
	"github.com/gin-gonic/gin"
	"jupiter/application/controller"
	"jupiter/router/middleware"
)

func ApiRouter() (router *gin.Engine) {
	router = gin.New()
	//这里不走中间件
	router.GET("/ping", controller.Ping)
	/*- 中间件 -*/
	router.Use(middleware.Logger()) //日志记录
	router.Use(middleware.Cors())   //跨域
	router.Use(gin.Recovery())      //错误恢复
	//需要走中间件的写在下面

	return
}
