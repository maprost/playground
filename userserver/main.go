package main

import (
	"net/http"

	"github.com/maprost/playground/userserver/core"
	"github.com/maprost/playground/userserver/core/data"
	"github.com/maprost/playground/userserver/core/sys"
)

func main() {

	sys.InitConfig()
	data.OpenDB()
	data.Register()

	router := core.InitServer()
	http.ListenAndServe(":"+sys.GetPort(), router)
}
