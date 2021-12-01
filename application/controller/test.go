package controller

import (
	"github.com/gin-gonic/gin"
	"jupiter/library"
)

//接口demo
func Ping(c *gin.Context) {
	library.Success(c, `pong`)
}
