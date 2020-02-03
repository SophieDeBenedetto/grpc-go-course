package main

import (
	context "context"
	"fmt"
	"log"

	"github.com/SophieDeBenedetto/grpc-go-course/sum/sumpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting client...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client could not connect: %v", err)
	}

	defer cc.Close()
	c := sumpb.NewSumServiceClient(cc)
	doUnary(c)
}

func doUnary(c sumpb.SumServiceClient) {
	nums := &sumpb.Nums{
		Nums: []int64{1, 2, 3, 4},
	}
	request := &sumpb.SumRequest{Nums: nums}
	resp, err := c.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("Client received error from server: %v", err)
	}
	fmt.Println(resp.GetSum())
} // func (*server) Sum(ctx context.Conddtext, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
