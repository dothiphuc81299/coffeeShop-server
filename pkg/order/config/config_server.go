package config

import "github.com/dothiphuc81299/coffeeShop-server/pkg/util/env"

const (
	defaultHTTPPort = "3089"
	defaultGRPCPort = "50004"
)

type Server struct {
	Host     string
	HTTPPort string
	GRPCPort string
}

func (cfg *Config) serverConfig() {
	cfg.Server.HTTPPort = env.GetEnvAsString("ORDER_HTTP_PORT", defaultHTTPPort)
	cfg.Server.GRPCPort = env.GetEnvAsString("ORDER_GRPC_PORT", defaultGRPCPort)
}
