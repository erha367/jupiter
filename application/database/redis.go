package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"jupiter/config"
	"log"
	"time"
)

var RedisCluster *redis.ClusterClient

func InitCluster() {
	//初始化集群
	RedisCluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        config.App.Redis.Cluster,
		Password:     config.App.Redis.Password,
		PoolSize:     config.App.Redis.PoolSize,
		MinIdleConns: config.App.Redis.MinIdle,
		IdleTimeout:  time.Second * 30,
	})
	//功能测试
	ctx := context.Background()
	val, err := RedisCluster.Do(ctx, `ping`).Result()
	if err != nil {
		log.Fatal(`redis conn err`, err)
	}
	log.Println(`redis 服务启动`, val)
}
