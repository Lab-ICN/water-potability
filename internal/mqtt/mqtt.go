package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/lab-icn/water-potability-sensor-service/internal/config"
	"github.com/rs/zerolog"
)

func NewClient(cfg *config.Config, log *zerolog.Logger) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions().
		AddBroker(fmt.Sprintf(
			"%s://%s:%d",
			cfg.MQTT.Protocol,
			cfg.MQTT.Host,
			cfg.MQTT.Port,
		)).
		SetUsername(cfg.MQTT.Username).
		SetPassword(cfg.MQTT.Password).
		SetClientID(cfg.MQTT.ClientID).
		SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {})
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return client, nil
}
