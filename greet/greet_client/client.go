package main

import (
	"context"
	"fmt"
	"github.com/matthewjamesboyle/grpc/greet/greetpb"
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
	c := greetpb.NewGreetServiceClient(conn)
	//res, err := c.Greet(context.Background(), &greetpb.GreetRequest{
	//	Greeting: &greetpb.Greeting{
	//		FirstName: "Matt",
	//		LastName:  "Boyle",
	//	},
	//})
	//
	//if err != nil {
	//	fmt.Println("Greet client call failed")
	//}
	//
	//fmt.Println(res.Result)

	doServerStreaming(c)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting stream service")
	s, err := c.GreetManyTimes(context.Background(), &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Matt",
			LastName:  "Boyle",
		},
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
			fmt.Println("wtf")
		}

		fmt.Println(msg.GetResult())
	}

}
