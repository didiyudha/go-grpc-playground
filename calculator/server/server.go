package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"

	pb "github.com/didiyudha/go-grpc-playground/calculator/protogo/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CalculatroServer - calculator server type
type CalculatroServer struct{}

// Add operation implementation
func (*CalculatroServer) Add(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	time.Sleep(time.Duration(2) * time.Second)

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("client cancel the the request")
		return nil, status.Errorf(codes.DeadlineExceeded, "Client cancelled, abandoning.")
	}
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

// FindMax - find maximum bidirectional streaming
func (*CalculatroServer) FindMax(stream pb.CalculatorService_FindMaxServer) error {
	var max int64 = -10000000
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		n := in.GetN()
		if n > max {
			max = n
		}
		res := &pb.FindMaxResponse{
			Max: max,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}

func (*CalculatroServer) SquareRoot(ctx context.Context, req *pb.SquareRootRequest) (*pb.SquareRootResponse, error) {
	fmt.Println("Received SquareRoot RPC")
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}
	return &pb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
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
