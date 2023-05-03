package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hellogrpc "grpc_class/pb"
)

func main() {
	dial, e := grpc.Dial("localhost:8888", grpc.WithInsecure())
	fmt.Println(e)
	defer dial.Close()
	client := hellogrpc.NewHelloGRPCClient(dial)
	req, _ := client.SayHi(context.Background(), &hellogrpc.Req{Message: "我是客户端:hello"})
	fmt.Println(req.Message)
}
