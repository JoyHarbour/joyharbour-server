package main

import (
	"log/slog"
	"os"

	"github.com/joyharbour/joyharbour-server/server/config"
	"github.com/joyharbour/joyharbour-server/server/database"
	"github.com/joyharbour/joyharbour-server/server/storage"
	"github.com/joyharbour/joyharbour-server/server/web"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	err := changeDirectoryToHome()
	if err != nil {
		slog.Error("failed change directory to home", "err", err)
		os.Exit(1)
	}

	err = initialSetup()
	if err != nil {
		os.Exit(1)
	}

	err = web.RunWebServer()
	if err != nil {
		os.Exit(1)
	}
}

func changeDirectoryToHome() (err error) {
	homePath, ok := os.LookupEnv("JOYHARBOUR_HOME")
	if !ok {
		homePath = "."
	}

	slog.Info("changing directory to home", "path", homePath)
	err = os.Chdir(homePath)
	if err != nil {
		return
	}

	return
}

func initialSetup() (err error) {
	err = config.InitConfig()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		return err
	}

	err = storage.InitStorage()
	if err != nil {
		slog.Error("failed to load storage", "error", err)
		return err
	}

	err = database.InitDatabase()
	if err != nil {
		slog.Error("failed to load database", "error", err)
		return err
	}

	return
}
