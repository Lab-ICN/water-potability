package mqtt

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

func NewClient(logger *zap.Logger) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions().
		AddBroker(fmt.Sprintf("%s://%s:%s", os.Getenv("MQTT_PROTOCOL"), os.Getenv("MQTT_HOST"), os.Getenv("MQTT_PORT"))).
		// SetUsername(os.Getenv("MQTT_USERNAME")).
		// SetPassword(os.Getenv("MQTT_PASSWORD")).
		SetClientID(os.Getenv("MQTT_CLIENT_ID")).
		SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
			logger.Info("message received", zap.String("topic", msg.Topic()), zap.ByteString("payload", msg.Payload()))
		})
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return client, nil
}
