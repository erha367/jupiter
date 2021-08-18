package controller

import (
	"github.com/gin-gonic/gin"
	"jupiter/application/library"
)

//接口demo
func Ping(c *gin.Context) {
	name := c.PostForm(`name`)
	library.OkJson(c, name)
}
