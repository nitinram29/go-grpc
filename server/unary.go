package main

import (
	"context"

	pb "github.com/nitinram29/go-grpc/proto"
)

func (s *learningServer) UnaryCall(ctx context.Context, req *pb.NoParam) (*pb.AckResponse, error) {
	return &pb.AckResponse{
		Message: "Hello Bro!!",
	}, nil;
}
