package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/didiyudha/go-grpc-playground/calculator/protogo/calculator"
	"google.golang.org/grpc"
)

// CalculatroServer - calculator server type
type CalculatroServer struct{}

// Add operation implementation
func (*CalculatroServer) Add(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()
	res := &pb.CalculatorResponse{
		Total: (firstNumber + secondNumber),
	}
	return res, nil
}

func main() {
	var calculatorServer CalculatroServer
	conn, err := net.Listen("tcp", "localhost:50052")

	if err != nil {
		log.Fatalln(err)
	}

	gRPCServer := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(gRPCServer, &calculatorServer)

	fmt.Println("Calculator gRPC server")

	if err := gRPCServer.Serve(conn); err != nil {
		log.Fatalln(err)
	}
}
