package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	http.Handle("/", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world!")
}
