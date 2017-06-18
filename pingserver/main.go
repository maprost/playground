package main

import (
	"fmt"
	"net/http"

	"github.com/maprost/playground/pingserver/internal"
	"github.com/maprost/playground/pingserver/internal/sys"
)

func main() {
	sys.InitConfig()

	router := internal.InitServer()
	fmt.Println("Port:" + sys.GetPort())
	http.ListenAndServe(":"+sys.GetPort(), router)
}
