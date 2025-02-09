package service

import (
	"context"

	"github.com/lab-icn/water-potability-sensor-service/internal/domain"
	pb "github.com/lab-icn/water-potability-sensor-service/internal/interface/rpc"
	"github.com/lab-icn/water-potability-sensor-service/internal/repository"
	"github.com/rs/zerolog"
)

type WaterPotabilityService struct {
	repository repository.WaterPotabilityRepositoryItf
	client     pb.WaterPotabilityServiceClient
	logger     zerolog.Logger
}

type WaterPotabilityServiceItf interface {
	PredictWaterPotability(ctx context.Context, wp domain.WaterPotability) error
}

func NewWaterPotabilityService(repository repository.WaterPotabilityRepositoryItf, client pb.WaterPotabilityServiceClient, logger zerolog.Logger) *WaterPotabilityService {
	return &WaterPotabilityService{repository, client, logger}
}

func (s *WaterPotabilityService) PredictWaterPotability(ctx context.Context, wp domain.WaterPotability) error {
	res, err := s.client.PredictWaterPotability(ctx, &pb.PredictWaterPotabilityRequest{
		Ph:                  wp.PH,
		TotalDissolveSolids: wp.TotalDissolvedSolids,
		Turbidity:           wp.Turbidity,
	})
	if err != nil {
		return err
	}

	s.logger.Info().
		Str("grpc", "model-water-potability").
		Any("data", wp).
		Float64("prediction", res.GetPrediction()).
		Msg("success predict water potability")

	data := domain.WaterPotabilityWithPrediction{
		Node:                 wp.Node,
		PH:                   wp.PH,
		TotalDissolvedSolids: wp.TotalDissolvedSolids,
		Turbidity:            wp.Turbidity,
		Prediction:           res.Prediction,
	}

	err = s.repository.WriteWaterPotabilityWithPrediction(ctx, data)
	if err != nil {
		return err
	}

	s.logger.Info().
		Str("influxdb", "wp").
		Any("data", data).
		Msg("success write water potability with prediction")

	return nil
}
