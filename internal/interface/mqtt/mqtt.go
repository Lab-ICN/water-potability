package mqtt

import (
	"context"
	"encoding/json"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/lab-icn/water-potability-sensor-service/internal/aes256"
	"github.com/lab-icn/water-potability-sensor-service/internal/config"
	"github.com/lab-icn/water-potability-sensor-service/internal/domain"
	"github.com/lab-icn/water-potability-sensor-service/internal/service"
	"github.com/rs/zerolog"
)

type subscriber struct {
	service service.WaterPotabilityServiceItf
	cfg     *config.AES
	log     *zerolog.Logger
}

type IMqttSubscriber interface {
	SensorSubscriber(client mqtt.Client, msg mqtt.Message)
}

func NewMqttSubscriber(
	service service.WaterPotabilityServiceItf,
	cfg *config.Config,
	log *zerolog.Logger,
) IMqttSubscriber {
	return &subscriber{service, &cfg.AES, log}
}

func (s *subscriber) SensorSubscriber(client mqtt.Client, msg mqtt.Message) {
	s.log.Debug().
		Str("topic", msg.Topic()).
		Msg(string(msg.Payload()))

	jsonstr, err := aes256.Decrypt(
		string(msg.Payload()),
		[]byte(s.cfg.Key),
		[]byte(s.cfg.IV),
	)
	if err != nil {
		s.log.Err(err).
			Bytes("payload", msg.Payload()).
			Msg("decrypting mqtt payload aes cipher")
		return
	}
	s.log.Debug().Msg(jsonstr)

	var potability domain.WaterPotability
	if err := json.Unmarshal([]byte(jsonstr), &potability); err != nil {
		s.log.Err(err).
			Str("payload", jsonstr).
			Msg("decoding mqtt json payload string to struct")
		return
	}

	potability.Node = msg.Topic()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := s.service.PredictWaterPotability(ctx, potability); err != nil {
		s.log.Err(err).Msg("predict water potability")
		return
	}
}
