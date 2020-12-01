package database

import (
	red "github.com/gomodule/redigo/redis"
	"jupiter/config"
	"log"
	"time"
)

type Redis struct {
	pool *red.Pool
}

var redis *Redis

func InitRedis() {
	redis = new(Redis)
	redis.pool = &red.Pool{
		MaxIdle:     50,
		MaxActive:   500,
		Wait:        true,
		IdleTimeout: time.Duration(120),
		Dial: func() (red.Conn, error) {
			return red.Dial(
				"tcp",
				config.App.Redis.Host,
				red.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				red.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				red.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				red.DialDatabase(0),
				red.DialPassword(config.App.Redis.Password),
			)
		},
	}
	log.Printf(`redis 连接成功`)
}

func RedExec(cmd string, key interface{}, args ...interface{}) (interface{}, error) {
	con := redis.pool.Get()
	if err := con.Err(); err != nil {
		return nil, err
	}
	defer con.Close()
	parmas := make([]interface{}, 0)
	parmas = append(parmas, key)
	if len(args) > 0 {
		for _, v := range args {
			parmas = append(parmas, v)
		}
	}
	return con.Do(cmd, parmas...)
}
