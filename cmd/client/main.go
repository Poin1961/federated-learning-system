package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Poin1961/federated-learning-system/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Starting Federated Learning Client...")

	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		serverAddr = "localhost:50051"
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewFederatedLearningClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Simulate client sending local model updates
	localModel := &proto.ModelUpdate{
		ClientId: "client-1",
		Weights:  map[string]float32{"layer1_w": 0.1, "layer1_b": 0.01},
		Metrics:  map[string]float32{"accuracy": 0.85, "loss": 0.15},
	}

	r, err := c.SendModelUpdate(ctx, localModel)
	if err != nil {
		log.Fatalf("Could not send model update: %v", err)
	}
	log.Printf("Server Response: %s", r.GetMessage())

	// Simulate client requesting global model
	globalModelReq := &proto.GlobalModelRequest{
		ClientId: "client-1",
	}

	g, err := c.GetGlobalModel(ctx, globalModelReq)
	if err != nil {
		log.Fatalf("Could not get global model: %v", err)
	}
	log.Printf("Global Model Weights: %v", g.GetWeights())
}
