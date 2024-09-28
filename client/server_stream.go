package main

import (
	"context"
	"io"
	"log"

	pb "github.com/nitinram29/go-grpc/proto"
)

func callServerStreamingCall(client pb.LearnServiceClient, names *pb.NamesList) {

	stream, err := client.ServerStreamingCall(context.Background(), names)
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("err")
		}
		log.Printf(msg.Message)
	}
	log.Printf("Streaming ends")
}
