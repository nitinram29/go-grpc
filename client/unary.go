package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nitinram29/go-grpc/proto"
)

func callUnaryCall(client pb.LearnServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.UnaryCall(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatal("HE %v",err)
	}
	log.Printf("%s", res.Message)
}