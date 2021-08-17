package test

import (
	"context"
	consulApi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"jupiter/application/proto"
	"net"
	"strconv"
	"testing"
	"time"
)

func TestGRPCClient(t *testing.T) {
	/*- 1.创建client -*/
	config := consulApi.DefaultConfig()
	config.Address = `consul.sit13.dom:61170`
	config.Token = `p2BE1AtpwPbrxZdC6k+eXA==`
	client, err := consulApi.NewClient(config)
	if err != nil {
		t.Log(err)
		return
	}
	/*- 2.获取健康地址 -*/
	services, _, err := client.Health().Service(`jupiter`, `jup`, true, nil)
	if err != nil {
		t.Log(err)
		return
	}
	var addrs []string
	for _, service := range services {
		t.Log(service.Service)
		addrs = append(addrs, net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port)))
	}
	/*- 3.获取了好多地址，随机、轮询、加权等自己实现 -*/
	ctx, out := context.WithTimeout(context.Background(), time.Second)
	defer out()
	conn, err := grpc.DialContext(ctx, addrs[0], grpc.WithInsecure())
	if err != nil {
		t.Log(err)
	}
	defer conn.Close()
	rpcCli := proto.NewHelloServiceClient(conn)
	res, err := rpcCli.SayHello(ctx, &proto.HelloRequest{
		Name: "welcome to the USA,have a good day!",
	})
	t.Log(res, err)
}
