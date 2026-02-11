package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Timeout                time.Duration `yaml:"timeout"`
	ServiceAccountFilePath string        `yaml:"googleServiceAccountFilePath"`
}

func Load(path string) (*Config, error) {

	cfgFile, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to load config: %w", err)
	}

	var cfg Config
	err = yaml.Unmarshal(cfgFile, &cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)
	}

	return &cfg, nil
}
