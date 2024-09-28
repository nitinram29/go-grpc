package main

import (
	"io"
	"log"

	pb "github.com/nitinram29/go-grpc/proto"
)

func (s *learningServer) BiDirectionalStreamingCall(stream pb.LearnService_BiDirectionalStreamingCallServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Failed ,", err)
		}
		log.Printf("Got req for %s", msg.Name)
		res := &pb.AckResponse{
			Message: "Great to see you, " + msg.Name,
		}
		if reqErr := stream.Send(res); reqErr != nil {
			log.Fatal("Failed ", reqErr)
			return reqErr
		}
	}
	return nil
}
