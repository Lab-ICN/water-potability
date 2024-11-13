api:
	@go run cmd/api/main.go

grpc:
	@go run cmd/grpc/server.go

generate-rpc:
	@protoc --proto_path=./proto ./proto/*.proto --go_out=. --go-grpc_out=.

compose/up:
	@docker compose --file deployment/dev.compose.yaml --env-file .env up --detach --no-deps

compose/down:
	@docker compose --file deployment/dev.compose.yaml --env-file .env down


.PHONY:
	api
	grpc
	generate-rpc
	compose/up
	compose/down
