package controller

import (
	"github.com/gin-gonic/gin"
	"jupiter/domain/repository"
	"jupiter/library"
)

//接口demo
func Ping(c *gin.Context) {
	library.Success(c, `pong`)
}

func GetFolder(c *gin.Context) {
	folder := repository.EeoCourseFoldersMgr(library.GetMasterDB(`eo_osfile`))
	library.Success(c, folder.GetTableName())
}
