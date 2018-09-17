package main

import (
	"context"
	"fmt"
	"github.com/matthewjamesboyle/grpc/calculator/calcproto"
	"google.golang.org/grpc"
	"io"
)

func main() {
	fmt.Println("Hello I'm a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed")
	}
	defer conn.Close()
	c := calcproto.NewSumServiceClient(conn)
	//res, err := c.Sum(context.Background(), &calcproto.SumRequest{
	//	Sum: &calcproto.Sum{
	//		FirstNum:  10,
	//		SecondNum: 3,
	//	},
	//})
	//if err != nil {
	//	fmt.Println("Greet client call failed")
	//}
	//
	//fmt.Println(res.Result)
	doServerStreaming(c)
}

func doServerStreaming(c calcproto.SumServiceClient) {
	fmt.Println("starting stream service")
	s, err := c.PrimeDecompStream(context.Background(), &calcproto.PrimeDecomposition{
		Num: 32000,
	})

	if err != nil {
		panic(err)
	}

	for {
		msg, err := s.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(msg.GetNum())
	}
}
