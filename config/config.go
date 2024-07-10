package config

type Config struct {
	Debug bool
}

func NewConfig() *Config {
	return &Config{}
}
