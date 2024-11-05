package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Port    string        `yaml:"port"`
	Host    string        `yaml:"host"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	var cfg Config

	configPath := fetchConfigPath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}
	return &cfg
}

func fetchConfigPath() string {
	defaultPath := "./config/default.yaml"

	v := os.Getenv("REDIS_CONFIG")
	if v != "" {
		return v
	} else {
		return defaultPath
	}
}
