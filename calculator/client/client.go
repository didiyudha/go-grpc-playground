package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/didiyudha/go-grpc-playground/calculator/protogo/calculator"
	"google.golang.org/grpc"
)

func main() {
	clientConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil {
		log.Fatalln(err)
	}

	defer clientConn.Close()

	client := pb.NewCalculatorServiceClient(clientConn)
	// Add(client)
	PrimeNumberDecomposition(client)
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
