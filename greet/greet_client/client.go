package main

import (
	"fmt"
	"log"
	context "context"

	"github.com/SophieDeBenedetto/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting client...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
		log.Fatalf("Client could not connect: %v", err)
	}

	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	greeting := &greetpb.Greeting{
		FirstName: "Sophie",
		LastName: "DeBenedetto",
	}
	request := &greetpb.GreetRequest{Greeting: greeting}
	resp, err := c.Greet(context.Background(), request)
	if err != nil {
		log.Fatalf("Client received error from server: %v", err)
	}
	fmt.Println(resp.GetResult())
}