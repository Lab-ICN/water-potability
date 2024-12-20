package service

import (
	"context"

	"github.com/lab-icn/water-potability-sensor-service/internal/domain"
	pb "github.com/lab-icn/water-potability-sensor-service/internal/interface/rpc"
	"github.com/lab-icn/water-potability-sensor-service/internal/repository"
)

type WaterPotabilityService struct {
	repository repository.WaterPotabilityRepositoryItf
	client     pb.WaterPotabilityServiceClient
}

type WaterPotabilityServiceItf interface {
	PredictWaterPotability(ctx context.Context, wp domain.WaterPotability) error
}

func NewWaterPotabilityService(repository repository.WaterPotabilityRepositoryItf, client pb.WaterPotabilityServiceClient) *WaterPotabilityService {
	return &WaterPotabilityService{repository, client}
}

func (s *WaterPotabilityService) PredictWaterPotability(ctx context.Context, wp domain.WaterPotability) error {
	// res, err := s.client.PredictWaterPotability(ctx, &pb.PredictWaterPotabilityRequest{
	// 	Ph:                  wp.PH,
	// 	TotalDissolveSolids: wp.TotalDissolvedSolids,
	// 	Turbidity:           wp.Turbidity,
	// })
	// if err != nil {
	// 	return err
	// }

	return s.repository.WriteWaterPotabilityWithPrediction(ctx, domain.WaterPotabilityWithPrediction{
		PH:                   wp.PH,
		TotalDissolvedSolids: wp.TotalDissolvedSolids,
		Turbidity:            wp.Turbidity,
		// Prediction:           res.Prediction,
	})
}
