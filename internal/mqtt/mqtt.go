package mqtt

import (
	"fmt"
	"time"

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
		SetClientID(cfg.MQTT.ClientID).
		SetUsername(cfg.MQTT.Username).
		SetPassword(cfg.MQTT.Password).
		SetKeepAlive(15 * time.Second).
		SetPingTimeout(10 * time.Second).
		SetAutoReconnect(true).
		SetCleanSession(false).
		SetDefaultPublishHandler(func(_ mqtt.Client, _ mqtt.Message) {}).
		SetConnectionLostHandler(func(_ mqtt.Client, _ error) {
			log.Warn().Msg("mqtt connection lost")
		}).
		SetReconnectingHandler(func(_ mqtt.Client, _ *mqtt.ClientOptions) {
			log.Info().Msg("mqtt reconnecting")
		}).
		SetOnConnectHandler(func(_ mqtt.Client) {
			log.Info().Msg("mqtt connected")
		})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return client, nil
}
