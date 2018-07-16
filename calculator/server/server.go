package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	pb "github.com/didiyudha/go-grpc-playground/calculator/protogo/calculator"
	"google.golang.org/grpc"
)

// CalculatroServer - calculator server type
type CalculatroServer struct{}

// Add operation implementation
func (*CalculatroServer) Add(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	fmt.Printf("Server receive request: %+v\n", req)
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()
	res := &pb.CalculatorResponse{
		Total: (firstNumber + secondNumber),
	}
	fmt.Printf("Server send response: %+v\n", res)
	return res, nil
}

// PrimeNumberDecomposition - server streaming
func (*CalculatroServer) PrimeNumberDecomposition(req *pb.PrimeNumberDecompositionRequest, stream pb.CalculatorService_PrimeNumberDecompositionServer) error {
	n := req.GetN()
	var k uint64 = 2
	for n > 1 {
		if n%k == 0 {
			res := &pb.PrimeNumberDecompositionResponse{
				Result: k,
			}
			if err := stream.Send(res); err != nil {
				return err
			}
			n = n / k
		} else {
			k++
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

// Average of total amount that client send to the server
func (*CalculatroServer) Average(stream pb.CalculatorService_AverageServer) error {
	var n int
	var total uint64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			avg := float64(total) / float64(n)
			res := &pb.AverageResponse{
				Result: avg,
			}
			return stream.SendAndClose(res)
		}
		if err != nil {
			return err
		}
		temp := req.GetNumber()
		total += temp
		n++
	}
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
