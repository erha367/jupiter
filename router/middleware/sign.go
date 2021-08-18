package middleware

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"jupiter/application/library"
	"log"
	"net/url"
	"sort"
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
				library.FailJson(c, err)
				c.Abort()
			}
			val := c.Request.PostForm
			sign := GetSignByParms(val)
			sign2 := c.PostForm(`sign`)
			if sign != sign2 {
				library.FailJson(c, errors.New(`sign value err`))
				c.Abort()
			}
		case `GET`:
			urlx := c.Request.RequestURI
			path := strings.Split(urlx, `?`)
			v, err := url.ParseQuery(path[len(path)-1])
			if err != nil {
				library.FailJson(c, err)
				c.Abort()
			}
			sign := GetSignByParms(v)
			//log.Println(sign)
			sign2 := c.Query(`sign`)
			if sign != sign2 {
				library.FailJson(c, errors.New(`sign value err`))
				c.Abort()
			}
		default:

		}
	}
}

//根据请求url获取sign
func GetSignByParms(req map[string][]string) string {
	ignores := []string{}
	if _, ok := req["ignoreSign"]; ok {
		ignores = strings.Split(req["ignoreSign"][0], `,`)
	}
	parms := []string{}
	for k, v := range req {
		if k == `sign` || k == `ignoreSign` {
			continue
		}
		if len(ignores) > 0 {
			if InArray(k, ignores) {
				continue
			}
		}
		tmp := k + `=` + v[0]
		parms = append(parms, tmp)
	}
	sort.Strings(parms)
	val := strings.Join(parms, "&")
	xurl := val + `&key=classin`
	log.Println(xurl)
	data := []byte(xurl)
	hash := md5.Sum(data)
	sign := fmt.Sprintf("%x", hash) //将[]byte转成16进制
	return sign
}

func InArray(val string, vals []string) bool {
	for _, v := range vals {
		if val == v {
			return true
		}
	}
	return false
}
