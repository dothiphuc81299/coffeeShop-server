package config

type Config struct {
	MongoDB        MongoDBConfig
	Server          Server
}

func FromEnv() (*Config, error) {
	cfg := &Config{}
	cfg.mongoDBConfig()
	cfg.serverConfig()


	return cfg, nil
}
