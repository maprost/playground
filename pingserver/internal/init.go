package internal

import "github.com/gorilla/mux"

func InitServer() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	initPing(router)

	return router
}
