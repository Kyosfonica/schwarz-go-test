package config

import (
	"coupon_service/internal/api"
	"github.com/brumhard/alligotor"
)

type Config struct {
	API api.Config
}

func New() (Config, error) {
	cfg := Config{}
	if err := alligotor.Get(&cfg); err != nil {
		return Config{}, err
	}

	cfg.API.Port = 8080
	cfg.API.Host = "localhost"

	return cfg, nil
}

func (cfg *Config) SetDefaults() {
	cfg.API.Port = 8080
	cfg.API.Host = "localhost"
}

func (cfg *Config) SetPort(port int) {
	cfg.API.Port = port
}

func (cfg *Config) SetHost(host string) {
	cfg.API.Host = host
}
