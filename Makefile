api:
	@go run cmd/api/main.go

grpc:
	@go run cmd/grpc/server.go

generate-rpc:
	@protoc --proto_path=./proto ./proto/*.proto --go_out=. --go-grpc_out=.

compose/up:
	@docker compose --file deployment/testing.compose.yaml --env-file .env up --detach --no-deps

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
