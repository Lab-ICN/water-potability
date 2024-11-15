package mqtt

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/lab-icn/water-potability-sensor-service/internal/aes"
	"github.com/lab-icn/water-potability-sensor-service/internal/domain"
	"github.com/lab-icn/water-potability-sensor-service/internal/service"
	"go.uber.org/zap"
)

type handler struct {
	service service.WaterPotabilityServiceItf
	logger  *zap.Logger
}

func NewMqttHandler(client mqtt.Client, logger *zap.Logger, service service.WaterPotabilityServiceItf) {
	handler := &handler{service, logger}
	if token := client.Subscribe("wp", 1, handler.sensorSubscriber); token.Wait() && token.Error() != nil {
		log.Fatalf("Error subscribing to topic: %v", token.Error())
	}
}

func (h *handler) sensorSubscriber(client mqtt.Client, msg mqtt.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	cipher, err := base64.StdEncoding.DecodeString(string(msg.Payload()))
	if err != nil {
		h.logger.Error(
			"failed to decode base64 encoded message",
			zap.String("topic", msg.Topic()),
			zap.ByteString("payload", msg.Payload()),
			zap.Error(err),
		)
		return
	}
	plaintext, err := aes.Decrypt(
		cipher,
		[]byte(os.Getenv("AES_KEY")),
		[]byte(os.Getenv("AES_IV")),
	)
	if err != nil {
		h.logger.Error(
			"failed to decrypt aes cipher",
			zap.ByteString("cipher", cipher),
			zap.Error(err),
		)
		return
	}
	h.logger.Debug(
		"decrypted aes message",
		zap.String("payload", plaintext),
	)

	var potability domain.WaterPotability
	if err := json.Unmarshal([]byte(plaintext), &potability); err != nil {
		h.logger.Error(
			"failed to unmarshal json to struct",
			zap.String("plaintext", plaintext),
			zap.Error(err),
		)
		return
	}
	if err := h.service.PredictWaterPotability(ctx, potability); err != nil {
		h.logger.Error("failed to predict water potability", zap.Error(err))
		return
	}
}
