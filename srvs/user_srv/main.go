package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"Bruce_shop/srvs/user_srv/global"
	"Bruce_shop/srvs/user_srv/handler"
	"Bruce_shop/srvs/user_srv/initialize"
	"Bruce_shop/srvs/user_srv/proto"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip")
	PORT := flag.Int("port", 50051, "port")

	// initialize
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	zap.S().Info(global.ServerConfig)

	flag.Parse()
	zap.S().Infof("ip: %s, port: %d\n", *IP, *PORT)
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *PORT))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	// 注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 服务注册
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	check := &api.AgentServiceCheck{
		GRPC:                           "127.0.0.1:50051",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	registration.ID = global.ServerConfig.Name
	registration.Port = *PORT
	registration.Tags = []string{"Bruce", "Hsu"}
	registration.Address = "127.0.0.1"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	err = server.Serve(listener)
	if err != nil {
		panic("failed to start grpc: " + err.Error())
	}
}
