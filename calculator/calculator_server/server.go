package main

import (
	"context"
	"log"
	"net"

	"github.com/rahulrana95/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) GetNNumbersSum(ctx context.Context, req *calculatorpb.NNumbersSumRequest) (*calculatorpb.NNumbersSumResponse, error) {
	numArr := req.GetValues()

	var result int32 = 0

	for _, val := range numArr {
		result += val
	}
	res := &calculatorpb.NNumbersSumResponse{
		Result: result,
	}
	return res, nil
}

func main() {

	ss, err := net.Listen("tcp", "localhost:5000")

	if err != nil {
		log.Fatalf("Failed to start server.")
	}

	ser := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(ser, &server{})

	if err := ser.Serve(ss); err != nil {
		log.Fatalf("Failed to serve", err)
	}

}
