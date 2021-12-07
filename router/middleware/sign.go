package middleware

import (
	"github.com/gin-gonic/gin"
	"jupiter/library"
	"net/url"
	"strings"
)

/**
 * @description：sign中间件验证
 * @author     ：yangsen
 * @date       ：2021/8/18 下午2:34
 * @company    ：eeo.cn
 */

//sign中间件验证
func Sign() gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.Request.Method {
		case `POST`:
			err := c.Request.ParseForm()
			if err != nil {
				library.Fail(c, 102)
				c.Abort()
			}
			val := c.Request.PostForm
			sign, _, _ := library.GetSignByParms(val)
			sign2 := c.PostForm(`sign`)
			if sign != sign2 {
				library.Fail(c, 102)
				c.Abort()
			}
		case `GET`:
			urlx := c.Request.RequestURI
			path := strings.Split(urlx, `?`)
			v, err := url.ParseQuery(path[len(path)-1])
			if err != nil {
				library.Fail(c, 102)
				c.Abort()
			}
			sign, _, _ := library.GetSignByParms(v)
			sign2 := c.Query(`sign`)
			if sign != sign2 {
				library.Fail(c, 102)
				c.Abort()
			}
		default:
			//啥也不干
		}
	}
}
