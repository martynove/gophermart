package config

import (
	"errors"
	"flag"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddress        string `env:"RUN_ADDRESS"`
	DatabaseURI          string `env:"DATABASE_URI"`
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	DebugMode            bool   `env:"GOPHER_DEBUG" envDefault:"False"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return &cfg, err
	}
	loadFromFlag(&cfg)
	switch {
	case cfg.ServerAddress == "":
		return &cfg, errors.New("doesn't set api api server address (env - RUN_ADDRESS or flag -a)")
	case cfg.DatabaseURI == "":
		return &cfg, errors.New("doesn't set database uri (env - DATABASE_URI or flag -d)")
	case cfg.AccrualSystemAddress == "":
		return &cfg, errors.New("doesn't set accrual system address (env - ACCRUAL_SYSTEM_ADDRESS or flag -r)")
	}
	return &cfg, nil
}

func loadFromFlag(cfg *Config) {
	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "api server address")
	flag.StringVar(&cfg.DatabaseURI, "d", cfg.DatabaseURI, "database uri")
	flag.StringVar(&cfg.AccrualSystemAddress, "r", cfg.AccrualSystemAddress, "accrual system address")
	flag.BoolVar(&cfg.DebugMode, "debug", cfg.DebugMode, "flag for debug mode")
	flag.Parse()
}
