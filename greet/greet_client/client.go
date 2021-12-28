package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/rahulrana95/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("I am a clinet")
	cc, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Faield strating error")
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Rahul",
			LastName:  "Rana",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		fmt.Println("Greeting error", err)
	}

	fmt.Println("Greeting res", res)

	doServerStreaming(c)

}

func doServerStreaming(c greetpb.GreetServiceClient) {

	fmt.Println("Starting a streaming rpc cleint")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Rahul",
			LastName:  "Rana",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatal("stream error", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("stream error 2", err)
		}
		result := msg.GetResult()

		log.Println("result from stream is: ", result)
	}

}
