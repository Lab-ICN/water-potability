package mqtt

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
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
	log.Printf("Received message on topic '%s'", msg.Topic())
	log.Printf("Payload received: %s", string(msg.Payload()))

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ciphertext, err := base64.StdEncoding.DecodeString(string(msg.Payload()))
	if err != nil {
		log.Printf("Error decoding base64 payload: %v", err)
		log.Printf("Failed payload: %s", string(msg.Payload()))
		return
	}

	plaintext, err := aes.DecryptAESGCM(ciphertext)
	if err != nil {
		log.Printf("Error decrypting message payload: %v", err)
		log.Printf("Failed payload: %s", string(msg.Payload()))
		return
	}

	log.Printf("Decrypted payload: %s", plaintext)

	var potability domain.WaterPotability
	if err := json.Unmarshal([]byte(plaintext), &potability); err != nil {
		log.Printf("Error decoding JSON from decrypted payload: %v", err)
		log.Printf("Failed plaintext: %s", plaintext)
		return
	}

	if err := h.service.PredictWaterPotability(ctx, potability); err != nil {
		log.Printf("Error predicting water potability data: %v", err)
		return
	}
}
