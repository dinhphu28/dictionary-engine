package config

import (
	"errors"
	"strings"

	"github.com/BurntSushi/toml"
)

var cfg Config

func LoadConfig(path string) error {
	if strings.HasSuffix(path, ".toml") {
		return loadTomlConfig(path)
	}

	return errors.New("config type not support")
}

func loadTomlConfig(path string) error {
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return err
	}

	if len(cfg.Priority) == 0 {
		return errors.New("priority list is empty")
	}
	if cfg.Paths.Resources == "" {
		return errors.New("paths.resources is required")
	}

	return nil
}

func GetConfig() Config {
	return cfg
}
