package main

import (
	"bytes"
	"crypto/aes"
	"encoding/json"
	stdlog "log"
	"math/rand"
	"os"
	"time"

	"github.com/lab-icn/water-potability-sensor-service/internal/aes256"
	"github.com/lab-icn/water-potability-sensor-service/internal/config"
	"github.com/lab-icn/water-potability-sensor-service/internal/domain"
	_mqtt "github.com/lab-icn/water-potability-sensor-service/internal/mqtt"
	"github.com/rs/zerolog"
	"go.uber.org/zap"
)

func main() {
	path := os.Getenv("CONFIG_FILEPATH")
	content, err := os.ReadFile(path)
	if err != nil {
		stdlog.Fatalf("opening config file at %s: %v\n", path, err)
	}
	cfg := new(config.Config)
	if err := json.Unmarshal(content, cfg); err != nil {
		stdlog.Fatalf("parsing config file content: %v\n", err)
	}

	log := zerolog.New(os.Stderr).With().Timestamp().Logger()

	logger, err := zap.NewProduction()
	if err != nil {
		stdlog.Fatalf("Failed to create logger instance: %v\n", err)
	}
	defer logger.Sync()

	mqttClient, err := _mqtt.NewClient(cfg, &log)
	if err != nil {
		stdlog.Fatalf("Failed to start MQTT connection: %v\n", err)
	}
	defer mqttClient.Disconnect(250)

	stdlog.Println("mqtt server is running...")
	for {
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(domain.WaterPotability{
			PH:                   4 + rand.Float64()*(10-4),
			TotalDissolvedSolids: 20 + rand.Float64()*(50-20),
			Turbidity:            5 + rand.Float64()*(10-5),
		}); err != nil {
			logger.Fatal("parsing map to json", zap.Error(err))
		}
		enc, err := aes256.Encrypt(buf.String(), cfg.AES.Key[:aes.BlockSize*2])
		if err != nil {
			logger.Fatal("encrypting json string", zap.Error(err))
		}
		buf.Reset()
		if _, err := buf.WriteString(enc); err != nil {
			logger.Fatal("writing ciphertext to buffer", zap.Error(err))
		}

		token := mqttClient.Publish("/wp", 1, false, buf.Bytes())
		token.Wait()
		stdlog.Println("message sent")
		time.Sleep(time.Second)
	}
}
