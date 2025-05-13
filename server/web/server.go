package web

import (
	"log/slog"
	"net/http"

	"github.com/joyharbour/joyharbour-server/server/config"
)

func RunWebServer() error {
	HandleRoutes()

	config := config.GetConfig().Web

	slog.Info("start web server", "address", config.Address)
	err := http.ListenAndServe(config.Address, nil)

	slog.Error("web server exited", "err", err)
	return err
}
