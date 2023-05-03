package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hellogrpc "grpc_class/pb"
	"net"
)

type server struct {
	hellogrpc.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(ctx context.Context, req *hellogrpc.Req) (*hellogrpc.Res, error) {
	fmt.Println(req.GetMessage())
	return &hellogrpc.Res{Message: "我是服务端:hello"}, nil
}

func main() {
	l, e := net.Listen("tcp", ":8888")
	fmt.Println(e)
	s := grpc.NewServer()
	hellogrpc.RegisterHelloGRPCServer(s, &server{})
	s.Serve(l)
}
