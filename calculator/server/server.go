package main

import (
	"context"
	"fmt"
	"github.com/matthewjamesboyle/grpc/calculator/calcproto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Sum(ctx context.Context, r *calcproto.SumRequest) (*calcproto.SumResponse, error) {
	return &calcproto.SumResponse{Result: r.GetSum().GetFirstNum() + r.GetSum().GetSecondNum()}, nil
}

func main() {
	fmt.Print("Calculator service started")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcproto.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("server failed")
	}
}
