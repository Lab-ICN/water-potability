package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/lab-icn/water-potability-sensor-service/internal/config"
	pb "github.com/lab-icn/water-potability-sensor-service/internal/interface/rpc"
	"google.golang.org/grpc"
)

type WaterPotabilityServer struct {
	pb.UnimplementedWaterPotabilityServiceServer
}

func (s *WaterPotabilityServer) PredictWaterPotability(ctx context.Context, in *pb.PredictWaterPotabilityRequest) (*pb.PredictWaterPotabilityResponse, error) {
	return &pb.PredictWaterPotabilityResponse{
		Prediction: 100,
	}, nil
}

func main() {
	path := os.Getenv("CONFIG_FILEPATH")
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("opening config file at %s: %v\n", path, err)
	}
	cfg := new(config.Config)
	if err := json.Unmarshal(content, cfg); err != nil {
		log.Fatalf("parsing config file content: %v\n", err)
	}

	server := grpc.NewServer()
	pb.RegisterWaterPotabilityServiceServer(server, &WaterPotabilityServer{})
	listener, err := net.Listen(cfg.GRPC.Protocol, fmt.Sprintf(
		"%s:%d",
		cfg.GRPC.Host,
		cfg.GRPC.Port,
	))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server running on %s\n", listener.Addr().String())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
