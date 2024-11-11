package main

import (
	"context"
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lab-icn/water-potability-sensor-service/internal/infra"
	_mqtt "github.com/lab-icn/water-potability-sensor-service/internal/water_potability/interface/mqtt"
	pb "github.com/lab-icn/water-potability-sensor-service/internal/water_potability/interface/rpc"
	"github.com/lab-icn/water-potability-sensor-service/internal/water_potability/repository"
	"github.com/lab-icn/water-potability-sensor-service/internal/water_potability/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Fail to create logger instance: %v\n", err)
	}
	defer logger.Sync()
	grpcClient, err := grpc.NewClient(
		os.Getenv("GRPC_ADDR"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Fail to start gRPC connection: %v\n", err)
	}
	defer grpcClient.Close()
	mqttOpts := mqtt.NewClientOptions().
		AddBroker(fmt.Sprintf("%s://%s:%s", os.Getenv("MQTT_PROTOCOL"), os.Getenv("MQTT_HOST"), os.Getenv("MQTT_PORT"))).
		SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
			logger.Info("message received", zap.String("topic", msg.Topic()), zap.ByteString("payload", msg.Payload()))
		})
	mqttClient := mqtt.NewClient(mqttOpts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Connection failed to MQTT Broker: %v\n", token.Error())
	}
	influxdb := infra.NewInfluxDB(os.Getenv("INFLUXDB_URL"), os.Getenv("INFLUXDB_TOKEN"))
	if running, err := influxdb.Ping(ctx); err != nil || !running {
		log.Fatalf("Fail to ping InfluxDB: %v\n", err)
	}
	defer influxdb.Close()
	wpClient := pb.NewWaterPotabilityServiceClient(grpcClient)
	wpRepository := repository.NewWaterPotabilityRepository(influxdb)
	wpService := service.NewWaterPotabilityService(wpRepository, wpClient)
	_mqtt.NewMQTT(wpService)
}
