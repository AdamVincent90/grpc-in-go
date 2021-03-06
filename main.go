package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"./src/calculate"
	"./src/greet"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Running Applicaiton..")
	const (
		target = "localhost:50051"
	)

	creds, credErr := credentials.NewClientTLSFromFile("./ssl/ca.crt", "")

	if credErr != nil {
		log.Fatalln("Error getting TLS Cerfication", credErr)
	}

	cc, err := grpc.Dial(target, grpc.WithTransportCredentials(creds))

	if err != nil {
		fmt.Println("Error validating SSL Certification...", err)
		fmt.Println("Running API with insecure..")
		cc, err = grpc.Dial(target, grpc.WithInsecure())
		if err != nil {
			log.Fatalln("Cannot load API with insecure, terminating program.", err)
		}
	}

	defer cc.Close()

	c := greet.NewGreetServiceClient(cc)
	//ca := calculate.NewCalculateServiceClient(cc)

	unaryGreetRequest(c, "Nikolai", "Ifrim")

	//unaryAddRequest(ca, 25, 10)

	//serverStreamGreetRequest(c, "Adam")

	//serverStreamPrimeNumberRequest(ca, 1087867680)

	//clientStreamingGreetRequest(c)

	//clientStreamingAverageNumber(ca)

	//BidirectionalStream(c)

	//BiDrectionalMaxNumber(ca)

	// UnarySquareRoot(ca, 10)
	// UnarySquareRoot(ca, -10)
	// UnaryWithDeadline(c, 5*time.Second)
	// UnaryWithDeadline(c, 2*time.Second)
}

func UnaryWithDeadline(c greet.GreetServiceClient, timer time.Duration) {
	user := &greet.DeadlineRequest{
		Greeting: &greet.Greeting{
			FirstName: "Ricky",
			LastName:  "Champ",
		},
	}
	ctx, ctxFunc := context.WithTimeout(context.Background(), timer)
	defer ctxFunc()

	res, err := c.Deadline(ctx, user)

	if err != nil {
		err2, ok := status.FromError(err)
		if ok {
			if err2.Code() == codes.DeadlineExceeded {
				fmt.Println("Response exceeded deadline, cancelling server response")
				fmt.Println(err2.Message())
				return
			} else {
				fmt.Println("Unexpected error:", err2.Message())
				return
			}
		} else {
			log.Fatalln(err2)
		}
	}

	fmt.Println("Received response within deadline", res)
}

func UnarySquareRoot(ca calculate.CalculateServiceClient, number float32) {

	req := &calculate.SquareRootRequest{
		Number: number,
	}

	res, err := ca.SquareRoot(context.Background(), req)

	if err != nil {
		resErr, ok := status.FromError(err)

		if ok {
			fmt.Println(resErr.Message())
			fmt.Println(resErr.Code())
			if resErr.Code() == codes.InvalidArgument {
				fmt.Println("Argument in Request - INVALID ARGUMENT")
				return
			}
			return
		} else {
			log.Fatalln(err)
		}

	}

	fmt.Println("Received result from server", res.GetResult())
}

func BiDrectionalMaxNumber(ca calculate.CalculateServiceClient) {
	stream, err := ca.MaxNumber(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	wait := make(chan struct{})

	req := []int32{1, 10, 3, 4, 5, 11, 10, 10, 8, 12}

	// GO ROUTINE TO STREAM MESSAGES TO SERVER (FUNCTIONAL LITERAL)
	go func() {
		for _, r := range req {
			fmt.Println("Streaming number to server..", r)
			err := stream.Send(&calculate.StreamNumberRequest{
				Num: r,
			})
			time.Sleep(1000 * time.Millisecond)
			if err == io.EOF {
				fmt.Println("END OF FILE")
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
		}

		err = stream.CloseSend()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	// GO ROUTINE TO RECEIVE A STREAM OF MESSAGES FROM SERVER (FUNCTION LITERAL)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("END OF FILE")
				break
			}
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println("The New Max Number is:", res.GetNum())
		}
		close(wait)
	}()

	fmt.Println(<-wait)
}

