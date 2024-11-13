package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	pb "github.com/lab-icn/water-potability-sensor-service/internal/water_potability/interface/rpc"
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
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	server := grpc.NewServer()
	pb.RegisterWaterPotabilityServiceServer(server, &WaterPotabilityServer{})
	listener, err := net.Listen(os.Getenv("GRPC_PROTOCOL"), os.Getenv("GRPC_ADDR"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server running on %s\n", listener.Addr().String())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
