package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "rpc_demo/myProto"
)

type server struct {}

func (this *server) SayHello(ctx context.Context, in *pb.HelloReq) (out *pb.HelloResponse,err error) {
	return &pb.HelloResponse{Msg:"hello"}, nil
}

func (this *server) SayName(ctx context.Context, in *pb.NameReq) (out *pb.NameResponse,err error){
	//return &pd.NameResponse{Msg:in.Name + "it is name"}, nil
	return &pb.NameResponse{Msg:in.Name + " is my name"}, nil
}

func main()  {
	ln, err := net.Listen("tcp", ":10080")
	if err != nil {
		fmt.Println("network error", err)
	}

	//创建grpc服务
	srv := grpc.NewServer()
	//注册服务
	pb.RegisterHelloServerServer(srv, &server{})
	err = srv.Serve(ln)
	if err != nil {
		fmt.Println("Serve error", err)
	}
}
