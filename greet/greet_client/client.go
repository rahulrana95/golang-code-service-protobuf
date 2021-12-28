package main

import (
	"context"
	"fmt"
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

}
