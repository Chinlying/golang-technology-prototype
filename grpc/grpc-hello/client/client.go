package main

import (
	"bufio"
	"context"
	"fmt"
	proto "golang-technology-prototype/grpc/grpc-hello/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func main() {
	var serviceHost = "127.0.0.1:8001"

	conn, err := grpc.Dial(serviceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := proto.NewHelloClient(conn)
	rsp, err := client.SayHello(context.TODO(), &proto.SayRequest{
		Name: "Chinlying",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp)

	fmt.Println("Press Enter to Abort...")
	in := bufio.NewReader(os.Stdin)
	_, _, _ = in.ReadLine()
}
