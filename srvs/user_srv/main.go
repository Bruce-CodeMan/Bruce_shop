package main

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"Bruce_shop/srvs/user_srv/handler"
	"Bruce_shop/srvs/user_srv/proto"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip")
	PORT := flag.Int("port", 50051, "port")
	flag.Parse()
	fmt.Printf("ip: %s, port: %d\n", *IP, *PORT)
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *PORT))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	err = server.Serve(listener)
	if err != nil {
		panic("failed to start grpc: " + err.Error())
	}
}