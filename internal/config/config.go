package config

type Config struct {
	AES      AES
	InfluxDB InfluxDB
	GRPC     GRPC
	MQTT     MQTT
}

type GRPC struct {
	Protocol string
	Host     string
	Port     uint
}

type InfluxDB struct {
	Protocol string
	Host     string
	Token    string
	Org      string
	Bucket   string
	Port     uint
}

type MQTT struct {
	Protocol     string
	Host         string
	Username     string
	Password     string
	ClientID     string
	SensorTopics []string
	Port         uint
	QOS          byte
}

type AES struct {
	Key string
	IV  string
}
