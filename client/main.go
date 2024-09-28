package main

import (
	"log"

	pb "github.com/nitinram29/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	
	client := pb.NewLearnServiceClient(conn);
	
	// Unary (Request-Response)
	// callUnaryCall(client);
	// log.Printf("\n\n")

	// // Server Stream
	names := &pb.NamesList {
		Name: []string{"Nitin", "Nikhil"},
	}
	// callServerStreamingCall(client, names);
	// log.Printf("\n\n")

	// // Client Stream
	// callClientStreamingCall(client, *names);
	// log.Printf("\n\n")	

	// Bi-Directional Stream
	callBiDirectionalStreamingCall(client, *names);
}
