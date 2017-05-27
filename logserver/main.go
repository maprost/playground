package main

import (
	"github.com/maprost/playground/userserver/core/data"
	"github.com/maprost/playground/userserver/core/sys"
	"net/http"
)

func main() {

	sys.InitConfig()

	router := core.InitServer()
	http.ListenAndServe(":"+sys.GetPort(), router)
}
