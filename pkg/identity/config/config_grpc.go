package config

import "github.com/dothiphuc81299/coffeeShop-server/pkg/util/env"

const (
	defaultOrder = "localhost:50003"
)

type GRPCClientConfig struct {
	Order string
}

func (cfg *Config) grpcClientConfig() {
	cfg.GrpcClient.Order = env.GetEnvAsString("GRPC_ORDER", defaultOrder)
}
