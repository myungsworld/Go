package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

const portNumber = "9000"

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %s", err)
	}

}