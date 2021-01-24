package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/Urotea/cloud-endpoints-sample/proto"
	"google.golang.org/grpc"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(context context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("called sayHello request = %#v \n", req)
	return &pb.HelloReply{Message: "reply"}, nil
}

var port = flag.String("port", "8080", "set port")

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", *port))
	if err != nil {
		log.Fatalf("portの指定に失敗しました。 err = %#v \n", err)
	}

	s := grpc.NewServer()
	greetrServer := &greeterServer{}
	pb.RegisterGreeterServer(s, greetrServer)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("サーバの起動に失敗しました。 err = %#v \n", err)
	}
}
