include .env

api:
	@go run cmd/api/main.go

grpc:
	@go run cmd/grpc/server.go

generate-rpc:
	@protoc --proto_path=./proto ./proto/*.proto --go_out=. --go-grpc_out=.

compose/up:
	@docker compose --file deployment/testing.compose.yaml --env-file .env up --detach --no-deps
	@docker exec --user root waterpotability-mosquitto-1 mosquitto_passwd -b -c /mosquitto/config/passwd ${MQTT_USERNAME} ${MQTT_PASSWORD}
	@docker exec --user root waterpotability-mosquitto-1 chown mosquitto:mosquitto /mosquitto/config/passwd

compose/fresh:
	@docker compose --file deployment/testing.compose.yaml --env-file .env up --detach --no-deps --build

compose/down:
	@docker compose --file deployment/testing.compose.yaml --env-file .env down

compose/ps:
	@docker compose --file deployment/testing.compose.yaml --env-file .env ps --all

.PHONY:
	api
	grpc
	generate-rpc
	compose/up
	compose/fresh
	compose/down
	compose/ps
