package infra

import influxdb2 "github.com/influxdata/influxdb-client-go/v2"

func NewInfluxDB(host, token string) influxdb2.Client {
	return influxdb2.NewClient(host, token)
}
