package config

type Config struct {
	MongoDB    MongoDBConfig
	Server     Server
	GrpcClient GRPCClientConfig
}

func FromEnv() (*Config, error) {
	cfg := &Config{}
	cfg.mongoDBConfig()
	cfg.serverConfig()
	cfg.grpcClientConfig()

	return cfg, nil
}
