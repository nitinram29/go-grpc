package main

import (
	"log"
	"time"

	pb "github.com/nitinram29/go-grpc/proto"
)

func (s *learningServer) ServerStreamingCall(req *pb.NamesList, stream pb.LearnService_ServerStreamingCallServer) error {
	log.Printf("Got req with names: %v", req.Name)

	for _, name := range req.Name {
		res := &pb.AckResponse{
			Message: "Hello, " + name,
		}
		if err := stream.Send(res); err != nil {
			log.Fatal(err)
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
