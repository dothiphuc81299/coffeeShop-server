package config

var cfg Config

// InitENV ...
func InitENV() {
	cfg = Config{
		IsDev: true,
	}
}

// Getcfg ...
func Getcfg() *Config {
	return &cfg
}
