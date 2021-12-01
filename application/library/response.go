package library

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const OK = 1
const NO = 0

func OkJson(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": OK,
		"msg":  "OK",
		"data": data,
	})
}

func FailJson(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code": NO,
		"msg":  err.Error(),
	})
}
