package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"./src/calculate"
	"./src/greet"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Running Applicaiton..")
	const (
		target = "127.0.0.1:50051"
	)

	cc, err := grpc.Dial(target, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("%+v", err)
	}

	defer cc.Close()

	c := greet.NewGreetServiceClient(cc)
	ca := calculate.NewCalculateServiceClient(cc)

	unaryGreetRequest(c, "Nikolai", "Ifrim")

	unaryAddRequest(ca, 25, 10)

	//serverStreamGreetRequest(c, "Adam")

	serverStreamPrimeNumberRequest(ca, 1080)

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
