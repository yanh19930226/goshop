package main

import (
	"context"
	"fmt"
	"golearn/grpctest/proto"

	"google.golang.org/grpc"
)

func main() {
	//stream
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	fmt.Println("启动成功！！")

	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "bobby"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
