package influxdb

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/lab-icn/water-potability-sensor-service/internal/config"
)

func NewClient(ctx context.Context, cfg *config.InfluxDB) (influxdb2.Client, error) {
	client := influxdb2.NewClient(fmt.Sprintf(
		"%s://%s:%d",
		cfg.Protocol,
		cfg.Host,
		cfg.Port,
	), cfg.Token)
	if ok, err := client.Ping(ctx); err != nil || !ok {
		return nil, err
	}
	return client, nil
}
