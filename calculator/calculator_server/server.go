package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/rahulrana95/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) GetNNumbersSum(ctx context.Context, req *calculatorpb.NNumbersSumRequest) (*calculatorpb.NNumbersSumResponse, error) {
	fmt.Println("GetNNumbersSum running")
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

func (*server) GetPrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_GetPrimeNumberDecompositionServer) error {
	log.Println("GetPrimeNumberDecomposition called, streaming from server")
	numFromReq := req.GetNum()

	var k int32 = 2
	var N int32 = numFromReq
	for N > 1 {
		if N%k == 0 { // if k evenly divides into N
			log.Println(k)
			resObj := &calculatorpb.PrimeNumberDecompositionResponse{
				Num: k,
			}
			stream.Send(resObj)
			N = N / k
		} else {
			k = k + 1
		}

		time.Sleep(1000 * time.Millisecond)
	}

	return nil

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
