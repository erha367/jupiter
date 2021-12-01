package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jupiter/application"
	"jupiter/application/database"
	"jupiter/config"
	"jupiter/router"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	//初始化日志
	application.Bootstrap()
	//初始化db
	database.InitDatabases()
	//初始化redis
	//database.InitCluster()
	//启动模式
	gin.SetMode(config.Mode())
	//系统初始化
	defer database.CloseDatabases()
	//路由
	apiRouter := router.ApiRouter()
	//pprof
	//pprof.Register(apiRouter)
	//启动
	errChan := make(chan error)
	go func() {
		err := apiRouter.Run(`:` + strconv.Itoa(config.App.HttpPort))
		if err != nil {
			errChan <- err
		}
	}()
	//信号接收类似：ctrl+c
	go func() {
		sigc := make(chan os.Signal)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sigc)
	}()
	e := <-errChan
	log.Println(`sys interrupt :`, e)
}
