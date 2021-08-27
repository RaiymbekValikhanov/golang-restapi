package apiserver

import (
	"github.com/RaiymbekValikhanov/golang-restapi/internal/app/store"
)

type Config struct {
	BindAddr string `yaml:"bind_addr"`
	LogLevel string `yaml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
