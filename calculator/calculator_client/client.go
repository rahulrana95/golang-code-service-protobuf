package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/rahulrana95/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("I am a clinet")
	cc, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Faield strating error")
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	req := &calculatorpb.NNumbersSumRequest{
		Values: []int32{1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, -14},
	}

	res, err := c.GetNNumbersSum(context.Background(), req)

	if err != nil {
		fmt.Println("Greeting error ->", err)
	}

	fmt.Println("Greeting res", res)

	doServerStreamingPrimeDecompostion(c)

}

func doServerStreamingPrimeDecompostion(c calculatorpb.CalculatorServiceClient) {

	fmt.Println("Starting a streaming rpc cleint")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Num: 1000,
	}

	resStream, err := c.GetPrimeNumberDecomposition(context.Background(), req)

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
		result := msg.GetNum()

		log.Println("result from stream is: ", result)
	}

}
