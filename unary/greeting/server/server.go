package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"

	"github.com/didiyudha/go-grpc-playground/unary/greeting/greet"
)

const (
	address     = ":50051"
	defaultName = "Didi"
)

// GreeterServer - greeter server gRPC
type GreeterServer struct{}

// SayHello implementation by GreeterServer
func (gs *GreeterServer) SayHello(ctx context.Context, req *greet.GreeetRequest) (*greet.GreeetResponse, error) {
	name := req.GetName()
	if len(strings.TrimSpace(name)) == 0 {
		name = defaultName
	}
	msg := fmt.Sprintf("Hi, %s nice to meet you!", name)
	resp := &greet.GreeetResponse{
		Msg: msg,
	}
	return resp, nil
}

func main() {
	var greeterServer GreeterServer
	conn, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("can not listen to %s\n", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	greet.RegisterGreeterServer(grpcServer, &greeterServer)

	fmt.Println("GreeterServer is up")

	if err := grpcServer.Serve(conn); err != nil {
		log.Fatalf("unable to serve %v\n", err)
	}

}
