package aggregator

import (
	"context"
	"log"
	"sync"

	"github.com/Poin1961/federated-learning-system/pkg/proto"
)

type FederatedAggregatorServer struct {
	proto.UnimplementedFederatedLearningServer
	mu            sync.Mutex
	globalModel   map[string]float32
	clientUpdates []*proto.ModelUpdate
}

func NewFederatedAggregatorServer() *FederatedAggregatorServer {
	return &FederatedAggregatorServer{
		globalModel: make(map[string]float32),
	}
}

func (s *FederatedAggregatorServer) SendModelUpdate(ctx context.Context, in *proto.ModelUpdate) (*proto.Acknowledgement, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Printf("Received model update from client %s", in.ClientId)
	s.clientUpdates = append(s.clientUpdates, in)

	// Simple aggregation: average weights (in a real system, this would be more complex)
	for key, value := range in.Weights {
		s.globalModel[key] += value
	}

	log.Printf("Current aggregated global model (sum): %v", s.globalModel)

	return &proto.Acknowledgement{Message: "Model update received and aggregated"}, nil
}

func (s *FederatedAggregatorServer) GetGlobalModel(ctx context.Context, in *proto.GlobalModelRequest) (*proto.GlobalModel, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	log.Printf("Client %s requested global model", in.ClientId)

	// In a real system, this would return the actual aggregated global model
	return &proto.GlobalModel{
		Weights: s.globalModel,
	},
	
}
