package main

import (
	"context"
	"fmt"
	"growing-into-an-excellent-golang-developer/project-layout/services/hello/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedHelloServer
}

func (s *server) SayHello(ctx context.Context, req *proto.SayRequest) (*proto.SayResponse, error) {
	fmt.Printf("request:%v", req.Name)
	return &proto.SayResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	proto.RegisterHelloServer(s, &server{})
	reflection.Register(s)

	defer func() {
		s.Stop()
		listen.Close()
	}()

	fmt.Println("Serving 50051...")
	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}

}
