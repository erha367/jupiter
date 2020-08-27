package controller

import (
	"github.com/gin-gonic/gin"
	"jupyter/application/library"
)

//接口demo
func Ping(c *gin.Context) {
	library.OkJson(c, `Pong`)
}
