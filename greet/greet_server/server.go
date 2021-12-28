package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/rahulrana95/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()

	result := "Hello! How are you " + firstName

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil

}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstName := req.GetGreeting().GetFirstName()

	result := "Hello! How are you " + firstName

	for i := 0; i < 10; i++ {
		res := &greetpb.GreetManyTimesResponse{
			Result: result + strconv.Itoa(i),
		}

		stream.Send(res)

		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:5000")
	log.Println("Listening server at", 5000)
	if err != nil {
		fmt.Println("Not able to start server")
		log.Fatalf("FAiled to listen", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve", err)
	}

}
