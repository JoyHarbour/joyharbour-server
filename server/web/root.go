package web

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func RootHandler() http.Handler {
	wwwStat, err := os.Stat("www")

	if errors.Is(err, os.ErrNotExist) {
		slog.Warn("www directory not found - using fallback handler")
		return fallbackRootHandler{reason: "www directory not found"}
	}

	if err != nil {
		slog.Warn("error checking www directory - using fallback handler", "err", err)
		return fallbackRootHandler{reason: "error checking www directory"}
	}

	if !wwwStat.IsDir() {
		slog.Warn("www path is not a directory - using fallback handler")
		return fallbackRootHandler{reason: "www path is not a directory"}
	}

	return http.FileServer(http.Dir("www"))
}

//
// fallback handler when there are no www files available
//

type fallbackRootHandler struct {
	reason string
}

func (handler fallbackRootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Server started without frontend: %v", handler.reason)
}
