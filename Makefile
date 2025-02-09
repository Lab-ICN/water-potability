api:
	@CONFIG_FILEPATH=config.json go run cmd/subsciber/main.go

api-windows:
	@set CONFIG_FILEPATH=config.json && go run cmd/subscriber/main.go

grpc:
	@go run cmd/grpc/server.go

generate-rpc:
	@protoc --proto_path=./proto ./proto/*.proto --go_out=. --go-grpc_out=.

compose/up:
	@docker compose --file deployment/compose.yaml --env-file deployment/.env up --detach --no-deps

compose/down:
	@docker compose --file deployment/dev.compose.yaml down

compose/testing:
	@docker compose \
		--file deployment/dev.compose.yaml \
		--file deployment/testing.compose.yaml \
		--env-file .env \
		up --detach --no-deps

compose/testing.down:
	@docker compose \
		--file deployment/dev.compose.yaml \
		--file deployment/testing.compose.yaml \
		--env-file .env \
		down --detach --no-deps

.PHONY: api grpc generate-rpc compose/up compose/down compose/testing
.PHONY: compose/testing.down
