package library

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
 * @description：业务成功/失败返回
 * @author     ：yangsen
 * @date       ：2021/12/1 下午1:38
 * @company    ：eeo.cn
 */

type Response struct {
	ErrorInfo ErrorInfo   `json:"error_info"`
	Data      interface{} `json:"data"`
}

type ErrorInfo struct {
	Errno int    `json:"errno"`
	Error string `json:"error"`
}

const SucessInfo = "程序正常执行"
const FailInfo = "系统异常"

//成功返回
func Success(c *gin.Context, data interface{}) {
	var res Response
	res.ErrorInfo.Errno = OK
	res.ErrorInfo.Error = SucessInfo
	res.Data = data
	c.JSON(http.StatusOK, res)
}

//异常返回
func Fail(c *gin.Context, code int) {
	var res Response
	codeStr := strconv.Itoa(code)
	res.ErrorInfo.Errno = code
	msg, err := GetErrorMsg(codeStr)
	if err != nil {
		res.ErrorInfo.Error = FailInfo
	} else {
		res.ErrorInfo.Error = msg
	}
	c.JSON(http.StatusOK, res)
}
