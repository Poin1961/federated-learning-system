package main

import (
	"fmt"
	"log"
	"net"

	"federated-learning-system/pkg/aggregator"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Federated Learning Server...")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	aggregator.RegisterAggregatorServer(s, &aggregator.Server{})

	fmt.Printf("Server listening on %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
