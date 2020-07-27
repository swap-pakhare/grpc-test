package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"services/greet/greetpb"
	"strings"
)

type server struct {

}

func (s server) Greet(ctx context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Function called")

	firstName := request.GetGreeting().FirstName
	lastName := request.GetGreeting().LastName

	result := "Hello " + firstName + " " + lastName

	return &greetpb.GreetResponse{Result: result}, nil
}

func (s server) GreetFullName(ctx context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Full Name function called")

	firstName := request.GetGreeting().FirstName
	lastName := request.GetGreeting().LastName

	results := "Hello " + strings.ToUpper(firstName + " " + lastName)

	return  &greetpb.GreetResponse{Result: results}, nil
}

func main() {
	fmt.Println("Listening on server...")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen, err is %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to listen, err is %v", err)
	}

}
