package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/didiyudha/go-grpc-playground/unary/greeting/greet"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to dial %v\n", err)
	}
	defer conn.Close()

	client := greet.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &greet.GreeetRequest{
		Name: "Didi",
	}
	res, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("error when call SayHello %v\n", err)
	}
	msg := res.GetMsg()
	fmt.Println("Response from server: ", msg)
}
