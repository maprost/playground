package internal

import (
	"github.com/gorilla/mux"
	"net/http"
)

func initPing(router *mux.Router) {
	router.Path("/ping").Methods(http.MethodGet).HandlerFunc(ping)
}

func ping(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusNoContent)
}
