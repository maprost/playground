package main

import (
	"github.com/maprost/playground/userserver/core/sys"
	"github.com/maprost/playground/userserver/core/data"
	"net/http"
)

func main() {

	sys.InitConfig()

	router := core.InitServer()
	http.ListenAndServe(":"+sys.GetPort(), router)
}
