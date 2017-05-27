package core

import "github.com/gorilla/mux"

func InitServer() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	initUser(router)

	return router
}
