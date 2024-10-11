package service

import (
	"context"

	"github.com/lab-icn/water-potability-sensor-service/internal/domain"
	pb "github.com/lab-icn/water-potability-sensor-service/internal/water_potability/interface/rpc"
	"github.com/lab-icn/water-potability-sensor-service/internal/water_potability/repository"
)

type WaterPotabilityService struct {
	repository          repository.WaterPotabilityRepositoryItf
	grpcWaterPotability pb.WaterPotabilityServiceClient
}

type WaterPotabilityServiceItf interface {
	PredictWaterPotability(ctx context.Context, wp domain.WaterPotability) error
}

func NewWaterPotabilityService(repository repository.WaterPotabilityRepositoryItf, grpcWaterPotability pb.WaterPotabilityServiceClient) *WaterPotabilityService {
	return &WaterPotabilityService{repository: repository, grpcWaterPotability: grpcWaterPotability}
}

func (s *WaterPotabilityService) PredictWaterPotability(ctx context.Context, wp domain.WaterPotability) error {
	res, err := s.grpcWaterPotability.PredictWaterPotability(ctx, &pb.PredictWaterPotabilityRequest{
		Ph:                  wp.PH,
		TotalDissolveSolids: wp.TotalDissolvedSolids,
		Turbidity:           wp.Turbidity,
	})

	if err != nil {
		return err
	}

	return s.repository.WriteWaterPotabilityWithPrediction(ctx, domain.WaterPotabilityWithPrediction{
		PH:                   wp.PH,
		TotalDissolvedSolids: wp.TotalDissolvedSolids,
		Turbidity:            wp.Turbidity,
		Prediction:           res.Prediction,
		Level:                res.Level,
	})
}
