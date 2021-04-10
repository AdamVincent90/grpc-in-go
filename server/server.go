package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

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

func (*server) MaxNumber(css calculate.CalculateService_MaxNumberServer) error {
	fmt.Println("Bi-drectional streaming initiated")
	result := int32(0)
	for {
		req, err := css.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalln(err)
		}

		if result <= req.GetNum() {
			result = req.GetNum()
			res := &calculate.StreamNumberResponse{
				Num: result,
			}
			fmt.Println("Sending New Max Number..", result, "to client..")
			err = css.Send(res)
		}

		if err != nil {
			log.Fatalln(err)
		}

	}
}

func (*server) GreetEveryone(stream greet.GreetService_GreetEveryoneServer) error {
	fmt.Println("Bi-directional streaming invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalln(err)
		}
		fn := req.GetGreeting().GetFirstName()
		result := "Hello " + fn

		res := &greet.GreetEveryoneResponse{
			Result: result,
		}
		fmt.Println("Sending Stream response", result, "to client stream")
		err = stream.Send(res)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (*server) AverageNumber(stream calculate.CalculateService_AverageNumberServer) error {
	counter := 0
	sum := float32(0)
	for {
		msg, err := stream.Recv()
		fmt.Println("Received message from client stream", counter, sum)
		if err == io.EOF {
			fmt.Println("END OF FILE")
			sum = sum / float32(counter)
			return stream.SendAndClose(&calculate.ClientStreamNumberResponse{
				Num: sum,
			})
		}
		if err != nil {
			log.Fatalln(err)
		}
		sum += msg.GetNum()
		counter += 1
	}
}

func (*server) LongGreet(stream greet.GreetService_LongGreetServer) error {

	fmt.Println("Client stream invoked..")
	var result string

	for {
		s, err := stream.Recv()
		fmt.Println("received message from client stream", s)

		if err == io.EOF {
			fmt.Println("END OF FILE")
			return stream.SendAndClose(&greet.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			return err
		}

		fname := s.GetGreeting().GetFirstName()

		result += "Hello " + fname + " Nice to meet you!\n"

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

func (*server) GreetManyTimes(req *greet.GreetManyTimesRequest, stream greet.GreetService_GreetManyTimesServer) error {
	fmt.Println(req, "invoked")

	name := req.GetGreeting().GetFirstName()

	for i := 0; i < 10; i++ {
		result := "Hello There " + name + ", this response has been sent to you " + strconv.Itoa(i) + " times!"
		res := &greet.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (*server) PrimeDecomposition(req *calculate.PrimeNumberRequest, stream calculate.CalculateService_PrimeDecompositionServer) error {
	fmt.Println(req, "invoked")
	var k int32
	number := req.GetNum()

	k = 2

	for number > 1 {
		if number%k == 0 {

			res := &calculate.PrimeNumberResponse{
				Result: k,
			}
			stream.Send(res)
			number = number / k
			time.Sleep(500 * time.Millisecond)
		} else {
			k = k + 1
		}
	}
	return nil
}
