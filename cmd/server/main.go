package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Poin1961/federated-learning-system/pkg/aggregator"
	"github.com/Poin1961/federated-learning-system/pkg/proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Federated Learning Server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server := aggregator.NewFederatedAggregatorServer()
	proto.RegisterFederatedLearningServer(s, server)

	go func() {
		fmt.Printf("Server listening on :%s\n", port)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop // Wait for interrupt signal

	fmt.Println("Shutting down server...")
	s.GracefulStop()
	fmt.Println("Server gracefully stopped.")
}
