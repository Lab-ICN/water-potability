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
	PredictWaterPotabilityData(ctx context.Context, wp domain.WaterPotability) error
}

func NewWaterPotabilityService(repository repository.WaterPotabilityRepositoryItf) WaterPotabilityService {
	return WaterPotabilityService{repository: repository}
}

func (s *WaterPotabilityService) PredictWaterPotabilityData(ctx context.Context, wp domain.WaterPotability) error {
	res, err := s.grpcWaterPotability.PredictWaterPotability(ctx, &pb.PredictWaterPotabilityRequest{
		Id:                  wp.ID,
		Ph:                  wp.PH,
		TotalDissolveSolids: wp.TotalDissolvedSolids,
		Turbidity:           wp.Turbidity,
	})

	if err != nil {
		return err
	}

	return s.repository.WriteWaterPotabilityWithPrediction(ctx, domain.WaterPotabilityWithPrediction{
		ID:                   wp.ID,
		PH:                   wp.PH,
		TotalDissolvedSolids: wp.TotalDissolvedSolids,
		Turbidity:            wp.Turbidity,
		Prediction:           res.Prediction,
	})
}
