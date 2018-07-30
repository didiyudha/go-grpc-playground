package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	pb "github.com/didiyudha/go-grpc-playground/calculator/protogo/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	clientConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil {
		log.Fatalln(err)
	}

	defer clientConn.Close()

	client := pb.NewCalculatorServiceClient(clientConn)
	// Add(client)
	// PrimeNumberDecomposition(client)
	// Average(client)
	// FindMax(client)
	// doErrorUnary(client)
	doErrorTimeoutCall(client)
}

// Add - Unary call
func Add(client pb.CalculatorServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	req := &pb.CalculatorRequest{
		FirstNumber:  10.0,
		SecondNumber: 25.0,
	}
	fmt.Printf("Client send request: %+v\n", req)
	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Client receive response: %+v\n", res)
	fmt.Println(res.GetTotal())
}

// PrimeNumberDecomposition - server streaming
func PrimeNumberDecomposition(client pb.CalculatorServiceClient) {
	req := &pb.PrimeNumberDecompositionRequest{
		N: 120,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	streamin, err := client.PrimeNumberDecomposition(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := streamin.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(res.GetResult())
	}
}

// Average - client streaming calls
func Average(client pb.CalculatorServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	stream, err := client.Average(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 1; i <= 10; i++ {
		req := &pb.AverageRuequest{
			Number: uint64(i),
		}
		if err := stream.Send(req); err != nil {
			log.Fatalln(err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Avg: ", strconv.FormatFloat(reply.GetResult(), 'f', 2, 64))
}

func FindMax(client pb.CalculatorServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	stream, err := client.FindMax(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("Maximum: ", res.GetMax())
		}
	}()
	numbers := []int64{
		10,
		5,
		100,
		7,
		6,
		101,
		1000,
	}
	for _, n := range numbers {
		r := &pb.FindMaxRequest{
			N: n,
		}
		log.Println("Client send: ", n)
		if err := stream.Send(r); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
}

func doErrorUnary(c pb.CalculatorServiceClient) {
	fmt.Println("Starting to do a SquareRoot Unary RPC...")

	// correct call
	doErrorCall(c, 10)

	// error call
	doErrorCall(c, -2)
}

func doErrorCall(c pb.CalculatorServiceClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &pb.SquareRootRequest{Number: n})

	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			fmt.Printf("Error message from server: %v\n", respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("Big Error calling SquareRoot: %v", err)
			return
		}
	}
	fmt.Printf("Result of square root of %v: %v\n", n, res.GetNumberRoot())
}

func doErrorTimeoutCall(client pb.CalculatorServiceClient) {
	req := pb.CalculatorRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Add(ctx, &req)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			fmt.Printf("Error message from server: %v\n", respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("Big Error calling SquareRoot: %v", err)
			return
		}
	}
	fmt.Println(res.GetTotal())
}
