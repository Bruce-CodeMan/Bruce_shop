package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"Bruce_shop/srvs/user_srv/global"
	"Bruce_shop/srvs/user_srv/handler"
	"Bruce_shop/srvs/user_srv/initialize"
	"Bruce_shop/srvs/user_srv/proto"
	"Bruce_shop/srvs/user_srv/utils"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip")
	PORT := flag.Int("port", 0, "port")

	// initialize
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	flag.Parse()
	if *PORT == 0 {
		*PORT, _ = utils.GetFreePort()
	}
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
		GRPC:                           fmt.Sprintf("127.0.0.1:%d", *PORT),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}
	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	serviceID := fmt.Sprintf("%s", uuid.NewV4())
	registration.ID = serviceID
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
