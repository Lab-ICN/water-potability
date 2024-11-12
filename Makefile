api:
	@go run cmd/api/main.go

grpc:
	@go run cmd/grpc/server.go

generate-rpc:
	@protoc --proto_path=./proto ./proto/*.proto --go_out=. --go-grpc_out=.

compose/up:
	@docker compose --file testing.compose.yaml up --detach --no-deps

compose/fresh:
	@docker compose --file testing.compose.yaml up --detach --no-deps --build

compose/down:
	@docker compose --file testing.compose.yaml down

compose/ps:
	@docker compose --file testing.compose.yaml ps --all

.PHONY:
	api
	grpc
	generate-rpc
	compose/up
	compose/fresh
	compose/down
	compose/ps
