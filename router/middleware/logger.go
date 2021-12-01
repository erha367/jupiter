package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	UUIDS "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"jupiter/library"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		//全局的uuid，默认取自 header
		var uuid = c.Request.Header.Get("uuid")
		if len(uuid) == 0 {
			uuid = UUIDS.NewV4().String()
		}
		//请求前的记录
		library.Logger.Info("接收到请求：",
			zap.String("ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.RequestURI),
			zap.String("uuid", uuid),
		)
		//埋点
		c.Set(`uuid`, uuid)
		//往下走
		c.Next()
		//请求后的记录
		latency := time.Since(t)
		status := c.Writer.Status()
		library.Logger.Info("请求结束",
			zap.String("consume", fmt.Sprintf("%v ms", latency)),
			zap.Int("status", status),
			zap.String("uuid", uuid),
		)
	}
}
