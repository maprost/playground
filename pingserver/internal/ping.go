package internal

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func initPing(router *mux.Router) {
	router.Path("/").Methods(http.MethodGet).HandlerFunc(ping)
}

func ping(writer http.ResponseWriter, _ *http.Request) {
	fmt.Println("Ping")

	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("{\"blob\":\"drop\"}"))
}
