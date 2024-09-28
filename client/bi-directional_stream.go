package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/nitinram29/go-grpc/proto"
)

func callBiDirectionalStreamingCall(client pb.LearnServiceClient, names pb.NamesList) {
	log.Printf("Streaming started")
	stream, err := client.BiDirectionalStreamingCall(context.Background())
	if err != nil {
		log.Fatal("Failed ", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal("Failed ", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Name {
		req := &pb.UserRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatal("Failed ", err)
		}
		time.Sleep(time.Second)
	}
	stream.CloseSend()
	<-waitc // whats the use of this line
	log.Printf("Bi Directional streaming finised")
}