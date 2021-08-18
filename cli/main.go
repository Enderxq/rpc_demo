package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "rpc_demo/myProto"
)

func main() {
	//客户端连接服务端
	conn, err := grpc.Dial("127.0.0.1:10080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("network error", err)
	}

	//网络延迟关闭
	defer  conn.Close()
	//获得grpc句柄
	c := pb.NewHelloServerClient(conn)
	//通过句柄进行调用服务端函数SayHello
	re1, err := c.SayHello(context.Background(),&pb.HelloReq{Name:"Ender"})
	if err != nil {
		fmt.Println("calling SayHello() error", err)
	}
	fmt.Println(re1.Msg)

	//通过句柄进行调用服务端函数SayName
	re2, err := c.SayName(context.Background(),&pb.NameReq{Name:"Ender"})
	if err != nil {
		fmt.Println("calling SayName() error", err)
	}
	fmt.Println(re2.Msg)
}
