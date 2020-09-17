package router

import (
	"github.com/gin-gonic/gin"
	"jupiter/application/controller"
	"jupiter/router/middleware"
)

func ApiRouter() (router *gin.Engine) {
	router = gin.New()
	//中间件
	router.Use(middleware.Logger())
	router.Use(middleware.Cors())
	router.Use(gin.Recovery())
	//demo
	router.GET("/ping", controller.Ping)
	//静态资源
	router.Static("/static", "./static")
	return
}
