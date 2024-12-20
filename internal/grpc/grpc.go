package grpc

import (
	"fmt"

	"github.com/lab-icn/water-potability-sensor-service/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(cfg *config.Config) (conn *grpc.ClientConn, err error) {
	return grpc.NewClient(
		fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}
