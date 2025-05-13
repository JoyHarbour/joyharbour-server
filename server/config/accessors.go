package config

import (
	"encoding/json"
	"log/slog"
	"os"
	"path"
)

var cachedConfig *Config = nil

func GetConfig() Config {
	if cachedConfig == nil {
		panic("GetConfig() called before LoadConfig()")
	}

	return *cachedConfig
}

func InitConfig() (err error) {
	filePath := path.Join("config", "config.json")
	slog.Info("init config")

	file, err := os.Open(filePath)
	if err != nil {
		return
	}

	config := Config{}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return
	}

	cachedConfig = &config
	return
}
