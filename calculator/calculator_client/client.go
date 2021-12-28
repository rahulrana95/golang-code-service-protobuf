package main

import (
	"context"
	"fmt"
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

	val := []int32{1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, -14}
	req := &calculatorpb.NNumbersSumRequest{
		Values: val,
	}

	res, err := c.GetNNumbersSum(context.Background(), req)

	if err != nil {
		fmt.Println("Greeting error", err)
	}

	fmt.Println("Greeting res", res.Result)

}
