run:
	@go run cmd/api/main.go

generate-rpc:
	@protoc --proto_path=./proto ./proto/*.proto --go_out=. --go-grpc_out=.