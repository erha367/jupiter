package main

import (
	"context"
	"github.com/fvbock/endless"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"jupiter/application"
	"jupiter/application/database"
	"jupiter/application/library"
	"jupiter/config"
	"jupiter/router"
	"log"
	"os"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//系统初始化
	defer database.CloseDatabases()
	application.Bootstrap()
	database.InitDatabases()
	database.InitRedis()
	gin.SetMode(config.Mode())
	//异步队列
	go library.Queue(ctx)
	//路由
	apiRouter := router.ApiRouter()
	//等待组-热重启相关
	pprof.Register(apiRouter)
	w := sync.WaitGroup{}
	w.Add(2)
	//启动80端口
	go func() {
		err := endless.ListenAndServe(`:`+strconv.Itoa(config.App.HttpPort), apiRouter)
		if err != nil {
			log.Println(err)
		}
		log.Println(`stopped`, config.App.HttpPort)
		w.Done()
	}()
	//启动443端口
	go func() {
		err := endless.ListenAndServeTLS(`:`+strconv.Itoa(config.App.HttpsPort), config.App.CertFile, config.App.KeyFile, apiRouter)
		if err != nil {
			log.Println(err)
		}
		log.Println(`stopped`, config.App.HttpsPort)
		w.Done()
	}()
	log.Printf("Actual pid is %d", syscall.Getpid())
	w.Wait()
	log.Println(`All servers stopped. Exiting.`)
	os.Exit(0)
}
