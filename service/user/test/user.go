package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	v1 "user/api/user/v1"
)

var userClient v1.UserClient
var conn *grpc.ClientConn

func main() {
	Init()

	TestCreateUser()

	conn.Close()
}

// Init 初始化 grpc 链接 注意这里链接的 端口
func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("grpc link err" + err.Error())
	}
	userClient = v1.NewUserClient(conn)
}

func TestCreateUser() {
	rsp, err := userClient.CreateUser(context.Background(), &v1.CreateUserInfo{
		Mobile:   fmt.Sprintf("1388888888%d", 1),
		Password: "admin123",
		NickName: fmt.Sprintf("YWWW%d", 1),
	})
	if err != nil {
		panic("grpc 创建用户失败" + err.Error())
	}
	fmt.Println(rsp.Id)
}
