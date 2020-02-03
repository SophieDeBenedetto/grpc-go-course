package main

import (
	context "context"
	"fmt"
	"io"
	"log"

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
	doStreaming(c)
	// doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Client making Unary call...")
	greeting := &greetpb.Greeting{
		FirstName: "Sophie",
		LastName:  "DeBenedetto",
	}
	request := &greetpb.GreetRequest{Greeting: greeting}
	resp, err := c.Greet(context.Background(), request)
	if err != nil {
		log.Fatalf("Client received error from server: %v", err)
	}
	fmt.Println(resp.GetResult())
}

func doStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Client making streaming request...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sophie",
			LastName:  "Huang",
		},
	}
	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Client received error from stream: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// end of stream
			fmt.Println("Done streaming.")
			break
		}
		if err != nil {
			log.Fatalf("Error reading from stream: %v", err)
		}
		log.Printf("Response from stream: %v", msg)
	}

}
