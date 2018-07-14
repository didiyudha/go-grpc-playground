package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/didiyudha/go-grpc-playground/greet/protogo/greet"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer cc.Close()
	client := pb.NewGreetServiceClient(cc)
	// unaryCall(client)
	// serverStreamingCall(client)
	clientStreamingCall(client)
}

func unaryCall(client pb.GreetServiceClient) {
	dataset := []struct {
		FirstName string
		LastName  string
	}{
		{
			FirstName: "Didi",
			LastName:  "Yudha",
		},
		{
			FirstName: "Eko",
			LastName:  "Budi",
		},
		{
			FirstName: "Yance",
			LastName:  "Parisman",
		},
		{
			FirstName: "Dwi",
			LastName:  "Susanti",
		},
		{
			FirstName: "Soegiarto",
			LastName:  "Tumiyam",
		},
	}
	for _, data := range dataset {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		req := &pb.GreetRequest{
			Greeting: &pb.Greeting{
				FirstName: data.FirstName,
				LastName:  data.LastName,
			},
		}
		fmt.Printf("Client send request FirstName: %s, Lastname: %s\n",
			req.GetGreeting().GetFirstName(), req.GetGreeting().GetLastName())
		res, err := client.Greet(ctx, req)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Client get response: ", res.GetResult())
		time.Sleep(2 * time.Second)
	}
}

func serverStreamingCall(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	req := pb.GreetManyTimesRequest{
		Greeting: &pb.Greeting{
			FirstName: "Didi",
			LastName:  "Yudha",
		},
	}
	stream, err := client.GreetManyTimes(ctx, &req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(res.GetResult())
	}
}

func clientStreamingCall(client pb.GreetServiceClient) {
	names := []struct {
		FirstName string
		LastName  string
	}{
		{
			FirstName: "Didi",
			LastName:  "Yudha",
		},
		{
			FirstName: "Eko",
			LastName:  "Budi",
		},
		{
			FirstName: "Yance",
			LastName:  "Parisman",
		},
		{
			FirstName: "Dwi",
			LastName:  "Susanti",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	stream, err := client.LongGreet(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range names {
		req := &pb.LongGreetRequest{
			Greeting: &pb.Greeting{
				FirstName: name.FirstName,
				LastName:  name.LastName,
			},
		}
		fmt.Printf("Client sends %v\n", req.GetGreeting())
		if err := stream.Send(req); err != nil {
			log.Fatalf("%v.LongGreet", stream)
		}
		time.Sleep(2 * time.Second)
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Response server: ", reply.GetResult())
}
