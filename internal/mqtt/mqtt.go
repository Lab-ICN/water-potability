package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/lab-icn/water-potability-sensor-service/internal/config"
	mqttDelivery "github.com/lab-icn/water-potability-sensor-service/internal/interface/mqtt"
	"github.com/rs/zerolog"
)

func Listen(
	subscriber mqttDelivery.IMqttSubscriber,
	cfg *config.Config,
	log *zerolog.Logger,
) mqtt.Client {
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
		SetDefaultPublishHandler(func(_ mqtt.Client, _ mqtt.Message) {}).
		SetConnectionLostHandler(func(_ mqtt.Client, _ error) {
			log.Warn().Msg("mqtt connection lost")
		}).
		SetReconnectingHandler(func(_ mqtt.Client, _ *mqtt.ClientOptions) {
			log.Info().Msg("mqtt reconnecting")
		}).
		SetOnConnectHandler(func(client mqtt.Client) {
			log.Info().Msg("mqtt connected")
			token := client.Subscribe(cfg.MQTT.SensorTopic, cfg.MQTT.QOS, subscriber.SensorSubscriber)
			<-token.Done()
			if err := token.Error(); err != nil {
				log.Error().Err(err).Msg("attempting to subscribe")
			}
		})

	return mqtt.NewClient(opts)
}
