package mqtt

import (
	"context"
	"crypto/aes"
	"encoding/json"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/lab-icn/water-potability-sensor-service/internal/aes256"
	"github.com/lab-icn/water-potability-sensor-service/internal/config"
	"github.com/lab-icn/water-potability-sensor-service/internal/domain"
	"github.com/lab-icn/water-potability-sensor-service/internal/service"
	"github.com/rs/zerolog"
)

type handler struct {
	service service.WaterPotabilityServiceItf
	cfg     *config.AES
	log     *zerolog.Logger
}

func NewMqttHandler(
	client mqtt.Client,
	cfg *config.Config,
	log *zerolog.Logger,
	service service.WaterPotabilityServiceItf,
) {
	handler := &handler{service, &cfg.AES, log}
	token := client.Subscribe(cfg.MQTT.SensorTopic, 1, handler.sensorSubscriber)
	<-token.Done()
	if token.Error() != nil {
		log.Err(token.Error()).Msg("subscribing to mqtt topic")
	}
}

func (h *handler) sensorSubscriber(client mqtt.Client, msg mqtt.Message) {
	h.log.Debug().
		Str("topic", msg.Topic()).
		Msg(string(msg.Payload()))

	jsonstr, err := aes256.Decrypt(
		string(msg.Payload()),
		h.cfg.Key[:aes.BlockSize*2],
	)
	if err != nil {
		h.log.Err(err).
			Bytes("payload", msg.Payload()).
			Msg("decrypting mqtt payload aes cipher")
		return
	}
	h.log.Debug().Msg(jsonstr)

	var potability domain.WaterPotability
	if err := json.Unmarshal([]byte(jsonstr), &potability); err != nil {
		h.log.Err(err).
			Str("payload", jsonstr).
			Msg("decoding mqtt json payload string to struct")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := h.service.PredictWaterPotability(ctx, potability); err != nil {
		h.log.Err(err).Msg("predict water potability")
		return
	}
}
