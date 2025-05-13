package storage

import (
	"errors"
	"log/slog"
	"os"

	"github.com/joyharbour/joyharbour-server/server/config"
)

func InitStorage() (err error) {
	config := config.GetConfig().Storage
	slog.Info("init storage", "type", config.Type)

	if config.Type == "local" {
		storage, err = loadStorageLocal()
		if err != nil {
			return
		}
	}

	return
}

func loadStorageLocal() (Storage, error) {
	path := "storage"

	config := config.GetConfig().Storage.Options
	if config != nil {
		configPath, ok := config["path"]
		if ok {
			path = configPath
		}
	}

	slog.Info("init local storage", "path", path)

	pathStat, err := os.Stat(path)

	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(path, 0750)
		if err != nil {
			return nil, err
		}

		slog.Debug("created new storage directory")
	} else if err != nil {
		return nil, err
	} else if !pathStat.IsDir() {
		return nil, errors.New("storage path is not a directory")
	}

	return &LocalStorage{Path: path}, nil
}
