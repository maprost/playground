package main

import (
	"net/http"

	"github.com/maprost/playground/userserver/internal"
	"github.com/maprost/playground/userserver/internal/data"
	"github.com/maprost/playground/userserver/internal/sys"
)

func main() {

	sys.InitConfig()
	data.OpenDB()
	data.Register()

	router := internal.InitServer()

	http.ListenAndServe(":"+sys.GetPort(), router)
}
