package config

import (
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/env"
)

const (
	defaultURL = "mongodb+srv://phucdt1280:dXM7ZrooZOSYWN2a@cluster0.6scequi.mongodb.net/"
)

type MongoDBConfig struct {
	URI string
	DBName string 
}

func (s *Config) mongoDBConfig() {
	s.MongoDB.URI = env.GetEnvAsString("MONGODB_URI", defaultURL)
	s.MongoDB.DBName = env.GetEnvAsString("MONGODB_DBNAME", "orderDB")
}
