package grpc

import (
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient() (conn *grpc.ClientConn, err error) {
	return grpc.NewClient(
		os.Getenv("GRPC_ADDR"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}
