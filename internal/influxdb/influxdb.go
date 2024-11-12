package influxdb

import (
	"context"
	"fmt"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func NewClient(ctx context.Context) (influxdb2.Client, error) {
	uri := fmt.Sprintf(
		"%s://%s:%s",
		os.Getenv("INFLUXDB_PROTOCOL"),
		os.Getenv("INFLUXDB_HOST"),
		os.Getenv("INFLUXDB_PORT"),
	)
	client := influxdb2.NewClient(uri, os.Getenv("INFLUXDB_TOKEN"))
	if ok, err := client.Ping(ctx); err != nil || !ok {
		return nil, err
	}
	return client, nil
}
