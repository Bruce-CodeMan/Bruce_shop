/**
 * @Author: Bruce
 * @Description: test the serverAPI
 * @Date: 2023/2/15 9:51 AM
 */

package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"Bruce_shop/srvs/user_srv/proto"
)

var (
	userClient proto.UserClient
	conn       *grpc.ClientConn
)

func Init() {
	var err error
	if conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		panic("连接出现问题")
	}
	userClient = proto.NewUserClient(conn)
}

func TestGetUserList() {
	resp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 5,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range resp.Data {
		fmt.Printf("Id: %d, Nickname: %s, Mobile: %s\n", user.Id, user.NickName, user.Mobile)
		res, err := userClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
			Password:         "admin123",
			EncryptoPassword: user.Password,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("核验密码:", res.Success)
	}
}

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		resp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			Mobile:   fmt.Sprintf("1883092027%d", i),
			NickName: fmt.Sprintf("Bruce--%d", i),
			Password: fmt.Sprintf("Bruce"),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("服务端返回的Id:", resp.Id)
	}
}

func TestGetUserByMobile() {
	for i := 0; i < 10; i++ {
		resp, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
			Mobile: fmt.Sprintf("1883092027%d", i),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("服务端返回的Id:", resp.Id)
	}
}

func TestGetUserById() {
	for i := 5; i < 15; i++ {
		resp, err := userClient.GetUserById(context.Background(), &proto.IdRequest{
			Id: int32(i),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("服务端返回的Mobile:", resp.Mobile)
	}
}

func TestUpdateUser() {
	for i := 1; i < 10; i++ {
		_, err := userClient.UpdateUser(context.Background(), &proto.UpdateUserInfo{
			Id:     int32(i),
			Gender: "female",
		})
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	Init()
	TestGetUserList()
	//TestGetUserByMobile()
	//TestGetUserById()
	//TestUpdateUser()
	defer conn.Close()
}
