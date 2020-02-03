package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/SophieDeBenedetto/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	fmt.Println("Starting server...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Calling `Greet` on the server...")
	name := req.GetGreeting().GetFirstName()
	return &greetpb.GreetResponse{Result: fmt.Sprintf("Hi, %v", name)}, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	name := req.GetGreeting().GetFirstName()
	for i := 0; i <= 10; i++ {
		result := fmt.Sprintf("Hi, %v, no. %d", name, i)
		res := &greetpb.GreetManyTimesResponse{Result: result}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}
