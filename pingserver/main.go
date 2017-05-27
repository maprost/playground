package main

import (
	"github.com/maprost/playground/pingserver/internal"
	"github.com/maprost/playground/pingserver/internal/sys"
	"net/http"
)

func main() {

	sys.InitConfig()

	router := internal.InitServer()
	http.ListenAndServe(":"+sys.GetPort(), router)
}
