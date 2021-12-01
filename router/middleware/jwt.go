package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"jupiter/library"
	"time"
)

/**
 * @description：TODO
 * @author     ：yangsen
 * @date       ：2021/8/19 上午11:41
 * @company    ：eeo.cn
 */

func Jwtm() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(`token`)
		claim, err := library.ParseToken(token)
		if err != nil {
			library.FailJson(c, err)
			c.Abort()
		}
		left := claim.ExpiresAt - time.Now().Unix()
		if left < 0 {
			library.FailJson(c,errors.New(`token is out time`))
			c.Abort()
		}
	}
}
