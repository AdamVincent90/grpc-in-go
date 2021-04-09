package main

import (
	"context"
	"fmt"
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
