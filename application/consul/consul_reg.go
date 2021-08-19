package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"jupiter/application/utils"
	"jupiter/config"
	"log"
	"time"
)

var ConsulClient *api.Client
var err error

func RegitserService() {
	//consul client
	consulConfig := api.DefaultConfig()
	consulConfig.Address = config.App.Consul.Host
	consulConfig.Token = config.App.Consul.Token
	ConsulClient, err = api.NewClient(consulConfig)
	if err != nil {
		fmt.Printf("NewClient error\n%v", err)
		return
	}
	agent := ConsulClient.Agent()
	interval := time.Duration(10) * time.Second
	deregister := time.Duration(1) * time.Minute
	ip, _ := utils.GetLocalIp()
	//register consul
	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", config.App.Name, ip, config.App.RpcPort), // 服务节点的名称
		Name:    config.App.Name,                                                  // 服务名称
		Tags:    []string{config.App.Consul.Tag},                                  // tag，可以为空
		Port:    config.App.RpcPort,                                               // 服务端口
		Address: config.App.IP,                                                    // 服务 IP
		Check: &api.AgentServiceCheck{ // 健康检查
			Interval:                       interval.String(),                                                // 健康检查间隔
			GRPC:                           fmt.Sprintf("%v:%v/%v", ip, config.App.RpcPort, config.App.Name), // grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
			DeregisterCriticalServiceAfter: deregister.String(),                                              // 注销时间，相当于过期时间
		},
	}
	fmt.Printf("registing to %v\n", config.App.Consul.Host)
	if err := agent.ServiceRegister(reg); err != nil {
		fmt.Printf("Service Register error\n%v", err)
		return
	}
}

//反注册
func UnRegService() {
	log.Println(`consul bye`)
	err := ConsulClient.Agent().ServiceDeregister(fmt.Sprintf("%v-%v-%v", config.App.Name, config.App.IP, config.App.RpcPort))
	if err != nil {
		log.Println(`consul unreg err`, err)
	}
}
