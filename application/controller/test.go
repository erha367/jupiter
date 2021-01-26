package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"jupiter/application/database"
	"jupiter/application/library"
	"time"
)

//接口demo
func Ping(c *gin.Context) {
	//redis-demo
	err := database.RedisCluster.Set(context.Background(), `name`, `pong`, time.Hour).Err()
	if err != nil {
		library.FailJson(c, err)
	}
	val := database.RedisCluster.Get(context.Background(), `name`).Val()
	//log.Println(val)
	library.OkJson(c, val)
}
