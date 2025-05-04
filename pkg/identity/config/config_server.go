package config

import "github.com/dothiphuc81299/coffeeShop-server/pkg/util/env"

const (
	defaultHTTPPort = "3089"
	defaultGRPCPort = "50003"
)

type Server struct {
	Host           string
	HTTPPort       string
	GRPCPort       string
	StaticRootPath string
}

func (cfg *Config) serverConfig() {
	cfg.Server.HTTPPort = env.GetEnvAsString("IDENTITY_HTTP_PORT", defaultHTTPPort)
	cfg.Server.GRPCPort = env.GetEnvAsString("IDENTITY_GRPC_PORT", defaultGRPCPort)
}
