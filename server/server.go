package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"../src/calculate"
	"../src/greet"
	"google.golang.org/grpc"
)

type server struct{}

// Server application (THIS CAN BE CONSIDERED A COMPLETLEY SEPERATE APP TO MAIN.GO)
func main() {

	const (
		protocol = "tcp"             // grpc uses TCP PROTOCOL
		address  = "127.0.0.1:50051" // 50051 port represents GRPC PORT
	)

	nl, err := net.Listen(protocol, address) // TAKES IN PROTOCOL AND ADDRESS, LISTENS ON THAT PORT FOR SERVICES
	if err != nil {
		log.Fatalln("Error establishing a", protocol, "connection on", address, "\nerr")
	}

	fmt.Println("Server Running at:", address) // if no error then specify to console server is running at address

	s := grpc.NewServer() // creates a pointer to the grpc.server

	greet.RegisterGreetServiceServer(s, &server{}) // registers a server with the grpc server s and the services
	calculate.RegisterCalculateServiceServer(s, &server{})

	if err := s.Serve(nl); err != nil {
		log.Fatalln(err)
	}
}

func (*server) Greet(ctx context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {

	fmt.Println(req, "invoked")

	name := req.GetGreeting().GetFirstName() + " " + req.GetGreeting().GetLastName()
	result := "Hello " + name

	res := &greet.GreetResponse{
		Result: result,
	}

	return res, nil
}

func (*server) Add(ctx context.Context, req *calculate.AddRequest) (*calculate.AddResponse, error) {
	fmt.Println(req, "invoked")

	sum := req.Sums.GetSum1() + req.Sums.GetSum2()

	res := &calculate.AddResponse{
		Result: sum,
	}

	return res, nil
}
