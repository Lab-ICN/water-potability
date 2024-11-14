package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lab-icn/water-potability-sensor-service/internal/domain"
	"github.com/lab-icn/water-potability-sensor-service/internal/grpc"
	"github.com/lab-icn/water-potability-sensor-service/internal/influxdb"
	mqttAdapter "github.com/lab-icn/water-potability-sensor-service/internal/interface/mqtt"
	pb "github.com/lab-icn/water-potability-sensor-service/internal/interface/rpc"
	_mqtt "github.com/lab-icn/water-potability-sensor-service/internal/mqtt"
	"github.com/lab-icn/water-potability-sensor-service/internal/repository"
	"github.com/lab-icn/water-potability-sensor-service/internal/service"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create logger instance: %v\n", err)
	}
	defer logger.Sync()

	grpcClient, err := grpc.NewClient()
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v\n", err)
	}
	defer grpcClient.Close()

	mqttClient, err := _mqtt.NewClient(logger)
	if err != nil {
		log.Fatalf("Failed to start MQTT connection: %v\n", err)
	}
	defer mqttClient.Disconnect(250)

	influxdb, err := influxdb.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to start InfluxDB connection: %v\n", err)
	}
	defer influxdb.Close()

	wpClient := pb.NewWaterPotabilityServiceClient(grpcClient)
	wpRepository := repository.NewWaterPotabilityRepository(influxdb)
	wpService := service.NewWaterPotabilityService(wpRepository, wpClient)
	mqttAdapter.NewMqttHandler(mqttClient, logger, wpService)

	log.Println("client server running...")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGILL, syscall.SIGTERM)
	done := make(chan struct{}, 1)
	go func() {
		<-sig
		log.Println("shutting down...")
		done <- struct{}{}
	}()
	<-done
	log.Println("exiting...")
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
		token := client.Publish("wp", 1, false, buffer.Bytes())
		token.Wait()
		time.Sleep(1 * time.Second)
	}
	return nil
}
