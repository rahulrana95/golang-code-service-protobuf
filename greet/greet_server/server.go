package main

import (
	"fmt"
	"log"
	"net"

	"github.com/rahulrana95/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
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
