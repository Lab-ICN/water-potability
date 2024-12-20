package main

import (
	"bytes"
	"context"
	"encoding/json"
	stdlog "log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lab-icn/water-potability-sensor-service/internal/config"
	"github.com/lab-icn/water-potability-sensor-service/internal/domain"
	"github.com/lab-icn/water-potability-sensor-service/internal/grpc"
	"github.com/lab-icn/water-potability-sensor-service/internal/influxdb"
	mqttAdapter "github.com/lab-icn/water-potability-sensor-service/internal/interface/mqtt"
	pb "github.com/lab-icn/water-potability-sensor-service/internal/interface/rpc"
	_mqtt "github.com/lab-icn/water-potability-sensor-service/internal/mqtt"
	"github.com/lab-icn/water-potability-sensor-service/internal/repository"
	"github.com/lab-icn/water-potability-sensor-service/internal/service"
	"github.com/rs/zerolog"
)

func main() {
	path := os.Getenv("CONFIG_FILEPATH")
	content, err := os.ReadFile(path)
	if err != nil {
		stdlog.Fatalf("opening config file at %s: %v\n", path, err)
	}
	cfg := new(config.Config)
	if err := json.Unmarshal(content, cfg); err != nil {
		stdlog.Fatalf("parsing config file content: %v\n", err)
	}
	ctx := context.Background()

	log := zerolog.
		New(os.Stderr).
		With().
		Timestamp().
		Logger().
		Level(zerolog.DebugLevel)

	grpcClient, err := grpc.NewClient(cfg)
	if err != nil {
		stdlog.Fatalf("Failed to start gRPC connection: %v\n", err)
	}
	defer grpcClient.Close()

	mqttClient, err := _mqtt.NewClient(cfg, &log)
	if err != nil {
		stdlog.Fatalf("Failed to start MQTT connection: %v\n", err)
	}
	defer mqttClient.Disconnect(250)

	influxdb, err := influxdb.NewClient(ctx, &cfg.InfluxDB)
	if err != nil {
		stdlog.Fatalf("Failed to start InfluxDB connection: %v\n", err)
	}
	defer influxdb.Close()

	wpClient := pb.NewWaterPotabilityServiceClient(grpcClient)
	wpRepository := repository.NewWaterPotabilityRepository(influxdb, &cfg.InfluxDB)
	wpService := service.NewWaterPotabilityService(wpRepository, wpClient)
	mqttAdapter.NewMqttHandler(mqttClient, &cfg.AES, &log, wpService)

	stdlog.Println("client server running...")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGILL, syscall.SIGTERM)
	done := make(chan struct{}, 1)
	go func() {
		<-sig
		stdlog.Println("shutting down...")
		done <- struct{}{}
	}()
	<-done
	stdlog.Println("exiting...")
}

func mockPublisher(client mqtt.Client) error {
	buffer := new(bytes.Buffer)
	if err := json.NewEncoder(buffer).Encode(domain.WaterPotability{
		PH:                   7.5,
		Turbidity:            5.5,
		TotalDissolvedSolids: 100,
	}); err != nil {
		return err
	}
	for range 1000 {
		token := client.Publish("/wp", 1, false, buffer.Bytes())
		token.Wait()
		time.Sleep(1 * time.Second)
	}
	return nil
}
