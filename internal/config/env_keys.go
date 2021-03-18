package config

// Config ...
type Config struct {
	ENV string `env:"ENV"`

	Mongo      MongoCfg `env:",prefix=MONGO_"`
	AuthSecret string   `env:"AUTH_SECRET,required"`
	AdminPort  string   `env:"ADMIN_PORT,required"`

	FileHost string `env:"FILE_HOST,required"`
}

// MongoCfg ...
type MongoCfg struct {
	URI  string `env:"URI,required"`
	Name string `env:"NAME,required"`
}
