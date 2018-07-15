package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "github.com/didiyudha/go-grpc-playground/greet/protogo/greet"
)

type GreetServer struct{}

// Greet unary call
func (gs *GreetServer) Greet(ctx context.Context, greetRequest *pb.GreetRequest) (*pb.GreetResponse, error) {
	fname := greetRequest.GetGreeting().GetFirstName()
	lnmae := greetRequest.GetGreeting().GetLastName()
	fmt.Printf("Server get request First Name: %s, Last Name: %s\n", fname, lnmae)
	msg := fmt.Sprintf("Hi, %s %s, Nice to meet you", fname, lnmae)
	return &pb.GreetResponse{
		Result: msg,
	}, nil
}

// GreetManyTimes - server streaming call
func (gs *GreetServer) GreetManyTimes(req *pb.GreetManyTimesRequest, stream pb.GreetService_GreetManyTimesServer) error {
	fname := req.GetGreeting().GetFirstName()
	lname := req.GetGreeting().GetLastName()
	msg := fmt.Sprintf("Hi, %s %s, Nice to meet you", fname, lname)

	for i := 1; i <= 10; i++ {
		resp := &pb.GreetManyTimesResponse{
			Result: msg + fmt.Sprintf(" %d ", i),
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}

// LongGreet - client streaming call
func (gs *GreetServer) LongGreet(stream pb.GreetService_LongGreetServer) error {
	msg := "Hi,"
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			msg += " nice to meet you"
			res := &pb.LongGreetResponse{
				Result: msg,
			}
			return stream.SendAndClose(res)
		}
		if err != nil {
			return err
		}
		fmt.Printf("Get request from client: %+v\n", req.GetGreeting())
		msg += fmt.Sprintf(" %s %s, ", req.GetGreeting().GetFirstName(), req.GetGreeting().GetLastName())
	}
}

// GreetEveryOne - Bidirectional call streaming
func (gs *GreetServer) GreetEveryOne(streamin pb.GreetService_GreetEveryOneServer) error {
	for {
		in, err := streamin.Recv()

		// Client finish send message to the server
		if err == io.EOF {
			return nil
		}

		// Error occured
		if err != nil {
			return err
		}

		// Extract request message
		fmt.Printf("Client send request %+v\n", in.GetGreeting())
		fname := in.GetGreeting().GetFirstName()
		lname := in.GetGreeting().GetLastName()

		// Construct the message response and set to the protobuf response type
		msg := fmt.Sprintf("Hi %s %s nice to meet you", fname, lname)
		res := &pb.GreetEveryOneResponse{
			Result: msg,
		}

		// Send the response to the client
		if err := streamin.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
}

func main() {
	var greetServer GreetServer
	conn, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &greetServer)

	fmt.Println("Greet server gRPC")

	if err := grpcServer.Serve(conn); err != nil {
		log.Fatalln(err)
	}

}
