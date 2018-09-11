package main

import (
	"context"
	"fmt"
	"github.com/matthewjamesboyle/grpc/calculator/calcproto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed")
	}
	defer conn.Close()
	c := calcproto.NewSumServiceClient(conn)
	res, err := c.Sum(context.Background(), &calcproto.SumRequest{
		Sum: &calcproto.Sum{
			FirstNum:  10,
			SecondNum: 3,
		},
	})
	if err != nil {
		fmt.Println("Greet client call failed")
	}

	fmt.Println(res.Result)
}
