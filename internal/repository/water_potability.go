package repository

import (
	"context"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/lab-icn/water-potability-sensor-service/internal/domain"
)

type WaterPotabilityRepository struct {
	influxdb       influxdb2.Client
	orgInfluxdb    string
	bucketInfluxdb string
}

type WaterPotabilityRepositoryItf interface {
	WriteWaterPotabilityWithPrediction(ctx context.Context, wp domain.WaterPotabilityWithPrediction) error
}

func NewWaterPotabilityRepository(influxdb influxdb2.Client) *WaterPotabilityRepository {
	orgInfluxdb := os.Getenv("INFLUXDB_ORG")
	bucketInfluxdb := os.Getenv("INFLUXDB_BUCKET")

	return &WaterPotabilityRepository{
		influxdb:       influxdb,
		orgInfluxdb:    orgInfluxdb,
		bucketInfluxdb: bucketInfluxdb,
	}
}

func (r *WaterPotabilityRepository) WriteWaterPotabilityWithPrediction(ctx context.Context, wp domain.WaterPotabilityWithPrediction) error {
	writeApi := influxdb2.NewPointWithMeasurement("water_potability_with_prediction").
		AddField("prediction", wp.Prediction).
		AddField("ph", wp.PH).
		AddField("total_dissolved_solids", wp.TotalDissolvedSolids).
		AddField("turbidity", wp.Turbidity).
		AddField("level", wp.Level).
		SetTime(time.Now())

	return r.influxdb.WriteAPIBlocking(r.orgInfluxdb, r.bucketInfluxdb).WritePoint(ctx, writeApi)
}
