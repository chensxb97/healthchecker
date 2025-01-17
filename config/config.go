package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Endpoint struct {
	URL string `yaml:"url"`
}

type Config struct {
	Endpoints []Endpoint `yaml:"endpoints"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
