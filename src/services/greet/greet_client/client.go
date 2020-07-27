package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"services/greet/greetpb"
)

func callGreet(client greetpb.GreetServiceClient)  {
	fmt.Println("In the callGreet function")

	req := greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Swapnil",
			LastName: "Pakhare",
		},
	}

	res, err := client.Greet(context.Background(), &req)

	if err != nil {
		log.Fatalf("Error in sending request : %v", err.Error())
	}

	fmt.Println("Result is ", res.Result)
}

func callGreetFullName(client greetpb.GreetServiceClient)  {
	fmt.Println("In the callGreetFullName function")

	req := greetpb.GreetRequest{Greeting: &greetpb.Greeting{FirstName: "Swapnil", LastName: "Pakhare"}}

	res, err := client.GreetFullName(context.Background(), &req)

	if err != nil {
		log.Fatalf("Error in sending full name request : %v", err.Error())
	}

	fn := res.Greet.FirstName
	ln := res.Greet.LastName

	result := fn + " " + ln

	fmt.Println("Response of full name is ", result)
}

func main() {
	fmt.Println("Client is online...")

	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Error in client : ", err.Error())
	}

	defer conn.Close()

	client := greetpb.NewGreetServiceClient(conn)

	callGreet(client)
	callGreetFullName(client)
}
