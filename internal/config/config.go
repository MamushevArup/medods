package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server  `yaml:"server"`
	MongoDB `yaml:"mongodb"`
}

type Server struct {
	Port string `yaml:"port"`
}
type MongoDB struct {
	Port       string `yaml:"port"`
	Host       string `yaml:"host"`
	Database   string `yaml:"database"`
	Collection string `yaml:"collection"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig("config.yml", &cfg); err != nil {
		return nil, fmt.Errorf("error with reading config files %v", err)
	}
	return &cfg, nil
}
