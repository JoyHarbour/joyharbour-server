package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRoutes() {
	r := mux.NewRouter()
	r.Handle("/", RootHandler())

	http.Handle("/", r)
}
