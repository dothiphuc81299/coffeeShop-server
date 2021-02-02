package config

// Config ...
type Config struct {
	ENV string `env:"ENV"`

	Mongo     MongoCfg `env:",prefix=MONGO_"`
	AdminPort string   `env:"ADMIN_PORT,required"`
}

// MongoCfg ...
type MongoCfg struct {
	URI  string `env:"URI,required"`
	Name string `env:"NAME,required"`
}

