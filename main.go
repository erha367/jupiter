package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
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
	//系统初始化
	defer database.CloseDatabases()
	application.Bootstrap()
	//初始化db
	database.InitDatabases()
	//初始化redis
	database.InitCluster()
	//启动模式
	gin.SetMode(config.Mode())
	//路由
	apiRouter := router.ApiRouter()
	//pprof
	pprof.Register(apiRouter)
	//启动
	errChan := make(chan error)
	go func() {
		err := apiRouter.Run(`:` + strconv.Itoa(config.App.HttpPort))
		if err != nil {
			errChan <- err
		}
	}()
	go func() {
		sigc := make(chan os.Signal)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sigc)
	}()
	e := <-errChan
	log.Println(`sys interrupt :`, e)
}
