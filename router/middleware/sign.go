package middleware

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"jupiter/application/library"
	"net/url"
	"sort"
	"strings"
)

/**
 * @description：TODO
 * @author     ：yangsen
 * @date       ：2021/8/18 下午2:34
 * @company    ：eeo.cn
 */

//websocket连接池插件
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
	parms := []string{}
	for k, v := range req {
		if k == `sign` {
			continue
		}
		tmp := k + `=` + v[0]
		parms = append(parms, tmp)
	}
	sort.Strings(parms)
	val := strings.Join(parms, "&")
	xurl := val + `&key=classin`
	data := []byte(xurl)
	hash := md5.Sum(data)
	sign := fmt.Sprintf("%x", hash) //将[]byte转成16进制
	return sign
}
