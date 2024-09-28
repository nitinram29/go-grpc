package main

import (
	"log"
	"net"

	pb "github.com/nitinram29/go-grpc/proto"
	"google.golang.org/grpc"
)

type learningServer struct {
	pb.LearnServiceServer
}

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLearnServiceServer(grpcServer, &learningServer{})
	log.Printf("Server started on port %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed in start : %v", err)
	}

}
