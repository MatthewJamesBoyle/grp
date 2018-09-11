package main

import (
	"context"
	"fmt"
	"github.com/matthewjamesboyle/grpc/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet was called")
	firstName := req.Greeting.FirstName
	lastName := req.Greeting.LastName

	result := "Hello" + firstName + lastName
	return &greetpb.GreetResponse{
		Result: result,
	}, nil
}

func main() {
	fmt.Print("hey")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("server failed")
	}
}
