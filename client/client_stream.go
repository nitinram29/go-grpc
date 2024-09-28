package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nitinram29/go-grpc/proto"
)

func callClientStreamingCall(client pb.LearnServiceClient, names pb.NamesList) {
	stream, err := client.ClientStreamingCall(context.Background())
	if err != nil {
		log.Fatal("Failed %v", err)
	}

	for _, name := range names.Name {
		req := &pb.UserRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatal("Failed %s", err)
		}
		log.Printf("Send the request for name : %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming done!")
	if err != nil {
		log.Fatal("Failed %v", err)
	}
	log.Printf("res, %s", res.Message)
}