func BidirectionalStream(c greet.GreetServiceClient) {
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	wait := make(chan struct{})

	req := []*greet.GreetEveryoneRequest{
		{
			Greeting: &greet.Greeting{
				FirstName: "Adam",
			},
		},
		{
			Greeting: &greet.Greeting{
				FirstName: "Alex",
			},
		},
		{
			Greeting: &greet.Greeting{
				FirstName: "Nikolai",
			},
		},
		{
			Greeting: &greet.Greeting{
				FirstName: "Ricky",
			},
		},
	}

	// GO ROUTINE TO STREAM MESSAGES TO SERVER (FUNCTIONAL LITERAL)
	go func() {
		for _, r := range req {
			fmt.Printf("Sending client stream request to server stream %v\n", r.GetGreeting().FirstName)
			err := stream.Send(r)
			if err != nil {
				log.Fatalln(err)
			}
			time.Sleep(1000 * time.Millisecond)
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	// GO ROUTINE TO RECEIVE A STREAM OF MESSAGES FROM SERVER (FUNCTION LITERAL)
	go func() {

		for {
			res, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("END OF FILE")
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println("received", res)
		}

		close(wait)
	}()

	fmt.Println(<-wait)
}

func clientStreamingAverageNumber(csc calculate.CalculateServiceClient) {
	res, err := csc.AverageNumber(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	req := []*calculate.ClientStreamNumberRequest{
		{
			Num: 10,
		},
		{
			Num: 43,
		},
		{
			Num: 49,
		},
		{
			Num: 12,
		},
		{
			Num: 102,
		},
	}

	for _, r := range req {
		fmt.Println("Sending", r.Num, "to sever")
		err := res.Send(r)
		time.Sleep(1000 * time.Millisecond)
		if err != nil {
			log.Fatalln(err)
		}
	}

	resp, err := res.CloseAndRecv()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Response from server averaging all client streaming numbers:", resp.Num)
}

func clientStreamingGreetRequest(gsc greet.GreetServiceClient) {

	//firstnames := []string{"Adam", "John", "Robbie", "Serin", "Nikolai", "Maruki"}

	res, error := gsc.LongGreet(context.Background())

	if error != nil {
		log.Fatalln(error)
	}

	req := []*greet.LongGreetRequest{
		{
			Greeting: &greet.Greeting{
				FirstName: "Adam",
			},
		},
		{
			Greeting: &greet.Greeting{
				FirstName: "Tim",
			},
		},
		{
			Greeting: &greet.Greeting{
				FirstName: "Alex",
			},
		},
	}

	for _, r := range req {

		err := res.Send(r)

		if err == io.EOF {
			fmt.Println("END OF FILE")
		}
		if err != nil {
			log.Fatalln(err)
		}
	}

	sr, err := res.CloseAndRecv()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(sr)
}

func unaryGreetRequest(gsc greet.GreetServiceClient, s1 string, s2 string) {
	req := &greet.GreetRequest{
		Greeting: &greet.Greeting{
			FirstName: s1,
			LastName:  s2,
		},
	}

	res, err := gsc.Greet(context.Background(), req)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}

func unaryAddRequest(gsc calculate.CalculateServiceClient, foo int32, bar int32) {
	req := &calculate.AddRequest{
		Sums: &calculate.Sums{
			Sum1: foo,
			Sum2: bar,
		},
	}

	res, err := gsc.Add(context.Background(), req)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}

func serverStreamGreetRequest(gsc greet.GreetServiceClient, name string) {

	req := &greet.GreetManyTimesRequest{
		Greeting: &greet.Greeting{
			FirstName: name,
		},
	}
	resStream, err := gsc.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalln(err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			fmt.Println("End of stream..")
			break
		}
		if err != nil {
			log.Fatalln("Error during stream", err)
		}
		fmt.Println(msg)
	}
}

func serverStreamPrimeNumberRequest(psc calculate.CalculateServiceClient, num int32) {

	req := &calculate.PrimeNumberRequest{
		Num: num,
	}

	resStream, err := psc.PrimeDecomposition(context.Background(), req)

	if err != nil {
		log.Fatalln(err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			fmt.Println("END OF STREAM")
		}
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(msg)
	}

}
