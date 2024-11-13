package mqtt

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/lab-icn/water-potability-sensor-service/internal/domain"
	"github.com/lab-icn/water-potability-sensor-service/internal/water_potability/service"
)

type handler struct {
	service service.WaterPotabilityServiceItf
}

func NewMqttHandler(client mqtt.Client, service service.WaterPotabilityServiceItf) {
	handler := &handler{service}
	if token := client.Subscribe("wp", byte(1), handler.sensorSubscriber); token.Wait() && token.Error() != nil {
		log.Fatalf("Error subscribing to topic: %v", token.Error())
	}
}

func (h *handler) sensorSubscriber(client mqtt.Client, msg mqtt.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	buf := new(bytes.Buffer)
	if _, err := buf.Write(msg.Payload()); err != nil {
		log.Printf("Error writing payload to buffer: %v", err)
		return
	}
	var potability domain.WaterPotability
	if err := json.NewDecoder(buf).Decode(&potability); err != nil {
		log.Printf("Error decode message payload: %v", err)
		return
	}
	if err := h.service.PredictWaterPotability(ctx, potability); err != nil {
		log.Printf("Error predicting water potability data: %v", err)
		return
	}
}
