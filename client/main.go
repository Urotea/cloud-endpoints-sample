package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Urotea/cloud-endpoints-sample/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	context := context.Background()
	conn, err := grpc.Dial(
		"SERVER_HOST_NAME:443",
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
		grpc.WithTimeout(time.Duration(10*time.Second)))
	if err != nil {
		log.Fatalf("client connection error: %#v \n", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	message := &pb.HelloRequest{Name: "World"}
	res, err := client.SayHello(context, message)
	if err != nil {
		log.Fatalf("sayHelloの送信に失敗しました。 err = %#v \n", err)
	}
	fmt.Printf("result:%#v \n", res)
}
