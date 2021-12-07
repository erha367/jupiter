package library

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

/**
 * @description：计算签名
 * @author     ：yangsen
 * @date       ：2021/9/23 上午10:10
 * @company    ：eeo.cn
 */

//根据请求url获取sign
func GetSignByParms(req map[string][]string) (string, string, string) {
	ignores := []string{}
	if _, ok := req["ignoreSign"]; ok {
		ignores = strings.Split(req["ignoreSign"][0], `,`)
	}
	parms := []string{}
	for k, v := range req {
		if k == `sign` || k == `ignoreSign` || v[0] == `` {
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
	data := []byte(xurl)
	hash := md5.Sum(data)
	sign := fmt.Sprintf("%x", hash) //将[]byte转成16进制
	return sign, xurl, val
}

func InArray(val string, vals []string) bool {
	for _, v := range vals {
		if val == v {
			return true
		}
	}
	return false
}
