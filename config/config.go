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

var cfg Config

func LoadConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return err
	}
	return nil
}

func GetEndpoints() []string {
	urls := make([]string, len(cfg.Endpoints))
	for i, endpoint := range cfg.Endpoints {
		urls[i] = endpoint.URL
	}
	return urls
}
