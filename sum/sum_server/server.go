package main

import (
	context "context"
	"fmt"
	"log"
	"net"

	"github.com/SophieDeBenedetto/grpc-go-course/sum/sumpb"
	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	fmt.Println("Starting Sum server...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	nums := req.GetNums().Nums
	fmt.Println(nums)
	var sum int64

	for i := 0; i <= len(nums)-1; i++ {
		sum += nums[i]
	}
	return &sumpb.SumResponse{Sum: sum}, nil
}
