package main

import (
	"context"
	"fmt"
	"log"
	pb "simple-api/gen/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)

	}

	client := pb.NewTestApiClient(conn)
	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Nsg: "hello vaibhav here"})
	if err != nil {
		log.Println(err)

	}
	fmt.Println(resp)
	fmt.Println(resp.Nsg)

}
