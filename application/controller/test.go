package controller

import (
	"github.com/gin-gonic/gin"
	"jupiter/application/validator"
	"jupiter/application/library"
	"log"
)

//接口demo
func Ping(c *gin.Context) {
	//查看cookie
	uid, err := c.Cookie(`_eeos_uid`)
	log.Println(uid, err)
	//设置cookie
	c.SetCookie(`cid`, `yangsen2021`, 3600, "/", ".eeo.im", false, true)
	name2, _ := c.Cookie(`cid`)
	log.Println(name2)
	library.OkJson(c, name2)
}

//登录成功返回json-web-token
func Login(c *gin.Context) {
	var uInfo validator.UserInfo
	if err := c.ShouldBind(&uInfo); err != nil {
		library.FailJson(c, err)
		return
	}
	//验证用户密码，略
	if token, err := library.GenerateToken(uInfo.UserName, uInfo.Password); err != nil {
		library.FailJson(c, err)
	} else {
		library.OkJson(c, token)
	}
	return
}

func Jwt(c *gin.Context) {
	library.OkJson(c, true)
}
