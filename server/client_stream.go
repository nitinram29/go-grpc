package main

import (
	"io"
	"log"

	pb "github.com/nitinram29/go-grpc/proto"
)

func (s *learningServer) ClientStreamingCall(stream pb.LearnService_ClientStreamingCallServer) error {
	var message []string
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Message: message})
		}
		if err != nil {
			log.Fatal("Failed ", (err==io.EOF))
		}
		log.Printf("Got request for name %s", msg.Name)
		message = append(message, "Hello, ", msg.Name)
	}
}
