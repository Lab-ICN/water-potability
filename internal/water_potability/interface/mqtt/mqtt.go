package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/lab-icn/water-potability-sensor-service/internal/domain"
	"github.com/lab-icn/water-potability-sensor-service/internal/water_potability/service"
)

type MQTT struct {
	service service.WaterPotabilityServiceItf
}

func (m *MQTT) messagePubHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s with id %d from topic: %s\n", string(msg.Payload()), msg.MessageID(), msg.Topic())
	wp, err := parseWaterPotabilityJsonData(msg.Payload())
	if err != nil {
		log.Printf("Error parsing message: %v", err)
		return
	}

	err = m.service.PredictWaterPotability(context.Background(), wp)
	if err != nil {
		log.Printf("Error predicting water potability data: %v", err)
		return
	}
}

func (m *MQTT) connectLostHandler(client mqtt.Client, err error) {
	log.Printf("Connection lost: %v", err)
}

func NewMQTT(waterPotability service.WaterPotabilityServiceItf) {
	m := &MQTT{
		service: waterPotability,
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s://%s:%s", os.Getenv("MQTT_PROTOCOL"), os.Getenv("MQTT_HOST"), os.Getenv("MQTT_PORT")))
	opts.SetUsername(os.Getenv("MQTT_USERNAME"))
	opts.SetPassword(os.Getenv("MQTT_PASSWORD"))
	opts.SetClientID(os.Getenv("MQTT_CLIENT_ID"))
	opts.OnConnectionLost = m.connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Printf("Error connecting: %v", token.Error())
	}

	topic := os.Getenv("MQTT_TOPIC")
	qosStr := os.Getenv("MQTT_QOS")
	qos, err := strconv.Atoi(qosStr)
	if err != nil {
		log.Fatalf("Error converting QoS: %v", err)
	}

	if token := client.Subscribe(topic, byte(qos), m.messagePubHandler); token.Wait() && token.Error() != nil {
		log.Fatalf("Error subscribing to topic: %v", token.Error())
	}

	log.Printf("Subscribed to topic: %s\n", topic)

	publish(client)

	client.Disconnect(250)
}

func publish(client mqtt.Client) {
	num := 100000
	topic := os.Getenv("MQTT_TOPIC")
	qosStr := os.Getenv("MQTT_QOS")
	qos, err := strconv.Atoi(qosStr)
	if err != nil {
		log.Fatalf("Error converting QoS: %v", err)
	}

	for i := 0; i < num; i++ {
		waterPotability := domain.WaterPotability{
			PH:                   7.5,
			Turbidity:            5.5,
			TotalDissolvedSolids: 100,
		}

		data, err := json.Marshal(waterPotability)
		if err != nil {
			log.Fatalf("Error marshalling data: %v", err)
			return
		}

		token := client.Publish(topic, byte(qos), false, string(data))
		token.Wait()
		time.Sleep(1 * time.Second)
	}
}
