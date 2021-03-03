package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"jupiter/application/proto"
	"jupiter/config"
	"log"
	"net"
	"strconv"
)

/**
 * @description：TODO
 * @author     ：yangsen
 * @date       ：2020/12/3 下午8:46
 * @company    ：eeo.cn
 */

type GrpcServer struct{}

func (s *GrpcServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	log.Println(`--- req ---`)
	//time.Sleep(time.Millisecond*600)
	return &proto.HelloResponse{
		Result:               "hello xxx " + req.Name,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}, nil
}

func GrpcInit() {
	port := `:` + strconv.Itoa(config.App.RpcPort)
	lis, err := net.Listen(`tcp`, port)
	if err != nil {
		log.Println(`net listen err`)
		return
	}
	s := grpc.NewServer()
	proto.RegisterHelloServiceServer(s, &GrpcServer{})
	grpc_health_v1.RegisterHealthServer(s, &HealthImpl{})
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		log.Println("failed to serve: %v", err)
		return
	}
}

//health
type HealthImpl struct{}

// Check 实现健康检查接口，这里直接返回健康状态，这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	fmt.Print("health checking\n")
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *HealthImpl) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return nil
}
