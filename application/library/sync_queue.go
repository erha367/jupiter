package library

import (
	"context"
	ytask "github.com/gojuukaze/YTask/v2"
	"github.com/gojuukaze/YTask/v2/server"
	"jupyter/config"
	"strings"
)

var YClient server.Client

func Queue(ctx context.Context) {
	val := strings.Split(config.App.Redis.Host, `:`)
	broker := ytask.Broker.NewRedisBroker(val[0], val[1], "", 0, 0)
	backend := ytask.Backend.NewRedisBackend(val[0], val[1], "", 0, 0)
	ser := ytask.Server.NewServer(
		ytask.Config.Broker(&broker),
		ytask.Config.Backend(&backend),
		ytask.Config.Debug(true),
		ytask.Config.StatusExpires(60*5),
		ytask.Config.ResultExpires(60*5),
	)
	YClient = ser.GetClient()
	//函数注册，类似laravel的event-listener
	ser.Add("group1", "addRec", Ad)
	ser.Run("group1", 1)

	//退出
	for {
		select {
		case <-ctx.Done():
			ser.Shutdown(context.Background())
			return
		}
	}
}

func Add() {

}
