package mqtt

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttHandler struct {
	client mqtt.Client
}

func NewMqttHandler(broker string, port int, clientID string) *MqttHandler {
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(clientID)
	opts.OnConnect = func(c mqtt.Client) {
		fmt.Println("Connected to MQTT broker")
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to MQTT broker: %v", token.Error())
	}

	return &MqttHandler{
		client: client,
	}
}

func (h *MqttHandler) Subscribe(topic string, callback mqtt.MessageHandler) {
	token := h.client.Subscribe(topic, 1, callback)
	token.Wait()
	if token.Error() != nil {
		log.Fatalf("Error subscribing to topic %s: %v", topic, token.Error())
	}
}

func (h *MqttHandler) Close() {
	h.client.Disconnect(250)
	fmt.Println("Disconnected from MQTT broker")
}
