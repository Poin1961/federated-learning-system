package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"federated-learning-system/pkg/aggregator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddr := flag.String("server-addr", "localhost:8080", "The server address in the format of host:port")
	dataPath := flag.String("data-path", "./data/client.csv", "Path to the client's local data file")
	flag.Parse()

	fmt.Printf("Starting Federated Learning Client. Connecting to %s with data from %s\n", *serverAddr, *dataPath)

	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := aggregator.NewAggregatorClient(conn)

	// Simulate local model training and update submission
	for i := 0; i < 5; i++ {
		fmt.Printf("\nClient training round %d...\n", i+1)
		// In a real scenario, this would involve loading data from dataPath,
		// training a local model, and generating model updates.
		localModelUpdate := &aggregator.ModelUpdate{
			ClientId: "client-1",
			Weights:  []float32{0.1 * float32(i), 0.2 * float32(i), 0.3 * float32(i)},
			Bias:     0.05 * float32(i),
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_, err := client.SubmitModelUpdate(ctx, localModelUpdate)
		if err != nil {
			log.Printf("Could not submit model update: %v", err)
		} else {
			fmt.Println("Model update submitted successfully.")
		}
		time.Sleep(2 * time.Second) // Simulate some work
	}

	fmt.Println("Client finished.")
}
